CREATE TABLE users (
                       id SERIAL PRIMARY KEY,
                       username TEXT NOT NULL UNIQUE,
                       password TEXT NOT NULL,
                       role TEXT DEFAULT 'user'
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
                          category_id INTEGER NOT NULL REFERENCES categories(id) ON DELETE CASCADE,
                          user_id INTEGER REFERENCES users(id) ON DELETE CASCADE
);

CREATE TABLE outfits (
                         id SERIAL PRIMARY KEY,
                         name VARCHAR(255) NOT NULL,
                         user_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE
);

CREATE TABLE outfit_products (
                                 outfit_id INTEGER NOT NULL REFERENCES outfits(id) ON DELETE CASCADE,
                                 product_id INTEGER NOT NULL REFERENCES products(id) ON DELETE CASCADE,
                                 PRIMARY KEY (outfit_id, product_id)
);