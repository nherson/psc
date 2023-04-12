-- Drop index "fighterresults_fight_id_fighter_id" from table: "fighter_results"
DROP INDEX "fighterresults_fight_id_fighter_id";
-- Create index "fighterresults_fighter_id_fight_id" to table: "fighter_results"
CREATE UNIQUE INDEX "fighterresults_fighter_id_fight_id" ON "fighter_results" ("fighter_id", "fight_id");
