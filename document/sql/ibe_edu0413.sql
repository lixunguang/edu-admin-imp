/*
Navicat MySQL Data Transfer

Source Server         : localhost_3306
Source Server Version : 80032
Source Host           : localhost:3306
Source Database       : ibe_edu

Target Server Type    : MYSQL
Target Server Version : 80032
File Encoding         : 65001

Date: 2023-04-13 17:04:36
*/

SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
DROP TABLE IF EXISTS `admin`;
CREATE TABLE `admin`
(
    `id`              int         NOT NULL AUTO_INCREMENT COMMENT '主键',
    `login_name`        varchar(50) NOT NULL DEFAULT '' COMMENT '登录id，唯一',
    `name`            varchar(50) NOT NULL DEFAULT '' COMMENT '实际姓名',
    `desc`    text COMMENT '介绍',
    `password`        varchar(50) NOT NULL DEFAULT '' COMMENT '密码',
    `organization_id` int         NOT NULL DEFAULT '0' COMMENT '组织id，包括学校院系班级',
    `phone_number`    varchar(50) NOT NULL DEFAULT '' COMMENT '电话号码',
    `is_deleted`      tinyint     NOT NULL DEFAULT '1' COMMENT '是否删除 1-否 2-是',
    `created_at`      timestamp   NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at`      timestamp   NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  AUTO_INCREMENT = 20
  DEFAULT CHARSET = utf8mb4 COMMENT ='管理员，包括超级管理员，管理员，可以登录后台进行编辑和管理';

-- ----------------------------
-- Records of teacher
-- ----------------------------

-- ----------------------------
-- Table structure for `class`
-- ----------------------------
DROP TABLE IF EXISTS `class`;
CREATE TABLE `class`
(
    `id`         int         NOT NULL AUTO_INCREMENT COMMENT '主键',
    `name`       varchar(50) NOT NULL DEFAULT '' COMMENT '班级名字',
    `created_at` timestamp   NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` timestamp   NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT ='班级表';

-- ----------------------------
-- Records of class
-- ----------------------------

-- ----------------------------
-- Table structure for `course`
-- ----------------------------
DROP TABLE IF EXISTS `course`;
CREATE TABLE `course`
(
    `id`         int          NOT NULL AUTO_INCREMENT COMMENT '主键',
    `title`      varchar(100) NOT NULL DEFAULT '0' COMMENT '课程标题',
    `desc`       text COMMENT '课程描述',
    `created_at` timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  AUTO_INCREMENT = 11
  DEFAULT CHARSET = utf8mb4 COMMENT ='课程表';

-- ----------------------------
-- Records of course
-- ----------------------------
INSERT INTO `course`
VALUES ('1', '课程1', '课程1 desc描述描述', '2023-03-21 14:38:04', '2023-03-21 14:38:04');
INSERT INTO `course`
VALUES ('2', '课程2', '课程2222222222222', '2023-03-21 14:39:38', '2023-03-21 14:39:38');
INSERT INTO `course`
VALUES ('3', '课程3', '课程33333333333333333333222', '2023-03-21 14:43:24', '2023-03-21 14:43:24');
INSERT INTO `course`
VALUES ('4', '课程4', '课程4444444444444', '2023-03-21 15:17:45', '2023-03-21 18:06:35');
INSERT INTO `course`
VALUES ('5', '课程8', '课程8888888888', '2023-03-29 14:06:17', '2023-03-29 14:06:17');
INSERT INTO `course`
VALUES ('6', '课程9', '课程9999', '2023-03-29 14:07:16', '2023-03-29 14:07:16');
INSERT INTO `course`
VALUES ('7', '课程13', '课程132', '2023-03-29 14:09:07', '2023-03-29 14:09:07');
INSERT INTO `course`
VALUES ('8', '课程12', '课程12132', '2023-03-29 14:11:58', '2023-03-29 14:11:58');
INSERT INTO `course`
VALUES ('9', '课程122', '课程12132', '2023-03-29 14:13:04', '2023-03-29 14:13:04');
INSERT INTO `course`
VALUES ('10', 'kecheng1',
        '课程12课程12132课程12132课程12132课程12132课程12132课程12132课程12132课程12132课程12132课程12132132课程12课程12132课程12132课程12132课程12132课程12132课程12132课程12132课程12132课程12132课程12132132课程12课程12132课程12132课程12132课程12132课程12132课程12132课程12132课程12132课程12132课程12132132课程12课程12132课程12132课程12132课程12132课程12132课程12132课程12132课程12132课程12132课程12132132',
        '2023-03-30 16:30:45', '2023-03-30 16:30:45');
INSERT INTO `course`
VALUES ('100', '塔身吊装实训课',
        '在电力建设施工过程中，抱杆是不可或缺的起重装置。它的主要功能是依托在建铁塔，将塔材提升至需要的位置进行施工安装，因此广泛应用于电力建设施工尤其是吊装施工并且是主要的受力单元。在施工过程中抱杆受力是否均衡合理，是否有超载使用，直接影响着施工的质量和安全。在传统的抱杆受力计算中，一般采用简化的方式，利用三角函数进行计算。然而这种简化方式只能粗略计算简单受力工况，当采用反向拉线等复杂工况或现场地形复杂时，不易分析各个结构的受力结果。对于现场施工人员，抱杆的力学计算所需的专业知识要求较高，繁琐并且无法保证可靠性，因此在施工中主要依靠工作经验来控制抱杆受力。',
        '2023-04-17 14:13:04', '2023-04-17 14:13:04');

-- ----------------------------
-- Table structure for `course_resource`
-- ----------------------------
DROP TABLE IF EXISTS `course_resource`;
CREATE TABLE `course_resource`
(
    `id`          int       NOT NULL AUTO_INCREMENT COMMENT '主键',
    `course_id`   int       NOT NULL DEFAULT '0' COMMENT '课程id',
    `resource_id` int       NOT NULL DEFAULT '0' COMMENT '资源id',
    `type`        tinyint   NOT NULL DEFAULT '0' COMMENT '资源用途，1展示图片，2 引导视频，3 课堂id',
    `index`       tinyint            DEFAULT NULL,
    `created_at`  timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at`  timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  AUTO_INCREMENT = 63
  DEFAULT CHARSET = utf8mb4 COMMENT ='课程-资源关系表';

-- ----------------------------
-- Records of course_resource
-- ----------------------------
INSERT INTO `course_resource`
VALUES ('1', '1', '9', '1', '0', '2023-03-22 10:56:59', '2023-03-22 11:16:47');
INSERT INTO `course_resource`
VALUES ('2', '1', '11', '2', '0', '2023-03-22 11:06:23', '2023-03-22 11:06:23');
INSERT INTO `course_resource`
VALUES ('3', '1', '1', '3', '1', '2023-03-22 11:12:49', '2023-03-23 16:01:40');
INSERT INTO `course_resource`
VALUES ('4', '1', '2', '3', '2', '2023-03-22 11:13:16', '2023-03-23 16:01:41');
INSERT INTO `course_resource`
VALUES ('5', '1', '3', '3', '3', '2023-03-22 11:13:22', '2023-03-23 16:01:43');
INSERT INTO `course_resource`
VALUES ('6', '1', '10', '2', '0', '2023-03-23 17:18:48', '2023-03-23 17:18:48');
INSERT INTO `course_resource`
VALUES ('7', '1', '11', '2', '0', '2023-03-30 17:45:23', '2023-04-06 16:56:50');
INSERT INTO `course_resource`
VALUES ('8', '2', '9', '1', '0', '2023-03-22 10:56:59', '2023-03-22 11:16:47');
INSERT INTO `course_resource`
VALUES ('9', '2', '11', '2', '0', '2023-03-22 11:06:23', '2023-03-22 11:06:23');
INSERT INTO `course_resource`
VALUES ('10', '2', '1', '3', '1', '2023-03-22 11:12:49', '2023-03-23 16:01:40');
INSERT INTO `course_resource`
VALUES ('11', '2', '2', '3', '2', '2023-03-22 11:13:16', '2023-03-23 16:01:41');
INSERT INTO `course_resource`
VALUES ('12', '2', '3', '3', '3', '2023-03-22 11:13:22', '2023-03-23 16:01:43');
INSERT INTO `course_resource`
VALUES ('13', '2', '10', '2', '0', '2023-03-23 17:18:48', '2023-03-23 17:18:48');
INSERT INTO `course_resource`
VALUES ('14', '2', '11', '2', '0', '2023-03-30 17:45:23', '2023-04-06 16:56:50');
INSERT INTO `course_resource`
VALUES ('15', '3', '9', '1', '0', '2023-03-22 10:56:59', '2023-03-22 11:16:47');
INSERT INTO `course_resource`
VALUES ('16', '3', '11', '2', '0', '2023-03-22 11:06:23', '2023-03-22 11:06:23');
INSERT INTO `course_resource`
VALUES ('17', '3', '1', '3', '1', '2023-03-22 11:12:49', '2023-03-23 16:01:40');
INSERT INTO `course_resource`
VALUES ('18', '3', '2', '3', '2', '2023-03-22 11:13:16', '2023-03-23 16:01:41');
INSERT INTO `course_resource`
VALUES ('19', '3', '3', '3', '3', '2023-03-22 11:13:22', '2023-03-23 16:01:43');
INSERT INTO `course_resource`
VALUES ('20', '3', '10', '2', '0', '2023-03-23 17:18:48', '2023-03-23 17:18:48');
INSERT INTO `course_resource`
VALUES ('21', '3', '11', '2', '0', '2023-03-30 17:45:23', '2023-04-06 16:56:50');
INSERT INTO `course_resource`
VALUES ('22', '4', '9', '1', '0', '2023-03-22 10:56:59', '2023-03-22 11:16:47');
INSERT INTO `course_resource`
VALUES ('23', '4', '11', '2', '0', '2023-03-22 11:06:23', '2023-03-22 11:06:23');
INSERT INTO `course_resource`
VALUES ('24', '4', '1', '3', '1', '2023-03-22 11:12:49', '2023-03-23 16:01:40');
INSERT INTO `course_resource`
VALUES ('25', '4', '2', '3', '2', '2023-03-22 11:13:16', '2023-03-23 16:01:41');
INSERT INTO `course_resource`
VALUES ('26', '4', '3', '3', '3', '2023-03-22 11:13:22', '2023-03-23 16:01:43');
INSERT INTO `course_resource`
VALUES ('27', '4', '10', '2', '0', '2023-03-23 17:18:48', '2023-03-23 17:18:48');
INSERT INTO `course_resource`
VALUES ('28', '4', '11', '2', '0', '2023-03-30 17:45:23', '2023-04-06 16:56:50');
INSERT INTO `course_resource`
VALUES ('29', '5', '9', '1', '0', '2023-03-22 10:56:59', '2023-03-22 11:16:47');
INSERT INTO `course_resource`
VALUES ('30', '5', '11', '2', '0', '2023-03-22 11:06:23', '2023-03-22 11:06:23');
INSERT INTO `course_resource`
VALUES ('31', '5', '1', '3', '1', '2023-03-22 11:12:49', '2023-03-23 16:01:40');
INSERT INTO `course_resource`
VALUES ('32', '5', '2', '3', '2', '2023-03-22 11:13:16', '2023-03-23 16:01:41');
INSERT INTO `course_resource`
VALUES ('33', '5', '3', '3', '3', '2023-03-22 11:13:22', '2023-03-23 16:01:43');
INSERT INTO `course_resource`
VALUES ('34', '5', '10', '2', '0', '2023-03-23 17:18:48', '2023-03-23 17:18:48');
INSERT INTO `course_resource`
VALUES ('35', '5', '11', '2', '0', '2023-03-30 17:45:23', '2023-04-06 16:56:50');
INSERT INTO `course_resource`
VALUES ('36', '6', '9', '1', '0', '2023-03-22 10:56:59', '2023-03-22 11:16:47');
INSERT INTO `course_resource`
VALUES ('37', '6', '11', '2', '0', '2023-03-22 11:06:23', '2023-03-22 11:06:23');
INSERT INTO `course_resource`
VALUES ('38', '6', '1', '3', '1', '2023-03-22 11:12:49', '2023-03-23 16:01:40');
INSERT INTO `course_resource`
VALUES ('39', '6', '2', '3', '2', '2023-03-22 11:13:16', '2023-03-23 16:01:41');
INSERT INTO `course_resource`
VALUES ('40', '6', '3', '3', '3', '2023-03-22 11:13:22', '2023-03-23 16:01:43');
INSERT INTO `course_resource`
VALUES ('41', '6', '10', '2', '0', '2023-03-23 17:18:48', '2023-03-23 17:18:48');
INSERT INTO `course_resource`
VALUES ('42', '6', '11', '2', '0', '2023-03-30 17:45:23', '2023-04-06 16:56:50');
INSERT INTO `course_resource`
VALUES ('43', '7', '9', '1', '0', '2023-03-22 10:56:59', '2023-03-22 11:16:47');
INSERT INTO `course_resource`
VALUES ('44', '7', '11', '2', '0', '2023-03-22 11:06:23', '2023-03-22 11:06:23');
INSERT INTO `course_resource`
VALUES ('45', '7', '1', '3', '1', '2023-03-22 11:12:49', '2023-03-23 16:01:40');
INSERT INTO `course_resource`
VALUES ('46', '7', '2', '3', '2', '2023-03-22 11:13:16', '2023-03-23 16:01:41');
INSERT INTO `course_resource`
VALUES ('47', '7', '3', '3', '3', '2023-03-22 11:13:22', '2023-03-23 16:01:43');
INSERT INTO `course_resource`
VALUES ('48', '7', '10', '2', '0', '2023-03-23 17:18:48', '2023-03-23 17:18:48');
INSERT INTO `course_resource`
VALUES ('49', '7', '11', '2', '0', '2023-03-30 17:45:23', '2023-04-06 16:56:50');
INSERT INTO `course_resource`
VALUES ('51', '8', '9', '1', '0', '2023-03-22 10:56:59', '2023-03-22 11:16:47');
INSERT INTO `course_resource`
VALUES ('52', '8', '11', '2', '0', '2023-03-22 11:06:23', '2023-03-22 11:06:23');
INSERT INTO `course_resource`
VALUES ('53', '8', '1', '3', '1', '2023-03-22 11:12:49', '2023-03-23 16:01:40');
INSERT INTO `course_resource`
VALUES ('54', '8', '2', '3', '2', '2023-03-22 11:13:16', '2023-03-23 16:01:41');
INSERT INTO `course_resource`
VALUES ('55', '8', '3', '3', '3', '2023-03-22 11:13:22', '2023-03-23 16:01:43');
INSERT INTO `course_resource`
VALUES ('56', '8', '10', '2', '0', '2023-03-23 17:18:48', '2023-03-23 17:18:48');
INSERT INTO `course_resource`
VALUES ('57', '8', '11', '2', '0', '2023-03-30 17:45:23', '2023-04-06 16:56:50');
INSERT INTO `course_resource`
VALUES ('61', '9', '9', '1', '0', '2023-03-22 10:56:59', '2023-03-22 11:16:47');
INSERT INTO `course_resource`
VALUES ('62', '9', '11', '2', '0', '2023-03-22 11:06:23', '2023-03-22 11:06:23');
INSERT INTO `course_resource`
VALUES ('63', '9', '1', '3', '1', '2023-03-22 11:12:49', '2023-03-23 16:01:40');
INSERT INTO `course_resource`
VALUES ('64', '9', '2', '3', '2', '2023-03-22 11:13:16', '2023-03-23 16:01:41');
INSERT INTO `course_resource`
VALUES ('65', '9', '3', '3', '3', '2023-03-22 11:13:22', '2023-03-23 16:01:43');
INSERT INTO `course_resource`
VALUES ('66', '9', '10', '2', '0', '2023-03-23 17:18:48', '2023-03-23 17:18:48');
INSERT INTO `course_resource`
VALUES ('67', '9', '11', '2', '0', '2023-03-30 17:45:23', '2023-04-06 16:56:50');

INSERT INTO `course_resource`
VALUES ('71', '10', '9', '1', '0', '2023-03-22 10:56:59', '2023-03-22 11:16:47');
INSERT INTO `course_resource`
VALUES ('72', '10', '11', '2', '0', '2023-03-22 11:06:23', '2023-03-22 11:06:23');
INSERT INTO `course_resource`
VALUES ('73', '10', '1', '3', '1', '2023-03-22 11:12:49', '2023-03-23 16:01:40');
INSERT INTO `course_resource`
VALUES ('74', '10', '2', '3', '2', '2023-03-22 11:13:16', '2023-03-23 16:01:41');
INSERT INTO `course_resource`
VALUES ('75', '10', '3', '3', '3', '2023-03-22 11:13:22', '2023-03-23 16:01:43');
INSERT INTO `course_resource`
VALUES ('76', '10', '10', '2', '0', '2023-03-23 17:18:48', '2023-03-23 17:18:48');
INSERT INTO `course_resource`
VALUES ('77', '10', '11', '2', '0', '2023-03-30 17:45:23', '2023-04-06 16:56:50');

INSERT INTO `course_resource`
VALUES ('100', '100', '100', '1', '0', '2023-03-30 17:45:23', '2023-04-06 16:56:50');
INSERT INTO `course_resource`
VALUES ('101', '100', '101', '3', '1', '2023-03-30 17:45:23', '2023-04-06 16:56:50');
INSERT INTO `course_resource`
VALUES ('102', '100', '102', '3', '2', '2023-03-30 17:45:23', '2023-04-06 16:56:50');
INSERT INTO `course_resource`
VALUES ('103', '100', '103', '3', '3', '2023-03-30 17:45:23', '2023-04-06 16:56:50');
INSERT INTO `course_resource`
VALUES ('104', '100', '104', '3', '4', '2023-03-30 17:45:23', '2023-04-06 16:56:50');
INSERT INTO `course_resource`
VALUES ('105', '100', '105', '3', '5', '2023-03-30 17:45:23', '2023-04-06 16:56:50');
INSERT INTO `course_resource`
VALUES ('106', '100', '106', '3', '6', '2023-03-30 17:45:23', '2023-04-06 16:56:50');
INSERT INTO `course_resource`
VALUES ('107', '100', '107', '3', '7', '2023-03-30 17:45:23', '2023-04-06 16:56:50');
INSERT INTO `course_resource`
VALUES ('108', '100', '108', '3', '8', '2023-03-30 17:45:23', '2023-04-06 16:56:50');
-- ----------------------------
-- Table structure for `course_student`
-- ----------------------------
DROP TABLE IF EXISTS `course_student`;
CREATE TABLE `course_student`
(
    `id`         int       NOT NULL AUTO_INCREMENT COMMENT '主键',
    `student_id` int       NOT NULL DEFAULT '0' COMMENT '学生id',
    `courese_id` int       NOT NULL DEFAULT '0' COMMENT '课程id',
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT ='学生-课程关系表';

-- ----------------------------
-- Records of course_student
-- ----------------------------

-- ----------------------------
-- Table structure for `course_teacher`
-- ----------------------------
DROP TABLE IF EXISTS `course_teacher`;
CREATE TABLE `course_teacher`
(
    `id`         int       NOT NULL AUTO_INCREMENT COMMENT '主键',
    `course_id`  int       NOT NULL DEFAULT '0' COMMENT '课程id',
    `teacher_id` int       NOT NULL DEFAULT '0' COMMENT '教师id',
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  AUTO_INCREMENT = 6
  DEFAULT CHARSET = utf8mb4 COMMENT ='课程-教师关系表';

-- ----------------------------
-- Records of course_teacher
-- ----------------------------
INSERT INTO `course_teacher`
VALUES ('1', '4', '14', '2023-03-22 17:33:21', '2023-03-22 17:33:21');
INSERT INTO `course_teacher`
VALUES ('2', '4', '12', '2023-03-22 18:32:43', '2023-03-22 18:32:43');
INSERT INTO `course_teacher`
VALUES ('3', '3', '19', '2023-03-22 18:33:08', '2023-03-22 18:33:08');
INSERT INTO `course_teacher`
VALUES ('4', '1', '19', '2023-03-23 15:17:01', '2023-03-23 15:17:01');
INSERT INTO `course_teacher`
VALUES ('5', '1', '12', '2023-03-23 15:17:07', '2023-03-23 15:17:07');

-- ----------------------------
-- Table structure for `department`
-- ----------------------------
DROP TABLE IF EXISTS `department`;
CREATE TABLE `department`
(
    `id`         int         NOT NULL AUTO_INCREMENT COMMENT '主键',
    `name`       varchar(50) NOT NULL DEFAULT '' COMMENT '院系名字',
    `created_at` timestamp   NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` timestamp   NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT ='院系表';

-- ----------------------------
-- Records of department
-- ----------------------------

-- ----------------------------
-- Table structure for `favorite`
-- ----------------------------
DROP TABLE IF EXISTS `favorite`;
CREATE TABLE `favorite`
(
    `id`          int       NOT NULL AUTO_INCREMENT COMMENT '主键',
    `user_id`     int       NOT NULL DEFAULT '0' COMMENT '所属用户id',
    `resource_id` int       NOT NULL DEFAULT '0' COMMENT '资源id',
    `useage`      tinyint   NOT NULL DEFAULT '0' COMMENT '资源所属类型，如资源，课程。',
    `created_at`  timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at`  timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT ='收藏表';

-- ----------------------------
-- Records of favorite
-- ----------------------------

-- ----------------------------
-- Table structure for `file`
-- ----------------------------
DROP TABLE IF EXISTS `file`;
CREATE TABLE `file`
(
    `id`            int           NOT NULL AUTO_INCREMENT COMMENT '主键',
    `original_name` varchar(1024) NOT NULL DEFAULT '' COMMENT '原始名字',
    `storage_name`  varchar(1024) NOT NULL DEFAULT '' COMMENT '存储名字',
    `created_at`    timestamp     NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at`    timestamp     NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  AUTO_INCREMENT = 3
  DEFAULT CHARSET = utf8mb4 COMMENT ='文件表，存储上传文件原始名字及存储路径';

-- ----------------------------
-- Records of file
-- ----------------------------
INSERT INTO `file`
VALUES ('1', 'ibe_edu20230329.sql', 'uploads/202304111117122714521008e20f8', '2023-04-11 11:18:18',
        '2023-04-11 11:18:18');
INSERT INTO `file`
VALUES ('2', '11.png', 'uploads/2023041111171227090160059b290', '2023-04-11 11:18:34', '2023-04-11 11:18:34');

-- ----------------------------
-- Table structure for `history`
-- ----------------------------
DROP TABLE IF EXISTS `history`;
CREATE TABLE `history`
(
    `id`            int       NOT NULL AUTO_INCREMENT COMMENT '主键',
    `user_id`       int       NOT NULL DEFAULT '0' COMMENT '所属用户id',
    `resource_id`   int       NOT NULL DEFAULT '0' COMMENT '资源id',
    `resource_type` tinyint   NOT NULL DEFAULT '0',
    `useage`        tinyint   NOT NULL DEFAULT '0' COMMENT '资源所属类型，如资源，课程。',
    `created_at`    timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at`    timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT ='历史记录表';

-- ----------------------------
-- Records of history
-- ----------------------------

-- ----------------------------
-- Table structure for `learning`
-- ----------------------------
DROP TABLE IF EXISTS `learning`;
CREATE TABLE `learning`
(
    `id`               int          NOT NULL AUTO_INCREMENT COMMENT '主键',
    `title`            varchar(100) NOT NULL DEFAULT '0' COMMENT '学习资源标题',
    `desc`             text COMMENT '学习资源描述',
    `author`           varchar(100) NOT NULL DEFAULT '0' COMMENT '作者',
    `cover_picture_id` int          NOT NULL DEFAULT '0' COMMENT '展示图片',
    `category_id`      int          NOT NULL DEFAULT '0' COMMENT '学习资源的分类',
    `created_at`       timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at`       timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  AUTO_INCREMENT = 15
  DEFAULT CHARSET = utf8mb4 COMMENT ='学习资源表';

-- ----------------------------
-- Records of learning
-- ----------------------------
INSERT INTO `learning`
VALUES ('1', '资源1', '资源1资源1资源1', 'xx', '6', '1', '2023-04-12 17:22:19', '2023-04-13 16:02:09');
INSERT INTO `learning`
VALUES ('2', '资源2', '资源2222', 'xx', '7', '1', '2023-04-12 17:22:51', '2023-04-13 16:02:11');
INSERT INTO `learning`
VALUES ('3', '资源3', '资源3资源3资源3', 'xx', '8', '1', '2023-04-12 17:31:31', '2023-04-13 16:01:39');
INSERT INTO `learning`
VALUES ('4', '资源4', '资源4资源4资源4', 'xx', '8', '2', '2023-04-12 17:31:31', '2023-04-13 16:01:39');
-- ----------------------------
-- Table structure for `learning_category`
-- ----------------------------
DROP TABLE IF EXISTS `learning_category`;
CREATE TABLE `learning_category`
(
    `id`         int          NOT NULL AUTO_INCREMENT COMMENT '主键',
    `title`      varchar(100) NOT NULL DEFAULT '0' COMMENT 'learning_category标题',
    `desc`       text COMMENT 'learning_category描述,可以为空',
    `created_at` timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  AUTO_INCREMENT = 14
  DEFAULT CHARSET = utf8mb4 COMMENT ='学习栏目分类表';

-- ----------------------------
-- Records of learning_category
-- ----------------------------
INSERT INTO `learning_category`
VALUES ('1', '基础视频', '基础视频基础视频基础视频', '2023-04-12 17:18:22', '2023-04-13 16:02:21');
INSERT INTO `learning_category`
VALUES ('2', '提高视频', '基础视频基础视频基础视频', '2023-04-12 17:19:29', '2023-04-13 16:02:24');
INSERT INTO `learning_category`
VALUES ('3', '进阶视频', '进阶视频进阶视频进阶视频', '2023-04-12 17:19:37', '2023-04-13 16:02:30');
INSERT INTO `learning_category`
VALUES ('4', '高级视频', '高级视频高级视频高级视频', '2023-04-12 17:19:37', '2023-04-13 16:02:30');
-- ----------------------------
-- Table structure for `learning_item`
-- ----------------------------
DROP TABLE IF EXISTS `learning_item`;
CREATE TABLE `learning_item`
(
    `id`          int          NOT NULL AUTO_INCREMENT COMMENT '主键',
    `title`       varchar(100) NOT NULL DEFAULT '0' COMMENT '学习资源标题',
    `desc`        text COMMENT '学习资源描述',
    `resource_id` int          NOT NULL DEFAULT '0' COMMENT '资源id，一般为视频',
    `learning_id` int          NOT NULL DEFAULT '0' COMMENT '所属learning id',
    `index`       int          NOT NULL DEFAULT '0' COMMENT '资源排序',
    `created_at`  timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at`  timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  AUTO_INCREMENT = 17
  DEFAULT CHARSET = utf8mb4 COMMENT ='学习资源条目表';

-- ----------------------------
-- Records of learning_item
-- ----------------------------
INSERT INTO `learning_item`
VALUES ('1', '条目2', '条目22222条目1条目1条目1', '11', '1', '2', '2023-04-12 17:37:34', '2023-04-12 17:37:34');
INSERT INTO `learning_item`
VALUES ('2', '条目1', '条目1条目1条目1条目1', '10', '1', '1', '2023-04-12 17:38:01', '2023-04-12 17:39:26');
INSERT INTO `learning_item`
VALUES ('3', '条目3', '条目3条目3条目3条目3', '12', '1', '3', '2023-04-12 17:39:16', '2023-04-12 17:39:16');
INSERT INTO `learning_item`
VALUES ('4', '条目4', '条目4条目4条目4条目4', '12', '1', '4', '2023-04-12 17:39:16', '2023-04-12 17:39:16');
INSERT INTO `learning_item`
VALUES ('5', '条目5', '条目5条目5条目5条目5', '12', '2', '1', '2023-04-12 17:39:16', '2023-04-12 17:39:16');
INSERT INTO `learning_item`
VALUES ('6', '条目6', '条目6条目6条目6条目6', '12', '2', '1', '2023-04-12 17:39:16', '2023-04-12 17:39:16');
-- ----------------------------
-- Table structure for `lesson`
-- ----------------------------
DROP TABLE IF EXISTS `lesson`;
CREATE TABLE `lesson`
(
    `id`         int          NOT NULL AUTO_INCREMENT COMMENT '主键',
    `title`      varchar(100) NOT NULL DEFAULT '0' COMMENT '课堂标题',
    `desc`       text COMMENT '课堂描述',
    `created_at` timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  AUTO_INCREMENT = 5
  DEFAULT CHARSET = utf8mb4 COMMENT ='课程-教师关系表';

-- ----------------------------
-- Records of lesson
-- ----------------------------
INSERT INTO `lesson`
VALUES ('1', '课程1-第1节课', '课程1-第1节课课程', '2023-03-23 16:35:47', '2023-04-03 12:28:30');
INSERT INTO `lesson`
VALUES ('2', '课程1-第2节课', '课程1-第2节课课程', '2023-03-23 16:36:18', '2023-04-03 12:28:31');
INSERT INTO `lesson`
VALUES ('3', '课程1-第3节课', '课程1-第3节课课程', '2023-03-23 16:36:24', '2023-04-03 12:28:34');
INSERT INTO `lesson`
VALUES ('4', '课程2-第1节课', '课程2-第1节课课程1-', '2023-03-23 16:36:33', '2023-03-23 16:36:33');

INSERT INTO `lesson`
VALUES ('101', '实验一：输电线路铁塔组立抱杆工程简介及仿真APP体验', '课程章节', '2023-03-23 16:36:33', '2023-03-23 16:36:33');

INSERT INTO `lesson`
VALUES ('102', '实验二：Simdroid软件梁壳基本操作学习实训指导书', '课程章节', '2023-03-23 16:36:33', '2023-03-23 16:36:33');

INSERT INTO `lesson`
VALUES ('103', '实验三：塔身吊装-参数化建模实训指导书', '课程章节', '2023-03-23 16:36:33', '2023-03-23 16:36:33');

INSERT INTO `lesson`
VALUES ('104', '实验四：塔身吊装-网格剖分实训指导书模板', '课程章节', '2023-03-23 16:36:33', '2023-03-23 16:36:33');
INSERT INTO `lesson`
VALUES ('105', '实验五：塔身吊装-边界与求解设置实训指导书', '课程章节', '2023-03-23 16:36:33', '2023-03-23 16:36:33');
INSERT INTO `lesson`
VALUES ('106', '实验六：塔身吊装-结果后处理实训指导书', '课程章节', '2023-03-23 16:36:33', '2023-03-23 16:36:33');
INSERT INTO `lesson`
VALUES ('107', '实验七：塔身吊装-APP封装实训指导书', '课程章节', '2023-03-23 16:36:33', '2023-03-23 16:36:33');
INSERT INTO `lesson`
VALUES ('108', '实验八：塔身吊装- APP实操', '课程章节', '2023-03-23 16:36:33', '2023-03-23 16:36:33');

-- ----------------------------
-- Table structure for `lesson_section`
-- ----------------------------
DROP TABLE IF EXISTS `lesson_section`;
CREATE TABLE `lesson_section`
(
    `id`         int          NOT NULL AUTO_INCREMENT COMMENT '主键',
    `lesson_id`  int          NOT NULL DEFAULT '0' COMMENT '课堂id',
    `title`      varchar(100) NOT NULL DEFAULT '0' COMMENT '课堂子部分标题',
    `index`      tinyint      NOT NULL,
    `created_at` timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  AUTO_INCREMENT = 11
  DEFAULT CHARSET = utf8mb4 COMMENT ='课堂的各个组成部分标题';

-- ----------------------------
-- Records of lesson_section
-- ----------------------------
INSERT INTO `lesson_section`
VALUES ('1', '1', '教学内容', '1', '2023-03-31 15:44:50', '2023-03-31 15:47:29');
INSERT INTO `lesson_section`
VALUES ('2', '1', '参考资源', '2', '2023-03-31 15:45:48', '2023-03-31 15:47:58');
INSERT INTO `lesson_section`
VALUES ('3', '1', '实验资源', '3', '2023-03-31 15:46:34', '2023-03-31 15:49:28');
INSERT INTO `lesson_section`
VALUES ('4', '1', '作业', '4', '2023-03-31 15:46:37', '2023-03-31 15:49:31');
INSERT INTO `lesson_section`
VALUES ('5', '2', '教学内容', '1', '2023-03-31 15:46:39', '2023-03-31 15:49:39');
INSERT INTO `lesson_section`
VALUES ('6', '2', '参考资源', '2', '2023-03-31 15:46:39', '2023-03-31 15:47:44');
INSERT INTO `lesson_section`
VALUES ('7', '2', '实验资源', '3', '2023-03-31 15:46:40', '2023-03-31 15:49:43');
INSERT INTO `lesson_section`
VALUES ('8', '3', '教学内容', '1', '2023-03-31 15:46:40', '2023-03-31 16:22:29');
INSERT INTO `lesson_section`
VALUES ('9', '4', 'XX', '1', '2023-04-03 13:01:34', '2023-04-04 10:46:55');
INSERT INTO `lesson_section`
VALUES ('10', '4', '作业', '2', '2023-04-03 13:02:07', '2023-04-04 10:46:53');

-- ----------------------------
-- Table structure for `lesson_section_resource`
-- ----------------------------
DROP TABLE IF EXISTS `lesson_section_resource`;
CREATE TABLE `lesson_section_resource`
(
    `id`                int       NOT NULL AUTO_INCREMENT COMMENT '主键',
    `lesson_section_id` int       NOT NULL DEFAULT '0' COMMENT 'lesson_section_id id',
    `resource_id`       int       NOT NULL DEFAULT '0' COMMENT '资源id',
    `index`             tinyint   NOT NULL DEFAULT '0' COMMENT '排序',
    `created_at`        timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at`        timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  AUTO_INCREMENT = 16
  DEFAULT CHARSET = utf8mb4 COMMENT ='课堂部分-资源关系表';

-- ----------------------------
-- Records of lesson_section_resource
-- ----------------------------
INSERT INTO `lesson_section_resource`
VALUES ('1', '1', '14', '1', '2023-03-28 18:16:11', '2023-03-28 18:16:11');
INSERT INTO `lesson_section_resource`
VALUES ('2', '1', '16', '2', '2023-03-28 18:25:50', '2023-03-28 18:25:50');
INSERT INTO `lesson_section_resource`
VALUES ('3', '1', '17', '3', '2023-03-28 18:26:14', '2023-03-28 18:26:14');
INSERT INTO `lesson_section_resource`
VALUES ('4', '2', '10', '1', '2023-03-28 18:26:27', '2023-04-03 12:06:07');
INSERT INTO `lesson_section_resource`
VALUES ('7', '1', '19', '4', '2023-03-28 18:44:30', '2023-03-31 18:34:21');
INSERT INTO `lesson_section_resource`
VALUES ('8', '2', '11', '2', '2023-03-28 18:44:37', '2023-03-28 18:44:37');
INSERT INTO `lesson_section_resource`
VALUES ('9', '2', '12', '3', '2023-03-28 18:44:42', '2023-03-28 18:44:42');
INSERT INTO `lesson_section_resource`
VALUES ('10', '3', '5', '3', '2023-03-28 18:46:21', '2023-03-31 18:08:53');
INSERT INTO `lesson_section_resource`
VALUES ('11', '4', '20', '41', '2023-03-28 18:46:28', '2023-04-10 11:18:16');
INSERT INTO `lesson_section_resource`
VALUES ('13', '4', '3', '42', '2023-04-03 12:32:16', '2023-04-04 10:50:14');

-- ----------------------------
-- Table structure for `mark`
-- ----------------------------
DROP TABLE IF EXISTS `mark`;
CREATE TABLE `mark`
(
    `id`               int       NOT NULL AUTO_INCREMENT COMMENT '主键',
    `lesson_id`        int       NOT NULL DEFAULT '0' COMMENT '课堂lesson id',
    `score_item_id`    int       NOT NULL DEFAULT '0' COMMENT '打分项目id',
    `score_subitem_id` int       NOT NULL DEFAULT '0' COMMENT '打分子项id',
    `point`            int       NOT NULL DEFAULT '0' COMMENT '子项分',
    `created_at`       timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at`       timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT ='打分表';

-- ----------------------------
-- Records of mark
-- ----------------------------

-- ----------------------------
-- Table structure for `news`
-- ----------------------------
DROP TABLE IF EXISTS `news`;
CREATE TABLE `news`
(
    `id`         int          NOT NULL AUTO_INCREMENT COMMENT '主键',
    `title`      varchar(100) NOT NULL DEFAULT '0' COMMENT '新闻标题',
    `content`    text COMMENT '新闻内容',
    `picture_id` int                   DEFAULT NULL,
    `publisher`  varchar(100) NOT NULL DEFAULT '0' COMMENT '发布人',
    `created_at` timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  AUTO_INCREMENT = 32
  DEFAULT CHARSET = utf8mb4 COMMENT ='新闻动态表';

-- ----------------------------
-- Records of news
-- ----------------------------
INSERT INTO `news`
VALUES ('1', '新闻通知1', '111', '0', 'admin', '2023-03-14 11:06:09', '2023-03-15 17:39:40');
INSERT INTO `news`
VALUES ('2', '新闻通知2', 'ppppppppppppppppppppppppdiv', '0', 'admin', '2023-03-14 11:06:53', '2023-03-15 17:39:41');
INSERT INTO `news`
VALUES ('3', '新闻通知3', '<p style=\"text-align: center;\"><strong>第一条新闻</strong></p>', '8', 'admin',
        '2023-03-14 11:12:18', '2023-03-30 15:35:45');
INSERT INTO `news`
VALUES ('4', '新闻通知4',
        '<p style=\"text-align: center;\"><strong>第一条新闻</strong></p><p style=\"text-align: center;\"><span style=\"color: #e03e2d;\"><em><strong>09:34:03 2023-03-14</strong></em></span></p><p>&nbsp; &nbsp;新闻内容段落1</p><p style=\"padding-left: 40px;\"><span style=\"font-family: \'arial black\', sans-serif;\">情况1</span></p><p style=\"padding-left: 40px;\"><span style=\"font-family: \'arial black\', sans-serif;\">情况2</span></p><p style=\"padding-left: 40px;\"><span style=\"font-family: \'arial black\', sans-serif;\">情况3</span></p><p>&nbsp;<span style=\"text-decoration: line-through;\"> &nbsp;新闻内容段落2</span></p><table style=\"border-collapse: collapse; width: 100%;\" border=\"1\"><tbody><tr><td style=\"width: 23.2395%;\">1</td><td style=\"width: 23.2395%;\">2</td><td style=\"width: 23.2395%;\">&nbsp;</td><td style=\"width: 23.2395%;\">&nbsp;</td></tr><tr><td style=\"width: 23.2395%;\">&nbsp;</td><td style=\"width: 23.2395%;\">3</td><td style=\"width: 23.2395%;\">4</td><td style=\"width: 23.2395%;\">4</td></tr><tr><td style=\"width: 23.2395%;\">&nbsp;</td><td style=\"width: 23.2395%;\">5</td><td style=\"width: 23.2395%;\">5</td><td style=\"width: 23.2395%;\">5</td></tr></tbody></table><p><span style=\"text-decoration: underline;\">&nbsp; &nbsp;新闻内容段落3</span></p><p><span style=\"text-decoration: underline;\"><img src=\"https://7072-prod-8g06gkn605f1c56a-1301228508.tcb.qcloud.la/resource/image/1.png\" alt=\"存储在nginx里面的图片\" width=\"300\" height=\"200\" /></span></p><p>&nbsp;</p><p><span style=\"text-decoration: underline;\">图片2</span></p><p><a href=\"http://www.baidu.com\"><strong><u>百度网站</u></strong></a></p>',
        '0', 'admin', '2023-03-15 09:08:28', '2023-03-15 17:39:44');
INSERT INTO `news`
VALUES ('5', '新闻通知4',
        '<p style=\"text-align: center;\"><strong>第2条新闻</strong></p><p style=\"text-align: center;\"><span style=\"color: #e03e2d;\"><em><strong>09:34:03 2023-03-14</strong></em></span></p><p>&nbsp; &nbsp;新闻内容段落1</p><p style=\"padding-left: 40px;\"><span style=\"font-family: \'arial black\', sans-serif;\">情况1</span></p><p style=\"padding-left: 40px;\"><span style=\"font-family: \'arial black\', sans-serif;\">情况2</span></p><p style=\"padding-left: 40px;\"><span style=\"font-family: \'arial black\', sans-serif;\">情况3</span></p><p>&nbsp;<span style=\"text-decoration: line-through;\"> &nbsp;新闻内容段落2</span></p><table style=\"border-collapse: collapse; width: 100%;\" border=\"1\"><tbody><tr><td style=\"width: 23.2395%;\">1</td><td style=\"width: 23.2395%;\">2</td><td style=\"width: 23.2395%;\">&nbsp;</td><td style=\"width: 23.2395%;\">&nbsp;</td></tr><tr><td style=\"width: 23.2395%;\">&nbsp;</td><td style=\"width: 23.2395%;\">3</td><td style=\"width: 23.2395%;\">4</td><td style=\"width: 23.2395%;\">4</td></tr><tr><td style=\"width: 23.2395%;\">&nbsp;</td><td style=\"width: 23.2395%;\">5</td><td style=\"width: 23.2395%;\">5</td><td style=\"width: 23.2395%;\">5</td></tr></tbody></table><p><span style=\"text-decoration: underline;\">&nbsp; &nbsp;新闻内容段落3</span></p><p><span style=\"text-decoration: underline;\"><img src=\"https://7072-prod-8g06gkn605f1c56a-1301228508.tcb.qcloud.la/resource/image/1.png\" alt=\"存储在nginx里面的图片\" width=\"300\" height=\"200\" /></span></p><p>&nbsp;</p><p><span style=\"text-decoration: underline;\">图片2</span></p><p><a href=\"http://www.baidu.com\"><strong><u>百度网站</u></strong></a></p>',
        '0', 'admin', '2023-03-15 09:09:34', '2023-03-15 17:39:45');
INSERT INTO `news`
VALUES ('6', '新闻通知5',
        '<p style=\"text-align: center;\"><strong>第3条新闻</strong></p><p style=\"text-align: center;\"><span style=\"color: #e03e2d;\"><em><strong>09:34:03 2023-03-14</strong></em></span></p><p>&nbsp; &nbsp;新闻内容段落1</p><p style=\"padding-left: 40px;\"><span style=\"font-family: \'arial black\', sans-serif;\">情况1</span></p><p style=\"padding-left: 40px;\"><span style=\"font-family: \'arial black\', sans-serif;\">情况2</span></p><p style=\"padding-left: 40px;\"><span style=\"font-family: \'arial black\', sans-serif;\">情况3</span></p><p>&nbsp;<span style=\"text-decoration: line-through;\"> &nbsp;新闻内容段落2</span></p><table style=\"border-collapse: collapse; width: 100%;\" border=\"1\"><tbody><tr><td style=\"width: 23.2395%;\">1</td><td style=\"width: 23.2395%;\">2</td><td style=\"width: 23.2395%;\">&nbsp;</td><td style=\"width: 23.2395%;\">&nbsp;</td></tr><tr><td style=\"width: 23.2395%;\">&nbsp;</td><td style=\"width: 23.2395%;\">3</td><td style=\"width: 23.2395%;\">4</td><td style=\"width: 23.2395%;\">4</td></tr><tr><td style=\"width: 23.2395%;\">&nbsp;</td><td style=\"width: 23.2395%;\">5</td><td style=\"width: 23.2395%;\">5</td><td style=\"width: 23.2395%;\">5</td></tr></tbody></table><p><span style=\"text-decoration: underline;\">&nbsp; &nbsp;新闻内容段落3</span></p><p><span style=\"text-decoration: underline;\"><img src=\"https://7072-prod-8g06gkn605f1c56a-1301228508.tcb.qcloud.la/resource/image/1.png\" alt=\"存储在nginx里面的图片\" width=\"300\" height=\"200\" /></span></p><p>&nbsp;</p><p><span style=\"text-decoration: underline;\">图片2</span></p><p><a href=\"http://www.baidu.com\"><strong><u>百度网站</u></strong></a></p>',
        '7', 'admin', '2023-03-15 09:11:03', '2023-03-30 15:31:46');
INSERT INTO `news`
VALUES ('7', '新闻通知6',
        '<p style=\"text-align: center;\"><strong>第4条新闻</strong></p><p style=\"text-align: center;\"><span style=\"color: #e03e2d;\"><em><strong>09:34:03 2023-03-14</strong></em></span></p><p>&nbsp; &nbsp;新闻内容段落1</p><p style=\"padding-left: 40px;\"><span style=\"font-family: \'arial black\', sans-serif;\">情况1</span></p><p style=\"padding-left: 40px;\"><span style=\"font-family: \'arial black\', sans-serif;\">情况2</span></p><p style=\"padding-left: 40px;\"><span style=\"font-family: \'arial black\', sans-serif;\">情况3</span></p><p>&nbsp;<span style=\"text-decoration: line-through;\"> &nbsp;新闻内容段落2</span></p><table style=\"border-collapse: collapse; width: 100%;\" border=\"1\"><tbody><tr><td style=\"width: 23.2395%;\">1</td><td style=\"width: 23.2395%;\">2</td><td style=\"width: 23.2395%;\">&nbsp;</td><td style=\"width: 23.2395%;\">&nbsp;</td></tr><tr><td style=\"width: 23.2395%;\">&nbsp;</td><td style=\"width: 23.2395%;\">3</td><td style=\"width: 23.2395%;\">4</td><td style=\"width: 23.2395%;\">4</td></tr><tr><td style=\"width: 23.2395%;\">&nbsp;</td><td style=\"width: 23.2395%;\">5</td><td style=\"width: 23.2395%;\">5</td><td style=\"width: 23.2395%;\">5</td></tr></tbody></table><p><span style=\"text-decoration: underline;\">&nbsp; &nbsp;新闻内容段落3</span></p><p><span style=\"text-decoration: underline;\"><img src=\"https://7072-prod-8g06gkn605f1c56a-1301228508.tcb.qcloud.la/resource/image/1.png\" alt=\"存储在nginx里面的图片\" width=\"300\" height=\"200\" /></span></p><p>&nbsp;</p><p><span style=\"text-decoration: underline;\">图片2</span></p><p><a href=\"http://www.baidu.com\"><strong><u>百度网站</u></strong></a></p>',
        '1', 'admin', '2023-03-15 09:11:08', '2023-03-16 17:15:04');
INSERT INTO `news`
VALUES ('8', 'news', 'abcdefghijklmnopqrstuvwxyz', '6', 'admin', '2023-03-17 10:01:00', '2023-03-17 10:01:00');
INSERT INTO `news`
VALUES ('9', '新闻通知9', '111', '0', 'admin', '2023-03-14 11:06:09', '2023-03-15 17:39:40');
INSERT INTO `news`
VALUES ('10', '新闻通知10', '111', '0', 'admin', '2023-03-14 11:06:09', '2023-03-15 17:39:40');
INSERT INTO `news`
VALUES ('11', '新闻通知11', '111', '0', 'admin', '2023-03-14 11:06:09', '2023-03-15 17:39:40');
INSERT INTO `news`
VALUES ('12', '新闻通知12', '111', '0', 'admin', '2023-03-14 11:06:09', '2023-03-15 17:39:40');
INSERT INTO `news`
VALUES ('13', '新闻通知13', '111', '0', 'admin', '2023-03-14 11:06:09', '2023-03-15 17:39:40');
INSERT INTO `news`
VALUES ('14', '新闻通知14', '111', '0', 'admin', '2023-03-14 11:06:09', '2023-03-15 17:39:40');
INSERT INTO `news`
VALUES ('15', '新闻通知15', '111', '0', 'admin', '2023-03-14 11:06:09', '2023-03-15 17:39:40');
INSERT INTO `news`
VALUES ('16', '新闻通知16', '111', '0', 'admin', '2023-03-14 11:06:09', '2023-03-15 17:39:40');
INSERT INTO `news`
VALUES ('17', '新闻通知17', '111', '0', 'admin', '2023-03-14 11:06:09', '2023-03-15 17:39:40');
INSERT INTO `news`
VALUES ('18', '新闻通知18', '111', '0', 'admin', '2023-03-14 11:06:09', '2023-03-15 17:39:40');
INSERT INTO `news`
VALUES ('19', '新闻通知19', '111', '0', 'admin', '2023-03-14 11:06:09', '2023-03-15 17:39:40');
INSERT INTO `news`
VALUES ('20', '新闻通知20', '111', '0', 'admin', '2023-03-14 11:06:09', '2023-03-15 17:39:40');
INSERT INTO `news`
VALUES ('21', '新闻通知21', '111', '0', 'admin', '2023-03-14 11:06:09', '2023-03-15 17:39:40');
INSERT INTO `news`
VALUES ('22', '新闻通知22', '111', '0', 'admin', '2023-03-14 11:06:09', '2023-03-15 17:39:40');
INSERT INTO `news`
VALUES ('23', '新闻通知23', '111', '0', 'admin', '2023-03-14 11:06:09', '2023-03-15 17:39:40');
INSERT INTO `news`
VALUES ('24', '新闻通知24', '111', '0', 'admin', '2023-03-14 11:06:09', '2023-03-15 17:39:40');
INSERT INTO `news`
VALUES ('25', '新闻通知25', '111', '0', 'admin', '2023-03-14 11:06:09', '2023-03-15 17:39:40');
INSERT INTO `news`
VALUES ('26', '新闻通知26', '111', '0', 'admin', '2023-03-14 11:06:09', '2023-03-15 17:39:40');
INSERT INTO `news`
VALUES ('27', '新闻通知27', '111', '0', 'admin', '2023-03-14 11:06:09', '2023-03-15 17:39:40');
INSERT INTO `news`
VALUES ('28', '新闻通知28', '111', '0', 'admin', '2023-03-14 11:06:09', '2023-03-15 17:39:40');
INSERT INTO `news`
VALUES ('29', '新闻通知29', '111', '0', 'admin', '2023-03-14 11:06:09', '2023-03-15 17:39:40');
INSERT INTO `news`
VALUES ('30', '新闻通知30', '111', '0', 'admin', '2023-03-14 11:06:09', '2023-03-15 17:39:40');
INSERT INTO `news`
VALUES ('31', '新闻通知31', '111', '0', 'admin', '2023-03-14 11:06:09', '2023-03-15 17:39:40');

INSERT INTO `news`
VALUES ('100', '我校虚拟仿真实验室建成启动', '<p><strong>本网讯 &nbsp;</strong>我校虚拟仿真实验室建成启用，新建成的虚拟仿真实验室是目前全国同类院校中建设标准和规格最高的实验室。3月22日，我校校长仲伟合一行参观了该实验室并就实验室应用、口译信息化教学等问题进行了深入的研讨。</p>
<p>&nbsp;</p>
<p>图片</p>
<p>&nbsp;</p>
<p>实验室利用先进的虚拟仿真和虚拟现实技术与多媒体设备相结合，根据外语教学的特点，形成了独具特色的外语虚拟仿真和虚拟现实的实践中心。实验室有四大功能：情景教学、情景微课录制、课堂录播、虚拟现实实训，全方位满足学生实践技能提升和教学资源建设两大方面的需求，同时对于创新课堂评估方式、提高学生积极性以及丰富翻转课堂实践等方面都有很大帮助。国家级同声传译实验教学中心副主任古煜奎和厂家代表冯大治对实验室的功能和应用进行了介绍。</p>
<p>&nbsp;</p>
<p>图片</p>
', '101', 'admin', '2023-04-17 11:06:09', '2023-04-17 17:39:40');
INSERT INTO `news`
VALUES ('101', '构建未来：建设国家虚拟仿真实验室的重要性与挑战', '<p id=\"1LB4TVGA\">随着科技的不断发展，虚拟仿真技术逐渐成为一个备受瞩目的领域。国家虚拟仿真实验室的建设，也成为了现代化国家的必然选择。</p>
<p id=\"1LB4TVGB\">首先，国家虚拟仿真实验室的建设具有重要意义。虚拟仿真技术可以模拟现实场景，使人们在虚拟环境中体验真实世界的各种情境，从而提高人们的认知水平。这项技术在军事、航空、医疗、教育等领域具有广泛的应用前景，可以有效地提高生产效率，降低事故风险，缩短研发周期，降低成本等方面的优势。</p>
<p>图片1</p>
<p id=\"1LB4TVGE\">其次，建设国家虚拟仿真实验室也面临着许多挑战。其中最大的挑战是技术方面的问题。虚拟仿真技术涉及到许多领域的知识，需要各种专业人才的参与。此外，虚拟仿真技术的成本较高，需要大量的投入。建设国家虚拟仿真实验室需要有政府的大力支持和专业的人才支持。</p>
<p id=\"1LB4TVGF\">在面对这些挑战时，政府需要采取积极的措施。政府可以制定相关政策，支持虚拟仿真技术的研发和应用。政府还可以通过投资，吸引专业人才的参与，提高技术水平。此外，政府还可以积极引导企业和机构参与虚拟仿真技术的研究和应用，形成产学研合作的良好机制。</p>
<p>图片2</p>
<p>2019年1月31号，教育部高等教育司发布了《关于2018年度国家虚拟仿真实验教学项目名单》的公示，共有298个项目入选，来自全国186所高校。这也是继2018年6月教育部公布首批105个国家级虚拟现实仿真实验教学项目后，教育部再次认定这一重量级教学项目。其中北京欧倍尔助力30所高校成功申报2018年度国家虚拟仿真实验教学项目，在同行业中居于领先水平，彰显了公司雄厚的技术实力。</p>
<p>图片3</p>
<p id=\"1LB4TVGL\">从入选名单来看，本次共有298项目进入公示，共涉及23个学科领域。机械类教学项目最多，共有30个，数量最多。临床医学类共有29个，位居第二位。生物科学类共有20个，位居第三位。化工与制药类共有18个，数量位居第四。医学基础类、药学类、动物类、能源动力类、食品科学与工程类、土木类等数量也较多。</p>
<p id=\"1LB4TVGM\">建设国家虚拟仿真实验室是一个具有重要意义的事业。在克服技术方面的困难和面对各种挑战的同时，政府和社会各界应积极支持和参与，共同推进虚拟仿真技术的发展和应用，为实现现代化建设和经济发展做出贡献。</p>',
        '102', 'admin', '2023-04-17 11:06:09', '2023-04-17 17:39:40');
INSERT INTO `news`
VALUES ('102', '虚拟仿真实验项目助力线上实验教学', '<p>&nbsp; &nbsp; &nbsp; &nbsp;5月3日晚，材料学部董超芳与实验中心黄鹏、商北雁使用材料科学与工程学院实验中心建设的材料虚拟仿真项目，为2020级材料化学专业25名学生在线进行了《材料科学基础实验C》课程线上虚拟实验教学，内容包括Fe-Fe3C相图与铁碳合金平衡组织显微观察、金相制样与显微观察等。学生课后表示，虚拟仿真实验很新奇也很好操作，还能在需要注意的实验操作上给我一些提示，并且可以让我很直观地看到实验结果。教师们积极使用虚拟仿真实验项目，以保障疫情防控期间实验教学顺利开展。</p>
<p>图片</p>
<p>&nbsp;</p>
<p>&nbsp; &nbsp; &nbsp; &nbsp;材料科学与工程学院实验中心自2007年获批国家级实验教学示范中心和2014年国家级虚拟仿真实验教学中心以来，在教育部与学校支持下，建成了虚拟仿真实验项目1.0版和2.0版。学校材料科学与工程学院、能源与环境工程学院、高等工程师学院等学院都已在课程教学、专业实习、学科竞赛中使用该项目，累计使用超过1万余人次。2017年，国内首次开设校级公选课程《材料虚拟实验》，并编写了课程讲义。2020年疫情期间项目免费向学校和国内高校开放，有力地支持了疫情期间的线上实验教学。同时，在大学生金相实验技能大赛中，创建了线上预赛、线下决赛，虚实结合比赛的新形式，为疫情期间学科竞赛提供了新赛道。其中虚拟仿真实验项目1.0版《系列材料虚拟实验》于2021年荣获第六届全国高等学校教师创新大赛一等奖。虚拟仿真实验项目2.0版&ldquo;金属材料成分、工艺、组织和性能一体化设计虚拟仿真实验&rdquo;于2021年6月申报国家级虚拟仿真一流课程，并建立了虚拟仿真沉浸式体验实验室。2021年8月，实验中心在国内率先牵头成立材料虚仿实验教学虚拟教研室，目前已有80余所国内高校加入，开展的教研活动获得了良好效果，带动了各高校推进融合式实验教学改革。2021年12月，项目获批教育部产学合作协同育人项目&ldquo;建设实验教学虚拟教研室&rdquo;，打造了材料类虚拟仿真项目共享应用云基地。​</p>
<p>图片</p>
<p>&nbsp; &nbsp; &nbsp; &nbsp;截至2022年5月3日，虚拟仿真实验1.0版已共享到国内国防科技大学、中山大学、青海大学等28所高校。2020年疫情出现以来，该版本支持了哈尔滨工业大学（威海）、烟台大学、江西科技师范大学、北华航天工业学院等高校开展线上实验教学。随着疫情防控常态化，山东大学、湖南大学、安徽大学等高校也开始采用虚拟仿真实验2.0版，通过虚拟仿真一流课开设了8门实验课。</p>
<p>&nbsp; &nbsp; &nbsp; &nbsp;材料科学与工程学院实验中心发挥双国家级实验教学中心示范引领优势，积极探索虚实结合实验教学新模式，深度挖掘我校优质实验教学资源，丰富虚拟仿真实验项目内容，进一步拓展推广应用途径，有力支持了我校乃至国内高校材料及相关专业的本科实验教学，推动了材料专业跨校、跨地区的虚拟仿真实验教学项目共建共享，开拓了疫情防控条件下实验教学新模式。</p>',
        '103', 'admin', '2023-04-17 11:06:09', '2023-04-17 17:39:40');

-- ----------------------------
-- Table structure for `organization`
-- ----------------------------
DROP TABLE IF EXISTS `organization`;
CREATE TABLE `organization`
(
    `id`            int       NOT NULL AUTO_INCREMENT COMMENT '主键',
    `school_id`     int       NOT NULL DEFAULT '0' COMMENT 'school id',
    `department_id` int       NOT NULL DEFAULT '0' COMMENT 'department id',
    `class_id`      int       NOT NULL DEFAULT '0' COMMENT 'class id',
    `created_at`    timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at`    timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT ='学校组织表';

-- ----------------------------
-- Records of organization
-- ----------------------------

-- ----------------------------
-- Table structure for `resource`
-- ----------------------------
DROP TABLE IF EXISTS `resource`;
CREATE TABLE `resource`
(
    `id`         int          NOT NULL AUTO_INCREMENT COMMENT '主键',
    `title`      varchar(100) NOT NULL DEFAULT '0' COMMENT '资源标题',
    `desc`       text COMMENT '资源描述',
    `type`       tinyint      NOT NULL DEFAULT '0' COMMENT '资源类型，1图片，2 word文档，3 pdf文档，4 apps，5 ibe，6留言7 富文本',
    `content`    tinytext COMMENT '资源内容，存储路径、url及少量的文字',
    `created_at` timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  AUTO_INCREMENT = 36
  DEFAULT CHARSET = utf8mb4 COMMENT ='存储系统中用到的基础资源';

-- ----------------------------
-- Records of resource
-- ----------------------------
INSERT INTO `resource`
VALUES ('1', '一个图片', '图片平铺', '1', '/image/1.webp', '2023-03-16 16:00:29', '2023-03-30 15:43:45');
INSERT INTO `resource`
VALUES ('2', '一个文件', '文件', '1', '/files/1.pdf', '2023-03-16 16:02:38', '2023-03-30 15:47:40');
INSERT INTO `resource`
VALUES ('3', '一个word文件', 'word文件', '2', '/files/1.docx', '2023-03-16 16:03:35', '2023-03-30 15:47:38');
INSERT INTO `resource`
VALUES ('4', '一个apps', 'apps', '4', '/ibe/1.app', '2023-03-16 16:04:29', '2023-03-30 15:47:25');
INSERT INTO `resource`
VALUES ('5', 'ibe文件', 'ibe', '5', '/ibe/1.ibe', '2023-03-16 16:19:17', '2023-03-30 15:47:16');
INSERT INTO `resource`
VALUES ('6', '第二个图片', '图片', '1', '/image/2.webp', '2023-03-17 09:26:35', '2023-03-30 15:43:43');
INSERT INTO `resource`
VALUES ('7', '第san个图片', '图片', '1', '/image/3.webp', '2023-03-17 11:25:51', '2023-03-30 15:43:47');
INSERT INTO `resource`
VALUES ('8', '第si个图片', '图片', '1', '/image/4.webp', '2023-03-17 11:26:46', '2023-03-30 15:43:48');
INSERT INTO `resource`
VALUES ('9', '第5个图片', '图片', '1', '/image/5.webp', '2023-03-17 11:30:37', '2023-03-30 15:43:53');
INSERT INTO `resource`
VALUES ('10', '第一个个视频', '视频x', '8', '/vedio/1.mp4', '2023-03-22 11:02:00', '2023-03-30 15:47:59');
INSERT INTO `resource`
VALUES ('11', '第2个个视频', '视频y', '8', '/vedio/2.mp4', '2023-03-22 11:03:02', '2023-03-30 15:44:32');
INSERT INTO `resource`
VALUES ('12', '第3个个shipin', '视频z', '8', '/vedio/3.mp4', '2023-03-22 11:04:00', '2023-03-30 15:44:22');
INSERT INTO `resource`
VALUES ('14', '实验目的', '富文本1', '7', '2', '2023-03-28 18:07:30', '2023-04-06 17:56:24');
INSERT INTO `resource`
VALUES ('15', '富文本2', '富文本2', '7', '3', '2023-03-28 18:08:40', '2023-04-06 17:54:35');
INSERT INTO `resource`
VALUES ('16', '实验要求', '富文本3', '7', '4', '2023-03-28 18:08:45', '2023-04-06 17:56:43');
INSERT INTO `resource`
VALUES ('17', '实验步骤和要点', '富文本4', '7', '5', '2023-03-28 18:08:52', '2023-04-06 17:56:52');
INSERT INTO `resource`
VALUES ('18', '注意事项', '富文本5', '7', '6', '2023-03-28 18:08:57', '2023-04-06 17:56:59');
INSERT INTO `resource`
VALUES ('19', '富文本6', '富文本6', '7', '7', '2023-03-28 18:09:04', '2023-04-06 17:54:46');
INSERT INTO `resource`
VALUES ('20', '作业要求1', '作业要求1', '7', '8', '2023-04-10 11:15:01', '2023-04-10 11:15:01');
INSERT INTO `resource`
VALUES ('21', '作业提交内容', '作业提交内容', '7', '9', '2023-04-10 14:48:01', '2023-04-10 14:53:21');
INSERT INTO `resource`
VALUES ('22', '未使用', '作业要求1', '7', '10', '2023-04-10 14:48:03', '2023-04-10 14:53:26');
INSERT INTO `resource`
VALUES ('23', '未使用', '作业要求1', '7', '11', '2023-04-10 14:48:04', '2023-04-10 14:53:32');
INSERT INTO `resource`
VALUES ('24', 'ibe_edu20230329.sql', '作业文件', '0', 'uploads/202304111117122714521008e20f8', '2023-04-11 11:18:31',
        '2023-04-11 11:18:31');
INSERT INTO `resource`
VALUES ('25', '11.png', '作业文件', '0', 'uploads/2023041111171227090160059b290', '2023-04-11 11:18:37',
        '2023-04-11 11:18:37');
INSERT INTO `resource`
VALUES ('26', '富文本', '富文本', '7', '13', '2023-04-11 15:46:31', '2023-04-11 15:46:31');
INSERT INTO `resource`
VALUES ('27', '作业要求a', '作业要求a', '7', '15', '2023-04-11 15:49:37', '2023-04-11 15:49:37');
INSERT INTO `resource`
VALUES ('28', '作业要求b', '作业要求b', '7', '16', '2023-04-11 15:50:12', '2023-04-11 15:50:12');
INSERT INTO `resource`
VALUES ('29', '作业要求b', '作业要求b', '7', '18', '2023-04-11 15:52:46', '2023-04-11 15:52:46');
INSERT INTO `resource`
VALUES ('30', '富文本', '富文本', '7', '19', '2023-04-11 16:17:56', '2023-04-11 16:17:56');
INSERT INTO `resource`
VALUES ('31', '富文本', '富文本', '7', '20', '2023-04-11 16:20:15', '2023-04-11 16:20:15');
INSERT INTO `resource`
VALUES ('32', 'ibe_edu20230329.sql', '作业文件', '0', 'uploads/202304121755165063266008e20f8', '2023-04-12 17:55:19',
        '2023-04-12 17:55:19');
INSERT INTO `resource`
VALUES ('33', 'ibe_edu20230329.sql', '作业文件', '0', 'uploads/202304121755403778440008e20f8', '2023-04-12 17:55:41',
        '2023-04-12 17:55:41');
INSERT INTO `resource`
VALUES ('34', '11.png', '作业文件', '0', 'uploads/2023041218005058953050059b290', '2023-04-12 18:00:52',
        '2023-04-12 18:00:52');
INSERT INTO `resource`
VALUES ('35', 'ibe_edu20230329.sql', '作业文件', '0', 'uploads/202304121800505901451008e20f8', '2023-04-12 18:00:52',
        '2023-04-12 18:00:52');

INSERT INTO `resource`
VALUES ('100', '抱杆展示图片', '图片', '1', '/image/抱杆.png', '2023-04-12 18:00:52',
        '2023-04-12 18:00:52');
INSERT INTO `resource`
VALUES ('101', 'banner展示图片', '图片', '1', '/image/banner1.png', '2023-04-12 18:00:52',
        '2023-04-12 18:00:52');

INSERT INTO `resource`
VALUES ('102', 'banner展示图片', '图片', '1', '/image/banner2.png', '2023-04-12 18:00:52',
        '2023-04-12 18:00:52');

INSERT INTO `resource`
VALUES ('103', 'banner展示图片', '图片', '1', '/image/banner3.png', '2023-04-12 18:00:52',
        '2023-04-12 18:00:52');



-- ----------------------------
-- Table structure for `resource_richtext`
-- ----------------------------
DROP TABLE IF EXISTS `resource_richtext`;
CREATE TABLE `resource_richtext`
(
    `id`         int       NOT NULL AUTO_INCREMENT COMMENT '主键',
    `content`    text COMMENT '基础资源富文本内容',
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  AUTO_INCREMENT = 21
  DEFAULT CHARSET = utf8mb4 COMMENT ='存储系统中用到的富文本资源';

-- ----------------------------
-- Records of resource_richtext
-- ----------------------------
INSERT INTO `resource_richtext`
VALUES ('2',
        '<p style=\"text-align: center;\"><strong>富文本1</strong></p><p style=\"text-align: center;\"><span style=\"color: #e03e2d;\"><em><strong>09:34:03 2023-03-14</strong></em></span></p><p>&nbsp; &nbsp;新闻内容段落1</p><p style=\"padding-left: 40px;\"><span style=\"font-family: \'arial black\', sans-serif;\">情况1</span></p><p style=\"padding-left: 40px;\"><span style=\"font-family: \'arial black\', sans-serif;\">情况2</span></p><p style=\"padding-left: 40px;\"><span style=\"font-family: \'arial black\', sans-serif;\">情况3</span></p><p>&nbsp;<span style=\"text-decoration: line-through;\"> &nbsp;新闻内容段落2</span></p><table style=\"border-collapse: collapse; width: 100%;\" border=\"1\"><tbody><tr><td style=\"width: 23.2395%;\">1</td><td style=\"width: 23.2395%;\">2</td><td style=\"width: 23.2395%;\">&nbsp;</td><td style=\"width: 23.2395%;\">&nbsp;</td></tr><tr><td style=\"width: 23.2395%;\">&nbsp;</td><td style=\"width: 23.2395%;\">3</td><td style=\"width: 23.2395%;\">4</td><td style=\"width: 23.2395%;\">4</td></tr><tr><td style=\"width: 23.2395%;\">&nbsp;</td><td style=\"width: 23.2395%;\">5</td><td style=\"width: 23.2395%;\">5</td><td style=\"width: 23.2395%;\">5</td></tr></tbody></table><p><span style=\"text-decoration: underline;\">&nbsp; &nbsp;新闻内容段落3</span></p><p><span style=\"text-decoration: underline;\"><img src=\"https://7072-prod-8g06gkn605f1c56a-1301228508.tcb.qcloud.la/resource/image/1.png\" alt=\"存储在nginx里面的图片\" width=\"300\" height=\"200\" /></span></p><p>&nbsp;</p><p><span style=\"text-decoration: underline;\">图片2</span></p><p><a href=\"http://www.baidu.com\"><strong><u>百度网站</u></strong></a></p>',
        '2023-03-28 18:07:30', '2023-03-28 18:07:30');
INSERT INTO `resource_richtext`
VALUES ('3',
        '<p style=\"text-align: center;\"><strong>富文本2</strong></p><p style=\"text-align: center;\"><span style=\"color: #e03e2d;\"><em><strong>09:34:03 2023-03-14</strong></em></span></p><p>&nbsp; &nbsp;新闻内容段落1</p><p style=\"padding-left: 40px;\"><span style=\"font-family: \'arial black\', sans-serif;\">情况1</span></p><p style=\"padding-left: 40px;\"><span style=\"font-family: \'arial black\', sans-serif;\">情况2</span></p><p style=\"padding-left: 40px;\"><span style=\"font-family: \'arial black\', sans-serif;\">情况3</span></p><p>&nbsp;<span style=\"text-decoration: line-through;\"> &nbsp;新闻内容段落2</span></p><table style=\"border-collapse: collapse; width: 100%;\" border=\"1\"><tbody><tr><td style=\"width: 23.2395%;\">1</td><td style=\"width: 23.2395%;\">2</td><td style=\"width: 23.2395%;\">&nbsp;</td><td style=\"width: 23.2395%;\">&nbsp;</td></tr><tr><td style=\"width: 23.2395%;\">&nbsp;</td><td style=\"width: 23.2395%;\">3</td><td style=\"width: 23.2395%;\">4</td><td style=\"width: 23.2395%;\">4</td></tr><tr><td style=\"width: 23.2395%;\">&nbsp;</td><td style=\"width: 23.2395%;\">5</td><td style=\"width: 23.2395%;\">5</td><td style=\"width: 23.2395%;\">5</td></tr></tbody></table><p><span style=\"text-decoration: underline;\">&nbsp; &nbsp;新闻内容段落3</span></p><p><span style=\"text-decoration: underline;\"><img src=\"https://7072-prod-8g06gkn605f1c56a-1301228508.tcb.qcloud.la/resource/image/1.png\" alt=\"存储在nginx里面的图片\" width=\"300\" height=\"200\" /></span></p><p>&nbsp;</p><p><span style=\"text-decoration: underline;\">图片2</span></p><p><a href=\"http://www.baidu.com\"><strong><u>百度网站</u></strong></a></p>',
        '2023-03-28 18:08:40', '2023-03-28 18:08:40');
INSERT INTO `resource_richtext`
VALUES ('4',
        '<p style=\"text-align: center;\"><strong>富文本2</strong></p><p style=\"text-align: center;\"><span style=\"color: #e03e2d;\"><em><strong>09:34:03 2023-03-14</strong></em></span></p><p>&nbsp; &nbsp;新闻内容段落1</p><p style=\"padding-left: 40px;\"><span style=\"font-family: \'arial black\', sans-serif;\">情况1</span></p><p style=\"padding-left: 40px;\"><span style=\"font-family: \'arial black\', sans-serif;\">情况2</span></p><p style=\"padding-left: 40px;\"><span style=\"font-family: \'arial black\', sans-serif;\">情况3</span></p><p>&nbsp;<span style=\"text-decoration: line-through;\"> &nbsp;新闻内容段落2</span></p><table style=\"border-collapse: collapse; width: 100%;\" border=\"1\"><tbody><tr><td style=\"width: 23.2395%;\">1</td><td style=\"width: 23.2395%;\">2</td><td style=\"width: 23.2395%;\">&nbsp;</td><td style=\"width: 23.2395%;\">&nbsp;</td></tr><tr><td style=\"width: 23.2395%;\">&nbsp;</td><td style=\"width: 23.2395%;\">3</td><td style=\"width: 23.2395%;\">4</td><td style=\"width: 23.2395%;\">4</td></tr><tr><td style=\"width: 23.2395%;\">&nbsp;</td><td style=\"width: 23.2395%;\">5</td><td style=\"width: 23.2395%;\">5</td><td style=\"width: 23.2395%;\">5</td></tr></tbody></table><p><span style=\"text-decoration: underline;\">&nbsp; &nbsp;新闻内容段落3</span></p><p><span style=\"text-decoration: underline;\"><img src=\"https://7072-prod-8g06gkn605f1c56a-1301228508.tcb.qcloud.la/resource/image/1.png\" alt=\"存储在nginx里面的图片\" width=\"300\" height=\"200\" /></span></p><p>&nbsp;</p><p><span style=\"text-decoration: underline;\">图片2</span></p><p><a href=\"http://www.baidu.com\"><strong><u>百度网站</u></strong></a></p>',
        '2023-03-28 18:08:45', '2023-03-28 18:08:45');
INSERT INTO `resource_richtext`
VALUES ('5',
        '<p style=\"text-align: center;\"><strong>富文本4</strong></p><p style=\"text-align: center;\"><span style=\"color: #e03e2d;\"><em><strong>09:34:03 2023-03-14</strong></em></span></p><p>&nbsp; &nbsp;新闻内容段落1</p><p style=\"padding-left: 40px;\"><span style=\"font-family: \'arial black\', sans-serif;\">情况1</span></p><p style=\"padding-left: 40px;\"><span style=\"font-family: \'arial black\', sans-serif;\">情况2</span></p><p style=\"padding-left: 40px;\"><span style=\"font-family: \'arial black\', sans-serif;\">情况3</span></p><p>&nbsp;<span style=\"text-decoration: line-through;\"> &nbsp;新闻内容段落2</span></p><table style=\"border-collapse: collapse; width: 100%;\" border=\"1\"><tbody><tr><td style=\"width: 23.2395%;\">1</td><td style=\"width: 23.2395%;\">2</td><td style=\"width: 23.2395%;\">&nbsp;</td><td style=\"width: 23.2395%;\">&nbsp;</td></tr><tr><td style=\"width: 23.2395%;\">&nbsp;</td><td style=\"width: 23.2395%;\">3</td><td style=\"width: 23.2395%;\">4</td><td style=\"width: 23.2395%;\">4</td></tr><tr><td style=\"width: 23.2395%;\">&nbsp;</td><td style=\"width: 23.2395%;\">5</td><td style=\"width: 23.2395%;\">5</td><td style=\"width: 23.2395%;\">5</td></tr></tbody></table><p><span style=\"text-decoration: underline;\">&nbsp; &nbsp;新闻内容段落3</span></p><p><span style=\"text-decoration: underline;\"><img src=\"https://7072-prod-8g06gkn605f1c56a-1301228508.tcb.qcloud.la/resource/image/1.png\" alt=\"存储在nginx里面的图片\" width=\"300\" height=\"200\" /></span></p><p>&nbsp;</p><p><span style=\"text-decoration: underline;\">图片2</span></p><p><a href=\"http://www.baidu.com\"><strong><u>百度网站</u></strong></a></p>',
        '2023-03-28 18:08:52', '2023-03-28 18:10:25');
INSERT INTO `resource_richtext`
VALUES ('6',
        '<p style=\"text-align: center;\"><strong>富文本5</strong></p><p style=\"text-align: center;\"><span style=\"color: #e03e2d;\"><em><strong>09:34:03 2023-03-14</strong></em></span></p><p>&nbsp; &nbsp;新闻内容段落1</p><p style=\"padding-left: 40px;\"><span style=\"font-family: \'arial black\', sans-serif;\">情况1</span></p><p style=\"padding-left: 40px;\"><span style=\"font-family: \'arial black\', sans-serif;\">情况2</span></p><p style=\"padding-left: 40px;\"><span style=\"font-family: \'arial black\', sans-serif;\">情况3</span></p><p>&nbsp;<span style=\"text-decoration: line-through;\"> &nbsp;新闻内容段落2</span></p><table style=\"border-collapse: collapse; width: 100%;\" border=\"1\"><tbody><tr><td style=\"width: 23.2395%;\">1</td><td style=\"width: 23.2395%;\">2</td><td style=\"width: 23.2395%;\">&nbsp;</td><td style=\"width: 23.2395%;\">&nbsp;</td></tr><tr><td style=\"width: 23.2395%;\">&nbsp;</td><td style=\"width: 23.2395%;\">3</td><td style=\"width: 23.2395%;\">4</td><td style=\"width: 23.2395%;\">4</td></tr><tr><td style=\"width: 23.2395%;\">&nbsp;</td><td style=\"width: 23.2395%;\">5</td><td style=\"width: 23.2395%;\">5</td><td style=\"width: 23.2395%;\">5</td></tr></tbody></table><p><span style=\"text-decoration: underline;\">&nbsp; &nbsp;新闻内容段落3</span></p><p><span style=\"text-decoration: underline;\"><img src=\"https://7072-prod-8g06gkn605f1c56a-1301228508.tcb.qcloud.la/resource/image/1.png\" alt=\"存储在nginx里面的图片\" width=\"300\" height=\"200\" /></span></p><p>&nbsp;</p><p><span style=\"text-decoration: underline;\">图片2</span></p><p><a href=\"http://www.baidu.com\"><strong><u>百度网站</u></strong></a></p>',
        '2023-03-28 18:08:57', '2023-03-28 18:10:17');
INSERT INTO `resource_richtext`
VALUES ('7',
        '<p style=\"text-align: center;\"><strong>富文本6</strong></p><p style=\"text-align: center;\"><span style=\"color: #e03e2d;\"><em><strong>09:34:03 2023-03-14</strong></em></span></p><p>&nbsp; &nbsp;新闻内容段落1</p><p style=\"padding-left: 40px;\"><span style=\"font-family: \'arial black\', sans-serif;\">情况1</span></p><p style=\"padding-left: 40px;\"><span style=\"font-family: \'arial black\', sans-serif;\">情况2</span></p><p style=\"padding-left: 40px;\"><span style=\"font-family: \'arial black\', sans-serif;\">情况3</span></p><p>&nbsp;<span style=\"text-decoration: line-through;\"> &nbsp;新闻内容段落2</span></p><table style=\"border-collapse: collapse; width: 100%;\" border=\"1\"><tbody><tr><td style=\"width: 23.2395%;\">1</td><td style=\"width: 23.2395%;\">2</td><td style=\"width: 23.2395%;\">&nbsp;</td><td style=\"width: 23.2395%;\">&nbsp;</td></tr><tr><td style=\"width: 23.2395%;\">&nbsp;</td><td style=\"width: 23.2395%;\">3</td><td style=\"width: 23.2395%;\">4</td><td style=\"width: 23.2395%;\">4</td></tr><tr><td style=\"width: 23.2395%;\">&nbsp;</td><td style=\"width: 23.2395%;\">5</td><td style=\"width: 23.2395%;\">5</td><td style=\"width: 23.2395%;\">5</td></tr></tbody></table><p><span style=\"text-decoration: underline;\">&nbsp; &nbsp;新闻内容段落3</span></p><p><span style=\"text-decoration: underline;\"><img src=\"https://7072-prod-8g06gkn605f1c56a-1301228508.tcb.qcloud.la/resource/image/1.png\" alt=\"存储在nginx里面的图片\" width=\"300\" height=\"200\" /></span></p><p>&nbsp;</p><p><span style=\"text-decoration: underline;\">图片2</span></p><p><a href=\"http://www.baidu.com\"><strong><u>百度网站</u></strong></a></p>',
        '2023-03-28 18:09:04', '2023-03-28 18:09:04');
INSERT INTO `resource_richtext`
VALUES ('8',
        '作业提交说明,\n要求文件名清晰明确，作业的提交形式包括：\n1）截图（*.jpg,*.png）\n2）word文档(*.doc,*.docx)\n3）simdroid工程文件（*.ibe）\n4)simdroid apps(*.app)',
        '2023-04-10 11:15:01', '2023-04-10 11:15:01');
INSERT INTO `resource_richtext`
VALUES ('9', '这是一段学生提交的作业说明', '2023-04-10 14:48:01', '2023-04-10 14:49:37');
INSERT INTO `resource_richtext`
VALUES ('10', '作业提交说明,\n要未使用', '2023-04-10 14:48:03', '2023-04-10 14:53:48');
INSERT INTO `resource_richtext`
VALUES ('11', '作业提交说明,\n未使用', '2023-04-10 14:48:04', '2023-04-10 14:53:54');
INSERT INTO `resource_richtext`
VALUES ('12', '作业提交 作业提交 这是一段富文本', '2023-04-11 15:12:10', '2023-04-11 15:12:10');
INSERT INTO `resource_richtext`
VALUES ('18',
        '作业提交说明b,\n要求文件名清晰明确，作业的提交形式包括：\n1）截图（*.jpg,*.png）\n2）word文档(*.doc,*.docx)\n3）simdroid工程文件（*.ibe）\n4)simdroid apps(*.app)',
        '2023-04-11 15:52:36', '2023-04-11 15:52:36');
INSERT INTO `resource_richtext`
VALUES ('19', '作业提交2 作业提交 这是一段富文本', '2023-04-11 16:17:55', '2023-04-11 16:17:55');
INSERT INTO `resource_richtext`
VALUES ('20', '作业提交2 作业提交 这是一段富文本', '2023-04-11 16:20:15', '2023-04-11 16:20:15');

-- ----------------------------
-- Table structure for `school`
-- ----------------------------
DROP TABLE IF EXISTS `school`;
CREATE TABLE `school`
(
    `id`         int         NOT NULL AUTO_INCREMENT COMMENT '主键',
    `name`       varchar(50) NOT NULL DEFAULT '' COMMENT '学校名字',
    `created_at` timestamp   NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` timestamp   NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT ='学校表';

-- ----------------------------
-- Records of school
-- ----------------------------

-- ----------------------------
-- Table structure for `score`
-- ----------------------------
DROP TABLE IF EXISTS `score`;
CREATE TABLE `score`
(
    `id`         int       NOT NULL AUTO_INCREMENT COMMENT '主键',
    `student_id` int       NOT NULL DEFAULT '0' COMMENT '学生id',
    `course_id`  int       NOT NULL DEFAULT '0' COMMENT '课程id',
    `work_id`    int       NOT NULL DEFAULT '0' COMMENT '课堂作业id',
    `mark_id`    int       NOT NULL DEFAULT '0' COMMENT '打分id',
    `score`      int       NOT NULL DEFAULT '0' COMMENT '课堂得分',
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT ='课堂成绩关系表';

-- ----------------------------
-- Records of score
-- ----------------------------

-- ----------------------------
-- Table structure for `score_item`
-- ----------------------------
DROP TABLE IF EXISTS `score_item`;
CREATE TABLE `score_item`
(
    `id`         int       NOT NULL AUTO_INCREMENT COMMENT '主键',
    `item`       int       NOT NULL DEFAULT '0' COMMENT '打分项',
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT ='打分项表';

-- ----------------------------
-- Records of score_item
-- ----------------------------

-- ----------------------------
-- Table structure for `score_subitem`
-- ----------------------------
DROP TABLE IF EXISTS `score_subitem`;
CREATE TABLE `score_subitem`
(
    `id`         int       NOT NULL AUTO_INCREMENT COMMENT '主键',
    `subitem`    int       NOT NULL DEFAULT '0' COMMENT '打分子项',
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT ='打分子项表';

-- ----------------------------
-- Records of score_subitem
-- ----------------------------
-- ----------------------------
-- Table structure for `teacher`
-- ----------------------------
DROP TABLE IF EXISTS `teacher`;
CREATE TABLE `teacher`
(
    `id`              int         NOT NULL AUTO_INCREMENT COMMENT '主键',
    `login_name`        varchar(50) NOT NULL DEFAULT '' COMMENT '登录id，唯一',
    `name`            varchar(50) NOT NULL DEFAULT '' COMMENT '实际姓名',
    `desc`    text COMMENT '介绍',
    `password`        varchar(50) NOT NULL DEFAULT '' COMMENT '密码',
    `organization_id` int         NOT NULL DEFAULT '0' COMMENT '组织id，包括学校院系班级',
    `phone_number`    varchar(50) NOT NULL DEFAULT '' COMMENT '电话号码',
    `is_deleted`      tinyint     NOT NULL DEFAULT '1' COMMENT '是否删除 1-否 2-是',
    `created_at`      timestamp   NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at`      timestamp   NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  AUTO_INCREMENT = 20
  DEFAULT CHARSET = utf8mb4 COMMENT ='教师表，编写或者讲解课程的作者';

-- ----------------------------
-- Records of teacher
-- ----------------------------

-- ----------------------------
-- Table structure for `user`
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`
(
    `id`              int         NOT NULL AUTO_INCREMENT COMMENT '主键',
    `login_id`        varchar(50) NOT NULL DEFAULT '' COMMENT '登录id，唯一',
-- `role`            tinyint     NOT NULL DEFAULT '1' COMMENT '账户类型：1学生，2教师，3管理员',
    `name`            varchar(50) NOT NULL DEFAULT '' COMMENT '姓名',
    `organization_id` int         NOT NULL DEFAULT '0' COMMENT '组织id，包括学校院系班级',
    `password`        varchar(50) NOT NULL DEFAULT '' COMMENT '密码',
    `phone_number`    varchar(50) NOT NULL DEFAULT '' COMMENT '电话号码',
    `creator_id`      int         NOT NULL DEFAULT '0' COMMENT '创建人id',
    `is_deleted`      tinyint     NOT NULL DEFAULT '1' COMMENT '是否删除 1-否 2-是',
    `created_at`      timestamp   NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at`      timestamp   NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  AUTO_INCREMENT = 20
  DEFAULT CHARSET = utf8mb4 COMMENT ='用户表';

-- ----------------------------
-- Records of user
-- ----------------------------
INSERT INTO `user`
VALUES ('1', '202311111111', '1', '123', '1', 'xxx1', '0', '0', '1', '2023-03-09 16:53:08', '2023-03-09 16:53:08');
INSERT INTO `user`
VALUES ('2', '202311111112', '1', 'lixunguang', '1', 'uua', '0', '0', '1', '2023-03-09 16:54:12',
        '2023-03-09 16:54:12');
INSERT INTO `user`
VALUES ('3', 'a1', '3', 'admin', '1', 'admin', '0', '0', '1', '2023-03-09 16:55:09', '2023-03-09 16:55:09');
INSERT INTO `user`
VALUES ('10', 'a1', '3', 'admin2', '1', 'admin', '0', '0', '1', '2023-03-09 17:07:51', '2023-03-22 17:32:12');
INSERT INTO `user`
VALUES ('12', 't1', '2', 'liteacher', '1', 'adminxx', '0', '0', '1', '2023-03-15 09:16:32', '2023-03-22 18:30:07');
INSERT INTO `user`
VALUES ('13', 'a1', '3', 'admin1', '1', 'adminxx', '0', '0', '1', '2023-03-15 09:17:27', '2023-03-15 09:17:27');
INSERT INTO `user`
VALUES ('14', 't1', '2', 'teacher1', '1', 'teacher1', '0', '0', '1', '2023-03-15 09:26:49', '2023-03-15 09:26:49');
INSERT INTO `user`
VALUES ('15', 'a1', '3', 'admin5', '1', 'admin5', '0', '0', '1', '2023-03-15 09:27:31', '2023-03-15 09:27:31');
INSERT INTO `user`
VALUES ('16', 'a1', '3', 'admin3', '1', 'admin5', '0', '0', '1', '2023-03-15 09:29:17', '2023-03-22 17:32:17');
INSERT INTO `user`
VALUES ('17', 'a1', '3', 'lxg11', '1', 'lxg11', '0', '0', '1', '2023-03-21 15:45:53', '2023-03-21 15:45:53');
INSERT INTO `user`
VALUES ('18', 'a1', '3', 'lxg2', '1', 'lxg2', '0', '0', '1', '2023-03-21 15:46:25', '2023-03-21 15:46:25');
INSERT INTO `user`
VALUES ('19', 't1', '2', 'teacherx', '1', 'teacherx', '0', '0', '1', '2023-03-22 18:33:58', '2023-03-22 18:33:58');

-- ----------------------------
-- Table structure for `work_commit`
-- ----------------------------
DROP TABLE IF EXISTS `work_commit`;
CREATE TABLE `work_commit`
(
    `id`         int       NOT NULL AUTO_INCREMENT COMMENT '主键',
    `lesson_id`  int       NOT NULL DEFAULT '0' COMMENT '所属课堂id',
    `user_id`    int       NOT NULL DEFAULT '0' COMMENT '所属用户id',
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  AUTO_INCREMENT = 9
  DEFAULT CHARSET = utf8mb4 COMMENT ='作业提交表';

-- ----------------------------
-- Records of work_commit
-- ----------------------------
INSERT INTO `work_commit`
VALUES ('8', '1', '11', '2023-04-11 16:20:13', '2023-04-11 16:20:13');

-- ----------------------------
-- Table structure for `work_commit_resource`
-- ----------------------------
DROP TABLE IF EXISTS `work_commit_resource`;
CREATE TABLE `work_commit_resource`
(
    `id`             int       NOT NULL AUTO_INCREMENT COMMENT '主键',
    `work_commit_id` int       NOT NULL DEFAULT '0' COMMENT '作业提交id',
    `resource_id`    int       NOT NULL DEFAULT '0' COMMENT '资源id',
    `type`           tinyint   NOT NULL DEFAULT '0' COMMENT '资源所属类型，如布置作业文件，提交作业文件，提交作业图片，留言',
    `can_modify`     tinyint   NOT NULL DEFAULT '1' COMMENT '是否可以再次修改，1 可以，2 不可以',
    `created_at`     timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at`     timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  AUTO_INCREMENT = 17
  DEFAULT CHARSET = utf8mb4 COMMENT ='作业提交-资源表';

-- ----------------------------
-- Records of work_commit_resource
-- ----------------------------
INSERT INTO `work_commit_resource`
VALUES ('10', '6', '12', '44', '0', '2023-04-11 15:12:40', '2023-04-11 15:12:40');
INSERT INTO `work_commit_resource`
VALUES ('11', '7', '1', '43', '0', '2023-04-11 16:17:41', '2023-04-11 16:17:41');
INSERT INTO `work_commit_resource`
VALUES ('12', '7', '3', '43', '0', '2023-04-11 16:17:44', '2023-04-11 16:17:44');
INSERT INTO `work_commit_resource`
VALUES ('13', '7', '30', '44', '0', '2023-04-11 16:17:58', '2023-04-11 16:17:58');
INSERT INTO `work_commit_resource`
VALUES ('14', '8', '1', '43', '0', '2023-04-11 16:20:14', '2023-04-11 16:20:14');
INSERT INTO `work_commit_resource`
VALUES ('15', '8', '3', '43', '0', '2023-04-11 16:20:14', '2023-04-11 16:20:14');
INSERT INTO `work_commit_resource`
VALUES ('16', '8', '31', '44', '0', '2023-04-11 16:20:15', '2023-04-11 16:20:15');
