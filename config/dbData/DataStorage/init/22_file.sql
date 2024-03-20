-- create table
CREATE TABLE IF NOT EXISTS public.file
(
    id         uuid unique DEFAULT uuid_generate_v4(),
    name       varchar(255) not null,
    url        varchar(255) not null,
    active     BOOLEAN     DEFAULT TRUE,
    parent_id  uuid        DEFAULT null,
    sort_order INTEGER     DEFAULT null,

    created_at TIMESTAMP   DEFAULT NOW()              NOT NULL,
    created_by UUID        DEFAULT '0e95efda-f9e2-4fac-8184-3ce2e8b7e0e1' REFERENCES public.user (id),
    updated_at TIMESTAMP   DEFAULT NOW()              NOT NULL,
    updated_by UUID        DEFAULT '0e95efda-f9e2-4fac-8184-3ce2e8b7e0e1' REFERENCES public.user (id),

    CONSTRAINT file_pkey PRIMARY KEY (id)
);

-- -- index, constraint and ownership
CREATE INDEX file_id ON public.file (id);
ALTER TABLE public.file
    OWNER TO postgres;
COMMENT ON TABLE public.file IS 'File table';

-- auto update updated_at
CREATE TRIGGER update_file_updated_at
    BEFORE UPDATE
    ON public.file
    FOR EACH ROW
EXECUTE PROCEDURE update_update_at_column();

-- auto set sort_order column
CREATE OR REPLACE FUNCTION set_order_column_to_file()
    RETURNS TRIGGER AS
$$
BEGIN
    NEW.sort_order = (SELECT COALESCE(MAX(sort_order)  + 1, 0) FROM file);
    RETURN NEW;
END;
$$ language 'plpgsql';

CREATE TRIGGER file_order
    BEFORE INSERT
    ON public.file
    FOR EACH ROW
EXECUTE FUNCTION set_order_column_to_file();
