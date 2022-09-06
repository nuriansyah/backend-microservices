CREATE TABLE  users (
    id INTEGER PRIMARY KEY,
    name varchar(255) NOT NULL,
    email varchar(255) NOT NULL UNIQUE,
    password varchar(255) NOT NULL,
    role varchar(255) NOT NULL,
    avatar varchar(255) null
);
CREATE TABLE user_details (
    user_id INTEGER  NOT NULL,
    nrp varchar(9) NOT NULL,
    prodi varchar(255) NOT NULL,
    program varchar(255) NOT NULL,
    company varchar(255) NULL,
    batch smallint NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id)
);
