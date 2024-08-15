-- phpMyAdmin SQL Dump
-- version 5.2.0
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1:3306
-- Generation Time: Aug 15, 2024 at 10:34 AM
-- Server version: 8.0.30
-- PHP Version: 8.1.9

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `rg-test`
--

-- --------------------------------------------------------

--
-- Table structure for table `inventory`
--

CREATE TABLE `inventory` (
  `id` int NOT NULL,
  `name` text NOT NULL,
  `id_executor` int NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Dumping data for table `inventory`
--

INSERT INTO `inventory` (`id`, `name`, `id_executor`) VALUES
(1, 'Перфоратор', 2),
(2, 'Шуруповерт', 1),
(3, 'Молоток', 2);

-- --------------------------------------------------------

--
-- Table structure for table `inventory_operations`
--

CREATE TABLE `inventory_operations` (
  `id` int NOT NULL,
  `src_executor` int NOT NULL,
  `dst_executor` int NOT NULL,
  `request_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `status_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `status` int NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Dumping data for table `inventory_operations`
--

INSERT INTO `inventory_operations` (`id`, `src_executor`, `dst_executor`, `request_time`, `status_time`, `status`) VALUES
(1, 2, 1, '2024-08-14 01:30:27', '2024-08-14 01:30:27', 4),
(2, 2, 5, '2024-08-14 23:03:23', '2024-08-14 23:03:23', 1),
(3, 1, 5, '2024-08-14 23:03:23', '2024-08-14 23:03:23', 2),
(4, 2, 5, '2024-08-15 00:27:38', '2024-08-15 00:27:38', 1);

-- --------------------------------------------------------

--
-- Table structure for table `inventory_operations_detail`
--

CREATE TABLE `inventory_operations_detail` (
  `id` int NOT NULL,
  `operation_id` int NOT NULL,
  `inventory_id` int NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Dumping data for table `inventory_operations_detail`
--

INSERT INTO `inventory_operations_detail` (`id`, `operation_id`, `inventory_id`) VALUES
(1, 1, 1),
(2, 2, 1),
(3, 3, 2),
(4, 4, 3);

-- --------------------------------------------------------

--
-- Table structure for table `inventory_operation_status`
--

CREATE TABLE `inventory_operation_status` (
  `id` int NOT NULL,
  `name` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Dumping data for table `inventory_operation_status`
--

INSERT INTO `inventory_operation_status` (`id`, `name`) VALUES
(1, 'accept'),
(2, 'cancel'),
(3, 'reject'),
(4, 'await');

-- --------------------------------------------------------

--
-- Table structure for table `user`
--

CREATE TABLE `user` (
  `id` int NOT NULL,
  `name` text NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Dumping data for table `user`
--

INSERT INTO `user` (`id`, `name`) VALUES
(1, 'Валера'),
(2, 'Антон'),
(4, 'афыафыа'),
(5, 'Геннадий');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `inventory`
--
ALTER TABLE `inventory`
  ADD PRIMARY KEY (`id`),
  ADD KEY `id_executor` (`id_executor`);

--
-- Indexes for table `inventory_operations`
--
ALTER TABLE `inventory_operations`
  ADD PRIMARY KEY (`id`),
  ADD KEY `dst_executor` (`dst_executor`),
  ADD KEY `src_executor` (`src_executor`),
  ADD KEY `status` (`status`);

--
-- Indexes for table `inventory_operations_detail`
--
ALTER TABLE `inventory_operations_detail`
  ADD PRIMARY KEY (`id`),
  ADD KEY `inventory_id` (`inventory_id`),
  ADD KEY `operation_id` (`operation_id`);

--
-- Indexes for table `inventory_operation_status`
--
ALTER TABLE `inventory_operation_status`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `user`
--
ALTER TABLE `user`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `inventory`
--
ALTER TABLE `inventory`
  MODIFY `id` int NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=4;

--
-- AUTO_INCREMENT for table `inventory_operations`
--
ALTER TABLE `inventory_operations`
  MODIFY `id` int NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=5;

--
-- AUTO_INCREMENT for table `inventory_operations_detail`
--
ALTER TABLE `inventory_operations_detail`
  MODIFY `id` int NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=5;

--
-- AUTO_INCREMENT for table `inventory_operation_status`
--
ALTER TABLE `inventory_operation_status`
  MODIFY `id` int NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=5;

--
-- AUTO_INCREMENT for table `user`
--
ALTER TABLE `user`
  MODIFY `id` int NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=503;

--
-- Constraints for dumped tables
--

--
-- Constraints for table `inventory`
--
ALTER TABLE `inventory`
  ADD CONSTRAINT `inventory_ibfk_1` FOREIGN KEY (`id_executor`) REFERENCES `user` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT;

--
-- Constraints for table `inventory_operations`
--
ALTER TABLE `inventory_operations`
  ADD CONSTRAINT `inventory_operations_ibfk_1` FOREIGN KEY (`dst_executor`) REFERENCES `user` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  ADD CONSTRAINT `inventory_operations_ibfk_2` FOREIGN KEY (`src_executor`) REFERENCES `user` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  ADD CONSTRAINT `inventory_operations_ibfk_3` FOREIGN KEY (`status`) REFERENCES `inventory_operation_status` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT;

--
-- Constraints for table `inventory_operations_detail`
--
ALTER TABLE `inventory_operations_detail`
  ADD CONSTRAINT `inventory_operations_detail_ibfk_1` FOREIGN KEY (`inventory_id`) REFERENCES `inventory` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  ADD CONSTRAINT `inventory_operations_detail_ibfk_2` FOREIGN KEY (`operation_id`) REFERENCES `inventory_operations` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
