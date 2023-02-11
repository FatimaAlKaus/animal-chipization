CREATE TABLE gender(
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);

CREATE TABLE life_status(
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);

CREATE TABLE animal_type(
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);

CREATE TABLE location_point(
    id BIGSERIAL PRIMARY KEY,
    latitude REAL NOT NULL,
    longitude REAL NOT NULL
);