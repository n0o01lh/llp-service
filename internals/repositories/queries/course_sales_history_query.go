package queries

const COURSE_SALES_HISTORY_QUERY = "select distinct rc.course_id, tc.title, sh.teacher_id, sum(sh.amount) as amount from public.resources_courses as rc left join public.sales_history as sh on rc.resource_id = sh.resource_id  left join public.courses as tc on tc.id = rc.course_id where sh.teacher_id = ? group by rc.course_id, tc.title, sh.teacher_id;"
