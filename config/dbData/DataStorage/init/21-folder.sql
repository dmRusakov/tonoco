-- create table
CREATE TABLE IF NOT EXISTS public.folder
(
    folder_id  uuid unique DEFAULT uuid_generate_v4(),
    name       varchar(255) not null,
    path       varchar(255) not null,
    parent_id  uuid        default null REFERENCES public.folder (folder_id),

    created_at TIMESTAMP   DEFAULT NOW(),
    created_by UUID        DEFAULT NULL REFERENCES public.user (id),
    updated_at TIMESTAMP   DEFAULT NOW(),
    updated_by UUID        DEFAULT NULL REFERENCES public.user (id),

    CONSTRAINT folder_pkey PRIMARY KEY (folder_id)
);

-- -- index, constraint and ownership
CREATE UNIQUE INDEX folder_id ON public.folder(folder_id);
ALTER TABLE public.folder OWNER TO postgres;
COMMENT ON TABLE public.folder IS 'File Folders table';

-- auto update updated_at
CREATE TRIGGER update_folder_updated_at BEFORE UPDATE ON public.folder FOR EACH ROW EXECUTE PROCEDURE update_update_at_column();

-- -- default data
INSERT INTO public.folder (folder_id, name, path, parent_id)
VALUES ('c475a6f3-55ad-4641-8caa-a76bfae13fbb', 'root', '/assets', null),
      ('0e22bcfd-4b38-4f2f-8d46-d71f10b04ac3', 'images', '/images', 'c475a6f3-55ad-4641-8caa-a76bfae13fbb');

-- get data
select * from public.folder;
