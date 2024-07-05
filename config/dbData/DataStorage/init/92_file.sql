-- create table
CREATE TABLE IF NOT EXISTS public.file
(
    id         UUID UNIQUE  DEFAULT uuid_generate_v4() NOT NULL,
    name       VARCHAR(255) DEFAULT NULL,
    url        VARCHAR(255) DEFAULT NULL,
    sort_order INTEGER      DEFAULT NULL,
    active     BOOLEAN      DEFAULT TRUE,

    created_at TIMESTAMP    DEFAULT NOW(),
    created_by UUID         DEFAULT NULL,
    updated_at TIMESTAMP    DEFAULT NOW(),
    updated_by UUID         DEFAULT NULL,

    CONSTRAINT file_pkey PRIMARY KEY (id)
);

-- ownership and index
ALTER TABLE public.file OWNER TO postgres;
CREATE INDEX file_id ON public.file (id);
CREATE INDEX file_url ON public.file (url);
CREATE INDEX file_sort_order ON public.file (sort_order);
CREATE INDEX file_updated_at ON public.file (updated_at);

-- add comments
COMMENT ON TABLE public.file IS 'File table';
COMMENT ON COLUMN public.file.id IS 'Unique identifier for file';
COMMENT ON COLUMN public.file.name IS 'Name of the file';
COMMENT ON COLUMN public.file.url IS 'URL of the file';
COMMENT ON COLUMN public.file.sort_order IS 'Order of the file';
COMMENT ON COLUMN public.file.active IS 'Active status of the file';
COMMENT ON COLUMN public.file.created_at IS 'Creation date of the file';
COMMENT ON COLUMN public.file.created_by IS 'Creator of the file';
COMMENT ON COLUMN public.file.updated_at IS 'Last update date of the file';
COMMENT ON COLUMN public.file.updated_by IS 'Last updater of the file';

-- auto set sort_order column
CREATE TRIGGER file_order
    BEFORE INSERT
    ON public.file
    FOR EACH ROW
    EXECUTE FUNCTION set_order_column_universal();

-- auto update updated_at
CREATE TRIGGER file_updated_at
    BEFORE UPDATE
    ON public.file
    FOR EACH ROW
    EXECUTE FUNCTION update_update_at_column();

-- demo data

-- get data
select * from public.file;

