Perquisite 
- go version 1.12.6 
- mysql 5.7.25

Tool used
- Table Plus and valentina studio
- GoLand
- Insomnia

Feature
- Can create todo
- Todo have many tags
- Todo have custom date and reminder
- todo have detail note in it


1. go mod init
```
$ github.com/hidayatullahap/go-todo-example
```
2. Create main package
3. Project structure

4. create database
```
CREATE DATABASE medium_todo;
```

5. create todos and tags table

```
-- CREATE TABLE "todos" ----------------------------------------
CREATE TABLE `todos` (
	`id` Int( 11 ) AUTO_INCREMENT  NOT NULL,
	`message` Text NOT NULL,
	`note` Text NULL,
	`custom_date` Timestamp NULL,
	`is_done` TinyInt( 1 ) NOT NULL DEFAULT 0,
	`is_reminded` TinyInt( 1 ) NOT NULL DEFAULT 0,
	`created_at` Timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
	`updated_at` Timestamp NOT NULL ON UPDATE CURRENT_TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	PRIMARY KEY ( `id` ) )
ENGINE = InnoDB;
-- -------------------------------------------------------------

INSERT INTO `todos` ( `message`) VALUES ( "Hello todo" );

-- CREATE TABLE "tags" -----------------------------------------
CREATE TABLE `tags` (
    `id` Int( 11 ) AUTO_INCREMENT NOT NULL,
    `name` VarChar( 255 ) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
    `created_at` Timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` Timestamp NOT NULL ON UPDATE CURRENT_TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY ( `id` ) )
    CHARACTER SET = utf8
    COLLATE = utf8_general_ci
    ENGINE = InnoDB;
-- -------------------------------------------------------------

INSERT INTO `tags`(`name`) VALUES ('Priority');


-- CREATE TABLE "todo_tags" ------------------------------------
CREATE TABLE `todo_tags` (
     `id` Int( 11 ) AUTO_INCREMENT NOT NULL,
     `todo_id` Int( 11 ) NOT NULL,
     `tag_id` Int( 11 ) NOT NULL,
     PRIMARY KEY ( `id` ) )
    CHARACTER SET = utf8
    COLLATE = utf8_general_ci
    ENGINE = InnoDB;
-- -------------------------------------------------------------

INSERT INTO `todo_tags`(`todo_id`,`tag_id`) VALUES (1,1);
```

6. add dotenv for configs
library for load env variables
```
$ go get github.com/joho/godotenv
```
get echo library for http server
```
$ go get github.com/labstack/echo/v4
```

7. add hello world server
8. move hello func to action folder
9. create mysql connection and app env
10. show todo list with created repo
11. Read todo request to insert into db
12. create repo for create todo -> e08d069
13. refactor todo find bad request error
14. add todo detail
15. refactor before update