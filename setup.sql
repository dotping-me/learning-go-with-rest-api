-- 0. Drops the database to make a fresh install
DROP DATABASE IF EXISTS learn_go;

-- 1. Create the database
CREATE DATABASE learn_go;

-- 2. Connect to the database
\c learn_go;

-- 3. Create tables
CREATE TABLE IF NOT EXISTS "user_profiles" (
    id            SERIAL       PRIMARY KEY,
    username      VARCHAR(100) NOT NULL UNIQUE,
    password      VARCHAR(255) NOT NULL,
    registered_at TIMESTAMP    DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS "posts" (
    id              SERIAL    PRIMARY KEY,
    content         TEXT      NOT NULL,
    likes           INT       DEFAULT 0,
    dislikes        INT       DEFAULT 0,
    posted_at       TIMESTAMP DEFAULT NOW(),
    user_profile_id INT       NOT NULL REFERENCES user_profiles(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS "comments" (
    id              SERIAL    PRIMARY KEY,
    content         TEXT      NOT NULL,
    commented_at    TIMESTAMP DEFAULT NOW(),
    user_profile_id INT       NOT NULL REFERENCES user_profiles(id),
    post_id         INT       NOT NULL REFERENCES posts(id) ON DELETE CASCADE
);

-- Indexing to speed up queries
CREATE INDEX idx_posts_user_profile_id ON posts(user_profile_id);
CREATE INDEX idx_comments_post_id ON comments(post_id);
CREATE INDEX idx_comments_user_profile_id ON comments(user_profile_id);