CREATE TABLE games (
    id SERIAL PRIMARY KEY,
    steam_id BIGINT NOT NULL UNIQUE,
    name TEXT NOT NULL,
    price DECIMAL(10,2)
);