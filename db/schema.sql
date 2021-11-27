-- Create tables.
DROP TABLE IF EXISTS balancers CASCADE;
CREATE TABLE balancers
(
    id SERIAL PRIMARY KEY
);

DROP TABLE IF EXISTS machines CASCADE;
CREATE TABLE machines
(
    id SERIAL PRIMARY KEY,
    is_working BOOLEAN NOT NULL
);

DROP TABLE IF EXISTS connections;
CREATE TABLE connections
(
    machine_id SERIAL REFERENCES machines (id) ON DELETE CASCADE,
    balancer_id SERIAL REFERENCES balancers (id) ON DELETE CASCADE,
    PRIMARY KEY (machine_id, balancer_id)
);
