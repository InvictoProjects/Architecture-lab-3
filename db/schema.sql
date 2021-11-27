-- Create tables.
DROP TABLE IF EXISTS balancers;
CREATE TABLE balancers
(
    id SERIAL PRIMARY KEY
);

DROP TABLE IF EXISTS machines;
CREATE TABLE machines
(
    id SERIAL PRIMARY KEY,
    is_working BOOLEAN NOT NULL
);
