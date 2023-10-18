CREATE TABLE IF NOT EXISTS public.category
(
    id                UUID UNIQUE   DEFAULT uuid_generate_v4(),
    name              VARCHAR(255)                NOT NULL,
    slug              VARCHAR(255) UNIQUE         NOT NULL,
    short_description VARCHAR(255)  DEFAULT NULL,
    description       VARCHAR(4000) DEFAULT NULL,
    "order"           INTEGER       DEFAULT NULL,
    active            BOOLEAN       DEFAULT TRUE,

    created_at        TIMESTAMP     default NOW() NOT NULL,
    created_by        UUID          DEFAULT NULL REFERENCES public.user (id),
    updated_at        TIMESTAMP     default NOW() NOT NULL,
    updated_by        UUID          DEFAULT NULL REFERENCES public.user (id),

    CONSTRAINT category_pkey PRIMARY KEY (id)
);

CREATE UNIQUE INDEX category_id ON public.category (id);
CREATE UNIQUE INDEX category_slug ON public.category (slug);
ALTER TABLE public.category
    OWNER TO postgres;
COMMENT ON TABLE public.category IS 'Product categories';
-- insert data
INSERT INTO public.category (id, slug, name, short_description, description, "order")
VALUES ('1f484cda-c00e-4ed8-a325-9c5e035f9920', 'island-range-hoods', 'Island Range Hoods', 'Some text', 'Some text ',
        1),
       ('1f484cda-c00e-4ed8-a325-9c5e035f9921', 'wall-range-hoods', 'Wall range hoods', 'Some text', 'Some text', 2),
       ('1f484cda-c00e-4ed8-a325-9c5e035f9922', 'built-in-range-hoods', 'Built-in Range Hoods', 'Some text',
        'Some text', 3),
       ('1f484cda-c00e-4ed8-a325-9c5e035f9923', 'range-hood-accessories', 'Accessories', 'Some text', 'Some text', 4),
       ('1f484cda-c00e-4ed8-a325-9c5e035f9924', 'discontinued-range-hoods', 'Discontinued', 'Some text', 'Some text',
        5);

