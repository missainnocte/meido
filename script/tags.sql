/*
 Navicat Premium Data Transfer

 Source Server         : meido_dev
 Source Server Type    : SQLite
 Source Server Version : 3030001
 Source Schema         : main

 Target Server Type    : SQLite
 Target Server Version : 3030001
 File Encoding         : 65001

 Date: 06/09/2020 23:15:56
*/

PRAGMA foreign_keys = false;

-- ----------------------------
-- Table structure for tags
-- ----------------------------
DROP TABLE IF EXISTS "tags";
CREATE TABLE "tags" (
  "id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,
  "title" text NOT NULL
);

-- ----------------------------
-- Indexes structure for table tags
-- ----------------------------
CREATE INDEX "index_id"
ON "tags" (
  "id" ASC
);
CREATE INDEX "index_title"
ON "tags" (
  "title" ASC
);

PRAGMA foreign_keys = true;
