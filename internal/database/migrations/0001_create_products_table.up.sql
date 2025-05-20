CREATE TABLE users (
                       id SERIAL PRIMARY KEY,
                       username TEXT NOT NULL UNIQUE,
                       password TEXT NOT NULL,
                       role TEXT
);

CREATE TABLE categories (
                            id SERIAL PRIMARY KEY,
                            name VARCHAR(255) NOT NULL UNIQUE
);

CREATE TABLE products (
                          id SERIAL PRIMARY KEY,
                          name VARCHAR(255) NOT NULL,
                          price DECIMAL(10,2) NOT NULL,
                          stock INTEGER DEFAULT 0,
                          category_id INTEGER NOT NULL REFERENCES categories(id) ON DELETE CASCADE
);
