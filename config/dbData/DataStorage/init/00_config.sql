
-- create uuid extension
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- auto-set updated at
CREATE OR REPLACE FUNCTION update_update_at_column()
    RETURNS TRIGGER AS $$
BEGIN
   NEW.updated_at = current_timestamp;
   RETURN NEW;
END;
$$ language 'plpgsql';


-- auto set sort_order column
CREATE OR REPLACE FUNCTION set_order_column_universal()
    RETURNS TRIGGER AS
$$
DECLARE
max_sort_order INTEGER;
BEGIN
    EXECUTE format('SELECT COALESCE(MAX(sort_order) + 1, 0) FROM %I', TG_TABLE_NAME) INTO max_sort_order;
    NEW.sort_order = max_sort_order;
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