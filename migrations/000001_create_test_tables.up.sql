CREATE TABLE something
(
    column1 INT,
    column2 INT,
    column3 INT,
    PRIMARY KEY (column1, column2)
);

INSERT INTO something (column1, column2, column3)
VALUES (1, 2, 3);
INSERT INTO something (column1, column2, column3)
VALUES (1, 3, 3);
