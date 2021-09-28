-- ---
-- Globals
-- ---

-- SET SQL_MODE="NO_AUTO_VALUE_ON_ZERO";
-- SET FOREIGN_KEY_CHECKS=0;
DROP TABLE IF EXISTS "cards" CASCADE;
DROP TABLE IF EXISTS "sets" CASCADE;
DROP TABLE IF EXISTS "card_sets" CASCADE;

-- ---
-- Table 'cards'
--
-- ---
DROP TABLE IF EXISTS "cards";

CREATE TABLE "cards" (
  "card_id" SERIAL NOT NULL,
  "card_name" VARCHAR(100) NOT NULL,
  "card_type" VARCHAR(100) NOT NULL,
  "card_artist" VARCHAR(100) NOT NULL,
  "card_effect" VARCHAR(200),
  "card_power" INTEGER,
  "card_intelligence" INTEGER,
  "card_endurance" INTEGER,

  PRIMARY KEY ("card_id")
)


-- ---
-- Table 'sets'
--
-- ---
DROP TABLE IF EXISTS "sets";

CREATE TABLE "sets" (
  "set_id" SERIAL NOT NULL,
  "set_name" VARCHAR(100) NOT NULL,

  PRIMARY KEY ("set_id")
)

-- ---
-- Table 'card_sets'
--
-- ---
DROP TABLE IF EXISTS "card_sets";

CREATE TABLE "card_sets" (
  "card_set_id" SERIAL NOT NULL,
  "card_id" INTEGER NOT NULL DEFAULT -1,
  "set_id" INTEGER NOT NULL DEFAULT -1,

  "card_set_number" INTEGER NOT NULL,
  "card_set_rarity" VARCHAR(100) NOT NULL,

  PRIMARY KEY ("card_set_id")
)


ALTER TABLE "card_sets" ADD FOREIGN KEY ("card_id") REFERENCES "cards" ("card_id") ON DELETE CASCADE;
ALTER TABLE "card_sets" ADD FOREIGN KEY ("set_id") REFERENCES "sets" ("set_id") ON DELETE CASCADE;
