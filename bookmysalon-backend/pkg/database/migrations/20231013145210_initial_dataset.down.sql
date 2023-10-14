-- pkg/database/migrations/000002_seed_data.down.sql

-- Removing seeded data from the services table
DELETE FROM services WHERE salon_id = 1 AND name = 'Hair Cut';
DELETE FROM services WHERE salon_id = 2 AND name = 'Hair Coloring';

-- Removing seeded data from the salons table
DELETE FROM salons WHERE name = 'Elegant Salon' AND address = '123 Beauty St, City';
DELETE FROM salons WHERE name = 'Trendy Cuts' AND address = '456 Barber Blvd, City';

-- Removing seeded data from the users table
DELETE FROM users WHERE username = 'alice' AND email = 'alice@email.com';
DELETE FROM users WHERE username = 'bob' AND email = 'bob@email.com';

-- ... you can continue this pattern for the other tables as needed ...
