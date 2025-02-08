alter table "public"."receiver" add column "updated_at" timestamptz
 not null default now();
