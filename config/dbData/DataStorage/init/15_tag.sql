CREATE TABLE IF NOT EXISTS public.tag_select
(
    id                  UUID UNIQUE DEFAULT uuid_generate_v4(),
    tag_type_id UUID    NOT NULL REFERENCES public.tag_type (id),
    name                VARCHAR(255)            NOT NULL,
    short_description   VARCHAR(255)            DEFAULT NULL,
    description         VARCHAR(6000)           DEFAULT NULL,
    url                 VARCHAR(255)            NOT NULL,
    active              BOOLEAN                 DEFAULT TRUE,
    sort_order          INTEGER                 NOT NULL,

    created_at          TIMESTAMP               DEFAULT NOW() NOT NULL,
    created_by          UUID                    DEFAULT '0e95efda-f9e2-4fac-8184-3ce2e8b7e0e1' REFERENCES public.user (id),
    updated_at          TIMESTAMP               DEFAULT NOW() NOT NULL,
    updated_by          UUID                    DEFAULT '0e95efda-f9e2-4fac-8184-3ce2e8b7e0e1' REFERENCES public.user (id),

    CONSTRAINT tag_select_pkey PRIMARY KEY (id)
);

-- ownership and index
ALTER TABLE public.tag_select OWNER TO postgres;
CREATE INDEX tag_select_id ON public.tag_select (id);
CREATE INDEX tag_select_tag_type_id ON public.tag_select (tag_type_id);
CREATE INDEX tag_select_url ON public.tag_select (url);
CREATE INDEX tag_select_sort_order ON public.tag_select (sort_order);
CREATE INDEX tag_select_updated_at ON public.tag_select (updated_at);

-- add comment to table
COMMENT ON TABLE public.tag_select IS 'Product Tag';
COMMENT ON COLUMN public.tag_select.id IS 'Unique identifier for product tag';
COMMENT ON COLUMN public.tag_select.tag_type_id IS 'Reference to tag type';
COMMENT ON COLUMN public.tag_select.name IS 'Name of product tag';
COMMENT ON COLUMN public.tag_select.url IS 'URL of product tag';
COMMENT ON COLUMN public.tag_select.short_description IS 'Short description of product tag';
COMMENT ON COLUMN public.tag_select.description IS 'Description of product tag';
COMMENT ON COLUMN public.tag_select.active IS 'Active status of product tag';
COMMENT ON COLUMN public.tag_select.sort_order IS 'Sort order of product tag';
COMMENT ON COLUMN public.tag_select.created_at IS 'Creation time of product tag';
COMMENT ON COLUMN public.tag_select.created_by IS 'Creator of product tag';
COMMENT ON COLUMN public.tag_select.updated_at IS 'Update time of product tag';
COMMENT ON COLUMN public.tag_select.updated_by IS 'Updater of product tag';

-- auto update updated_at
CREATE TRIGGER tag_select_updated_at
    BEFORE UPDATE
    ON public.user
    FOR EACH ROW
EXECUTE FUNCTION update_update_at_column();

-- auto set sort_order column
CREATE TRIGGER tag_select_order
    BEFORE INSERT
    ON public.tag_select
    FOR EACH ROW
EXECUTE FUNCTION set_order_column_universal();

-- insert data
INSERT INTO public.tag_select (id, specification_id, name, url, sort_order)
VALUES ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd382160', (select id from public.tag_type where url = 'additional-lighting'), '4x energy-efficient LED lights', '4x-energy-efficient-led-lights', '2160'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd381875', (select id from public.tag_type where url = 'additional-lighting'), '4x Internal Body Illumination LED Lights', '4x-internal-body-illumination-led-lights', '1875'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd381844', (select id from public.tag_type where url = 'airflow'), '150-800 CFM. CFM can be reduced upon customer request, to accommodate local building code requirements.', '150-800-cfm-cfm-can-be-reduced-upon-customer-request-to-accommodate-local-building-code-requirements', '1844'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd381895', (select id from public.tag_type where url = 'airflow'), '150-940 CFM (zero static pressure)', '150-940-cfm-zero-static-pressure', '1895'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd381572', (select id from public.tag_type where url = 'airflow'), '150-940 CFM. CFM can be reduced upon customer request, to accommodate local building code requirements.', '150-940-cfm-cfm-can-be-reduced-upon-customer-request-to-accommodate-local-building-code-requirements', '1572'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd381573', (select id from public.tag_type where url = 'blower-type'), 'Internal Whisper-Quiet tangential (INCLUDED)', 'internal-whisper-quiet-tangential-included', '1573'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd381671', (select id from public.tag_type where url = 'brand'), 'FuturoFuturo', 'futurofuturo', '1671'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd381579', (select id from public.tag_type where url = 'certifications'), 'UL CSA Listed for USA and Canada', 'ul-csa-listed-for-usa-and-canada', '1579'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd381628', (select id from public.tag_type where url = 'color-finish'), 'Black', 'black', '1628'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd383229', (select id from public.tag_type where url = 'color-finish'), 'Black gunmetal', 'black-gunmetal', '3229'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd383206', (select id from public.tag_type where url = 'color-finish'), 'Black, matte', 'black-matte', '3206'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd382062', (select id from public.tag_type where url = 'color-finish'), 'Brass', 'brass', '2062'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd382031', (select id from public.tag_type where url = 'color-finish'), 'Copper', 'copper', '2031'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd382063', (select id from public.tag_type where url = 'color-finish'), 'Glass', 'glass', '2063'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd383232', (select id from public.tag_type where url = 'color-finish'), 'Gold', 'gold', '3232'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd383180', (select id from public.tag_type where url = 'color-finish'), 'Graphite', 'graphite', '3180'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd382056', (select id from public.tag_type where url = 'color-finish'), 'Grey', 'grey', '2056'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd382038', (select id from public.tag_type where url = 'color-finish'), 'Iron', 'iron', '2038'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd382064', (select id from public.tag_type where url = 'color-finish'), 'Orange', 'orange', '2064'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd381604', (select id from public.tag_type where url = 'color-finish'), 'Stainless Steel', 'stainless-steel', '1604'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd381615', (select id from public.tag_type where url = 'color-finish'), 'White', 'white', '1615'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd383178', (select id from public.tag_type where url = 'color-finish'), 'White (Semi-Gloss)', 'white-semi-gloss', '3178'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd382065', (select id from public.tag_type where url = 'color-finish'), 'Wood', 'wood', '2065'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd383260', (select id from public.tag_type where url = 'color-finish'), 'Wood (raw ash)', 'wood-raw-ash', '3260'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd381598', (select id from public.tag_type where url = 'control-panel-type'), 'Convenient 3-speed slider control', 'convenient-3-speed-slider-control', '1598'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd381570', (select id from public.tag_type where url = 'control-panel-type'), 'Electronic, illuminated control panel', 'electronic-illuminated-control-panel', '1570'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd381686', (select id from public.tag_type where url = 'control-panel-type'), 'Electronic, illuminated control panel. Wireless Remote Control (INCLUDED)', 'electronic-illuminated-control-panel-wireless-remote-control-included', '1686'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd381701', (select id from public.tag_type where url = 'control-panel-type'), 'Futuristic touch-sensitive optical control panel', 'futuristic-touch-sensitive-optical-control-panel', '1701'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd382011', (select id from public.tag_type where url = 'control-panel-type'), 'Illuminated 5-button control panel + rotary dimmer', 'illuminated-5-button-control-panel-rotary-dimmer', '2011'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd382027', (select id from public.tag_type where url = 'control-panel-type'), 'Illuminated rotary switch', 'illuminated-rotary-switch', '2027'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd383274', (select id from public.tag_type where url = 'control-panel-type'), 'Wireless Remote Control (INCLUDED)', 'wireless-remote-control-included', '3274'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd381584', (select id from public.tag_type where url = 'country-of-production'), 'Made in Italy', 'made-in-italy', '1584'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd381978', (select id from public.tag_type where url = 'design'), 'Concealed - fits all', 'concealed-fits-all', '1978'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd381565', (select id from public.tag_type where url = 'design'), 'Contemporary', 'contemporary', '1565'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd381605', (select id from public.tag_type where url = 'design'), 'Modern', 'modern', '1605'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd381593', (select id from public.tag_type where url = 'design'), 'Traditional', 'traditional', '1593'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd382368', (select id from public.tag_type where url = 'diameter'), '6″', '6', '2368'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd382370', (select id from public.tag_type where url = 'diameter'), '8″', '8', '2370'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd383010', (select id from public.tag_type where url = 'duct-size'), '5″ (rigid round or equivalent)', '5-rigid-round-or-equivalent', '3010'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd382145', (select id from public.tag_type where url = 'duct-size'), '6″ (rigid round or equivalent)', '6-rigid-round-or-equivalent', '2145'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd383137', (select id from public.tag_type where url = 'duct-size'), 'Input - 8 1/2″ x 3 1/2″, Output - 6″', 'input-8-1-2-x-3-1-2-output-6', '3137'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd382003', (select id from public.tag_type where url = 'duct-size'), 'N/A - Ductless', 'n-a-ductless', '2003'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd381576', (select id from public.tag_type where url = 'exhaust-type'), 'Ducted', 'ducted', '1576'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd381807', (select id from public.tag_type where url = 'exhaust-type'), 'Ducted only', 'ducted-only', '1807'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd382146', (select id from public.tag_type where url = 'exhaust-type'), 'Ducted or Ductless', 'ducted-or-ductless', '2146'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd381687', (select id from public.tag_type where url = 'exhaust-type'), 'Ducted, for Ductless please contact our Customer Service', 'ducted-for-ductless-please-contact-our-customer-service', '1687'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd381577', (select id from public.tag_type where url = 'exhaust-type'), 'Ductless', 'ductless', '1577'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd382004', (select id from public.tag_type where url = 'exhaust-type'), 'Ductless, with high-performance MCIF filter', 'ductless-with-high-performance-mcif-filter', '2004'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd382531', (select id from public.tag_type where url = 'filter-accessories'), 'Chimney Extension', 'extension', '2531'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd382530', (select id from public.tag_type where url = 'filter-accessories'), 'Filter', 'filter', '2530'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd382532', (select id from public.tag_type where url = 'filter-accessories'), 'Lights', 'lights', '2532'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd382533', (select id from public.tag_type where url = 'filter-accessories'), 'Other', 'other', '2533'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd382086', (select id from public.tag_type where url = 'filter-color'), 'Black', 'black', '2086'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd383234', (select id from public.tag_type where url = 'filter-color'), 'Copper', 'copper', '3234'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd383257', (select id from public.tag_type where url = 'filter-color'), 'Glass', 'glass', '3257'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd383233', (select id from public.tag_type where url = 'filter-color'), 'Gold', 'gold', '3233'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd382074', (select id from public.tag_type where url = 'filter-color'), 'SEE ALL', 'see-all', '2074'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd382075', (select id from public.tag_type where url = 'filter-color'), 'Stainless Steel', 'stainless-steel', '2075'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd382076', (select id from public.tag_type where url = 'filter-color'), 'White', 'white', '2076'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd383237', (select id from public.tag_type where url = 'filter-color'), 'Wood', 'wood', '3237'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd382105', (select id from public.tag_type where url = 'filter-design'), 'Air Loop', 'air-loop', '2105'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd382096', (select id from public.tag_type where url = 'filter-design'), 'Built-in', 'built-in', '2096'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd382085', (select id from public.tag_type where url = 'filter-design'), 'Modern', 'modern', '2085'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd383135', (select id from public.tag_type where url = 'filter-design'), 'Recessed ceiling', 'recessed-ceiling', '3135'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd382082', (select id from public.tag_type where url = 'filter-design'), 'SEE ALL', 'see-all', '2082'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd382089', (select id from public.tag_type where url = 'filter-design'), 'Traditional', 'traditional', '2089'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd382106', (select id from public.tag_type where url = 'filter-exhaust-type'), 'Air Loop', 'air-loop', '2106'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd382080', (select id from public.tag_type where url = 'filter-exhaust-type'), 'Ducted', 'ducted', '2080'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd382081', (select id from public.tag_type where url = 'filter-exhaust-type'), 'Ductless', 'ductless', '2081'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd382079', (select id from public.tag_type where url = 'filter-exhaust-type'), 'SEE ALL', 'see-all', '2079'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd382094', (select id from public.tag_type where url = 'filter-material'), 'Copper', 'copper', '2094'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd382084', (select id from public.tag_type where url = 'filter-material'), 'Glass', 'glass', '2084'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd382092', (select id from public.tag_type where url = 'filter-material'), 'Painted steel', 'painted-steel', '2092'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd382077', (select id from public.tag_type where url = 'filter-material'), 'SEE ALL', 'see-all', '2077'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd382078', (select id from public.tag_type where url = 'filter-material'), 'Stainless Steel', 'stainless-steel', '2078'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd382088', (select id from public.tag_type where url = 'filter-material'), 'Wood', 'wood', '2088'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd382000', (select id from public.tag_type where url = 'filter-type'), '2-Stage Filtering System', '2-stage-filtering-system', '2000'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd381808', (select id from public.tag_type where url = 'filter-type'), 'Charcoal, non-rechargeable', 'charcoal-non-rechargeable', '1808'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd381948', (select id from public.tag_type where url = 'filter-type'), 'Combination charcoal metal mesh, reusable/rechargeable.', 'combination-charcoal-metal-mesh-reusable-rechargeable', '1948'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd382144', (select id from public.tag_type where url = 'filter-type'), 'Commercial Baffle (x 2)', 'commercial-baffle-x-2', '2144'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd382150', (select id from public.tag_type where url = 'filter-type'), 'Commercial Baffle (x 3)', 'commercial-baffle-x-3', '2150'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd381785', (select id from public.tag_type where url = 'filter-type'), 'Dishwasher-safe baffle filters', 'dishwasher-safe-baffle-filters', '1785'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd381599', (select id from public.tag_type where url = 'filter-type'), 'Dishwasher-safe, metal mesh filters', 'dishwasher-safe-metal-mesh-filters', '1599'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd381571', (select id from public.tag_type where url = 'filter-type'), 'Perimeter Suction System with dishwasher-safe concealed filters', 'perimeter-suction-system-with-dishwasher-safe-concealed-filters', '1571'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd383270', (select id from public.tag_type where url = 'height-above-cooktop'), '0″ - 6″', '0-6', '3270'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd383204', (select id from public.tag_type where url = 'height-above-cooktop'), '16″ - 18″', '16-18', '3204'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd382520', (select id from public.tag_type where url = 'height-above-cooktop'), '21″', '21', '2520'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd382142', (select id from public.tag_type where url = 'height-above-cooktop'), '26″ - 30″ for best performance', '26-30-for-best-performance', '2142'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd382208', (select id from public.tag_type where url = 'height-above-cooktop'), '28″ - 30″ for best performance', '28-30-for-best-performance', '2208'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd383134', (select id from public.tag_type where url = 'height-above-cooktop'), 'Minimum 28″ to unlimited height', 'minimum-28-to-unlimited-height', '3134'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd381784', (select id from public.tag_type where url = 'item-weight'), '27 lbs', '27-lbs', '1784'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd381932', (select id from public.tag_type where url = 'item-weight'), '28 lbs', '28-lbs', '1932'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd381757', (select id from public.tag_type where url = 'item-weight'), '45 lbs', '45-lbs', '1757'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd381763', (select id from public.tag_type where url = 'item-weight'), '47 lbs', '47-lbs', '1763'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd381766', (select id from public.tag_type where url = 'item-weight'), '48 lbs', '48-lbs', '1766'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd381959', (select id from public.tag_type where url = 'item-weight'), '50 lbs', '50-lbs', '1959'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd381752', (select id from public.tag_type where url = 'item-weight'), '52 lbs', '52-lbs', '1752'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd381775', (select id from public.tag_type where url = 'item-weight'), '60 lbs', '60-lbs', '1775'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd381746', (select id from public.tag_type where url = 'item-weight'), '61 lbs', '61-lbs', '1746'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd381769', (select id from public.tag_type where url = 'item-weight'), '62 lbs', '62-lbs', '1769'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd381743', (select id from public.tag_type where url = 'item-weight'), '67 lbs', '67-lbs', '1743'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd383280', (select id from public.tag_type where url = 'item-weight'), '7 lbs', '7-lbs', '3280'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd383267', (select id from public.tag_type where url = 'lighting-type'), 'Dynamic LED Light (2700K - 5600K)', 'dynamic-led-light-2700k-5600k', '3267'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd383296', (select id from public.tag_type where url = 'lighting-type'), 'Dynamic LED light (2700K-6000K)', 'dynamic-led-light-2700k-6000k', '3296'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd383292', (select id from public.tag_type where url = 'lighting-type'), 'Dynamic LED Light (2700K-6500K)', 'dynamic-led-light-2700k-6500k', '3292'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd381567', (select id from public.tag_type where url = 'lighting-type'), 'Fluorescent', 'fluorescent', '1567'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd381642', (select id from public.tag_type where url = 'lighting-type'), 'Halogen', 'halogen', '1642'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd381595', (select id from public.tag_type where url = 'lighting-type'), 'Incandescent', 'incandescent', '1595'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd381606', (select id from public.tag_type where url = 'lighting-type'), 'LED', 'led', '1606'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd383271', (select id from public.tag_type where url = 'lighting-type'), 'LED (3000K)', 'led-3000k-2', '3271'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd383256', (select id from public.tag_type where url = 'lighting-type'), 'LED (4000K)', 'led-4000k', '3256'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd383223', (select id from public.tag_type where url = 'lighting-type'), 'LED (4700K)', 'led-4700k', '3223'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd383179', (select id from public.tag_type where url = 'lighting-type'), 'LED strip', 'led-strip', '3179'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd383208', (select id from public.tag_type where url = 'lighting-type'), 'LED strip (24V, 10W, 4700K)', 'led-strip-24v-10w-4700k', '3208'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd383215', (select id from public.tag_type where url = 'lighting-type'), 'LED strip (3000K)', 'led-strip-3000k', '3215'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd382143', (select id from public.tag_type where url = 'lighting-type'), 'LED Strip (3200K)', 'led-strip-3200k', '2143'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd383217', (select id from public.tag_type where url = 'lighting-type'), 'LED strip (4000K)', 'led-strip-4000k', '3217'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd383209', (select id from public.tag_type where url = 'lighting-type'), 'LED strip (5000K)', 'led-strip-5000k', '3209'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd383197', (select id from public.tag_type where url = 'lighting-type'), 'LED, 3000K', 'led-3000k', '3197'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd383199', (select id from public.tag_type where url = 'lighting-type'), 'LED, 3200K', 'led-3200k', '3199'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd383196', (select id from public.tag_type where url = 'lighting-type'), 'LED, 4200K', 'led-4200k', '3196'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd383207', (select id from public.tag_type where url = 'lighting-type'), 'LED, adjustable light temperature', 'led-adjustable-light-temperature', '3207'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd383203', (select id from public.tag_type where url = 'lighting-type'), 'Non-dimmable', 'non-dimmable', '3203'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd381984', (select id from public.tag_type where url = 'materials'), 'Black tempered glass', 'black-tempered-glass', '1984'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd381829', (select id from public.tag_type where url = 'materials'), 'Clear tempered glass', 'clear-tempered-glass', '1829'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd382043', (select id from public.tag_type where url = 'materials'), 'Concrete', 'concrete', '2043'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd382032', (select id from public.tag_type where url = 'materials'), 'Copper, white tempered glass', 'copper-white-tempered-glass', '2032'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd383255', (select id from public.tag_type where url = 'materials'), 'Glass', 'glass', '3255'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd381804', (select id from public.tag_type where url = 'materials'), 'Painted steel', 'painted-steel', '1804'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd382029', (select id from public.tag_type where url = 'materials'), 'Painted steel, black tempered glass', 'painted-steel-black-tempered-glass', '2029'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd383210', (select id from public.tag_type where url = 'materials'), 'Painted steel, white tempered glass', 'painted-steel-white-tempered-glass', '3210'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd383136', (select id from public.tag_type where url = 'materials'), 'PVC', 'pvc', '3136'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd381594', (select id from public.tag_type where url = 'materials'), 'Stainless Steel', 'stainless-steel', '1594'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd381810', (select id from public.tag_type where url = 'materials'), 'Stainless steel, aluminum mesh', 'stainless-steel-aluminum-mesh', '1810'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd382104', (select id from public.tag_type where url = 'materials'), 'Stainless Steel, black tempered glass', 'stainless-steel-black-tempered-glass', '2104'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd381566', (select id from public.tag_type where url = 'materials'), 'Stainless Steel, Glass', 'stainless-steel-glass', '1566'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd382101', (select id from public.tag_type where url = 'materials'), 'Stainless Steel, tempered glass', 'stainless-steel-tempered-glass', '2101'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd382025', (select id from public.tag_type where url = 'materials'), 'Stainless steel, white tempered glass', 'stainless-steel-white-tempered-glass', '2025'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd383295', (select id from public.tag_type where url = 'materials'), 'Tempered Glass', 'tempered-glass', '3295'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd383261', (select id from public.tag_type where url = 'materials'), 'Wood (raw ash)', 'wood-raw-ash', '3261'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd382152', (select id from public.tag_type where url = 'max-usable-length'), '134″', '134-inch', '2152'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd381682', (select id from public.tag_type where url = 'mounting-type'), 'Ceiling / Soffit / Cabinet Mount', 'ceiling-soffit-cabinet-mount', '1682'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd381781', (select id from public.tag_type where url = 'mounting-type'), 'In-Cabinet / Under-Cabinet Mount', 'in-cabinet-under-cabinet-mount', '1781'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd381558', (select id from public.tag_type where url = 'mounting-type'), 'Island (Ceiling) Mount', 'island-ceiling-mount', '1558'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd381587', (select id from public.tag_type where url = 'mounting-type'), 'Wall Mount', 'wall-mount', '1587'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd381694', (select id from public.tag_type where url = 'mounting-type'), 'Wall Mount or Undercabinet Mount', 'wall-mount-or-undercabinet-mount', '1694'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd381560', (select id from public.tag_type where url = 'category'), 'Wall', 'wall', '1660'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd381561', (select id from public.tag_type where url = 'category'), 'Island', 'island', '1661'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd381574', (select id from public.tag_type where url = 'noise-level'), '0.5 - 3.2 sones', '0-5-3-2-sones', '1574'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd381596', (select id from public.tag_type where url = 'of-lights'), '1', '1', '1596'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd381953', (select id from public.tag_type where url = 'of-lights'), '12', '12', '1953'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd381973', (select id from public.tag_type where url = 'of-lights'), '1x full-width LED light strip', '1x-full-width-led-light-strip', '1973'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd383246', (select id from public.tag_type where url = 'of-lights'), '1x LED light strip', '1x-led-light-strip', '3246'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd383216', (select id from public.tag_type where url = 'of-lights'), '1x perimeter LED light strip', '1x-perimeter-led-light-strip', '3216'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd383214', (select id from public.tag_type where url = 'of-lights'), '1x perimeter LED light strip (3000K)', '1x-perimeter-led-light-strip-3000k', '3214'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd381568', (select id from public.tag_type where url = 'of-lights'), '2', '2', '1568'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd381991', (select id from public.tag_type where url = 'of-lights'), '2x full-width LED light strips', '2x-full-width-led-light-strips', '1991'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd383225', (select id from public.tag_type where url = 'of-lights'), '2x LED light strips', '2x-led-light-strips', '3225'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd382010', (select id from public.tag_type where url = 'of-lights'), '2x worklights + dimmable internal light', '2x-worklights-dimmable-internal-light', '2010'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd381616', (select id from public.tag_type where url = 'of-lights'), '3', '3', '1616'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd383194', (select id from public.tag_type where url = 'of-lights'), '3 work lights + ambient light ring', '3-work-lights-ambient-light-ring', '3194'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd381607', (select id from public.tag_type where url = 'of-lights'), '4', '4', '1607'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd382048', (select id from public.tag_type where url = 'of-lights'), '4 work lights + internal light', '4-work-lights-internal-light', '2048'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd383224', (select id from public.tag_type where url = 'of-lights'), '4x LED light strips', '4x-led-light-strips', '3224'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd382022', (select id from public.tag_type where url = 'of-lights'), '4x work lights + 2x ambient lights', '4x-work-lights-2x-ambient-lights', '2022'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd382026', (select id from public.tag_type where url = 'of-lights'), '4x work lights + ambient light ring', '4x-work-lights-ambient-light-ring', '2026'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd382017', (select id from public.tag_type where url = 'of-lights'), '4x work lights + dimmable internal light', '4x-work-lights-dimmable-internal-light', '2017'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd382006', (select id from public.tag_type where url = 'of-lights'), '4x worklights + 2x accent lights', '4x-worklights-2x-accent-lights', '2006'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd383272', (select id from public.tag_type where url = 'of-lights'), '5', '5', '3272'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd381921', (select id from public.tag_type where url = 'of-lights'), '6', '6', '1921'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd381690', (select id from public.tag_type where url = 'of-lights'), '8', '8', '1690'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd381999', (select id from public.tag_type where url = 'of-lights'), '8 total - 6x worklights + 2x accent lights', '8-total-6x-worklights-2x-accent-lights', '1999'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd382053', (select id from public.tag_type where url = 'of-lights'), 'Illuminated perimeter glass panel', 'illuminated-perimeter-glass-panel', '2053'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd381597', (select id from public.tag_type where url = 'of-speeds'), '3', '3', '1597'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd381569', (select id from public.tag_type where url = 'of-speeds'), '4', '4', '1569'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd381581', (select id from public.tag_type where url = 'order-processing-time'), 'Same day shipping', 'same-day-shipping', '1581'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd382013', (select id from public.tag_type where url = 'power-requirements'), 'Direct Wire Connection, US/Canada 110-120 Volts, 60 Hz, 3 Amp', 'direct-wire-connection-us-canada-110-120-volts-60-hz-3-amp', '2013'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd381949', (select id from public.tag_type where url = 'power-requirements'), 'E-14 (Euro Base), AC 110-140V', 'e-14-euro-base-ac-110-140v', '1949'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd381823', (select id from public.tag_type where url = 'power-requirements'), 'T5, 21W', 't5-21w', '1823'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd381967', (select id from public.tag_type where url = 'power-requirements'), 'T5, 28W', 't5-28w', '1967'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd381578', (select id from public.tag_type where url = 'power-requirements'), 'US Standard 3-pin plug, 110-120 Volts, 60 Hz, 3 Amp', 'us-standard-3-pin-plug-110-120-volts-60-hz-3-amp', '1578'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd381822', (select id from public.tag_type where url = 'power-requirements'), 'Voltage: 12V, Power: 20W, Base Type: G4', 'voltage-12v-power-20w-base-type-g4', '1822'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd382265', (select id from public.tag_type where url = 'recommended-range-width'), 'Up to 24″', 'up-to-24', '2265'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd382255', (select id from public.tag_type where url = 'recommended-range-width'), 'Up to 30 ″', 'up-to-30-', '2255'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd381591', (select id from public.tag_type where url = 'recommended-range-width'), 'Up to 36&Prime;', 'up-to-36', '1591'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd382213', (select id from public.tag_type where url = 'recommended-range-width'), 'Up to 36″', 'up-to-36', '2213'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd382207', (select id from public.tag_type where url = 'recommended-range-width'), 'Up to 48 ″', 'up-to-48-', '2207'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd383132', (select id from public.tag_type where url = 'recommended-range-width'), 'Up to 48″', 'up-to-48', '3132'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd381582', (select id from public.tag_type where url = 'shipping-speed'), 'Estimated delivery time within continental 48 states: 1-7 business days', 'estimated-delivery-time-within-continental-48-states-1-7-business-days', '1582'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd381620', (select id from public.tag_type where url = 'shipping-weight'), '105 lbs', '105-lbs', '1620'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd381805', (select id from public.tag_type where url = 'shipping-weight'), '15 lbs', '15-lbs', '1805'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd381786', (select id from public.tag_type where url = 'shipping-weight'), '32 lbs', '32-lbs', '1786'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd381933', (select id from public.tag_type where url = 'shipping-weight'), '33 lbs', '33-lbs', '1933'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd381764', (select id from public.tag_type where url = 'shipping-weight'), '52 lbs', '52-lbs', '1764'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd381767', (select id from public.tag_type where url = 'shipping-weight'), '53 lbs', '53-lbs', '1767'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd381753', (select id from public.tag_type where url = 'shipping-weight'), '57 lbs', '57-lbs', '1753'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd381758', (select id from public.tag_type where url = 'shipping-weight'), '65 lbs', '65-lbs', '1758'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd381770', (select id from public.tag_type where url = 'shipping-weight'), '67 lbs', '67-lbs', '1770'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd381960', (select id from public.tag_type where url = 'shipping-weight'), '78 lbs', '78-lbs', '1960'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd381776', (select id from public.tag_type where url = 'shipping-weight'), '80 lbs', '80-lbs', '1776'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd381747', (select id from public.tag_type where url = 'shipping-weight'), '81 lbs', '81-lbs', '1747'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd381744', (select id from public.tag_type where url = 'shipping-weight'), '87 lbs', '87-lbs', '1744'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd381583', (select id from public.tag_type where url = 'ships-via'), 'Freight', 'freight', '1583'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd381619', (select id from public.tag_type where url = 'ships-via'), 'Ground', 'ground', '1619'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd381580', (select id from public.tag_type where url = 'warranty'), '1-year US standard + 2-year parts', '1-year-us-standard-2-year-parts', '1580'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd381664', (select id from public.tag_type where url = 'width-group'), '14″-24″', 'up-to-24', '1664'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd381653', (select id from public.tag_type where url = 'width-group'), '25″-30″', 'up-to-30', '1653'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd381600', (select id from public.tag_type where url = 'width-group'), '32″-40″', 'up-to-36', '1600'),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd381585', (select id from public.tag_type where url = 'width-group'), '42″-84″', 'up-to-48', '1585'),
       ('a0eebc98-9c0b-4ef8-bb6d-5bb9bd381500', (select id from public.tag_type where url = 'category'), 'Island range hoods', 'island', '1'),
       ('a0eebc98-9c0b-4ef8-bb6d-5bb9bd381501', (select id from public.tag_type where url = 'category'), 'Wall range hoods', 'wall', '2'),
       ('a0eebc98-9c0b-4ef8-bb6d-5bb9bd381502', (select id from public.tag_type where url = 'category'), 'Air loop range hoods', 'ait-loop', '3'),
       ('a0eebc98-9c0b-4ef8-bb6d-5bb9bd381503', (select id from public.tag_type where url = 'category'), 'Built-in range hoods', 'built-in', '4'),
       ('a0eebc98-9c0b-4ef8-bb6d-5bb9bd381504', (select id from public.tag_type where url = 'category'), 'Under Cabinet range hoods', 'under-cabinet', '5'),
       ('a0eebc98-9c0b-4ef8-bb6d-5bb9bd381505', (select id from public.tag_type where url = 'category'), 'Accessories', 'accessories', '6'),
       ('a0eebc98-9c0b-4ef8-bb6d-5bb9bd381506', (select id from public.tag_type where url = 'shipping-class'), 'Freight', 'freight', '1'),
       ('a0eebc98-9c0b-4ef8-bb6d-5bb9bd381507', (select id from public.tag_type where url = 'shipping-class'), 'Oversize Freight ', 'oversize-freight', '2'),
       ('a0eebc98-9c0b-4ef8-bb6d-5bb9bd381508', (select id from public.tag_type where url = 'shipping-class'), 'Ground', 'ground', '3'),
       ('a0eebc98-9c0b-4ef8-bb6d-5bb9bd381509', (select id from public.tag_type where url = 'status'), 'Public', 'public', '1'),
       ('a0eebc98-9c0b-4ef8-bb6d-5bb9bd381509', (select id from public.tag_type where url = 'status'), 'Private', 'private', '2'),
       ('a0eebc98-9c0b-4ef8-bb6d-5bb9bd381509', (select id from public.tag_type where url = 'status'), 'Out of stock', 'out-of-stock', '3'),
       ('a0eebc98-9c0b-4ef8-bb6d-5bb9bd381509', (select id from public.tag_type where url = 'status'), 'Discontinued', 'discontinued', '4'),
       ('a0eebc98-9c0b-4ef8-bb6d-5bb9bd381509', (select id from public.tag_type where url = 'status'), 'Archived', 'archived', '5'),




;

-- get data
select *
from public.tag_select;

-- get data form WooCommerce DB
-- SELECT
--     CONCAT('a0eebc99-9c0b-4ef8-bb6d-6bb9bd38', tt.term_id) as id,
--     CONCAT('(select id from public.tag_type where url = ''', SUBSTRING(tt.taxonomy FROM 4), ''')') as specification_id,
--     t.name as name,
--     t.url as url,
--     tt.term_id as sort_order
-- FROM wp_posts p
--          JOIN wp_term_relationships tr ON p.ID = tr.object_id
--          JOIN wp_term_taxonomy tt ON tr.term_taxonomy_id = tt.term_taxonomy_id AND tt.taxonomy LIKE 'pa_%'
--          JOIN wp_terms t ON tt.term_id = t.term_id
-- WHERE p.post_type = 'product'
-- GROUP BY
--     tt.taxonomy, t.name
-- ORDER BY
--     tt.taxonomy;