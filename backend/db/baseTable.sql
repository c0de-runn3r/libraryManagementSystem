CREATE TABLE "users"
(
    "id" CHAR NOT NULL PRIMARY KEY,
    "email" VARCHAR NOT NULL 
    "password" VARCHAR NOT NULL,
    "roles" VARCHAR NOT NULL
);

CREATE TABLE "books"
(
    "title" VARCHAR NOT NULL PRIMARY KEY,
    "authors" VARCHAR,
    "publisher" VARCHAR,
    "year" YEAR,
    "ISBN" VARCHAR (13),
    "other_cordes" TEXT,
    "page_count" INT,
    "genres" TEXT
);
