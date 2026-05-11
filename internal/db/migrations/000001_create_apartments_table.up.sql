CREATE TABLE apartments (
                            id              UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                            name            TEXT NOT NULL,
                            address         TEXT NOT NULL,
                            status          TEXT NOT NULL DEFAULT 'clean'
                                CHECK (status IN ('clean','dirty','in_progress','maintenance','blocked')),
                            license_number  TEXT NOT NULL UNIQUE,
                            cadastral_ref   TEXT,
                            door_code       TEXT,
                            next_check_in   TIMESTAMPTZ,
                            checkout_date   DATE,
                            guest_name      TEXT,
                            notes           TEXT,
                            owner_id        UUID,              -- Supabase auth.users UUID
                            created_at      TIMESTAMPTZ NOT NULL DEFAULT NOW(),
                            updated_at      TIMESTAMPTZ NOT NULL DEFAULT NOW()
);