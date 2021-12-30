-- phpMyAdmin SQL Dump
-- version 5.1.1
-- https://www.phpmyadmin.net/
--
-- Host: localhost
-- Waktu pembuatan: 26 Des 2021 pada 16.00
-- Versi server: 10.4.21-MariaDB
-- Versi PHP: 7.4.23

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `go-startup`
--

-- --------------------------------------------------------

--
-- Struktur dari tabel `campaigns`
--

CREATE TABLE `campaigns` (
  `id` int(11) UNSIGNED NOT NULL,
  `user_id` int(11) DEFAULT NULL,
  `name` varchar(255) DEFAULT NULL,
  `short_description` varchar(255) DEFAULT NULL,
  `description` text DEFAULT NULL,
  `perks` text DEFAULT NULL,
  `backer_count` int(11) DEFAULT NULL,
  `goal_amount` int(11) DEFAULT NULL,
  `current_amount` int(11) DEFAULT NULL,
  `slug` varchar(255) DEFAULT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data untuk tabel `campaigns`
--

INSERT INTO `campaigns` (`id`, `user_id`, `name`, `short_description`, `description`, `perks`, `backer_count`, `goal_amount`, `current_amount`, `slug`, `created_at`, `updated_at`) VALUES
(1, 1, 'campaign', 'shorttt ', 'longg', 'prtks1,perks2,perks3', 0, 1000000, 0, 'campaign1', '2021-12-05 07:57:43', '2021-12-05 07:57:43'),
(2, 1, 'campaign2', 'shorttt ', 'longg', 'prtks1,perks2,perks3', 0, 1000000, 0, 'campaign2', '2021-12-05 07:57:43', '2021-12-05 07:57:43'),
(3, 19, 'campaign2', 'shorttt ', 'longg', 'prtks1,perks2,perks3', 0, 2000000000, 0, 'campaign3', '2021-12-05 07:57:43', '2021-12-05 07:57:43'),
(4, 1, 'penggalanga  dana JKT48', 'macet', 'macet cokkkkkk', 'hadiah,hadiah2,hadiah3', 0, 1000000000, 0, 'penggalanga-dana-jkt48-s-int-1', '2021-12-09 21:25:32', '2021-12-09 21:25:32'),
(5, 17, 'penggalanga  dana JKT48 agar tidak bubar', 'macet', 'macet cokkkkkk', 'hadiah,hadiah2,hadiah3', 0, 1000000000, 0, 'penggalanga-dana-jkt48-agar-tidak-bubar-1', '2021-12-09 21:31:07', '2021-12-09 21:31:07'),
(6, 17, 'konser cho tokimeki sendenbu indonesia', 'konser tokimeki sendenbu', 'konser tokimeki sendenbu di indonesia', 'lightstick, t-shirt, towel, photopack', 0, 2000000, 0, 'konser-cho-tokimeki-sendenbu-17', '2021-12-09 22:27:39', '2021-12-10 22:41:10'),
(7, 17, 'tokisen konser in indonesia', 'konser tokisen', 'konser yang dilaksanakn di jakarta indonesia', 't-shirt,lightstick,towel', 0, 1000000, 0, 'tokisen-konser-in-indonesia-17', '2021-12-20 23:05:57', '2021-12-20 23:05:57');

-- --------------------------------------------------------

--
-- Struktur dari tabel `campaign_images`
--

CREATE TABLE `campaign_images` (
  `id` int(11) UNSIGNED NOT NULL,
  `campaign_id` int(11) DEFAULT NULL,
  `file_name` varchar(255) DEFAULT NULL,
  `is_primary` tinyint(4) DEFAULT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data untuk tabel `campaign_images`
--

INSERT INTO `campaign_images` (`id`, `campaign_id`, `file_name`, `is_primary`, `created_at`, `updated_at`) VALUES
(1, 1, 'campaign-images/foto.jpg', 0, '2021-12-05 08:10:04', '2021-12-05 08:10:04'),
(2, 1, 'campaign-images/fotoq.jpg', 1, '2021-12-05 08:10:04', '2021-12-05 08:10:04'),
(3, 1, 'fotoq.jpg', 0, '2021-12-05 08:10:04', '2021-12-05 08:10:04'),
(4, 17, 'images/17-insomnia.png', 0, '2021-12-13 23:35:43', '2021-12-14 16:31:10'),
(5, 17, 'images/17-insomnia.png', 0, '2021-12-13 23:58:48', '2021-12-14 16:31:10'),
(6, 17, 'images/17-insomnia.png', 0, '2021-12-13 23:59:54', '2021-12-14 16:31:10'),
(7, 17, 'images/17-insomnia.png', 0, '2021-12-14 00:01:20', '2021-12-14 16:31:10'),
(8, 17, 'images/17-08_ria_kurumi_9_1.jpg', 0, '2021-12-14 16:27:12', '2021-12-14 16:31:10'),
(9, 17, 'images/17-08_ria_kurumi_9_1.jpg', 0, '2021-12-14 16:30:46', '2021-12-14 16:31:10'),
(10, 5, 'images/17-1068795.png', 1, '2021-12-14 16:31:10', '2021-12-14 16:31:10'),
(11, 6, 'images/17-insomnia.png', 1, '2021-12-14 23:59:12', '2021-12-14 23:59:12'),
(12, 6, 'images/17-insomnia.png', 0, '2021-12-15 00:03:34', '2021-12-15 00:03:34');

-- --------------------------------------------------------

--
-- Struktur dari tabel `transactions`
--

CREATE TABLE `transactions` (
  `id` int(11) UNSIGNED NOT NULL,
  `campaign_id` int(11) DEFAULT NULL,
  `user_id` int(11) DEFAULT NULL,
  `amount` int(11) DEFAULT NULL,
  `status` varchar(255) DEFAULT NULL,
  `code` varchar(255) DEFAULT NULL,
  `payment_url` varchar(255) DEFAULT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data untuk tabel `transactions`
--

INSERT INTO `transactions` (`id`, `campaign_id`, `user_id`, `amount`, `status`, `code`, `payment_url`, `created_at`, `updated_at`) VALUES
(1, 6, 17, 1000000, 'paid', NULL, NULL, '2021-12-15 17:21:20', '2021-12-15 17:21:20'),
(2, 6, 17, 2000000, 'pending', NULL, NULL, '2021-12-15 17:45:33', '2021-12-15 17:45:33'),
(3, 5, 17, 1233333, 'paid', NULL, NULL, '2021-12-16 17:44:36', '2021-12-16 17:44:36'),
(4, 18, 18, 5000000, '', '', NULL, '2021-12-20 21:15:37', '2021-12-20 21:15:37'),
(5, 18, 18, 5000000, '', '', NULL, '2021-12-20 21:16:38', '2021-12-20 21:16:38'),
(6, 18, 18, 9000000, '', '', NULL, '2021-12-20 21:17:13', '2021-12-20 21:17:13'),
(7, 18, 18, 92000000, 'Pending', '', NULL, '2021-12-20 21:23:40', '2021-12-20 21:23:40'),
(8, 18, 17, 12345000, 'Pending', '', NULL, '2021-12-20 23:01:59', '2021-12-20 23:01:59'),
(9, 7, 17, 12345000, 'Pending', '', NULL, '2021-12-20 23:06:35', '2021-12-20 23:06:35'),
(10, 7, 17, 13345000, 'Pending', '', 'https://app.sandbox.midtrans.com/snap/v2/vtweb/8e7d8f5b-8193-430c-a313-ef0cbc18890c', '2021-12-22 23:45:27', '2021-12-22 23:45:29'),
(11, 7, 17, 100000, 'Pending', '', 'https://app.sandbox.midtrans.com/snap/v2/vtweb/7c43de0d-9d02-44ee-b33f-a2c1734984a3', '2021-12-24 15:31:06', '2021-12-24 15:31:07');

-- --------------------------------------------------------

--
-- Struktur dari tabel `users`
--

CREATE TABLE `users` (
  `id` int(11) UNSIGNED NOT NULL,
  `name` varchar(255) DEFAULT NULL,
  `occupation` varchar(255) DEFAULT NULL,
  `email` varchar(255) DEFAULT NULL,
  `password_hash` varchar(255) DEFAULT NULL,
  `avatar_file_name` varchar(255) DEFAULT NULL,
  `role` varchar(255) DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data untuk tabel `users`
--

INSERT INTO `users` (`id`, `name`, `occupation`, `email`, `password_hash`, `avatar_file_name`, `role`, `created_at`, `updated_at`) VALUES
(1, 'sakura endo', 'idol', 'sakura@gmail.com', '1234', 'images/avatar.jpg', 'member', '2021-11-09 19:07:00', '2021-12-05 14:59:30'),
(2, 'Test', '', '', '', '', '', NULL, NULL),
(3, 'sakura endo', '', '', '', '', '', NULL, NULL),
(4, 'sakura endoooo', '', '', '', '', '', NULL, NULL),
(5, 'test simpan service', 'idol', 'email@email.com', '$2a$04$nHh9HU4TbwaGe3U.lG2DMeRkJZrHouIBXdjfHAOjtJbQNd77tegDu', '', 'user', NULL, NULL),
(6, 'Sakuraaaaa', 'idol', 'sakuraa@gmail.com', '$2a$04$bej674W66IUAyHTgubM9v.qmQP2/LeksKKR0jjZGNyO7KeZH0dVvK', '', 'user', NULL, NULL),
(7, 'Sakuraaaaa', 'idol', 'sakuraa@gmail.com', '$2a$04$HKw10IxL8G5dBB.AKyKZuup30.3/3Lrd4D2Oxa8tEi9AsNKBTd0UO', '', 'user', NULL, NULL),
(8, 'Sakuraaaaa', 'idol', 'sakuraa@gmail.com', '$2a$04$dAxtsce.WWagbGR0C0QrTOCEFZLAf/0XkZ32hH5lGZnElNipICKJi', '', 'user', NULL, NULL),
(9, 'Sakuraaaaa miyawaki', 'idol', 'sakuraa48@gmail.com', '$2a$04$ISPFObHop7NnnL/Y7YnKPOxfLLe/BNOKwpHJaIyk4Uq1DTVx3EaCS', '', 'user', NULL, NULL),
(10, 'Sakuraaaaa miyawaki', 'idol', 'sakuraa48@gmail.com', '$2a$04$U6zJTOA/Y/.7tZfRPLhEV./SvqQlPiAIHNdlCv/TNsHMPQQ9k.6na', '', 'user', NULL, NULL),
(11, 'Sakuraaaaa miyawaki', 'idol', 'sakuraa48@gmail.com', '$2a$04$6chPwIZi/.pIVzt8l78SvOYi0auiYhvub3.YWs1edR9ZEk2gHAyXy', '', 'user', NULL, NULL),
(13, 'saki seto', 'idol', 'sakichan@challange.com', '$2a$04$3gZXxOwNH1OAgkdSjkbc9O4f6L3j7wlS09SLWIrmTfk9P1cMJs.9i', '', 'user', NULL, NULL),
(14, 'saki seto1', 'idol', 'sakichan@challange.com', '$2a$04$IQVnHLEW1.tFW6ByvpQ/.OtnCTG3ApDzZaYWD7.5DAwrarlOB7BcK', '', 'user', '0000-00-00 00:00:00', '0000-00-00 00:00:00'),
(15, 'saki seto1', 'idol', 'sakichan@challange.com', '$2a$04$B9ZlPWZkp36nBuT52cJhauViq9KTcroRmkr1QB22pc9sdcgfRCqP2', '', 'user', '0000-00-00 00:00:00', '0000-00-00 00:00:00'),
(16, 'saki ', 'idol', 'sakichan@challange.com', '$2a$04$.lWk2fDF9PdNg5Db6QnqB.fWuIqL8Z0UcEiOGpXkwacRhuMNQT6V.', '', 'user', '2021-11-23 22:49:18', '2021-11-23 22:49:18'),
(17, 'aki suda', 'idol', 'akisuda@tokisen.com', '$2a$04$VupfIutRg2RkxC4KwJN/e.JUnSsdKycrGQTNx54xYQhskP/KkVu3K', '', 'user', '2021-11-28 11:56:06', '2021-11-28 11:56:06'),
(18, 'aki suda', 'idol', 'akisuda@tokisen.com', '$2a$04$e44jXNC6gLxLdlatpq5VDO4b2ct/iuZAP/fyMG8K88N6ANxu6VgkS', '', 'user', '2021-11-28 13:44:30', '2021-11-28 13:44:30'),
(19, 'aki suda', 'idol', 'akisuda@tokisen.com', '$2a$04$DBvKeDSpha04ZHtb7XchHu14a0TgekoVyiJSHCGXsnni9QMi0cnD2', '', 'user', '2021-11-29 21:18:52', '2021-11-29 21:18:52'),
(20, 'akari kito', 'idol', 'akarin@kawaii.com', '$2a$04$9CuwJ3btasPcuJ2YzokJS.IsHkfGcSIGf4txv17OzUI35Tmj4o/tS', '', 'user', '2021-11-29 21:48:57', '2021-11-29 21:48:57'),
(21, 'misaki chan', 'idol', 'misaki@kawaii.com', '$2a$04$J0.F/I.NqQ81Oey/fu4vM.UlHKHh6Ivqx.Cyu5qdHgxSfbfcl0pPW', 'images/21-insomnia.png', 'user', '2021-12-01 23:31:13', '2021-12-03 20:44:17');

--
-- Indexes for dumped tables
--

--
-- Indeks untuk tabel `campaigns`
--
ALTER TABLE `campaigns`
  ADD PRIMARY KEY (`id`);

--
-- Indeks untuk tabel `campaign_images`
--
ALTER TABLE `campaign_images`
  ADD PRIMARY KEY (`id`);

--
-- Indeks untuk tabel `transactions`
--
ALTER TABLE `transactions`
  ADD PRIMARY KEY (`id`);

--
-- Indeks untuk tabel `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT untuk tabel yang dibuang
--

--
-- AUTO_INCREMENT untuk tabel `campaigns`
--
ALTER TABLE `campaigns`
  MODIFY `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=8;

--
-- AUTO_INCREMENT untuk tabel `campaign_images`
--
ALTER TABLE `campaign_images`
  MODIFY `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=13;

--
-- AUTO_INCREMENT untuk tabel `transactions`
--
ALTER TABLE `transactions`
  MODIFY `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=12;

--
-- AUTO_INCREMENT untuk tabel `users`
--
ALTER TABLE `users`
  MODIFY `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=22;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
