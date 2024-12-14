BEGIN;
    CREATE TYPE "currency_type" AS ENUM ('BTC', 'ETH');

    CREATE TABLE "wallets" (
        "id" SERIAL PRIMARY KEY,
        "user_id" INT REFERENCES "users" ("id"),
        "currency" currency_type NOT NULL,
        "balance" NUMERIC(10, 8) NOT NULL DEFAULT 0.0,
        "locked" NUMERIC(10, 8) NOT NULL DEFAULT 0.0,
        "blocked" BOOLEAN DEFAULT false,

        "created_at" TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
        "updated_at" TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,

        UNIQUE ("user_id", "currency")
    );

    CREATE TYPE "wallet_action_type" AS ENUM ('increase', 'decrease', 'lock', 'release', 'commit');

    CREATE TABLE "wallet_histories" (
        "id" SERIAL PRIMARY KEY,
        "wallet_id" INT REFERENCES "wallets" ("id"),
        "amount" NUMERIC(10, 8) NOT NULL DEFAULT 0.0,
        "action" wallet_action_type NOT NULL,
        "created_at" TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
    );
COMMIT;
