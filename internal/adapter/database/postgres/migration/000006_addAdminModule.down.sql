BEGIN;
    DROP INDEX IF exists "admin_username";

    DROP TABLE IF EXISTS "admin_users";
COMMIT;