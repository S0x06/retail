/*
Navicat MySQL Data Transfer

Source Server         : localhost
Source Server Version : 50528
Source Host           : localhost:3306
Source Database       : retail

Target Server Type    : MYSQL
Target Server Version : 50528
File Encoding         : 65001

Date: 2019-04-21 23:34:45
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for coupon
-- ----------------------------
DROP TABLE IF EXISTS `coupon`;
CREATE TABLE `coupon` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) DEFAULT NULL COMMENT '优惠券',
  `issuer_id` int(11) DEFAULT '0' COMMENT '发布的店家',
  `desc` varchar(255) DEFAULT NULL COMMENT '描述',
  `quantity` int(11) DEFAULT '0' COMMENT '发放数量 0无限 ',
  `duration` int(11) DEFAULT '0' COMMENT '时长',
  `begin_time` timestamp NULL DEFAULT NULL COMMENT '活动开始时间',
  `status` tinyint(4) DEFAULT '0' COMMENT '是否上线 0 否 1是',
  `delete` tinyint(4) DEFAULT '0' COMMENT '是否删除 0 否 1是',
  `create_at` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '创建时间',
  `delete_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',
  `type_id` int(11) DEFAULT '0' COMMENT '优惠券类型id',
  `amount` decimal(11,2) DEFAULT '0.00' COMMENT '抵扣金额',
  `limit` int(11) DEFAULT '1' COMMENT '每个人可以领多少次 默认1',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of coupon
-- ----------------------------

-- ----------------------------
-- Table structure for type
-- ----------------------------
DROP TABLE IF EXISTS `type`;
CREATE TABLE `type` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) DEFAULT NULL COMMENT '类型名称',
  `status` tinyint(4) DEFAULT '0' COMMENT '是否使用 0使用 1不使用',
  `create_at` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '创建时间',
  `delete` tinyint(4) DEFAULT '0' COMMENT '是否删除 0否 1是',
  `delete_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='优惠券类型';

-- ----------------------------
-- Records of type
-- ----------------------------
