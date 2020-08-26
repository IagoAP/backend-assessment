CREATE TABLE IF NOT EXISTS model
(
    id_externalApp INT,
    id_product VARCHAR ( 255 ),
    id_superUser INT,
    description VARCHAR ( 255 ),
    customer_mid INT,
    customer_email VARCHAR ( 255 ),
    activated BOOLEAN DEFAULT NULL
);
