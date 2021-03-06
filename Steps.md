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
15. refactor before update endpoint
16. add update endpoint
17. add crud tags and validator
18. create custom validator message
19. add todo store validator
20. create custom error message
21. add index in todo_tags 
```
-- CREATE INDEX "i_tag_todos_tag_id" ---------------------------
CREATE INDEX `i_tag_todos_tag_id` ON `todo_tags` (`todo_id`);
-- -------------------------------------------------------------
```
21. read tag array request
22. create tags builder
23. TodoTags gorm:"-" on create, change create todo repo with tx
24. show detail todo with tag relationship, add omitempty in tags_json (dont give space after "," it will not work)
25. add updating todo-tag relationship
26. edit detail todo function for search relation
27. add list tags 
28. refactor validator
29. add update status todo