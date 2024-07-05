-- create table
CREATE TABLE IF NOT EXISTS public.folder
(
    id         UUID UNIQUE  DEFAULT uuid_generate_v4() NOT NULL,
    name       varchar(255) DEFAULT NULL,
    url        varchar(255) DEFAULT NULL,
    parent_id  uuid         DEFAULT null,
    active     BOOLEAN      DEFAULT TRUE,
    sort_order INTEGER      DEFAULT NULL,

    created_at TIMESTAMP    DEFAULT NOW()              NOT NULL,
    created_by UUID         DEFAULT NULL,
    updated_at TIMESTAMP    DEFAULT NOW()              NOT NULL,
    updated_by UUID         DEFAULT NULL,

    CONSTRAINT folder_pkey PRIMARY KEY (id)
);

-- ownership and index
ALTER TABLE public.folder OWNER TO postgres;
CREATE INDEX folder_id ON public.folder (id);
CREATE INDEX folder_url ON public.folder (url);
CREATE INDEX folder_sort_order ON public.folder (sort_order);
CREATE INDEX folder_updated_at ON public.folder (updated_at);

-- add comments
COMMENT ON TABLE public.folder IS 'Folder table';
COMMENT ON COLUMN public.folder.id IS 'Unique identifier';
COMMENT ON COLUMN public.folder.name IS 'Name of the folder';
COMMENT ON COLUMN public.folder.url IS 'URL of the folder';
COMMENT ON COLUMN public.folder.parent_id IS 'Parent folder ID';
COMMENT ON COLUMN public.folder.active IS 'Active status of the folder';
COMMENT ON COLUMN public.folder.sort_order IS 'Sort order of the folder';

-- auto update updated_at
CREATE TRIGGER folder_updated_at
    BEFORE UPDATE
    ON public.folder
    FOR EACH ROW
    EXECUTE FUNCTION update_update_at_column();

-- auto set sort_order column
CREATE TRIGGER folder_order
    BEFORE INSERT
    ON public.folder
    FOR EACH ROW
    EXECUTE FUNCTION set_order_column_universal();

-- default data
INSERT INTO public.folder (id, name, url, parent_id)
VALUES ('c475a6f3-55ad-4641-8caa-a76bfae13fb0', 'root', '/assets', null),
       ('c475a6f3-55ad-4641-8caa-a76bfae13fb1', 'images', '/images', 'c475a6f3-55ad-4641-8caa-a76bfae13fb0'),
       ('c475a6f3-55ad-4641-8caa-a76bfae13fb2', 'videos', '/videos', 'c475a6f3-55ad-4641-8caa-a76bfae13fb0'),
       ('c475a6f3-55ad-4641-8caa-a76bfae13fb3', 'documents', '/documents', 'c475a6f3-55ad-4641-8caa-a76bfae13fb0'),
       ('c475a6f3-55ad-4641-8caa-a76bfae13fb4', 'templates', '/templates', 'c475a6f3-55ad-4641-8caa-a76bfae13fb0'),
       ('c475a6f3-55ad-4641-8caa-a76bfae13fb5', 'css', '/css', 'c475a6f3-55ad-4641-8caa-a76bfae13fb0'),
       ('c475a6f3-55ad-4641-8caa-a76bfae13fb6', 'js', '/js', 'c475a6f3-55ad-4641-8caa-a76bfae13fb0');

-- get data
select * from public.folder;
