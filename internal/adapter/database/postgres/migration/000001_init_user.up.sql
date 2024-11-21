CREATE TYPE "users_type_enum" AS ENUM ('free', 'premium');

CREATE TABLE "users" (
    "id" BIGSERIAL PRIMARY KEY,
    "email" varchar UNIQUE NOT NULL,
    "user_type" users_type_enum DEFAULT 'free',
    "display_name" varchar,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "deleted_at" TIMESTAMP
);

CREATE UNIQUE INDEX "user_email" ON "users" ("email");
