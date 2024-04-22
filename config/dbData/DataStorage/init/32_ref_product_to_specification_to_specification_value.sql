-- create table
CREATE TABLE IF NOT EXISTS public.ref_product_to_specification_to_specification_value (
    id uuid NOT NULL,
    product_id uuid NOT NULL,
    specification_id uuid NOT NULL,
    specification_value_id uuid NOT NULL,
    CONSTRAINT ref_product_to_specification_to_specification_value_pkey PRIMARY KEY (id)
);

ALTER TABLE public.ref_product_to_specification_to_specification_value OWNER TO postgres;
CREATE INDEX ref_pssv_id ON public.ref_product_to_specification_to_specification_value USING btree (id);
CREATE INDEX ref_pssv_product_id ON public.ref_product_to_specification_to_specification_value USING btree (product_id);
CREATE INDEX ref_pssv_specification_id ON public.ref_product_to_specification_to_specification_value USING btree (specification_id);
ALTER TABLE public.ref_product_to_specification_to_specification_value ADD CONSTRAINT ref_pssv_product_id_fkey FOREIGN KEY (product_id) REFERENCES public.product_info(id);
ALTER TABLE public.ref_product_to_specification_to_specification_value ADD CONSTRAINT ref_pssv_specification_id_fkey FOREIGN KEY (specification_id) REFERENCES public.specification(id);
ALTER TABLE public.ref_product_to_specification_to_specification_value ADD CONSTRAINT ref_pssv_specification_value_id_fkey FOREIGN KEY (specification_value_id) REFERENCES public.specification_value(id);
COMMENT ON TABLE public.ref_product_to_specification_to_specification_value IS 'Reference table for product to specification to specification value';




