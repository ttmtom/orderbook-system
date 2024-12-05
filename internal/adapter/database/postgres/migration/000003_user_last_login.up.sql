BEGIN;
    ALTER TABLE users ADD last_login_at TIMESTAMPTZ default null;
    ALTER TABLE users ADD last_access_at TIMESTAMPTZ default null;
COMMIT;
