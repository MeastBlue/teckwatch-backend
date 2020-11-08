CREATE extension IF NOT EXISTS "uuid-ossp";

DROP TYPE IF EXISTS "prio_t";
CREATE TYPE "prio_t" AS ENUM (
  'NONE',
  'HIGH',
  'MEDIUM',
  'LOW'
);

CREATE TABLE "users" (
  "id" uuid PRIMARY KEY NOT NULL DEFAULT (uuid_generate_v4()),
  "lastname" text NOT NULL,
  "firstname" text NOT NULL,
  "surname" text NOT NULL,
  "avatar" text NOT NULL,
  "email" text NOT NULL,
  "password" text NOT NULL,
  "created_at" date NOT NULL DEFAULT (now()),
  "updated_at" date NOT NULL DEFAULT (now())
);

CREATE TABLE "techs" (
  "id" uuid PRIMARY KEY NOT NULL DEFAULT (uuid_generate_v4()),
  "label" text NOT NULL,
  "description" text NOT NULL,
  "created_at" date NOT NULL DEFAULT (now()),
  "updated_at" date NOT NULL DEFAULT (now())
);

CREATE TABLE "task" (
  "id" uuid PRIMARY KEY NOT NULL DEFAULT (uuid_generate_v4()),
  "label" text NOT NULL,
  "prio" prio_t NOT NULL DEFAULT 'NONE',
  "progress" smallint NOT NULL DEFAULT 0,
  "done" boolean NOT NULL DEFAULT false,
  "user_id" uuid NOT NULL,
  "tech_id" uuid NOT NULL,
  "created_at" date NOT NULL DEFAULT (now()),
  "updated_at" date NOT NULL DEFAULT (now())
);

ALTER TABLE "task" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "task" ADD FOREIGN KEY ("tech_id") REFERENCES "techs" ("id");
