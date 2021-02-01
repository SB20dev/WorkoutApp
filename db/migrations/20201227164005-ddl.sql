
-- +migrate Up
CREATE TABLE IF NOT EXISTS "users" (
    "system_id" serial primary key,
    "id" varchar(32) unique not null,
    "password" varchar(100) not null,
    "last_login" timestamp with time zone,
    "created" timestamp with time zone
);

CREATE TABLE IF NOT EXISTS "commitments" (
    "id" serial primary key,
    "user_id" varchar(32) references "users" ("id"),
    "score" integer,
    "committed" timestamp with time zone
);

CREATE TABLE IF NOT EXISTS "menus" (
    "id" serial primary key,
    "user_id" varchar(32) references "users" ("id"),
    "name" varchar(32)
);

CREATE TABLE IF NOT EXISTS "commitment_menus" (
    "id" serial primary key,
    "commitment_id" integer references "commitments" ("id"),
    "menu_id" integer references "menus" ("id"),
    "amount" varchar(32) --定量的にできないか
);

CREATE TABLE IF NOT EXISTS "parts" (
    "id" serial primary key,
    "class" varchar(32),
    "detail" varchar(32)
);

CREATE TABLE IF NOT EXISTS "menu_parts" (
    "id" serial primary key,
    "menu_id" integer references "menus" ("id"),
    "part_id" integer references "parts" ("id")
);

CREATE TABLE IF NOT EXISTS "statuses" (
    "id" serial primary key,
    "user_id" varchar(32) references "users" ("id"),
    "part_id" integer references "parts" ("id"),
    "last_commited" timestamp with time zone
);

-- +migrate Down
DROP TABLE IF EXISTS
    "users",
    "commitments",
    "menus",
    "commitment_menus",
    "parts",
    "menu_parts",
    "statuses"
    CASCADE;