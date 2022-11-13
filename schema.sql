DROP TABLE IF EXISTS todos;

CREATE TABLE todos (
    id uuid NOT NULL,
    text text,
    done bool,
    PRIMARY KEY ("id")
);
