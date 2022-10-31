CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    balance INTEGER
);

CREATE TABLE buffer (
    id SERIAL PRIMARY KEY,
    balance INTEGER,
    user_id INTEGER,
    FOREIGN KEY (user_id) REFERENCES users(id)
);

CREATE TABLE avito (
    id SERIAL PRIMARY KEY,
    balance INTEGER
);

INSERT INTO users (balance) VALUES (0);
INSERT INTO users (balance) VALUES (0);
INSERT INTO avito (balance) VALUES (0);