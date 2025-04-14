CREATE TYPE test_type AS ENUM ('first_order_test', 'second_order_test', 'third_order_test');

CREATE TABLE consumers
(
    id         UUID           NOT NULL,
    email      VARCHAR UNIQUE NULLS NOT DISTINCT,
    login      VARCHAR UNIQUE NOT NULL,
    password   VARCHAR        NOT NULL,
    created_at TIMESTAMP      NOT NULL,
    CONSTRAINT consumers PRIMARY KEY (id)
);

CREATE TABLE resolved
(
    id            UUID      NOT NULL,
    user_id       UUID      NOT NULL,
    resolved_type test_type NOT NULL,
    is_active     BOOL      NOT NULL DEFAULT TRUE,
    created_at    TIMESTAMP NOT NULL,
    passed_at     TIMESTAMP          DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT resolved_pk PRIMARY KEY (id),
    CONSTRAINT resolved_consumers_fk FOREIGN KEY (user_id) REFERENCES consumers (id) ON DELETE CASCADE
);

CREATE TABLE resolved_questions
(
    resolved_id     UUID     NOT NULL,
    question_order  SMALLINT NOT NULL,
    question_text   VARCHAR  NOT NULL,
    question_answer VARCHAR,
    image_location  VARCHAR,
    mark            SMALLINT NOT NULL,
    CONSTRAINT questions_pk PRIMARY KEY (resolved_id, question_order),
    CONSTRAINT questions_resolved_fk FOREIGN KEY (resolved_id) REFERENCES resolved (id) ON DELETE CASCADE
);

CREATE TABLE test_results
(
    id             UUID      NOT NULL,
    user_id        UUID      NOT NULL,
    resolved_id    UUID      NOT NULL,
    image_location VARCHAR,
    profession     VARCHAR[] NOT NULL,
    created_at     TIMESTAMP NOT NULL,
    CONSTRAINT test_results_pk PRIMARY KEY (id, user_id, resolved_id),
    CONSTRAINT test_results_consumers_fk FOREIGN KEY (user_id) REFERENCES consumers (id) ON DELETE CASCADE,
    CONSTRAINT test_results_resolved_fk FOREIGN KEY (resolved_id) REFERENCES resolved (id) ON DELETE CASCADE
);
---- create above / drop below ----
DROP TYPE test_type;
DROP TABLE consumers;
DROP TABLE resolved;
DROP TABLE resolved_questions;
