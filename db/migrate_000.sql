DROP TABLE IF EXISTS `pale_rabbit_mst`.`item`;
CREATE TABLE `pale_rabbit_mst`.`item` (
  `code` VARCHAR(10) NOT NULL COMMENT '',
  `name` VARCHAR(100) NOT NULL COMMENT '',
  `price` DECIMAL NULL DEFAULT 0 COMMENT '',
  `updated_by` VARCHAR(40) NULL COMMENT '',
  `updated_at` VARCHAR(45) NULL COMMENT '',
  PRIMARY KEY (`code`)  COMMENT '');