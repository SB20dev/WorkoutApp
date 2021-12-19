
-- +migrate Up
CREATE TABLE IF NOT EXISTS "users" (
    "system_id" serial primary key,
    "id" varchar(32) unique not null,
    "password" varchar(100) not null,
    "last_login" timestamp with time zone,
    "created" timestamp with time zone
);
SELECT SETVAL("users_system_id_seq", 1000, false);

CREATE TABLE IF NOT EXISTS "commitments" (
    "id" serial primary key,
    "user_id" integer references "users" ("system_id") on delete cascade,
    "score" integer,
    "committed" timestamp with time zone
);
SELECT SETVAL("commitments_id_seq", 1000, false);

CREATE TABLE IF NOT EXISTS "menus" (
    "id" serial primary key,
    "user_id" integer references "users" ("system_id") on delete cascade,
    "name" varchar(32)
);
SELECT SETVAL("menus_id_seq", 1000, false);

CREATE TABLE IF NOT EXISTS "commitment_menus" (
    "id" serial primary key,
    "commitment_id" integer references "commitments" ("id") on delete cascade,
    "menu_id" integer references "menus" ("id") on delete cascade,
    "amount" varchar(32) --定量的にできないか
);
SELECT SETVAL("commitment_menus_id_seq", 1000, false);

CREATE TABLE IF NOT EXISTS "common_classes" (
    "id" serial primary key,
    "class" varchar(32)
);
SELECT SETVAL("common_classes_id_seq", 1000, false);

CREATE TABLE IF NOT EXISTS "classes" (
    "id" serial primary key,
    "class" varchar(32),
    "common_class_id" integer,
    "user_id" integer references "users" ("system_id") on delete cascade,
    "deleted" boolean default false
);
SELECT SETVAL("classes_id_seq", 1000, false);

CREATE TABLE IF NOT EXISTS "common_parts" (
    "id" serial primary key,
    "class_id" integer references "common_classes" ("id"),
    "part" varchar(32) not null,
);
SELECT SETVAL("common_parts_id_seq", 1000, false);

CREATE TABLE IF NOT EXISTS "parts" (
    "id" serial primary key,
    "class_id" integer references "class" ("id") on delete cascade,
    "part" varchar(32),
    "state_id" integer references "status" ("id") not null,
    "common_part_id" integer,
    "user_id" integer references "users" ("system_id") on delete cascade,
    "deleted" boolean default false
);
SELECT SETVAL("parts_id_seq", 1000, false);

CREATE TABLE IF NOT EXISTS "menu_parts" (
    "id" serial primary key,
    "menu_id" integer references "menus" ("id") on delete cascade,
    "part_id" integer references "parts" ("id") on delete cascade
);
SELECT SETVAL("menu_parts_id_seq", 1000, false);

CREATE TABLE IF NOT EXISTS "status" (
    "id" integer primary key,
    "state" varchar(32),
);
SELECT SETVAL("status", 1000, false);

-- +migrate Down
DROP TABLE IF EXISTS
    "users",
    "commitments",
    "menus",
    "commitment_menus",
    "common_classes",
    "classes",
    "common_parts",
    "parts",
    "menu_parts",
    "status"
    CASCADE;