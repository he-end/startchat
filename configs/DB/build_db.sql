-- Active: 1751385645775@@127.0.0.1@5432@startchat

-- ==============================================================
-- =====================                    =====================
-- =====================    OTP services    ===================== 
-- =====================                    =====================
-- ==============================================================
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS otp_requests (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),

    email TEXT,
    phone TEXT,

    otp_code TEXT NOT NULL,
    expires_at TIMESTAMPTZ NOT NULL,
    verified BOOLEAN NOT NULL DEFAULT FALSE,

    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    CHECK (
        (email IS NOT NULL AND phone IS NULL)
        OR
        (email IS NULL AND phone IS NOT NULL)
    )
);

-- Partial unique index untuk email
CREATE UNIQUE INDEX IF NOT EXISTS uniq_otp_email
ON otp_requests(email)
WHERE email IS NOT NULL;

-- Partial unique index untuk phone
CREATE UNIQUE INDEX IF NOT EXISTS uniq_otp_phone
ON otp_requests(phone)
WHERE phone IS NOT NULL;

-- sample insert
-- INSERT INTO otp_requests (email, otp_code, expires_at) VALUES ('user@example.com', 'bcrypt-otp', NOW() + INTERVAL '5 minutes');

-- sample delete
-- DELETE from otp_requests WHERE email = 'user@example.com';

-- verify true
-- UPDATE public.otp_requests set verified = TRUE WHERE email = 'user@example.com';




-- ==============================================================
-- =====================                    =====================
-- =====================    BASE    USER    ===================== 
-- =====================                    =====================
-- ==============================================================

CREATE TABLE IF NOT EXISTS users (
    id BIGINT PRIMARY KEY, -- ID user 10-12 digit, angka, bisa untuk login / kontak

    email TEXT UNIQUE,
    phone TEXT UNIQUE,

    password TEXT NOT NULL,
    name TEXT,

    is_verified BOOLEAN NOT NULL DEFAULT FALSE,
    is_active BOOLEAN NOT NULL DEFAULT TRUE,

    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
-- sample insert
-- INSERT INTO users (id, email, password, name) VALUES (8273920154, 'user@example.com', '<hashed_password>', 'Hendri');

-- sample delete
-- DELETE from public.users WHERE email = 'user@example.com';

-- sample unactivade account
-- UPDATE public.users set is_active = FALSE WHERE email = 'user@example.com';
