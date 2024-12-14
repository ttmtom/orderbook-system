BEGIN;
    DROP TABLE wallet_histories;
    DROP TYPE wallet_action_type;

    DROP TABLE wallets;
    DROP TYPE currency_type;
COMMIT;