ALTER TABLE mst_api DROP PRIMARY KEY, CHANGE id id VARCHAR(36), ADD PRIMARY KEY (id);

ALTER TABLE mst_access DROP CONSTRAINT usertype_api_UNIQUE, MODIFY api_id VARCHAR(36), ADD UNIQUE usertype_api_UNIQUE (user_type_id, api_id);

DELETE FROM mst_access;

DELETE FROM mst_api;