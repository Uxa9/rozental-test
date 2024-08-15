-- phpMyAdmin SQL Dump
-- version 5.2.0
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1:3306
-- Generation Time: Aug 15, 2024 at 09:35 AM
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

--
-- Indexes for dumped tables
--

--
-- Indexes for table `inventory_operations`
--
ALTER TABLE `inventory_operations`
  ADD PRIMARY KEY (`id`),
  ADD KEY `dst_executor` (`dst_executor`),
  ADD KEY `src_executor` (`src_executor`),
  ADD KEY `status` (`status`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `inventory_operations`
--
ALTER TABLE `inventory_operations`
  MODIFY `id` int NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=5;

--
-- Constraints for dumped tables
--

--
-- Constraints for table `inventory_operations`
--
ALTER TABLE `inventory_operations`
  ADD CONSTRAINT `inventory_operations_ibfk_1` FOREIGN KEY (`dst_executor`) REFERENCES `user` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  ADD CONSTRAINT `inventory_operations_ibfk_2` FOREIGN KEY (`src_executor`) REFERENCES `user` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  ADD CONSTRAINT `inventory_operations_ibfk_3` FOREIGN KEY (`status`) REFERENCES `inventory_operation_status` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
