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

DROP TABLE IF EXISTS connections;
CREATE TABLE connections
(
    machine_id int REFERENCES machines (id),
    balancer_id int REFERENCES balancers (id),
    PRIMARY KEY (machine_id, balancer_id)
);
