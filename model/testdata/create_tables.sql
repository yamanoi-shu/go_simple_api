DROP DATABASE IF EXISTS test_db;

CREATE DATABASE IF NOT EXISTS test_db;

USE test_db;

CREATE TABLE users(
    id INT AUTO_INCREMENT,
    first_name VARCHAR(255),
    last_name VARCHAR(255),
    PRIMARY KEY(id),
    UNIQUE name_index (first_name, last_name)
);

INSERT INTO users (id, first_name, last_name) VALUES (1, 'Michael', 'Jordan');
