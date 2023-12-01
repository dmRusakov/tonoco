-- create table
CREATE TABLE IF NOT EXISTS public.product_status
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

    CONSTRAINT product_status_pkey PRIMARY KEY (id)
);

-- ownership, index and comment
ALTER TABLE public.product_status
    OWNER TO postgres;
CREATE UNIQUE INDEX product_status_id ON public.product_status (id);
CREATE UNIQUE INDEX product_status_slug ON public.product_status (slug);
COMMENT ON TABLE public.product_status IS 'Product Statuses';

-- auto update updated_at
CREATE TRIGGER product_status_updated_at
    BEFORE UPDATE
    ON public.user
    FOR EACH ROW
EXECUTE FUNCTION update_update_at_column();

-- auto set "order" column
CREATE OR REPLACE FUNCTION set_order_column_to_product_status()
    RETURNS TRIGGER AS $$
BEGIN
    NEW.order = (SELECT COALESCE(MAX("order"), 0) + 1 FROM product_status);
    RETURN NEW;
END;
$$ language 'plpgsql';

CREATE TRIGGER product_status_order
    BEFORE INSERT
    ON public.product_status
    FOR EACH ROW EXECUTE FUNCTION set_order_column_to_product_status();

-- demo data
INSERT INTO public.product_status (id, name, slug)
VALUES ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11', 'Public', 'public'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a12', 'Privet', 'private'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a13', 'Out of stock', 'out-of-stock'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a14', 'Discontinued', 'discontinued'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a15', 'Archived', 'archived');

-- get data
select *
from public.product_status;