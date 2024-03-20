CREATE TABLE IF NOT EXISTS public.specification_type
(
    id         UUID UNIQUE  DEFAULT uuid_generate_v4(),
    name       VARCHAR(255)               NOT NULL,
    url        VARCHAR(255)               NOT NULL,
    unit       VARCHAR(255) DEFAULT NULL,
    active     BOOLEAN      DEFAULT TRUE,
    sort_order INTEGER                    NOT NULL,

    created_at TIMESTAMP    DEFAULT NOW() NOT NULL,
    created_by UUID         DEFAULT '0e95efda-f9e2-4fac-8184-3ce2e8b7e0e1' REFERENCES public.user (id),
    updated_at TIMESTAMP    DEFAULT NOW() NOT NULL,
    updated_by UUID         DEFAULT '0e95efda-f9e2-4fac-8184-3ce2e8b7e0e1' REFERENCES public.user (id),

    CONSTRAINT specification_type_pkey PRIMARY KEY (id)
);

-- ownership and index
ALTER TABLE public.specification_type OWNER TO postgres;
CREATE INDEX specification_type_id ON public.specification_type (id);
CREATE INDEX specification_type_url ON public.specification_type (url);
CREATE INDEX specification_type_sort_order ON public.specification_type (sort_order);
CREATE INDEX specification_type_updated_at ON public.specification_type (updated_at);

-- add comment to table
COMMENT ON TABLE public.specification_type IS 'Specification Types';
COMMENT ON COLUMN public.specification_type.id IS 'Unique Identifier';
COMMENT ON COLUMN public.specification_type.name IS 'Name of the Specification Type';
COMMENT ON COLUMN public.specification_type.url IS 'URL of the Specification Type';
COMMENT ON COLUMN public.specification_type.unit IS 'Unit of the Specification Type';
COMMENT ON COLUMN public.specification_type.active IS 'Active Status of the Specification Type';
COMMENT ON COLUMN public.specification_type.sort_order IS 'Sort Order of the Specification Type';
COMMENT ON COLUMN public.specification_type.created_at IS 'Created Date of the Specification Type';
COMMENT ON COLUMN public.specification_type.created_by IS 'Created By of the Specification Type';
COMMENT ON COLUMN public.specification_type.updated_at IS 'Updated Date of the Specification Type';
COMMENT ON COLUMN public.specification_type.updated_by IS 'Updated By of the Specification Type';

-- auto update updated_at
CREATE TRIGGER specification_type_updated_at
    BEFORE UPDATE
    ON public.user
    FOR EACH ROW
EXECUTE FUNCTION update_update_at_column();

-- auto set sort_order column
CREATE TRIGGER specification_type_order
    BEFORE INSERT
    ON public.specification_type
    FOR EACH ROW
EXECUTE FUNCTION set_order_column_universal();

-- insert data
INSERT INTO public.specification_type (id, name, url, unit)
VALUES ('a0eebc99-9c0b-4ef8-bb6d-7ab9bd380a11', 'Inch', 'inch', '″'),
       ('a0eebc99-9c0b-4ef8-bb6d-7ab9bd380a12', 'Pound', 'pound', 'lb'),
       ('a0eebc99-9c0b-4ef8-bb6d-7ab9bd380a13', 'Select', 'select', null),
       ('a0eebc99-9c0b-4ef8-bb6d-7ab9bd380a14', 'Text', 'text', null);

-- get data
select * from public.specification_type;