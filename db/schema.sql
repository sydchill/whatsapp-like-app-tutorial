CREATE SCHEMA IF NOT EXISTS hidotcom;


-- user
CREATE TABLE IF NOT EXISTS hidotcom.user (
    user_id SERIAL PRIMARY KEY,
    username TEXT NOT NULL,
    email TEXT NOT NULL,
    user_password TEXT NOT NULL,
    token TEXT NOT NULL,
    refresh_token TEXT NOT NULL,
    creation_date TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT(now() AT TIME ZONE 'utc')
)
