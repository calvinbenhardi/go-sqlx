CREATE TABLE "organizations" (
  "id" uuid PRIMARY KEY NOT NULL,
  "name" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "organization_branches" (
  "id" uuid PRIMARY KEY NOT NULL,
  "name" varchar NOT NULL,
  "organization_id" uuid NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "users" (
  "id" uuid PRIMARY KEY NOT NULL,
  "first_name" varchar NOT NULL,
  "email" varchar NOT NULL,
  "password" varchar NOT NULL,
  "organization_branch_id" uuid NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "wallets" (
  "id" varchar PRIMARY KEY NOT NULL,
  "balance" numeric(20,2) NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "ledgers" (
  "id" varchar PRIMARY KEY NOT NULL,
  "sender_id" uuid NOT NULL,
  "receiver_id" uuid NOT NULL,
  "amount" numeric(20,2) NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

ALTER TABLE "organization_branches" ADD FOREIGN KEY ("organization_id") REFERENCES "organizations" ("id");

ALTER TABLE "users" ADD FOREIGN KEY ("organization_branch_id") REFERENCES "organization_branches" ("id");

ALTER TABLE "ledgers" ADD FOREIGN KEY ("sender_id") REFERENCES "users" ("id");

ALTER TABLE "ledgers" ADD FOREIGN KEY ("receiver_id") REFERENCES "users" ("id");
