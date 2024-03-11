-- create table
CREATE TABLE IF NOT EXISTS public.shipping_class
(
    id          UUID UNIQUE DEFAULT uuid_generate_v4() NOT NULL,
    name        VARCHAR(255)                           NOT NULL,
    url         VARCHAR(255)                           NOT NULL,
    sort_order  INTEGER                                NOT NULL,
    active      BOOLEAN     DEFAULT TRUE               NOT NULL,

    created_at TIMESTAMP   DEFAULT NOW()               NOT NULL,
    created_by UUID        DEFAULT '0e95efda-f9e2-4fac-8184-3ce2e8b7e0e1' REFERENCES public.user (id),
    updated_at TIMESTAMP   DEFAULT NOW()               NOT NULL,
    updated_by UUID        DEFAULT '0e95efda-f9e2-4fac-8184-3ce2e8b7e0e1' REFERENCES public.user (id),

    CONSTRAINT shipping_class_pkey PRIMARY KEY (id)
);

-- ownership, index and comment
ALTER TABLE public.shipping_class OWNER TO postgres;
CREATE INDEX shipping_class_id ON public.shipping_class (id);
CREATE INDEX shipping_class_url ON public.shipping_class (url);
CREATE INDEX shipping_class_sort_order ON public.shipping_class (sort_order);
CREATE INDEX shipping_class_updated_at ON public.shipping_class (updated_at);

-- add comments
COMMENT ON TABLE public.shipping_class IS 'Table of shipping classes.';
COMMENT ON COLUMN public.shipping_class.id IS 'Unique identifier';
COMMENT ON COLUMN public.shipping_class.name IS 'Name of the shipping class';
COMMENT ON COLUMN public.shipping_class.url IS 'URL of the shipping class';
COMMENT ON COLUMN public.shipping_class.sort_order IS 'Sort order of the shipping class';
COMMENT ON COLUMN public.shipping_class.active IS 'Active status of the shipping class';

-- auto update updated_at
CREATE TRIGGER shipping_class_updated_at
    BEFORE UPDATE
    ON public.user
    FOR EACH ROW
EXECUTE FUNCTION update_update_at_column();

-- demo data
INSERT INTO public.shipping_class (id, name, url, sort_order)
VALUES ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11', 'Freight', 'freight', 0),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a12', 'Ground', 'ground', 1),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a13', 'Ground - small', 'ground-small', 2);

-- get data
select * from public.shipping_class;