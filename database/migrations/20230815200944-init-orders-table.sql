-- +migrate Up
CREATE TABLE IF NOT EXISTS `orders`
(
    `id`                    INT UNSIGNED NOT NULL AUTO_INCREMENT,
    `customers_id`          INT UNSIGNED NOT NULL,
    `customer_addresses_id` INT UNSIGNED NOT NULL,
    `total_price`           DECIMAL      NOT NULL,
    `status`                VARCHAR(45)  NOT NULL,
    `deleted_at`            TIMESTAMP    NULL DEFAULT NULL,
    PRIMARY KEY (`id`),
    UNIQUE INDEX `id_UNIQUE` (`id` ASC) VISIBLE,
    INDEX `fk_orders_customers1_idx` (`customers_id` ASC) VISIBLE,
    INDEX `fk_orders_customer_addresses1_idx` (`customer_addresses_id` ASC) VISIBLE,
    CONSTRAINT `fk_orders_customers1`
        FOREIGN KEY (`customers_id`)
            REFERENCES `customers` (`id`)
            ON DELETE NO ACTION
            ON UPDATE NO ACTION,
    CONSTRAINT `fk_orders_customer_addresses1`
        FOREIGN KEY (`customer_addresses_id`)
            REFERENCES `customer_addresses` (`id`)
            ON DELETE NO ACTION
            ON UPDATE NO ACTION
)
    ENGINE = InnoDB;
-- +migrate Down
DROP TABLE IF EXISTS `orders`;