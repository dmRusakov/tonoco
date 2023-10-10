-- drop if needed
DROP TABLE IF EXISTS public.user CASCADE;

-- create table
CREATE TABLE IF NOT EXISTS public.user
(
    user_id    uuid    DEFAULT uuid_generate_v4(),
    email      character varying(255),
    first_name character varying(255),
    last_name  character varying(255),
    password   character varying(60),
    active     integer DEFAULT 0,
    created_at timestamp DEFAULT now(),
    updated_at timestamp DEFAULT now()
);
ALTER TABLE public.user OWNER TO postgres;
ALTER TABLE ONLY public.user ADD CONSTRAINT user_pkey PRIMARY KEY (user_id);
CREATE UNIQUE INDEX user_id ON public.user(user_id);
CREATE UNIQUE INDEX email ON public.user(email);

-- insert test data
INSERT INTO "public"."user" (user_id, email, first_name, last_name, password, active)
VALUES ('0e95efda-f9e2-4fac-8184-3ce2e8b7e0e2', 'mike@yaronia.com','Mike','Ross','$2a$12$1zGLuYDDNvATh4RA4avbKuheAMpb1svexSzrQm7up.bnpwQHs0jNe',1)
ON CONFLICT (email) DO UPDATE
    SET
        first_name = EXCLUDED.first_name,
        last_name = EXCLUDED.last_name,
        password = EXCLUDED.password,
        active = EXCLUDED.active;