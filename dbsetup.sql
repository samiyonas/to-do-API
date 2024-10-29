#!/usr/bin/mysql
-- create a user
-- create a database
-- Grant the user all the necessary privileges

CREATE USER IF NOT EXISTS 'todo_user'@'localhost' IDENTIFIED BY '$USER_PASSWD';
CREATE DATABASE IF NOT EXISTS tododb;
GRANT ALL PRIVILEGES ON tododb.* TO 'todo_user'@'localhost';
FLUSH PRIVILEGES
