-- phpMyAdmin SQL Dump
-- version 4.4.15.10
-- https://www.phpmyadmin.net
--
-- Host: localhost
-- Generation Time: 2024 年 3 月 31 日 12:06
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
-- テーブルの構造 `customer_master`
--

CREATE TABLE IF NOT EXISTS `customer_master` (
  `ITEM0` int NOT NULL COMMENT 'プライマリーキー',
  `ITEM1` varchar(128) NOT NULL COMMENT '氏名',
  `ITEM2` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT 'フリガナ',
  `ITEM3` varchar(1024) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '住所',
  `ITEM4` int NOT NULL DEFAULT '0' COMMENT '年齢',
  `ITEM5` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '固定電話',
  `ITEM6` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '携帯電話',
  `ITEM7` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT 'メールアドレス',
  `ITEM8` varchar(2048) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '備考欄',
  `ITEM9` varchar(64) NOT NULL COMMENT '作成者',
  `ITEM10` datetime NOT NULL COMMENT '作成日付'
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- テーブルのデータのダンプ `customer_master`
--

INSERT INTO `customer_master` (`ITEM0`, `ITEM1`, `ITEM2`, `ITEM3`, `ITEM4`, `ITEM5`, `ITEM6`, `ITEM7`, `ITEM8`, `ITEM9`, `ITEM10`) VALUES
(1, '織田裕二', 'オダユウジ', '東京都港区麻布十番１－１－１', 56, '0333335555', '0903335555', 'y.oda@gmail.com', 'テスト登録１です。', 'system', '2024-03-31 00:00:00'),
(2, '水谷豊', 'ミズタニユタカ', '東京都港区六本木１－１－１', 71, '0311112222', '09011112222', 'y.mizutani@gmail.com', 'テスト登録２です。', 'system', '2024-03-31 00:00:00');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `customer_master`
--
ALTER TABLE `customer_master`
  ADD PRIMARY KEY (`ITEM0`),
  ADD KEY `ITEM1` (`ITEM1`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `customer_master`
--
ALTER TABLE `customer_master`
  MODIFY `ITEM0` int NOT NULL AUTO_INCREMENT COMMENT 'プライマリーキー',AUTO_INCREMENT=3;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
