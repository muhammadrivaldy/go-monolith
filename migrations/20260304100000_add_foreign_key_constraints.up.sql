ALTER TABLE `mst_user`
  MODIFY `status` BIGINT NOT NULL,
  MODIFY `user_type_id` BIGINT NOT NULL;

ALTER TABLE `mst_access`
  MODIFY `user_type_id` BIGINT NOT NULL;

ALTER TABLE `mst_access`
  ADD INDEX `idx_mst_access_api_id` (`api_id`);

ALTER TABLE `mst_api`
  ADD CONSTRAINT `fk_mst_api_service_id`
  FOREIGN KEY (`service_id`) REFERENCES `mst_service` (`id`)
  ON UPDATE CASCADE
  ON DELETE RESTRICT;

ALTER TABLE `mst_user`
  ADD CONSTRAINT `fk_mst_user_status`
  FOREIGN KEY (`status`) REFERENCES `mst_user_status` (`id`)
  ON UPDATE CASCADE
  ON DELETE RESTRICT,
  ADD CONSTRAINT `fk_mst_user_user_type_id`
  FOREIGN KEY (`user_type_id`) REFERENCES `mst_user_type` (`id`)
  ON UPDATE CASCADE
  ON DELETE RESTRICT;

ALTER TABLE `mst_access`
  ADD CONSTRAINT `fk_mst_access_user_type_id`
  FOREIGN KEY (`user_type_id`) REFERENCES `mst_user_type` (`id`)
  ON UPDATE CASCADE
  ON DELETE RESTRICT,
  ADD CONSTRAINT `fk_mst_access_api_id`
  FOREIGN KEY (`api_id`) REFERENCES `mst_api` (`id`)
  ON UPDATE CASCADE
  ON DELETE CASCADE;
