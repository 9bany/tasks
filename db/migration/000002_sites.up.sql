CREATE TABLE "sites" (
  "id" SERIAL NOT NULL,
  "url" varchar PRIMARY KEY NOT NULL,
  "meta_data" JSONB NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (now())
);