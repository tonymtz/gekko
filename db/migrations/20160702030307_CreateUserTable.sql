
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied

CREATE TABLE "user" (
  "id" SERIAL PRIMARY KEY NOT NULL,
  "id_provider" TEXT NOT NULL,
  "display_name" TEXT NOT NULL,
  "email" TEXT NOT NULL,
  "profile_picture" TEXT NOT NULL,
  "role" INT NOT NULL DEFAULT 1,
  "token" TEXT,
  "jwt" TEXT
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back

DROP TABLE "user";
