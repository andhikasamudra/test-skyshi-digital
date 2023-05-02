CREATE TABLE IF NOT EXISTS activities
(
    id int          NOT NULL AUTO_INCREMENT,
    title       varchar(255) NOT NULL,
    email       varchar(255) NOT NULL,
    created_at timestamp default now(),
    updated_at timestamp default null,
    PRIMARY KEY (id)
);