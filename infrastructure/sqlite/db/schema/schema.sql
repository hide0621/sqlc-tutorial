CREATE TABLE books (
    id INTEGER PRIMARY KEY,
    name text NOT NULL,
    release_year INTEGER NOT NULL,
    total_page INTEGER NOT NULL
);

CREATE TABLE categories (
    id INTEGER PRIMARY KEY,
    name text NOT NULL
);

CREATE TABLE book_categories (
    id INTEGER PRIMARY KEY,
    book_id INTEGER,
    category_id INTEGER,
    FOREIGN KEY (book_id) REFERENCES books(id),
    FOREIGN KEY (category_id) REFERENCES categories(id)
);