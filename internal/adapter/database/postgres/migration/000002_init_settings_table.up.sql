BEGIN;
    CREATE TABLE "time_limits"
    (
        "id"   varchar PRIMARY KEY,
        "time" numeric CHECK ("time" >= 0)
    );
COMMIT;
