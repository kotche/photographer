CREATE TABLE IF NOT EXISTS photographers (
                     id SERIAL PRIMARY KEY,
                     name TEXT NOT NULL,
                     created_at TIMESTAMP NOT NULL DEFAULT NOW()
            );

CREATE TABLE IF NOT EXISTS clients (
        id SERIAL PRIMARY KEY,
        name TEXT NOT NULL,
        created_at TIMESTAMP NOT NULL DEFAULT NOW(),
        updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
        deleted_at TIMESTAMP
    );

CREATE TABLE IF NOT EXISTS debt_client (
       id SERIAL PRIMARY KEY,
       photographer_id INTEGER,
       client_id INTEGER,
       amount INTEGER,
       occurred_at TIMESTAMP NOT NULL DEFAULT NOW(),
       CONSTRAINT fk_photographer_id FOREIGN KEY (photographer_id) REFERENCES photographers(id) ON DELETE CASCADE,
       CONSTRAINT fk_client_id FOREIGN KEY (client_id) REFERENCES clients(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS money_client (
       id SERIAL PRIMARY KEY,
       photographer_id INTEGER,
       client_id INTEGER,
       amount INTEGER,
       occurred_at TIMESTAMP NOT NULL DEFAULT NOW(),
       CONSTRAINT fk_photographer_id FOREIGN KEY (photographer_id) REFERENCES photographers(id) ON DELETE CASCADE,
       CONSTRAINT fk_client_id FOREIGN KEY (client_id) REFERENCES clients(id) ON DELETE CASCADE
)