BEGIN;
    DROP TABLE transaction_events;
    DROP TYPE transaction_event_type;
    DROP TABLE transactions;
    DROP TYPE transaction_type;
    ALTER TABLE wallet_histories DROP COLUMN transaction_id;
COMMIT;