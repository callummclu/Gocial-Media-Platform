/*these arent currently correct*/

CREATE TABLE users (
    id bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    name varchar(255) NOT NULL,
    surname varchar(255) NOT NULL,
    username varchar(255) NOT NULL UNIQUE,
    email varchar(255) NOT NULL UNIQUE,
    email_verified_at timestamp DEFAULT NULL,
    password varchar(255) NOT NULL,
    created_at timestamp DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    UNIQUE KEY users_username_unique (username),
    UNIQUE KEY users_email_unique (email)
) 

CREATE TABLE posts (
    id bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    title varchar(255) NOT NULL,
    content varchar(255) NOT NULL,
    PRIMARY KEY (id),

);

CREATE TABLE images (
    uuid bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    uri varchar(255) NOT NULL,
    PRIMARY KEY (uuid),
);