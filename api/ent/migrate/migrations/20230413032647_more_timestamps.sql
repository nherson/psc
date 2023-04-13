-- Modify "events" table
ALTER TABLE "events" ALTER COLUMN "created_at", ALTER COLUMN "updated_at";
-- Modify "fighter_alias" table
ALTER TABLE "fighter_alias" ADD COLUMN "created_at" timestamptz NOT NULL, ADD COLUMN "updated_at" timestamptz NOT NULL;
-- Modify "fighter_results" table
ALTER TABLE "fighter_results" ADD COLUMN "created_at" timestamptz NOT NULL, ADD COLUMN "updated_at" timestamptz NOT NULL;
-- Modify "fighters" table
ALTER TABLE "fighters" ADD COLUMN "created_at" timestamptz NOT NULL, ADD COLUMN "updated_at" timestamptz NOT NULL;
-- Modify "fights" table
ALTER TABLE "fights" ADD COLUMN "created_at" timestamptz NOT NULL, ADD COLUMN "updated_at" timestamptz NOT NULL;
-- Modify "upcoming_events" table
ALTER TABLE "upcoming_events" ADD COLUMN "created_at" timestamptz NOT NULL, ADD COLUMN "updated_at" timestamptz NOT NULL;
-- Modify "upcoming_fights" table
ALTER TABLE "upcoming_fights" ADD COLUMN "created_at" timestamptz NOT NULL, ADD COLUMN "updated_at" timestamptz NOT NULL;
