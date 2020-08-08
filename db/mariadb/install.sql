CREATE TABLE `files` (
    `id` INT(16) NOT NULL auto_increment,
    `original_filename` VARCHAR(1024) NOT NULL default '',
    `filename` VARCHAR(1024) NOT NULL default '',
    `stored_filename` VARCHAR(1024) NOT NULL default '',
    `md5` VARCHAR(256) NOT NULL default '',
    `sha256` VARCHAR(256) NOT NULL default '',
    `upload_date` DATETIME NOT NULL default '0000-00-00 00:00:00',
    `last_update` DATETIME NOT NULL default '0000-00-00 00:00:00',
    `status` TINYINT(1) NOT NULL default 0,
    PRIMARY KEY(`id`)
);

CREATE TABLE `users` (
    `id` INT(16) NOT NULL auto_increment,
    `username` VARCHAR(1024) NOT NULL default '',
    `email` VARCHAR(1024) NOT NULL default '',
    `password` VARCHAR(256) NOT NULL default '',
    `register_date` DATETIME NOT NULL default '0000-00-00 00:00:00',
    `last_login_date` DATETIME NOT NULL default '0000-00-00 00:00:00',
    `status` TINYINT(1) NOT NULL default 0,
    PRIMARY KEY(`id`)
);

CREATE TABLE `file_actions` (
    `id` BIGINT(16) NOT NULL auto_increment,
    `file_id` INT(16) NOT NULL default 0,
    `user_id` INT(16) NOT NULL default 0,
    `action` VARCHAR(256) NOT NULL default 0,
    `date` DATETIME NOT NULL default '0000-00-00 00:00:00',
    PRIMARY KEY(`id`)
);

CREATE TABLE `user_actions` (
    `id` BIGINT(16) NOT NULL auto_increment,
    `target_id` INT(16) NOT NULL default 0,
    `user_id` INT(16) NOT NULL default 0,
    `action` VARCHAR(256) NOT NULL default 0,
    `date` DATETIME NOT NULL default '0000-00-00 00:00:00',
    PRIMARY KEY(`id`)
);
