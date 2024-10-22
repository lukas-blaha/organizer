## Create database

CREATE DATABASE organizer;

## Create table for reminders

CREATE TABLE reminders (
    id SERIAL PRIMARY KEY,
    start TEXT NOT NULL,
    repeat INTEGER NOT NULL,
    last TEXT NOT NULL,
    category TEXT NOT NULL,
    name TEXT NOT NULL
);
