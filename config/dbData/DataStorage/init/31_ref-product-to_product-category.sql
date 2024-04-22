-- create table
CREATE TABLE IF NOT EXISTS public.ref_product_to_product_category (
    id uuid NOT NULL,
    product_id uuid NOT NULL,
    product_category_id uuid NOT NULL,
    active           BOOLEAN      DEFAULT TRUE,
    sort_order          INTEGER      DEFAULT null,
    created_at TIMESTAMP   DEFAULT NOW(),
    created_by UUID        DEFAULT NULL REFERENCES public.user (id),
    updated_at TIMESTAMP   DEFAULT NOW(),
    updated_by UUID        DEFAULT NULL REFERENCES public.user (id),
    CONSTRAINT ref_product_to_product_category_pkey PRIMARY KEY (id)
);

-- index, constraint and ownership
ALTER TABLE public.ref_product_to_product_category OWNER TO postgres;
CREATE INDEX ref_product_to_product_category_id ON public.ref_product_to_product_category USING btree (id);
CREATE INDEX ref_product_to_product_category_product_id ON public.ref_product_to_product_category USING btree (product_id);
CREATE INDEX ref_product_to_product_category_product_category_id ON public.ref_product_to_product_category USING btree (product_category_id);
ALTER TABLE public.ref_product_to_product_category ADD CONSTRAINT ref_product_to_product_category_product_id_fkey FOREIGN KEY (product_id) REFERENCES public.product_info(id);
ALTER TABLE public.ref_product_to_product_category ADD CONSTRAINT ref_product_to_product_category_product_category_id_fkey FOREIGN KEY (product_category_id) REFERENCES public.product_category(id);
COMMENT ON TABLE public.ref_product_to_product_category IS 'Reference table for product and product category';

-- auto set sort_order column by product_category_id (Order of product in product category)
CREATE OR REPLACE FUNCTION public.ref_product_to_product_category_set_order()
    RETURNS trigger
    LANGUAGE 'plpgsql'
    COST 100
    VOLATILE NOT LEAKPROOF
AS $BODY$
BEGIN
    UPDATE public.ref_product_to_product_category
    SET sort_order = (SELECT COUNT(*) FROM public.ref_product_to_product_category WHERE product_category_id = NEW.product_category_id) + 1
    WHERE id = NEW.id;
    RETURN NEW;
END;
$BODY$;

CREATE TRIGGER ref_product_to_product_category_set_order
    BEFORE INSERT
    ON public.ref_product_to_product_category
    FOR EACH ROW
    EXECUTE PROCEDURE public.ref_product_to_product_category_set_order();

