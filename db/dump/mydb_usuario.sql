-- MySQL dump 10.13  Distrib 5.7.27, for Linux (x86_64)
--
-- Host: localhost    Database: mydb
-- ------------------------------------------------------
-- Server version	5.7.27-0ubuntu0.18.04.1

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `usuario`
--

DROP TABLE IF EXISTS `usuario`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `usuario` (
  `id_user` int(6) NOT NULL AUTO_INCREMENT,
  `username` varchar(45) NOT NULL,
  `password` varchar(60) NOT NULL,
  PRIMARY KEY (`id_user`,`username`),
  UNIQUE KEY `username_UNIQUE` (`username`)
) ENGINE=InnoDB AUTO_INCREMENT=81 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `usuario`
--

LOCK TABLES `usuario` WRITE;
/*!40000 ALTER TABLE `usuario` DISABLE KEYS */;
INSERT INTO `usuario` VALUES (1,'MSFT.MEX','$2a$08$4Dtt2WUrGvfZHlHztFfSfu28A/Pn1YhsEBJbabbSfVo4jrMwJWfqW'),(2,'aurelio','$2a$08$r9adBgIjz/yNLUJTJCBj5u1pG/AY1wyCuP2HwfP9eFiLcQYGOQwxu'),(3,'prof','$2a$08$dnWD7ILwZPedQkNZzurv0e9wOqXZmG1KAiRH9bLL90isprFX/QEBu'),(4,'GhJ8Y!z0%qDPkyLb','$2a$08$ssfgyth3l79H2z1uuretF./WYzVKqSv36T0olaQ/YzT.la.5NJUg6'),(5,'BOT','$2a$08$7p2yOkymASKvuSvgmS3F.OfAMd77xlw8Q/01ASAgWmSW9txIfBTgC'),(7,'ADMIN','$2a$08$wAwE/5TcRySZs.foSdp1MeP8ejuQ0k7npQdTnL5ZUSYODHL8wyFsi'),(26,'teste','$2a$08$uMVRO1fzrOZD/cgMOFhvb.mzXGiJLstjpTwDmBHqtCvZfU4fGDSe2'),(27,'','$2a$08$CvZndqq1C2NDaexfiU18ZuECgWiSilukO2Ay05/JsqFaSISIGgvsW'),(28,'teste2','$2a$08$b6yZN0jbgLa/MIwfgj2...5QKcp8Ar.T5KtGfxtQqNcgEpz.QY3MS');
/*!40000 ALTER TABLE `usuario` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2019-11-22 19:59:22
