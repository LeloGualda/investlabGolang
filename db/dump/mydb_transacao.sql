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
-- Table structure for table `transacao`
--

DROP TABLE IF EXISTS `transacao`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `transacao` (
  `id_transacao` int(11) NOT NULL AUTO_INCREMENT,
  `data` date NOT NULL,
  `valor` float NOT NULL,
  `id_user` int(6) NOT NULL,
  `tipo` int(1) DEFAULT NULL,
  `descricao` varchar(45) DEFAULT NULL,
  PRIMARY KEY (`id_transacao`),
  KEY `fk_transacao_usuario1_idx` (`id_user`),
  CONSTRAINT `fk_transacao_usuario1` FOREIGN KEY (`id_user`) REFERENCES `usuario` (`id_user`) ON DELETE NO ACTION ON UPDATE NO ACTION
) ENGINE=InnoDB AUTO_INCREMENT=18 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `transacao`
--

LOCK TABLES `transacao` WRITE;
/*!40000 ALTER TABLE `transacao` DISABLE KEYS */;
INSERT INTO `transacao` VALUES (1,'2019-09-08',3000.57,2,1,'deposito'),(2,'2019-09-08',30000.6,2,1,'deposito'),(3,'2019-09-08',3023.9,1,2,'venda acao:MSFT.MEX'),(4,'2019-09-08',-3023.9,2,3,'compra acaoMSFT.MEX'),(5,'2019-09-08',3023.9,1,2,'venda acao:MSFT.MEX'),(6,'2019-09-08',-3023.9,2,3,'compra acaoMSFT.MEX'),(7,'2019-09-08',3023.9,1,2,'venda acao:MSFT.MEX'),(8,'2019-09-08',-3023.9,2,3,'compra acaoMSFT.MEX'),(9,'2019-09-08',3051.39,1,2,'venda acao:MSFT.MEX'),(10,'2019-09-08',-3051.39,2,3,'compra acaoMSFT.MEX'),(11,'2019-09-08',10000,3,1,'deposito'),(12,'2019-09-08',3023.9,1,2,'venda acao:MSFT.MEX'),(13,'2019-09-08',-3023.9,2,3,'compra acaoMSFT.MEX'),(14,'2019-09-08',3051.39,1,2,'venda acao:MSFT.MEX'),(15,'2019-09-08',-3051.39,2,3,'compra acaoMSFT.MEX'),(16,'2019-09-08',2776.49,2,2,'venda acao:MSFT.MEX'),(17,'2019-09-08',-2776.49,1,3,'compra acaoMSFT.MEX');
/*!40000 ALTER TABLE `transacao` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2019-11-21 19:17:39
