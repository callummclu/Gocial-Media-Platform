/*these arent currently correct*/

CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name varchar(255) NOT NULL,
    surname varchar(255) NOT NULL,
    username varchar(255) NOT NULL UNIQUE,
    email varchar(255) NOT NULL UNIQUE,
    email_verified_at timestamp DEFAULT NULL,
    description varchar(255),
    friends text[] DEFAULT NULL,
    received_invitations text[] DEFAULT NULL,
    sent_invitations text[] DEFAULT NULL,
    password varchar(255) NOT NULL,
    created_at timestamp DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
) 

CREATE TABLE posts (
    id SERIAL PRIMARY KEY,
    title varchar(255) NOT NULL,
    content varchar(255) NOT NULL,
    username varchar(255) NOT NULL,
    created_at timestamp DEFAULT CURRENT_TIMESTAMP

);

CREATE TABLE images (
    uuid SERIAL PRIMARY KEY,
    uri varchar(255) NOT NULL,
);