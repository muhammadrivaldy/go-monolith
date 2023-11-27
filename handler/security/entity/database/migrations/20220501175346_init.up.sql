CREATE TABLE IF NOT EXISTS `mst_service` (
  `id` INT(11) NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(100) NOT NULL,
  `created_at` DATETIME NOT NULL,
  `updated_at` DATETIME NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `name_UNIQUE` (`name`)
);

CREATE TABLE IF NOT EXISTS `mst_api` (
  `id` INT(11) NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(100) CHARACTER SET latin1 NOT NULL,
  `endpoint` VARCHAR(255) CHARACTER SET latin1 NOT NULL,
  `method` VARCHAR(15) CHARACTER SET latin1 NOT NULL,
  `service_id` INT(11) NOT NULL,
  `created_at` DATETIME NOT NULL,
  `updated_at` DATETIME NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `name_UNIQUE` (`name`),
  UNIQUE KEY `endpoint_method_UNIQUE` (`endpoint`,`method`)
);

CREATE TABLE IF NOT EXISTS `mst_access` (
  `id` INT(11) NOT NULL AUTO_INCREMENT,
  `user_type_id` INT(2) NOT NULL,
  `api_id` INT(3) NOT NULL,
  `created_by` INT NOT NULL,
  `created_at` DATETIME NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `usertype_api_UNIQUE` (`user_type_id`,`api_id`)
);