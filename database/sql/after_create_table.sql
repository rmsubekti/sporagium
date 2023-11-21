---- empty for now ----
create or replace view spora.secrets as select c.id,c.secret,c."domain",c."data"  from spora.client c 