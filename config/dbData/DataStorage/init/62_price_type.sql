-- drop table if exists
DROP TABLE IF EXISTS public.price_type CASCADE;

-- create table
CREATE TABLE IF NOT EXISTS public.price_type
(
    id          UUID UNIQUE  DEFAULT uuid_generate_v4() ,
    name        VARCHAR(255) DEFAULT NULL,
    url         VARCHAR(255) DEFAULT NULL,
    sort_order  INTEGER      DEFAULT NULL,
    is_public   BOOLEAN      DEFAULT TRUE,
    active      BOOLEAN      DEFAULT TRUE,

    created_at TIMESTAMP     DEFAULT NOW(),
    created_by UUID          DEFAULT NULL,
    updated_at TIMESTAMP     DEFAULT NOW(),
    updated_by UUID          DEFAULT NULL,

    CONSTRAINT price_type_pkey PRIMARY KEY (id)
);

-- ownership and index
ALTER TABLE public.price_type OWNER TO postgres;
CREATE INDEX price_type_id ON public.price_type USING btree (id);
CREATE INDEX price_type_name ON public.price_type USING btree (name);
CREATE INDEX price_type_url ON public.price_type USING btree (url);
CREATE INDEX price_type_is_public ON public.price_type USING btree (is_public);
CREATE INDEX price_type_active ON public.price_type USING btree (active);
CREATE INDEX price_type_sort_order ON public.price_type USING btree (sort_order);
CREATE INDEX price_type_updated_at ON public.price_type USING btree (updated_at);

-- comment on table
COMMENT ON TABLE public.price_type IS 'Reference table for price type';
COMMENT ON COLUMN public.price_type.id IS 'Unique identifier for price type';
COMMENT ON COLUMN public.price_type.name IS 'Name of price type';
COMMENT ON COLUMN public.price_type.url IS 'URL of price type';
COMMENT ON COLUMN public.price_type.sort_order IS 'Sort order of price type';
COMMENT ON COLUMN public.price_type.is_public IS 'Public status of price type';
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
INSERT INTO public.price_type (id, name, url, sort_order, is_public)
VALUES ('0d1d9e2a-5cf7-4dc3-81e1-47c1dee4b2f1', 'Special', 'special', 1, true),
       ('0d1d9e2a-5cf7-4dc3-81e1-47c1dee4b2f2', 'Sale', 'sale', 2, true),
       ('0d1d9e2a-5cf7-4dc3-81e1-47c1dee4b2f3', 'Regular', 'regular', 3, true),
       ('0d1d9e2a-5cf7-4dc3-81e1-47c1dee4b2f4', 'Purchase/Factory', 'purchase-factory', 4, false);

-- get data
select * from public.price_type;
