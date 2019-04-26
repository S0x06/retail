/*
 Navicat Premium Data Transfer

 Source Server         : 127.0.0.1
 Source Server Type    : MySQL
 Source Server Version : 50553
 Source Host           : localhost:3306
 Source Schema         : retail

 Target Server Type    : MySQL
 Target Server Version : 50553
 File Encoding         : 65001

 Date: 26/04/2019 19:02:40
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for activity
-- ----------------------------
DROP TABLE IF EXISTS `activity`;
CREATE TABLE `activity`  (
  `id` int(11) UNSIGNED NOT NULL,
  `coupon_id` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '优惠卷id',
  `activity_type` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '活动类型',
  `start_time` timestamp NULL DEFAULT NULL COMMENT '活动开始时间',
  `end_time` timestamp NULL DEFAULT NULL COMMENT '活动结束时间',
  `issuer_id` int(11) NULL DEFAULT NULL COMMENT '商家id',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = MyISAM CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for activity_type
-- ----------------------------
DROP TABLE IF EXISTS `activity_type`;
CREATE TABLE `activity_type`  (
  `id` int(10) UNSIGNED NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = MyISAM CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Fixed;

-- ----------------------------
-- Table structure for coupon
-- ----------------------------
DROP TABLE IF EXISTS `coupon`;
CREATE TABLE `coupon`  (
  `id` int(11) UNSIGNED NOT NULL,
  `name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '优惠券',
  `issuer_id` int(11) NULL DEFAULT 0 COMMENT '发布的店家',
  `desc` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '描述',
  `quantity` int(11) NULL DEFAULT 0 COMMENT '发放数量 0无限 ',
  `duration` int(11) NULL DEFAULT 0 COMMENT '时长',
  `begin_time` timestamp NULL DEFAULT NULL COMMENT '活动开始时间',
  `status` tinyint(4) NULL DEFAULT 0 COMMENT '是否上线 0 否 1是',
  `delete` tinyint(4) NULL DEFAULT 0 COMMENT '是否删除 0 否 1是',
  `create_at` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '创建时间',
  `delete_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',
  `type_id` int(11) NULL DEFAULT 0 COMMENT '优惠券类型id',
  `limit` int(11) NULL DEFAULT 1 COMMENT '每个人可以领多少次 默认1',
  `channel` tinyint(4) NULL DEFAULT 0 COMMENT '发放渠道 0商家 1平台',
  `accumulate` int(11) NULL DEFAULT 1 COMMENT '是否累加  默认累加使用1张',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Compact;

-- ----------------------------
-- Table structure for coupon_log
-- ----------------------------
DROP TABLE IF EXISTS `coupon_log`;
CREATE TABLE `coupon_log`  (
  `id` int(11) NOT NULL,
  `conpon_id` int(11) NULL DEFAULT 0 COMMENT '优惠卷id',
  `user_id` int(11) NULL DEFAULT 0 COMMENT '用户id',
  `issuer_id` int(11) NULL DEFAULT NULL COMMENT '商家id',
  `desc` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `create_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = MyISAM CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for coupon_receive
-- ----------------------------
DROP TABLE IF EXISTS `coupon_receive`;
CREATE TABLE `coupon_receive`  (
  `id` int(10) UNSIGNED NOT NULL,
  `receive_id` int(11) NULL DEFAULT 0 COMMENT '用户id',
  `conpon_id` int(11) NULL DEFAULT 0 COMMENT '优惠卷id',
  `issuer_id` int(11) NULL DEFAULT 0 COMMENT '商家id',
  `used` tinyint(255) NULL DEFAULT 0 COMMENT '是否已使用 0否 1是',
  `create_at` timestamp NULL DEFAULT NULL COMMENT '领取时间',
  `delete` tinyint(255) NULL DEFAULT 0 COMMENT '删除0否 1是',
  `delete_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = MyISAM CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Fixed;

-- ----------------------------
-- Table structure for release
-- ----------------------------
DROP TABLE IF EXISTS `release`;
CREATE TABLE `release`  (
  `id` int(11) UNSIGNED NOT NULL,
  `crowd` int(255) NULL DEFAULT 0 COMMENT '人群 0不限 1男性 2女性',
  `area` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '0' COMMENT '地区 0 无地区限制',
  `service` tinyint(255) NULL DEFAULT 0 COMMENT '具体业务方，比如 1激活中心，2新手券，3下单券，用户生日券',
  `issuer_id` int(11) NULL DEFAULT 0 COMMENT '商家id',
  `create_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',
  `release_at` timestamp NULL DEFAULT NULL COMMENT '发布时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = MyISAM CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for type
-- ----------------------------
DROP TABLE IF EXISTS `type`;
CREATE TABLE `type`  (
  `id` int(11) UNSIGNED NOT NULL,
  `name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '类型名称',
  `status` tinyint(4) NULL DEFAULT 0 COMMENT '是否使用 0使用 1不使用',
  `create_at` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '创建时间',
  `delete` tinyint(4) NULL DEFAULT 0 COMMENT '是否删除 0否 1是',
  `delete_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '优惠券类型' ROW_FORMAT = Compact;

-- ----------------------------
-- Records of type
-- ----------------------------
INSERT INTO `type` VALUES (1, '满减券', 1, '2019-04-22 10:29:20', 0, NULL);
INSERT INTO `type` VALUES (2, '折扣券', 1, '2019-04-22 10:29:40', 0, NULL);

SET FOREIGN_KEY_CHECKS = 1;
