package queries

const RESOURCE_SALES_HISTORY_MOST_POPULAR_QUERY = `
select distinct sh.resource_id, r.title, sh.teacher_id, count(*) as sales_count from sales_history as sh
left join resources as r on r.id = sh.resource_id 
where sh.teacher_id = ?
GROUP BY sh.resource_id, r.title, sh.teacher_id
ORDER BY sales_count DESC
LIMIT ?;
`
