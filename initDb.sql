-- Connect to the default database as a superuser (e.g., postgres)
-- CREATE USER
CREATE USER someone WITH PASSWORD 'a';

-- CREATE DATABASE
CREATE DATABASE random_db;

-- GRANT PRIVILEGES
GRANT ALL PRIVILEGES ON DATABASE random_db TO someone;
GRANT ALL ON users TO someone;
