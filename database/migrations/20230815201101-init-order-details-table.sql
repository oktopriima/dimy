-- +migrate Up
CREATE TABLE IF NOT EXISTS `order_details`
(
    `id`          INT UNSIGNED NOT NULL AUTO_INCREMENT,
    `orders_id`   INT UNSIGNED NOT NULL,
    `products_id` INT UNSIGNED NOT NULL,
    `quantity`    DECIMAL      NOT NULL,
    `price`       DECIMAL      NOT NULL,
    `deleted_at`  TIMESTAMP    NULL DEFAULT NULL,
    PRIMARY KEY (`id`),
    UNIQUE INDEX `id_UNIQUE` (`id` ASC) VISIBLE,
    INDEX `fk_order_details_orders1_idx` (`orders_id` ASC) VISIBLE,
    INDEX `fk_order_details_products1_idx` (`products_id` ASC) VISIBLE,
    CONSTRAINT `fk_order_details_orders1`
        FOREIGN KEY (`orders_id`)
            REFERENCES `orders` (`id`)
            ON DELETE NO ACTION
            ON UPDATE NO ACTION,
    CONSTRAINT `fk_order_details_products1`
        FOREIGN KEY (`products_id`)
            REFERENCES `products` (`id`)
            ON DELETE NO ACTION
            ON UPDATE NO ACTION
)
    ENGINE = InnoDB;
-- +migrate Down
