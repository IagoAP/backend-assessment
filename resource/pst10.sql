CREATE TABLE IF NOT EXISTS entities
(
    id SERIAL PRIMARY KEY,
    user_type VARCHAR ( 50 ) NOT NULL,
    name VARCHAR ( 50 ) NOT NULL,
    username VARCHAR ( 50 ) UNIQUE NOT NULL,
    password VARCHAR ( 50 ) NOT NULL
);

CREATE TABLE IF NOT EXISTS tokens
(
    id_entities INT REFERENCES entities (id),
    token VARCHAR ( 255 ) NOT NULL,
    expiration_time timestamptz DEFAULT current_timestamp(0)
);

CREATE TABLE IF NOT EXISTS product
(
    id          VARCHAR ( 255 ),
    description VARCHAR ( 255 ),
    customer_mid INT,
    customer_email VARCHAR ( 255 ),
    externalApp_id INT REFERENCES entities (id) DEFAULT NULL,
    superuser_id INT REFERENCES entities (id) DEFAULT NULL,
    activated BOOLEAN DEFAULT NULL
);

insert into entities (user_type, name, username, password) values ('ExternalApp', 'Empresa1', 'Empresa1', 'Empresa1');
insert into entities (user_type, name, username, password) values ('ExternalApp', 'Empresa2', 'Empresa2', 'Empresa2');
insert into entities (user_type, name, username, password) values ('ExternalApp', 'Empresa3', 'Empresa3', 'Empresa3');
insert into entities (user_type, name, username, password) values ('ExternalApp', 'Empresa4', 'Empresa4', 'Empresa4');
insert into entities (user_type, name, username, password) values ('ExternalApp', 'Empresa5', 'Empresa5', 'Empresa5');
insert into entities (user_type, name, username, password) values ('ExternalApp', 'Empresa6', 'Empresa6', 'Empresa6');
insert into entities (user_type, name, username, password) values ('ExternalApp', 'Empresa7', 'Empresa7', 'Empresa7');
insert into entities (user_type, name, username, password) values ('ExternalApp', 'Empresa8', 'Empresa8', 'Empresa8');
insert into entities (user_type, name, username, password) values ('ExternalApp', 'Empresa9', 'Empresa9', 'Empresa9');
insert into entities (user_type, name, username, password) values ('ExternalApp', 'Empresa10', 'Empresa10', 'Empresa10');
insert into entities (user_type, name, username, password) values ('ExternalApp', 'Empresa11', 'Empresa11', 'Empresa11');
insert into entities (user_type, name, username, password) values ('ExternalApp', 'Empresa12', 'Empresa12', 'Empresa12');
insert into entities (user_type, name, username, password) values ('ExternalApp', 'Empresa13', 'Empresa13', 'Empresa13');
insert into entities (user_type, name, username, password) values ('ExternalApp', 'Empresa14', 'Empresa14', 'Empresa14');
insert into entities (user_type, name, username, password) values ('ExternalApp', 'Empresa15', 'Empresa15', 'Empresa15');
insert into entities (user_type, name, username, password) values ('ExternalApp', 'Empresa16', 'Empresa16', 'Empresa16');
insert into entities (user_type, name, username, password) values ('ExternalApp', 'Empresa17', 'Empresa17', 'Empresa17');
insert into entities (user_type, name, username, password) values ('ExternalApp', 'Empresa18', 'Empresa18', 'Empresa18');
insert into entities (user_type, name, username, password) values ('ExternalApp', 'Empresa19', 'Empresa19', 'Empresa19');
insert into entities (user_type, name, username, password) values ('ExternalApp', 'Empresa20', 'Empresa20', 'Empresa20');
insert into entities (user_type, name, username, password) values ('ExternalApp', 'Empresa21', 'Empresa21', 'Empresa21');
insert into entities (user_type, name, username, password) values ('ExternalApp', 'Empresa22', 'Empresa22', 'Empresa22');
insert into entities (user_type, name, username, password) values ('ExternalApp', 'Empresa23', 'Empresa23', 'Empresa23');
insert into entities (user_type, name, username, password) values ('ExternalApp', 'Empresa24', 'Empresa24', 'Empresa24');
insert into entities (user_type, name, username, password) values ('ExternalApp', 'Empresa25', 'Empresa25', 'Empresa25');
insert into entities (user_type, name, username, password) values ('ExternalApp', 'Empresa26', 'Empresa26', 'Empresa26');
insert into entities (user_type, name, username, password) values ('ExternalApp', 'Empresa27', 'Empresa27', 'Empresa27');
insert into entities (user_type, name, username, password) values ('ExternalApp', 'Empresa28', 'Empresa28', 'Empresa28');
insert into entities (user_type, name, username, password) values ('ExternalApp', 'Empresa29', 'Empresa29', 'Empresa29');
insert into entities (user_type, name, username, password) values ('ExternalApp', 'Empresa30', 'Empresa30', 'Empresa30');

insert into entities (user_type, name, username, password) values ('SuperUser', 'User1', 'User1', 'User1');
insert into entities (user_type, name, username, password) values ('SuperUser', 'User2', 'User2', 'User2');
insert into entities (user_type, name, username, password) values ('SuperUser', 'User3', 'User3', 'User3');
insert into entities (user_type, name, username, password) values ('SuperUser', 'User4', 'User4', 'User4');
insert into entities (user_type, name, username, password) values ('SuperUser', 'User5', 'User5', 'User5');
insert into entities (user_type, name, username, password) values ('SuperUser', 'User6', 'User6', 'User6');
insert into entities (user_type, name, username, password) values ('SuperUser', 'User7', 'User7', 'User7');
insert into entities (user_type, name, username, password) values ('SuperUser', 'User8', 'User8', 'User8');
insert into entities (user_type, name, username, password) values ('SuperUser', 'User9', 'User9', 'User9');
insert into entities (user_type, name, username, password) values ('SuperUser', 'User10', 'User10', 'User10');
