CREATE TABLE IF NOT EXISTS reagent (
    id UUID PRIMARY KEY,
    guid UUID NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    deleted_at TIMESTAMP NULL
);
