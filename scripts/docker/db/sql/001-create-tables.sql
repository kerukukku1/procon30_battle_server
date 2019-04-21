---- drop ----
DROP TABLE IF EXISTS `users`;

---- create ----
create table IF not exists `users`
(
    `id`               INT(20) AUTO_INCREMENT,
    `token`            VARCHAR(256) NOT NULL,
    `name`             VARCHAR(256) NOT NULL,
    PRIMARY KEY (`id`)
) DEFAULT CHARSET=utf8 COLLATE=utf8_bin;
