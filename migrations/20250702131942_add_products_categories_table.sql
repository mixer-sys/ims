-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE TABLE categories (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);

CREATE TABLE products (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    category_id INT REFERENCES categories(id),
    price DECIMAL(10, 2) NOT NULL
);

INSERT INTO categories (name) VALUES
('Electronics'),
('Clothing'),
('Books');

INSERT INTO products (name, category_id, price) VALUES
('Smartphone', 1, 100000),
('Laptop', 1, 120000),
('T-shirt', 2, 1200),
('Jeans', 2, 3500),
('Fiction Book', 3, 700),
('Non-Fiction Book', 3, 800);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DROP TABLE products;
DROP TABLE categories;
-- +goose StatementEnd
