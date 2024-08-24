CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
-- Create "users" table
CREATE TABLE "users" (
  "id" uuid NOT NULL DEFAULT uuid_generate_v4(),
  "created_at" timestamptz NOT NULL,
  "updated_at" timestamptz NOT NULL,
  "deleted_at" timestamptz NULL,
  "first_name" text NOT NULL,
  "last_name" text NOT NULL,
  "phone_number" text NOT NULL,
  "address" text NOT NULL,
  "pin" text NOT NULL,
  "balance" bigint NOT NULL DEFAULT 0,
  PRIMARY KEY ("id"),
  CONSTRAINT "uni_users_phone_number" UNIQUE ("phone_number")
);
-- Create index "idx_users_deleted_at" to table: "users"
CREATE INDEX "idx_users_deleted_at" ON "users" ("deleted_at");
-- Create "transactions" table
CREATE TABLE "transactions" (
  "id" uuid NOT NULL DEFAULT uuid_generate_v4(),
  "status" text NOT NULL,
  "user_id" uuid NOT NULL,
  "transaction_type" text NOT NULL,
  "amount" bigint NOT NULL,
  "remarks" text NOT NULL,
  "balance_before" bigint NOT NULL,
  "balance_after" bigint NOT NULL,
  "action" text NULL,
  "created_at" timestamptz NOT NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_transactions_user" FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
