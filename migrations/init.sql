CREATE TABLE users
(
    id              SERIAL PRIMARY KEY,
    passport_number VARCHAR(50) UNIQUE NOT NULL,
    surname         VARCHAR(100),
    name            VARCHAR(100),
    patronymic      VARCHAR(100),
    address         VARCHAR(255),
    created_at      TIMESTAMP,
    updated_at      TIMESTAMP
);

CREATE TABLE tasks
(
    id         SERIAL PRIMARY KEY,
    user_id    INTEGER REFERENCES users (id),
    name       VARCHAR(255),
    duration   INTEGER,
    status     VARCHAR(50),
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);
