-- drop table if exists
DROP TABLE IF EXISTS public.product_image CASCADE;

-- create table
CREATE TABLE IF NOT EXISTS public.product_image
(
    id             UUID UNIQUE   DEFAULT uuid_generate_v4(),
    product_id     UUID          DEFAULT NULL,
    image_id       UUID          DEFAULT NULL,

    type           VARCHAR(255)  DEFAULT NULL,
    sort_order     INTEGER       DEFAULT NULL,

    created_at     TIMESTAMP     DEFAULT NOW() NOT NULL,
    created_by     UUID          DEFAULT NULL,
    updated_at     TIMESTAMP     DEFAULT NOW() NOT NULL,
    updated_by     UUID          DEFAULT NULL,

    CONSTRAINT product_image_pkey PRIMARY KEY (id)
);

-- ownership and index
ALTER TABLE public.product_image OWNER TO postgres;
CREATE INDEX product_image_id ON public.product_image USING btree (id);
CREATE INDEX product_image_product_id ON public.product_image USING btree (product_id);
CREATE INDEX product_image_image_id ON public.product_image USING btree (image_id);
CREATE INDEX product_image_type ON public.product_image USING btree (type);

-- comment on table
COMMENT ON TABLE public.product_image IS 'Reference table for product image';
COMMENT ON COLUMN public.product_image.id IS 'Unique identifier for product image';
COMMENT ON COLUMN public.product_image.product_id IS 'Product id of product image';
COMMENT ON COLUMN public.product_image.image_id IS 'Image id of product image';
COMMENT ON COLUMN public.product_image.type IS 'Type of product image';
COMMENT ON COLUMN public.product_image.sort_order IS 'Sort order of product image';
COMMENT ON COLUMN public.product_image.created_at IS 'Creation time of product image';
COMMENT ON COLUMN public.product_image.created_by IS 'Creator of product image';
COMMENT ON COLUMN public.product_image.updated_at IS 'Update time of product image';
COMMENT ON COLUMN public.product_image.updated_by IS 'Updater of product image';


-- auto update updated_at
CREATE OR REPLACE TRIGGER product_image_updated_at
    BEFORE UPDATE OR INSERT
    ON public.product_image
    FOR EACH ROW
EXECUTE FUNCTION update_update_at_column();

-- auto set created_by
CREATE OR REPLACE TRIGGER product_image_created_by
    BEFORE INSERT
    ON public.product_image
    FOR EACH ROW
EXECUTE FUNCTION set_created_by_if_null();

-- auto set updated_by
CREATE OR REPLACE TRIGGER product_image_updated_by
    BEFORE INSERT OR UPDATE
    ON public.product_image
    FOR EACH ROW
EXECUTE FUNCTION set_updated_by_if_null();


select * from public.product_image;

-- SELECT * FROM
--     (-- thumbnail
--         SELECT
--             product_sku.meta_value AS sku,
--             pm.meta_value          AS image_id,
--             'main' as type,
--             0 as sort_order
--         FROM wp_posts product
--                  JOIN wp_postmeta pm ON product.ID = pm.post_id
--                  JOIN wp_postmeta product_sku ON product.ID = product_sku.post_id
--                  JOIN wp_posts imageInfo ON pm.meta_value = imageInfo.ID
--         WHERE product.post_type = 'product'
--           AND pm.meta_key = '_thumbnail_id'
--           AND product_sku.meta_key = '_sku'
--           AND pm.meta_value IS NOT NULL
--           AND pm.meta_value != ''
--         UNION ALL
-- -- gallery
--         SELECT
--             product_sku.meta_value AS sku,
--             SUBSTRING_INDEX(SUBSTRING_INDEX(pm.meta_value, ',', numbers.n), ',', -1) AS image_id,
--             'gallery' as type,
--             numbers.n - 1 AS sort_order
--         FROM wp_posts product_gall
--                  JOIN wp_postmeta pm ON product_gall.ID = pm.post_id
--                  JOIN wp_postmeta product_sku ON product_gall.ID = product_sku.post_id
--                  JOIN (
--             SELECT 1 n UNION ALL SELECT 2 UNION ALL SELECT 3 UNION ALL SELECT 4 UNION ALL SELECT 5
--             UNION ALL SELECT 6 UNION ALL SELECT 7 UNION ALL SELECT 8 UNION ALL SELECT 9 UNION ALL SELECT 10
--             UNION ALL SELECT 11 UNION ALL SELECT 12 UNION ALL SELECT 13 UNION ALL SELECT 14 UNION ALL SELECT 15
--             UNION ALL SELECT 16 UNION ALL SELECT 17 UNION ALL SELECT 18 UNION ALL SELECT 19 UNION ALL SELECT 20
--             UNION ALL SELECT 21 UNION ALL SELECT 22 UNION ALL SELECT 23 UNION ALL SELECT 24 UNION ALL SELECT 25
--             UNION ALL SELECT 26 UNION ALL SELECT 27 UNION ALL SELECT 28 UNION ALL SELECT 29 UNION ALL SELECT 30
--         ) numbers ON CHAR_LENGTH(pm.meta_value) - CHAR_LENGTH(REPLACE(pm.meta_value, ',', '')) >= numbers.n - 1
--         WHERE product_gall.post_type = 'product'
--           AND pm.meta_key = '_product_image_gallery'
--           AND product_sku.meta_key = '_sku'
--           AND pm.meta_value IS NOT NULL
--           AND pm.meta_value != ''
--         UNION ALL
-- -- dimension
--         SELECT
--             product_sku.meta_value AS sku,
--             pm.meta_value AS image_id,
--             'dimension' as type,
--             0 as sort_order
--         FROM wp_posts product
--                  JOIN wp_postmeta pm ON product.ID = pm.post_id
--                  JOIN wp_postmeta product_sku ON product.ID = product_sku.post_id
--         WHERE product.post_type = 'product'
--           AND pm.meta_key = 'product_dimensions'
--           AND product_sku.meta_key = '_sku'
--           AND pm.meta_value IS NOT NULL
--           AND pm.meta_value != ''
--         UNION ALL
-- -- hover
--         SELECT
--             product_sku.meta_value AS sku,
--             pm.meta_value AS image_id,
--             'hover' as type,
--             0 as sort_order
--         FROM wp_posts product
--                  JOIN wp_postmeta pm ON product.ID = pm.post_id
--                  JOIN wp_postmeta product_sku ON product.ID = product_sku.post_id
--         WHERE product.post_type = 'product'
--           AND pm.meta_key = 'hover_image'
--           AND product_sku.meta_key = '_sku'
--           AND pm.meta_value IS NOT NULL
--           AND pm.meta_value != ''
--         UNION ALL
-- -- feed_general_image
--         SELECT
--             product_sku.meta_value AS sku,
--             pm.meta_value AS image_id,
--             'feed_general_image' as type,
--             0 as sort_order
--         FROM wp_posts product
--                  JOIN wp_postmeta pm ON product.ID = pm.post_id
--                  JOIN wp_postmeta product_sku ON product.ID = product_sku.post_id
--         WHERE product.post_type = 'product'
--           AND pm.meta_key = 'general_feed-image'
--           AND product_sku.meta_key = '_sku'
--           AND pm.meta_value IS NOT NULL
--           AND pm.meta_value != ''
--         UNION ALL
-- -- feed_dimensions
--         SELECT
--             product_sku.meta_value AS sku,
--             pm.meta_value AS image_id,
--             'feed_dimensions' as type,
--             0 as sort_order
--         FROM wp_posts product
--                  JOIN wp_postmeta pm ON product.ID = pm.post_id
--                  JOIN wp_postmeta product_sku ON product.ID = product_sku.post_id
--         WHERE product.post_type = 'product'
--           AND pm.meta_key = 'feed-dimensions'
--           AND product_sku.meta_key = '_sku'
--           AND pm.meta_value IS NOT NULL
--           AND pm.meta_value != ''
--         UNION ALL
-- -- feed_lifestyle_image
--         SELECT
--             product_sku.meta_value AS sku,
--             pm.meta_value AS image_id,
--             'feed_lifestyle_image' as type,
--             0 as sort_order
--         FROM wp_posts product
--                  JOIN wp_postmeta pm ON product.ID = pm.post_id
--                  JOIN wp_postmeta product_sku ON product.ID = product_sku.post_id
--         WHERE product.post_type = 'product'
--           AND pm.meta_key = 'feed_lifestyle_image'
--           AND product_sku.meta_key = '_sku'
--           AND pm.meta_value IS NOT NULL
--           AND pm.meta_value != ''
-- -- feed_gallery
--         UNION ALL
--         SELECT
--             sku,
--             value,
--             CAST((sort_order / 2 - 0.5) AS SIGNED) AS sort_order,
--             'feed_gallery' as type
--         FROM
--             (WITH RECURSIVE numbers AS (
--                 SELECT 1 AS n
--                 UNION ALL
--                 SELECT n + 1
--                 FROM numbers
--                 WHERE n < 100
--             )
--              SELECT
--                  product_sku.meta_value AS sku,
--                  SUBSTRING_INDEX(SUBSTRING_INDEX(pm.meta_value, '"', numbers.n), 's:5:"', -1) AS value,
--                  numbers.n - 1 AS sort_order
--              FROM wp_postmeta pm
--                       JOIN wp_postmeta product_sku ON pm.post_id = product_sku.post_id
--                       JOIN numbers ON CHAR_LENGTH(pm.meta_value) - CHAR_LENGTH(REPLACE(pm.meta_value, '"', '')) >= numbers.n - 1
--              WHERE pm.meta_key = 'feed-image'
--                AND product_sku.meta_key = '_sku') T
--         WHERE value NOT LIKE '%s:%'
--           AND value NOT LIKE '%}%'
--           AND value != ''
--     ) T ORDER BY sku, sort_order;