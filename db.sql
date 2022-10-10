CREATE TABLE albums (
    "id" VARCHAR(30) PRIMARY KEY,
    "name" TEXT NOT NULL,
    "year" INT NOT NULL
)

CREATE TABLE songs (
    "id" VARCHAR(30) PRIMARY KEY,
    "title" TEXT NOT NULL,
    "year" INT NOT NULL,
    "genre"  TEXT NOT NULL,
    "performer"  TEXT NOT NULL,
    "duration" INT NOT NULL,
    "albumId" VARCHAR(30) NULL REFERENCES albums(id) ON DELETE CASCADE
);