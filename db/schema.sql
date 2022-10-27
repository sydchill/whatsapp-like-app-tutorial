CREATE SCHEMA IF NOT EXISTS hidotcom;


-- user
CREATE TABLE IF NOT EXISTS hidotcom.user (
    user_id SERIAL PRIMARY KEY,
    user_password TEXT NOT NULL,
    token TEXT NOT NULL,
    refresh_token TEXT NOT NULL,
    creation_date TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT(now() AT TIME ZONE 'utc')
);

--  user_profile
CREATE TABLE IF NOT EXISTS hidotcom.user_profile (
    user_id BIGINT NOT NULL REFERENCES hidotcom.user(user_id),
    profile_image BYTEA,
    username TEXT NOT NULL,
    email TEXT NOT NULL,
    "name" TEXT,
    surname TEXT,
    about TEXT
);
