CREATE TABLE IF NOT EXISTS public.specification_value
(
    id               UUID UNIQUE DEFAULT uuid_generate_v4(),
    specification_id UUID         NOT NULL REFERENCES public.specification (id),
    name             VARCHAR(255) NOT NULL,
    slug             VARCHAR(255) NOT NULL,
    active           BOOLEAN     DEFAULT TRUE,
    "order"          INTEGER     DEFAULT 9999,

    created_at       TIMESTAMP   DEFAULT NOW(),
    created_by       UUID        DEFAULT NULL REFERENCES public.user (id),
    updated_at       TIMESTAMP   DEFAULT NOW(),
    updated_by       UUID        DEFAULT NULL REFERENCES public.user (id),

    CONSTRAINT specification_value_pkey PRIMARY KEY (id)
);

CREATE UNIQUE INDEX specification_value_id ON public.specification_value (id);
CREATE UNIQUE INDEX specification_value_slug ON public.specification_value (slug);
ALTER TABLE public.specification_value
    OWNER TO postgres;
COMMENT ON TABLE public.specification_value IS 'Product Specification values';

-- insert data
INSERT INTO public.specification_value (id, specification_id, name, slug, "order")
VALUES ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11', 'a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11', 'Inox', 'inox', 1),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a12', 'a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11', 'Black', 'black', 2),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a13', 'a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11', 'White', 'white', 3),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a14', 'a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a12', '30', '30', 1),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a15', 'a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a12', '36', '36', 2),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a16', 'a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a12', '48', '48', 3),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a17', 'a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a12', '60', '60', 1),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a18', 'a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a12', '70', '70', 2),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a19', 'a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a13', '80', '80', 3);