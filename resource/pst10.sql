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
