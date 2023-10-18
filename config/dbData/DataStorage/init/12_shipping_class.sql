-- create table
CREATE TABLE IF NOT EXISTS public.shipping_class
(
    id         UUID UNIQUE DEFAULT uuid_generate_v4() NOT NULL,
    name       VARCHAR(255)                           NULL,
    "order"    INTEGER                                NULL,
    active     BOOLEAN     DEFAULT TRUE               NOT NULL,

    created_at TIMESTAMP   DEFAULT NOW()              NOT NULL,
    created_by UUID        DEFAULT NULL REFERENCES public.user (id),
    updated_at TIMESTAMP   DEFAULT NOW()              NOT NULL,
    updated_by UUID        DEFAULT NULL REFERENCES public.user (id),

    CONSTRAINT shipping_class_pkey PRIMARY KEY (id)
);

-- ownership, index and comment
ALTER TABLE public.shipping_class
    OWNER TO postgres;
CREATE UNIQUE INDEX shipping_class_id ON public.shipping_class (id);
COMMENT ON TABLE public.shipping_class IS 'Table of shipping classes.';

-- demo data
INSERT INTO public.shipping_class (id, name, "order")
VALUES ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11', 'Freight', 1),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a12', 'Ground', 2),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a13', 'Ground - Accessory', 3);