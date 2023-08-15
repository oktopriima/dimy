-- +migrate Up
CREATE TABLE IF NOT EXISTS `order_payments`
(
    `id`                 INT UNSIGNED NOT NULL AUTO_INCREMENT,
    `orders_id`          INT UNSIGNED NOT NULL,
    `payment_methods_id` INT UNSIGNED NOT NULL,
    `value`              DECIMAL      NOT NULL,
    `status`             VARCHAR(45)  NOT NULL,
    `deleted_at`         TIMESTAMP    NULL DEFAULT NULL,
    PRIMARY KEY (`id`),
    UNIQUE INDEX `id_UNIQUE` (`id` ASC) VISIBLE,
    INDEX `fk_order_payments_orders1_idx` (`orders_id` ASC) VISIBLE,
    INDEX `fk_order_payments_payment_methods1_idx` (`payment_methods_id` ASC) VISIBLE,
    CONSTRAINT `fk_order_payments_orders1`
        FOREIGN KEY (`orders_id`)
            REFERENCES `orders` (`id`)
            ON DELETE NO ACTION
            ON UPDATE NO ACTION,
    CONSTRAINT `fk_order_payments_payment_methods1`
        FOREIGN KEY (`payment_methods_id`)
            REFERENCES `payment_methods` (`id`)
            ON DELETE NO ACTION
            ON UPDATE NO ACTION
)
    ENGINE = InnoDB;
-- +migrate Down
DROP TABLE IF EXISTS `order_payments`;