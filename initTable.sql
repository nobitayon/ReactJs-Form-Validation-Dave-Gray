-- CREATE TABLE
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE users (
    uid uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
    username VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL
);