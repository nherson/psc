-- Modify "fighter_results" table
ALTER TABLE "fighter_results" ADD COLUMN "corner" character varying NOT NULL DEFAULT 'red', ADD COLUMN "win" boolean NOT NULL DEFAULT false;
