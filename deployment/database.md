## Create database

CREATE DATABASE organizer;

## Create table for reminders

CREATE TABLE reminders (
    id SERIAL PRIMARY KEY,
    start TEXT NOT NULL,
    repeat INTEGER NOT NULL,
    next TEXT NOT NULL,
    category TEXT NOT NULL,
    name TEXT NOT NULL,
    done BOOLEAN NOT NULL
);
