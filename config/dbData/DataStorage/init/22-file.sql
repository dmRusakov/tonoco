-- create table
CREATE TABLE IF NOT EXISTS public.file
(
    id         uuid unique DEFAULT uuid_generate_v4(),
    name       varchar(255) not null,
    slug       varchar(255) not null,

    created_at TIMESTAMP   DEFAULT NOW(),
    created_by UUID        DEFAULT NULL REFERENCES public.user (id),
    updated_at TIMESTAMP   DEFAULT NOW(),
    updated_by UUID        DEFAULT NULL REFERENCES public.user (id),

    CONSTRAINT file_pkey PRIMARY KEY (id)
);

-- -- index, constraint and ownership
CREATE UNIQUE INDEX file_id ON public.file (id);
ALTER TABLE public.file OWNER TO postgres;
COMMENT ON TABLE public.file IS 'File table';

-- auto update updated_at
CREATE TRIGGER update_file_updated_at
    BEFORE UPDATE
    ON public.file
    FOR EACH ROW
EXECUTE PROCEDURE update_update_at_column();