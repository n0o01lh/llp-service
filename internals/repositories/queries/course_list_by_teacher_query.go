package queries

const COURSE_LIST_BY_TEACHER_QUERY = `select r.title, r.created_at, r.price, r.type, cs.id, rc.order from resources_courses rc
left join resources r on r.id=rc.resource_id
left join courses cs on cs.id=rc.course_id
where cs.teacher_id=?;`
