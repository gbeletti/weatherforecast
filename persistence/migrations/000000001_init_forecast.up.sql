BEGIN;

CREATE TABLE IF NOT EXISTS "forecast" (
    date timestamp without time zone PRIMARY KEY,
    wind_speed numeric(6,2) NOT NULL,
    wind_direction smallint NOT NULL,
    alert boolean NOT NULL default false
);

CREATE TABLE IF NOT EXISTS "configuration"(
    id int PRIMARY KEY,
    last_update timestamp with time zone NULL
);

CREATE INDEX IF NOT EXISTS "forecast_alert_idx" ON "forecast" (alert);

INSERT INTO "configuration" (id, last_update) VALUES (1, NULL);

COMMIT;