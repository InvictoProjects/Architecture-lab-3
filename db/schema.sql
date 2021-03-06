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

-- Insert demo data.
INSERT INTO balancers VALUES (DEFAULT);
INSERT INTO balancers VALUES (DEFAULT);
INSERT INTO balancers VALUES (DEFAULT);
INSERT INTO balancers VALUES (DEFAULT);

INSERT INTO machines (is_working) VALUES (TRUE);
INSERT INTO machines (is_working) VALUES (TRUE);
INSERT INTO machines (is_working) VALUES (TRUE);
INSERT INTO machines (is_working) VALUES (FALSE);
INSERT INTO machines (is_working) VALUES (TRUE);
INSERT INTO machines (is_working) VALUES (FALSE);
INSERT INTO machines (is_working) VALUES (TRUE);
INSERT INTO machines (is_working) VALUES (FALSE);
INSERT INTO machines (is_working) VALUES (TRUE);
INSERT INTO machines (is_working) VALUES (FALSE);

INSERT INTO connections (machine_id, balancer_id) VALUES (1, 1);
INSERT INTO connections (machine_id, balancer_id) VALUES (2, 1);
INSERT INTO connections (machine_id, balancer_id) VALUES (3, 2);
INSERT INTO connections (machine_id, balancer_id) VALUES (4, 2);
INSERT INTO connections (machine_id, balancer_id) VALUES (5, 2);
INSERT INTO connections (machine_id, balancer_id) VALUES (6, 3);
INSERT INTO connections (machine_id, balancer_id) VALUES (7, 3);
INSERT INTO connections (machine_id, balancer_id) VALUES (8, 4);
INSERT INTO connections (machine_id, balancer_id) VALUES (9, 4);
INSERT INTO connections (machine_id, balancer_id) VALUES (10, 4);

