BEGIN;
    DROP INDEX IF exists "admin_email";

    DROP TABLE IF EXISTS "admin_users";
COMMIT;