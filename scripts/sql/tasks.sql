-- name: GetAllTasks :many
SELECT
    tasks.id as id,
    tasks.uuid as uuid,
    tasks.name as name,
    tasks.description as description,
    tasks.status as status,
    tasks.due_date as due_date,
    tasks.order as task_order,
    categories.name as category_name
FROM tasks
LEFT JOIN categories
    ON tasks.category_id = categories.id
WHERE
    tasks.deleted_at IS NULL
ORDER BY task_order ASC;
