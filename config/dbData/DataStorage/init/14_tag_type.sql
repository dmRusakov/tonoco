-- create table
CREATE TABLE IF NOT EXISTS public.tag_type
(
    id                UUID UNIQUE DEFAULT uuid_generate_v4() NOT NULL,
    name              varchar(255)  NOT NULL,
    url               varchar(255)  NOT NULL,
    short_description VARCHAR(255)  DEFAULT NULL,
    description       VARCHAR(6000) DEFAULT NULL,
    required          BOOLEAN       DEFAULT FALSE,
    active            BOOLEAN       DEFAULT TRUE,
    prime             BOOLEAN       DEFAULT FALSE,
    list_item         BOOLEAN       DEFAULT FALSE,
    filter            BOOLEAN       DEFAULT FALSE,
    sort_order        INTEGER       DEFAULT NULL,
    type              VARCHAR(50)   DEFAULT NULL,
    prefix            VARCHAR(50)   DEFAULT NULL,
    suffix            VARCHAR(50)   DEFAULT NULL,

    created_at TIMESTAMP    DEFAULT NOW()               NOT NULL,
    created_by UUID         DEFAULT '0e95efda-f9e2-4fac-8184-3ce2e8b7e0e1' REFERENCES public.user (id),
    updated_at TIMESTAMP    DEFAULT NOW()               NOT NULL,
    updated_by UUID         DEFAULT '0e95efda-f9e2-4fac-8184-3ce2e8b7e0e1' REFERENCES public.user (id),

    CONSTRAINT tag_type_pkey PRIMARY KEY (id)
    );

-- ownership and index
ALTER TABLE public.tag_type OWNER TO postgres;
CREATE INDEX tag_type_id ON public.tag_type (id);
CREATE INDEX tag_type_url ON public.tag_type (url);
CREATE INDEX tag_type_prime ON public.tag_type (prime);
CREATE INDEX tag_type_active ON public.tag_type (active);
CREATE INDEX tag_type_list_item ON public.tag_type (list_item);
CREATE INDEX tag_type_filter ON public.tag_type (filter);
CREATE INDEX tag_type_type ON public.tag_type (type);
CREATE INDEX tag_type_sort_order ON public.tag_type (sort_order);
CREATE INDEX tag_type_updated_at ON public.tag_type (updated_at);

-- add comments
COMMENT ON TABLE  public.tag_type IS 'Tag type table';
COMMENT ON COLUMN public.tag_type.id IS 'Unique identifier';
COMMENT ON COLUMN public.tag_type.name IS 'Name of the tag type';
COMMENT ON COLUMN public.tag_type.url IS 'URL of the tag type';
COMMENT ON COLUMN public.tag_type.short_description IS 'Short description of the tag type';
COMMENT ON COLUMN public.tag_type.description IS 'Description of the tag type';
COMMENT ON COLUMN public.tag_type.required IS 'Required status of the tag type';
COMMENT ON COLUMN public.tag_type.prime IS 'Prime status of the tag type';
COMMENT ON COLUMN public.tag_type.active IS 'Active status of the tag type';
COMMENT ON COLUMN public.tag_type.list_item IS 'In list item status of the tag type';
COMMENT ON COLUMN public.tag_type.filter IS 'Filter status of the tag type';
COMMENT ON COLUMN public.tag_type.sort_order IS 'Sort order of the tag type';
COMMENT ON COLUMN public.tag_type.type IS 'Type of the tag type';
COMMENT ON COLUMN public.tag_type.prefix IS 'Prefix of the tag type';
COMMENT ON COLUMN public.tag_type.suffix IS 'Suffix of the tag type';

-- auto update updated_at
CREATE TRIGGER tag_type_updated_at
    BEFORE UPDATE
    ON public.tag_type
    FOR EACH ROW
    EXECUTE FUNCTION update_update_at_column();

-- auto set sort_order column
CREATE TRIGGER tag_type_order
    BEFORE INSERT
    ON public.tag_type
    FOR EACH ROW
    EXECUTE FUNCTION set_order_column_universal();

-- default data
-- insert data
INSERT INTO public.tag_type (id, name, url, type, prime, list_item, filter, required, suffix)
VALUES ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380101', 'Category', 'category', 'select', true, true, true,  true, null),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380102', 'Status', 'status', 'select', true, false, false,  true, null),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380103', 'Shipping class', 'shipping-class', 'select', true, false, false,  true, null),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380106', 'Mounting Type', 'mounting-type', 'select', true, true, true,  true, null),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380108', 'Width', 'width', 'text', false, false, false,  false, '″'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380109', 'Depth', 'depth', 'text', false, false, false,  false, '″'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380110', 'Height', 'height', 'text', false, false, false,  false, '″'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380111', 'Recommended Range Width', 'recommended-range-width', 'select', false, false, false,  false, null),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380112', 'Height Above Cooktop', 'height-above-cooktop', 'select', false, false, false,  false, null),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380113', 'Color / Finish', 'color-finish', 'select', false, false, true,  true, null),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380114', 'Design', 'design', 'select', false, false, false,  false, null),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380115', 'Materials', 'materials', 'select', false, false, false,  false, null),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380116', 'Lighting Type', 'lighting-type', 'select', false, false, false,  false, null),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380117', '# of Lights', 'of-lights', 'select', false, false, false,  false, null),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380118', '# of Speeds', 'of-speeds', 'select', false, false, false,  false, null),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380119', 'Control Panel Type', 'control-panel-type', 'select', false, false, false,  false, null),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380120', 'Filter Type', 'filter-type', 'select', false, false, false,  false, null),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380121', 'Airflow', 'airflow', 'select', false, false, false,  false, null),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380122', 'Blower Type', 'blower-type', 'select', false, false, false,  false, null),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380123', 'Noise Level', 'noise-level', 'select', false, false, false,  false, null),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380124', 'Duct Size', 'duct-size', 'select', false, false, false,  false, null),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380125', 'Exhaust Type', 'exhaust-type', 'select', false, false, false,  false, null),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380126', 'Power Requirements', 'power-requirements', 'select', false, false, false,  false, null),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380127', 'Certifications', 'certifications', 'select', false, false, false,  false, null),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380128', 'Warranty', 'warranty', 'select', false, false, false,  false, null),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380129', 'Order Processing Time', 'order-processing-time', 'select', false, false, false,  false, null),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380130', 'Shipping Speed', 'shipping-speed', 'select', false, false, false,  false, null),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380131', 'Ships Via', 'ships-via', 'select', false, false, false,  false, null),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380132', 'Country of Production', 'country-of-production', 'select', false, false, false,  false, null),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380133', 'Filter - Width (Group)', 'width-group', 'select', false, false, false,  false, null),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380135', 'Shipping Weight', 'shipping-weight', 'select', false, false, false,  false, null),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380136', 'Brand', 'brand', 'select', false, false, false,  false, null),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380137', 'Item Weight', 'item-weight', 'select', false, false, false,  false, null),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380138', 'Diameter', 'diameter', 'select', false, false, false,  false, null),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380139', 'Additional Lighting', 'additional-lighting', 'select', false, false, false,  false, null),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380142', 'Filter – Color', 'filter-color', 'select', false, false, true,  false, null),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380143', 'Filter – Material', 'filter-material', 'select', false, false, true,  false, null),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380144', 'Filter – Exhaust Type', 'filter-exhaust-type', 'select', false, false, true,  false, null),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380145', 'Filter – Design', 'filter-design', 'select', false, false, true,  false, null),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380146', 'Length', 'max-usable-length', 'select', false, false, true,  false, null),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380147', 'Filter – Accessories', 'filter-accessories', 'select', false, false, true,  false, null);

-- get data
select * from public.tag_type;
