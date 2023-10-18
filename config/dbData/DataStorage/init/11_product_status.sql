-- create table
CREATE TABLE IF NOT EXISTS public.product_status
(
    id         UUID UNIQUE DEFAULT uuid_generate_v4() NOT NULL,
    name       VARCHAR(255)                           NULL,
    "order"    INTEGER                                NULL,
    active     BOOLEAN DEFAULT TRUE                   NOT NULL,

    created_at TIMESTAMP   DEFAULT NOW()              NOT NULL,
    created_by UUID        DEFAULT NULL REFERENCES public.user (id),
    updated_at TIMESTAMP   DEFAULT NOW()              NOT NULL,
    updated_by UUID        DEFAULT NULL REFERENCES public.user (id),

    CONSTRAINT product_status_pkey PRIMARY KEY (id)
);

-- ownership, index and comment
ALTER TABLE public.product_status
    OWNER TO postgres;
CREATE UNIQUE INDEX product_status_id ON public.product_status (id);
COMMENT ON TABLE public.product_status IS 'Product Statuses';

-- demo data
INSERT INTO public.product_status (id, name, "order")
VALUES ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11', 'Public', 1),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a12', 'Privet', 2),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a13', 'Out of stock', 3),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a14', 'Discontinued', 4),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a15', 'Archived', 5);