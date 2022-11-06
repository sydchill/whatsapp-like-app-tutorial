CREATE OR REPLACE FUNCTION hidotcom.create_user (
    name_param TEXT,
    surname_param TEXT,
    username_param TEXT,
    email_param TEXT,
    hash_password_param TEXT,
    token_param TEXT,
    refresh_token_param TEXT,
)
RETURNS TABLE (
    user_id BIGINT
) AS $$
DECLARE
    user_id_variable BIGINT;
BEGIN

    INSERT INTO hidotcom.user AS u
    (hash_password, token, refresh_token)
    VALUES
    (hash_password_param, token_param, refresh_token_param)
    RETURNING u.user_id INTO user_id_variable;

    INSERT INTO hidotcom.user_profile
    (user_id, username, "name", surname, email)
    VALUES
    (user_id_variable, username_param, name_param, surname_param, email_param);

    RETURN QUERY
    SELECT user_id_variable;
END;
$$ LANGUAGE plpgsql;