BEGIN;
    CREATE TABLE "users" (
        "id" SERIAL PRIMARY KEY,
        "id_hash" VARCHAR UNIQUE NOT NULL,
        "password_hash" VARCHAR NOT NULL,
        "email" VARCHAR UNIQUE NOT NULL,
        "display_name" VARCHAR(255),
        "created_at" TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
        "updated_at" TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
        "deleted_at" TIMESTAMPTZ
    );

    CREATE UNIQUE INDEX "user_email" ON "users" ("email");
    CREATE UNIQUE INDEX "user_hash_id" ON "users" ("id_hash");

COMMIT;
