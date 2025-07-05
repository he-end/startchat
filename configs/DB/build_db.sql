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
-- ==============================================================
-- =====================                    =====================
-- =====================    PENDING USER    ===================== 
-- =====================                    =====================
-- ==============================================================
CREATE TABLE IF NOT EXISTS pending_users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    email TEXT NOT NULL UNIQUE,
    password TEXT NOT NULL,
    ip_address TEXT NOT NULL,
    token TEXT NOT NULL UNIQUE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    expires_at TIMESTAMPTZ NOT NULL
);
-- sample insert
INSERT INTO pending_users (email, ip_address, token, expires_at) VALUES ('hend@testasd.com', '12.123.123.123', 'hashtoken', NOW() + INTERVAL '5 minutes');
-- sample get pending_users
select * from pending_users where email = 'hend@testasd.com' and ip_address = '12.123.123.123' and token = 'hashtoken';
-- sample delete
-- DELETE from pending_users where email = 'hend@testasd.com' and ip_address = '12.123.123.123' and token = 'hashtoken';
-- DELETE from pending_users where id = '635d23a9-1276-4377-ab1e-b27012d99ad4';