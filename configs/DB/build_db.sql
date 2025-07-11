-- Active: 1751385645775@@127.0.0.1@5432@startchat

-- ==============================================================
-- =====================                    =====================
-- =====================    OTP services    ===================== 
-- =====================                    =====================
-- ==============================================================
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TYPE otp_purpose as ENUM ('login', 'register','forgot_password','delete_account'); -- you can add something for purpose otp

CREATE TABLE IF NOT EXISTS otp_requests (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),

    email TEXT,
    phone TEXT,

    purpose otp_purpose NOT NULL DEFAULT 'login', -- tambahkan tujuan OTP (optional tapi bagus)
    otp_code TEXT NOT NULL,

    expires_at TIMESTAMPTZ NOT NULL,
    verified BOOLEAN NOT NULL DEFAULT FALSE,

    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    verified_at TIMESTAMPTZ,
    
    CHECK (
        (email IS NOT NULL AND phone IS NULL)
        OR
        (email IS NULL AND phone IS NOT NULL)
    )
);



-- sample insert
INSERT INTO otp_requests (email, otp_code, purpose, expires_at) VALUES ('user@example.com', 'bcrypt-otp','register', NOW() + INTERVAL '5 minutes');

-- sample delete
-- DELETE from otp_requests WHERE email = 'user@example.com';

-- check rate limit insert
-- select count(*) from otp_requests where email = 'user@example.com' and created_at > now() - INTERVAL '5 minute';

-- sample verify true
-- UPDATE public.otp_requests set verified = TRUE, verified_at = NOW() WHERE email = 'user@example.com' and otp_code = '123456';
-- sample get otp
-- SELECT * from otp_requests where email = 'hend41234@proton.me' and purpose = 'register' ORDER BY created_at DESC limit 1;


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
    token TEXT NOT NULL UNIQUE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    expires_at TIMESTAMPTZ NOT NULL,
    verified BOOLEAN NOT NULL DEFAULT FALSE
);
-- create index
CREATE INDEX IF NOT EXISTS idx_pending_verified_expires
ON pending_users (verified, expires_at);

-- sample insert
-- INSERT INTO pending_users (email, ip_address, token, expires_at) VALUES ('hend@testasd.com', '12.123.123.123', 'hashtoken', NOW() + INTERVAL '5 minutes');
-- sample get pending_users
-- select * from pending_users where email = 'hend@testasd.com' and ip_address = '12.123.123.123' and token = 'hashtoken';
-- sample delete
-- DELETE from pending_users where email = 'hend@testasd.com' and ip_address = '12.123.123.123' and token = 'hashtoken';
-- DELETE from pending_users where id = '635d23a9-1276-4377-ab1e-b27012d99ad4';
SELECT * from pending_users WHERE verified = TRUE AND expires_at < NOW();