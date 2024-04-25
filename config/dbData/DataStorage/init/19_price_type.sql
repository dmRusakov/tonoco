-- create table
CREATE TABLE IF NOT EXISTS public.price_type
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

    CONSTRAINT price_type_pkey PRIMARY KEY (id)
);

-- ownership and index
ALTER TABLE public.price_type OWNER TO postgres;
CREATE INDEX price_type_id ON public.price_type USING btree (id);
CREATE INDEX price_type_name ON public.price_type USING btree (name);
CREATE INDEX price_type_url ON public.price_type USING btree (url);
CREATE INDEX price_type_sort_order ON public.price_type USING btree (sort_order);
CREATE INDEX price_type_updated_at ON public.price_type USING btree (updated_at);

-- comment on table
COMMENT ON TABLE public.price_type IS 'Reference table for price type';
COMMENT ON COLUMN public.price_type.id IS 'Unique identifier for price type';
COMMENT ON COLUMN public.price_type.name IS 'Name of price type';
COMMENT ON COLUMN public.price_type.url IS 'URL of price type';
COMMENT ON COLUMN public.price_type.sort_order IS 'Sort order of price type';
COMMENT ON COLUMN public.price_type.active IS 'Active status of price type';
COMMENT ON COLUMN public.price_type.created_at IS 'Creation time of price type';
COMMENT ON COLUMN public.price_type.created_by IS 'Creator of price type';
COMMENT ON COLUMN public.price_type.updated_at IS 'Update time of price type';
COMMENT ON COLUMN public.price_type.updated_by IS 'Updater of price type';

-- auto set sort_order column
CREATE TRIGGER price_type_order
    BEFORE INSERT
    ON public.price_type
    FOR EACH ROW
    EXECUTE FUNCTION set_order_column_universal();

-- auto update updated_at
CREATE TRIGGER price_type_updated_at
    BEFORE UPDATE
    ON public.price_type
    FOR EACH ROW
    EXECUTE FUNCTION update_update_at_column();

-- demo data
INSERT INTO public.price_type (name, url, sort_order)
VALUES ('Regular', 'regular', 1),
       ('Sale', 'sale', 2),
       ('Special', 'special', 3),
       ('Purchase/Factory', 'purchase-factory', 4);

-- get data
select * from public.price_type;