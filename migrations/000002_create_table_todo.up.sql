CREATE TABLE IF NOT EXISTS todos
(
    id                int          NOT NULL AUTO_INCREMENT,
    title             varchar(255) NOT NULL,
    priority          varchar(255),
    activity_group_id int          NOT NULL,
    is_active         bool,
    created_at        timestamp default now(),
    updated_at        timestamp default null,

    PRIMARY KEY (id),
    FOREIGN KEY (activity_group_id) REFERENCES activities (id)
);

