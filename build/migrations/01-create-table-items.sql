CREATE TABLE IF NOT EXISTS items (
    id SERIAL PRIMARY KEY,
    name varchar(100) NOT NULL, 
    description varchar(255),
    price DECIMAL(10, 2) NOT NULL,
    stock INT NOT NULL,
    status VARCHAR(20) NOT NULL
);