BEGIN;

CREATE SCHEMA IF NOT EXISTS tgSubChecker;

CREATE TYPE tgSubChecker.event_type AS ENUM ('subscribed', 'unsubscribed');

CREATE TABLE IF NOT EXISTS tgSubChecker.users (
    id BIGINT PRIMARY KEY,
    username VARCHAR(80),
    first_name VARCHAR(80),
    last_name VARCHAR(80),
    is_bot BOOLEAN,
    is_premium BOOLEAN,
    notifications BOOLEAN
);

CREATE TABLE if NOT EXISTS tgSubChecker.channels (
    id BIGINT PRIMARY KEY,
    name VARCHAR(80) UNIQUE,
    owner_id BIGINT REFERENCES tgSubChecker.users(id)
);

CREATE TABLE IF NOT EXISTS tgSubChecker.subscribers (
    id BIGINT PRIMARY KEY,
    username VARCHAR(80),
    first_name VARCHAR(80),
    last_name VARCHAR(80),
    invate_link VARCHAR(80),
    is_bot BOOLEAN,
    is_premium BOOLEAN
);

CREATE TABLE IF NOT EXISTS tgSubChecker.sub_events (
    id SERIAL PRIMARY KEY,
    chat_id BIGINT REFERENCES tgSubChecker.channels(id),
    subscriber_id BIGINT REFERENCES tgSubChecker.subscribers(id),
    event_type tgSubChecker.event_type,
    event_time TIMESTAMP
);


COMMIT;