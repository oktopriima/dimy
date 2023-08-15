-- +migrate Up
CREATE TABLE IF NOT EXISTS `payment_methods`
(
    `id`         INT UNSIGNED NOT NULL AUTO_INCREMENT,
    `name`       VARCHAR(45)  NOT NULL,
    `is_active`  TINYINT      NOT NULL DEFAULT 1,
    `deleted_at` TIMESTAMP    NULL DEFAULT NULL,
    PRIMARY KEY (`id`),
    UNIQUE INDEX `id_UNIQUE` (`id` ASC) VISIBLE
)
    ENGINE = InnoDB;
-- +migrate Down
DROP TABLE IF EXISTS payment_methods;