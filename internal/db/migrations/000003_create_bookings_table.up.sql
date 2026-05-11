CREATE TABLE bookings (
                          id              UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                          apartment_id    UUID NOT NULL REFERENCES apartments(id) ON DELETE CASCADE,
                          guest_name      TEXT NOT NULL,
                          check_in        DATE NOT NULL,
                          check_out       DATE NOT NULL,
                          status          TEXT NOT NULL DEFAULT 'upcoming'
                              CHECK (status IN ('upcoming','active','completed','cancelled')),
                          registered      BOOLEAN NOT NULL DEFAULT FALSE,
                          guest_count     INT DEFAULT 1,
                          notes           TEXT,
                          created_at      TIMESTAMPTZ NOT NULL DEFAULT NOW(),
                          updated_at      TIMESTAMPTZ NOT NULL DEFAULT NOW()
);