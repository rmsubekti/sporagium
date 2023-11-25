---- empty for now ----
create or replace view spora.secrets as 
select s2.id::text as id,s.secret,
s2."data", s2.callback_url as domain
from spora.secret s 
join spora.spora s2 on s.spora_id = s2.id 