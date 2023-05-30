create table `items` (
                          `id` bigint unsigned auto_increment,
                          `name` varchar(100) default null,
                          `sell_in` int unsigned default null,
                          `quality` int unsigned default null,
                          `created_at` timestamp default CURRENT_TIMESTAMP,
                          `updated_at` timestamp default CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP,
                          PRIMARY KEY (`id`)
) ENGINE='InnoDB' AUTO_INCREMENT=1 DEFAULT CHARSET='utf8mb4' COLLATE='utf8mb4_0900_ai_ci';

create table `images` (
    `id` bigint unsigned auto_increment,
    `url` varchar(200) default null,
    `item_id` bigint unsigned not null,
    `created_at` timestamp default CURRENT_TIMESTAMP,
    `updated_at` timestamp default CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    KEY `fk_images_item_id_items` (`item_id`),
    CONSTRAINT `fk_images_item_id_items` FOREIGN KEY (`item_id`) REFERENCES `items` (`id`) ON DELETE CASCADE
) ENGINE='InnoDB' AUTO_INCREMENT=1 DEFAULT CHARSET='utf8mb4' COLLATE='utf8mb4_0900_ai_ci';

create table `users` (
                          `id` bigint unsigned auto_increment,
                          `username` varchar(200) not null,
                          `password` varchar(200) default null,
                          `created_at` timestamp default CURRENT_TIMESTAMP,
                          `updated_at` timestamp default CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP,
                          PRIMARY KEY (`id`)
) ENGINE='InnoDB' AUTO_INCREMENT=1 DEFAULT CHARSET='utf8mb4' COLLATE='utf8mb4_0900_ai_ci';