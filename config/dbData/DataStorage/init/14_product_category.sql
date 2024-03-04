CREATE TABLE IF NOT EXISTS public.product_category
(
    id                UUID UNIQUE   DEFAULT uuid_generate_v4(),
    name              VARCHAR(255)                NOT NULL,
    url               VARCHAR(255) UNIQUE         NOT NULL,
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

-- permissions
ALTER TABLE public.product_category
    OWNER TO postgres;

-- indexes
CREATE INDEX category_id ON public.product_category (id);
CREATE INDEX category_url ON public.product_category (url);
CREATE INDEX category_sort_order ON public.product_category (sort_order);
CREATE INDEX category_updated_at ON public.product_category (updated_at);

-- comments
COMMENT ON TABLE public.product_category IS 'Product categories';
COMMENT ON COLUMN public.product_category.id IS 'Unique identifier';
COMMENT ON COLUMN public.product_category.name IS 'Name';
COMMENT ON COLUMN public.product_category.url IS 'URL';
COMMENT ON COLUMN public.product_category.short_description IS 'Short description';
COMMENT ON COLUMN public.product_category.description IS 'Description';
COMMENT ON COLUMN public.product_category.sort_order IS 'Sort order';
COMMENT ON COLUMN public.product_category.prime IS 'Prime';
COMMENT ON COLUMN public.product_category.active IS 'Active';
COMMENT ON COLUMN public.product_category.created_at IS 'Created at';
COMMENT ON COLUMN public.product_category.created_by IS 'Created by';
COMMENT ON COLUMN public.product_category.updated_at IS 'Updated at';
COMMENT ON COLUMN public.product_category.updated_by IS 'Updated by';

-- auto update updated_at
CREATE TRIGGER product_category_updated_at
    BEFORE UPDATE
    ON public.user
    FOR EACH ROW
EXECUTE FUNCTION update_update_at_column();

-- insert data
INSERT INTO public.product_category (id, url, name, short_description, description, prime, active)
VALUES ('1f484cda-c00e-4ed8-a325-9c5e035f9901', 'island', 'Island range hoods', 'Some text', 'Some text', true, true),
       ('1f484cda-c00e-4ed8-a325-9c5e035f9902', 'wall', 'Wall range hoods', 'Some text', 'Some text', true, true),
       ('1f484cda-c00e-4ed8-a325-9c5e035f9903', 'ait-loop', 'Air loop range hoods', 'Some text', 'Some text', true, true),
       ('1f484cda-c00e-4ed8-a325-9c5e035f9904', 'built-in', 'Built-in range hoods', 'Some text', 'Some text', true, true),
       ('1f484cda-c00e-4ed8-a325-9c5e035f9905', 'under-cabinet', 'Under Cabinet range hoods', 'Some text', 'Some text', true, true),
       ('1f484cda-c00e-4ed8-a325-9c5e035f9906', 'accessories', 'Accessories', 'Some text', 'Some text', true, true),
       ('1f484cda-c00e-4ed8-a325-9c5e035f9907', 'black', 'Black range hoods', 'Some text', 'Some text', false, true),
       ('1f484cda-c00e-4ed8-a325-9c5e035f9908', 'white', 'White range hoods', 'Some text', 'Some text', false, true),
       ('1f484cda-c00e-4ed8-a325-9c5e035f9909', 'wood', 'Wood range hoods', 'Some text', 'Some text', false, true),
       ('1f484cda-c00e-4ed8-a325-9c5e035f9910', 'stainless-steel', 'Stainless Steel range hoods', 'Some text', 'Some text', false, true),
       ('1f484cda-c00e-4ed8-a325-9c5e035f9911', 'glass', 'Glass range hoods', 'Some text', 'Some text', false, true),
       ('1f484cda-c00e-4ed8-a325-9c5e035f9912', 'perimeter-filter', 'Perimeter Filter range hoods', 'Some text', 'Some text', false, true),
       ('1f484cda-c00e-4ed8-a325-9c5e035f9913', 'murano', 'Murano range hoods', 'Some text', 'Some text', false, true),
       ('1f484cda-c00e-4ed8-a325-9c5e035f9914', 'ductless', 'Ductless range hoods', 'Some text', 'Some text', false, true),
       ('1f484cda-c00e-4ed8-a325-9c5e035f9915', 'ducted', 'Ducted range hoods', 'Some text', 'Some text', false, true),
       ('1f484cda-c00e-4ed8-a325-9c5e035f9999', 'discontinued', 'Discontinued', 'Some text', 'Some text', false, false);

-- get data
select *
from public.product_category;
