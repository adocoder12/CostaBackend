CREATE TABLE cleaners (
                          id         UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                          name       TEXT NOT NULL,
                          phone      TEXT NOT NULL,
                          email      TEXT,
                          zone       TEXT,
                          active     BOOLEAN NOT NULL DEFAULT TRUE,
                          created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
                          updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);