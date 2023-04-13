-- Modify "events" table
ALTER TABLE "events" ADD COLUMN "created_at" timestamptz NOT NULL, ADD COLUMN "updated_at" timestamptz NOT NULL;
