CREATE TABLE `users`
(
  `id` char(20)  NOT NULL,
  `name` varchar(10) NOT NULL,
  `description` varchar(100) NOT NULL DEFAULT '',
  `created_at` bigint(20) unsigned NOT NULL,
  `updated_at` bigint(20) unsigned NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
