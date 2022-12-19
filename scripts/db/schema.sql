CREATE TABLE `shorturl` (
                            `id` int NOT NULL AUTO_INCREMENT,
                            `short_id` varchar(126) NOT NULL,
                            `url` varchar(256) NOT NULL,
                            `created_at` datetime NOT NULL,
                            PRIMARY KEY (`id`),
                            UNIQUE KEY `shorturl_pk` (`short_id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `access_log` (
                              `id` int NOT NULL AUTO_INCREMENT,
                              `created_at` datetime NOT NULL,
                              `shorturl_id` int NOT NULL,
                              `ip` varchar(15) NOT NULL,
                              PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;