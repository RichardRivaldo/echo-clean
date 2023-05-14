-- MySQL dump 10.13  Distrib 8.0.27, for Win64 (x86_64)
--
-- Host: localhost    Database: online_shop
-- ------------------------------------------------------
-- Server version	8.0.27

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Current Database: `online_shop`
--

CREATE DATABASE /*!32312 IF NOT EXISTS*/ `online_shop` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci */ /*!80016 DEFAULT ENCRYPTION='N' */;

USE `online_shop`;

--
-- Table structure for table `like_reviews`
--

DROP TABLE IF EXISTS `like_reviews`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `like_reviews` (
  `review_id` varchar(50) NOT NULL,
  `member_id` varchar(50) NOT NULL,
  PRIMARY KEY (`review_id`,`member_id`),
  KEY `FK_LikeReview_Member` (`member_id`),
  CONSTRAINT `FK_LikeReview_Member` FOREIGN KEY (`member_id`) REFERENCES `members` (`member_id`) ON DELETE CASCADE,
  CONSTRAINT `FK_LikeReview_Review` FOREIGN KEY (`review_id`) REFERENCES `reviews` (`review_id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `like_reviews`
--

LOCK TABLES `like_reviews` WRITE;
/*!40000 ALTER TABLE `like_reviews` DISABLE KEYS */;
INSERT INTO `like_reviews` (`review_id`, `member_id`) VALUES ('3c4c8cac-9635-458b-806d-f3e03342c32e','3cbdb3a5-e825-4f41-9eed-ae66746e54e3'),('3c4c8cac-9635-458b-806d-f3e03342c32e','c1068326-97ea-4b7b-905a-927fe93a08d6');
/*!40000 ALTER TABLE `like_reviews` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `members`
--

DROP TABLE IF EXISTS `members`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `members` (
  `member_id` varchar(50) NOT NULL,
  `username` varchar(50) NOT NULL,
  `gender` enum('Male','Female') DEFAULT NULL,
  `skin_type` enum('Normal','Sensitive','Dry','Oily') NOT NULL,
  `skin_color` enum('White','Yellow','Black') NOT NULL,
  PRIMARY KEY (`member_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `members`
--

LOCK TABLES `members` WRITE;
/*!40000 ALTER TABLE `members` DISABLE KEYS */;
INSERT INTO `members` (`member_id`, `username`, `gender`, `skin_type`, `skin_color`) VALUES ('3cbdb3a5-e825-4f41-9eed-ae66746e54e3','test female member','Female','Normal','Yellow'),('c1068326-97ea-4b7b-905a-927fe93a08d6','test32','Female','Normal','Yellow');
/*!40000 ALTER TABLE `members` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `products`
--

DROP TABLE IF EXISTS `products`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `products` (
  `product_id` varchar(50) NOT NULL,
  `product_name` varchar(255) NOT NULL,
  `price` decimal(12,3) NOT NULL,
  PRIMARY KEY (`product_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `products`
--

LOCK TABLES `products` WRITE;
/*!40000 ALTER TABLE `products` DISABLE KEYS */;
INSERT INTO `products` (`product_id`, `product_name`, `price`) VALUES ('8e756f80-2dd1-4f63-bca3-a081d25e231a','other product test',10121.000),('8e756f80-2dd1-4f63-bca3-a081d25e83b2','productest',10000.000);
/*!40000 ALTER TABLE `products` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `reviews`
--

DROP TABLE IF EXISTS `reviews`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `reviews` (
  `review_id` varchar(50) NOT NULL,
  `member_id` varchar(50) NOT NULL,
  `product_id` varchar(50) NOT NULL,
  `desc_review` varchar(500) DEFAULT NULL,
  PRIMARY KEY (`review_id`),
  KEY `FK_Review_Member` (`member_id`),
  KEY `FK_Review_Product` (`product_id`),
  CONSTRAINT `FK_Review_Member` FOREIGN KEY (`member_id`) REFERENCES `members` (`member_id`) ON DELETE CASCADE,
  CONSTRAINT `FK_Review_Product` FOREIGN KEY (`product_id`) REFERENCES `products` (`product_id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `reviews`
--

LOCK TABLES `reviews` WRITE;
/*!40000 ALTER TABLE `reviews` DISABLE KEYS */;
INSERT INTO `reviews` (`review_id`, `member_id`, `product_id`, `desc_review`) VALUES ('3c4c8cac-9635-458b-806d-f3e03342c32e','3cbdb3a5-e825-4f41-9eed-ae66746e54e3','8e756f80-2dd1-4f63-bca3-a081d25e83b2','testttt'),('3c4c8cac-9635-458b-806d-f3e03342c332','c1068326-97ea-4b7b-905a-927fe93a08d6','8e756f80-2dd1-4f63-bca3-a081d25e83b2','test2'),('3c4c8cac-9635-458b-806d-f3e03342fg1a','3cbdb3a5-e825-4f41-9eed-ae66746e54e3','8e756f80-2dd1-4f63-bca3-a081d25e231a','good');
/*!40000 ALTER TABLE `reviews` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2023-05-14 12:29:59
