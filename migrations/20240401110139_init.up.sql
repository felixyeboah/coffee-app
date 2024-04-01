CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS "coffees" (
    "id" uuid PRIMARY KEY NOT NULL DEFAULT (uuid_generate_v4()),
    "name" VARCHAR(255) NOT NULL,
    "roast" VARCHAR(255) NOT NULL,
    "region" VARCHAR(255) NOT NULL,
    "image" VARCHAR(255) NOT NULL,
    "price" FLOAT NOT NULL,
    "grind_unit" INT NOT NULL,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);