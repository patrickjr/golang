CREATE TABLE IF NOT EXISTS users  (
    email character varying(100),
    password character varying(40),
    forgot_password boolean,
    user_name character varying(40)
);