-- drop table if exists
DROP TABLE IF EXISTS public.shop_tag_type CASCADE;

-- create table
CREATE TABLE IF NOT EXISTS public.shop_tag_type
(
    id          UUID UNIQUE DEFAULT uuid_generate_v4() NOT NULL,
    shop_id     UUID        DEFAULT NULL,
    tag_type_id UUID        DEFAULT NULL,
    source        VARCHAR(50) DEFAULT NULL,
    sort_order  INTEGER     DEFAULT NULL,
    active      BOOLEAN     DEFAULT TRUE,

    created_at  TIMESTAMP   DEFAULT NOW()              NOT NULL,
    created_by  UUID        DEFAULT NULL,
    updated_at  TIMESTAMP   DEFAULT NOW()              NOT NULL,
    updated_by  UUID        DEFAULT NULL,

    CONSTRAINT shop_tag_type_pkey PRIMARY KEY (id)
);

-- ownership and index
ALTER TABLE public.shop_tag_type OWNER TO postgres;
CREATE INDEX shop_tag_type_id ON public.shop_tag_type USING btree (id);
CREATE INDEX shop_tag_type_shop_id ON public.shop_tag_type USING btree (shop_id);
CREATE INDEX shop_tag_type_tag_type_id ON public.shop_tag_type USING btree (tag_type_id);
CREATE INDEX shop_tag_type_source ON public.shop_tag_type USING btree (source);
CREATE INDEX shop_tag_type_sort_order ON public.shop_tag_type USING btree (sort_order);
CREATE INDEX shop_tag_type_updated_at ON public.shop_tag_type USING btree (updated_at);

-- comment on table
COMMENT ON TABLE public.shop_tag_type IS 'Reference table for shop and tag type';
COMMENT ON COLUMN public.shop_tag_type.id IS 'Unique identifier for shop tag type';
COMMENT ON COLUMN public.shop_tag_type.shop_id IS 'Shop id';
COMMENT ON COLUMN public.shop_tag_type.tag_type_id IS 'Tag type id';
COMMENT ON COLUMN public.shop_tag_type.source IS 'Source of shop tag type';
COMMENT ON COLUMN public.shop_tag_type.sort_order IS 'Sort order of shop tag type';
COMMENT ON COLUMN public.shop_tag_type.active IS 'Active status of shop tag type';
COMMENT ON COLUMN public.shop_tag_type.created_at IS 'Creation time of shop tag type';
COMMENT ON COLUMN public.shop_tag_type.created_by IS 'Creator of shop tag type';
COMMENT ON COLUMN public.shop_tag_type.updated_at IS 'Update time of shop tag type';
COMMENT ON COLUMN public.shop_tag_type.updated_by IS 'Updater of shop tag type';

-- auto update updated_at
CREATE OR REPLACE TRIGGER shop_tag_type_updated_at
    BEFORE UPDATE OR INSERT
    ON public.shop_tag_type
    FOR EACH ROW
EXECUTE FUNCTION update_update_at_column();

-- auto set created_by
CREATE OR REPLACE TRIGGER shop_tag_type_created_by
    BEFORE INSERT
    ON public.shop_tag_type
    FOR EACH ROW
EXECUTE FUNCTION set_created_by_if_null();

-- auto set updated_by
CREATE OR REPLACE TRIGGER shop_tag_type_updated_by
    BEFORE INSERT OR UPDATE
    ON public.shop_tag_type
    FOR EACH ROW
EXECUTE FUNCTION set_updated_by_if_null();

-- data insert
INSERT INTO public.shop_tag_type (id, shop_id, tag_type_id, source, sort_order, active, created_at, created_by, updated_at, updated_by) VALUES ('cfa5c8d2-e48d-4b37-ad31-fe5d296e7004', '79997faf-fd52-4bc9-bda2-696b20d29973', 'a0eebc99-9c0b-4ef8-bb6d-6bb9bd381558', 'grid-tag', 1, true, '2024-12-08 23:45:58.168794', '0e95efda-f9e2-4fac-8184-3ce2e8b7e0e1', '2024-12-08 23:53:07.071560', '0e95efda-f9e2-4fac-8184-3ce2e8b7e0e1');
INSERT INTO public.shop_tag_type (id, shop_id, tag_type_id, source, sort_order, active, created_at, created_by, updated_at, updated_by) VALUES ('c2cd0223-3e84-4e50-b71e-bf7d88a3dc3e', '79997faf-fd52-4bc9-bda2-696b20d29973', 'a0eebc99-9c0b-4ef8-bb6d-6bb9bd382157', 'grid-tag', 2, true, '2024-12-08 23:51:46.069246', '0e95efda-f9e2-4fac-8184-3ce2e8b7e0e1', '2024-12-08 23:53:07.071560', '0e95efda-f9e2-4fac-8184-3ce2e8b7e0e1');
INSERT INTO public.shop_tag_type (id, shop_id, tag_type_id, source, sort_order, active, created_at, created_by, updated_at, updated_by) VALUES ('2c5ac012-359f-4cb9-af14-d54bbe87877d', '79997faf-fd52-4bc9-bda2-696b20d29973', 'a0eebc99-9c0b-4ef8-bb6d-6bb9bd381001', 'grid-tag', 3, true, '2024-12-08 23:50:46.824618', '0e95efda-f9e2-4fac-8184-3ce2e8b7e0e1', '2024-12-08 23:53:07.071560', '0e95efda-f9e2-4fac-8184-3ce2e8b7e0e1');

-- get data
select * from public.shop_tag_type;
