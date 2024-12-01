BEGIN;
    DROP TABLE wallet_histories;
    DROP TABLE instruments;
    DROP TABLE wallets;

    DROP TYPE status;
    DROP TYPE instrument_type;
    DROP TYPE currency_type;

COMMIT;