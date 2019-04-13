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
