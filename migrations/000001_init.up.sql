CREATE SCHEMA IF NOT EXISTS core;

CREATE TABLE core.user(
    id            bigserial   PRIMARY KEY,
    password_hash text        NOT NULL,
    email         text        NOT NULL UNIQUE,
    role          text        NOT NULL DEFAULT 'user'
                          CHECK ( role in ('user', 'admin') ),
    created_at    timestamptz NOT NULL DEFAULT now()
);


CREATE TABLE core.resource(
    id bigserial PRIMARY KEY,
    name text NOT NULL,
    description text,
    location text,
    is_active boolean NOT NULL DEFAULT TRUE,
    created_at    timestamptz NOT NULL DEFAULT now()
);


CREATE TABLE core.booking(
    id bigserial PRIMARY KEY,
    user_id bigint NOT NULL REFERENCES core.user(id) ON DELETE  CASCADE,
    resource_id bigint NOT NULL REFERENCES core.resource(id) ON DELETE RESTRICT,
    start_time timestamptz NOT NULL ,
    end_time  timestamptz NOT NULL,
    created_at    timestamptz NOT NULL DEFAULT now(),
    CHECK ( end_time > start_time )
);


CREATE INDEX idx_booking_user_id     ON core.booking(user_id);
CREATE INDEX idx_bookings_resource_id ON core.booking(resource_id);