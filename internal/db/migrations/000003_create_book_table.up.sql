CREATE TABLE book(
                     id INTEGER AUTO_INCREMENT PRIMARY KEY,
                     name VARCHAR NOT NULL,
                     released DATE NOT NULL,
                     coast INTEGER NOT NULL,
                     pages INTEGER NOT NULL,
                     poster VARCHAR NOT NULL,
                     author_id INTEGER REFERENCES author(id) ON DELETE CASCADE ON UPDATE CASCADE,
                     genre_id INTEGER REFERENCES genre(id) ON DELETE CASCADE ON UPDATE CASCADE
);
