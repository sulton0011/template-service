CREATE TABLE users (
    id BIGSERIAL PRIMARY KEY,
    fname VARCHAR(64),
    lname VARCHAR(64),
    username VARCHAR(64),
    email VARCHAR(124)
);