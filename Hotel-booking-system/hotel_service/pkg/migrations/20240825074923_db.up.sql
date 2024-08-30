CREATE TABLE IF NOT EXISTS hotels(
    id SERIAL PRIMARY KEY,
    name VARCHAR(250),
    location VARCHAR(250),
    rating INT,
    address VARCHAR(250)
);

CREATE TABLE IF NOT EXISTS rooms (
    hotel_id INT,
    id SERIAL PRIMARY KEY,
    room_type VARCHAR(100),
    price_per_night FLOAT,
    available BOOLEAN DEFAULT TRUE
);