-- drop table if exists
DROP TABLE IF EXISTS public.tag_type CASCADE;

-- create table
CREATE TABLE IF NOT EXISTS public.tag_type
(
    id                UUID UNIQUE DEFAULT uuid_generate_v4() NOT NULL,
    name              varchar(255)  DEFAULT NULL,
    url               varchar(255)  DEFAULT NULL,
    active            BOOLEAN       DEFAULT TRUE,
    sort_order        INTEGER       DEFAULT NULL,
    type              VARCHAR(50)   DEFAULT NULL,
    prefix            VARCHAR(50)   DEFAULT NULL,
    suffix            VARCHAR(50)   DEFAULT NULL,

    created_at TIMESTAMP    DEFAULT NOW()               NOT NULL,
    created_by UUID         DEFAULT NULL,
    updated_at TIMESTAMP    DEFAULT NOW()               NOT NULL,
    updated_by UUID         DEFAULT NULL,

    CONSTRAINT tag_type_pkey PRIMARY KEY (id)
);

-- ownership and index
ALTER TABLE public.tag_type OWNER TO postgres;
CREATE INDEX IF NOT EXISTS tag_type_id ON public.tag_type (id);
CREATE INDEX IF NOT EXISTS tag_type_url ON public.tag_type (url);
CREATE INDEX IF NOT EXISTS tag_type_active ON public.tag_type (active);
CREATE INDEX IF NOT EXISTS tag_type_type ON public.tag_type (type);
CREATE INDEX IF NOT EXISTS tag_type_sort_order ON public.tag_type (sort_order);
CREATE INDEX IF NOT EXISTS tag_type_updated_at ON public.tag_type (updated_at);

-- add comments
COMMENT ON TABLE  public.tag_type IS 'Tag type table';
COMMENT ON COLUMN public.tag_type.id IS 'Unique identifier';
COMMENT ON COLUMN public.tag_type.name IS 'Name of the tag type';
COMMENT ON COLUMN public.tag_type.url IS 'URL of the tag type';
COMMENT ON COLUMN public.tag_type.active IS 'Active status of the tag type';
COMMENT ON COLUMN public.tag_type.sort_order IS 'Sort order of the tag type';
COMMENT ON COLUMN public.tag_type.type IS 'Type of the tag type';
COMMENT ON COLUMN public.tag_type.prefix IS 'Prefix of the tag type';
COMMENT ON COLUMN public.tag_type.suffix IS 'Suffix of the tag type';

CREATE OR REPLACE TRIGGER update_update_at_column
    BEFORE UPDATE
    ON public.tag_type
    FOR EACH ROW
EXECUTE FUNCTION update_update_at_column();

-- auto set sort_order column
CREATE OR REPLACE FUNCTION set_order_column_universal()
    RETURNS TRIGGER AS $$
BEGIN
    IF NEW.sort_order IS NULL OR NEW.sort_order = 0 THEN
        NEW.sort_order = (SELECT COALESCE(MAX(sort_order), 0) + 1 FROM public.tag_type);
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- auto update updated_at
CREATE OR REPLACE TRIGGER tag_type_updated_at
    BEFORE UPDATE
    ON public.tag_type
    FOR EACH ROW
EXECUTE FUNCTION update_update_at_column();

-- auto set created_by
CREATE OR REPLACE TRIGGER tag_type_created_by
    BEFORE INSERT
    ON public.tag_type
    FOR EACH ROW
EXECUTE FUNCTION set_created_by_if_null();

-- auto set updated_by
CREATE OR REPLACE TRIGGER tag_type_updated_by
    BEFORE INSERT
    ON public.tag_type
    FOR EACH ROW
EXECUTE FUNCTION set_updated_by_if_null();

--
INSERT INTO public.tag_type (id, url, name, type)
VALUES
    ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd381558', 'mounting-type', 'Mounting type', 'select'),
    ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd381001', 'category', 'Category', 'select'),
    ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd381002', 'status', 'Status', 'select'),
    ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd381003', 'shipping-class', 'Shipping Class', 'select'),
    ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd381565', 'design', 'Design', 'select'),
    ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd381566', 'materials', 'Materials', 'select'),
    ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd381567', 'lighting-type', 'Lighting type', 'select'),
    ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd381568', 'of-lights', 'Of lights', 'select'),
    ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd381569', 'of-speeds', 'Of speeds', 'select'),
    ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd381570', 'control-panel-type', 'Control panel type', 'select'),
    ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd381571', 'filter-type', 'Filter type', 'select'),
    ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd381572', 'airflow', 'Airflow', 'select'),
    ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd381573', 'blower-type', 'Blower type', 'select'),
    ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd381574', 'noise-level', 'Noise level', 'select'),
    ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd381575', 'duct-size', 'Duct size', 'select'),
    ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd381576', 'exhaust-type', 'Exhaust type', 'select'),
    ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd381578', 'power-requirements', 'Power requirements', 'select'),
    ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd381579', 'certifications', 'Certifications', 'select'),
    ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd381580', 'warranty', 'Warranty', 'select'),
    ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd381581', 'order-processing-time', 'Order processing time', 'select'),
    ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd381582', 'shipping-speed', 'Shipping speed', 'select'),
    ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd381583', 'ships-via', 'Ships via', 'select'),
    ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd381584', 'country-of-tag_typeion', 'Country of tag_typeion', 'select'),
    ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd381585', 'width-group', 'Width group', 'select'),
    ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd381591', 'recommended-range-width', 'Recommended range width', 'select'),
    ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd381602', 'depth', 'Depth', 'select'),
    ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd381604', 'color-finish', 'Color finish', 'select'),
    ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd381620', 'shipping-weight', 'Shipping weight', 'select'),
    ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd381641', 'height', 'Height', 'select'),
    ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd381671', 'brand', 'Brand', 'select'),
    ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd381743', 'item-weight', 'Item weight', 'select'),
    ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd381875', 'additional-lighting', 'Additional lighting', 'select'),
    ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd382001', 'filter-stage-1', 'Filter stage 1', 'select'),
    ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd382002', 'filter-stage-2', 'Filter stage 2', 'select'),
    ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd382074', 'filter-color', 'Filter color', 'select'),
    ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd382077', 'filter-material', 'Filter material', 'select'),
    ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd382079', 'filter-exhaust-type', 'Filter exhaust type', 'select'),
    ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd382082', 'filter-design', 'Filter design', 'select'),
    ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd382142', 'height-above-cooktop', 'Height above cooktop', 'select'),
    ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd382152', 'max-usable-length', 'Max usable length', 'select'),
    ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd382157', 'width', 'Width', 'select'),
    ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd382368', 'diameter', 'Diameter', 'select'),
    ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd382530', 'filter-accessories', 'Filter accessories', 'select'),
    ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd383138', 'qty-per-box', 'Qty per box', 'select'),
    ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd383139', 'keywords', 'Keywords', 'select')
ON CONFLICT (id) DO UPDATE
    SET url = EXCLUDED.url,
        name = EXCLUDED.name,
        type = EXCLUDED.type,
        suffix = EXCLUDED.suffix;

-- get data
select * from public.tag_type;

-- woocommerce data
-- SELECT
--     CONCAT('a0eebc99-9c0b-4ef8-bb6d-6bb9bd38', tt.term_taxonomy_id) as id,
--     SUBSTRING(tt.taxonomy, 4) AS url,
--     CONCAT(UCASE(SUBSTRING(REPLACE(SUBSTRING(tt.taxonomy, 4), '-', ' '), 1, 1)),
--            LOWER(SUBSTRING(REPLACE(SUBSTRING(tt.taxonomy, 4), '-', ' '), 2))) AS name,
--     'select' as type,
--     false as prime,
--     false as list_item,
--     false as filter,
--     false as required,
--     '' as suffix
-- FROM
--     wp_term_taxonomy tt
-- WHERE
--     tt.taxonomy LIKE 'pa_%'
--   AND tt.count > 0
-- GROUP BY tt.taxonomy
--
-- UNION ALL
--
-- SELECT
--     'a0eebc99-9c0b-4ef8-bb6d-6bb9bd381001' AS id,
--     'category' AS url,
--     'Category' AS name,
--     'select' as type,
--     true as prime,
--     true as list_item,
--     true as filter,
--     true as required,
--     '' as suffix
--
-- UNION ALL
--
-- SELECT
--     'a0eebc99-9c0b-4ef8-bb6d-6bb9bd381002' AS id,
--     'status' AS url,
--     'Status' AS name,
--     'select' as type,
--     true as prime,
--     true as list_item,
--     true as filter,
--     true as required,
--     '' as suffix
--
-- UNION ALL
--
-- SELECT
--     'a0eebc99-9c0b-4ef8-bb6d-6bb9bd381003' AS id,
--     'shipping-class' AS url,
--     'Shipping Class' AS name,
--     'select' as type,
--     true as prime,
--     true as list_item,
--     true as filter,
--     true as required,
--     '' as suffix
--
-- ORDER BY id;