package queries

const RESOURCE_SALES_HISTORY_BY_TEACHER_QUERY = `select distinct sh.resource_id, sh.teacher_id, r.title, sum(amount) as amount 
from public.sales_history sh 
left join public.resources as r on r.id = sh.resource_id 
where sh.teacher_id = ? 
group by sh.resource_id, sh.teacher_id, r.title`
