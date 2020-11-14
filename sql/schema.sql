CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE EXTENSION IF NOT EXISTS "pgcrypto";
 
CREATE DATABASE guestbuku;
 \c guestbuku

CREATE TABLE "tamus"(
    id UUID NOT NULL DEFAULT gen_random_uuid(),
    email TEXT NOT NULL DEFAULT '',
    name TEXT NOT NULL DEFAULT '',
    date TIMESTAMPTZ NULL DEFAULT now()::TIMESTAMPTZ,
    discription TEXT NOT NULL DEFAULT '',
    PRIMARY KEY (id)
);
 