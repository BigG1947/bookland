CREATE TABLE author(
                       id INTEGER AUTO_INCREMENT PRIMARY KEY,
                       last_name VARCHAR NOT NULL,
                       first_name VARCHAR NOT NULL,
                       birthday DATE NOT NULL,
                       bio TEXT
);