CREATE TABLE guests (
                        id              UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                        booking_id      UUID NOT NULL REFERENCES bookings(id) ON DELETE CASCADE,
                        name            TEXT NOT NULL,
                        surname         TEXT NOT NULL,
                        id_type         TEXT NOT NULL CHECK (id_type IN ('dni','passport','nie')),
                        id_number       TEXT NOT NULL,
                        support_number  TEXT NOT NULL,
                        nationality     TEXT NOT NULL,     -- ISO 3166-1 alpha-2
                        birth_date      DATE NOT NULL,
                        residence       TEXT NOT NULL,
                        signature_url   TEXT,              -- S3 key
                        registered_at   TIMESTAMPTZ,       -- set when transmitted to SES.HOSPEDAJES
                        created_at      TIMESTAMPTZ NOT NULL DEFAULT NOW()
);