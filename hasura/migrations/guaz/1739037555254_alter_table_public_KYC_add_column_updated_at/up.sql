alter table "public"."KYC" add column "updated_at" timestamptz
 not null default now();
