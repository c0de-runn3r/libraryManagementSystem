CREATE TABLE IF NOT EXISTS "users"
(
    "id" CHAR NOT NULL PRIMARY KEY,
    "email" VARCHAR NOT NULL 
    "password" VARCHAR NOT NULL,
    "roles" VARCHAR NOT NULL
);

CREATE TABLE IF NOT EXISTS "books"
(
    "title" VARCHAR NOT NULL PRIMARY KEY,
    "authors" VARCHAR,
    "genres" TEXT,
    "publisher" VARCHAR,
    "year" YEAR,
    "ISBN" VARCHAR (13),
    "UDK_number" TEXT,
    "page_count" INT
);

CREATE TABLE IF NOT EXISTS "authors"
(
    "author" VARCHAR NOT NULL
);

CREATE TABLE IF NOT EXISTS "BooksAuthorsRel"
(
    "authors" VARCHAR NOT NULL,
    "author" VARCHAR NOT NULL,
    FOREIGN KEY "authors" REFERENCES books("authors"),
    FOREIGN KEY "author" REFERENCES author("author")
);
