CREATE TABLE IF NOT EXISTS public.specification_type
(
    id         UUID UNIQUE DEFAULT uuid_generate_v4(),
    name       VARCHAR(255) NOT NULL,
    unit       VARCHAR(255) DEFAULT NULL,
    active     BOOLEAN     DEFAULT TRUE,
    "order"    INTEGER     DEFAULT 9999,

    created_at TIMESTAMP   DEFAULT NOW(),
    created_by UUID        DEFAULT NULL REFERENCES public.user (id),
    updated_at TIMESTAMP   DEFAULT NOW(),
    updated_by UUID        DEFAULT NULL REFERENCES public.user (id),

    CONSTRAINT specification_type_pkey PRIMARY KEY (id)
);

CREATE UNIQUE INDEX specification_type_id ON public.specification_type (id);
ALTER TABLE public.specification_type
    OWNER TO postgres;
COMMENT ON TABLE public.specification_type IS 'Specification Types';

-- insert data
INSERT INTO public.specification_type (id, name, unit, "order")
VALUES ('a0eebc99-9c0b-4ef8-bb6d-7ab9bd380a11', 'Inch', '″', 1),
       ('a0eebc99-9c0b-4ef8-bb6d-7ab9bd380a12', 'Pound', 'lb', 2),
       ('a0eebc99-9c0b-4ef8-bb6d-7ab9bd380a13', 'Select', null, 3),
       ('a0eebc99-9c0b-4ef8-bb6d-7ab9bd380a14', 'Text',  null, 4);
