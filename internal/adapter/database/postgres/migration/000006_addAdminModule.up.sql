BEGIN;
CREATE TABLE "admin_users" (
    "id" SERIAL PRIMARY KEY,
    "password_hash" VARCHAR NOT NULL,
    "username" VARCHAR UNIQUE NOT NULL,
    "created_at" TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    "deleted_at" TIMESTAMPTZ
);

CREATE UNIQUE INDEX "admin_username" ON "users" ("username");

COMMIT;
