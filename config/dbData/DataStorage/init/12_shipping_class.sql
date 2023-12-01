-- create table
CREATE TABLE IF NOT EXISTS public.shipping_class
(
    id         UUID UNIQUE DEFAULT uuid_generate_v4() NOT NULL,
    name       VARCHAR(255)                           NULL,
    slug       VARCHAR(255)                           NULL,
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

-- auto update updated_at
CREATE TRIGGER shipping_class_updated_at
    BEFORE UPDATE
    ON public.user
    FOR EACH ROW
EXECUTE FUNCTION update_update_at_column();

-- auto set "order" column
CREATE OR REPLACE FUNCTION set_order_column_to_shipping_class()
    RETURNS TRIGGER AS $$
BEGIN
    NEW.order = (SELECT COALESCE(MAX("order"), 0) + 1 FROM shipping_class);
    RETURN NEW;
END;
$$ language 'plpgsql';

CREATE TRIGGER shipping_class_order
    BEFORE INSERT
    ON public.shipping_class
    FOR EACH ROW EXECUTE FUNCTION set_order_column_to_shipping_class();

-- demo data
INSERT INTO public.shipping_class (id, name, slug)
VALUES ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11', 'Freight', 'freight'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a12', 'Ground', 'ground'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a13', 'Ground - small', 'ground-small');

-- get data
select * from public.shipping_class;