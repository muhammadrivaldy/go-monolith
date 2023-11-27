CREATE TABLE IF NOT EXISTS `mst_version` (
    `id` INT(11) NOT NULL AUTO_INCREMENT,
    `version` VARCHAR(10) NOT NULL,
    `support` TINYINT(1) NOT NULL,
    `created_by` INT NOT NULL,
    `created_at` DATETIME NOT NULL,
    `updated_by` INT NOT NULL,
    `updated_at` DATETIME NOT NULL,
    PRIMARY KEY (`id`)
);

INSERT INTO `mst_version` (
    `id`, 
    `version`, 
    `support`,
    `created_by`,
    `created_at`,
    `updated_by`,
    `updated_at`
) VALUES (1, "v1.0.0", 1, 1, NOW(), 1, NOW());