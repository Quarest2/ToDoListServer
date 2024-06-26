CREATE TABLE IF NOT EXISTS tasks
(
    id       SERIAL PRIMARY KEY,
    username TEXT,
    task     TEXT,
    deadline TIMESTAMP WITHOUT TIME ZONE,
    is_done  BOOLEAN DEFAULT false
);