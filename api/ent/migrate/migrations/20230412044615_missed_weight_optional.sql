-- Modify "fighter_results" table
ALTER TABLE "fighter_results" ALTER COLUMN "missed_weight" DROP NOT NULL, ALTER COLUMN "missed_weight" SET DEFAULT false;
