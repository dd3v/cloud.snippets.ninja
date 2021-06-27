CREATE TABLE `sessions`
(
    `id`            int          NOT NULL AUTO_INCREMENT,
    `user_id`       int          NOT NULL,
    `refresh_token` char(36)     NOT NULL,
    `exp`           timestamp    NOT NULL,
    `ip`            varchar(25)  NOT NULL,
    `user_agent`    varchar(500) NOT NULL,
    `created_at`    datetime     NOT NULL,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE utf8mb4_unicode_ci;

CREATE TABLE `snippets`
(
    `id`                    int          NOT NULL AUTO_INCREMENT,
    `user_id`               int          NOT NULL,
    `favorite`              tinyint(1) NOT NULL DEFAULT '0',
    `access_level`          tinyint(1) NOT NULL DEFAULT '0',
    `title`                 varchar(500) NOT NULL,
    `content`               text         NOT NULL,
    `tags`                  json NULL,
    `language`              varchar(20)  NOT NULL,
    `custom_editor_options` json         NOT NULL,
    `created_at`            datetime     NOT NULL,
    `updated_at`            datetime     NOT NULL,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE utf8mb4_unicode_ci;

CREATE TABLE `users`
(
    `id`         int          NOT NULL AUTO_INCREMENT,
    `password`   varchar(60)  NOT NULL,
    `login`      varchar(100) NOT NULL,
    `email`      varchar(100) NOT NULL,
    `created_at` datetime     NOT NULL,
    `updated_at` datetime     NOT NULL,
    PRIMARY KEY (`id`),
    UNIQUE KEY `email_idx` (`email`) USING BTREE,
    UNIQUE KEY `login_idx` (`login`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE utf8mb4_unicode_ci;


CREATE FULLTEXT INDEX `title_content` ON snippets(title,content)
