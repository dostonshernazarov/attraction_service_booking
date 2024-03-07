-- Drop foreign key constraints
ALTER TABLE "images_table" DROP CONSTRAINT IF EXISTS "images_table_attraction_id_foreign";
ALTER TABLE "reviews_table" DROP CONSTRAINT IF EXISTS "reviews_table_attraction_id_foreign";
ALTER TABLE "attractions_table" DROP CONSTRAINT IF EXISTS "attractions_table_owner_id_foreign";
ALTER TABLE "attractions_table" DROP CONSTRAINT IF EXISTS "attractions_table_location_id_foreign";
ALTER TABLE "attractions_table" DROP CONSTRAINT IF EXISTS "attractions_table_location_id_unique";

-- Drop primary keys
ALTER TABLE "attractions_table" DROP CONSTRAINT IF EXISTS "attractions_table_pkey";
ALTER TABLE "locations_table" DROP CONSTRAINT IF EXISTS "locations_table_pkey";
ALTER TABLE "reviews_table" DROP CONSTRAINT IF EXISTS "reviews_table_pkey";
ALTER TABLE "images_table" DROP CONSTRAINT IF EXISTS "images_table_pkey";

-- Drop tables
DROP TABLE IF EXISTS "attractions_table";
DROP TABLE IF EXISTS "locations_table";
DROP TABLE IF EXISTS "reviews_table";
DROP TABLE IF EXISTS "owner_table";
DROP TABLE IF EXISTS "images_table";