CREATE TABLE `user` (
    `id` varchar(36) NOT NULL,
    `email` varchar(36) NOT NULL UNIQUE,
    `password` varchar(255) NOT NULL,
    `role` varchar(30) DEFAULT 'attender',
    `status` varchar(25) NOT NULL DEFAULT 'inactive',
    `created_at` timestamp(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
    `updated_at` timestamp(3) NULL DEFAULT NULL,
    `created_by` varchar(25) NOT NULL DEFAULT 'admin',
    `updated_by` varchar(25) NULL DEFAULT NULL,
    PRIMARY KEY (`id`),
    KEY `status` (`status`),
    KEY `role` (`role`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;