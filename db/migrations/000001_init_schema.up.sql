CREATE TABLE "users" (
  "id" bigserial PRIMARY KEY,
  "email" varchar NOT NULL,
  "password" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "bills" (
  "id" bigserial PRIMARY KEY,
  "user_id" bigint NOT NULL,
  "concept" varchar NOT NULL,
  "price" bigint NOT NULL,
  "fecha" timestamptz NOT NULL DEFAULT (now())
);

CREATE INDEX ON "users" ("email");

CREATE INDEX ON "bills" ("concept");

COMMENT ON COLUMN "bills"."price" IS 'price must be positive';

ALTER TABLE "bills" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");
