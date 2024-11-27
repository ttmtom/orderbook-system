BEGIN;
    ALTER TABLE users DROP COLUMN last_login_at;
    ALTER TABLE users DROP COLUMN last_access_at;
COMMIT;
