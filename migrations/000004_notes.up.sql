CREATE TABLE IF NOT EXISTS notes (
    id SERIAL NOT NULL PRIMARY KEY,
    title varchar,
    note varchar,
    user_id int,
    FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE
)