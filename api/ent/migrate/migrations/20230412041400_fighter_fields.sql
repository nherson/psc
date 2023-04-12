-- Modify "fighters" table
ALTER TABLE "fighters" ADD COLUMN "ufc_fighter_id" character varying NOT NULL, ADD COLUMN "mma_id" bigint NOT NULL, ADD COLUMN "first_name" character varying NOT NULL, ADD COLUMN "last_name" character varying NOT NULL, ADD COLUMN "nick_name" character varying NOT NULL;
-- Create index "fighters_ufc_fighter_id_key" to table: "fighters"
CREATE UNIQUE INDEX "fighters_ufc_fighter_id_key" ON "fighters" ("ufc_fighter_id");
