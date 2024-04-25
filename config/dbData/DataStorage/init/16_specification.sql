CREATE TABLE IF NOT EXISTS public.specification
(
    id                 UUID UNIQUE DEFAULT uuid_generate_v4(),
    name               VARCHAR(255)              NOT NULL,
    url                VARCHAR(255) UNIQUE       NOT NULL,
    specification_type UUID                      NOT NULL REFERENCES public.specification_type (id),
    active             BOOLEAN     DEFAULT TRUE,
    sort_order         INTEGER                   NOT NULL,

    created_at         TIMESTAMP   DEFAULT NOW() NOT NULL,
    created_by         UUID        DEFAULT '0e95efda-f9e2-4fac-8184-3ce2e8b7e0e1' REFERENCES public.user (id),
    updated_at         TIMESTAMP   DEFAULT NOW() NOT NULL,
    updated_by         UUID        DEFAULT '0e95efda-f9e2-4fac-8184-3ce2e8b7e0e1' REFERENCES public.user (id),

    CONSTRAINT specification_pkey PRIMARY KEY (id)
);

-- ownership and index
ALTER TABLE public.specification OWNER TO postgres;
CREATE INDEX specification_id ON public.specification (id);
CREATE INDEX specification_url ON public.specification (url);
CREATE INDEX specification_specification_type ON public.specification (specification_type);
CREATE INDEX specification_sort_order ON public.specification (sort_order);
CREATE INDEX specification_updated_at ON public.specification (updated_at);

-- add comment to table
COMMENT ON TABLE public.specification IS 'Product Specification';
COMMENT ON COLUMN public.specification.id IS 'Unique Identifier';
COMMENT ON COLUMN public.specification.name IS 'Specification Name';
COMMENT ON COLUMN public.specification.url IS 'Specification URL';
COMMENT ON COLUMN public.specification.specification_type IS 'Specification Type';
COMMENT ON COLUMN public.specification.active IS 'Active Status';
COMMENT ON COLUMN public.specification.sort_order IS 'Sort Order';
COMMENT ON COLUMN public.specification.created_at IS 'Record Created Date';
COMMENT ON COLUMN public.specification.created_by IS 'Record Created By';
COMMENT ON COLUMN public.specification.updated_at IS 'Record Updated Date';
COMMENT ON COLUMN public.specification.updated_by IS 'Record Updated By';

-- auto update updated_at
CREATE TRIGGER specification_updated_at
    BEFORE UPDATE
    ON public.specification
    FOR EACH ROW
EXECUTE FUNCTION update_update_at_column();

-- auto set sort_order column
CREATE TRIGGER specification_order
    BEFORE INSERT
    ON public.specification
    FOR EACH ROW
EXECUTE FUNCTION set_order_column_universal();

-- insert data
INSERT INTO public.specification (id, name, url, specification_type)
VALUES ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380107', 'Mounting Type', 'mounting-type', 'a0eebc99-9c0b-4ef8-bb6d-7ab9bd380a13'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380108', 'Width', 'width', 'a0eebc99-9c0b-4ef8-bb6d-7ab9bd380a13'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380109', 'Depth', 'depth', 'a0eebc99-9c0b-4ef8-bb6d-7ab9bd380a13'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380110', 'Height', 'height', 'a0eebc99-9c0b-4ef8-bb6d-7ab9bd380a13'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380111', 'Recommended Range Width', 'recommended-range-width', 'a0eebc99-9c0b-4ef8-bb6d-7ab9bd380a13'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380112', 'Height Above Cooktop', 'height-above-cooktop', 'a0eebc99-9c0b-4ef8-bb6d-7ab9bd380a13'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380113', 'Color / Finish', 'color-finish', 'a0eebc99-9c0b-4ef8-bb6d-7ab9bd380a13'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380114', 'Design', 'design', 'a0eebc99-9c0b-4ef8-bb6d-7ab9bd380a13'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380115', 'Materials', 'materials', 'a0eebc99-9c0b-4ef8-bb6d-7ab9bd380a13'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380116', 'Lighting Type', 'lighting-type', 'a0eebc99-9c0b-4ef8-bb6d-7ab9bd380a13'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380117', '# of Lights', 'of-lights', 'a0eebc99-9c0b-4ef8-bb6d-7ab9bd380a13'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380118', '# of Speeds', 'of-speeds', 'a0eebc99-9c0b-4ef8-bb6d-7ab9bd380a13'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380119', 'Control Panel Type', 'control-panel-type', 'a0eebc99-9c0b-4ef8-bb6d-7ab9bd380a13'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380120', 'Filter Type', 'filter-type', 'a0eebc99-9c0b-4ef8-bb6d-7ab9bd380a13'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380121', 'Airflow', 'airflow', 'a0eebc99-9c0b-4ef8-bb6d-7ab9bd380a13'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380122', 'Blower Type', 'blower-type', 'a0eebc99-9c0b-4ef8-bb6d-7ab9bd380a13'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380123', 'Noise Level', 'noise-level', 'a0eebc99-9c0b-4ef8-bb6d-7ab9bd380a13'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380124', 'Duct Size', 'duct-size', 'a0eebc99-9c0b-4ef8-bb6d-7ab9bd380a13'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380125', 'Exhaust Type', 'exhaust-type', 'a0eebc99-9c0b-4ef8-bb6d-7ab9bd380a13'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380126', 'Power Requirements', 'power-requirements', 'a0eebc99-9c0b-4ef8-bb6d-7ab9bd380a13'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380127', 'Certifications', 'certifications', 'a0eebc99-9c0b-4ef8-bb6d-7ab9bd380a13'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380128', 'Warranty', 'warranty', 'a0eebc99-9c0b-4ef8-bb6d-7ab9bd380a13'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380129', 'Order Processing Time', 'order-processing-time', 'a0eebc99-9c0b-4ef8-bb6d-7ab9bd380a13'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380130', 'Shipping Speed', 'shipping-speed', 'a0eebc99-9c0b-4ef8-bb6d-7ab9bd380a13'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380131', 'Ships Via', 'ships-via', 'a0eebc99-9c0b-4ef8-bb6d-7ab9bd380a13'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380132', 'Country of Production', 'country-of-production', 'a0eebc99-9c0b-4ef8-bb6d-7ab9bd380a13'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380133', 'Filter - Width (Group)', 'width-group', 'a0eebc99-9c0b-4ef8-bb6d-7ab9bd380a13'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380135', 'Shipping Weight', 'shipping-weight', 'a0eebc99-9c0b-4ef8-bb6d-7ab9bd380a13'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380136', 'Brand', 'brand', 'a0eebc99-9c0b-4ef8-bb6d-7ab9bd380a13'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380137', 'Item Weight', 'item-weight', 'a0eebc99-9c0b-4ef8-bb6d-7ab9bd380a13'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380138', 'Diameter', 'diameter', 'a0eebc99-9c0b-4ef8-bb6d-7ab9bd380a13'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380139', 'Additional Lighting', 'additional-lighting', 'a0eebc99-9c0b-4ef8-bb6d-7ab9bd380a13'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380142', 'Filter – Color', 'filter-color', 'a0eebc99-9c0b-4ef8-bb6d-7ab9bd380a13'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380143', 'Filter – Material', 'filter-material', 'a0eebc99-9c0b-4ef8-bb6d-7ab9bd380a13'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380144', 'Filter – Exhaust Type', 'filter-exhaust-type', 'a0eebc99-9c0b-4ef8-bb6d-7ab9bd380a13'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380145', 'Filter – Design', 'filter-design', 'a0eebc99-9c0b-4ef8-bb6d-7ab9bd380a13'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380146', 'Length', 'max-usable-length', 'a0eebc99-9c0b-4ef8-bb6d-7ab9bd380a13'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380147', 'Filter – Accessories', 'filter-accessories', 'a0eebc99-9c0b-4ef8-bb6d-7ab9bd380a13');

-- get data 
select * from public.specification;

-- get data from WooCommerce DataBase
-- select
--     CONCAT('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380', attribute_id) as id,
--     attribute_label as name,
--     attribute_name as url,
--     attribute_id as sort_order,
--     'a0eebc99-9c0b-4ef8-bb6d-7ab9bd380a13' as type
-- from wp_woocommerce_attribute_taxonomies;