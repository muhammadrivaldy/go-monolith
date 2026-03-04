ALTER TABLE `mst_access`
  DROP FOREIGN KEY `fk_mst_access_api_id`,
  DROP FOREIGN KEY `fk_mst_access_user_type_id`;

ALTER TABLE `mst_user`
  DROP FOREIGN KEY `fk_mst_user_user_type_id`,
  DROP FOREIGN KEY `fk_mst_user_status`;

ALTER TABLE `mst_api`
  DROP FOREIGN KEY `fk_mst_api_service_id`;

ALTER TABLE `mst_access`
  DROP INDEX `idx_mst_access_api_id`;

ALTER TABLE `mst_access`
  MODIFY `user_type_id` INT(2) NOT NULL;

ALTER TABLE `mst_user`
  MODIFY `status` INT(2) NOT NULL,
  MODIFY `user_type_id` INT(2) NOT NULL;
