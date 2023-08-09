CREATE TABLE "keys" (
  "id" SERIAL NOT NULL,
  "key" varchar PRIMARY KEY,
  "usage_count" INT NOT NULL DEFAULT 0,
  "created_at" timestamp NOT NULL DEFAULT (now())
);