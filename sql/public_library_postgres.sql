/*
 Navicat Premium Dump SQL

 Source Server         : Postgres Brew
 Source Server Type    : PostgreSQL
 Source Server Version : 170006 (170006)
 Source Host           : localhost:5999
 Source Catalog        : library
 Source Schema         : public

 Target Server Type    : PostgreSQL
 Target Server Version : 170006 (170006)
 File Encoding         : 65001

 Date: 03/11/2025 15:49:07
*/


-- ----------------------------
-- Table structure for authors
-- ----------------------------
DROP TABLE IF EXISTS "public"."authors";
CREATE TABLE "public"."authors" (
  "id" uuid NOT NULL,
  "name" varchar(100) COLLATE "pg_catalog"."default" NOT NULL,
  "created" timestamp(6) NOT NULL DEFAULT now(),
  "updated" timestamp(6)
)
;
ALTER TABLE "public"."authors" OWNER TO "sketch";

-- ----------------------------
-- Records of authors
-- ----------------------------
BEGIN;
INSERT INTO "public"."authors" ("id", "name", "created", "updated") VALUES ('019a48c6-e7a0-7eb4-89c3-d64000745659', 'J.K. Rowling', '2025-11-03 15:13:06.336672', NULL);
COMMIT;

-- ----------------------------
-- Table structure for books
-- ----------------------------
DROP TABLE IF EXISTS "public"."books";
CREATE TABLE "public"."books" (
  "id" uuid NOT NULL,
  "author_id" uuid NOT NULL,
  "title" varchar(100) COLLATE "pg_catalog"."default" NOT NULL,
  "amount" int8 NOT NULL DEFAULT 0,
  "created" timestamp(6) DEFAULT now(),
  "updated" timestamp(6)
)
;
ALTER TABLE "public"."books" OWNER TO "sketch";

-- ----------------------------
-- Records of books
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Function structure for update_updated_column
-- ----------------------------
DROP FUNCTION IF EXISTS "public"."update_updated_column"();
CREATE FUNCTION "public"."update_updated_column"()
  RETURNS "pg_catalog"."trigger" AS $BODY$
BEGIN
  NEW.updated = NOW();
  RETURN NEW;
END;
$BODY$
  LANGUAGE plpgsql VOLATILE
  COST 100;
ALTER FUNCTION "public"."update_updated_column"() OWNER TO "sketch";

-- ----------------------------
-- Indexes structure for table authors
-- ----------------------------
CREATE UNIQUE INDEX "authors_name_idx" ON "public"."authors" USING btree (
  "name" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
);

-- ----------------------------
-- Triggers structure for table authors
-- ----------------------------
CREATE TRIGGER "set_updated_timestamp" BEFORE UPDATE ON "public"."authors"
FOR EACH ROW
EXECUTE PROCEDURE "public"."update_updated_column"();

-- ----------------------------
-- Primary Key structure for table authors
-- ----------------------------
ALTER TABLE "public"."authors" ADD CONSTRAINT "authors_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Triggers structure for table books
-- ----------------------------
CREATE TRIGGER "set_updated_timestamp" BEFORE UPDATE ON "public"."books"
FOR EACH ROW
EXECUTE PROCEDURE "public"."update_updated_column"();

-- ----------------------------
-- Primary Key structure for table books
-- ----------------------------
ALTER TABLE "public"."books" ADD CONSTRAINT "books_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Foreign Keys structure for table books
-- ----------------------------
ALTER TABLE "public"."books" ADD CONSTRAINT "books_author_id_fkey" FOREIGN KEY ("author_id") REFERENCES "public"."authors" ("id") ON DELETE NO ACTION ON UPDATE NO ACTION;
