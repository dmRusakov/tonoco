DROP TABLE IF EXISTS public.store CASCADE;

-- create table
CREATE TABLE IF NOT EXISTS public.store
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
    currency_url   VARCHAR(10)  DEFAULT NULL,

    created_at     TIMESTAMP    DEFAULT NOW()              NOT NULL,
    created_by     UUID         DEFAULT NULL,
    updated_at     TIMESTAMP    DEFAULT NOW()              NOT NULL,
    updated_by     UUID         DEFAULT NULL,

    CONSTRAINT store_pkey PRIMARY KEY (id)
);

-- ownership and index
ALTER TABLE public.store OWNER TO postgres;
CREATE INDEX IF NOT EXISTS store_id ON public.store USING btree (id);
CREATE INDEX IF NOT EXISTS store_name ON public.store USING btree (name);
CREATE INDEX IF NOT EXISTS store_url ON public.store USING btree (url);
CREATE INDEX IF NOT EXISTS store_abbreviation ON public.store USING btree (abbreviation);
CREATE INDEX IF NOT EXISTS store_sort_order ON public.store USING btree (sort_order);
CREATE INDEX IF NOT EXISTS store_updated_at ON public.store USING btree (updated_at);

-- comment on table
COMMENT ON TABLE public.store IS 'Reference table for store';
COMMENT ON COLUMN public.store.id IS 'Unique identifier for store';
COMMENT ON COLUMN public.store.name IS 'Name of store';
COMMENT ON COLUMN public.store.url IS 'URL of store';
COMMENT ON COLUMN public.store.abbreviation IS 'Abbreviation of store';
COMMENT ON COLUMN public.store.sort_order IS 'Sort order of store';
COMMENT ON COLUMN public.store.active IS 'Active status of store';
COMMENT ON COLUMN public.store.address_line_1 IS 'Address line 1 of store';
COMMENT ON COLUMN public.store.address_line_2 IS 'Address line 2 of store';
COMMENT ON COLUMN public.store.city IS 'City of store';
COMMENT ON COLUMN public.store.state IS 'State of store';
COMMENT ON COLUMN public.store.zip IS 'Zip code of store';
COMMENT ON COLUMN public.store.country IS 'Country of store';
COMMENT ON COLUMN public.store.web_site IS 'Website of store';
COMMENT ON COLUMN public.store.phone IS 'Phone number of store';
COMMENT ON COLUMN public.store.email IS 'Email of store';
COMMENT ON COLUMN public.store.created_at IS 'Creation time of store';
COMMENT ON COLUMN public.store.created_by IS 'Creator of store';
COMMENT ON COLUMN public.store.updated_at IS 'Update time of store';
COMMENT ON COLUMN public.store.updated_by IS 'Updater of store';

-- auto set sort_order column
CREATE OR REPLACE TRIGGER store_order
    BEFORE INSERT
    ON public.store
    FOR EACH ROW
    EXECUTE FUNCTION set_order_column_universal();

-- auto update updated_at
CREATE OR REPLACE TRIGGER store_updated_at
    BEFORE UPDATE
    ON public.store
    FOR EACH ROW
EXECUTE FUNCTION update_update_at_column();

-- auto set created_by
CREATE OR REPLACE TRIGGER store_created_by
    BEFORE INSERT
    ON public.store
    FOR EACH ROW
EXECUTE FUNCTION set_created_by_if_null();

-- auto set updated_by
CREATE OR REPLACE TRIGGER store_updated_by
    BEFORE INSERT
    ON public.store
    FOR EACH ROW
EXECUTE FUNCTION set_updated_by_if_null();

-- default data
INSERT INTO public.store (name, url, abbreviation, sort_order, address_line_1, city, state, zip, country, web_site, phone, email, currency_url)
VALUES ('Futuro Factory Direct', 'fll', 'FLL', 1, '2201 John P Lyons Lane', 'Hallandale', 'FL', '33009', 'USA', 'https://www.futurofuturo.com', '800-230-3565', 'general@futurofuturo.com', 'usd');

-- get data
SELECT * FROM public.store;