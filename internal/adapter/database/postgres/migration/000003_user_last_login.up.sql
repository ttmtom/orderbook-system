BEGIN;
    ALTER TABLE users ADD last_login_at TIMESTAMPTZ;
    ALTER TABLE users ADD last_access_at TIMESTAMPTZ;
COMMIT;
