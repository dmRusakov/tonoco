-- create table
CREATE TABLE IF NOT EXISTS public.store
(
    id             UUID UNIQUE DEFAULT uuid_generate_v4() NOT NULL,
    name           VARCHAR(255)                           NOT NULL,
    sort_order     INTEGER                                NOT NULL,
    active         BOOLEAN     DEFAULT TRUE               NOT NULL,
    address_line_1 VARCHAR(255)                           NOT NULL,
    address_line_2 VARCHAR(255)                           DEFAULT NULL,
    city           VARCHAR(255)                           NOT NULL,
    state          VARCHAR(255)                           NOT NULL,
    zip            VARCHAR(255)                           NOT NULL,
    country        VARCHAR(255)                           NOT NULL,
    url            VARCHAR(255)                           NOT NULL,
    phone          VARCHAR(255)                           NOT NULL,
    email          VARCHAR(255)                           NOT NULL,

    created_at TIMESTAMP   DEFAULT NOW()               NOT NULL,
    created_by UUID        DEFAULT '0e95efda-f9e2-4fac-8184-3ce2e8b7e0e1' REFERENCES public.user (id),
    updated_at TIMESTAMP   DEFAULT NOW()               NOT NULL,
    updated_by UUID        DEFAULT '0e95efda-f9e2-4fac-8184-3ce2e8b7e0e1' REFERENCES public.user (id),

    CONSTRAINT store_pkey PRIMARY KEY (id)
);

-- ownership and index
ALTER TABLE public.store OWNER TO postgres;
CREATE INDEX store_id ON public.store USING btree (id);
CREATE INDEX store_name ON public.store USING btree (name);
CREATE INDEX store_url ON public.store USING btree (url);
CREATE INDEX store_sort_order ON public.store USING btree (sort_order);
CREATE INDEX store_updated_at ON public.store USING btree (updated_at);

-- comment on table
COMMENT ON TABLE public.store IS 'Reference table for store';
COMMENT ON COLUMN public.store.id IS 'Unique identifier for store';
COMMENT ON COLUMN public.store.name IS 'Name of store';
COMMENT ON COLUMN public.store.url IS 'URL of store';
COMMENT ON COLUMN public.store.sort_order IS 'Sort order of store';
COMMENT ON COLUMN public.store.active IS 'Active status of store';
COMMENT ON COLUMN public.store.address_line_1 IS 'Address line 1 of store';
COMMENT ON COLUMN public.store.address_line_2 IS 'Address line 2 of store';
COMMENT ON COLUMN public.store.city IS 'City of store';
COMMENT ON COLUMN public.store.state IS 'State of store';
COMMENT ON COLUMN public.store.zip IS 'Zip code of store';
COMMENT ON COLUMN public.store.country IS 'Country of store';
COMMENT ON COLUMN public.store.url IS 'URL of store';
COMMENT ON COLUMN public.store.phone IS 'Phone number of store';
COMMENT ON COLUMN public.store.email IS 'Email of store';
COMMENT ON COLUMN public.store.created_at IS 'Creation time of store';
COMMENT ON COLUMN public.store.created_by IS 'Creator of store';
COMMENT ON COLUMN public.store.updated_at IS 'Update time of store';
COMMENT ON COLUMN public.store.updated_by IS 'Updater of store';

-- auto update updated_at
CREATE TRIGGER store_updated_at
    BEFORE UPDATE
    ON public.store
    FOR EACH ROW
    EXECUTE FUNCTION update_update_at_column();

-- auto set sort_order column
CREATE TRIGGER store_order
    BEFORE INSERT
    ON public.store
    FOR EACH ROW
    EXECUTE FUNCTION set_order_column_universal();

-- default data
INSERT INTO public.store (name, sort_order, address_line_1, city, state, zip, country, url, phone, email)
VALUES ('Futuro Factory Direct ', 1, '2201 John P Lyons Lane', 'Hallandale', 'FL', '33009', 'USA', 'https://www.futurofuturo.com', '800-230-3565', 'general@futurofuturo.com');

-- get data
SELECT * FROM public.store;