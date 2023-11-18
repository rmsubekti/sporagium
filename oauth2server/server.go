package oauth2server

import (
	"context"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/go-oauth2/oauth2/v4/errors"
	"github.com/go-oauth2/oauth2/v4/manage"
	"github.com/go-oauth2/oauth2/v4/server"
	session "github.com/go-session/session/v3"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v4"
	"github.com/rmsubekti/sporagium/database"
	pg "github.com/vgarvardt/go-oauth2-pg/v4"
	"github.com/vgarvardt/go-pg-adapter/pgx4adapter"
)

func Setup(r *mux.Router) {
	_, dbUri := database.GetDBConnectionString()
	pgxConn, _ := pgx.Connect(context.TODO(), dbUri)

	manager := manage.NewDefaultManager()
	manager.SetAuthorizeCodeTokenCfg(manage.DefaultAuthorizeCodeTokenCfg)

	// use PostgreSQL token store with pgx.Connection adapter
	adapter := pgx4adapter.NewConn(pgxConn)
	tokenStore, _ := pg.NewTokenStore(adapter, pg.WithTokenStoreTableName(`token`), pg.WithTokenStoreGCInterval(time.Minute))
	defer tokenStore.Close()
	clientStore, _ := pg.NewClientStore(adapter, pg.WithClientStoreTableName(`'spora'.client`), pg.WithClientStoreInitTableDisabled())

	manager.MapTokenStorage(tokenStore)
	manager.MapClientStorage(clientStore)

	srv := server.NewDefaultServer(manager)
	srv.SetAllowGetAccessRequest(true)
	srv.SetClientInfoHandler(server.ClientFormHandler)

	srv.SetInternalErrorHandler(func(err error) (re *errors.Response) {
		log.Println("Internal Error:", err.Error())
		return
	})

	srv.SetResponseErrorHandler(func(re *errors.Response) {
		log.Println("Response Error:", re.Error.Error())
	})

	oauth := r.PathPrefix("/o").Subrouter()
	oauth.HandleFunc("/authorize", func(w http.ResponseWriter, r *http.Request) {
		store, err := session.Start(r.Context(), w, r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		var form url.Values
		if v, ok := store.Get("ReturnUri"); ok {
			form = v.(url.Values)
		}
		r.Form = form

		store.Delete("ReturnUri")
		store.Save()

		err = srv.HandleAuthorizeRequest(w, r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	})

	oauth.HandleFunc("/token", func(w http.ResponseWriter, r *http.Request) {
		srv.HandleTokenRequest(w, r)
	})
}
