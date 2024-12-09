BEGIN;
    CREATE TYPE "currency_type" AS ENUM ('BTC', 'ETH');

    CREATE TABLE "wallets" (
        "id" SERIAL PRIMARY KEY,
        "user_id" SERIAL REFERENCES "users" ("id"),
        "currency" currency_type NOT NULL,
        "balance" NUMERIC(10, 8) NOT NULL DEFAULT 0.0,
        "locked" BOOLEAN NOT NULL DEFAULT false,
        "hold" NUMERIC(10, 8) NOT NULL DEFAULT 0.0,

        "created_at" TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
        "updated_at" TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,

        UNIQUE ("user_id", "currency")
    );

    CREATE TYPE "instrument_type" AS ENUM ('deposit', 'withdrawal');

    CREATE TABLE "instruments" (
        "id" SERIAL PRIMARY KEY,
        "wallet_id" SERIAL REFERENCES  "wallets" ("id"),
        "amount" NUMERIC(10, 8) NOT NULL,
        "type" instrument_type NOT NULL,
        "created_at" TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
    );

    CREATE TYPE "status" AS ENUM ('pending', 'completed', 'failed', 'canceled');

    CREATE TABLE "wallet_histories" (
        "id" SERIAL PRIMARY KEY,
        "wallet_id" SERIAL REFERENCES "wallets" ("id"),
        "instrument_id" SERIAL REFERENCES "instruments" ("id"),
        "status" status NOT NULL,
        "created_at" TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
    );

COMMIT;
