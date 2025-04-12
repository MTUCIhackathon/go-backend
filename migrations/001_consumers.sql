CREATE TABLE consumers(
    id UUID PRIMARY KEY,
    login string NOT NULL,
    password string NOT NULL,
    created_at timestamp NOT NULL,
);
---- create above / drop below ----
DROP TABLE consumers;
