CREATE TABLE IF NOT EXISTS environments(
    id serial PRIMARY KEY,
    name VARCHAR (45) UNIQUE NOT NULL,
    preview BOOL
);
