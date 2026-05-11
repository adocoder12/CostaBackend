CREATE TABLE cleaning_tasks (
                                id                  UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                                apartment_id        UUID NOT NULL REFERENCES apartments(id) ON DELETE CASCADE,
                                assigned_cleaner_id UUID REFERENCES cleaners(id) ON DELETE SET NULL,
                                task_type           TEXT NOT NULL DEFAULT 'cleaning'
                                    CHECK (task_type IN ('cleaning','maintenance')),
                                status              TEXT NOT NULL DEFAULT 'pending'
                                    CHECK (status IN ('pending','in_progress','done')),
                                priority            TEXT NOT NULL DEFAULT 'normal'
                                    CHECK (priority IN ('normal','urgent')),
                                notes               TEXT,
                                repair_cost         NUMERIC(10,2),  -- maintenance only
                                photo_url           TEXT,           -- S3 key
                                scheduled_for       TIMESTAMPTZ NOT NULL,
                                completed_at        TIMESTAMPTZ,
                                created_at          TIMESTAMPTZ NOT NULL DEFAULT NOW(),
                                updated_at          TIMESTAMPTZ NOT NULL DEFAULT NOW()
);