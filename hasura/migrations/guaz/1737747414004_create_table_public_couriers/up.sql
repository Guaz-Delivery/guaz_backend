CREATE TABLE "public"."couriers" ("id" uuid NOT NULL DEFAULT gen_random_uuid(), "first_name" varchar NOT NULL, "middle_name" varchar NOT NULL, "last_name" varchar NOT NULL, "location" point NOT NULL, "rate" integer NOT NULL, "is_verified" boolean NOT NULL DEFAULT false, "phone_number" text NOT NULL, "email" text NOT NULL, "created_at" timestamptz NOT NULL DEFAULT now(), "updated_at" timestamptz NOT NULL DEFAULT now(), PRIMARY KEY ("id") , UNIQUE ("id"), UNIQUE ("email"), UNIQUE ("phone_number"));COMMENT ON TABLE "public"."couriers" IS E'a delivery guys table';
CREATE EXTENSION IF NOT EXISTS pgcrypto;
