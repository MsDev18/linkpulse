CREATE TABLE
    `otp` (
        `id` INT NOT NULL UNIQUE PRIMARY KEY AUTO_INCREMENT,
        `user_id` INT NOT NULL UNIQUE,
        `code` VARCHAR(20),
        `expires_at` TIMESTAMP NOT NULL,
        FOREIGN KEY (`user_id`) REFERENCES `user`(`id`)
    )