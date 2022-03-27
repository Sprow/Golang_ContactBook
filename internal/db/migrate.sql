CREATE TABLE IF NOT EXISTS users(
    id uuid PRIMARY KEY NOT NULL,
    created_at TIMESTAMPTZ NOT NULL
);

CREATE TABLE IF NOT EXISTS sessions(
    token VARCHAR(40) NOT NULL,
    user_id uuid REFERENCES users (id) on delete cascade on update cascade
);

CREATE TABLE IF NOT EXISTS contacts(
    id uuid PRIMARY KEY NOT NULL,
    name VARCHAR(50),
    phone_number VARCHAR(40)
);

CREATE TABLE IF NOT EXISTS user_contacts(
    user_id uuid REFERENCES users (id) on delete cascade on update cascade,
    contact_id uuid REFERENCES contacts (id) on delete cascade on update cascade
);

-- DROP TABLE users, sessions, contacts, user_contacts