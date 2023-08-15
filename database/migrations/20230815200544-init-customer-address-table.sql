-- +migrate Up
CREATE TABLE IF NOT EXISTS `customer_addresses`
(
    `id`           INT UNSIGNED NOT NULL AUTO_INCREMENT,
    `customers_id` INT UNSIGNED NOT NULL,
    `address`      TEXT         NOT NULL,
    `deleted_at`   TIMESTAMP    NULL DEFAULT NULL,
    UNIQUE INDEX `id_UNIQUE` (`id` ASC) VISIBLE,
    PRIMARY KEY (`id`),
    INDEX `fk_customer_addresses_customers_idx` (`customers_id` ASC) VISIBLE,
    CONSTRAINT `fk_customer_addresses_customers`
        FOREIGN KEY (`customers_id`)
            REFERENCES `customers` (`id`)
            ON DELETE NO ACTION
            ON UPDATE NO ACTION
)
    ENGINE = InnoDB;
-- +migrate Down
DROP TABLE IF EXISTS customer_addresses;