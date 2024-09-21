-- drop table if exists
DROP TABLE IF EXISTS public.image CASCADE;

-- create table
CREATE TABLE IF NOT EXISTS public.image
(
    id             UUID UNIQUE   DEFAULT uuid_generate_v4(),
    title          VARCHAR(255)  DEFAULT NULL,
    alt_text       VARCHAR(4000) DEFAULT NULL,

    origin_path    VARCHAR(255)  DEFAULT NULL,
    full_path      VARCHAR(255)  DEFAULT NULL,
    large_path     VARCHAR(255)  DEFAULT NULL,
    medium_path    VARCHAR(255)  DEFAULT NULL,
    grid_path      VARCHAR(255)  DEFAULT NULL,
    thumbnail_path VARCHAR(255)  DEFAULT NULL,

    sort_order     INTEGER       DEFAULT NULL,
    is_webp        BOOLEAN       DEFAULT FALSE,
    image_type     VARCHAR(255)  DEFAULT NULL,

    created_at     TIMESTAMP     DEFAULT NOW() NOT NULL,
    created_by     UUID          DEFAULT NULL,
    updated_at     TIMESTAMP     DEFAULT NOW() NOT NULL,
    updated_by     UUID          DEFAULT NULL,

    CONSTRAINT image_pkey PRIMARY KEY (id)
);

-- ownership and index
ALTER TABLE public.image OWNER TO postgres;
CREATE INDEX image_id ON public.image USING btree (id);
CREATE INDEX image_type ON public.image USING btree (image_type);

-- comment on table
COMMENT ON TABLE public.image IS 'Reference table for price type';
COMMENT ON COLUMN public.image.id IS 'Unique identifier for price type';
COMMENT ON COLUMN public.image.origin_path IS 'Origin path of image';
COMMENT ON COLUMN public.image.full_path IS 'Full path of image 2000x2000';
COMMENT ON COLUMN public.image.large_path IS 'Large path of image 1500x1500';
COMMENT ON COLUMN public.image.medium_path IS 'Medium path of image 800x800';
COMMENT ON COLUMN public.image.grid_path IS 'Medium path of image 400x400';
COMMENT ON COLUMN public.image.thumbnail_path IS 'Thumbnail path of image 200x200';

COMMENT ON COLUMN public.image.created_at IS 'Creation time of price type';
COMMENT ON COLUMN public.image.created_by IS 'Creator of price type';
COMMENT ON COLUMN public.image.updated_at IS 'Update time of price type';
COMMENT ON COLUMN public.image.updated_by IS 'Updater of price type';

-- auto update updated_at
CREATE OR REPLACE TRIGGER image_updated_at
    BEFORE UPDATE OR INSERT
    ON public.image
    FOR EACH ROW
EXECUTE FUNCTION update_update_at_column();

-- auto set created_by
CREATE OR REPLACE TRIGGER image_created_by
    BEFORE INSERT
    ON public.image
    FOR EACH ROW
EXECUTE FUNCTION set_created_by_if_null();

-- auto set updated_by
CREATE OR REPLACE TRIGGER image_updated_by
    BEFORE INSERT OR UPDATE
    ON public.image
    FOR EACH ROW
EXECUTE FUNCTION set_updated_by_if_null();

select * from public.image;

-- SELECT
--     img.ID as sort_order,
--     replace(replace(img.guid, 'https://futurofuturo.com/wp-content/uploads/', ''), 'http://futurofuturo.com/wp-content/uploads/', '') AS origin_path,
--     img.post_title AS title,
--     img.post_content AS alt_text,
--     replace(img.post_mime_type, 'image/', '') AS image_type
-- FROM wp_posts img
-- WHERE img.post_mime_type LIKE 'image/%';
