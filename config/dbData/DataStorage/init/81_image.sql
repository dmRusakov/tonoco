-- drop table if exists
DROP TABLE IF EXISTS public.image CASCADE;

-- create table
CREATE TABLE IF NOT EXISTS public.image
(
    id                  UUID UNIQUE   DEFAULT uuid_generate_v4(),

    filename            VARCHAR(255)  DEFAULT NULL,
    extension           VARCHAR(10)   DEFAULT NULL,

    is_compressed       BOOLEAN       DEFAULT FALSE,
    is_webp             BOOLEAN       DEFAULT FALSE,

    folder_id           UUID          DEFAULT NULL,
    sort_order          INTEGER       DEFAULT NULL,

    title               VARCHAR(255)  DEFAULT NULL,
    alt_text            VARCHAR(510)  DEFAULT NULL,
    copyright           VARCHAR(100)  DEFAULT NULL,
    creator             VARCHAR(100)  DEFAULT NULL,
    rating              FLOAT         DEFAULT NULL,

    origin_path         VARCHAR(255)  DEFAULT NULL,

    created_at          TIMESTAMP     DEFAULT NOW() NOT NULL,
    created_by          UUID          DEFAULT NULL,
    updated_at          TIMESTAMP     DEFAULT NOW() NOT NULL,
    updated_by          UUID          DEFAULT NULL,

    CONSTRAINT image_pkey PRIMARY KEY (id)
);

-- ownership and index
ALTER TABLE public.image OWNER TO postgres;
CREATE INDEX image_id ON public.image USING btree (id);
CREATE INDEX image_is_compressed ON public.image USING btree (is_compressed);
CREATE INDEX image_folder_id ON public.image USING btree (folder_id);

-- comment on table
COMMENT ON TABLE public.image IS 'Reference table for price type';

COMMENT ON COLUMN public.image.id IS 'Unique identifier for price type';
COMMENT ON COLUMN public.image.filename IS 'Filename of image';
COMMENT ON COLUMN public.image.extension IS 'Extension of image';

COMMENT ON COLUMN public.image.is_compressed IS 'Compressed status of image';
COMMENT ON COLUMN public.image.is_webp IS 'Webp status of image';

COMMENT ON COLUMN public.image.folder_id IS 'Folder of image';
COMMENT ON COLUMN public.image.sort_order IS 'Sort order of image';

COMMENT ON COLUMN public.image.title IS 'Title of image';
COMMENT ON COLUMN public.image.alt_text IS 'Alt text of image';
COMMENT ON COLUMN public.image.copyright IS 'Copyright of image';
COMMENT ON COLUMN public.image.creator IS 'Creator of image';
COMMENT ON COLUMN public.image.rating IS 'Rating of image';

COMMENT ON COLUMN public.image.origin_path IS 'Origin path of image';

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
--     img.post_content AS alt_text
-- FROM wp_posts img
-- WHERE img.post_mime_type LIKE 'image/%';
