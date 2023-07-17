DROP TABLE IF EXISTS users;
CREATE TABLE users (
    id INTEGER PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(100) NOT NULL
);

DROP TABLE IF EXISTS follows;
CREATE TABLE follows (
    id INTEGER PRIMARY KEY AUTO_INCREMENT,
    follower_id INTEGER NOT NULL,
    followee_id INTEGER NOT NULL,
    UNIQUE KEY (follower_id, followee_id),
    FOREIGN KEY (follower_id) REFERENCES users(id),
    FOREIGN KEY (followee_id) REFERENCES users(id)
);

DROP TABLE IF EXISTS posts;
CREATE TABLE posts (
    id INTEGER PRIMARY KEY AUTO_INCREMENT,
    user_id INTEGER NOT NULL,
    body TEXT NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id)
);
