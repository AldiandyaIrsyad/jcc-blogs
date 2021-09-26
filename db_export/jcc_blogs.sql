-- phpMyAdmin SQL Dump
-- version 5.1.0
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: Sep 26, 2021 at 07:47 AM
-- Server version: 10.4.18-MariaDB
-- PHP Version: 7.3.27

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `sql6440119`
--

-- --------------------------------------------------------

--
-- Table structure for table `categories`
--

CREATE TABLE `categories` (
  `id` int(10) UNSIGNED NOT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `title` varchar(255) DEFAULT NULL,
  `description` varchar(255) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `categories`
--

INSERT INTO `categories` (`id`, `created_at`, `updated_at`, `deleted_at`, `title`, `description`) VALUES
(1, '2021-09-26 10:48:41', '2021-09-26 10:48:41', '2021-09-26 10:49:18', 'concept', 'Language agnostik post'),
(2, '2021-09-26 10:48:58', '2021-09-26 12:06:15', NULL, 'implementation', 'an engineer approach!'),
(3, '2021-09-26 10:50:10', '2021-09-26 11:41:48', NULL, 'concept', 'Language agnostic post');

-- --------------------------------------------------------

--
-- Table structure for table `comments`
--

CREATE TABLE `comments` (
  `id` int(10) UNSIGNED NOT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `user_id` int(10) UNSIGNED DEFAULT NULL,
  `post_id` int(10) UNSIGNED DEFAULT NULL,
  `content` text DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `comments`
--

INSERT INTO `comments` (`id`, `created_at`, `updated_at`, `deleted_at`, `user_id`, `post_id`, `content`) VALUES
(1, '2021-09-26 11:45:04', '2021-09-26 11:45:04', '2021-09-26 11:54:52', 2, 10, 'This is why python is the best language!'),
(2, '2021-09-26 11:55:23', '2021-09-26 11:55:23', '2021-09-26 11:55:38', 2, 10, 'This is why python is the best language!');

-- --------------------------------------------------------

--
-- Table structure for table `posts`
--

CREATE TABLE `posts` (
  `id` int(10) UNSIGNED NOT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `user_id` int(10) UNSIGNED DEFAULT NULL,
  `title` varchar(255) DEFAULT NULL,
  `summary` varchar(255) DEFAULT NULL,
  `content` text DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `posts`
--

INSERT INTO `posts` (`id`, `created_at`, `updated_at`, `deleted_at`, `user_id`, `title`, `summary`, `content`) VALUES
(10, '2021-09-26 11:41:25', '2021-09-26 11:41:25', NULL, 4, 'Python shortest path on a graph', 'Creating a DFS Implementation to find shortest path on a graph', 'To create this you have to ...'),
(11, '2021-09-26 11:41:48', '2021-09-26 11:41:48', NULL, 5, 'Understanding Linkedlist with haskell', 'Using haskell to explain the concept of linkedlist', 'Now we\'re going to explain how to create...'),
(12, '2021-09-26 12:04:45', '2021-09-26 12:04:45', NULL, 4, 'Python shortest path on a graph', 'Creating a DFS Implementation to find shortest path on a graph', 'To create this you have to ...'),
(13, '2021-09-26 12:04:54', '2021-09-26 12:04:54', NULL, 4, 'Python shortest path on a graph', 'Creating a DFS Implementation to find shortest path on a graph', 'To create this you have to ...'),
(14, '2021-09-26 12:05:08', '2021-09-26 12:05:08', NULL, 4, 'Python shortest path on a graph', 'Creating a DFS Implementation to find shortest path on a graph', 'To create this you have to ...'),
(15, '2021-09-26 12:05:23', '2021-09-26 12:05:23', NULL, 4, 'Python shortest path on a graph', 'Creating a DFS Implementation to find shortest path on a graph', 'To create this you have to ...'),
(16, '2021-09-26 12:06:15', '2021-09-26 12:06:15', NULL, 4, 'Python shortest path on a graph', 'Creating a DFS Implementation to find shortest path on a graph', 'To create this you have to ...');

-- --------------------------------------------------------

--
-- Table structure for table `post_categories`
--

CREATE TABLE `post_categories` (
  `post_id` int(10) UNSIGNED NOT NULL,
  `category_id` int(10) UNSIGNED NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `post_categories`
--

INSERT INTO `post_categories` (`post_id`, `category_id`) VALUES
(0, 0),
(7, 2),
(8, 2),
(9, 2),
(10, 2),
(11, 3),
(12, 2),
(13, 2),
(14, 2),
(15, 2),
(16, 2);

-- --------------------------------------------------------

--
-- Table structure for table `post_tags`
--

CREATE TABLE `post_tags` (
  `post_id` int(10) UNSIGNED NOT NULL,
  `tag_id` int(10) UNSIGNED NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `post_tags`
--

INSERT INTO `post_tags` (`post_id`, `tag_id`) VALUES
(0, 0),
(9, 2),
(9, 5),
(10, 2),
(10, 5),
(11, 3),
(11, 4),
(12, 2),
(12, 5),
(13, 2),
(13, 5),
(14, 2),
(14, 5),
(15, 2),
(15, 5),
(16, 2),
(16, 5);

-- --------------------------------------------------------

--
-- Table structure for table `tags`
--

CREATE TABLE `tags` (
  `id` int(10) UNSIGNED NOT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `title` varchar(255) DEFAULT NULL,
  `description` varchar(255) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `tags`
--

INSERT INTO `tags` (`id`, `created_at`, `updated_at`, `deleted_at`, `title`, `description`) VALUES
(1, '2021-09-26 11:02:25', '2021-09-26 11:02:25', NULL, 'r', 'A Language for data science'),
(2, '2021-09-26 11:02:28', '2021-09-26 12:06:15', NULL, 'python', 'A Language of simplicity'),
(3, '2021-09-26 11:02:31', '2021-09-26 11:41:48', NULL, 'linked list', 'connect 1 by 1'),
(4, '2021-09-26 11:02:34', '2021-09-26 11:41:48', NULL, 'haskell', 'A Language for koolkidz'),
(5, '2021-09-26 11:02:36', '2021-09-26 12:06:15', NULL, 'graph', 'let\'s get connected'),
(6, '2021-09-26 11:02:39', '2021-09-26 11:02:39', NULL, 'binary tree', 'A tree that split into two');

-- --------------------------------------------------------

--
-- Table structure for table `users`
--

CREATE TABLE `users` (
  `id` int(10) UNSIGNED NOT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `username` varchar(255) DEFAULT NULL,
  `password` varchar(255) DEFAULT NULL,
  `privillege` bigint(20) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `users`
--

INSERT INTO `users` (`id`, `created_at`, `updated_at`, `deleted_at`, `username`, `password`, `privillege`) VALUES
(1, '2021-09-26 07:53:33', '2021-09-26 07:53:33', NULL, 'aldiandya', 'password', 1),
(2, '2021-09-26 07:54:07', '2021-09-26 07:54:07', NULL, 'reader1', 'random', 3),
(3, '2021-09-26 07:54:18', '2021-09-26 07:54:18', NULL, 'reader2', 'bedalagi', 3),
(4, '2021-09-26 07:54:26', '2021-09-26 11:16:27', NULL, 'poster1', 'gantilagi', 2),
(5, '2021-09-26 07:54:35', '2021-09-26 11:16:39', NULL, 'poster2', 'gantilagi', 2),
(6, '2021-09-26 07:57:28', '2021-09-26 11:16:55', NULL, 'admin', 'admin', 1);

--
-- Indexes for dumped tables
--

--
-- Indexes for table `categories`
--
ALTER TABLE `categories`
  ADD PRIMARY KEY (`id`),
  ADD KEY `idx_categories_deleted_at` (`deleted_at`);

--
-- Indexes for table `comments`
--
ALTER TABLE `comments`
  ADD PRIMARY KEY (`id`),
  ADD KEY `idx_comments_deleted_at` (`deleted_at`);

--
-- Indexes for table `posts`
--
ALTER TABLE `posts`
  ADD PRIMARY KEY (`id`),
  ADD KEY `idx_posts_deleted_at` (`deleted_at`);

--
-- Indexes for table `post_categories`
--
ALTER TABLE `post_categories`
  ADD PRIMARY KEY (`post_id`,`category_id`);

--
-- Indexes for table `post_tags`
--
ALTER TABLE `post_tags`
  ADD PRIMARY KEY (`post_id`,`tag_id`);

--
-- Indexes for table `tags`
--
ALTER TABLE `tags`
  ADD PRIMARY KEY (`id`),
  ADD KEY `idx_tags_deleted_at` (`deleted_at`);

--
-- Indexes for table `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`id`),
  ADD KEY `idx_users_deleted_at` (`deleted_at`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `categories`
--
ALTER TABLE `categories`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=4;

--
-- AUTO_INCREMENT for table `comments`
--
ALTER TABLE `comments`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;

--
-- AUTO_INCREMENT for table `posts`
--
ALTER TABLE `posts`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=17;

--
-- AUTO_INCREMENT for table `tags`
--
ALTER TABLE `tags`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=7;

--
-- AUTO_INCREMENT for table `users`
--
ALTER TABLE `users`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=7;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;

GRANT USAGE ON *.* TO `jcc_blogs`@`%` IDENTIFIED BY PASSWORD '*253874315E2FF061B3CF618B0A1CCFB2749E64F5';

GRANT ALL PRIVILEGES ON `jcc_blogs`.* TO `jcc_blogs`@`%`;
