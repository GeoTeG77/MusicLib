CREATE TABLE IF NOT EXISTS song_details (
    id SERIAL PRIMARY KEY,
    release_date DATE NOT NULL,
    text TEXT NOT NULL,
    link TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS songs (
    group_name TEXT NOT NULL,
    song_name TEXT NOT NULL,
    detail_id INT NOT NULL,
    PRIMARY KEY (group_name, song_name),
    FOREIGN KEY (detail_id) REFERENCES song_details (id) ON DELETE CASCADE
);

