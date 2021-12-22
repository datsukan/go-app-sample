CREATE TABLE IF NOT EXISTS users(
    id bigint AUTO_INCREMENT NOT NULL primary key,
    username varchar(50) UNIQUE NOT NULL,
    password varchar(50) NOT NULL,
    email varchar(300) UNIQUE NOT NULL,
    birthday date NOT NULL,
    image_url  varchar(2000) NOT NULL,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);