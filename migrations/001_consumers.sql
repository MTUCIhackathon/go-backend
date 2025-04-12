CREATE TABLE consumers(
    id UUID PRIMARY KEY,
    login varchar NOT NULL,
    password varchar NOT NULL,
    created_at timestamp NOT NULL
);
---- create above / drop below ----
DROP TABLE consumers;
