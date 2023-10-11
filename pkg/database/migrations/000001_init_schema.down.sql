-- Drop tables in reverse order of their dependencies to avoid foreign key constraint violations

-- Review and Rating Service
DROP TABLE IF EXISTS reviews;

-- Payment Service
DROP TABLE IF EXISTS promotions;
DROP TABLE IF EXISTS invoices;
DROP TABLE IF EXISTS transactions;

-- Availability Service
DROP TABLE IF EXISTS availabilities;

-- Appointment Service
DROP TABLE IF EXISTS appointments;

-- Salon Service
DROP TABLE IF EXISTS services;
DROP TABLE IF EXISTS salons;

-- User Service
DROP TABLE IF EXISTS users;
