-- +migrate Up
CREATE TABLE IF NOT EXISTS `customers`
(
    `id`         INT UNSIGNED NOT NULL AUTO_INCREMENT,
    `name`       VARCHAR(255) NOT NULL,
    `deleted_at` TIMESTAMP    NULL DEFAULT NULL,
    UNIQUE INDEX `id_UNIQUE` (`id` ASC) VISIBLE,
    PRIMARY KEY (`id`)
)
    ENGINE = InnoDB;
-- +migrate Down
DROP TABLE IF EXISTS `customers`;
