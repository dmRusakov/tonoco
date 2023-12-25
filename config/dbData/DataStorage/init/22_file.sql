-- create table
CREATE TABLE IF NOT EXISTS public.file
(
    id         uuid unique DEFAULT uuid_generate_v4(),
    name       varchar(255) not null,
    slug       varchar(255) not null,
    active           BOOLEAN      DEFAULT TRUE,
    "order"          INTEGER      DEFAULT null,

    created_at TIMESTAMP   DEFAULT NOW(),
    created_by UUID        DEFAULT NULL REFERENCES public.user (id),
    updated_at TIMESTAMP   DEFAULT NOW(),
    updated_by UUID        DEFAULT NULL REFERENCES public.user (id),

    CONSTRAINT file_pkey PRIMARY KEY (id)
);

-- -- index, constraint and ownership
CREATE INDEX file_id ON public.file (id);
ALTER TABLE public.file OWNER TO postgres;
COMMENT ON TABLE public.file IS 'File table';

-- auto update updated_at
CREATE TRIGGER update_file_updated_at
    BEFORE UPDATE
    ON public.file
    FOR EACH ROW
EXECUTE PROCEDURE update_update_at_column();

-- auto set "order" column
CREATE OR REPLACE FUNCTION set_order_column_to_file()
    RETURNS TRIGGER AS $$
BEGIN
    NEW.order = (SELECT COALESCE(MAX("order"), 0) + 1 FROM file);
    RETURN NEW;
END;
$$ language 'plpgsql';

CREATE TRIGGER file_order
    BEFORE INSERT
    ON public.file
    FOR EACH ROW EXECUTE FUNCTION set_order_column_to_file();
