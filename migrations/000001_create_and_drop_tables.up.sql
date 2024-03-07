CREATE TABLE "images_table"(
    "image_id" UUID NOT NULL,
    "attraction_id" UUID NOT NULL,
    "image_url" VARCHAR(255) NOT NULL,
    "caption" TEXT NOT NULL,
    "created_at" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL
);
ALTER TABLE
    "images_table" ADD PRIMARY KEY("image_id");
CREATE TABLE "owner_table"(
    "owner_id" UUID NOT NULL,
    "full_name" VARCHAR(255) NOT NULL,
    "email" VARCHAR(255) NOT NULL,
    "password" VARCHAR(255) NOT NULL,
    "birthday" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL,
    "phone_number" VARCHAR(255) NOT NULL,
    "image_url" VARCHAR(255) NOT NULL,
    "refresh_token" VARCHAR(255) NOT NULL,
    "created_at" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "deleted_at" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL,
    "role" VARCHAR(255) NOT NULL
);
ALTER TABLE
    "owner_table" ADD PRIMARY KEY("owner_id");
CREATE TABLE "reviews_table"(
    "review_id" INTEGER NOT NULL,
    "attraction_id" UUID NOT NULL,
    "user_id" UUID NOT NULL,
    "rating" DOUBLE PRECISION NOT NULL,
    "comment" TEXT NOT NULL,
    "created_at" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL,
    "updated_at" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL
);
ALTER TABLE
    "reviews_table" ADD PRIMARY KEY("review_id");
CREATE TABLE "locations_table"(
    "location_id" UUID NOT NULL,
    "name" VARCHAR(255) NOT NULL,
    "latitude" DOUBLE PRECISION NOT NULL,
    "longitude" DOUBLE PRECISION NOT NULL,
    "country" VARCHAR(255) NOT NULL,
    "city" VARCHAR(255) NOT NULL,
    "state/province" VARCHAR(255) NOT NULL,
    "address" VARCHAR(255) NOT NULL
);
ALTER TABLE
    "locations_table" ADD PRIMARY KEY("location_id");
CREATE TABLE "attractions_table"(
    "attraction_id" UUID NOT NULL,
    "name" VARCHAR(255) NOT NULL,
    "description" TEXT NOT NULL,
    "location_id" UUID NOT NULL,
    "opening_hours" VARCHAR(255) NOT NULL,
    "closing_hours" VARCHAR(255) NOT NULL,
    "category" VARCHAR(255) NOT NULL,
    "rating" DOUBLE PRECISION NOT NULL,
    "image_url" VARCHAR(255) NOT NULL,
    "website_url" VARCHAR(255) NOT NULL,
    "contact_information" VARCHAR(255) NOT NULL,
    "created_at" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "owner_id" UUID NOT NULL
);
ALTER TABLE
    "attractions_table" ADD PRIMARY KEY("attraction_id");
ALTER TABLE
    "attractions_table" ADD CONSTRAINT "attractions_table_location_id_unique" UNIQUE("location_id");
ALTER TABLE
    "attractions_table" ADD CONSTRAINT "attractions_table_owner_id_foreign" FOREIGN KEY("owner_id") REFERENCES "owner_table"("owner_id");
ALTER TABLE
    "attractions_table" ADD CONSTRAINT "attractions_table_location_id_foreign" FOREIGN KEY("location_id") REFERENCES "locations_table"("location_id");
ALTER TABLE
    "reviews_table" ADD CONSTRAINT "reviews_table_attraction_id_foreign" FOREIGN KEY("attraction_id") REFERENCES "attractions_table"("attraction_id");
ALTER TABLE
    "images_table" ADD CONSTRAINT "images_table_attraction_id_foreign" FOREIGN KEY("attraction_id") REFERENCES "attractions_table"("attraction_id");