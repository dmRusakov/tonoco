-- create table
CREATE TABLE IF NOT EXISTS public.folder
(
    id         uuid unique DEFAULT uuid_generate_v4(),
    name       varchar(255) not null,
    slug       varchar(255) not null,
    parent_id  uuid        default null REFERENCES public.folder (id),
    active     BOOLEAN     DEFAULT TRUE,
    sort_order INTEGER     DEFAULT null,

    created_at TIMESTAMP   DEFAULT NOW(),
    created_by UUID        DEFAULT NULL REFERENCES public.user (id),
    updated_at TIMESTAMP   DEFAULT NOW(),
    updated_by UUID        DEFAULT NULL REFERENCES public.user (id),

    CONSTRAINT folder_pkey PRIMARY KEY (id)
);

-- index, constraint and ownership
CREATE INDEX folder_id ON public.folder (id);
ALTER TABLE public.folder
    OWNER TO postgres;
COMMENT ON TABLE public.folder IS 'File Folders table';

-- auto set sort_order column
CREATE OR REPLACE FUNCTION set_order_column_to_folder()
    RETURNS TRIGGER AS
$$
BEGIN
    NEW.sort_order = (SELECT COALESCE(MAX(sort_order), 0) + 1 FROM folder);
    RETURN NEW;
END;
$$ language 'plpgsql';

CREATE TRIGGER folder_order
    BEFORE INSERT
    ON public.folder
    FOR EACH ROW
EXECUTE FUNCTION set_order_column_to_folder();

-- auto update updated_at
CREATE TRIGGER update_folder_updated_at
    BEFORE UPDATE
    ON public.folder
    FOR EACH ROW
EXECUTE PROCEDURE update_update_at_column();

-- -- default data
INSERT INTO public.folder (id, name, slug, parent_id)
VALUES ('c475a6f3-55ad-4641-8caa-a76bfae13fb0', 'root', '/assets', null),
       ('c475a6f3-55ad-4641-8caa-a76bfae13fb1', 'images', '/images', 'c475a6f3-55ad-4641-8caa-a76bfae13fb0'),
       ('c475a6f3-55ad-4641-8caa-a76bfae13fb2', 'videos', '/videos', 'c475a6f3-55ad-4641-8caa-a76bfae13fb0'),
       ('c475a6f3-55ad-4641-8caa-a76bfae13fb3', 'documents', '/documents', 'c475a6f3-55ad-4641-8caa-a76bfae13fb0'),
       ('c475a6f3-55ad-4641-8caa-a76bfae13fb4', 'templates', '/templates', 'c475a6f3-55ad-4641-8caa-a76bfae13fb0'),
       ('c475a6f3-55ad-4641-8caa-a76bfae13fb5', 'css', '/css', 'c475a6f3-55ad-4641-8caa-a76bfae13fb0'),
       ('c475a6f3-55ad-4641-8caa-a76bfae13fb6', 'js', '/js', 'c475a6f3-55ad-4641-8caa-a76bfae13fb0');

-- get data
select *
from public.folder;
