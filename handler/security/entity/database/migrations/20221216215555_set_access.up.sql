INSERT INTO mst_access (id, user_type_id, api_id, created_at, updated_at) SELECT 1, 1, id, now(), now() FROM mst_api WHERE `name` = 'Refresh JWT';
INSERT INTO mst_access (id, user_type_id, api_id, created_at, updated_at) SELECT 2, 1, id, now(), now() FROM mst_api WHERE `name` = 'Create User';
INSERT INTO mst_access (id, user_type_id, api_id, created_at, updated_at) SELECT 3, 1, id, now(), now() FROM mst_api WHERE `name` = 'Get User By Id';
INSERT INTO mst_access (id, user_type_id, api_id, created_at, updated_at) SELECT 4, 1, id, now(), now() FROM mst_api WHERE `name` = 'Edit User';
INSERT INTO mst_access (id, user_type_id, api_id, created_at, updated_at) SELECT 5, 1, id, now(), now() FROM mst_api WHERE `name` = 'Edit Password User';
INSERT INTO mst_access (id, user_type_id, api_id, created_at, updated_at) SELECT 6, 1, id, now(), now() FROM mst_api WHERE `name` = 'Remove User';
INSERT INTO mst_access (id, user_type_id, api_id, created_at, updated_at) SELECT 7, 1, id, now(), now() FROM mst_api WHERE `name` = 'Get Order';
INSERT INTO mst_access (id, user_type_id, api_id, created_at, updated_at) SELECT 8, 2, id, now(), now() FROM mst_api WHERE `name` = 'Create Order';
INSERT INTO mst_access (id, user_type_id, api_id, created_at, updated_at) SELECT 9, 3, id, now(), now() FROM mst_api WHERE `name` = 'Create Order';
INSERT INTO mst_access (id, user_type_id, api_id, created_at, updated_at) SELECT 10, 2, id, now(), now() FROM mst_api WHERE `name` = 'Refresh JWT';
INSERT INTO mst_access (id, user_type_id, api_id, created_at, updated_at) SELECT 11, 3, id, now(), now() FROM mst_api WHERE `name` = 'Refresh JWT';