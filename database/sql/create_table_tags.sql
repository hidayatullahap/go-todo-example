-- CREATE TABLE "tags" -----------------------------------------
CREATE TABLE `tags`
(
    `id`         Int(11) AUTO_INCREMENT                                  NOT NULL,
    `name`       VarChar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
    `created_at` Timestamp                                               NOT NULL                             DEFAULT CURRENT_TIMESTAMP,
    `updated_at` Timestamp                                               NOT NULL ON UPDATE CURRENT_TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
)
    CHARACTER SET = utf8
    COLLATE = utf8_general_ci
    ENGINE = InnoDB;
-- -------------------------------------------------------------

INSERT INTO `tags`(`name`)
VALUES ('Priority');


-- CREATE TABLE "todo_tags" ------------------------------------
CREATE TABLE `todo_tags`
(
    `id`      Int(11) AUTO_INCREMENT NOT NULL,
    `todo_id` Int(11)                NOT NULL,
    `tag_id`  Int(11)                NOT NULL,
    PRIMARY KEY (`id`),
    CONSTRAINT `u_todo_tag_id` UNIQUE (`todo_id`, `tag_id`)
)
    CHARACTER SET = utf8
    COLLATE = utf8_general_ci
    ENGINE = InnoDB;
-- -------------------------------------------------------------

INSERT INTO `todo_tags`(`todo_id`, `tag_id`)
VALUES (1, 1);
