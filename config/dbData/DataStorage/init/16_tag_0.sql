CREATE TABLE IF NOT EXISTS public.tag
(
    id            UUID UNIQUE  DEFAULT uuid_generate_v4(),
    product_id    UUID         NOT NULL REFERENCES public.product_info (id),
    tag_type_id   UUID         NOT NULL REFERENCES public.tag_type (id),
    tag_select_id UUID         DEFAULT NULL REFERENCES public.tag_select (id),
    value         VARCHAR(255) DEFAULT NULL,
    active        BOOLEAN      DEFAULT TRUE,
    sort_order    INTEGER      DEFAULT NULL,
    created_at    TIMESTAMP    DEFAULT NOW() NOT NULL,
    created_by    UUID         DEFAULT '0e95efda-f9e2-4fac-8184-3ce2e8b7e0e1' REFERENCES public.user (id),
    updated_at    TIMESTAMP    DEFAULT NOW() NOT NULL,
    updated_by    UUID         DEFAULT '0e95efda-f9e2-4fac-8184-3ce2e8b7e0e1' REFERENCES public.user (id),

    CONSTRAINT tag_pkey PRIMARY KEY (id)
);

-- ownership and index
ALTER TABLE public.tag  OWNER TO postgres;
CREATE INDEX tag_id ON public.tag (id);
CREATE INDEX tag_product_id ON public.tag (product_id);
CREATE INDEX tag_tag_type_id ON public.tag (tag_type_id);
CREATE INDEX tag_tag_select_id ON public.tag (tag_select_id);
CREATE INDEX tag_sort_order ON public.tag (sort_order);
CREATE INDEX tag_updated_at ON public.tag (updated_at);

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
CREATE TRIGGER tag_updated_at
    BEFORE UPDATE
    ON public.user
    FOR EACH ROW
EXECUTE FUNCTION update_update_at_column();

-- SELECT
--     CONCAT('(select id from public.product_info where sku = "', T.product_sku, '")') as product_id,
--     CONCAT('(select id from public.tag_type where url = "', T.tag_type_url, '")') as tag_type_id,
--     CONCAT('(select id from public.tag_select where tag_type_id = (select id from public.tag_type where url = "', T.tag_type_url, '") and url = "', T.tag_select_url , '")') as tag_select_id
-- FROM (
--     SELECT
--         pm.meta_value as product_sku,
--         SUBSTRING(tt.taxonomy, 4) as tag_type_url,
--         t.slug as tag_select_url
--     FROM wp_posts p
--     JOIN wp_postmeta pm ON p.ID = pm.post_id AND pm.meta_key = '_sku'
--     JOIN wp_term_relationships tr ON p.ID = tr.object_id
--     JOIN wp_term_taxonomy tt ON tr.term_taxonomy_id = tt.term_taxonomy_id AND tt.taxonomy LIKE 'pa_%'
--     JOIN wp_terms t ON tt.term_id = t.term_id
--     WHERE p.post_type = 'product'
-- ) AS T;