CREATE TABLE
    `user` (
        `id` INT NOT NULL UNIQUE PRIMARY KEY AUTO_INCREMENT,
        `name` VARCHAR(255) NULL,
        `phone_number` VARCHAR(20) NOT NULL
    )