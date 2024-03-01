CREATE TABLE IF NOT EXISTS public.product_category
(
    id                UUID UNIQUE   DEFAULT uuid_generate_v4(),
    name              VARCHAR(255)                NOT NULL,
    url              VARCHAR(255) UNIQUE         NOT NULL,
    short_description VARCHAR(255)  DEFAULT ''    NOT NULL,
    description       VARCHAR(4000) DEFAULT ''    NOT NULL,
    sort_order        INTEGER       DEFAULT NULL,
    prime             BOOLEAN       DEFAULT TRUE,
    active            BOOLEAN       DEFAULT TRUE,

    created_at        TIMESTAMP     default NOW() NOT NULL,
    created_by        UUID          DEFAULT '0e95efda-f9e2-4fac-8184-3ce2e8b7e0e1' REFERENCES public.user (id),
    updated_at        TIMESTAMP     default NOW() NOT NULL,
    updated_by        UUID          DEFAULT '0e95efda-f9e2-4fac-8184-3ce2e8b7e0e1' REFERENCES public.user (id),

    CONSTRAINT product_category_pkey PRIMARY KEY (id)
);

CREATE INDEX category_id ON public.product_category (id);
CREATE INDEX category_url ON public.product_category (url);
ALTER TABLE public.product_category
    OWNER TO postgres;
COMMENT ON TABLE public.product_category IS 'Product categories';

-- auto update updated_at
CREATE TRIGGER product_category_updated_at
    BEFORE UPDATE
    ON public.user
    FOR EACH ROW
EXECUTE FUNCTION update_update_at_column();

-- auto set sort_order column
CREATE OR REPLACE FUNCTION set_order_column_to_product_category()
    RETURNS TRIGGER AS
$$
BEGIN
    NEW.sort_order = (SELECT COALESCE(MAX(sort_order), 0) + 1 FROM product_category);
    RETURN NEW;
END;
$$ language 'plpgsql';
CREATE TRIGGER product_category_order
    BEFORE INSERT
    ON public.product_category
    FOR EACH ROW
EXECUTE FUNCTION set_order_column_to_product_category();

-- insert data
INSERT INTO public.product_category (id, url, name, short_description, description)
VALUES ('1f484cda-c00e-4ed8-a325-9c5e035f9901', 'island', 'Island Range Hoods', 'Some text', 'Some text'),
       ('1f484cda-c00e-4ed8-a325-9c5e035f9902', 'wall', 'Wall range hoods', 'Some text', 'Some text'),
       ('1f484cda-c00e-4ed8-a325-9c5e035f9903', 'ait-loop', 'Air loop range hoods', 'Some text', 'Some text'),
       ('1f484cda-c00e-4ed8-a325-9c5e035f9904', 'built-in', 'Built-in Range Hoods', 'Some text', 'Some text'),
       ('1f484cda-c00e-4ed8-a325-9c5e035f9905', 'under-cabinet', 'Under Cabinet Range hood', 'Some text', 'Some text'),
       ('1f484cda-c00e-4ed8-a325-9c5e035f9906', 'accessories', 'Accessories', 'Some text', 'Some text'),
       ('1f484cda-c00e-4ed8-a325-9c5e035f9907', 'black', 'Black Range hood', 'Some text', 'Some text'),
       ('1f484cda-c00e-4ed8-a325-9c5e035f9908', 'white', 'White Range hood', 'Some text', 'Some text'),
       ('1f484cda-c00e-4ed8-a325-9c5e035f9909', 'wood', 'Wood Range hood', 'Some text', 'Some text'),
       ('1f484cda-c00e-4ed8-a325-9c5e035f9910', 'stainless-steel', 'Stainless Steel Range hood', 'Some text',
        'Some text'),
       ('1f484cda-c00e-4ed8-a325-9c5e035f9911', 'glass', 'Glass Range hood', 'Some text', 'Some text'),
       ('1f484cda-c00e-4ed8-a325-9c5e035f9912', 'perimeter-filter', 'Perimeter Filter Range Hoods', 'Some text',
        'Some text'),
       ('1f484cda-c00e-4ed8-a325-9c5e035f9913', 'murano', 'Murano Range Hoods', 'Some text', 'Some text'),
       ('1f484cda-c00e-4ed8-a325-9c5e035f9914', 'ductless-range-hoods', 'Ductless Range hood', 'Some text',
        'Some text'),
       ('1f484cda-c00e-4ed8-a325-9c5e035f9915', 'ducted-range-hoods', 'Ducted Range hood', 'Some text', 'Some text'),
       ('1f484cda-c00e-4ed8-a325-9c5e035f9999', 'discontinued-range-hoods', 'Discontinued', 'Some text', 'Some text');

-- get data
select *
from public.product_category;
