-- drop if needed
DROP TABLE IF EXISTS public.folder;

-- create table
CREATE TABLE IF NOT EXISTS public.folder
(
    folder_id  uuid unique            DEFAULT uuid_generate_v4(),
    name         varchar(255)          not null,
    path         varchar(255)          not null,
    parent_id    uuid                  default null REFERENCES public.folder (folder_id),

    created_at        timestamp              default now(),
    created_by        uuid                   default null REFERENCES public.user (user_id),
    updated_at        timestamp              default now(),
    updated_by        uuid                   default null REFERENCES public.user (user_id)
);

-- index, constraint and ownership
ALTER TABLE ONLY public.folder ADD CONSTRAINT folder_pkey PRIMARY KEY (folder_id);
CREATE UNIQUE INDEX folder_id ON public.folder(folder_id);
CREATE UNIQUE INDEX parent_id ON public.folder(parent_id);
ALTER TABLE public.folder OWNER TO postgres;

-- comment on table
COMMENT ON TABLE public.folder IS 'File Folders table';

-- default data
INSERT INTO public.folder (folder_id, name, path, parent_id)
VALUES ('c475a6f3-55ad-4641-8caa-a76bfae13fbb', 'root', '/assets', null),
      ('0e22bcfd-4b38-4f2f-8d46-d71f10b04ac3', 'images', '/images', 'c475a6f3-55ad-4641-8caa-a76bfae13fbb');

select * from public.folder;

