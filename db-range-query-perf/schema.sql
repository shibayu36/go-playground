-- mysqladmin -u root drop db-range-query-perf
-- mysqladmin -u root create db-range-query-perf
-- mysql -u root db-range-query-perf < db-range-query-perf/schema.sql

CREATE TABLE users (
    id INTEGER PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(100) NOT NULL
);

CREATE TABLE follows (
    id INTEGER PRIMARY KEY AUTO_INCREMENT,
    follower_id INTEGER NOT NULL,
    followee_id INTEGER NOT NULL,
    UNIQUE KEY (follower_id, followee_id),
    FOREIGN KEY (follower_id) REFERENCES users(id),
    FOREIGN KEY (followee_id) REFERENCES users(id)
);

CREATE TABLE posts (
    id INTEGER PRIMARY KEY AUTO_INCREMENT,
    user_id INTEGER NOT NULL,
    body TEXT NOT NULL,
    posted_at DATETIME NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id)
);
