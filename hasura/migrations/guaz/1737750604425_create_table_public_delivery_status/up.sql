CREATE TABLE "public"."delivery_status" ("status_name" varchar NOT NULL, "status_description" text NOT NULL, PRIMARY KEY ("status_name") , UNIQUE ("status_name"));COMMENT ON TABLE "public"."delivery_status" IS E'status for delivery request from sender to couriers';
