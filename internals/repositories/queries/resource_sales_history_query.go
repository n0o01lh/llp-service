package queries

const RESOURCE_SALES_HISTORY_QUERY = "select distinct sh.resource_id, sh.teacher_id, r.title, sum(amount) as amount from public.sales_history sh left join public.resources as r on r.id = resource_id where resource_id = ? group by sh.resource_id, sh.teacher_id, r.title"
