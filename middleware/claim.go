package middleware

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/rmsubekti/sporagium/helper"

	"github.com/golang-jwt/jwt/v4"
)

type Principal struct {
	ID         string  `json:"id"`
	Name       string  `json:"name"`
	ExpireDays float32 `json:"expire_days,omitempty"`
	Token      string  `json:"token,omitempty"`
}

type customClaims struct {
	jwt.RegisteredClaims
	Principal Principal
}

func (p *Principal) CreateToken() (err error) {
	var jwtToken string
	mySigningKey := []byte(helper.GetEnv("SPORAGIUM_JWT_SECRET_KEY", "5up3rSP0r4G"))
	// Create the Claim
	claim := customClaims{
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add((time.Hour * 24) * time.Duration(p.ExpireDays))),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			ID:        p.ID,
			Issuer:    p.Name,
		},
		*p,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	if jwtToken, err = token.SignedString(mySigningKey); err != nil {
		return
	}
	p.Token = jwtToken
	return
}

func (p *Principal) Parse() error {

	if len(p.Token) < 1 {
		return errors.New("no token provided")
	}

	token := strings.SplitN(p.Token, " ", 2)

	if (len(token) < 2) || (token[0] != "Bearer") {
		return errors.New("incorrect format authorization header")
	}

	key, err := jwt.ParseWithClaims(token[1], &customClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(helper.GetEnv("SPORAGIUM_JWT_SECRET_KEY", "5up3rSP0r4G")), nil
	})

	if !key.Valid && err != nil {
		return err
	}

	if claim, ok := key.Claims.(*customClaims); ok {
		*p = claim.Principal
		return nil
	}

	return fmt.Errorf("invalid token %s ", err.Error())
}
