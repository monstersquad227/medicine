/*
 Navicat Premium Dump SQL

 Source Server         : LinkHub
 Source Server Type    : MySQL
 Source Server Version : 80025 (8.0.25)
 Source Host           : 47.103.98.61:3306
 Source Schema         : medicine

 Target Server Type    : MySQL
 Target Server Version : 80025 (8.0.25)
 File Encoding         : 65001

 Date: 03/11/2025 15:46:32
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for medicine_course
-- ----------------------------
DROP TABLE IF EXISTS `medicine_course`;
CREATE TABLE `medicine_course` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '自增ID',
  `user_id` int unsigned NOT NULL COMMENT '用户ID',
  `medicine_name` varchar(255) NOT NULL COMMENT '药物名称',
  `medicine_image` varchar(512) DEFAULT NULL COMMENT '药物图片',
  `medicine_type` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '药物方式：0-内服；1-外用',
  `medicine_timing` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '用药时机：0-不限；1-饭前用药；2-饭后用药；3-随餐用药；4-睡前用药',
  `course_start_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '用药开始时间',
  `status` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '方案状态：0-生效；1-废弃',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `idx_user_id` (`user_id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='用药方案表';

-- ----------------------------
-- Records of medicine_course
-- ----------------------------

-- ----------------------------
-- Table structure for medicine_plan
-- ----------------------------
DROP TABLE IF EXISTS `medicine_plan`;
CREATE TABLE `medicine_plan` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '自增ID',
  `medicine_id` int unsigned NOT NULL COMMENT '用药方案ID，关联 medicine_course.id',
  `amount` int NOT NULL COMMENT '用药数量',
  `type` varchar(16) NOT NULL COMMENT '剂量单位',
  `plan_time` varchar(64) NOT NULL COMMENT '用药时间',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `idx_medicine_id` (`medicine_id`),
  CONSTRAINT `fk_medicine_plan_course` FOREIGN KEY (`medicine_id`) REFERENCES `medicine_course` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='用药计划明细表';

-- ----------------------------
-- Records of medicine_plan
-- ----------------------------
BEGIN;
INSERT INTO `medicine_plan` (`id`, `medicine_id`, `amount`, `type`, `plan_time`, `created_at`, `updated_at`) VALUES (179, 139, 1, '包', '10:30', '2025-11-02 10:29:08', '2025-11-02 10:29:08');
INSERT INTO `medicine_plan` (`id`, `medicine_id`, `amount`, `type`, `plan_time`, `created_at`, `updated_at`) VALUES (184, 140, 2, '粒', '08:00', '2025-11-02 21:31:45', '2025-11-02 21:31:45');
INSERT INTO `medicine_plan` (`id`, `medicine_id`, `amount`, `type`, `plan_time`, `created_at`, `updated_at`) VALUES (185, 140, 2, '粒', '10:30', '2025-11-02 21:31:45', '2025-11-02 21:31:45');
INSERT INTO `medicine_plan` (`id`, `medicine_id`, `amount`, `type`, `plan_time`, `created_at`, `updated_at`) VALUES (186, 140, 2, '粒', '17:30', '2025-11-02 21:31:45', '2025-11-02 21:31:45');
COMMIT;

-- ----------------------------
-- Table structure for medicine_plan_record
-- ----------------------------
DROP TABLE IF EXISTS `medicine_plan_record`;
CREATE TABLE `medicine_plan_record` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '自增ID',
  `user_id` int unsigned NOT NULL COMMENT '用户ID',
  `plan_id` int unsigned DEFAULT NULL COMMENT '用药计划ID，关联 medicine_plan.id',
  `medicine_name` varchar(255) NOT NULL COMMENT '药物名称',
  `actual_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '实际用药时间',
  `memo` text COMMENT '打卡备注信息',
  `is_checked` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否打卡（0=未打卡，1=已打卡）',
  `status` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否异常（0: 正常；1: 异常），与计划时间相差15分钟记为异常',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '记录创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '记录更新时间',
  PRIMARY KEY (`id`),
  KEY `idx_user_id` (`user_id`),
  KEY `idx_plan_id` (`plan_id`),
  CONSTRAINT `fk_medicine_plan_record_medicine_plan` FOREIGN KEY (`plan_id`) REFERENCES `medicine_plan` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='用药打卡记录表';

-- ----------------------------
-- Records of medicine_plan_record
-- ----------------------------

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '自增ID',
  `phone_num` varchar(20) NOT NULL COMMENT '手机号',
  `push_token` varchar(128) DEFAULT NULL COMMENT 'PushKitToken',
  `notify_enabled` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否启用通知：1-启用，0-禁用',
  `huawei_id` varchar(100) DEFAULT NULL COMMENT '华为用户ID',
  `nickname` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '用户昵称',
  `password` varchar(255) NOT NULL COMMENT '用户密码（加密存储）',
  `image` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '用户头像 URL',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_phone_num` (`phone_num`),
  KEY `idx_huawei_id` (`huawei_id`)
) ENGINE=InnoDB AUTO_INCREMENT=100 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='用户表';

-- ----------------------------
-- Records of user
-- ----------------------------

SET FOREIGN_KEY_CHECKS = 1;
