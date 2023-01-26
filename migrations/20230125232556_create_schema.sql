-- +goose Up
CREATE TABLE IF NOT EXISTS users (
    user_id varchar(60) NOT NULL,
    email varchar(255) NOT NULL,
    password varchar(255) NOT NULL,
    first_name varchar(60) DEFAULT NULL,
    last_name varchar(60) DEFAULT NULL,
    created_at DATETIME DEFAULT NOW(),
    updated_at DATETIME DEFAULT NULL,
    PRIMARY KEY(user_id),
    UNIQUE(email)
);

CREATE TABLE IF NOT EXISTS user_email_confirmations(
    user_id varchar(60) NOT NULL,
    code varchar(255) NOT NULL,
    created_at DATETIME NOT NULL DEFAULT NOW(),
    FOREIGN KEY(user_id) REFERENCES users(user_id) ON DELETE CASCADE,
    UNIQUE(user_id, code)
);

CREATE TABLE IF NOT EXISTS categories(
    category_id SMALLINT NOT NULL AUTO_INCREMENT,
    parent_id SMALLINT NOT NULL DEFAULT 0,
    created_at DATETIME DEFAULT NOW(),
    update_at DATETIME DEFAULT NULL,
    name varchar(255),
    PRIMARY KEY(category_id)
);

CREATE TABLE IF NOT EXISTS ad(
    ad_id INT(11) NOT NULL AUTO_INCREMENT,
    category_id SMALLINT,
    creator_id varchar(60) NOT NULL,
    title varchar(255) NOT NULL,
    description text,
    location_lat_long varchar(60) DEFAULT NULL,
    create_at DATETIME DEFAULT NOW(),
    updated_at DATETIME DEFAULT NULL,
    PRIMARY KEY(ad_id),
    FOREIGN KEY(creator_id) REFERENCES users(user_id) ON DELETE CASCADE,
    FOREIGN KEY(category_id) REFERENCES categories(category_id) ON DELETE SET NULL
);

