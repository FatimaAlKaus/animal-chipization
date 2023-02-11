CREATE TABLE animal(
    id BIGSERIAL PRIMARY KEY,
    weight REAL NOT NULL,
    length REAL NOT NULL,
    height REAL NOT NULL,
    gender_id INT NOT NULL REFERENCES gender(id),
    life_status_id INT NOT NULL REFERENCES life_status(id),
    chipper_id INT NOT NULL REFERENCES account(id),
    chipping_date INT NOT NULL,
    chipping_location BIGINT NOT NULL REFERENCES location_point(id),
    death_date BIGINT NULL
);