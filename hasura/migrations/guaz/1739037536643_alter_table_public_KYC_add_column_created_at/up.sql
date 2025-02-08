alter table "public"."KYC" add column "created_at" timestamptz
 not null default now();
