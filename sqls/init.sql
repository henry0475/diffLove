CREATE TABLE `diffLove_db`.`users` (
	`id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
    `username` varchar(32) NOT NULL,
	`nickname` varchar(32) NOT NULL,
	`password` varchar(40) NOT NULL,
	`add_time` int(10) UNSIGNED NOT NULL,
	`gender` tinyint(1) DEFAULT 1 COMMENT '0girl-1boy',
    `status` tinyint(1) DEFAULT 0,
	PRIMARY KEY (`id`),
	KEY `name`(`username`) USING BTREE
) ENGINE=InnoDB
DEFAULT CHARACTER SET=utf8mb4 COLLATE=utf8mb4_general_ci
COMMENT='Sys Users';

CREATE TABLE `diffLove_db`.`msg_board` (
	`id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
	`bid` bigint NOT NULL COMMENT 'board id for different couples',
    `user_id` bigint NOT NULL,
	`content` text NOT NULL,
	`is_public` tinyint(1) DEFAULT 0,
	`add_time` int(10) UNSIGNED NOT NULL,
    `status` tinyint(1) DEFAULT 0,
	PRIMARY KEY (`id`),
	KEY `uid`(`user_id`) USING BTREE
) ENGINE=InnoDB
DEFAULT CHARACTER SET=utf8mb4 COLLATE=utf8mb4_general_ci
COMMENT='Message Board';

CREATE TABLE `diffLove_db`.`couples` (
	`id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
	`bid` bigint NOT NULL COMMENT 'board id for different couples',
	`vid` bigint NOT NULL COMMENT 'visited map',
    `user_id_1` bigint NOT NULL,
	`user_id_2` bigint NOT NULL,
	`add_time` int(10) UNSIGNED NOT NULL,
    `status` tinyint(1) DEFAULT 0,
	PRIMARY KEY (`id`)
) ENGINE=InnoDB
DEFAULT CHARACTER SET=utf8mb4 COLLATE=utf8mb4_general_ci
COMMENT='Couples';

CREATE TABLE `diffLove_db`.`visited_map` (
	`id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
	`vid` bigint NOT NULL COMMENT '',
    `user_id_1` bigint NOT NULL,
	`user_id_2` bigint NOT NULL,
	`add_time` int(10) UNSIGNED NOT NULL,
    `status` tinyint(1) DEFAULT 0,
	`title` varchar(32) DEFAULT '',
	`content` text DEFAULT '',
	`special_words` varchar(128) DEFAULT '',
	`point_lat` double(12,6) default 0.0,
  	`point_long` double(12,6) default 0.0,
	PRIMARY KEY (`id`)
) ENGINE=InnoDB
DEFAULT CHARACTER SET=utf8mb4 COLLATE=utf8mb4_general_ci
COMMENT='Visited points on map';