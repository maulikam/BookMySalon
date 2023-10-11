-- User Service
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL, -- hashed password
    email VARCHAR(255) UNIQUE NOT NULL,
    profile_image VARCHAR(255),
    date_joined TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    last_login TIMESTAMP
);

-- Salon Service
CREATE TABLE salons (
    salon_id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    address TEXT NOT NULL,
    contact_details TEXT,
    photos TEXT, -- URLs to photos
    average_rating FLOAT DEFAULT 0
);

CREATE TABLE services (
    service_id SERIAL PRIMARY KEY,
    salon_id INTEGER REFERENCES salons(salon_id), -- link service to salon
    name VARCHAR(255) NOT NULL,
    description TEXT,
    duration INTERVAL,
    price DECIMAL(10, 2) NOT NULL
);

-- Appointment Service
CREATE TABLE appointments (
    appointment_id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES users(id),
    salon_id INTEGER REFERENCES salons(salon_id),
    service_id INTEGER REFERENCES services(service_id),
    date_time TIMESTAMP NOT NULL,
    status VARCHAR(50) CHECK (status IN ('booked', 'cancelled', 'completed')),
    notification_settings TEXT
);

-- Availability Service
CREATE TABLE availabilities (
    availability_id SERIAL PRIMARY KEY,
    salon_id INTEGER REFERENCES salons(salon_id),
    service_id INTEGER REFERENCES services(service_id),
    start_date_time TIMESTAMP NOT NULL,
    end_date_time TIMESTAMP NOT NULL,
    status VARCHAR(50) CHECK (status IN ('available', 'booked'))
);

-- Payment Service
CREATE TABLE transactions (
    transaction_id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES users(id),
    amount DECIMAL(10, 2) NOT NULL,
    date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    status VARCHAR(50) CHECK (status IN ('pending', 'successful', 'failed')),
    payment_method VARCHAR(255) -- e.g., credit card, PayPal
);

CREATE TABLE invoices (
    invoice_id SERIAL PRIMARY KEY,
    transaction_id INTEGER REFERENCES transactions(transaction_id),
    details TEXT,
    date_issued TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE promotions (
    promotion_id SERIAL PRIMARY KEY,
    description TEXT,
    discount_amount DECIMAL(10, 2),
    valid_from TIMESTAMP,
    valid_to TIMESTAMP
);

-- Review and Rating Service
CREATE TABLE reviews (
    review_id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES users(id),
    salon_id INTEGER REFERENCES salons(salon_id),
    rating INTEGER CHECK (rating BETWEEN 1 AND 5),
    comment TEXT,
    date_posted TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
