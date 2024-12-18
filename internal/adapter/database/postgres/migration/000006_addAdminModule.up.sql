BEGIN;
    CREATE TABLE "admin_users" (
        "id" SERIAL PRIMARY KEY,
        "password_hash" VARCHAR NOT NULL,
        "email" VARCHAR UNIQUE NOT NULL,
        "role" VARCHAR NOT NULL DEFAULT 'admin',
        "created_at" TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
        "updated_at" TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
        "deleted_at" TIMESTAMPTZ
    );

    CREATE UNIQUE INDEX "admin_email" ON "admin_users" ("email");

COMMIT;
