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