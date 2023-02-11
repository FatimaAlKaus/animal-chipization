CREATE TABLE animal_animaltype(
    animal_id BIGINT REFERENCES animal(id) ON UPDATE CASCADE ON DELETE CASCADE,
    type_id INT REFERENCES animal_type(id) ON UPDATE CASCADE ON DELETE CASCADE,
    CONSTRAINT animal_animaltype_pkey PRIMARY KEY (animal_id, type_id)
);

CREATE TABLE animal_visited_locations(
    animal_id BIGINT REFERENCES animal(id) ON UPDATE CASCADE ON DELETE CASCADE,
    location_id BIGINT REFERENCES location_point(id) ON UPDATE CASCADE ON DELETE CASCADE,
    CONSTRAINT animal_location_pkey PRIMARY KEY (animal_id, location_id)
);