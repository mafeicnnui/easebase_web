/*
SQLyog Ultimate v11.24 (64 bit)
MySQL - 5.7.30-33-57-log : Database - puppet2.0
*********************************************************************
*/

/*!40101 SET NAMES utf8 */;

/*!40101 SET SQL_MODE=''*/;

/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;
CREATE DATABASE /*!32312 IF NOT EXISTS*/`puppet2.0` /*!40100 DEFAULT CHARACTER SET utf8mb4 */;

USE `puppet2.0`;

/*Table structure for table `t_dmlx` */

CREATE TABLE `t_dmlx` (
  `dm` varchar(10) NOT NULL COMMENT '大类代码',
  `mc` varchar(100) DEFAULT NULL COMMENT '大类名称',
  `flag` varchar(1) DEFAULT NULL COMMENT '大类状态',
  `create_time` datetime DEFAULT NULL COMMENT '创建时间',
  `update_time` datetime DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`dm`),
  KEY `idx_t_dmlx` (`dm`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/*Data for the table `t_dmlx` */

insert  into `t_dmlx`(`dm`,`mc`,`flag`,`create_time`,`update_time`) values ('01','部门','1','2020-05-11 09:14:29','2020-05-14 15:28:20');
insert  into `t_dmlx`(`dm`,`mc`,`flag`,`create_time`,`update_time`) values ('02','数据源类型','1','2020-05-11 09:14:29','2020-05-11 09:14:29');
insert  into `t_dmlx`(`dm`,`mc`,`flag`,`create_time`,`update_time`) values ('03','数据库环境','1','2020-05-11 09:14:29','2020-05-11 09:14:29');
insert  into `t_dmlx`(`dm`,`mc`,`flag`,`create_time`,`update_time`) values ('04','性别','1','2020-05-11 09:14:29','2020-05-11 09:14:29');
insert  into `t_dmlx`(`dm`,`mc`,`flag`,`create_time`,`update_time`) values ('05','项目编码','1','2020-05-11 09:14:29','2020-05-11 09:14:29');
insert  into `t_dmlx`(`dm`,`mc`,`flag`,`create_time`,`update_time`) values ('06','服务器类型','1','2020-05-11 09:14:29','2020-05-11 09:14:29');
insert  into `t_dmlx`(`dm`,`mc`,`flag`,`create_time`,`update_time`) values ('07','数据库实例类型','1','2020-05-11 09:14:29','2020-05-11 09:14:29');
insert  into `t_dmlx`(`dm`,`mc`,`flag`,`create_time`,`update_time`) values ('08','同步业务类型','1','2020-05-11 09:14:29','2020-05-11 09:14:29');
insert  into `t_dmlx`(`dm`,`mc`,`flag`,`create_time`,`update_time`) values ('09','同步数据类型','1','2020-05-11 09:14:29','2020-05-11 09:14:29');
insert  into `t_dmlx`(`dm`,`mc`,`flag`,`create_time`,`update_time`) values ('10','同步时间类型','1','2020-05-11 09:14:29','2020-05-11 09:14:29');
insert  into `t_dmlx`(`dm`,`mc`,`flag`,`create_time`,`update_time`) values ('11','数据源类型','1','2020-05-11 09:14:29','2020-05-11 09:14:29');
insert  into `t_dmlx`(`dm`,`mc`,`flag`,`create_time`,`update_time`) values ('12','版本号','1','2020-05-11 09:14:29','2020-05-11 09:14:29');
insert  into `t_dmlx`(`dm`,`mc`,`flag`,`create_time`,`update_time`) values ('13','工单类型','1','2020-05-11 09:14:29','2020-05-11 09:14:29');
insert  into `t_dmlx`(`dm`,`mc`,`flag`,`create_time`,`update_time`) values ('14','测试大类','1','2020-05-11 09:14:29','2020-05-11 09:14:29');
insert  into `t_dmlx`(`dm`,`mc`,`flag`,`create_time`,`update_time`) values ('15','zookeeper地址','1','2020-05-11 09:14:29','2020-05-11 09:14:29');
insert  into `t_dmlx`(`dm`,`mc`,`flag`,`create_time`,`update_time`) values ('16','hbase thrift接口地址','1','2020-05-11 09:14:29','2020-05-11 09:14:29');
insert  into `t_dmlx`(`dm`,`mc`,`flag`,`create_time`,`update_time`) values ('17','工单类型','1','2020-05-11 09:14:29','2020-05-11 09:14:29');
insert  into `t_dmlx`(`dm`,`mc`,`flag`,`create_time`,`update_time`) values ('18','项目组','1','2020-05-11 09:14:29','2020-05-11 09:14:29');
insert  into `t_dmlx`(`dm`,`mc`,`flag`,`create_time`,`update_time`) values ('19','工单状态','1','2020-05-11 09:14:29','2020-05-11 09:14:29');
insert  into `t_dmlx`(`dm`,`mc`,`flag`,`create_time`,`update_time`) values ('20','归档时间类型','1','2020-05-11 09:14:29','2020-05-11 09:14:29');
insert  into `t_dmlx`(`dm`,`mc`,`flag`,`create_time`,`update_time`) values ('21','归档策略','1','2020-05-11 09:14:29','2020-05-11 09:14:29');
insert  into `t_dmlx`(`dm`,`mc`,`flag`,`create_time`,`update_time`) values ('22','迁移策略','1','2020-05-11 09:14:29','2020-05-11 09:14:29');
insert  into `t_dmlx`(`dm`,`mc`,`flag`,`create_time`,`update_time`) values ('23','指标类型','1','2020-05-11 09:14:29','2020-05-11 09:14:29');
insert  into `t_dmlx`(`dm`,`mc`,`flag`,`create_time`,`update_time`) values ('24','阀值类型','1','2020-05-11 09:14:29','2020-05-11 09:14:29');
insert  into `t_dmlx`(`dm`,`mc`,`flag`,`create_time`,`update_time`) values ('25','数据库用户状态','1','2020-05-11 09:14:29','2020-05-11 09:14:29');
insert  into `t_dmlx`(`dm`,`mc`,`flag`,`create_time`,`update_time`) values ('26','数据源代理','1','2020-07-28 17:09:43','2020-07-28 17:09:43');
insert  into `t_dmlx`(`dm`,`mc`,`flag`,`create_time`,`update_time`) values ('27','mysql版本','1','2020-07-29 13:58:41','2020-07-29 13:58:41');
insert  into `t_dmlx`(`dm`,`mc`,`flag`,`create_time`,`update_time`) values ('28','数据库实例sql状态','1','2020-07-30 17:20:10','2020-07-30 17:20:10');
insert  into `t_dmlx`(`dm`,`mc`,`flag`,`create_time`,`update_time`) values ('29','数据库-控制台-SQL状态','1','2020-08-05 17:34:20','2020-08-05 17:34:20');
insert  into `t_dmlx`(`dm`,`mc`,`flag`,`create_time`,`update_time`) values ('30','mysql5.6配置文件','1','2020-08-06 14:32:50','2020-08-06 14:32:50');
insert  into `t_dmlx`(`dm`,`mc`,`flag`,`create_time`,`update_time`) values ('31','mysql5.6权限列表','1','2020-08-06 14:32:50','2020-08-06 14:32:50');
insert  into `t_dmlx`(`dm`,`mc`,`flag`,`create_time`,`update_time`) values ('32','mysql实例状态','1','2020-08-17 18:05:36','2020-08-17 18:05:38');
insert  into `t_dmlx`(`dm`,`mc`,`flag`,`create_time`,`update_time`) values ('33','mysql下载目录','1','2020-08-19 17:38:31','2020-08-19 17:38:31');
insert  into `t_dmlx`(`dm`,`mc`,`flag`,`create_time`,`update_time`) values ('34','MinIO同步类型','1','2020-09-21 11:07:03','2020-09-21 11:07:03');
insert  into `t_dmlx`(`dm`,`mc`,`flag`,`create_time`,`update_time`) values ('35','端口映射类型','1','2020-11-13 11:31:56','2020-11-13 11:31:56');
insert  into `t_dmlx`(`dm`,`mc`,`flag`,`create_time`,`update_time`) values ('36','数据库代理本地端口','1','2020-11-21 13:37:43','2020-11-21 13:37:45');
insert  into `t_dmlx`(`dm`,`mc`,`flag`,`create_time`,`update_time`) values ('37','慢日志采集方式','1','2021-02-19 16:41:52','2021-02-19 16:41:54');
insert  into `t_dmlx`(`dm`,`mc`,`flag`,`create_time`,`update_time`) values ('38','实例类型','1','2021-04-16 15:54:35','2021-04-16 15:54:37');

/*Table structure for table `t_dmmx` */

CREATE TABLE `t_dmmx` (
  `dm` varchar(10) NOT NULL DEFAULT '' COMMENT '代码大类',
  `dmm` varchar(200) NOT NULL DEFAULT '' COMMENT '代码小类',
  `dmmc` varchar(100) DEFAULT NULL COMMENT '小类名称',
  `flag` varchar(1) DEFAULT NULL COMMENT '小类状态',
  `create_time` datetime DEFAULT NULL COMMENT '创建时间',
  `update_time` datetime DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`dm`,`dmm`),
  KEY `idx_t_dmmx` (`dm`,`dmm`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/*Data for the table `t_dmmx` */

insert  into `t_dmmx`(`dm`,`dmm`,`dmmc`,`flag`,`create_time`,`update_time`) values ('01','01','研发部','1','2020-05-11 09:15:35','2020-05-11 09:15:35');
insert  into `t_dmmx`(`dm`,`dmm`,`dmmc`,`flag`,`create_time`,`update_time`) values ('01','02','测试部','1','2020-05-11 09:15:35','2020-05-11 09:15:35');
insert  into `t_dmmx`(`dm`,`dmm`,`dmmc`,`flag`,`create_time`,`update_time`) values ('01','03','项目部','1','2020-05-11 09:15:35','2020-05-11 09:15:35');
insert  into `t_dmmx`(`dm`,`dmm`,`dmmc`,`flag`,`create_time`,`update_time`) values ('01','04','人力部','1','2020-05-11 09:15:35','2020-05-11 09:15:35');
insert  into `t_dmmx`(`dm`,`dmm`,`dmmc`,`flag`,`create_time`,`update_time`) values ('01','05','行政部','1','2020-05-11 09:15:35','2020-05-11 09:15:35');
insert  into `t_dmmx`(`dm`,`dmm`,`dmmc`,`flag`,`create_time`,`update_time`) values ('01','06','平台组','1','2020-05-11 09:15:35','2020-05-11 09:15:35');
insert  into `t_dmmx`(`dm`,`dmm`,`dmmc`,`flag`,`create_time`,`update_time`) values ('01','07','实施组','1','2020-05-11 09:15:35','2020-05-11 09:15:35');
insert  into `t_dmmx`(`dm`,`dmm`,`dmmc`,`flag`,`create_time`,`update_time`) values ('02','0','mysql','1','2020-05-11 09:15:35','2020-05-11 09:15:35');
insert  into `t_dmmx`(`dm`,`dmm`,`dmmc`,`flag`,`create_time`,`update_time`) values ('02','1','oracle','1','2020-05-11 09:15:35','2020-05-11 09:15:35');
insert  into `t_dmmx`(`dm`,`dmm`,`dmmc`,`flag`,`create_time`,`update_time`) values ('02','2','mssql','1','2020-05-11 09:15:35','2020-05-11 09:15:35');
insert  into `t_dmmx`(`dm`,`dmm`,`dmmc`,`flag`,`create_time`,`update_time`) values ('02','3','postgresql','1','2020-05-11 09:15:35','2020-05-11 09:15:35');
insert  into `t_dmmx`(`dm`,`dmm`,`dmmc`,`flag`,`create_time`,`update_time`) values ('02','4','elasticsearch','1','2020-05-11 09:15:35','2020-05-11 09:15:35');
insert  into `t_dmmx`(`dm`,`dmm`,`dmmc`,`flag`,`create_time`,`update_time`) values ('02','5','redis','1','2020-05-11 09:15:35','2020-05-11 09:15:35');
insert  into `t_dmmx`(`dm`,`dmm`,`dmmc`,`flag`,`create_time`,`update_time`) values ('02','6','mongodb','1','2020-05-11 09:15:35','2020-05-11 09:15:35');
insert  into `t_dmmx`(`dm`,`dmm`,`dmmc`,`flag`,`create_time`,`update_time`) values ('02','7','mysql-pxc','1','2020-10-29 08:54:58','2020-10-29 08:55:01');
insert  into `t_dmmx`(`dm`,`dmm`,`dmmc`,`flag`,`create_time`,`update_time`) values ('03','1','生产环境','1','2020-05-11 09:15:35','2020-05-11 09:15:35');
insert  into `t_dmmx`(`dm`,`dmm`,`dmmc`,`flag`,`create_time`,`update_time`) values ('03','2','测试环境','1','2020-05-11 09:15:35','2020-05-11 09:15:35');
insert  into `t_dmmx`(`dm`,`dmm`,`dmmc`,`flag`,`create_time`,`update_time`) values ('03','3','开发环境','1','2020-05-11 09:15:35','2020-05-11 09:15:35');
insert  into `t_dmmx`(`dm`,`dmm`,`dmmc`,`flag`,`create_time`,`update_time`) values ('03','4','预生产环境','1','2020-05-11 09:15:35','2020-05-11 09:15:35');
insert  into `t_dmmx`(`dm`,`dmm`,`dmmc`,`flag`,`create_time`,`update_time`) values ('03','5','大数据环境','1','2020-05-11 09:15:35','2020-05-11 09:15:35');
insert  into `t_dmmx`(`dm`,`dmm`,`dmmc`,`flag`,`create_time`,`update_time`) values ('03','6','平台组环境','1','2020-05-11 09:15:35','2020-05-11 09:15:35');
insert  into `t_dmmx`(`dm`,`dmm`,`dmmc`,`flag`,`create_time`,`update_time`) values ('03','7','克隆环境','1','2020-11-18 10:03:22','2020-11-18 10:03:24');
insert  into `t_dmmx`(`dm`,`dmm`,`dmmc`,`flag`,`create_time`,`update_time`) values ('04','1','男','1','2020-05-11 09:15:35','2020-05-11 09:15:35');
insert  into `t_dmmx`(`dm`,`dmm`,`dmmc`,`flag`,`create_time`,`update_time`) values ('04','2','女','1','2020-05-11 09:15:35','2020-05-11 09:15:35');
insert  into `t_dmmx`(`dm`,`dmm`,`dmmc`,`flag`,`create_time`,`update_time`) values ('05','000','合生通项目','2','2020-05-11 09:15:35','2020-05-11 09:15:35');
insert  into `t_dmmx`(`dm`,`dmm`,`dmmc`,`flag`,`create_time`,`update_time`) values ('05','001','好房项目','2','2020-05-11 09:15:35','2020-05-11 09:15:35');
insert  into `t_dmmx`(`dm`,`dmm`,`dmmc`,`flag`,`create_time`,`update_time`) values ('05','002','会付通项目','2','2020-05-11 09:15:35','2020-05-11 09:15:35');
insert  into `t_dmmx`(`dm`,`dmm`,`dmmc`,`flag`,`create_time`,`update_time`) values ('05','003','平台组项目','2','2020-05-15 11:27:57','2020-05-15 11:27:57');
insert  into `t_dmmx`(`dm`,`dmm`,`dmmc`,`flag`,`create_time`,`update_time`) values ('05','10230','西安时代广场	','1','2020-05-11 09:15:35','2020-05-11 09:15:35');
insert  into `t_dmmx`(`dm`,`dmm`,`dmmc`,`flag`,`create_time`,`update_time`) values ('05','108','成都珠江广场','1','2020-05-11 09:15:35','2020-05-11 09:15:35');
insert  into `t_dmmx`(`dm`,`dmm`,`dmmc`,`flag`,`create_time`,`update_time`) values ('05','110','上海五角场','1','2020-05-11 09:15:35','2020-05-11 09:15:35');
insert  into `t_dmmx`(`dm`,`dmm`,`dmmc`,`flag`,`create_time`,`update_time`) values ('05','132','广州嘉和南项目','1','2020-05-11 09:15:35','2020-05-11 09:15:35');
insert  into `t_dmmx`(`dm`,`dmm`,`dmmc`,`flag`,`create_time`,`update_time`) values ('05','141','北京合生财富广场','1','2020-05-11 09:15:35','2020-05-11 09:15:35');
insert  into `t_dmmx`(`dm`,`dmm`,`dmmc`,`flag`,`create_time`,`update_time`) values ('05','145','广州珠江投资大厦','1','2020-05-11 09:15:35','2020-05-11 09:15:35');
insert  into `t_dmmx`(`dm`,`dmm`,`dmmc`,`flag`,`create_time`,`update_time`) values ('05','150','广州合生骏景广场','1','2020-05-11 09:15:35','2020-05-11 09:15:35');
insert  into `t_dmmx`(`dm`,`dmm`,`dmmc`,`flag`,`create_time`,`update_time`) values ('05','159','上海静安珠江创意中心','1','2020-06-15 14:37:18','2020-06-15 14:37:18');
insert  into `t_dmmx`(`dm`,`dmm`,`dmmc`,`flag`,`create_time`,`update_time`) values ('05','164','北京合生广场(木樨园)','1','2020-05-11 09:15:35','2020-05-11 09:15:35');
insert  into `t_dmmx`(`dm`,`dmm`,`dmmc`,`flag`,`create_time`,`update_time`) values ('05','170','广州嘉和北项目','1','2020-05-11 09:15:35','2020-05-11 09:15:35');
insert  into `t_dmmx`(`dm`,`dmm`,`dmmc`,`flag`,`create_time`,`update_time`) values ('05','183','上海海云天合生财富广场','1','2020-05-11 09:15:35','2020-05-11 09:15:35');
insert  into `t_dmmx`(`dm`,`dmm`,`dmmc`,`flag`,`create_time`,`update_time`) values ('05','188','广州越华珠江国际大厦','1','2020-05-11 09:15:35','2020-05-11 09:15:35');
insert  into `t_dmmx`(`dm`,`dmm`,`dmmc`,`flag`,`create_time`,`update_time`) values ('05','194','广州南方花园	','1','2020-05-11 09:15:35','2020-05-11 09:15:35');
insert  into `t_dmmx`(`dm`,`dmm`,`dmmc`,`flag`,`create_time`,`update_time`) values ('05','20229','广州珠江国际纺织城	','1','2020-05-11 09:15:35','2020-05-11 09:15:35');
insert  into `t_dmmx`(`dm`,`dmm`,`dmmc`,`flag`,`create_time`,`update_time`) values ('05','203','杭州科华合生国贸中心','1','2020-05-11 09:15:35','2020-05-11 09:15:35');
insert  into `t_dmmx`(`dm`,`dmm`,`dmmc`,`flag`,`create_time`,`update_time`) values ('05','207','北京德胜财富广场','1','2020-05-11 09:15:35','2020-05-11 09:15:35');
insert  into `t_dmmx`(`dm`,`dmm`,`dmmc`,`flag`,`create_time`,`update_time`) values ('05','213','合生新天地','1','2020-05-11 09:15:35','2020-05-11 09:15:35');
insert  into `t_dmmx`(`dm`,`dmm`,`dmmc`,`flag`,`create_time`,`update_time`) values ('05','218','北京朝阳合生汇','1','2020-05-11 09:15:35','2020-05-11 09:15:35');
insert  into `t_dmmx`(`dm`,`dmm`,`dmmc`,`flag`,`create_time`,`update_time`) values ('05','234','上海青浦米格	','1','2020-05-11 09:15:35','2020-05-11 09:15:35');
insert  into `t_dmmx`(`dm`,`dmm`,`dmmc`,`flag`,`create_time`,`update_time`) values ('05','237','广州增城合生汇','1','2020-05-11 09:15:35','2020-05-11 09:15:35');
insert  into `t_dmmx`(`dm`,`dmm`,`dmmc`,`flag`,`create_time`,`update_time`) values ('05','246','天津合生国际大厦','1','2020-09-07 10:22:54','2020-09-07 10:22:54');
insert  into `t_dmmx`(`dm`,`dmm`,`dmmc`,`flag`,`create_time`,`update_time`) values ('05','249','上海茶岸园同步服务器','1','2020-06-01 16:19:40','2020-06-01 16:19:42');
insert  into `t_dmmx`(`dm`,`dmm`,`dmmc`,`flag`,`create_time`,`update_time`) values ('05','666','商管项目','2','2020-12-17 16:28:54','2020-12-17 16:28:56');
insert  into `t_dmmx`(`dm`,`dmm`,`dmmc`,`flag`,`create_time`,`update_time`) values ('05','777','集团项目','2','2020-05-11 09:15:35','2020-05-11 09:15:35');
insert  into `t_dmmx`(`dm`,`dmm`,`dmmc`,`flag`,`create_time`,`update_time`) values ('05','888','德利多富租赁系统','2','2020-05-11 09:15:35','2020-05-11 09:15:35');
insert  into `t_dmmx`(`dm`,`dmm`,`dmmc`,`flag`,`create_time`,`update_time`) values ('05','992','合生通内网备份服务器','2',NULL,NULL);
insert  into `t_dmmx`(`dm`,`dmm`,`dmmc`,`flag`,`create_time`,`update_time`) values ('05','994','合生通业务库同步服务器','2',NULL,NULL);
insert  into `t_dmmx`(`dm`,`dmm`,`dmmc`,`flag`,`create_time`,`update_time`) values ('05','995','合生通内网同步服务器','2',NULL,NULL);
insert  into `t_dmmx`(`dm`,`dmm`,`dmmc`,`flag`,`create_time`,`update_time`) values ('05','999','阿里云同步中转服务器','2','2020-05-11 09:15:35','2020-05-11 09:15:35');
insert  into `t_dmmx`(`dm`,`dmm`,`dmmc`,`flag`,`create_time`,`update_time`) values ('06','1','备份服务器','1','2020-05-11 09:15:35','2020-05-11 09:15:35');
insert  into `t_dmmx`(`dm`,`dmm`,`dmmc`,`flag`,`create_time`,`update_time`) values ('06','2','同步服务器','1','2020-05-11 09:15:35','2020-05-11 09:15:35');
insert  into `t_dmmx`(`dm`,`dmm`,`dmmc`,`flag`,`create_time`,`update_time`) values ('06','3','数据库服务器','1','2020-05-11 09:15:35','2020-05-11 09:15:35');
insert  into `t_dmmx`(`dm`,`dmm`,`dmmc`,`flag`,`create_time`,`update_time`) values ('06','4','应用服务器','1','2020-07-16 15:14:18','2020-07-16 15:14:18');
insert  into `t_dmmx`(`dm`,`dmm`,`dmmc`,`flag`,`create_time`,`update_time`) values ('07','1','ECS','1','2020-05-11 09:15:35','2020-05-11 09:15:35');
insert  into `t_dmmx`(`dm`,`dmm`,`dmmc`,`flag`,`create_time`,`update_time`) values ('07','2','RDS','1','2020-05-11 09:15:35','2020-05-11 09:15:35');
insert  into `t_dmmx`(`dm`,`dmm`,`dmmc`,`flag`,`create_time`,`update_time`) values ('08','1','离线客流','1','2020-05-11 09:15:35','2020-05-11 09:15:35');
insert  into `t_dmmx`(`dm`,`dmm`,`dmmc`,`flag`,`create_time`,`update_time`) values ('08','10','CMS数据','1','2020-05-11 09:15:35','2020-05-11 09:15:35');
insert  into `t_dmmx`(`dm`,`dmm`,`dmmc`,`flag`,`create_time`,`update_time`) values ('08','11','商户数据','1','2020-05-11 09:15:35','2020-05-11 09:15:35');
insert  into `t_dmmx`(`dm`,`dmm`,`dmmc`,`flag`,`create_time`,`update_time`) values ('08','12','卡券数据','1','2020-05-11 09:15:35','2020-05-11 09:15:35');
insert  into `t_dmmx`(`dm`,`dmm`,`dmmc`,`flag`,`create_time`,`update_time`) values ('08','13','订单数据','1','2020-05-11 09:15:35','2020-05-11 09:15:35');
insert  into `t_dmmx`(`dm`,`dmm`,`dmmc`,`flag`,`create_time`,`update_time`) values ('08','14','水单数据','1','2020-05-11 09:15:35','2020-05-11 09:15:35');
insert  into `t_dmmx`(`dm`,`dmm`,`dmmc`,`flag`,`create_time`,`update_time`) values ('08','15','好房业务','1','2020-05-11 09:15:35','2020-05-11 09:15:35');
insert  into `t_dmmx`(`dm`,`dmm`,`dmmc`,`flag`,`create_time`,`update_time`) values ('08','16','dataX同步','1','2020-05-11 09:15:35','2020-05-11 09:15:35');
insert  into `t_dmmx`(`dm`,`dmm`,`dmmc`,`flag`,`create_time`,`update_time`) values ('08','17','店铺商户关系','1','2020-05-11 09:15:35','2020-05-11 09:15:35');
insert  into `t_dmmx`(`dm`,`dmm`,`dmmc`,`flag`,`create_time`,`update_time`) values ('08','18','BI统计','1','2020-05-11 09:15:35','2020-05-11 09:15:35');
insert  into `t_dmmx`(`dm`,`dmm`,`dmmc`,`flag`,`create_time`,`update_time`) values ('08','19','收费员收费口','1','2020-05-11 09:15:35','2020-05-11 09:15:35');
insert  into `t_dmmx`(`dm`,`dmm`,`dmmc`,`flag`,`create_time`,`update_time`) values ('08','2','实时客流','1','2020-05-11 09:15:35','2020-05-11 09:15:35');
insert  into `t_dmmx`(`dm`,`dmm`,`dmmc`,`flag`,`create_time`,`update_time`) values ('08','3','离线车流','1','2020-05-11 09:15:35','2020-05-11 09:15:35');
insert  into `t_dmmx`(`dm`,`dmm`,`dmmc`,`flag`,`create_time`,`update_time`) values ('08','4','实时车流','1','2020-05-11 09:15:35','2020-05-11 09:15:35');
insert  into `t_dmmx`(`dm`,`dmm`,`dmmc`,`flag`,`create_time`,`update_time`) values ('08','5','客流设备','1','2020-05-11 09:15:35','2020-05-11 09:15:35');
insert  into `t_dmmx`(`dm`,`dmm`,`dmmc`,`flag`,`create_time`,`update_time`) values ('08','6','反向寻车','1','2020-05-11 09:15:35','2020-05-11 09:15:35');
insert  into `t_dmmx`(`dm`,`dmm`,`dmmc`,`flag`,`create_time`,`update_time`) values ('08','7','收费员结算','1','2020-05-11 09:15:35','2020-05-11 09:15:35');
insert  into `t_dmmx`(`dm`,`dmm`,`dmmc`,`flag`,`create_time`,`update_time`) values ('08','8','业绩上报','1','2020-05-11 09:15:35','2020-05-11 09:15:35');
insert  into `t_dmmx`(`dm`,`dmm`,`dmmc`,`flag`,`create_time`,`update_time`) values ('08','9','会员数据','1','2020-05-11 09:15:35','2020-05-11 09:15:35');
insert  into `t_dmmx`(`dm`,`dmm`,`dmmc`,`flag`,`create_time`,`update_time`) values ('09','1','mssql->mysql','1','2020-05-11 09:15:35','2020-05-11 09:15:35');
insert  into `t_dmmx`(`dm`,`dmm`,`dmmc`,`flag`,`create_time`,`update_time`) values ('09','2','mysql->mysql','1','2020-05-11 09:15:35','2020-05-11 09:15:35');
insert  into `t_dmmx`(`dm`,`dmm`,`dmmc`,`flag`,`create_time`,`update_time`) values ('09','3','mssql->mssql','1','2020-05-11 09:15:35','2020-05-11 09:15:35');
insert  into `t_dmmx`(`dm`,`dmm`,`dmmc`,`flag`,`create_time`,`update_time`) values ('09','4','mongo->mongo','1','2020-05-11 09:15:35','2020-05-11 09:15:35');
insert  into `t_dmmx`(`dm`,`dmm`,`dmmc`,`flag`,`create_time`,`update_time`) values ('09','5','mysql->hbase','1','2020-05-11 09:15:35','2020-05-11 09:15:35');
insert  into `t_dmmx`(`dm`,`dmm`,`dmmc`,`flag`,`create_time`,`update_time`) values ('09','6','mysql->elasticsearch','1','2020-10-26 14:50:49','2020-10-26 14:50:52');
insert  into `t_dmmx`(`dm`,`dmm`,`dmmc`,`flag`,`create_time`,`update_time`) values ('10','day','天','1','2020-05-11 09:15:35','2020-05-11 09:15:35');
insert  into `t_dmmx`(`dm`,`dmm`,`dmmc`,`flag`,`create_time`,`update_time`) values ('10','hour','小时','1','2020-05-11 09:15:35','2020-05-11 09:15:35');
insert  into `t_dmmx`(`dm`,`dmm`,`dmmc`,`flag`,`create_time`,`update_time`) values ('10','min','分','1','2020-05-11 09:15:35','2020-05-11 09:15:35');
insert  into `t_dmmx`(`dm`,`dmm`,`dmmc`,`flag`,`create_time`,`update_time`) values ('11','1','备份数据源','1','2020-05-11 09:15:35','2020-05-11 09:15:35');
insert  into `t_dmmx`(`dm`,`dmm`,`dmmc`,`flag`,`create_time`,`update_time`) values ('11','2','同步数据源','1','2020-05-11 09:15:35','2020-05-11 09:15:35');
insert  into `t_dmmx`(`dm`,`dmm`,`dmmc`,`flag`,`create_time`,`update_time`) values ('12','1','V3.7.5','1','2020-05-11 09:15:35','2020-05-11 09:15:35');
insert  into `t_dmmx`(`dm`,`dmm`,`dmmc`,`flag`,`create_time`,`update_time`) values ('12','2','V3.7.6','1','2020-05-11 09:15:35','2020-05-11 09:15:35');
insert  into `t_dmmx`(`dm`,`dmm`,`dmmc`,`flag`,`create_time`,`update_time`) values ('13','1','DDL工单','1','2020-05-11 09:15:35','2020-05-11 09:15:35');
insert  into `t_dmmx`(`dm`,`dmm`,`dmmc`,`flag`,`create_time`,`update_time`) values ('13','2','DML工单','1','2020-05-11 09:15:35','2020-05-11 09:15:35');
insert  into `t_dmmx`(`dm`,`dmm`,`dmmc`,`flag`,`create_time`,`update_time`) values ('13','3','DCL工单','1','2020-05-11 09:15:35','2020-05-11 09:15:35');
insert  into `t_dmmx`(`dm`,`dmm`,`dmmc`,`flag`,`create_time`,`update_time`) values ('13','4','存储过程','1','2020-05-11 09:15:35','2020-05-11 09:15:35');
insert  into `t_dmmx`(`dm`,`dmm`,`dmmc`,`flag`,`create_time`,`update_time`) values ('15','10.2.39.165:2181,10.2.39.166:2181,10.2.39.182:2181','大数据开发zookeeper地址','1','2020-05-11 09:15:35','2020-05-11 09:15:35');
insert  into `t_dmmx`(`dm`,`dmm`,`dmmc`,`flag`,`create_time`,`update_time`) values ('15','10.2.39.84:2181,10.2.39.89:2181,10.2.39.67:2181','大数据测试zookeeper地址','1','2020-05-11 09:15:35','2020-05-11 09:15:35');
insert  into `t_dmmx`(`dm`,`dmm`,`dmmc`,`flag`,`create_time`,`update_time`) values ('15','192.168.100.63:2181,192.168.100.64:2181,192.168.100.69:2181','大数据生产zookeeper地址','1','2020-05-11 09:15:35','2020-05-11 09:15:35');
insert  into `t_dmmx`(`dm`,`dmm`,`dmmc`,`flag`,`create_time`,`update_time`) values ('16','10.2.39.165:9090','大数据开发thrift接口地址','1','2020-05-11 09:15:35','2020-05-11 09:15:35');
insert  into `t_dmmx`(`dm`,`dmm`,`dmmc`,`flag`,`create_time`,`update_time`) values ('16','10.2.39.84:9090','大数据测试thrift接口地址','1','2020-05-11 09:15:35','2020-05-11 09:15:35');
insert  into `t_dmmx`(`dm`,`dmm`,`dmmc`,`flag`,`create_time`,`update_time`) values ('16','192.168.100.63:9090','大数据生产thrift接口地址','1','2020-05-11 09:15:35','2020-05-11 09:15:35');
insert  into `t_dmmx`(`dm`,`dmm`,`dmmc`,`flag`,`create_time`,`update_time`) values ('17','1','数据传输','1','2020-05-11 09:15:35','2020-05-11 09:15:35');
insert  into `t_dmmx`(`dm`,`dmm`,`dmmc`,`flag`,`create_time`,`update_time`) values ('17','2','数据同步','1','2020-05-11 09:15:35','2020-05-11 09:15:35');
insert  into `t_dmmx`(`dm`,`dmm`,`dmmc`,`flag`,`create_time`,`update_time`) values ('17','3','账号权限','1','2020-05-11 09:15:35','2020-05-11 09:15:35');
insert  into `t_dmmx`(`dm`,`dmm`,`dmmc`,`flag`,`create_time`,`update_time`) values ('17','4','数据迁移','1','2020-05-11 09:15:35','2020-05-11 09:15:35');
insert  into `t_dmmx`(`dm`,`dmm`,`dmmc`,`flag`,`create_time`,`update_time`) values ('17','5','主从同步','1','2020-05-11 09:15:35','2020-05-11 09:15:35');
insert  into `t_dmmx`(`dm`,`dmm`,`dmmc`,`flag`,`create_time`,`update_time`) values ('17','6','数据备份','1','2020-05-11 09:15:35','2020-05-11 09:15:35');
insert  into `t_dmmx`(`dm`,`dmm`,`dmmc`,`flag`,`create_time`,`update_time`) values ('17','7','查询支持  ','1','2020-05-11 09:15:35','2020-05-11 09:15:35');
insert  into `t_dmmx`(`dm`,`dmm`,`dmmc`,`flag`,`create_time`,`update_time`) values ('17','8','查询优化','1','2020-05-11 09:15:35','2020-05-11 09:15:35');
insert  into `t_dmmx`(`dm`,`dmm`,`dmmc`,`flag`,`create_time`,`update_time`) values ('17','9','其它类型','1','2020-05-11 09:15:35','2020-05-11 09:15:35');
insert  into `t_dmmx`(`dm`,`dmm`,`dmmc`,`flag`,`create_time`,`update_time`) values ('18','1','合生通','1','2020-05-11 09:15:35','2020-05-11 09:15:35');
insert  into `t_dmmx`(`dm`,`dmm`,`dmmc`,`flag`,`create_time`,`update_time`) values ('18','2','好房','1','2020-05-11 09:15:35','2020-05-11 09:15:35');
insert  into `t_dmmx`(`dm`,`dmm`,`dmmc`,`flag`,`create_time`,`update_time`) values ('18','3','好生活','1','2020-05-11 09:15:35','2020-05-11 09:15:35');
insert  into `t_dmmx`(`dm`,`dmm`,`dmmc`,`flag`,`create_time`,`update_time`) values ('19','1','已新增','1','2020-05-11 09:15:35','2020-05-11 09:15:35');
insert  into `t_dmmx`(`dm`,`dmm`,`dmmc`,`flag`,`create_time`,`update_time`) values ('19','2','已发布','1','2020-05-11 09:15:35','2020-05-11 09:15:35');
insert  into `t_dmmx`(`dm`,`dmm`,`dmmc`,`flag`,`create_time`,`update_time`) values ('19','3','处理中','1','2020-05-11 09:15:35','2020-05-11 09:15:35');
insert  into `t_dmmx`(`dm`,`dmm`,`dmmc`,`flag`,`create_time`,`update_time`) values ('19','4','已完成','1','2020-05-11 09:15:35','2020-05-11 09:15:35');
insert  into `t_dmmx`(`dm`,`dmm`,`dmmc`,`flag`,`create_time`,`update_time`) values ('19','5','已驳回','1','2020-05-11 09:15:35','2020-05-11 09:15:35');
insert  into `t_dmmx`(`dm`,`dmm`,`dmmc`,`flag`,`create_time`,`update_time`) values ('19','6','已转派','1','2020-05-11 09:15:35','2020-05-11 09:15:35');
insert  into `t_dmmx`(`dm`,`dmm`,`dmmc`,`flag`,`create_time`,`update_time`) values ('19','7','已撤销','1','2020-05-11 09:15:35','2020-05-11 09:15:35');
insert  into `t_dmmx`(`dm`,`dmm`,`dmmc`,`flag`,`create_time`,`update_time`) values ('19','8','已关闭','1','2020-05-11 09:15:35','2020-05-11 09:15:35');
insert  into `t_dmmx`(`dm`,`dmm`,`dmmc`,`flag`,`create_time`,`update_time`) values ('20','1','hour','1','2020-05-11 09:15:35','2020-05-11 09:15:35');
insert  into `t_dmmx`(`dm`,`dmm`,`dmmc`,`flag`,`create_time`,`update_time`) values ('20','2','day','1','2020-05-11 09:15:35','2020-05-11 09:15:35');
insert  into `t_dmmx`(`dm`,`dmm`,`dmmc`,`flag`,`create_time`,`update_time`) values ('20','3','month','1','2020-05-11 09:15:35','2020-05-11 09:15:35');
insert  into `t_dmmx`(`dm`,`dmm`,`dmmc`,`flag`,`create_time`,`update_time`) values ('21','1','删除','1','2020-05-11 09:15:35','2020-05-11 09:15:35');
insert  into `t_dmmx`(`dm`,`dmm`,`dmmc`,`flag`,`create_time`,`update_time`) values ('21','2','迁移','1','2020-05-11 09:15:35','2020-05-11 09:15:35');
insert  into `t_dmmx`(`dm`,`dmm`,`dmmc`,`flag`,`create_time`,`update_time`) values ('22','1','周期','1','2020-05-11 09:15:35','2020-05-11 09:15:35');
insert  into `t_dmmx`(`dm`,`dmm`,`dmmc`,`flag`,`create_time`,`update_time`) values ('22','2','范围','1','2020-05-11 09:15:35','2020-05-11 09:15:35');
insert  into `t_dmmx`(`dm`,`dmm`,`dmmc`,`flag`,`create_time`,`update_time`) values ('23','1','服务器','1','2020-05-11 09:15:35','2020-05-11 09:15:35');
insert  into `t_dmmx`(`dm`,`dmm`,`dmmc`,`flag`,`create_time`,`update_time`) values ('23','2','数据库','1','2020-05-11 09:15:35','2020-05-11 09:15:35');
insert  into `t_dmmx`(`dm`,`dmm`,`dmmc`,`flag`,`create_time`,`update_time`) values ('24','1','百分比','1','2020-05-11 09:15:35','2020-05-11 09:15:35');
insert  into `t_dmmx`(`dm`,`dmm`,`dmmc`,`flag`,`create_time`,`update_time`) values ('24','2','计算型','1','2020-05-11 09:15:35','2020-05-11 09:15:35');
insert  into `t_dmmx`(`dm`,`dmm`,`dmmc`,`flag`,`create_time`,`update_time`) values ('24','3','标量值','1','2020-05-11 09:15:35','2020-05-11 09:15:35');
insert  into `t_dmmx`(`dm`,`dmm`,`dmmc`,`flag`,`create_time`,`update_time`) values ('25','1','正常','1','2020-05-14 18:57:14','2020-05-14 18:57:20');
insert  into `t_dmmx`(`dm`,`dmm`,`dmmc`,`flag`,`create_time`,`update_time`) values ('25','2','	停用','1','2020-05-14 18:57:17','2020-05-14 18:57:22');
insert  into `t_dmmx`(`dm`,`dmm`,`dmmc`,`flag`,`create_time`,`update_time`) values ('25','3','锁定','1','2020-05-15 11:27:57','2020-05-15 11:27:57');
insert  into `t_dmmx`(`dm`,`dmm`,`dmmc`,`flag`,`create_time`,`update_time`) values ('26','10.16.33.10:8888','北京合生广场代理服务','1','2020-11-12 10:39:41','2020-11-12 10:39:44');
insert  into `t_dmmx`(`dm`,`dmm`,`dmmc`,`flag`,`create_time`,`update_time`) values ('26','10.16.34.14:8888','北京朝阳合生汇代理服务','1','2020-11-02 14:26:59','2020-11-02 14:27:01');
insert  into `t_dmmx`(`dm`,`dmm`,`dmmc`,`flag`,`create_time`,`update_time`) values ('26','10.19.53.205:8888','广州南方花园代理服务','1','2020-11-12 10:53:08','2020-11-12 10:53:12');
insert  into `t_dmmx`(`dm`,`dmm`,`dmmc`,`flag`,`create_time`,`update_time`) values ('26','10.19.72.205:8888','广州合生骏景广场代理服务','1','2020-11-12 10:53:05','2020-11-12 10:53:10');
insert  into `t_dmmx`(`dm`,`dmm`,`dmmc`,`flag`,`create_time`,`update_time`) values ('26','10.25.128.84:8888','广州嘉和南代理服务','1','2020-07-28 17:11:15','2020-07-28 17:11:53');
insert  into `t_dmmx`(`dm`,`dmm`,`dmmc`,`flag`,`create_time`,`update_time`) values ('26','10.28.8.233:8888','成都珠江广场代理服务','1','2020-11-12 10:20:44','2020-11-12 10:20:46');
insert  into `t_dmmx`(`dm`,`dmm`,`dmmc`,`flag`,`create_time`,`update_time`) values ('26','10.5.232.214:8888','广州增城合生汇代理服务','1','2020-11-12 14:16:18','2020-11-12 14:16:20');
insert  into `t_dmmx`(`dm`,`dmm`,`dmmc`,`flag`,`create_time`,`update_time`) values ('26','10.5.3.26:8888','上海青浦米格代理服务','1','2020-11-12 10:29:51','2020-11-12 10:29:53');
insert  into `t_dmmx`(`dm`,`dmm`,`dmmc`,`flag`,`create_time`,`update_time`) values ('26','10.61.101.241:8888','上海五角场代理服务','1','2020-11-05 10:29:46','2020-11-05 10:29:48');
insert  into `t_dmmx`(`dm`,`dmm`,`dmmc`,`flag`,`create_time`,`update_time`) values ('26','39.106.226.140:8888','合生通生产业务库代理服务','1','2020-11-18 16:31:40','2020-11-18 16:31:42');
insert  into `t_dmmx`(`dm`,`dmm`,`dmmc`,`flag`,`create_time`,`update_time`) values ('26','47.94.197.105:8888','阿里云中转库代理服务','1','2020-11-12 14:33:15','2020-11-12 14:33:17');
insert  into `t_dmmx`(`dm`,`dmm`,`dmmc`,`flag`,`create_time`,`update_time`) values ('26','bjcfpark.hopsontong.com:60138','北京合生财富广场代理服务','1','2020-11-21 13:28:05','2020-11-21 13:28:07');
insert  into `t_dmmx`(`dm`,`dmm`,`dmmc`,`flag`,`create_time`,`update_time`) values ('26','bjdespark.hopsontong.com:60128','北京德胜合生财富广场代理服务','1','2020-11-21 13:28:47','2020-11-21 13:28:50');
insert  into `t_dmmx`(`dm`,`dmm`,`dmmc`,`flag`,`create_time`,`update_time`) values ('26','bjwjqlspark.hopsontong.com:60039','合生新天地代理服务','1','2020-11-21 13:25:55','2020-11-21 13:25:58');
insert  into `t_dmmx`(`dm`,`dmm`,`dmmc`,`flag`,`create_time`,`update_time`) values ('26','gzyhpark.hopsontong.com:60088','广州越华珠江广场代理服务','1','2020-11-21 13:26:30','2020-11-21 13:26:33');
insert  into `t_dmmx`(`dm`,`dmm`,`dmmc`,`flag`,`create_time`,`update_time`) values ('26','hzkhpark.hopsontong.com:60016','科华数码广场有限公司代理服务','1','2020-11-21 13:30:18','2020-11-21 13:30:20');
insert  into `t_dmmx`(`dm`,`dmm`,`dmmc`,`flag`,`create_time`,`update_time`) values ('26','shhytpark.hopsontong.com:60118','上海海云天合生财富广场代理服务','1','2020-11-21 13:29:23','2020-11-21 13:29:25');
insert  into `t_dmmx`(`dm`,`dmm`,`dmmc`,`flag`,`create_time`,`update_time`) values ('26','shjazjpark.hopsontong.com:60158','上海静安珠江创意中心代理服务','1','2020-11-21 13:29:45','2020-11-21 13:29:47');
insert  into `t_dmmx`(`dm`,`dmm`,`dmmc`,`flag`,`create_time`,`update_time`) values ('27','1','5.6','1','2020-07-29 13:59:16','2020-07-29 13:59:16');
insert  into `t_dmmx`(`dm`,`dmm`,`dmmc`,`flag`,`create_time`,`update_time`) values ('27','2','5.7','1','2020-07-29 13:59:22','2020-07-29 13:59:22');
insert  into `t_dmmx`(`dm`,`dmm`,`dmmc`,`flag`,`create_time`,`update_time`) values ('28','1','已就绪','1','2020-07-30 17:20:29','2020-07-30 17:20:29');
insert  into `t_dmmx`(`dm`,`dmm`,`dmmc`,`flag`,`create_time`,`update_time`) values ('28','2','运行中','1','2020-07-30 17:20:35','2020-07-30 17:20:35');
insert  into `t_dmmx`(`dm`,`dmm`,`dmmc`,`flag`,`create_time`,`update_time`) values ('28','3','已完成','1','2020-07-30 17:20:42','2020-07-30 17:20:42');
insert  into `t_dmmx`(`dm`,`dmm`,`dmmc`,`flag`,`create_time`,`update_time`) values ('28','4','运行失败','1','2020-07-30 17:20:49','2020-07-30 17:20:49');
insert  into `t_dmmx`(`dm`,`dmm`,`dmmc`,`flag`,`create_time`,`update_time`) values ('30','basedir=/usr/local/mysql5.6.44','mysql服务主目录','1','2020-08-06 14:38:30','2020-08-24 14:08:32');
insert  into `t_dmmx`(`dm`,`dmm`,`dmmc`,`flag`,`create_time`,`update_time`) values ('30','binlog-format=row','binlog日志格式','1','2020-08-06 14:37:59','2020-08-06 14:37:59');
insert  into `t_dmmx`(`dm`,`dmm`,`dmmc`,`flag`,`create_time`,`update_time`) values ('30','character-set-server=utf8mb4','mysql服务端字符集','1','2020-08-06 14:40:02','2020-08-06 14:40:02');
insert  into `t_dmmx`(`dm`,`dmm`,`dmmc`,`flag`,`create_time`,`update_time`) values ('30','datadir=/home/hopson/apps/usr/webserver/mysql/data/{}/{}','mysql数据主目录','1','2020-08-06 14:38:44','2020-08-31 16:28:49');
insert  into `t_dmmx`(`dm`,`dmm`,`dmmc`,`flag`,`create_time`,`update_time`) values ('30','default-character-set=utf8mb4','客户端字符集','1','2020-08-24 13:33:16','2020-08-24 13:33:16');
insert  into `t_dmmx`(`dm`,`dmm`,`dmmc`,`flag`,`create_time`,`update_time`) values ('30','expire_logs_days=7','binlog保留时间(天)','1','2020-08-06 14:37:12','2020-08-06 14:37:12');
insert  into `t_dmmx`(`dm`,`dmm`,`dmmc`,`flag`,`create_time`,`update_time`) values ('30','innodb_buffer_pool_size=4G','innodb缓存池大小','1','2020-08-06 14:44:47','2020-08-06 14:44:47');
insert  into `t_dmmx`(`dm`,`dmm`,`dmmc`,`flag`,`create_time`,`update_time`) values ('30','innodb_log_buffer_size=8m','innodb日志缓存池大小','1','2020-08-06 14:45:00','2020-08-06 14:45:00');
insert  into `t_dmmx`(`dm`,`dmm`,`dmmc`,`flag`,`create_time`,`update_time`) values ('30','innodb_log_file_size=512m','innodb日志文件大小','1','2020-08-06 14:45:11','2020-08-06 14:45:11');
insert  into `t_dmmx`(`dm`,`dmm`,`dmmc`,`flag`,`create_time`,`update_time`) values ('30','interactive_timeout=86400','交互等待时间','1','2020-09-30 08:12:43','2020-09-30 08:12:43');
insert  into `t_dmmx`(`dm`,`dmm`,`dmmc`,`flag`,`create_time`,`update_time`) values ('30','log-bin=mysql-bin','mysql binlog日志名称','1','2020-08-06 14:36:53','2020-08-06 14:36:53');
insert  into `t_dmmx`(`dm`,`dmm`,`dmmc`,`flag`,`create_time`,`update_time`) values ('30','log-error=/home/hopson/apps/usr/webserver/mysql/data/{}/{}/mysql.err','错误日志','1','2020-08-31 16:32:35','2020-08-31 16:52:23');
insert  into `t_dmmx`(`dm`,`dmm`,`dmmc`,`flag`,`create_time`,`update_time`) values ('30','log_bin_trust_function_creators=on','开启日志后是否能创建存储过程','1','2020-08-06 14:36:32','2020-09-03 11:12:05');
insert  into `t_dmmx`(`dm`,`dmm`,`dmmc`,`flag`,`create_time`,`update_time`) values ('30','lower_case_table_names=1','mysql是否区分大小写','1','2020-08-24 13:30:25','2020-08-24 13:30:25');
insert  into `t_dmmx`(`dm`,`dmm`,`dmmc`,`flag`,`create_time`,`update_time`) values ('30','max_allowed_packet=512M','最大网络包大小','1','2020-08-06 14:43:14','2020-09-03 11:35:29');
insert  into `t_dmmx`(`dm`,`dmm`,`dmmc`,`flag`,`create_time`,`update_time`) values ('30','max_connections=2000','最大连接数','1','2020-08-06 14:44:30','2020-08-06 14:44:30');
insert  into `t_dmmx`(`dm`,`dmm`,`dmmc`,`flag`,`create_time`,`update_time`) values ('30','pid-file=/home/hopson/apps/usr/webserver/mysql/data/{}/{}/mysql{}.pid','mysql进程文件','1','2020-09-01 16:13:33','2020-09-01 16:13:33');
insert  into `t_dmmx`(`dm`,`dmm`,`dmmc`,`flag`,`create_time`,`update_time`) values ('30','port={}','端口号','1','2020-08-06 14:39:35','2020-09-01 16:25:16');
insert  into `t_dmmx`(`dm`,`dmm`,`dmmc`,`flag`,`create_time`,`update_time`) values ('30','server_id=1','mysql服务id','1','2020-08-06 14:36:05','2020-08-06 14:36:05');
insert  into `t_dmmx`(`dm`,`dmm`,`dmmc`,`flag`,`create_time`,`update_time`) values ('30','socket=/home/hopson/apps/usr/webserver/mysql/data/{}/{}/mysql.sock','socket文件名','1','2020-08-06 14:40:33','2020-08-31 16:26:02');
insert  into `t_dmmx`(`dm`,`dmm`,`dmmc`,`flag`,`create_time`,`update_time`) values ('30','sql_mode=NO_ENGINE_SUBSTITUTION,STRICT_TRANS_TABLES','sql模式','1','2020-08-06 14:40:56','2020-08-06 14:40:56');
insert  into `t_dmmx`(`dm`,`dmm`,`dmmc`,`flag`,`create_time`,`update_time`) values ('30','symbolic-links=0','禁用符号链接','1','2020-08-06 14:39:21','2020-08-06 14:39:21');
insert  into `t_dmmx`(`dm`,`dmm`,`dmmc`,`flag`,`create_time`,`update_time`) values ('30','table_definition_cache=16384','表定义缓存的大小','1','2020-08-06 14:44:16','2020-08-06 14:44:16');
insert  into `t_dmmx`(`dm`,`dmm`,`dmmc`,`flag`,`create_time`,`update_time`) values ('30','table_open_cache=16384','表高速缓存的大小','1','2020-08-06 14:43:57','2020-08-06 14:43:57');
insert  into `t_dmmx`(`dm`,`dmm`,`dmmc`,`flag`,`create_time`,`update_time`) values ('30','transaction_isolation =\'READ-COMMITTED\'','事务隔离级别','1','2020-08-06 14:42:48','2020-08-06 14:42:48');
insert  into `t_dmmx`(`dm`,`dmm`,`dmmc`,`flag`,`create_time`,`update_time`) values ('30','user=hopson','mysql服务启动用户','1','2020-08-06 14:39:05','2020-08-06 14:39:05');
insert  into `t_dmmx`(`dm`,`dmm`,`dmmc`,`flag`,`create_time`,`update_time`) values ('30','wait_timeout=86400','锁等待时间','1','2020-09-30 08:12:31','2020-09-30 08:12:31');
insert  into `t_dmmx`(`dm`,`dmm`,`dmmc`,`flag`,`create_time`,`update_time`) values ('31','create','创建对象','1','2020-08-06 14:39:05','2020-08-06 14:39:05');
insert  into `t_dmmx`(`dm`,`dmm`,`dmmc`,`flag`,`create_time`,`update_time`) values ('31','delete','删除操作','1','2020-08-06 14:39:05','2020-08-06 14:39:05');
insert  into `t_dmmx`(`dm`,`dmm`,`dmmc`,`flag`,`create_time`,`update_time`) values ('31','drop','删除对象','1','2020-08-06 14:39:05','2020-08-06 14:39:05');
insert  into `t_dmmx`(`dm`,`dmm`,`dmmc`,`flag`,`create_time`,`update_time`) values ('31','insert','插入操作','1','2020-08-06 14:39:05','2020-08-06 14:39:05');
insert  into `t_dmmx`(`dm`,`dmm`,`dmmc`,`flag`,`create_time`,`update_time`) values ('31','select','查询操作','1','2020-08-06 14:39:05','2020-08-06 14:39:05');
insert  into `t_dmmx`(`dm`,`dmm`,`dmmc`,`flag`,`create_time`,`update_time`) values ('31','update','更新操作','1','2020-08-06 14:39:05','2020-08-06 14:39:05');
insert  into `t_dmmx`(`dm`,`dmm`,`dmmc`,`flag`,`create_time`,`update_time`) values ('32','1','未创建','1','2020-08-06 14:39:05','2020-08-06 14:39:05');
insert  into `t_dmmx`(`dm`,`dmm`,`dmmc`,`flag`,`create_time`,`update_time`) values ('32','2','已创建','1','2020-08-06 14:39:05','2020-08-06 14:39:05');
insert  into `t_dmmx`(`dm`,`dmm`,`dmmc`,`flag`,`create_time`,`update_time`) values ('32','3','已运行','1','2020-08-06 14:39:05','2020-08-06 14:39:05');
insert  into `t_dmmx`(`dm`,`dmm`,`dmmc`,`flag`,`create_time`,`update_time`) values ('32','4','已停止','1','2020-08-06 14:39:05','2020-08-06 14:39:05');
insert  into `t_dmmx`(`dm`,`dmm`,`dmmc`,`flag`,`create_time`,`update_time`) values ('32','5','已销毁','1','2020-09-02 09:50:25','2020-09-02 09:50:25');
insert  into `t_dmmx`(`dm`,`dmm`,`dmmc`,`flag`,`create_time`,`update_time`) values ('33','http://10.2.39.18/downloads/mysql-5.7.27-linux-glibc2.12-x86_64.tar','mysql5.7_download_url','1','2020-08-19 17:41:51','2020-08-19 17:41:51');
insert  into `t_dmmx`(`dm`,`dmm`,`dmmc`,`flag`,`create_time`,`update_time`) values ('33','http://124.127.103.190:65480/downloads/mysql-5.6.44-linux-glibc2.12-x86_64.tar.gz','mysql5.6_download_url','1','2020-08-19 17:42:09','2020-09-07 10:36:11');
insert  into `t_dmmx`(`dm`,`dmm`,`dmmc`,`flag`,`create_time`,`update_time`) values ('34','1','目录同步','1','2020-09-21 11:07:44','2020-09-21 11:07:44');
insert  into `t_dmmx`(`dm`,`dmm`,`dmmc`,`flag`,`create_time`,`update_time`) values ('34','2','服务同步','1','2020-09-21 11:07:48','2020-09-21 11:07:48');
insert  into `t_dmmx`(`dm`,`dmm`,`dmmc`,`flag`,`create_time`,`update_time`) values ('35','1','数据库代理','1','2020-11-13 11:32:19','2020-11-13 11:32:19');
insert  into `t_dmmx`(`dm`,`dmm`,`dmmc`,`flag`,`create_time`,`update_time`) values ('36','01','8888','1','2020-11-21 13:38:26','2020-11-21 13:38:28');
insert  into `t_dmmx`(`dm`,`dmm`,`dmmc`,`flag`,`create_time`,`update_time`) values ('37','1','进程','1','2021-02-19 16:42:49','2021-02-19 16:42:49');
insert  into `t_dmmx`(`dm`,`dmm`,`dmmc`,`flag`,`create_time`,`update_time`) values ('37','2','慢日志','1','2021-02-19 16:42:49','2021-02-19 16:42:49');
insert  into `t_dmmx`(`dm`,`dmm`,`dmmc`,`flag`,`create_time`,`update_time`) values ('38','1','自建库','1','2021-04-16 15:55:29','2021-04-16 15:55:33');
insert  into `t_dmmx`(`dm`,`dmm`,`dmmc`,`flag`,`create_time`,`update_time`) values ('38','2','数据源','1','2021-04-16 15:55:31','2021-04-16 15:55:35');

/*Table structure for table `t_role` */

CREATE TABLE `t_role` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(20) NOT NULL DEFAULT '',
  `status` varchar(1) NOT NULL DEFAULT '',
  `creator` varchar(10) NOT NULL DEFAULT '',
  `updater` varchar(10) NOT NULL DEFAULT '',
  `create_date` datetime NOT NULL COMMENT '创建时间',
  `last_update_date` datetime NOT NULL COMMENT '修改时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=54 DEFAULT CHARSET=utf8mb4;

/*Data for the table `t_role` */

insert  into `t_role`(`id`,`name`,`status`,`creator`,`updater`,`create_date`,`last_update_date`) values (33,'系统管理员','1','DBA','DBA','2021-05-11 16:19:23','2021-05-11 16:19:23');
insert  into `t_role`(`id`,`name`,`status`,`creator`,`updater`,`create_date`,`last_update_date`) values (42,'数据库管理员','1','DBA','DBA','2021-05-11 16:23:16','2021-05-11 16:23:16');
insert  into `t_role`(`id`,`name`,`status`,`creator`,`updater`,`create_date`,`last_update_date`) values (45,'大数据开发','1','DBA','DBA','2021-05-11 16:30:49','2021-05-11 16:30:49');
insert  into `t_role`(`id`,`name`,`status`,`creator`,`updater`,`create_date`,`last_update_date`) values (48,'开发人员	','1','DBA','DBA','2021-05-11 16:32:02','2021-05-11 16:32:02');

/*Table structure for table `t_role_privs` */

CREATE TABLE `t_role_privs` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `role_id` int(11) NOT NULL DEFAULT '0',
  `priv_id` varchar(11) NOT NULL DEFAULT '0',
  `creator` varchar(10) NOT NULL DEFAULT '',
  `updater` varchar(10) NOT NULL DEFAULT '',
  `create_date` datetime NOT NULL COMMENT '创建时间',
  `last_update_date` datetime NOT NULL COMMENT '修改时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=648 DEFAULT CHARSET=utf8mb4;

/*Data for the table `t_role_privs` */

insert  into `t_role_privs`(`id`,`role_id`,`priv_id`,`creator`,`updater`,`create_date`,`last_update_date`) values (123,33,'00101','DBA','DBA','2021-05-11 16:19:24','2021-05-11 16:19:24');
insert  into `t_role_privs`(`id`,`role_id`,`priv_id`,`creator`,`updater`,`create_date`,`last_update_date`) values (126,33,'00102','DBA','DBA','2021-05-11 16:19:24','2021-05-11 16:19:24');
insert  into `t_role_privs`(`id`,`role_id`,`priv_id`,`creator`,`updater`,`create_date`,`last_update_date`) values (129,33,'00103','DBA','DBA','2021-05-11 16:19:24','2021-05-11 16:19:24');
insert  into `t_role_privs`(`id`,`role_id`,`priv_id`,`creator`,`updater`,`create_date`,`last_update_date`) values (132,33,'00201','DBA','DBA','2021-05-11 16:19:24','2021-05-11 16:19:24');
insert  into `t_role_privs`(`id`,`role_id`,`priv_id`,`creator`,`updater`,`create_date`,`last_update_date`) values (135,33,'00202','DBA','DBA','2021-05-11 16:19:24','2021-05-11 16:19:24');
insert  into `t_role_privs`(`id`,`role_id`,`priv_id`,`creator`,`updater`,`create_date`,`last_update_date`) values (138,33,'00203','DBA','DBA','2021-05-11 16:19:24','2021-05-11 16:19:24');
insert  into `t_role_privs`(`id`,`role_id`,`priv_id`,`creator`,`updater`,`create_date`,`last_update_date`) values (141,33,'00301','DBA','DBA','2021-05-11 16:19:25','2021-05-11 16:19:25');
insert  into `t_role_privs`(`id`,`role_id`,`priv_id`,`creator`,`updater`,`create_date`,`last_update_date`) values (144,33,'00302','DBA','DBA','2021-05-11 16:19:25','2021-05-11 16:19:25');
insert  into `t_role_privs`(`id`,`role_id`,`priv_id`,`creator`,`updater`,`create_date`,`last_update_date`) values (147,33,'00303','DBA','DBA','2021-05-11 16:19:25','2021-05-11 16:19:25');
insert  into `t_role_privs`(`id`,`role_id`,`priv_id`,`creator`,`updater`,`create_date`,`last_update_date`) values (150,33,'00501','DBA','DBA','2021-05-11 16:19:25','2021-05-11 16:19:25');
insert  into `t_role_privs`(`id`,`role_id`,`priv_id`,`creator`,`updater`,`create_date`,`last_update_date`) values (153,33,'00502','DBA','DBA','2021-05-11 16:19:25','2021-05-11 16:19:25');
insert  into `t_role_privs`(`id`,`role_id`,`priv_id`,`creator`,`updater`,`create_date`,`last_update_date`) values (156,33,'00503','DBA','DBA','2021-05-11 16:19:25','2021-05-11 16:19:25');
insert  into `t_role_privs`(`id`,`role_id`,`priv_id`,`creator`,`updater`,`create_date`,`last_update_date`) values (159,33,'00601','DBA','DBA','2021-05-11 16:19:26','2021-05-11 16:19:26');
insert  into `t_role_privs`(`id`,`role_id`,`priv_id`,`creator`,`updater`,`create_date`,`last_update_date`) values (162,33,'01901','DBA','DBA','2021-05-11 16:19:26','2021-05-11 16:19:26');
insert  into `t_role_privs`(`id`,`role_id`,`priv_id`,`creator`,`updater`,`create_date`,`last_update_date`) values (165,33,'01902','DBA','DBA','2021-05-11 16:19:26','2021-05-11 16:19:26');
insert  into `t_role_privs`(`id`,`role_id`,`priv_id`,`creator`,`updater`,`create_date`,`last_update_date`) values (168,33,'01903','DBA','DBA','2021-05-11 16:19:26','2021-05-11 16:19:26');
insert  into `t_role_privs`(`id`,`role_id`,`priv_id`,`creator`,`updater`,`create_date`,`last_update_date`) values (171,33,'02002','DBA','DBA','2021-05-11 16:19:26','2021-05-11 16:19:26');
insert  into `t_role_privs`(`id`,`role_id`,`priv_id`,`creator`,`updater`,`create_date`,`last_update_date`) values (174,33,'02003','DBA','DBA','2021-05-11 16:19:27','2021-05-11 16:19:27');
insert  into `t_role_privs`(`id`,`role_id`,`priv_id`,`creator`,`updater`,`create_date`,`last_update_date`) values (177,33,'02101','DBA','DBA','2021-05-11 16:19:27','2021-05-11 16:19:27');
insert  into `t_role_privs`(`id`,`role_id`,`priv_id`,`creator`,`updater`,`create_date`,`last_update_date`) values (180,33,'02102','DBA','DBA','2021-05-11 16:19:27','2021-05-11 16:19:27');
insert  into `t_role_privs`(`id`,`role_id`,`priv_id`,`creator`,`updater`,`create_date`,`last_update_date`) values (183,33,'02103','DBA','DBA','2021-05-11 16:19:27','2021-05-11 16:19:27');
insert  into `t_role_privs`(`id`,`role_id`,`priv_id`,`creator`,`updater`,`create_date`,`last_update_date`) values (186,33,'02104','DBA','DBA','2021-05-11 16:19:27','2021-05-11 16:19:27');
insert  into `t_role_privs`(`id`,`role_id`,`priv_id`,`creator`,`updater`,`create_date`,`last_update_date`) values (189,33,'02105','DBA','DBA','2021-05-11 16:19:27','2021-05-11 16:19:27');
insert  into `t_role_privs`(`id`,`role_id`,`priv_id`,`creator`,`updater`,`create_date`,`last_update_date`) values (192,33,'06001','DBA','DBA','2021-05-11 16:19:28','2021-05-11 16:19:28');
insert  into `t_role_privs`(`id`,`role_id`,`priv_id`,`creator`,`updater`,`create_date`,`last_update_date`) values (195,33,'06002','DBA','DBA','2021-05-11 16:19:28','2021-05-11 16:19:28');
insert  into `t_role_privs`(`id`,`role_id`,`priv_id`,`creator`,`updater`,`create_date`,`last_update_date`) values (198,33,'06003','DBA','DBA','2021-05-11 16:19:28','2021-05-11 16:19:28');
insert  into `t_role_privs`(`id`,`role_id`,`priv_id`,`creator`,`updater`,`create_date`,`last_update_date`) values (201,36,'00101','DBA','DBA','2021-05-11 16:20:03','2021-05-11 16:20:03');
insert  into `t_role_privs`(`id`,`role_id`,`priv_id`,`creator`,`updater`,`create_date`,`last_update_date`) values (279,39,'00101','DBA','DBA','2021-05-11 16:21:44','2021-05-11 16:21:44');
insert  into `t_role_privs`(`id`,`role_id`,`priv_id`,`creator`,`updater`,`create_date`,`last_update_date`) values (282,39,'00102','DBA','DBA','2021-05-11 16:21:44','2021-05-11 16:21:44');
insert  into `t_role_privs`(`id`,`role_id`,`priv_id`,`creator`,`updater`,`create_date`,`last_update_date`) values (285,39,'00103','DBA','DBA','2021-05-11 16:21:44','2021-05-11 16:21:44');
insert  into `t_role_privs`(`id`,`role_id`,`priv_id`,`creator`,`updater`,`create_date`,`last_update_date`) values (288,39,'00201','DBA','DBA','2021-05-11 16:21:44','2021-05-11 16:21:44');
insert  into `t_role_privs`(`id`,`role_id`,`priv_id`,`creator`,`updater`,`create_date`,`last_update_date`) values (291,39,'00202','DBA','DBA','2021-05-11 16:21:44','2021-05-11 16:21:44');
insert  into `t_role_privs`(`id`,`role_id`,`priv_id`,`creator`,`updater`,`create_date`,`last_update_date`) values (294,39,'00203','DBA','DBA','2021-05-11 16:21:44','2021-05-11 16:21:44');
insert  into `t_role_privs`(`id`,`role_id`,`priv_id`,`creator`,`updater`,`create_date`,`last_update_date`) values (297,39,'00301','DBA','DBA','2021-05-11 16:21:45','2021-05-11 16:21:45');
insert  into `t_role_privs`(`id`,`role_id`,`priv_id`,`creator`,`updater`,`create_date`,`last_update_date`) values (300,39,'00302','DBA','DBA','2021-05-11 16:21:45','2021-05-11 16:21:45');
insert  into `t_role_privs`(`id`,`role_id`,`priv_id`,`creator`,`updater`,`create_date`,`last_update_date`) values (303,39,'00303','DBA','DBA','2021-05-11 16:21:45','2021-05-11 16:21:45');
insert  into `t_role_privs`(`id`,`role_id`,`priv_id`,`creator`,`updater`,`create_date`,`last_update_date`) values (306,39,'00501','DBA','DBA','2021-05-11 16:21:45','2021-05-11 16:21:45');
insert  into `t_role_privs`(`id`,`role_id`,`priv_id`,`creator`,`updater`,`create_date`,`last_update_date`) values (309,39,'00502','DBA','DBA','2021-05-11 16:21:45','2021-05-11 16:21:45');
insert  into `t_role_privs`(`id`,`role_id`,`priv_id`,`creator`,`updater`,`create_date`,`last_update_date`) values (312,39,'00503','DBA','DBA','2021-05-11 16:21:45','2021-05-11 16:21:45');
insert  into `t_role_privs`(`id`,`role_id`,`priv_id`,`creator`,`updater`,`create_date`,`last_update_date`) values (315,39,'00601','DBA','DBA','2021-05-11 16:21:46','2021-05-11 16:21:46');
insert  into `t_role_privs`(`id`,`role_id`,`priv_id`,`creator`,`updater`,`create_date`,`last_update_date`) values (318,39,'01901','DBA','DBA','2021-05-11 16:21:46','2021-05-11 16:21:46');
insert  into `t_role_privs`(`id`,`role_id`,`priv_id`,`creator`,`updater`,`create_date`,`last_update_date`) values (321,39,'01902','DBA','DBA','2021-05-11 16:21:46','2021-05-11 16:21:46');
insert  into `t_role_privs`(`id`,`role_id`,`priv_id`,`creator`,`updater`,`create_date`,`last_update_date`) values (324,39,'01903','DBA','DBA','2021-05-11 16:21:46','2021-05-11 16:21:46');
insert  into `t_role_privs`(`id`,`role_id`,`priv_id`,`creator`,`updater`,`create_date`,`last_update_date`) values (327,39,'02002','DBA','DBA','2021-05-11 16:21:46','2021-05-11 16:21:46');
insert  into `t_role_privs`(`id`,`role_id`,`priv_id`,`creator`,`updater`,`create_date`,`last_update_date`) values (330,39,'02003','DBA','DBA','2021-05-11 16:21:46','2021-05-11 16:21:46');
insert  into `t_role_privs`(`id`,`role_id`,`priv_id`,`creator`,`updater`,`create_date`,`last_update_date`) values (333,39,'02101','DBA','DBA','2021-05-11 16:21:47','2021-05-11 16:21:47');
insert  into `t_role_privs`(`id`,`role_id`,`priv_id`,`creator`,`updater`,`create_date`,`last_update_date`) values (336,39,'02102','DBA','DBA','2021-05-11 16:21:47','2021-05-11 16:21:47');
insert  into `t_role_privs`(`id`,`role_id`,`priv_id`,`creator`,`updater`,`create_date`,`last_update_date`) values (339,39,'02103','DBA','DBA','2021-05-11 16:21:47','2021-05-11 16:21:47');
insert  into `t_role_privs`(`id`,`role_id`,`priv_id`,`creator`,`updater`,`create_date`,`last_update_date`) values (342,39,'02104','DBA','DBA','2021-05-11 16:21:47','2021-05-11 16:21:47');
insert  into `t_role_privs`(`id`,`role_id`,`priv_id`,`creator`,`updater`,`create_date`,`last_update_date`) values (345,39,'02105','DBA','DBA','2021-05-11 16:21:47','2021-05-11 16:21:47');
insert  into `t_role_privs`(`id`,`role_id`,`priv_id`,`creator`,`updater`,`create_date`,`last_update_date`) values (348,39,'06001','DBA','DBA','2021-05-11 16:21:47','2021-05-11 16:21:47');
insert  into `t_role_privs`(`id`,`role_id`,`priv_id`,`creator`,`updater`,`create_date`,`last_update_date`) values (351,39,'06002','DBA','DBA','2021-05-11 16:21:48','2021-05-11 16:21:48');
insert  into `t_role_privs`(`id`,`role_id`,`priv_id`,`creator`,`updater`,`create_date`,`last_update_date`) values (354,39,'06003','DBA','DBA','2021-05-11 16:21:48','2021-05-11 16:21:48');
insert  into `t_role_privs`(`id`,`role_id`,`priv_id`,`creator`,`updater`,`create_date`,`last_update_date`) values (357,42,'00101','DBA','DBA','2021-05-11 16:23:16','2021-05-11 16:23:16');
insert  into `t_role_privs`(`id`,`role_id`,`priv_id`,`creator`,`updater`,`create_date`,`last_update_date`) values (360,42,'00102','DBA','DBA','2021-05-11 16:23:16','2021-05-11 16:23:16');
insert  into `t_role_privs`(`id`,`role_id`,`priv_id`,`creator`,`updater`,`create_date`,`last_update_date`) values (363,42,'00103','DBA','DBA','2021-05-11 16:23:17','2021-05-11 16:23:17');
insert  into `t_role_privs`(`id`,`role_id`,`priv_id`,`creator`,`updater`,`create_date`,`last_update_date`) values (366,42,'00201','DBA','DBA','2021-05-11 16:23:17','2021-05-11 16:23:17');
insert  into `t_role_privs`(`id`,`role_id`,`priv_id`,`creator`,`updater`,`create_date`,`last_update_date`) values (369,42,'00202','DBA','DBA','2021-05-11 16:23:17','2021-05-11 16:23:17');
insert  into `t_role_privs`(`id`,`role_id`,`priv_id`,`creator`,`updater`,`create_date`,`last_update_date`) values (372,42,'00203','DBA','DBA','2021-05-11 16:23:17','2021-05-11 16:23:17');
insert  into `t_role_privs`(`id`,`role_id`,`priv_id`,`creator`,`updater`,`create_date`,`last_update_date`) values (375,42,'00301','DBA','DBA','2021-05-11 16:23:17','2021-05-11 16:23:17');
insert  into `t_role_privs`(`id`,`role_id`,`priv_id`,`creator`,`updater`,`create_date`,`last_update_date`) values (378,42,'00302','DBA','DBA','2021-05-11 16:23:17','2021-05-11 16:23:17');
insert  into `t_role_privs`(`id`,`role_id`,`priv_id`,`creator`,`updater`,`create_date`,`last_update_date`) values (381,42,'00303','DBA','DBA','2021-05-11 16:23:18','2021-05-11 16:23:18');
insert  into `t_role_privs`(`id`,`role_id`,`priv_id`,`creator`,`updater`,`create_date`,`last_update_date`) values (384,42,'00401','DBA','DBA','2021-05-11 16:23:18','2021-05-11 16:23:18');
insert  into `t_role_privs`(`id`,`role_id`,`priv_id`,`creator`,`updater`,`create_date`,`last_update_date`) values (387,42,'00402','DBA','DBA','2021-05-11 16:23:18','2021-05-11 16:23:18');
insert  into `t_role_privs`(`id`,`role_id`,`priv_id`,`creator`,`updater`,`create_date`,`last_update_date`) values (390,42,'00403','DBA','DBA','2021-05-11 16:23:18','2021-05-11 16:23:18');
insert  into `t_role_privs`(`id`,`role_id`,`priv_id`,`creator`,`updater`,`create_date`,`last_update_date`) values (393,42,'00404','DBA','DBA','2021-05-11 16:23:18','2021-05-11 16:23:18');
insert  into `t_role_privs`(`id`,`role_id`,`priv_id`,`creator`,`updater`,`create_date`,`last_update_date`) values (396,42,'00501','DBA','DBA','2021-05-11 16:23:18','2021-05-11 16:23:18');
insert  into `t_role_privs`(`id`,`role_id`,`priv_id`,`creator`,`updater`,`create_date`,`last_update_date`) values (399,42,'00502','DBA','DBA','2021-05-11 16:23:19','2021-05-11 16:23:19');
insert  into `t_role_privs`(`id`,`role_id`,`priv_id`,`creator`,`updater`,`create_date`,`last_update_date`) values (402,42,'00503','DBA','DBA','2021-05-11 16:23:19','2021-05-11 16:23:19');
insert  into `t_role_privs`(`id`,`role_id`,`priv_id`,`creator`,`updater`,`create_date`,`last_update_date`) values (405,42,'00601','DBA','DBA','2021-05-11 16:23:19','2021-05-11 16:23:19');
insert  into `t_role_privs`(`id`,`role_id`,`priv_id`,`creator`,`updater`,`create_date`,`last_update_date`) values (408,42,'00602','DBA','DBA','2021-05-11 16:23:19','2021-05-11 16:23:19');
insert  into `t_role_privs`(`id`,`role_id`,`priv_id`,`creator`,`updater`,`create_date`,`last_update_date`) values (411,42,'00603','DBA','DBA','2021-05-11 16:23:20','2021-05-11 16:23:20');
insert  into `t_role_privs`(`id`,`role_id`,`priv_id`,`creator`,`updater`,`create_date`,`last_update_date`) values (414,42,'00604','DBA','DBA','2021-05-11 16:23:20','2021-05-11 16:23:20');
insert  into `t_role_privs`(`id`,`role_id`,`priv_id`,`creator`,`updater`,`create_date`,`last_update_date`) values (417,42,'00607','DBA','DBA','2021-05-11 16:23:20','2021-05-11 16:23:20');
insert  into `t_role_privs`(`id`,`role_id`,`priv_id`,`creator`,`updater`,`create_date`,`last_update_date`) values (420,42,'00701','DBA','DBA','2021-05-11 16:23:20','2021-05-11 16:23:20');
insert  into `t_role_privs`(`id`,`role_id`,`priv_id`,`creator`,`updater`,`create_date`,`last_update_date`) values (423,42,'00702','DBA','DBA','2021-05-11 16:23:20','2021-05-11 16:23:20');
insert  into `t_role_privs`(`id`,`role_id`,`priv_id`,`creator`,`updater`,`create_date`,`last_update_date`) values (426,42,'00703','DBA','DBA','2021-05-11 16:23:21','2021-05-11 16:23:21');
insert  into `t_role_privs`(`id`,`role_id`,`priv_id`,`creator`,`updater`,`create_date`,`last_update_date`) values (429,42,'00704','DBA','DBA','2021-05-11 16:23:21','2021-05-11 16:23:21');
insert  into `t_role_privs`(`id`,`role_id`,`priv_id`,`creator`,`updater`,`create_date`,`last_update_date`) values (432,42,'00705','DBA','DBA','2021-05-11 16:23:21','2021-05-11 16:23:21');
insert  into `t_role_privs`(`id`,`role_id`,`priv_id`,`creator`,`updater`,`create_date`,`last_update_date`) values (435,42,'00801','DBA','DBA','2021-05-11 16:23:21','2021-05-11 16:23:21');
insert  into `t_role_privs`(`id`,`role_id`,`priv_id`,`creator`,`updater`,`create_date`,`last_update_date`) values (438,42,'00802','DBA','DBA','2021-05-11 16:23:21','2021-05-11 16:23:21');
insert  into `t_role_privs`(`id`,`role_id`,`priv_id`,`creator`,`updater`,`create_date`,`last_update_date`) values (441,42,'00901','DBA','DBA','2021-05-11 16:23:21','2021-05-11 16:23:21');
insert  into `t_role_privs`(`id`,`role_id`,`priv_id`,`creator`,`updater`,`create_date`,`last_update_date`) values (444,42,'00902','DBA','DBA','2021-05-11 16:23:21','2021-05-11 16:23:21');
insert  into `t_role_privs`(`id`,`role_id`,`priv_id`,`creator`,`updater`,`create_date`,`last_update_date`) values (447,42,'00903','DBA','DBA','2021-05-11 16:23:22','2021-05-11 16:23:22');
insert  into `t_role_privs`(`id`,`role_id`,`priv_id`,`creator`,`updater`,`create_date`,`last_update_date`) values (450,42,'00904','DBA','DBA','2021-05-11 16:23:22','2021-05-11 16:23:22');
insert  into `t_role_privs`(`id`,`role_id`,`priv_id`,`creator`,`updater`,`create_date`,`last_update_date`) values (453,42,'00905','DBA','DBA','2021-05-11 16:23:22','2021-05-11 16:23:22');
insert  into `t_role_privs`(`id`,`role_id`,`priv_id`,`creator`,`updater`,`create_date`,`last_update_date`) values (456,42,'01001','DBA','DBA','2021-05-11 16:23:22','2021-05-11 16:23:22');
insert  into `t_role_privs`(`id`,`role_id`,`priv_id`,`creator`,`updater`,`create_date`,`last_update_date`) values (459,42,'01002','DBA','DBA','2021-05-11 16:23:22','2021-05-11 16:23:22');
insert  into `t_role_privs`(`id`,`role_id`,`priv_id`,`creator`,`updater`,`create_date`,`last_update_date`) values (462,42,'01003','DBA','DBA','2021-05-11 16:23:22','2021-05-11 16:23:22');
insert  into `t_role_privs`(`id`,`role_id`,`priv_id`,`creator`,`updater`,`create_date`,`last_update_date`) values (465,42,'01101','DBA','DBA','2021-05-11 16:23:23','2021-05-11 16:23:23');
insert  into `t_role_privs`(`id`,`role_id`,`priv_id`,`creator`,`updater`,`create_date`,`last_update_date`) values (468,42,'01102','DBA','DBA','2021-05-11 16:23:23','2021-05-11 16:23:23');
insert  into `t_role_privs`(`id`,`role_id`,`priv_id`,`creator`,`updater`,`create_date`,`last_update_date`) values (471,42,'01103','DBA','DBA','2021-05-11 16:23:23','2021-05-11 16:23:23');
insert  into `t_role_privs`(`id`,`role_id`,`priv_id`,`creator`,`updater`,`create_date`,`last_update_date`) values (474,42,'01104','DBA','DBA','2021-05-11 16:23:23','2021-05-11 16:23:23');
insert  into `t_role_privs`(`id`,`role_id`,`priv_id`,`creator`,`updater`,`create_date`,`last_update_date`) values (477,42,'01105','DBA','DBA','2021-05-11 16:23:23','2021-05-11 16:23:23');
insert  into `t_role_privs`(`id`,`role_id`,`priv_id`,`creator`,`updater`,`create_date`,`last_update_date`) values (480,42,'01108','DBA','DBA','2021-05-11 16:23:23','2021-05-11 16:23:23');
insert  into `t_role_privs`(`id`,`role_id`,`priv_id`,`creator`,`updater`,`create_date`,`last_update_date`) values (483,42,'01109','DBA','DBA','2021-05-11 16:23:24','2021-05-11 16:23:24');
insert  into `t_role_privs`(`id`,`role_id`,`priv_id`,`creator`,`updater`,`create_date`,`last_update_date`) values (486,42,'01201','DBA','DBA','2021-05-11 16:23:24','2021-05-11 16:23:24');
insert  into `t_role_privs`(`id`,`role_id`,`priv_id`,`creator`,`updater`,`create_date`,`last_update_date`) values (489,42,'01202','DBA','DBA','2021-05-11 16:23:24','2021-05-11 16:23:24');
insert  into `t_role_privs`(`id`,`role_id`,`priv_id`,`creator`,`updater`,`create_date`,`last_update_date`) values (492,42,'01203','DBA','DBA','2021-05-11 16:23:24','2021-05-11 16:23:24');
insert  into `t_role_privs`(`id`,`role_id`,`priv_id`,`creator`,`updater`,`create_date`,`last_update_date`) values (495,42,'01204','DBA','DBA','2021-05-11 16:23:24','2021-05-11 16:23:24');
insert  into `t_role_privs`(`id`,`role_id`,`priv_id`,`creator`,`updater`,`create_date`,`last_update_date`) values (498,42,'01301','DBA','DBA','2021-05-11 16:23:24','2021-05-11 16:23:24');
insert  into `t_role_privs`(`id`,`role_id`,`priv_id`,`creator`,`updater`,`create_date`,`last_update_date`) values (501,42,'01302','DBA','DBA','2021-05-11 16:23:25','2021-05-11 16:23:25');
insert  into `t_role_privs`(`id`,`role_id`,`priv_id`,`creator`,`updater`,`create_date`,`last_update_date`) values (504,42,'01303','DBA','DBA','2021-05-11 16:23:25','2021-05-11 16:23:25');
insert  into `t_role_privs`(`id`,`role_id`,`priv_id`,`creator`,`updater`,`create_date`,`last_update_date`) values (507,42,'01304','DBA','DBA','2021-05-11 16:23:25','2021-05-11 16:23:25');
insert  into `t_role_privs`(`id`,`role_id`,`priv_id`,`creator`,`updater`,`create_date`,`last_update_date`) values (510,42,'01501','DBA','DBA','2021-05-11 16:23:25','2021-05-11 16:23:25');
insert  into `t_role_privs`(`id`,`role_id`,`priv_id`,`creator`,`updater`,`create_date`,`last_update_date`) values (513,42,'01502','DBA','DBA','2021-05-11 16:23:25','2021-05-11 16:23:25');
insert  into `t_role_privs`(`id`,`role_id`,`priv_id`,`creator`,`updater`,`create_date`,`last_update_date`) values (516,42,'01503','DBA','DBA','2021-05-11 16:23:25','2021-05-11 16:23:25');
insert  into `t_role_privs`(`id`,`role_id`,`priv_id`,`creator`,`updater`,`create_date`,`last_update_date`) values (519,42,'01504','DBA','DBA','2021-05-11 16:23:26','2021-05-11 16:23:26');
insert  into `t_role_privs`(`id`,`role_id`,`priv_id`,`creator`,`updater`,`create_date`,`last_update_date`) values (522,42,'01601','DBA','DBA','2021-05-11 16:23:26','2021-05-11 16:23:26');
insert  into `t_role_privs`(`id`,`role_id`,`priv_id`,`creator`,`updater`,`create_date`,`last_update_date`) values (525,42,'01602','DBA','DBA','2021-05-11 16:23:26','2021-05-11 16:23:26');
insert  into `t_role_privs`(`id`,`role_id`,`priv_id`,`creator`,`updater`,`create_date`,`last_update_date`) values (528,42,'01603','DBA','DBA','2021-05-11 16:23:26','2021-05-11 16:23:26');
insert  into `t_role_privs`(`id`,`role_id`,`priv_id`,`creator`,`updater`,`create_date`,`last_update_date`) values (531,42,'01604','DBA','DBA','2021-05-11 16:23:26','2021-05-11 16:23:26');
insert  into `t_role_privs`(`id`,`role_id`,`priv_id`,`creator`,`updater`,`create_date`,`last_update_date`) values (534,42,'01701','DBA','DBA','2021-05-11 16:23:26','2021-05-11 16:23:26');
insert  into `t_role_privs`(`id`,`role_id`,`priv_id`,`creator`,`updater`,`create_date`,`last_update_date`) values (537,42,'01702','DBA','DBA','2021-05-11 16:23:26','2021-05-11 16:23:26');
insert  into `t_role_privs`(`id`,`role_id`,`priv_id`,`creator`,`updater`,`create_date`,`last_update_date`) values (540,42,'01703','DBA','DBA','2021-05-11 16:23:27','2021-05-11 16:23:27');
insert  into `t_role_privs`(`id`,`role_id`,`priv_id`,`creator`,`updater`,`create_date`,`last_update_date`) values (543,42,'01704','DBA','DBA','2021-05-11 16:23:27','2021-05-11 16:23:27');
insert  into `t_role_privs`(`id`,`role_id`,`priv_id`,`creator`,`updater`,`create_date`,`last_update_date`) values (546,42,'01705','DBA','DBA','2021-05-11 16:23:27','2021-05-11 16:23:27');
insert  into `t_role_privs`(`id`,`role_id`,`priv_id`,`creator`,`updater`,`create_date`,`last_update_date`) values (549,42,'01706','DBA','DBA','2021-05-11 16:23:27','2021-05-11 16:23:27');
insert  into `t_role_privs`(`id`,`role_id`,`priv_id`,`creator`,`updater`,`create_date`,`last_update_date`) values (552,42,'01801','DBA','DBA','2021-05-11 16:23:27','2021-05-11 16:23:27');
insert  into `t_role_privs`(`id`,`role_id`,`priv_id`,`creator`,`updater`,`create_date`,`last_update_date`) values (555,42,'01802','DBA','DBA','2021-05-11 16:23:27','2021-05-11 16:23:27');
insert  into `t_role_privs`(`id`,`role_id`,`priv_id`,`creator`,`updater`,`create_date`,`last_update_date`) values (558,42,'01803','DBA','DBA','2021-05-11 16:23:28','2021-05-11 16:23:28');
insert  into `t_role_privs`(`id`,`role_id`,`priv_id`,`creator`,`updater`,`create_date`,`last_update_date`) values (561,42,'01901','DBA','DBA','2021-05-11 16:23:28','2021-05-11 16:23:28');
insert  into `t_role_privs`(`id`,`role_id`,`priv_id`,`creator`,`updater`,`create_date`,`last_update_date`) values (564,42,'01902','DBA','DBA','2021-05-11 16:23:28','2021-05-11 16:23:28');
insert  into `t_role_privs`(`id`,`role_id`,`priv_id`,`creator`,`updater`,`create_date`,`last_update_date`) values (567,42,'01903','DBA','DBA','2021-05-11 16:23:28','2021-05-11 16:23:28');
insert  into `t_role_privs`(`id`,`role_id`,`priv_id`,`creator`,`updater`,`create_date`,`last_update_date`) values (570,42,'02001','DBA','DBA','2021-05-11 16:23:28','2021-05-11 16:23:28');
insert  into `t_role_privs`(`id`,`role_id`,`priv_id`,`creator`,`updater`,`create_date`,`last_update_date`) values (573,42,'02002','DBA','DBA','2021-05-11 16:23:28','2021-05-11 16:23:28');
insert  into `t_role_privs`(`id`,`role_id`,`priv_id`,`creator`,`updater`,`create_date`,`last_update_date`) values (576,42,'02003','DBA','DBA','2021-05-11 16:23:29','2021-05-11 16:23:29');
insert  into `t_role_privs`(`id`,`role_id`,`priv_id`,`creator`,`updater`,`create_date`,`last_update_date`) values (579,42,'02101','DBA','DBA','2021-05-11 16:23:29','2021-05-11 16:23:29');
insert  into `t_role_privs`(`id`,`role_id`,`priv_id`,`creator`,`updater`,`create_date`,`last_update_date`) values (582,42,'02102','DBA','DBA','2021-05-11 16:23:29','2021-05-11 16:23:29');
insert  into `t_role_privs`(`id`,`role_id`,`priv_id`,`creator`,`updater`,`create_date`,`last_update_date`) values (585,42,'02103','DBA','DBA','2021-05-11 16:23:29','2021-05-11 16:23:29');
insert  into `t_role_privs`(`id`,`role_id`,`priv_id`,`creator`,`updater`,`create_date`,`last_update_date`) values (588,42,'02104','DBA','DBA','2021-05-11 16:23:29','2021-05-11 16:23:29');
insert  into `t_role_privs`(`id`,`role_id`,`priv_id`,`creator`,`updater`,`create_date`,`last_update_date`) values (591,42,'02105','DBA','DBA','2021-05-11 16:23:29','2021-05-11 16:23:29');
insert  into `t_role_privs`(`id`,`role_id`,`priv_id`,`creator`,`updater`,`create_date`,`last_update_date`) values (594,42,'06001','DBA','DBA','2021-05-11 16:23:30','2021-05-11 16:23:30');
insert  into `t_role_privs`(`id`,`role_id`,`priv_id`,`creator`,`updater`,`create_date`,`last_update_date`) values (597,42,'06002','DBA','DBA','2021-05-11 16:23:30','2021-05-11 16:23:30');
insert  into `t_role_privs`(`id`,`role_id`,`priv_id`,`creator`,`updater`,`create_date`,`last_update_date`) values (600,42,'06003','DBA','DBA','2021-05-11 16:23:30','2021-05-11 16:23:30');
insert  into `t_role_privs`(`id`,`role_id`,`priv_id`,`creator`,`updater`,`create_date`,`last_update_date`) values (603,42,'06101','DBA','DBA','2021-05-11 16:23:30','2021-05-11 16:23:30');
insert  into `t_role_privs`(`id`,`role_id`,`priv_id`,`creator`,`updater`,`create_date`,`last_update_date`) values (606,42,'06102','DBA','DBA','2021-05-11 16:23:30','2021-05-11 16:23:30');
insert  into `t_role_privs`(`id`,`role_id`,`priv_id`,`creator`,`updater`,`create_date`,`last_update_date`) values (609,42,'06103','DBA','DBA','2021-05-11 16:23:30','2021-05-11 16:23:30');
insert  into `t_role_privs`(`id`,`role_id`,`priv_id`,`creator`,`updater`,`create_date`,`last_update_date`) values (612,45,'01501','DBA','DBA','2021-05-11 16:30:49','2021-05-11 16:30:49');
insert  into `t_role_privs`(`id`,`role_id`,`priv_id`,`creator`,`updater`,`create_date`,`last_update_date`) values (615,45,'01502','DBA','DBA','2021-05-11 16:30:49','2021-05-11 16:30:49');
insert  into `t_role_privs`(`id`,`role_id`,`priv_id`,`creator`,`updater`,`create_date`,`last_update_date`) values (618,45,'01503','DBA','DBA','2021-05-11 16:30:49','2021-05-11 16:30:49');
insert  into `t_role_privs`(`id`,`role_id`,`priv_id`,`creator`,`updater`,`create_date`,`last_update_date`) values (621,45,'01504','DBA','DBA','2021-05-11 16:30:50','2021-05-11 16:30:50');
insert  into `t_role_privs`(`id`,`role_id`,`priv_id`,`creator`,`updater`,`create_date`,`last_update_date`) values (624,48,'01701','DBA','DBA','2021-05-11 16:32:03','2021-05-11 16:32:03');
insert  into `t_role_privs`(`id`,`role_id`,`priv_id`,`creator`,`updater`,`create_date`,`last_update_date`) values (627,48,'01702','DBA','DBA','2021-05-11 16:32:03','2021-05-11 16:32:03');
insert  into `t_role_privs`(`id`,`role_id`,`priv_id`,`creator`,`updater`,`create_date`,`last_update_date`) values (630,48,'01703','DBA','DBA','2021-05-11 16:32:03','2021-05-11 16:32:03');
insert  into `t_role_privs`(`id`,`role_id`,`priv_id`,`creator`,`updater`,`create_date`,`last_update_date`) values (633,48,'00102','DBA','DBA','2021-05-11 16:32:03','2021-05-11 16:32:03');
insert  into `t_role_privs`(`id`,`role_id`,`priv_id`,`creator`,`updater`,`create_date`,`last_update_date`) values (636,51,'00101','DBA','DBA','2021-05-11 16:33:11','2021-05-11 16:33:11');
insert  into `t_role_privs`(`id`,`role_id`,`priv_id`,`creator`,`updater`,`create_date`,`last_update_date`) values (639,51,'00102','DBA','DBA','2021-05-11 16:33:11','2021-05-11 16:33:11');
insert  into `t_role_privs`(`id`,`role_id`,`priv_id`,`creator`,`updater`,`create_date`,`last_update_date`) values (642,51,'00103','DBA','DBA','2021-05-11 16:33:12','2021-05-11 16:33:12');
insert  into `t_role_privs`(`id`,`role_id`,`priv_id`,`creator`,`updater`,`create_date`,`last_update_date`) values (645,51,'00201','DBA','DBA','2021-05-11 16:33:12','2021-05-11 16:33:12');

/*Table structure for table `t_user` */

CREATE TABLE `t_user` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(20) NOT NULL DEFAULT '',
  `emp_no` varchar(10) NOT NULL DEFAULT '',
  `gender` varchar(1) NOT NULL DEFAULT '',
  `email` varchar(50) NOT NULL DEFAULT '',
  `phone` varchar(20) DEFAULT NULL,
  `dept_no` varchar(2) NOT NULL DEFAULT '',
  `expire_date` date NOT NULL,
  `password` varchar(100) NOT NULL DEFAULT '',
  `status` varchar(1) NOT NULL DEFAULT '',
  `login_name` varchar(20) NOT NULL DEFAULT '',
  `creator` varchar(10) NOT NULL DEFAULT '',
  `updater` varchar(10) NOT NULL DEFAULT '',
  `create_date` datetime NOT NULL,
  `last_update_date` datetime NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=372 DEFAULT CHARSET=utf8mb4;

/*Data for the table `t_user` */

insert  into `t_user`(`id`,`name`,`emp_no`,`gender`,`email`,`phone`,`dept_no`,`expire_date`,`password`,`status`,`login_name`,`creator`,`updater`,`create_date`,`last_update_date`) values (159,'马飞','190343','1','190343@lifeat.cn','15810420106','06','2021-06-30','E0E0C5305A8AEB8159E5C35EFA9275F6','1','mafei','DBA','DBA','2018-07-03 00:00:00','2021-02-07 00:00:00');
insert  into `t_user`(`id`,`name`,`emp_no`,`gender`,`email`,`phone`,`dept_no`,`expire_date`,`password`,`status`,`login_name`,`creator`,`updater`,`create_date`,`last_update_date`) values (162,'管理员','190343','1','190343@lifeat.cn','15801620810','03','2021-05-30','84178541B43712FCFD5DCF43D62B6FDD','1','admin','DBA','DBA','2018-08-27 00:00:00','2020-12-03 00:00:00');
insert  into `t_user`(`id`,`name`,`emp_no`,`gender`,`email`,`phone`,`dept_no`,`expire_date`,`password`,`status`,`login_name`,`creator`,`updater`,`create_date`,`last_update_date`) values (165,'王芳芳','190241','2','190241@lifeat.com','18310189890','01','2020-12-31','DD955894EEC804122C45E8782F51FF9C','1','wangfangfang','DBA','DBA','2019-03-19 00:00:00','2020-04-09 00:00:00');
insert  into `t_user`(`id`,`name`,`emp_no`,`gender`,`email`,`phone`,`dept_no`,`expire_date`,`password`,`status`,`login_name`,`creator`,`updater`,`create_date`,`last_update_date`) values (168,'陆俊彪','609012','1','609012@hopson.com.cn','15652630399','01','2020-12-31','727267D7BA387691399B90520C47EC81','1','lujunbiao','DBA','DBA','2019-03-19 00:00:00','2020-04-09 00:00:00');
insert  into `t_user`(`id`,`name`,`emp_no`,`gender`,`email`,`phone`,`dept_no`,`expire_date`,`password`,`status`,`login_name`,`creator`,`updater`,`create_date`,`last_update_date`) values (177,'顾晓','608545','1','608545@hopson.com.cn','13699105067','01','2023-12-28','116C9E3012CCCD3732C5F356CBB19322','1','guxiao','DBA','DBA','2019-03-19 00:00:00','2021-05-07 18:45:55');
insert  into `t_user`(`id`,`name`,`emp_no`,`gender`,`email`,`phone`,`dept_no`,`expire_date`,`password`,`status`,`login_name`,`creator`,`updater`,`create_date`,`last_update_date`) values (180,'张世超','190320','1','190320@lifeat.cn','15510443664','01','2019-12-31','7C307431EF4397C21B89AA2767781771','1','zhangshichao','DBA','DBA','2019-03-19 00:00:00','2020-03-03 00:00:00');
insert  into `t_user`(`id`,`name`,`emp_no`,`gender`,`email`,`phone`,`dept_no`,`expire_date`,`password`,`status`,`login_name`,`creator`,`updater`,`create_date`,`last_update_date`) values (183,'冯甜甜','609701','2','609701@lifeat.cn','13031738066','03','2020-12-31','6CA0C9ABC031A37D7C6720BCB190E693','1','fengtiantian','DBA','DBA','2019-03-19 00:00:00','2021-05-10 08:28:42');
insert  into `t_user`(`id`,`name`,`emp_no`,`gender`,`email`,`phone`,`dept_no`,`expire_date`,`password`,`status`,`login_name`,`creator`,`updater`,`create_date`,`last_update_date`) values (189,'申豹','190291','1','190291@lifeat.cn','13460690607','01','2020-12-31','496C6C4913A8BF11F4148FE2ADE988C9','1','shenbao','DBA','DBA','2019-03-19 00:00:00','2020-03-06 00:00:00');
insert  into `t_user`(`id`,`name`,`emp_no`,`gender`,`email`,`phone`,`dept_no`,`expire_date`,`password`,`status`,`login_name`,`creator`,`updater`,`create_date`,`last_update_date`) values (192,'王志鹏','608520','1','608520@hopson.com.cn','15011380929','01','2020-12-31','FDDBA867D66993270BDB7A3B9A95C275','1','wangzhipeng','DBA','DBA','2019-03-19 00:00:00','2020-03-05 00:00:00');
insert  into `t_user`(`id`,`name`,`emp_no`,`gender`,`email`,`phone`,`dept_no`,`expire_date`,`password`,`status`,`login_name`,`creator`,`updater`,`create_date`,`last_update_date`) values (195,'郭张书','190205','1','190205@lifeat.cn','18210171811','06','2021-12-31','563FF21CA32937FFF097317BED75A188','1','guozhangshu','DBA','DBA','2019-03-19 00:00:00','2021-02-05 00:00:00');
insert  into `t_user`(`id`,`name`,`emp_no`,`gender`,`email`,`phone`,`dept_no`,`expire_date`,`password`,`status`,`login_name`,`creator`,`updater`,`create_date`,`last_update_date`) values (198,'龚涛','609479','1','609479@hopson.com.cn','13811818420','01','2020-12-31','E8D8C75B8090A5CB80CFBC57A7BF7218','1','gongtao','DBA','DBA','2019-03-20 00:00:00','2020-04-29 00:00:00');
insert  into `t_user`(`id`,`name`,`emp_no`,`gender`,`email`,`phone`,`dept_no`,`expire_date`,`password`,`status`,`login_name`,`creator`,`updater`,`create_date`,`last_update_date`) values (201,'龚海鹏','300146','1','gonghaipeng@hopson.com.cn','18611116528','01','2020-12-31','1B1B3DC3C40A5C922E73FDCE000F5F77','1','gonghaipeng','DBA','DBA','2019-03-20 00:00:00','2020-05-15 00:00:00');
insert  into `t_user`(`id`,`name`,`emp_no`,`gender`,`email`,`phone`,`dept_no`,`expire_date`,`password`,`status`,`login_name`,`creator`,`updater`,`create_date`,`last_update_date`) values (204,'李锁臣','190379','1','190379@lifeat.cn','1234234343','06','2020-12-31','225B3CE8F94B9A51CB2C513DC1F213F0','1','lisuochen','DBA','DBA','2019-07-27 00:00:00','2020-04-27 00:00:00');
insert  into `t_user`(`id`,`name`,`emp_no`,`gender`,`email`,`phone`,`dept_no`,`expire_date`,`password`,`status`,`login_name`,`creator`,`updater`,`create_date`,`last_update_date`) values (207,'吕瑞超','190654','1','190654@lifeat.cn','17800166063','06','2020-12-31','86CBEDF6C6330911006564A8CDE73207','1','lvruichao','DBA','DBA','2020-02-28 00:00:00','2020-10-12 00:00:00');
insert  into `t_user`(`id`,`name`,`emp_no`,`gender`,`email`,`phone`,`dept_no`,`expire_date`,`password`,`status`,`login_name`,`creator`,`updater`,`create_date`,`last_update_date`) values (210,'刘洪林','190386','1','190386@lifeat.cn','13699277961','06','2022-01-31','3A022D75E6D5ED45D66221F0365AF613','1','liuhonglin','DBA','DBA','2020-03-04 00:00:00','2021-04-01 00:00:00');
insert  into `t_user`(`id`,`name`,`emp_no`,`gender`,`email`,`phone`,`dept_no`,`expire_date`,`password`,`status`,`login_name`,`creator`,`updater`,`create_date`,`last_update_date`) values (213,'李彬','190375','1','190375@lifeat.cn','18500048699','06','2020-12-31','6640FB3CC9094D61DE5E0AC71E77F113','1','libin','DBA','DBA','2020-03-04 00:00:00','2020-03-04 00:00:00');
insert  into `t_user`(`id`,`name`,`emp_no`,`gender`,`email`,`phone`,`dept_no`,`expire_date`,`password`,`status`,`login_name`,`creator`,`updater`,`create_date`,`last_update_date`) values (216,'李世林','190435','1','190435@lifeat.cn','17603437443','06','2020-12-31','2110435F38A220BF559D1B19BB6E6543','1','lishilin','DBA','DBA','2020-03-04 00:00:00','2020-03-04 00:00:00');
insert  into `t_user`(`id`,`name`,`emp_no`,`gender`,`email`,`phone`,`dept_no`,`expire_date`,`password`,`status`,`login_name`,`creator`,`updater`,`create_date`,`last_update_date`) values (219,'宋胜阳','190438','1','190438@lifeat.cn','18610931339','06','2020-12-31','75653369C3061A1FA12537174CFA2572','1','songshengyang','DBA','DBA','2020-03-04 00:00:00','2020-03-04 00:00:00');
insert  into `t_user`(`id`,`name`,`emp_no`,`gender`,`email`,`phone`,`dept_no`,`expire_date`,`password`,`status`,`login_name`,`creator`,`updater`,`create_date`,`last_update_date`) values (222,'刘旭','190453','1','190453@lifeat.cn','123123123123','06','2020-12-31','2195E71D3F77DA8FAA3A8C10F1D154FC','1','liuxu','DBA','DBA','2020-03-04 00:00:00','2020-04-09 00:00:00');
insert  into `t_user`(`id`,`name`,`emp_no`,`gender`,`email`,`phone`,`dept_no`,`expire_date`,`password`,`status`,`login_name`,`creator`,`updater`,`create_date`,`last_update_date`) values (225,'孙鹏','190656','1','190656@lifeat.cn','17853100297','06','2020-12-31','E77B892D6B0E105602483174B98AA7F8','1','sunpeng','DBA','DBA','2020-03-04 00:00:00','2020-03-04 00:00:00');
insert  into `t_user`(`id`,`name`,`emp_no`,`gender`,`email`,`phone`,`dept_no`,`expire_date`,`password`,`status`,`login_name`,`creator`,`updater`,`create_date`,`last_update_date`) values (228,'侯鑫','190655','1','190655@lifeat.cn','18035541232','06','2020-12-31','5681C9460803235FBBC9EE84F3094A61','1','houxin','DBA','DBA','2020-03-04 00:00:00','2020-03-04 00:00:00');
insert  into `t_user`(`id`,`name`,`emp_no`,`gender`,`email`,`phone`,`dept_no`,`expire_date`,`password`,`status`,`login_name`,`creator`,`updater`,`create_date`,`last_update_date`) values (231,'田杰','190487','1','190487@lifeat.cn','18811152389','06','2020-12-31','5A15CBE91C8DB9A466BE7113A9B95906','1','tianjie','DBA','DBA','2020-03-04 00:00:00','2020-03-24 00:00:00');
insert  into `t_user`(`id`,`name`,`emp_no`,`gender`,`email`,`phone`,`dept_no`,`expire_date`,`password`,`status`,`login_name`,`creator`,`updater`,`create_date`,`last_update_date`) values (234,'黄元涛','190491','1','190491@lifeat.cn','18345559046','06','2020-12-31','5397D56998A2A1281A7559E9B580FA29','1','huangyuantao','DBA','DBA','2020-03-04 00:00:00','2020-12-24 00:00:00');
insert  into `t_user`(`id`,`name`,`emp_no`,`gender`,`email`,`phone`,`dept_no`,`expire_date`,`password`,`status`,`login_name`,`creator`,`updater`,`create_date`,`last_update_date`) values (237,'索其涛','190573','1','190573@lifeat.cn','15701625037','06','2020-12-31','058D1F6721E33124D3CBBB75383F120D','1','suoqitao','DBA','DBA','2020-03-04 00:00:00','2020-03-24 00:00:00');
insert  into `t_user`(`id`,`name`,`emp_no`,`gender`,`email`,`phone`,`dept_no`,`expire_date`,`password`,`status`,`login_name`,`creator`,`updater`,`create_date`,`last_update_date`) values (240,'肖振男','190657','1','190657@lifeat.cn','17610613001','06','2020-12-31','8B87830A780E4205B38710F7C66FB655','1','xiaozhennan','DBA','DBA','2020-03-04 00:00:00','2020-03-04 00:00:00');
insert  into `t_user`(`id`,`name`,`emp_no`,`gender`,`email`,`phone`,`dept_no`,`expire_date`,`password`,`status`,`login_name`,`creator`,`updater`,`create_date`,`last_update_date`) values (243,'王华杰','190648','1','190648@lifeat.cn','18500193260','06','2020-12-31','478B53471FF51FA3D9DD14D04A5480B6','1','wanghuajie','DBA','DBA','2020-03-04 00:00:00','2020-03-28 00:00:00');
insert  into `t_user`(`id`,`name`,`emp_no`,`gender`,`email`,`phone`,`dept_no`,`expire_date`,`password`,`status`,`login_name`,`creator`,`updater`,`create_date`,`last_update_date`) values (246,'李鉴衡','190579','1','190579@lifeat.cn','15503355529','06','2020-12-31','103EB52811020E31EA05D5C089ABE15D','1','lijianheng','DBA','DBA','2020-03-04 00:00:00','2020-03-04 00:00:00');
insert  into `t_user`(`id`,`name`,`emp_no`,`gender`,`email`,`phone`,`dept_no`,`expire_date`,`password`,`status`,`login_name`,`creator`,`updater`,`create_date`,`last_update_date`) values (249,'刘影','609064','2','609064@hopson.com.cn','13524366252','02','2020-12-31','D60DC143A5CAB07855AB76B1B2586960','1','liuying','DBA','DBA','2020-03-04 00:00:00','2020-03-04 00:00:00');
insert  into `t_user`(`id`,`name`,`emp_no`,`gender`,`email`,`phone`,`dept_no`,`expire_date`,`password`,`status`,`login_name`,`creator`,`updater`,`create_date`,`last_update_date`) values (255,'冉爽','190423','1','190423@lifeat.cn','18600750013','02','2020-12-31','ECF4899ADC1CEF69AF400F9DBCF430EC','1','ranshuang','DBA','DBA','2020-03-04 00:00:00','2020-03-04 00:00:00');
insert  into `t_user`(`id`,`name`,`emp_no`,`gender`,`email`,`phone`,`dept_no`,`expire_date`,`password`,`status`,`login_name`,`creator`,`updater`,`create_date`,`last_update_date`) values (264,'王保磊','190577','1','190577@lifeat.cn','18600347582','02','2020-12-31','5A0B8CE52D586081F5EE387801A79E33','1','wangbaolei','DBA','DBA','2020-03-04 00:00:00','2020-03-04 00:00:00');
insert  into `t_user`(`id`,`name`,`emp_no`,`gender`,`email`,`phone`,`dept_no`,`expire_date`,`password`,`status`,`login_name`,`creator`,`updater`,`create_date`,`last_update_date`) values (267,'姚计东','190653','1','190653@lifeat.cn','17191183007','02','2020-12-31','1BEA97B0CC4B22DE386FD6A4B5F404E8','1','yaojidong','DBA','DBA','2020-03-04 00:00:00','2020-03-04 00:00:00');
insert  into `t_user`(`id`,`name`,`emp_no`,`gender`,`email`,`phone`,`dept_no`,`expire_date`,`password`,`status`,`login_name`,`creator`,`updater`,`create_date`,`last_update_date`) values (273,'王博','190393','1','190393@lifeat.cn','﻿18813078624','01','2020-12-31','79E89A1CD02D37C2D1C92A6407AC024B','1','wangbo','DBA','DBA','2020-03-04 00:00:00','2020-03-24 00:00:00');
insert  into `t_user`(`id`,`name`,`emp_no`,`gender`,`email`,`phone`,`dept_no`,`expire_date`,`password`,`status`,`login_name`,`creator`,`updater`,`create_date`,`last_update_date`) values (279,'刘祖鹏','190459','1','190459@lifeat.cn','18500371823','01','2020-12-31','1BA9E8D27CD8BB0DAE65DFCF55D2591E','1','liuzupeng','DBA','DBA','2020-03-04 00:00:00','2020-03-04 00:00:00');
insert  into `t_user`(`id`,`name`,`emp_no`,`gender`,`email`,`phone`,`dept_no`,`expire_date`,`password`,`status`,`login_name`,`creator`,`updater`,`create_date`,`last_update_date`) values (282,'杨德欣','190376','1','190376@lifeat.cn','18600171529','01','2020-12-31','FC96C8ECB1DFA223B69F6F0972975CA9','1','yangdexin','DBA','DBA','2020-03-04 00:00:00','2020-03-04 00:00:00');
insert  into `t_user`(`id`,`name`,`emp_no`,`gender`,`email`,`phone`,`dept_no`,`expire_date`,`password`,`status`,`login_name`,`creator`,`updater`,`create_date`,`last_update_date`) values (291,'王露','190593','2','190593@lifeat.cn','18845573941','01','2020-12-31','A1567BD2B8FD7F24FE717A4E8B85FE24','1','wanglu','DBA','DBA','2020-03-04 00:00:00','2020-03-04 00:00:00');
insert  into `t_user`(`id`,`name`,`emp_no`,`gender`,`email`,`phone`,`dept_no`,`expire_date`,`password`,`status`,`login_name`,`creator`,`updater`,`create_date`,`last_update_date`) values (366,'','','','','','','2021-05-12','','1','','DBA','DBA','2021-05-07 17:58:00','2021-05-07 17:58:00');

/*Table structure for table `t_user_role` */

CREATE TABLE `t_user_role` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) NOT NULL DEFAULT '0',
  `role_id` int(11) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=144 DEFAULT CHARSET=utf8mb4;

/*Data for the table `t_user_role` */

insert  into `t_user_role`(`id`,`user_id`,`role_id`) values (48,363,1);
insert  into `t_user_role`(`id`,`user_id`,`role_id`) values (51,363,5);
insert  into `t_user_role`(`id`,`user_id`,`role_id`) values (54,363,2);
insert  into `t_user_role`(`id`,`user_id`,`role_id`) values (57,363,3);
insert  into `t_user_role`(`id`,`user_id`,`role_id`) values (60,183,1);
insert  into `t_user_role`(`id`,`user_id`,`role_id`) values (63,183,1);
insert  into `t_user_role`(`id`,`user_id`,`role_id`) values (66,183,2);
insert  into `t_user_role`(`id`,`user_id`,`role_id`) values (69,183,4);
insert  into `t_user_role`(`id`,`user_id`,`role_id`) values (72,177,4);
insert  into `t_user_role`(`id`,`user_id`,`role_id`) values (75,177,4);
insert  into `t_user_role`(`id`,`user_id`,`role_id`) values (78,183,1);
insert  into `t_user_role`(`id`,`user_id`,`role_id`) values (81,183,1);
insert  into `t_user_role`(`id`,`user_id`,`role_id`) values (84,183,2);
insert  into `t_user_role`(`id`,`user_id`,`role_id`) values (87,369,2);
insert  into `t_user_role`(`id`,`user_id`,`role_id`) values (90,270,3);
insert  into `t_user_role`(`id`,`user_id`,`role_id`) values (93,285,2);
insert  into `t_user_role`(`id`,`user_id`,`role_id`) values (96,285,3);
insert  into `t_user_role`(`id`,`user_id`,`role_id`) values (99,294,4);
insert  into `t_user_role`(`id`,`user_id`,`role_id`) values (102,288,2);
insert  into `t_user_role`(`id`,`user_id`,`role_id`) values (105,276,2);
insert  into `t_user_role`(`id`,`user_id`,`role_id`) values (108,276,3);
insert  into `t_user_role`(`id`,`user_id`,`role_id`) values (111,276,4);

/*Table structure for table `t_xtqx` */

CREATE TABLE `t_xtqx` (
  `id` varchar(255) NOT NULL,
  `name` varchar(20) NOT NULL DEFAULT '',
  `parent_id` varchar(10) NOT NULL DEFAULT '',
  `url` varchar(100) NOT NULL DEFAULT '',
  `status` varchar(1) NOT NULL DEFAULT '',
  `icon` varchar(50) DEFAULT '',
  `creator` varchar(10) DEFAULT NULL,
  `updater` varchar(10) DEFAULT NULL,
  `create_date` datetime DEFAULT NULL,
  `last_update_date` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

/*Data for the table `t_xtqx` */

insert  into `t_xtqx`(`id`,`name`,`parent_id`,`url`,`status`,`icon`,`creator`,`updater`,`create_date`,`last_update_date`) values ('0','根节点','','','1',NULL,'DBA','DBA','2021-05-10 08:51:59','2021-05-10 08:51:59');
insert  into `t_xtqx`(`id`,`name`,`parent_id`,`url`,`status`,`icon`,`creator`,`updater`,`create_date`,`last_update_date`) values ('001','用户管理','0','','1','fa fa-user','DBA','DBA','2021-05-10 08:51:59','2021-05-10 08:51:59');
insert  into `t_xtqx`(`id`,`name`,`parent_id`,`url`,`status`,`icon`,`creator`,`updater`,`create_date`,`last_update_date`) values ('00101','用户新增','001','/user/add','1',NULL,'DBA','DBA','2021-05-10 08:51:59','2021-05-10 08:51:59');
insert  into `t_xtqx`(`id`,`name`,`parent_id`,`url`,`status`,`icon`,`creator`,`updater`,`create_date`,`last_update_date`) values ('00102','用户查询','001','/user/query','1',NULL,'DBA','DBA','2021-05-10 08:51:59','2021-05-10 08:51:59');
insert  into `t_xtqx`(`id`,`name`,`parent_id`,`url`,`status`,`icon`,`creator`,`updater`,`create_date`,`last_update_date`) values ('00103','用户变更','001','/user/change','1',NULL,'DBA','DBA','2021-05-10 08:51:59','2021-05-10 08:51:59');
insert  into `t_xtqx`(`id`,`name`,`parent_id`,`url`,`status`,`icon`,`creator`,`updater`,`create_date`,`last_update_date`) values ('002','角色管理','0','','1','mdi mdi-account-switch','DBA','DBA','2021-05-10 08:51:59','2021-05-10 08:51:59');
insert  into `t_xtqx`(`id`,`name`,`parent_id`,`url`,`status`,`icon`,`creator`,`updater`,`create_date`,`last_update_date`) values ('00201','角色查询','002','/role/query','1',NULL,'DBA','DBA','2021-05-10 08:51:59','2021-05-10 08:51:59');
insert  into `t_xtqx`(`id`,`name`,`parent_id`,`url`,`status`,`icon`,`creator`,`updater`,`create_date`,`last_update_date`) values ('00202','角色新增','002','/role/add','1',NULL,'DBA','DBA','2021-05-10 08:51:59','2021-05-10 08:51:59');
insert  into `t_xtqx`(`id`,`name`,`parent_id`,`url`,`status`,`icon`,`creator`,`updater`,`create_date`,`last_update_date`) values ('00203','角色变更','002','/role/change','1',NULL,'DBA','DBA','2021-05-10 08:51:59','2021-05-10 08:51:59');
insert  into `t_xtqx`(`id`,`name`,`parent_id`,`url`,`status`,`icon`,`creator`,`updater`,`create_date`,`last_update_date`) values ('003','菜单管理','0','','1','mdi mdi-file-tree','DBA','DBA','2021-05-10 08:51:59','2021-05-10 08:51:59');
insert  into `t_xtqx`(`id`,`name`,`parent_id`,`url`,`status`,`icon`,`creator`,`updater`,`create_date`,`last_update_date`) values ('00301','菜单查询','003','/menu/query','1',NULL,'DBA','DBA','2021-05-10 08:51:59','2021-05-10 08:51:59');
insert  into `t_xtqx`(`id`,`name`,`parent_id`,`url`,`status`,`icon`,`creator`,`updater`,`create_date`,`last_update_date`) values ('00302','菜单新增','003','/menu/add','1',NULL,'DBA','DBA','2021-05-10 08:51:59','2021-05-10 08:51:59');
insert  into `t_xtqx`(`id`,`name`,`parent_id`,`url`,`status`,`icon`,`creator`,`updater`,`create_date`,`last_update_date`) values ('00303','菜单变更','003','/menu/change','1',NULL,'DBA','DBA','2021-05-10 08:51:59','2021-05-10 08:51:59');
insert  into `t_xtqx`(`id`,`name`,`parent_id`,`url`,`status`,`icon`,`creator`,`updater`,`create_date`,`last_update_date`) values ('004','数据源管理','0','','1','mdi mdi-chemical-weapon','DBA','DBA','2021-05-10 08:51:59','2021-05-10 08:51:59');
insert  into `t_xtqx`(`id`,`name`,`parent_id`,`url`,`status`,`icon`,`creator`,`updater`,`create_date`,`last_update_date`) values ('00401','数据源查询','004','/ds/query','1',NULL,'DBA','DBA','2021-05-10 08:51:59','2021-05-10 08:51:59');
insert  into `t_xtqx`(`id`,`name`,`parent_id`,`url`,`status`,`icon`,`creator`,`updater`,`create_date`,`last_update_date`) values ('00402','数据源新增','004','/ds/add','1',NULL,'DBA','DBA','2021-05-10 08:51:59','2021-05-10 08:51:59');
insert  into `t_xtqx`(`id`,`name`,`parent_id`,`url`,`status`,`icon`,`creator`,`updater`,`create_date`,`last_update_date`) values ('00403','数据源维护','004','/ds/change','1',NULL,'DBA','DBA','2021-05-10 08:51:59','2021-05-10 08:51:59');
insert  into `t_xtqx`(`id`,`name`,`parent_id`,`url`,`status`,`icon`,`creator`,`updater`,`create_date`,`last_update_date`) values ('00404','数据源测试','004','/ds/test','1',NULL,'DBA','DBA','2021-05-10 08:51:59','2021-05-10 08:51:59');
insert  into `t_xtqx`(`id`,`name`,`parent_id`,`url`,`status`,`icon`,`creator`,`updater`,`create_date`,`last_update_date`) values ('005','服务器管理','0','','1','mdi mdi-monitor','DBA','DBA','2021-05-10 08:51:59','2021-05-10 08:51:59');
insert  into `t_xtqx`(`id`,`name`,`parent_id`,`url`,`status`,`icon`,`creator`,`updater`,`create_date`,`last_update_date`) values ('00501','服务器查询','005','/server/query','1',NULL,'DBA','DBA','2021-05-10 08:51:59','2021-05-10 08:51:59');
insert  into `t_xtqx`(`id`,`name`,`parent_id`,`url`,`status`,`icon`,`creator`,`updater`,`create_date`,`last_update_date`) values ('00502','新增服务器','005','/server/add','1',NULL,'DBA','DBA','2021-05-10 08:51:59','2021-05-10 08:51:59');
insert  into `t_xtqx`(`id`,`name`,`parent_id`,`url`,`status`,`icon`,`creator`,`updater`,`create_date`,`last_update_date`) values ('00503','服务器变更','005','/server/change','1',NULL,'DBA','DBA','2021-05-10 08:51:59','2021-05-10 08:51:59');
insert  into `t_xtqx`(`id`,`name`,`parent_id`,`url`,`status`,`icon`,`creator`,`updater`,`create_date`,`last_update_date`) values ('006','数据库管理','0','','1','mdi mdi-settings','DBA','DBA','2021-05-10 08:51:59','2021-05-10 08:51:59');
insert  into `t_xtqx`(`id`,`name`,`parent_id`,`url`,`status`,`icon`,`creator`,`updater`,`create_date`,`last_update_date`) values ('00601','实例创建','006','/db/inst/crt/query','1','','DBA','DBA','2021-05-10 08:51:59','2021-05-10 08:51:59');
insert  into `t_xtqx`(`id`,`name`,`parent_id`,`url`,`status`,`icon`,`creator`,`updater`,`create_date`,`last_update_date`) values ('00602','实例管理','006','/db/inst/query','1','','DBA','DBA','2021-05-10 08:51:59','2021-05-10 08:51:59');
insert  into `t_xtqx`(`id`,`name`,`parent_id`,`url`,`status`,`icon`,`creator`,`updater`,`create_date`,`last_update_date`) values ('00603','账号管理','006','/db/user/query','1','','DBA','DBA','2021-05-10 08:51:59','2021-05-10 08:51:59');
insert  into `t_xtqx`(`id`,`name`,`parent_id`,`url`,`status`,`icon`,`creator`,`updater`,`create_date`,`last_update_date`) values ('00604','配置管理','006','/db/inst/cfg/query','1','','DBA','DBA','2021-05-10 08:51:59','2021-05-10 08:51:59');
insert  into `t_xtqx`(`id`,`name`,`parent_id`,`url`,`status`,`icon`,`creator`,`updater`,`create_date`,`last_update_date`) values ('00605','权限管理','006','/db/inst/priv','0','','DBA','DBA','2021-05-10 08:51:59','2021-05-10 08:51:59');
insert  into `t_xtqx`(`id`,`name`,`parent_id`,`url`,`status`,`icon`,`creator`,`updater`,`create_date`,`last_update_date`) values ('00606','参数管理','006','/db/inst/para/query','0','','DBA','DBA','2021-05-10 08:51:59','2021-05-10 08:51:59');
insert  into `t_xtqx`(`id`,`name`,`parent_id`,`url`,`status`,`icon`,`creator`,`updater`,`create_date`,`last_update_date`) values ('00607','操作日志','006','/db/inst/opt/log/query','1','','DBA','DBA','2021-05-10 08:51:59','2021-05-10 08:51:59');
insert  into `t_xtqx`(`id`,`name`,`parent_id`,`url`,`status`,`icon`,`creator`,`updater`,`create_date`,`last_update_date`) values ('007','数据库监控','0','','1','ion-eye','DBA','DBA','2021-05-10 08:51:59','2021-05-10 08:51:59');
insert  into `t_xtqx`(`id`,`name`,`parent_id`,`url`,`status`,`icon`,`creator`,`updater`,`create_date`,`last_update_date`) values ('00701','指标管理','007','/monitor/index/query','1','','DBA','DBA','2021-05-10 08:51:59','2021-05-10 08:51:59');
insert  into `t_xtqx`(`id`,`name`,`parent_id`,`url`,`status`,`icon`,`creator`,`updater`,`create_date`,`last_update_date`) values ('00702','模板管理','007','/monitor/templete/query','1','','DBA','DBA','2021-05-10 08:51:59','2021-05-10 08:51:59');
insert  into `t_xtqx`(`id`,`name`,`parent_id`,`url`,`status`,`icon`,`creator`,`updater`,`create_date`,`last_update_date`) values ('00703','任务管理','007','/monitor/task/query','1','','DBA','DBA','2021-05-10 08:51:59','2021-05-10 08:51:59');
insert  into `t_xtqx`(`id`,`name`,`parent_id`,`url`,`status`,`icon`,`creator`,`updater`,`create_date`,`last_update_date`) values ('00704','监控图表','007','/monitor/graph/query','1','','DBA','DBA','2021-05-10 08:51:59','2021-05-10 08:51:59');
insert  into `t_xtqx`(`id`,`name`,`parent_id`,`url`,`status`,`icon`,`creator`,`updater`,`create_date`,`last_update_date`) values ('00705','监控大屏','007','/monitor/view','1',NULL,'DBA','DBA','2021-05-10 08:51:59','2021-05-10 08:51:59');
insert  into `t_xtqx`(`id`,`name`,`parent_id`,`url`,`status`,`icon`,`creator`,`updater`,`create_date`,`last_update_date`) values ('008','数据库工具','0','','0','mdi mdi-database','DBA','DBA','2021-05-10 08:51:59','2021-05-10 08:51:59');
insert  into `t_xtqx`(`id`,`name`,`parent_id`,`url`,`status`,`icon`,`creator`,`updater`,`create_date`,`last_update_date`) values ('00801','字典生成','008','/dba/dict','1','','DBA','DBA','2021-05-10 08:51:59','2021-05-10 08:51:59');
insert  into `t_xtqx`(`id`,`name`,`parent_id`,`url`,`status`,`icon`,`creator`,`updater`,`create_date`,`last_update_date`) values ('00802','Redis迁移','008','/redis/migrate','1','','DBA','DBA','2021-05-10 08:51:59','2021-05-10 08:51:59');
insert  into `t_xtqx`(`id`,`name`,`parent_id`,`url`,`status`,`icon`,`creator`,`updater`,`create_date`,`last_update_date`) values ('009','数据库备份','0','','1','mdi mdi-content-copy','DBA','DBA','2021-05-10 08:51:59','2021-05-10 08:51:59');
insert  into `t_xtqx`(`id`,`name`,`parent_id`,`url`,`status`,`icon`,`creator`,`updater`,`create_date`,`last_update_date`) values ('00901','新建备份','009','/backup/add','1','','DBA','DBA','2021-05-10 08:51:59','2021-05-10 08:51:59');
insert  into `t_xtqx`(`id`,`name`,`parent_id`,`url`,`status`,`icon`,`creator`,`updater`,`create_date`,`last_update_date`) values ('00902','备份维护','009','/backup/change','1','','DBA','DBA','2021-05-10 08:51:59','2021-05-10 08:51:59');
insert  into `t_xtqx`(`id`,`name`,`parent_id`,`url`,`status`,`icon`,`creator`,`updater`,`create_date`,`last_update_date`) values ('00903','任务查询','009','/backup/query','1','','DBA','DBA','2021-05-10 08:51:59','2021-05-10 08:51:59');
insert  into `t_xtqx`(`id`,`name`,`parent_id`,`url`,`status`,`icon`,`creator`,`updater`,`create_date`,`last_update_date`) values ('00904','日志查询','009','/backup/log/query','1','','DBA','DBA','2021-05-10 08:51:59','2021-05-10 08:51:59');
insert  into `t_xtqx`(`id`,`name`,`parent_id`,`url`,`status`,`icon`,`creator`,`updater`,`create_date`,`last_update_date`) values ('00905','日志分析','009','/backup/log/analyze','1','','DBA','DBA','2021-05-10 08:51:59','2021-05-10 08:51:59');
insert  into `t_xtqx`(`id`,`name`,`parent_id`,`url`,`status`,`icon`,`creator`,`updater`,`create_date`,`last_update_date`) values ('010','数据库恢复','0','','0','mdi mdi-backup-restore','DBA','DBA','2021-05-10 08:51:59','2021-05-10 08:51:59');
insert  into `t_xtqx`(`id`,`name`,`parent_id`,`url`,`status`,`icon`,`creator`,`updater`,`create_date`,`last_update_date`) values ('01001','恢复向导','010','/recover/guide','1','','DBA','DBA','2021-05-10 08:51:59','2021-05-10 08:51:59');
insert  into `t_xtqx`(`id`,`name`,`parent_id`,`url`,`status`,`icon`,`creator`,`updater`,`create_date`,`last_update_date`) values ('01002','恢复配置','010','/recover/config','1','','DBA','DBA','2021-05-10 08:51:59','2021-05-10 08:51:59');
insert  into `t_xtqx`(`id`,`name`,`parent_id`,`url`,`status`,`icon`,`creator`,`updater`,`create_date`,`last_update_date`) values ('01003','恢复查询','010','/recover/query','1','','DBA','DBA','2021-05-10 08:51:59','2021-05-10 08:51:59');
insert  into `t_xtqx`(`id`,`name`,`parent_id`,`url`,`status`,`icon`,`creator`,`updater`,`create_date`,`last_update_date`) values ('011','数据库同步','0','','1','mdi mdi-sync','DBA','DBA','2021-05-10 08:51:59','2021-05-10 08:51:59');
insert  into `t_xtqx`(`id`,`name`,`parent_id`,`url`,`status`,`icon`,`creator`,`updater`,`create_date`,`last_update_date`) values ('01101','离线同步','011','/sync/add','1','','DBA','DBA','2021-05-10 08:51:59','2021-05-10 08:51:59');
insert  into `t_xtqx`(`id`,`name`,`parent_id`,`url`,`status`,`icon`,`creator`,`updater`,`create_date`,`last_update_date`) values ('01102','同步维护','011','/sync/change','1','','DBA','DBA','2021-05-10 08:51:59','2021-05-10 08:51:59');
insert  into `t_xtqx`(`id`,`name`,`parent_id`,`url`,`status`,`icon`,`creator`,`updater`,`create_date`,`last_update_date`) values ('01103','任务查询','011','/sync/query','1','','DBA','DBA','2021-05-10 08:51:59','2021-05-10 08:51:59');
insert  into `t_xtqx`(`id`,`name`,`parent_id`,`url`,`status`,`icon`,`creator`,`updater`,`create_date`,`last_update_date`) values ('01104','日志查询','011','/sync/log/query','1','','DBA','DBA','2021-05-10 08:51:59','2021-05-10 08:51:59');
insert  into `t_xtqx`(`id`,`name`,`parent_id`,`url`,`status`,`icon`,`creator`,`updater`,`create_date`,`last_update_date`) values ('01105','日志分析','011','/sync/log/analyze','1','','DBA','DBA','2021-05-10 08:51:59','2021-05-10 08:51:59');
insert  into `t_xtqx`(`id`,`name`,`parent_id`,`url`,`status`,`icon`,`creator`,`updater`,`create_date`,`last_update_date`) values ('01107','同步拓扑','011','/sync/log/graph','0','','DBA','DBA','2021-05-10 08:51:59','2021-05-10 08:51:59');
insert  into `t_xtqx`(`id`,`name`,`parent_id`,`url`,`status`,`icon`,`creator`,`updater`,`create_date`,`last_update_date`) values ('01108','实时同步','011','/sync/log/real','1','','DBA','DBA','2021-05-10 08:51:59','2021-05-10 08:51:59');
insert  into `t_xtqx`(`id`,`name`,`parent_id`,`url`,`status`,`icon`,`creator`,`updater`,`create_date`,`last_update_date`) values ('01109','日志同步','011','/sync/log/kafaca','1',NULL,'DBA','DBA','2021-05-10 08:51:59','2021-05-10 08:51:59');
insert  into `t_xtqx`(`id`,`name`,`parent_id`,`url`,`status`,`icon`,`creator`,`updater`,`create_date`,`last_update_date`) values ('012','数据库传输','0','','1','mdi mdi-swap-horizontal','DBA','DBA','2021-05-10 08:51:59','2021-05-10 08:51:59');
insert  into `t_xtqx`(`id`,`name`,`parent_id`,`url`,`status`,`icon`,`creator`,`updater`,`create_date`,`last_update_date`) values ('01201','新建传输','012','/transfer/add','1','','DBA','DBA','2021-05-10 08:51:59','2021-05-10 08:51:59');
insert  into `t_xtqx`(`id`,`name`,`parent_id`,`url`,`status`,`icon`,`creator`,`updater`,`create_date`,`last_update_date`) values ('01202','传输查询','012','/transfer/query','1','','DBA','DBA','2021-05-10 08:51:59','2021-05-10 08:51:59');
insert  into `t_xtqx`(`id`,`name`,`parent_id`,`url`,`status`,`icon`,`creator`,`updater`,`create_date`,`last_update_date`) values ('01203','传输维护','012','/transfer/change','1','','DBA','DBA','2021-05-10 08:51:59','2021-05-10 08:51:59');
insert  into `t_xtqx`(`id`,`name`,`parent_id`,`url`,`status`,`icon`,`creator`,`updater`,`create_date`,`last_update_date`) values ('01204','日志查询','012','/transfer/log/query','1','','DBA','DBA','2021-05-10 08:51:59','2021-05-10 08:51:59');
insert  into `t_xtqx`(`id`,`name`,`parent_id`,`url`,`status`,`icon`,`creator`,`updater`,`create_date`,`last_update_date`) values ('013','数据库归档','0','','1','mdi mdi-lan-connect','DBA','DBA','2021-05-10 08:51:59','2021-05-10 08:51:59');
insert  into `t_xtqx`(`id`,`name`,`parent_id`,`url`,`status`,`icon`,`creator`,`updater`,`create_date`,`last_update_date`) values ('01301','新建归档','013','/archive/add','1','','DBA','DBA','2021-05-10 08:51:59','2021-05-10 08:51:59');
insert  into `t_xtqx`(`id`,`name`,`parent_id`,`url`,`status`,`icon`,`creator`,`updater`,`create_date`,`last_update_date`) values ('01302','归档查询','013','/archive/query','1','','DBA','DBA','2021-05-10 08:51:59','2021-05-10 08:51:59');
insert  into `t_xtqx`(`id`,`name`,`parent_id`,`url`,`status`,`icon`,`creator`,`updater`,`create_date`,`last_update_date`) values ('01303','归档维护','013','/archive/change','1','','DBA','DBA','2021-05-10 08:51:59','2021-05-10 08:51:59');
insert  into `t_xtqx`(`id`,`name`,`parent_id`,`url`,`status`,`icon`,`creator`,`updater`,`create_date`,`last_update_date`) values ('01304','日志查询','013','/archive/log/query','1','','DBA','DBA','2021-05-10 08:51:59','2021-05-10 08:51:59');
insert  into `t_xtqx`(`id`,`name`,`parent_id`,`url`,`status`,`icon`,`creator`,`updater`,`create_date`,`last_update_date`) values ('014','数据库部署','0','','0','mdi mdi-polymer','DBA','DBA','2021-05-10 08:51:59','2021-05-10 08:51:59');
insert  into `t_xtqx`(`id`,`name`,`parent_id`,`url`,`status`,`icon`,`creator`,`updater`,`create_date`,`last_update_date`) values ('015','大数据同步','0','','1','ion-ios7-pulse-strong','DBA','DBA','2021-05-10 08:51:59','2021-05-10 08:51:59');
insert  into `t_xtqx`(`id`,`name`,`parent_id`,`url`,`status`,`icon`,`creator`,`updater`,`create_date`,`last_update_date`) values ('01501','新增同步','015','/bigdata/add','1','','DBA','DBA','2021-05-10 08:51:59','2021-05-10 08:51:59');
insert  into `t_xtqx`(`id`,`name`,`parent_id`,`url`,`status`,`icon`,`creator`,`updater`,`create_date`,`last_update_date`) values ('01502','任务查询','015','/bigdata/query','1','','DBA','DBA','2021-05-10 08:51:59','2021-05-10 08:51:59');
insert  into `t_xtqx`(`id`,`name`,`parent_id`,`url`,`status`,`icon`,`creator`,`updater`,`create_date`,`last_update_date`) values ('01503','同步维护','015','/bigdata/change','1','','DBA','DBA','2021-05-10 08:51:59','2021-05-10 08:51:59');
insert  into `t_xtqx`(`id`,`name`,`parent_id`,`url`,`status`,`icon`,`creator`,`updater`,`create_date`,`last_update_date`) values ('01504','日志分析','015','/bigdata/log/analyze','1','','DBA','DBA','2021-05-10 08:51:59','2021-05-10 08:51:59');
insert  into `t_xtqx`(`id`,`name`,`parent_id`,`url`,`status`,`icon`,`creator`,`updater`,`create_date`,`last_update_date`) values ('016','慢日志管理','0','','1','mdi mdi-crop','DBA','DBA','2021-05-10 08:51:59','2021-05-10 08:51:59');
insert  into `t_xtqx`(`id`,`name`,`parent_id`,`url`,`status`,`icon`,`creator`,`updater`,`create_date`,`last_update_date`) values ('01601','慢日志新增','016','/slow/add','1','','DBA','DBA','2021-05-10 08:51:59','2021-05-10 08:51:59');
insert  into `t_xtqx`(`id`,`name`,`parent_id`,`url`,`status`,`icon`,`creator`,`updater`,`create_date`,`last_update_date`) values ('01602','慢日志维护','016','/slow/change','1','','DBA','DBA','2021-05-10 08:51:59','2021-05-10 08:51:59');
insert  into `t_xtqx`(`id`,`name`,`parent_id`,`url`,`status`,`icon`,`creator`,`updater`,`create_date`,`last_update_date`) values ('01603','慢日志查询','016','/slow/log/query','1','','DBA','DBA','2021-05-10 08:51:59','2021-05-10 08:51:59');
insert  into `t_xtqx`(`id`,`name`,`parent_id`,`url`,`status`,`icon`,`creator`,`updater`,`create_date`,`last_update_date`) values ('01604','慢日志分析','016','/slow/log/analyze','1','','DBA','DBA','2021-05-10 08:51:59','2021-05-10 08:51:59');
insert  into `t_xtqx`(`id`,`name`,`parent_id`,`url`,`status`,`icon`,`creator`,`updater`,`create_date`,`last_update_date`) values ('017','工单管理','0','','1','mdi mdi-format-align-left','DBA','DBA','2021-05-10 08:51:59','2021-05-10 08:51:59');
insert  into `t_xtqx`(`id`,`name`,`parent_id`,`url`,`status`,`icon`,`creator`,`updater`,`create_date`,`last_update_date`) values ('01701','我的工单','017','/order/query','1','','DBA','DBA','2021-05-10 08:51:59','2021-05-10 08:51:59');
insert  into `t_xtqx`(`id`,`name`,`parent_id`,`url`,`status`,`icon`,`creator`,`updater`,`create_date`,`last_update_date`) values ('01702','SQL查询','017','/sql/query','1','','DBA','DBA','2021-05-10 08:51:59','2021-05-10 08:51:59');
insert  into `t_xtqx`(`id`,`name`,`parent_id`,`url`,`status`,`icon`,`creator`,`updater`,`create_date`,`last_update_date`) values ('01703','SQL发布','017','/sql/release','1','','DBA','DBA','2021-05-10 08:51:59','2021-05-10 08:51:59');
insert  into `t_xtqx`(`id`,`name`,`parent_id`,`url`,`status`,`icon`,`creator`,`updater`,`create_date`,`last_update_date`) values ('01704','SQL审核','017','/sql/audit','1','','DBA','DBA','2021-05-10 08:51:59','2021-05-10 08:51:59');
insert  into `t_xtqx`(`id`,`name`,`parent_id`,`url`,`status`,`icon`,`creator`,`updater`,`create_date`,`last_update_date`) values ('01705','SQL执行','017','/sql/run','1','','DBA','DBA','2021-05-10 08:51:59','2021-05-10 08:51:59');
insert  into `t_xtqx`(`id`,`name`,`parent_id`,`url`,`status`,`icon`,`creator`,`updater`,`create_date`,`last_update_date`) values ('01706','SQL导出','017','/sql/exp','1','','DBA','DBA','2021-05-10 08:51:59','2021-05-10 08:51:59');
insert  into `t_xtqx`(`id`,`name`,`parent_id`,`url`,`status`,`icon`,`creator`,`updater`,`create_date`,`last_update_date`) values ('01707','phoenix查询','017','/sql/phoenix/sql','0','','DBA','DBA','2021-05-10 08:51:59','2021-05-10 08:51:59');
insert  into `t_xtqx`(`id`,`name`,`parent_id`,`url`,`status`,`icon`,`creator`,`updater`,`create_date`,`last_update_date`) values ('01708','phoenix发布','017','/sql/phoenix/release','0','','DBA','DBA','2021-05-10 08:51:59','2021-05-10 08:51:59');
insert  into `t_xtqx`(`id`,`name`,`parent_id`,`url`,`status`,`icon`,`creator`,`updater`,`create_date`,`last_update_date`) values ('018','任务管理','0','','0','mdi mdi-alarm','DBA','DBA','2021-05-10 08:51:59','2021-05-10 08:51:59');
insert  into `t_xtqx`(`id`,`name`,`parent_id`,`url`,`status`,`icon`,`creator`,`updater`,`create_date`,`last_update_date`) values ('01801','新建任务','018','/task/new','1','','DBA','DBA','2021-05-10 08:51:59','2021-05-10 08:51:59');
insert  into `t_xtqx`(`id`,`name`,`parent_id`,`url`,`status`,`icon`,`creator`,`updater`,`create_date`,`last_update_date`) values ('01802','任务查询','018','/task/query','1','','DBA','DBA','2021-05-10 08:51:59','2021-05-10 08:51:59');
insert  into `t_xtqx`(`id`,`name`,`parent_id`,`url`,`status`,`icon`,`creator`,`updater`,`create_date`,`last_update_date`) values ('01803','任务变更','018','/task/change','1','','DBA','DBA','2021-05-10 08:51:59','2021-05-10 08:51:59');
insert  into `t_xtqx`(`id`,`name`,`parent_id`,`url`,`status`,`icon`,`creator`,`updater`,`create_date`,`last_update_date`) values ('019','端口管理','0','','1','mdi mdi-code-brackets','DBA','DBA','2021-05-10 08:51:59','2021-05-10 08:51:59');
insert  into `t_xtqx`(`id`,`name`,`parent_id`,`url`,`status`,`icon`,`creator`,`updater`,`create_date`,`last_update_date`) values ('01901','端口新增','019','/port/add','1','','DBA','DBA','2021-05-10 08:51:59','2021-05-10 08:51:59');
insert  into `t_xtqx`(`id`,`name`,`parent_id`,`url`,`status`,`icon`,`creator`,`updater`,`create_date`,`last_update_date`) values ('01902','端口查询','019','/port/query','1','','DBA','DBA','2021-05-10 08:51:59','2021-05-10 08:51:59');
insert  into `t_xtqx`(`id`,`name`,`parent_id`,`url`,`status`,`icon`,`creator`,`updater`,`create_date`,`last_update_date`) values ('01903','端口维护','019','/port/change','1','','DBA','DBA','2021-05-10 08:51:59','2021-05-10 08:51:59');
insert  into `t_xtqx`(`id`,`name`,`parent_id`,`url`,`status`,`icon`,`creator`,`updater`,`create_date`,`last_update_date`) values ('020','功能管理','0','','1','mdi mdi-function','DBA','DBA','2021-05-10 08:51:59','2021-05-10 08:51:59');
insert  into `t_xtqx`(`id`,`name`,`parent_id`,`url`,`status`,`icon`,`creator`,`updater`,`create_date`,`last_update_date`) values ('02001','功能查询','020','/func/query','1','','DBA','DBA','2021-05-10 08:51:59','2021-05-10 08:51:59');
insert  into `t_xtqx`(`id`,`name`,`parent_id`,`url`,`status`,`icon`,`creator`,`updater`,`create_date`,`last_update_date`) values ('02002','功能新增','020','/func/add','1','','DBA','DBA','2021-05-10 08:51:59','2021-05-10 08:51:59');
insert  into `t_xtqx`(`id`,`name`,`parent_id`,`url`,`status`,`icon`,`creator`,`updater`,`create_date`,`last_update_date`) values ('02003','功能变更','020','/func/change','1','','DBA','DBA','2021-05-10 08:51:59','2021-05-10 08:51:59');
insert  into `t_xtqx`(`id`,`name`,`parent_id`,`url`,`status`,`icon`,`creator`,`updater`,`create_date`,`last_update_date`) values ('021','图片管理','0','','1','ion-images','DBA','DBA','2021-05-10 08:51:59','2021-05-10 08:51:59');
insert  into `t_xtqx`(`id`,`name`,`parent_id`,`url`,`status`,`icon`,`creator`,`updater`,`create_date`,`last_update_date`) values ('02101','新建任务','021','/minio/add','1','','DBA','DBA','2021-05-10 08:51:59','2021-05-10 08:51:59');
insert  into `t_xtqx`(`id`,`name`,`parent_id`,`url`,`status`,`icon`,`creator`,`updater`,`create_date`,`last_update_date`) values ('02102','任务查询','021','/minio/query','1','','DBA','DBA','2021-05-10 08:51:59','2021-05-10 08:51:59');
insert  into `t_xtqx`(`id`,`name`,`parent_id`,`url`,`status`,`icon`,`creator`,`updater`,`create_date`,`last_update_date`) values ('02103','任务维护','021','/minio/change','1','','DBA','DBA','2021-05-10 08:51:59','2021-05-10 08:51:59');
insert  into `t_xtqx`(`id`,`name`,`parent_id`,`url`,`status`,`icon`,`creator`,`updater`,`create_date`,`last_update_date`) values ('02104','日志查询','021','/minio/log/query','1','','DBA','DBA','2021-05-10 08:51:59','2021-05-10 08:51:59');
insert  into `t_xtqx`(`id`,`name`,`parent_id`,`url`,`status`,`icon`,`creator`,`updater`,`create_date`,`last_update_date`) values ('02105','日志分析','021','/minio/log/analyze','1','','DBA','DBA','2021-05-10 08:51:59','2021-05-10 08:51:59');
insert  into `t_xtqx`(`id`,`name`,`parent_id`,`url`,`status`,`icon`,`creator`,`updater`,`create_date`,`last_update_date`) values ('060','公告管理','0','','0','mdi mdi-message-text-outline','DBA','DBA','2021-05-10 08:51:59','2021-05-10 08:51:59');
insert  into `t_xtqx`(`id`,`name`,`parent_id`,`url`,`status`,`icon`,`creator`,`updater`,`create_date`,`last_update_date`) values ('06001','公告发布','060','/message/release','1','','DBA','DBA','2021-05-10 08:51:59','2021-05-10 08:51:59');
insert  into `t_xtqx`(`id`,`name`,`parent_id`,`url`,`status`,`icon`,`creator`,`updater`,`create_date`,`last_update_date`) values ('06002','公告查询','060','/message/query','1','','DBA','DBA','2021-05-10 08:51:59','2021-05-10 08:51:59');
insert  into `t_xtqx`(`id`,`name`,`parent_id`,`url`,`status`,`icon`,`creator`,`updater`,`create_date`,`last_update_date`) values ('06003','公告变更','060','/message/change','1','','DBA','DBA','2021-05-10 08:51:59','2021-05-10 08:51:59');
insert  into `t_xtqx`(`id`,`name`,`parent_id`,`url`,`status`,`icon`,`creator`,`updater`,`create_date`,`last_update_date`) values ('061','系统管理','0','','1','mdi mdi-settings','DBA','DBA','2021-05-10 08:51:59','2021-05-10 08:51:59');
insert  into `t_xtqx`(`id`,`name`,`parent_id`,`url`,`status`,`icon`,`creator`,`updater`,`create_date`,`last_update_date`) values ('06101','系统设置','061','/sys/setting','1','','DBA','DBA','2021-05-10 08:51:59','2021-05-10 08:51:59');
insert  into `t_xtqx`(`id`,`name`,`parent_id`,`url`,`status`,`icon`,`creator`,`updater`,`create_date`,`last_update_date`) values ('06102','代码管理','061','/sys/code','1','mdi mdi-code-brackets','DBA','DBA','2021-05-10 08:51:59','2021-05-10 08:51:59');
insert  into `t_xtqx`(`id`,`name`,`parent_id`,`url`,`status`,`icon`,`creator`,`updater`,`create_date`,`last_update_date`) values ('06103','审核规则','061','/sys/audit_rule','1','mdi mdi-crop','DBA','DBA','2021-05-10 08:51:59','2021-05-10 08:51:59');

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
