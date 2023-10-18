CREATE TABLE IF NOT EXISTS public.specification
(
    id         UUID UNIQUE DEFAULT uuid_generate_v4(),
    name       VARCHAR(255)        NOT NULL,
    slug       VARCHAR(255) UNIQUE NOT NULL,
    type       UUID                NOT NULL REFERENCES public.specification_type (id),
    active     BOOLEAN     DEFAULT TRUE,
    "order"    INTEGER     DEFAULT 9999,

    created_at TIMESTAMP   DEFAULT NOW(),
    created_by UUID        DEFAULT NULL REFERENCES public.user (id),
    updated_at TIMESTAMP   DEFAULT NOW(),
    updated_by UUID        DEFAULT NULL REFERENCES public.user (id),

    CONSTRAINT specification_pkey PRIMARY KEY (id)
);

CREATE UNIQUE INDEX specification_id ON public.specification (id);
CREATE UNIQUE INDEX specification_slug ON public.specification (slug);
ALTER TABLE public.specification
    OWNER TO postgres;
COMMENT ON TABLE public.specification IS 'Product Specification';

-- insert data
INSERT INTO public.specification (id, name, slug, "order", type)
VALUES ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380107','Mounting Type','mounting-type',1, 'a0eebc99-9c0b-4ef8-bb6d-7ab9bd380a13'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380108','Width','width',2, 'a0eebc99-9c0b-4ef8-bb6d-7ab9bd380a13'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380109','Depth','depth',3, 'a0eebc99-9c0b-4ef8-bb6d-7ab9bd380a13'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380110','Height','height',4, 'a0eebc99-9c0b-4ef8-bb6d-7ab9bd380a13'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380111','Recommended Range Width','recommended-range-width',5, 'a0eebc99-9c0b-4ef8-bb6d-7ab9bd380a13'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380112','Height Above Cooktop','height-above-cooktop',6, 'a0eebc99-9c0b-4ef8-bb6d-7ab9bd380a13'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380113','Color / Finish','color-finish',7, 'a0eebc99-9c0b-4ef8-bb6d-7ab9bd380a13'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380114','Design','design',8, 'a0eebc99-9c0b-4ef8-bb6d-7ab9bd380a13'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380115','Materials','materials',9, 'a0eebc99-9c0b-4ef8-bb6d-7ab9bd380a13'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380116','Lighting Type','lighting-type',10, 'a0eebc99-9c0b-4ef8-bb6d-7ab9bd380a13'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380117','# of Lights','of-lights',11, 'a0eebc99-9c0b-4ef8-bb6d-7ab9bd380a13'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380118','# of Speeds','of-speeds',12, 'a0eebc99-9c0b-4ef8-bb6d-7ab9bd380a13'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380119','Control Panel Type','control-panel-type',13, 'a0eebc99-9c0b-4ef8-bb6d-7ab9bd380a13'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380120','Filter Type','filter-type',14, 'a0eebc99-9c0b-4ef8-bb6d-7ab9bd380a13'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380121','Airflow','airflow',15, 'a0eebc99-9c0b-4ef8-bb6d-7ab9bd380a13'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380122','Blower Type','blower-type',16, 'a0eebc99-9c0b-4ef8-bb6d-7ab9bd380a13'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380123','Noise Level','noise-level',17, 'a0eebc99-9c0b-4ef8-bb6d-7ab9bd380a13'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380124','Duct Size','duct-size',18, 'a0eebc99-9c0b-4ef8-bb6d-7ab9bd380a13'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380125','Exhaust Type','exhaust-type',19, 'a0eebc99-9c0b-4ef8-bb6d-7ab9bd380a13'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380126','Power Requirements','power-requirements',20, 'a0eebc99-9c0b-4ef8-bb6d-7ab9bd380a13'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380127','Certifications','certifications',21, 'a0eebc99-9c0b-4ef8-bb6d-7ab9bd380a13'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380128','Warranty','warranty',22, 'a0eebc99-9c0b-4ef8-bb6d-7ab9bd380a13'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380129','Order Processing Time','order-processing-time',23, 'a0eebc99-9c0b-4ef8-bb6d-7ab9bd380a13'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380130','Shipping Speed','shipping-speed',24, 'a0eebc99-9c0b-4ef8-bb6d-7ab9bd380a13'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380131','Ships Via','ships-via',25, 'a0eebc99-9c0b-4ef8-bb6d-7ab9bd380a13'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380132','Country of Production','country-of-production',26, 'a0eebc99-9c0b-4ef8-bb6d-7ab9bd380a13'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380133','Filter - Width (Group)','width-group',27, 'a0eebc99-9c0b-4ef8-bb6d-7ab9bd380a13'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380135','Shipping Weight','shipping-weight',28, 'a0eebc99-9c0b-4ef8-bb6d-7ab9bd380a13'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380136','Brand','brand',29, 'a0eebc99-9c0b-4ef8-bb6d-7ab9bd380a13'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380137','Item Weight','item-weight',30, 'a0eebc99-9c0b-4ef8-bb6d-7ab9bd380a13'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380138','Diameter','diameter',31, 'a0eebc99-9c0b-4ef8-bb6d-7ab9bd380a13'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380139','Additional Lighting','additional-lighting',32, 'a0eebc99-9c0b-4ef8-bb6d-7ab9bd380a13'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380142','Filter – Color','filter-color',33, 'a0eebc99-9c0b-4ef8-bb6d-7ab9bd380a13'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380143','Filter – Material','filter-material',34, 'a0eebc99-9c0b-4ef8-bb6d-7ab9bd380a13'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380144','Filter – Exhaust Type','filter-exhaust-type',35, 'a0eebc99-9c0b-4ef8-bb6d-7ab9bd380a13'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380145','Filter – Design','filter-design',36, 'a0eebc99-9c0b-4ef8-bb6d-7ab9bd380a13'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380146','Length','max-usable-length',37, 'a0eebc99-9c0b-4ef8-bb6d-7ab9bd380a13'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380147','Filter – Accessories','filter-accessories',38, 'a0eebc99-9c0b-4ef8-bb6d-7ab9bd380a13');


-- get data 
select * from public.specification;

-- get data from WooCommerce DataBase
-- select
--     CONCAT('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380', attribute_id) as id,
--     attribute_label as name,
--     attribute_name as slug,
--     attribute_id as sort_order,
--     'a0eebc99-9c0b-4ef8-bb6d-7ab9bd380a13' as type
-- from wp_woocommerce_attribute_taxonomies;
