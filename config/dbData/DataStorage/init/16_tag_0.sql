-- drop table if exists
DROP TABLE IF EXISTS public.tag;

-- create table
CREATE TABLE IF NOT EXISTS public.tag
(
    id            UUID UNIQUE  DEFAULT uuid_generate_v4(),
    product_id    UUID         DEFAULT NULL,
    tag_type_id   UUID         DEFAULT NULL,
    tag_select_id UUID         DEFAULT NULL,
    value         VARCHAR(255) DEFAULT NULL,
    active        BOOLEAN      DEFAULT TRUE,
    sort_order    INTEGER      DEFAULT NULL,
    created_at    TIMESTAMP    DEFAULT NOW() NOT NULL,
    created_by    UUID         DEFAULT NULL,
    updated_at    TIMESTAMP    DEFAULT NOW() NOT NULL,
    updated_by    UUID         DEFAULT NULL,

    CONSTRAINT tag_pkey PRIMARY KEY (id)
);

-- ownership and index
ALTER TABLE public.tag  OWNER TO postgres;
CREATE INDEX IF NOT EXISTS tag_id ON public.tag (id);
CREATE INDEX IF NOT EXISTS tag_product_id ON public.tag (product_id);
CREATE INDEX IF NOT EXISTS tag_tag_type_id ON public.tag (tag_type_id);
CREATE INDEX IF NOT EXISTS tag_tag_select_id ON public.tag (tag_select_id);
CREATE INDEX IF NOT EXISTS tag_sort_order ON public.tag (sort_order);
CREATE INDEX IF NOT EXISTS tag_updated_at ON public.tag (updated_at);

-- add comment to table
COMMENT ON TABLE public.tag IS 'Product Tag';
COMMENT ON COLUMN public.tag.id IS 'Primary Key';
COMMENT ON COLUMN public.tag.product_id IS 'Product ID, Reference to Product Info';
COMMENT ON COLUMN public.tag.tag_type_id IS 'Tag Type ID, Reference to Tag Type';
COMMENT ON COLUMN public.tag.tag_select_id IS 'Tag Select ID, Reference to Tag Select';
COMMENT ON COLUMN public.tag.value IS 'Tag Value';
COMMENT ON COLUMN public.tag.active IS 'Active';
COMMENT ON COLUMN public.tag.sort_order IS 'Sort Order';
COMMENT ON COLUMN public.tag.created_at IS 'Created At';
COMMENT ON COLUMN public.tag.created_by IS 'Created By';
COMMENT ON COLUMN public.tag.updated_at IS 'Updated At';
COMMENT ON COLUMN public.tag.updated_by IS 'Updated By';

-- auto update updated_at
CREATE OR REPLACE TRIGGER tag_updated_at
    BEFORE UPDATE
    ON public.tag
    FOR EACH ROW
EXECUTE FUNCTION update_update_at_column();

-- auto update sort_order by product_id, tag_type_id and tag_select_id
CREATE OR REPLACE FUNCTION update_tag_sort_order_for_tag()
    RETURNS TRIGGER
    LANGUAGE PLPGSQL
AS $$
BEGIN
    UPDATE public.tag
    SET sort_order = sort_order + 1
    WHERE product_id = NEW.product_id
      AND tag_type_id = NEW.tag_type_id
      AND tag_select_id = NEW.tag_select_id
      AND sort_order >= NEW.sort_order;

    RETURN NEW;
END;
$$;

CREATE OR REPLACE TRIGGER update_tag_sort_order_for_tag
    BEFORE INSERT
    ON public.tag
    FOR EACH ROW
EXECUTE FUNCTION update_tag_sort_order_for_tag();

-- SELECT
--     CONCAT('(select id from public.product_info where sku = "', T.product_sku, '")') as product_id,
--     CONCAT('(select id from public.tag_type where url = "', T.tag_type_url, '")') as tag_type_id,
--     CONCAT('(select id from public.tag_select where tag_type_id = (select id from public.tag_type where url = "', T.tag_type_url, '") and url = "', T.tag_select_url , '")') as tag_select_id
-- FROM (
--          SELECT
--              pm.meta_value as product_sku,
--              SUBSTRING(tt.taxonomy, 4) as tag_type_url,
--              t.slug as tag_select_url
--          FROM wp_posts p
--                   JOIN wp_postmeta pm ON p.ID = pm.post_id AND pm.meta_key = '_sku'
--                   JOIN wp_term_relationships tr ON p.ID = tr.object_id
--                   JOIN wp_term_taxonomy tt ON tr.term_taxonomy_id = tt.term_taxonomy_id AND tt.taxonomy LIKE 'pa_%'
--                   JOIN wp_terms t ON tt.term_id = t.term_id
--          WHERE p.post_type = 'product'
--
--          UNION ALL
--
--          SELECT
--              pm.meta_value AS product_sku,
--              'category' AS tag_type_url,
--              t.slug AS tag_select_url
--          FROM
--              wp_posts p
--                  JOIN
--              wp_postmeta pm ON p.ID = pm.post_id AND pm.meta_key = '_sku'
--                  JOIN
--              wp_term_relationships tr ON p.ID = tr.object_id
--                  JOIN
--              wp_term_taxonomy tt ON tr.term_taxonomy_id = tt.term_taxonomy_id AND tt.taxonomy = 'product_cat'
--                  JOIN
--              wp_terms t ON tt.term_id = t.term_id
--          WHERE
--              p.post_type = 'product'
--
--          UNION ALL
--
--          SELECT
--              pm.meta_value AS product_sku,
--              'shipping-class' AS tag_type_url,
--              t.slug AS tag_select_url
--          FROM
--              wp_posts p
--                  JOIN
--              wp_postmeta pm ON p.ID = pm.post_id AND pm.meta_key = '_sku'
--                  JOIN
--              wp_term_relationships tr ON p.ID = tr.object_id
--                  JOIN
--              wp_term_taxonomy tt ON tr.term_taxonomy_id = tt.term_taxonomy_id AND tt.taxonomy = 'product_shipping_class'
--                  JOIN
--              wp_terms t ON tt.term_id = t.term_id
--          WHERE
--              p.post_type = 'product') as T
--     ) AS T;
