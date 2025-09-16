-- Active: 1757069669177@@127.0.0.1@5432@social_network_db
CREATE DATABASE social_network_db;

CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    user_name VARCHAR(50) NOT NULL UNIQUE,
    email VARCHAR(100) NOT NULL UNIQUE,
    age INT CHECK (age >= 0)
);

INSERT INTO users (user_name, email, age) VALUES
('alice',  'alice@gmail.com', 25),
('bob',    'bob@yahoo.com', 30),
('charlie','charlie@outlook.com', 28),
('diana',  'diana@gmail.com', 22),
('edward', 'edward@hotmail.com', 35),
('fiona',  'fiona@yandex.ru', 27),
('george', 'george@icloud.com', 40),
('hannah', 'hannah@protonmail.com', 31),
('ivan',   'ivan@mail.ru', 29),
('julia',  'julia@zoho.com', 26);

DROP TABLE users;

DROP DATABASE social_network_db;