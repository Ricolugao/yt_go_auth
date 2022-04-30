create table users (
	id serial,
    nome varchar(255),
    email varchar(255) unique,
    password varchar(255)
);