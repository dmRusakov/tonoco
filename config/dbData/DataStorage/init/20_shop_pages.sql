DROP TABLE IF EXISTS public.store_pages CASCADE;

-- create table
CREATE TABLE IF NOT EXISTS public.store_pages
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

    created_at        TIMESTAMP    DEFAULT NOW(),
    created_by        UUID         DEFAULT NULL,
    updated_at        TIMESTAMP    DEFAULT NOW(),
    updated_by        UUID         DEFAULT NULL,

    CONSTRAINT store_pages_pkey PRIMARY KEY (id)
);

-- ownership and index
ALTER TABLE public.store_pages OWNER TO postgres;
CREATE INDEX store_pages_id ON public.store_pages (id);
CREATE INDEX store_pages_url ON public.store_pages (url);
CREATE INDEX store_pages_sort_order ON public.store_pages (sort_order);
CREATE INDEX store_pages_updated_at ON public.store_pages (updated_at);

-- add comments
COMMENT ON TABLE public.store_pages IS 'Store pages table';
COMMENT ON COLUMN public.store_pages.id IS 'Unique identifier';
COMMENT ON COLUMN public.store_pages.name IS 'Name of the store page';
COMMENT ON COLUMN public.store_pages.seo_title IS 'SEO title of the store page';
COMMENT ON COLUMN public.store_pages.short_description IS 'Short description of the store page';
COMMENT ON COLUMN public.store_pages.description IS 'Description of the store page';
COMMENT ON COLUMN public.store_pages.url IS 'URL of the store page';
COMMENT ON COLUMN public.store_pages.image_id IS 'Image id of the store page';
COMMENT ON COLUMN public.store_pages.hover_image_id IS 'Hover image id of the store page';
COMMENT ON COLUMN public.store_pages.page IS 'Page number';
COMMENT ON COLUMN public.store_pages.per_page IS 'Number of items per page';
COMMENT ON COLUMN public.store_pages.sort_order IS 'Sort order of the store page';
COMMENT ON COLUMN public.store_pages.active IS 'Active status of the store page';
COMMENT ON COLUMN public.store_pages.created_at IS 'Creation date of the store page';
COMMENT ON COLUMN public.store_pages.created_by IS 'Creator of the store page';
COMMENT ON COLUMN public.store_pages.updated_at IS 'Last update date of the store page';
COMMENT ON COLUMN public.store_pages.updated_by IS 'Last updater of the store page';

-- auto set sort_order column
CREATE OR REPLACE TRIGGER store_pages_order
    BEFORE INSERT
    ON public.store_pages
    FOR EACH ROW
    EXECUTE FUNCTION set_order_column_universal();

-- auto update updated_at
CREATE OR REPLACE TRIGGER store_pages_updated_at
    BEFORE UPDATE
    ON public.store_pages
    FOR EACH ROW
EXECUTE FUNCTION update_update_at_column();

-- auto set created_by
CREATE OR REPLACE TRIGGER store_pages_created_by
    BEFORE INSERT
    ON public.store_pages
    FOR EACH ROW
EXECUTE FUNCTION set_created_by_if_null();

-- auto set updated_by
CREATE OR REPLACE TRIGGER store_pages_updated_by
    BEFORE INSERT
    ON public.store_pages
    FOR EACH ROW
EXECUTE FUNCTION set_updated_by_if_null();

-- get data
select * from public.store_pages;
