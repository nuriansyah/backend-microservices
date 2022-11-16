CREATE TABLE IF NOT EXISTS posts (
    id serial PRIMARY KEY,
    author_id integer NOT NULL,
    title varchar(255) NOT NULL,
    description varchar(255) NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    FOREIGN KEY (author_id) REFERENCES users(id)
);