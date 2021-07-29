CREATE TABLE IF NOT EXISTS services(
    id serial PRIMARY KEY,
    repo_url TEXT,
    name VARCHAR (45) UNIQUE NOT NULL
);
