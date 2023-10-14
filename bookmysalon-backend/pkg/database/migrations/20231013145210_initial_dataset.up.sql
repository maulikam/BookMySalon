-- pkg/database/migrations/000002_seed_data.up.sql

-- Seeding for users table
INSERT INTO users (username, password, email, profile_image, date_joined, last_login)
VALUES 
    ('alice', '$2a$12$nlQQ9u6ICDv3NV1VA3XRHOrW6H7xlGojcTT.KXt/lyFg5.tlTcl5y', 'alice@email.com', 'path/to/alice_profile_image.jpg', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('bob', '$2a$12$oZDRUSMeBGKe3fRFnMhlvus3Gvm7B3tkI0teZYNWf43Rn.p7C5Flu', 'bob@email.com', 'path/to/bob_profile_image.jpg', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

-- Seeding for salons table
INSERT INTO salons (name, address, contact_details, photos, average_rating)
VALUES 
    ('Elegant Salon', '123 Beauty St, City', '123-456-7890', 'path/to/salon1_photo.jpg', 4.5),
    ('Trendy Cuts', '456 Barber Blvd, City', '987-654-3210', 'path/to/salon2_photo.jpg', 4.7);

-- Seeding for services table
-- Note: assuming salon_id's from above are 1 and 2 respectively
INSERT INTO services (salon_id, name, description, duration, price)
VALUES 
    (1, 'Hair Cut', 'A simple hair cut.', '00:30:00', 20.50),
    (2, 'Hair Coloring', 'Full hair coloring.', '01:30:00', 65.00);

-- ... you can continue this pattern for the other tables as needed ...

-- Reminder: Passwords in the users table are placeholder 'hashed_password_for_alice' etc.
-- In a real-world scenario, you'd actually hash the passwords and then insert them.
