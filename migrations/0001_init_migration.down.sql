BEGIN;

DROP TABLE IF EXISTS tgSubChecker.sub_events;
DROP TABLE IF EXISTS tgSubChecker.subscribers;
DROP TABLE IF EXISTS tgSubChecker.channels;
DROP TABLE IF EXISTS tgSubChecker.users;
DROP TYPE IF EXISTS tgSubChecker.event_type;
DROP SCHEMA IF EXISTS tgSubChecker CASCADE;

COMMIT;