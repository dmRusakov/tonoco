DROP TABLE IF EXISTS public.shop CASCADE;

-- create table
CREATE TABLE IF NOT EXISTS public.shop
(
    id                UUID UNIQUE  DEFAULT uuid_generate_v4() NOT NULL,
    name              UUID         DEFAULT NULL,
    seo_title         UUID         DEFAULT NULL,
    short_description UUID         DEFAULT NULL,
    description       UUID         DEFAULT NULL,
    url               VARCHAR(255) DEFAULT NULL,

    image_id          UUID         DEFAULT NULL,
    hover_image_id     UUID         DEFAULT NULL,

    page              integer      DEFAULT 1,
    per_page          integer      DEFAULT 18,

    sort_order        INTEGER      DEFAULT NULL,
    active            BOOLEAN      DEFAULT TRUE,
    prime             BOOLEAN      DEFAULT FALSE,

    created_at        TIMESTAMP    DEFAULT NOW(),
    created_by        UUID         DEFAULT NULL,
    updated_at        TIMESTAMP    DEFAULT NOW(),
    updated_by        UUID         DEFAULT NULL,

    CONSTRAINT shop_pkey PRIMARY KEY (id)
);

-- ownership and index
ALTER TABLE public.shop OWNER TO postgres;
CREATE INDEX shop_id ON public.shop (id);
CREATE INDEX shop_url ON public.shop (url);
CREATE INDEX shop_sort_order ON public.shop (sort_order);
CREATE INDEX shop_updated_at ON public.shop (updated_at);

-- add comments
COMMENT ON TABLE public.shop IS 'Store pages table';
COMMENT ON COLUMN public.shop.id IS 'Unique identifier';
COMMENT ON COLUMN public.shop.name IS 'Name of the shop page';
COMMENT ON COLUMN public.shop.seo_title IS 'SEO title of the shop page';
COMMENT ON COLUMN public.shop.short_description IS 'Short description of the shop page';
COMMENT ON COLUMN public.shop.description IS 'Description of the shop page';
COMMENT ON COLUMN public.shop.url IS 'URL of the shop page';
COMMENT ON COLUMN public.shop.image_id IS 'Image id of the shop page';
COMMENT ON COLUMN public.shop.hover_image_id IS 'Hover image id of the shop page';
COMMENT ON COLUMN public.shop.page IS 'Page number';
COMMENT ON COLUMN public.shop.per_page IS 'Number of items per page';
COMMENT ON COLUMN public.shop.sort_order IS 'Sort order of the shop page';
COMMENT ON COLUMN public.shop.active IS 'Active status of the shop page';
COMMENT ON COLUMN public.shop.prime IS 'Prime status of the shop page';
COMMENT ON COLUMN public.shop.created_at IS 'Creation date of the shop page';
COMMENT ON COLUMN public.shop.created_by IS 'Creator of the shop page';
COMMENT ON COLUMN public.shop.updated_at IS 'Last update date of the shop page';
COMMENT ON COLUMN public.shop.updated_by IS 'Last updater of the shop page';

-- auto set sort_order column
CREATE OR REPLACE TRIGGER store_pages_order
    BEFORE INSERT
    ON public.shop
    FOR EACH ROW
    EXECUTE FUNCTION set_order_column_universal();

-- auto update updated_at
CREATE OR REPLACE TRIGGER store_pages_updated_at
    BEFORE UPDATE
    ON public.shop
    FOR EACH ROW
EXECUTE FUNCTION update_update_at_column();

-- auto set created_by
CREATE OR REPLACE TRIGGER store_pages_created_by
    BEFORE INSERT
    ON public.shop
    FOR EACH ROW
EXECUTE FUNCTION set_created_by_if_null();

-- auto set updated_by
CREATE OR REPLACE TRIGGER store_pages_updated_by
    BEFORE INSERT
    ON public.shop
    FOR EACH ROW
EXECUTE FUNCTION set_updated_by_if_null();

-- insert data
INSERT INTO public.shop (id, name, seo_title, short_description, description, url, image_id, hover_image_id, page, per_page, sort_order, active, created_at, created_by, updated_at, updated_by) VALUES ('79997faf-fd52-4bc9-bda2-696b20d29973', 'ba30af8e-ad87-4f75-b81c-829396b739e3', '3184cf72-a92c-463e-9517-939712abfeac', '03fe5a5f-d40c-4543-8e36-26bcf29cd563', 'fe8d1d54-2286-4f9d-9c65-bbc5728d30fc', 'range-hoods', null, null, 1, 18, 1, true, '2024-11-27 03:12:00.529855', '0e95efda-f9e2-4fac-8184-3ce2e8b7e0e1', '2024-11-27 03:13:28.796350', '0e95efda-f9e2-4fac-8184-3ce2e8b7e0e1');

-- get data
select * from public.shop;
