-- phpMyAdmin SQL Dump
-- version 4.4.15.10
-- https://www.phpmyadmin.net
--
-- Host: localhost
-- Generation Time: 2024 年 3 月 31 日 08:30
-- サーバのバージョン： 8.0.36
-- PHP Version: 7.3.33

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `golang_user`
--

-- --------------------------------------------------------

--
-- テーブルの構造 `user_id`
--

CREATE TABLE IF NOT EXISTS `user_id` (
  `ITEM0` int NOT NULL COMMENT 'プライマリーキー',
  `ITEM1` varchar(64) NOT NULL COMMENT 'ユーザID',
  `ITEM2` varchar(64) NOT NULL COMMENT 'ユーザPW',
  `ITEM3` varchar(64) NOT NULL COMMENT '作成者',
  `ITEM4` datetime NOT NULL COMMENT '作成日付'
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- テーブルのデータのダンプ `user_id`
--

INSERT INTO `user_id` (`ITEM0`, `ITEM1`, `ITEM2`, `ITEM3`, `ITEM4`) VALUES
(1, 'testuserid', 'testuserpw', 'system', '2024-03-31 00:00:00');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `user_id`
--
ALTER TABLE `user_id`
  ADD PRIMARY KEY (`ITEM0`),
  ADD KEY `ITEM1` (`ITEM1`(32),`ITEM2`) USING BTREE;

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `user_id`
--
ALTER TABLE `user_id`
  MODIFY `ITEM0` int NOT NULL AUTO_INCREMENT COMMENT 'プライマリーキー',AUTO_INCREMENT=2;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
