CREATE TABLE IF NOT EXISTS public.specification_type
(
    id         UUID UNIQUE DEFAULT uuid_generate_v4(),
    name       VARCHAR(255) NOT NULL,
    slug      VARCHAR(255) NOT NULL,
    unit       VARCHAR(255) DEFAULT NULL,
    active     BOOLEAN     DEFAULT TRUE,
    "order"    INTEGER     DEFAULT 9999,

    created_at TIMESTAMP   DEFAULT NOW(),
    created_by UUID        DEFAULT NULL REFERENCES public.user (id),
    updated_at TIMESTAMP   DEFAULT NOW(),
    updated_by UUID        DEFAULT NULL REFERENCES public.user (id),

    CONSTRAINT specification_type_pkey PRIMARY KEY (id)
);

CREATE UNIQUE INDEX specification_type_id ON public.specification_type (id);
ALTER TABLE public.specification_type
    OWNER TO postgres;
COMMENT ON TABLE public.specification_type IS 'Specification Types';

-- auto update updated_at
CREATE TRIGGER specification_type_updated_at
    BEFORE UPDATE
    ON public.user
    FOR EACH ROW
EXECUTE FUNCTION update_update_at_column();

-- auto set "order" column
CREATE OR REPLACE FUNCTION set_order_column_to_specification_type()
    RETURNS TRIGGER AS $$
BEGIN
    NEW.order = (SELECT COALESCE(MAX("order"), 0) + 1 FROM specification_type);
    RETURN NEW;
END;
$$ language 'plpgsql';

CREATE TRIGGER specification_type_order
    BEFORE INSERT
    ON public.specification_type
    FOR EACH ROW EXECUTE FUNCTION set_order_column_to_specification_type();

-- insert data
INSERT INTO public.specification_type (id, name, slug, unit)
VALUES ('a0eebc99-9c0b-4ef8-bb6d-7ab9bd380a11', 'Inch', 'inch', '″'),
       ('a0eebc99-9c0b-4ef8-bb6d-7ab9bd380a12', 'Pound', 'pound', 'lb'),
       ('a0eebc99-9c0b-4ef8-bb6d-7ab9bd380a13', 'Select', 'select', null),
       ('a0eebc99-9c0b-4ef8-bb6d-7ab9bd380a14', 'Text', 'text',  null);

-- get data
select * from public.specification_type;