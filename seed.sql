INSERT INTO `users` (`id`, `username`, `password`, `firstname`, `lastname`, `email`, `phone`, `user_role`, `profile_picture`, `is_deleted`, `refresh_token`, `reset_password`, `created_at`, `updated_at`)
VALUES
	('3b2c1306-12d5-4357-aab8-bbf49c9a6a20', 'user', '$2a$14$kZgfh5ttpvw9YfqSx496UO8FGLd08WQ/NWZtXFweuIhZ5A9KPeprC', 'angga', 'pamungkas', 'kap21kapuser@gmail.com', '3434534', 'user', '576113_realsam_o2_logo_partner.png', 0, NULL, NULL, '2022-07-07 15:54:52', '2022-07-07 15:54:52'),
	('fc7f19bd-a9a8-445c-bf5f-6e73936ff24f', 'admin', '$2a$14$YcBNkbqjbzUzEH3EQzPIA.0U.KXzFOFn0xAc6vgNhAvygbjmCxrhq', 'angga', 'pamungkas', 'kap21kap@gmail.com', '3434534', 'admin', '831252_realsam_o2_logo_partner.png', 0, NULL, NULL, '2022-07-07 15:54:24', '2022-07-07 15:54:24');


-- goose mysql "root:password_root_betest123@tcp(mysql)/betest?parseTime=true" up