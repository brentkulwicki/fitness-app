CREATE TYPE "excercise_review" AS ENUM (
  'review',
  'approved',
  'duplicate',
  'declined',
  'archived'
);

CREATE TYPE "weights" AS ENUM (
  'lb',
  'kg'
);

CREATE TABLE "exercises" (
  "id" integer PRIMARY KEY,
  "name" varchar NOT NULL,
  "details" varchar NOT NULL,
  "status" excercise_review DEFAULT 'review',
  "reference" varchar,
  "created_at" timestamp DEFAULT (now()),
  "updated_at" timestamp
);

CREATE TABLE "user_exercises" (
  "id" integer PRIMARY KEY,
  "exercise_id" integer NOT NULL,
  "reps_expected" integer NOT NULL DEFAULT 10,
  "reps_completed" integer NOT NULL DEFAULT 0,
  "weight" integer DEFAULT 0,
  "weight_type" weights DEFAULT 'lb',
  "notes" varchar,
  "created_at" timestamp DEFAULT (now()),
  "updated_at" timestamp,
  "finished_at" timestamp
);

CREATE TABLE "users" (
  "id" integer PRIMARY KEY,
  "name" varchar NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "created_at" timestamp DEFAULT (now())
);

CREATE TABLE "workouts" (
  "id" integer PRIMARY KEY,
  "user_id" integer,
  "exercises" array[integer],
  "created_at" timestamp DEFAULT (now()),
  "started_at" timestamp,
  "updated_at" timestamp,
  "finished_at" timestamp
);

CREATE INDEX ON "user_exercises" ("exercise_id");

CREATE INDEX ON "workouts" ("user_id");

CREATE INDEX ON "workouts" ("exercises");

ALTER TABLE "workouts" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "workouts" ADD FOREIGN KEY ("exercises") REFERENCES "user_exercises" ("id");

ALTER TABLE "user_exercises" ADD FOREIGN KEY ("exercise_id") REFERENCES "exercises" ("id");
