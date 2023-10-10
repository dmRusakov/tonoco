-- drop if needed
DROP TABLE IF EXISTS public.product;

-- create table
CREATE TABLE IF NOT EXISTS public.product
(
    product_id              uuid unique                   DEFAULT uuid_generate_v4(),
    sku                     character varying(100) unique default null,
    name                    character varying(255) COLLATE pg_catalog."default" NOT NULL,
    short_description       character varying(255)        default null,
    description             character varying(4000)       default null,
    "order"                 integer                       default null,

    regular_price           REAL                          default null,
    sale_price              REAL                          default null,
    factory_price           REAL                          default null,

    tax_status_id           integer                       default null,
    tax_class_id            integer                       default null,

    quantity                integer                       default null,
    return_to_stock_date    timestamp                     default null,
    is_track_stock          boolean                       default null,
    is_visible              boolean                       default null,
    is_discontinued         boolean                       default null,

    shipping_class_id       integer                       default null,
    shipping_weight         REAL                          default null,
    shipping_width          REAL                          default null,
    shipping_height         REAL                          default null,
    shipping_length         REAL                          default null,

    purchase_note_id        integer                       default null,
    is_enable_for_reviews   boolean                       default null,

    seo_title               character varying(100)        default null,
    seo_description         character varying(4000)       default null,
    gtin                    character varying(50)         default null,
    google_product_category character varying(50)         default null,
    google_product_type     character varying(100)        default null,
    language                character varying(2)          default null,

    created_at              timestamp                     default now(),
    created_by              uuid                          default null REFERENCES public.user (user_id),
    updated_at              timestamp                     default now(),
    updated_by              uuid                          default null REFERENCES public.user (user_id)
);
ALTER TABLE ONLY public.product ADD CONSTRAINT product_pkey PRIMARY KEY (product_id);
ALTER TABLE public.product OWNER TO postgres;
CREATE UNIQUE INDEX product_id ON public.product (product_id);
CREATE UNIQUE INDEX sku ON public.product (sku);

-- add demo data to product table
INSERT INTO public.product (product_id, sku, name, short_description, description, "order", regular_price, sale_price,
                            factory_price, tax_status_id, tax_class_id, quantity, return_to_stock_date, is_track_stock,
                            is_visible, is_discontinued, shipping_class_id, shipping_weight, shipping_width,
                            shipping_height, shipping_length, purchase_note_id, is_enable_for_reviews, seo_title,
                            seo_description, gtin, google_product_category, google_product_type, language)
VALUES ('c5506e76-01f4-4b8a-a066-e3205986a813', 'sku1', 'name1', 'short_description1', 'description1', 1, 1.1, 1.2, 1.3, 1, 1, 1, '2019-01-01 00:00:00', true, true, true, 1, 1.1, 1.2, 1.3, 1.4, 1, true, 'seo_title1', 'seo_description1', 'gtin1', 'google_product_category1', 'google_product_type1', 'en'),
       ('f18fcc54-4b42-4511-a827-b6cc3e48fa8d', 'sku2', 'name2', 'short_description2', 'description2', 2, 2.1, 2.2, 2.3, 2, 2, 2, '2019-01-02 00:00:00', false, false, false, 2, 2.1, 2.2, 2.3, 2.4, 2, false, 'seo_title2', 'seo_description2', 'gtin2', 'google_product_category2', 'google_product_type2', 'en'),
       ('01879ba4-14e2-4198-9217-8c86bd4d3717', 'sku3', 'name3', 'short_description3', 'description3', 3, 3.1, 3.2, 3.3, 3, 3, 3, '2019-01-03 00:00:00', true, true, true, 3, 3.1, 3.2, 3.3, 3.4, 3, true, 'seo_title3', 'seo_description3', 'gtin3', 'google_product_category3', 'google_product_type3', 'en'),
       ('0b0b1b01-f61b-4be2-9439-80be7e00a476', 'sku4', 'name4', 'short_description4', 'description4', 4, 4.1, 4.2, 4.3, 4, 4, 4, '2019-01-04 00:00:00', false, false, false, 4, 4.1, 4.2, 4.3, 4.4, 4, false, 'seo_title4', 'seo_description4', 'gtin4', 'google_product_category4', 'google_product_type4', 'en'),
       ('5cb12adf-d8de-4f8b-a04f-88a812c7c1a5', 'sku5', 'name5', 'short_description5', 'description5', 5, 5.1, 5.2, 5.3, 5, 5, 5, '2019-01-05 00:00:00', true, true, true, 5, 5.1, 5.2, 5.3, 5.4, 5, true, 'seo_title5', 'seo_description5', 'gtin5', 'google_product_category5', 'google_product_type5', 'en'),
       ('29d7499a-b223-44e0-acf2-3b4cd82e1af4', 'sku6', 'name6', 'short_description6', 'description6', 6, 6.1, 6.2, 6.3, 6, 6, 6, '2019-01-06 00:00:00', false, false, false, 6, 6.1, 6.2, 6.3, 6.4, 6, false, 'seo_title6', 'seo_description6', 'gtin6', 'google_product_category6', 'google_product_type6', 'en'),
       ('6b564ffb-4b14-4fb5-a904-25e66abeb294', 'sku7', 'name7', 'short_description7', 'description7', 7, 7.1, 7.2, 7.3, 7, 7, 7, '2019-01-07 00:00:00', true, true, true, 7, 7.1, 7.2, 7.3, 7.4, 7, true, 'seo_title7', 'seo_description7', 'gtin7', 'google_product_category7', 'google_product_type7', 'en');