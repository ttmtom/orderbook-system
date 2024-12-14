BEGIN;

    CREATE TYPE "transaction_type" AS ENUM ('deposit', 'withdrawal' , 'transfer');

    CREATE TABLE "transactions" (
        "id" SERIAL PRIMARY KEY,
        "id_hash" VARCHAR UNIQUE NOT NULL,
        "amount" NUMERIC(10, 8) NOT NULL,
        "type" transaction_type NOT NULL,
        "from_id" INT REFERENCES wallets(id),
        "to_id" INT REFERENCES wallets(id),
        "description" TEXT,
        "created_at" TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
    );

    CREATE TYPE "transaction_event_type" AS ENUM ('pending', 'success', 'rejected', 'cancel');

    CREATE TABLE "transaction_events" (
        "id" SERIAL PRIMARY KEY,
        "transaction_id" INT REFERENCES "transactions" ("id"),
        "type" transaction_event_type NOT NULL,
        "created_at" TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
    );

    ALTER TABLE wallet_histories ADD transaction_id INT REFERENCES transactions(id);

COMMIT;