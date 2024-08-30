CREATE TABLE IF NOT EXISTS booked(
    id SERIAL PRIMARY KEY,
    user_id INT,
    hotel_id INT,
    room_id INT,
    room_type VARCHAR(250),
    enterydate DATE,
    leavingdate DATE,
    totalcost FLOAT,
    status TEXT
);

CREATE TABLE IF NOT EXISTS waitinglist(
    id SERIAL PRIMARY KEY,
    user_id INT,
    hotel_id INT,
    room_type TEXT,
    user_email VARCHAR(250) UNIQUE,
    enterydate DATE,
    leavingdate DATE,
    status TEXT
);