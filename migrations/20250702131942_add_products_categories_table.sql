-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE categories (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(255) NOT NULL
);

CREATE TABLE products (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(255) NOT NULL,
    category_id UUID REFERENCES categories(id),
    price DECIMAL(10, 2) NOT NULL
);

INSERT INTO categories (id, name) VALUES
(uuid_generate_v4(), 'Electronics'),
(uuid_generate_v4(), 'Clothing'),
(uuid_generate_v4(), 'Books');

INSERT INTO products (id, name, category_id, price) VALUES
(uuid_generate_v4(), 'Smartphone', (SELECT id FROM categories WHERE name = 'Electronics'), 100000),
(uuid_generate_v4(), 'Laptop', (SELECT id FROM categories WHERE name = 'Electronics'), 120000),
(uuid_generate_v4(), 'T-shirt', (SELECT id FROM categories WHERE name = 'Clothing'), 1200),
(uuid_generate_v4(), 'Jeans', (SELECT id FROM categories WHERE name = 'Clothing'), 3500),
(uuid_generate_v4(), 'Fiction Book', (SELECT id FROM categories WHERE name = 'Books'), 700),
(uuid_generate_v4(), 'Non-Fiction Book', (SELECT id FROM categories WHERE name = 'Books'), 800);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DROP TABLE products;
DROP TABLE categories;
-- +goose StatementEnd
