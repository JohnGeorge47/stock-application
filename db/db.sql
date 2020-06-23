CREATE TABLE `stockapp`.`users` (
  `user_id` INT NOT NULL AUTO_INCREMENT,
  `user_name` VARCHAR(45) NOT NULL,
  `email_id` VARCHAR(45) NOT NULL,
  `created_at` TIMESTAMP NULL,
  PRIMARY KEY (`user_id`));

CREATE TABLE `stockapp`.`passwords` (
  `idpasswords` INT NOT NULL AUTO_INCREMENT,
  `user_id` INT NOT NULL,
  `email_id` VARCHAR(45) NOT NULL,
  `password` VARCHAR(256) NOT NULL,
  `updated_at` VARCHAR(45) NULL,
  PRIMARY KEY (`idpasswords`),
  UNIQUE INDEX `user_id_UNIQUE` (`user_id` ASC),
  UNIQUE INDEX `email_id_UNIQUE` (`email_id` ASC),
  CONSTRAINT `fk_passwords_1`
    FOREIGN KEY (`user_id`)
    REFERENCES `stockapp`.`users` (`user_id`)
    ON DELETE CASCADE
    ON UPDATE NO ACTION);

CREATE TABLE `stockapp`.`stocks` (
  `stock_id` INT NOT NULL AUTO_INCREMENT,
  `stock_name` VARCHAR(45) NOT NULL,
  `stock_val` INT NOT NULL,
  `max_val` INT NULL,
  `min_val` INT NULL,
  `updated_at` TIMESTAMP NULL DEFAULT current_timestamp on update current_timestamp,
  PRIMARY KEY (`stock_id`),
  UNIQUE INDEX `stock_id_UNIQUE` (`stock_id` ASC),
  UNIQUE INDEX `stock_name_UNIQUE` (`stock_name` ASC));

CREATE TABLE `stockapp`.`user_stock` (
  `iduser_stock` INT NOT NULL AUTO_INCREMENT,
  `user_id` INT NULL,
  `stock_id` INT NULL,
  PRIMARY KEY (`iduser_stock`),
  INDEX `fk_user_stock_1_idx` (`user_id` ASC),
  INDEX `fk_user_stock_2_idx` (`stock_id` ASC),
  CONSTRAINT `fk_user_stock_1`
    FOREIGN KEY (`user_id`)
    REFERENCES `stockapp`.`users` (`user_id`)
    ON DELETE  CASCADE
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_user_stock_2`
    FOREIGN KEY (`stock_id`)
    REFERENCES `stockapp`.`stocks` (`stock_id`)
    ON DELETE CASCADE
    ON UPDATE NO ACTION);