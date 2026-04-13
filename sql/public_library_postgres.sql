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

 Date: 13/04/2026 15:54:45
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
INSERT INTO "public"."authors" ("id", "name", "created", "updated") VALUES ('019c08d1-c060-7537-9dcf-8d683662f22d', 'Keigo Higashino', '2026-01-29 15:14:49.952378', NULL);
INSERT INTO "public"."authors" ("id", "name", "created", "updated") VALUES ('019d85c8-294e-727c-b36f-7bdd35908fd2', 'Fyodor Dostoyevsky', '2026-04-13 14:39:40.750748', NULL);
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
-- Function structure for _navicat_temp_stored_proc
-- ----------------------------
DROP FUNCTION IF EXISTS "public"."_navicat_temp_stored_proc"("p_sort_field" text, "p_sort_order" text, "p_limit" int4, "p_offset" int4, "p_search" text);
CREATE OR REPLACE FUNCTION "public"."_navicat_temp_stored_proc"("p_sort_field" text, "p_sort_order" text, "p_limit" int4, "p_offset" int4, "p_search" text)
  RETURNS SETOF "public"."authors" AS $BODY$
DECLARE
    v_query TEXT;
    v_sort_field TEXT;
    v_sort_order TEXT;
BEGIN
    -- whitelist field
    IF p_sort_field NOT IN ('id', 'name', 'created_at') THEN
        v_sort_field := 'id';
    ELSE
        v_sort_field := p_sort_field;
    END IF;

    -- whitelist order
    IF LOWER(p_sort_order) NOT IN ('asc', 'desc') THEN
        v_sort_order := 'asc';
    ELSE
        v_sort_order := p_sort_order;
    END IF;

    v_query := format(
        'SELECT 
					id, 
					name
         FROM public.authors
         WHERE ($1 IS NULL OR $1 = '''' OR name ILIKE ''%%'' || $1 || ''%%'')
         ORDER BY %I %s
         LIMIT $2 OFFSET $3',
        v_sort_field,
        v_sort_order
    );

    RETURN QUERY EXECUTE v_query
    USING p_search, p_limit, p_offset;
END;
$BODY$
  LANGUAGE plpgsql VOLATILE
  COST 100
  ROWS 1000;
ALTER FUNCTION "public"."_navicat_temp_stored_proc"("p_sort_field" text, "p_sort_order" text, "p_limit" int4, "p_offset" int4, "p_search" text) OWNER TO "sketch";

-- ----------------------------
-- Function structure for get_count_author
-- ----------------------------
DROP FUNCTION IF EXISTS "public"."get_count_author"("p_search" text);
CREATE OR REPLACE FUNCTION "public"."get_count_author"("p_search" text)
  RETURNS "pg_catalog"."int8" AS $BODY$
DECLARE
    v_count BIGINT;
BEGIN
    SELECT COUNT(*)
    INTO v_count
    FROM public.authors
    WHERE (p_search IS NULL OR p_search = '' OR name ILIKE '%' || p_search || '%');

    RETURN v_count;
END;
$BODY$
  LANGUAGE plpgsql VOLATILE
  COST 100;
ALTER FUNCTION "public"."get_count_author"("p_search" text) OWNER TO "sketch";

-- ----------------------------
-- Function structure for get_list_author
-- ----------------------------
DROP FUNCTION IF EXISTS "public"."get_list_author"("p_sort_field" text, "p_sort_order" text, "p_limit" int4, "p_offset" int4, "p_search" text);
CREATE OR REPLACE FUNCTION "public"."get_list_author"("p_sort_field" text, "p_sort_order" text, "p_limit" int4, "p_offset" int4, "p_search" text)
  RETURNS TABLE("id" uuid, "name" varchar, "created" timestamp, "updated" timestamp) AS $BODY$
DECLARE
    v_query TEXT;
    v_sort_field TEXT;
    v_sort_order TEXT;
BEGIN
    -- whitelist field
    IF p_sort_field NOT IN ('id', 'name', 'created_at') THEN
        v_sort_field := 'id';
    ELSE
        v_sort_field := p_sort_field;
    END IF;

    -- whitelist order
    IF LOWER(p_sort_order) NOT IN ('asc', 'desc') THEN
        v_sort_order := 'asc';
    ELSE
        v_sort_order := p_sort_order;
    END IF;

    v_query := format(
        'SELECT 
            id,
            name,
            created,
						updated
         FROM public.authors
         WHERE ($1 IS NULL OR $1 = '''' OR name ILIKE ''%%'' || $1 || ''%%'')
         ORDER BY %I %s
         LIMIT $2 OFFSET $3',
        v_sort_field,
        v_sort_order
    );

    RETURN QUERY EXECUTE v_query
    USING p_search, p_limit, p_offset;
END;
$BODY$
  LANGUAGE plpgsql VOLATILE
  COST 100
  ROWS 1000;
ALTER FUNCTION "public"."get_list_author"("p_sort_field" text, "p_sort_order" text, "p_limit" int4, "p_offset" int4, "p_search" text) OWNER TO "sketch";

-- ----------------------------
-- Function structure for update_updated_column
-- ----------------------------
DROP FUNCTION IF EXISTS "public"."update_updated_column"();
CREATE OR REPLACE FUNCTION "public"."update_updated_column"()
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
