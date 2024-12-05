BEGIN;
    CREATE TABLE "time_limits"
    (
        "id"   varchar PRIMARY KEY,
        "time" numeric CHECK ("time" >= 0)
    );

    INSERT INTO "time_limits" (id, time) VALUES ('refresh_token', '1814400');
    INSERT INTO "time_limits" (id, time) VALUES ('access_token', '600');
COMMIT;
