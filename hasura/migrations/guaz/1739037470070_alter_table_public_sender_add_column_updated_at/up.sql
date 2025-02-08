alter table "public"."sender" add column "updated_at" timestamptz
 not null default now();
