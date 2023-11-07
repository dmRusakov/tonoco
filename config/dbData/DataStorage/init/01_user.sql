-- create table
CREATE TABLE IF NOT EXISTS public.user
(
    id         UUID UNIQUE DEFAULT uuid_generate_v4(),
    email      VARCHAR(255) NOT NULL UNIQUE,
    first_name VARCHAR(255),
    last_name  VARCHAR(255),
    password   VARCHAR(60),
    active     BOOLEAN     DEFAULT TRUE,

    created_at TIMESTAMP   DEFAULT NOW() NOT NULL,
    created_by UUID        DEFAULT NULL REFERENCES public.user (id),
    updated_at TIMESTAMP   DEFAULT NOW() NOT NULL,
    updated_by UUID        DEFAULT NULL REFERENCES public.user (id),

    CONSTRAINT user_pkey PRIMARY KEY (id)
);

-- index, constraint and ownership
ALTER TABLE public.user
    OWNER TO postgres;
CREATE UNIQUE INDEX user_id ON public.user (id);
CREATE UNIQUE INDEX email ON public.user (email);

-- auto update updated_at
CREATE TRIGGER user_set_updated_at
    BEFORE UPDATE ON public.user
    FOR EACH ROW
EXECUTE FUNCTION update_update_at_column();

-- insert test data
INSERT INTO public.user (id, email, first_name, last_name, password, active)
VALUES ('0e95efda-f9e2-4fac-8184-3ce2e8b7e0e1', 'mike@yaronia.com', 'Tonoco', 'Ross',
        '$2a$12$1zGLuYDDNvATh4RA4avbKuheAMpb1svexSzrQm7up.bnpwQHs0jNe', true)
ON CONFLICT (email) DO UPDATE
    SET first_name = EXCLUDED.first_name,
        last_name  = EXCLUDED.last_name,
        password   = EXCLUDED.password,
        active     = EXCLUDED.active;

-- get data
select * from public.user;