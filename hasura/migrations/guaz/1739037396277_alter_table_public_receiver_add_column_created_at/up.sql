alter table "public"."receiver" add column "created_at" timestamptz
 not null default now();
