-- create table
CREATE TABLE IF NOT EXISTS public.warehouse
(
    id             UUID UNIQUE  DEFAULT uuid_generate_v4() NOT NULL,
    name           VARCHAR(255) DEFAULT NULL,
    url            VARCHAR(255) UNIQUE                     NOT NULL,
    abbreviation   VARCHAR(10)  DEFAULT NULL,
    sort_order     INTEGER      DEFAULT NULL,
    active         BOOLEAN      DEFAULT TRUE,
    address_line_1 VARCHAR(255) DEFAULT NULL,
    address_line_2 VARCHAR(255) DEFAULT NULL,
    city           VARCHAR(255) DEFAULT NULL,
    state          VARCHAR(255) DEFAULT NULL,
    zip            VARCHAR(255) DEFAULT NULL,
    country        VARCHAR(255) DEFAULT NULL,
    web_site       VARCHAR(255) DEFAULT NULL,
    phone          VARCHAR(255) DEFAULT NULL,
    email          VARCHAR(255) DEFAULT NULL,

    created_at     TIMESTAMP    DEFAULT NOW()              NOT NULL,
    created_by     UUID         DEFAULT NULL,
    updated_at     TIMESTAMP    DEFAULT NOW()              NOT NULL,
    updated_by     UUID         DEFAULT NULL,

    CONSTRAINT warehouse_pkey PRIMARY KEY (id)
);

-- ownership and index
ALTER TABLE public.warehouse OWNER TO postgres;
CREATE INDEX IF NOT EXISTS warehouse_id ON public.warehouse USING btree (id);
CREATE INDEX IF NOT EXISTS warehouse_name ON public.warehouse USING btree (name);
CREATE INDEX IF NOT EXISTS warehouse_url ON public.warehouse USING btree (url);
CREATE INDEX IF NOT EXISTS warehouse_abbreviation ON public.warehouse USING btree (abbreviation);
CREATE INDEX IF NOT EXISTS warehouse_sort_order ON public.warehouse USING btree (sort_order);
CREATE INDEX IF NOT EXISTS warehouse_updated_at ON public.warehouse USING btree (updated_at);

-- comment on table
COMMENT ON TABLE public.warehouse IS 'Reference table for warehouse';
COMMENT ON COLUMN public.warehouse.id IS 'Unique identifier for warehouse';
COMMENT ON COLUMN public.warehouse.name IS 'Name of warehouse';
COMMENT ON COLUMN public.warehouse.url IS 'URL of warehouse';
COMMENT ON COLUMN public.warehouse.abbreviation IS 'Abbreviation of warehouse';
COMMENT ON COLUMN public.warehouse.sort_order IS 'Sort order of warehouse';
COMMENT ON COLUMN public.warehouse.active IS 'Active status of warehouse';
COMMENT ON COLUMN public.warehouse.address_line_1 IS 'Address line 1 of warehouse';
COMMENT ON COLUMN public.warehouse.address_line_2 IS 'Address line 2 of warehouse';
COMMENT ON COLUMN public.warehouse.city IS 'City of warehouse';
COMMENT ON COLUMN public.warehouse.state IS 'State of warehouse';
COMMENT ON COLUMN public.warehouse.zip IS 'Zip code of warehouse';
COMMENT ON COLUMN public.warehouse.country IS 'Country of warehouse';
COMMENT ON COLUMN public.warehouse.web_site IS 'Website of warehouse';
COMMENT ON COLUMN public.warehouse.phone IS 'Phone number of warehouse';
COMMENT ON COLUMN public.warehouse.email IS 'Email of warehouse';
COMMENT ON COLUMN public.warehouse.created_at IS 'Creation time of warehouse';
COMMENT ON COLUMN public.warehouse.created_by IS 'Creator of warehouse';
COMMENT ON COLUMN public.warehouse.updated_at IS 'Update time of warehouse';
COMMENT ON COLUMN public.warehouse.updated_by IS 'Updater of warehouse';

-- auto update updated_at
CREATE TRIGGER warehouse_updated_at
    BEFORE UPDATE
    ON public.warehouse
    FOR EACH ROW
    EXECUTE FUNCTION update_update_at_column();

-- auto set sort_order column
CREATE TRIGGER warehouse_order
    BEFORE INSERT
    ON public.warehouse
    FOR EACH ROW
    EXECUTE FUNCTION set_order_column_universal();

-- default data
INSERT INTO public.warehouse (name, url, abbreviation, sort_order, address_line_1, city, state, zip, country, web_site, phone, email)
VALUES ('Futuro Factory Direct', 'fll', 'FLL', 1, '2201 John P Lyons Lane', 'Hallandale', 'FL', '33009', 'USA', 'https://www.futurofuturo.com', '800-230-3565', 'general@futurofuturo.com');

-- get data
SELECT * FROM public.warehouse;