
-- create uuid extension
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE OR REPLACE FUNCTION update_update_at_column()
    RETURNS TRIGGER AS $$
BEGIN
   NEW.updated_at = current_timestamp;
   RETURN NEW;
END;
$$ language 'plpgsql';

-- drop table
-- DROP TABLE IF EXISTS public.product;
-- DROP TABLE IF EXISTS public.product_status;
-- DROP TABLE IF EXISTS public.shipping_class;
-- DROP TABLE IF EXISTS public.category;
-- DROP TABLE IF EXISTS public.specification;
-- DROP TABLE IF EXISTS public.specification_value;
-- DROP TABLE IF EXISTS public.user;
-- DROP TABLE IF EXISTS public.folder;