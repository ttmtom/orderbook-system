BEGIN;
    ALTER TABLE users ADD last_login_at TIMESTAMP;
    ALTER TABLE users ADD last_access_at TIMESTAMP;
COMMIT;
