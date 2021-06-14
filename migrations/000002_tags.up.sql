CREATE TABLE `tags`
(
    `id`         int          NOT NULL AUTO_INCREMENT,
    `user_id`    int          NOT NULL,
    `name`       varchar(100) NOT NULL,
    `color`      char(7) DEFAULT NULL,
    `created_at` datetime     NOT NULL,
    `updated_at` datetime     NOT NULL,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;

CREATE
UNIQUE INDEX `user_id_tag_name_idx`
ON `snippets`.`tags` (`user_id`,`name`)
USING BTREE;


CREATE TABLE `snippet_tags`
(
    `id`         int      NOT NULL AUTO_INCREMENT,
    `snippet_id` int      NOT NULL,
    `tag_id`     int      NOT NULL,
    `created_at` datetime NOT NULL,
    `updated_at` datetime NOT NULL,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;

CREATE
UNIQUE INDEX `snippet_tag_idx`
ON `snippets`.`snippet_tags` (`snippet_id`,`tag_id`)
USING BTREE;

