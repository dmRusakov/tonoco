
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
    max_sort_order BIGINT;
BEGIN
    EXECUTE format('SELECT COALESCE(MAX(sort_order), 0) + 1 FROM %I', TG_TABLE_NAME) INTO max_sort_order;
    IF NEW.sort_order IS NULL OR NEW.sort_order = 0 THEN
        NEW.sort_order = max_sort_order;
    END IF;
    RETURN NEW;
END;
$$ language 'plpgsql';

CREATE OR REPLACE FUNCTION set_created_by_if_null()
    RETURNS TRIGGER AS $$
BEGIN
    IF NEW.created_by IS NULL THEN
        SELECT id INTO NEW.created_by FROM public.user LIMIT 1;
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION set_updated_by_if_null()
    RETURNS TRIGGER AS $$
BEGIN
    IF NEW.updated_by IS NULL THEN
        SELECT id INTO NEW.updated_by FROM public.user LIMIT 1;
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;



-- drop table
-- DROP TABLE IF EXISTS public.product_info;
-- DROP TABLE IF EXISTS public.product_status;
-- DROP TABLE IF EXISTS public.shipping_class;
-- DROP TABLE IF EXISTS public.category;
-- DROP TABLE IF EXISTS public.tag;
-- DROP TABLE IF EXISTS public.tag_select;
-- DROP TABLE IF EXISTS public.user;
-- DROP TABLE IF EXISTS public.folder;