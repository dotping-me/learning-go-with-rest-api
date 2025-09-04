-- 1. Create the database
CREATE DATABASE learn_go;

-- 2. Connect to the database
\c learn_go;

-- 3. Create the UserProfile table
CREATE TABLE IF NOT EXISTS "user_profiles" (
    id       SERIAL       PRIMARY KEY,
    username VARCHAR(100) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL
);
