INSERT INTO mst_api (`name`, `endpoint`, `method`, `service_id`, `created_at`, `updated_at`)
VALUES ('Refresh JWT', '/refresh', 'POST', 1, now(), now());

INSERT INTO mst_access (user_type_id, api_id, created_by, created_at) 
SELECT 1, id, 0, now() FROM mst_api
WHERE `name` in ('Refresh JWT');

INSERT INTO mst_access (user_type_id, api_id, created_by, created_at) 
SELECT 2, id, 0, now() FROM mst_api
WHERE `name` in ('Refresh JWT');

INSERT INTO mst_access (user_type_id, api_id, created_by, created_at) 
SELECT 3, id, 0, now() FROM mst_api
WHERE `name` in ('Refresh JWT');