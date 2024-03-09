-- create table
CREATE TABLE IF NOT EXISTS public.product_status
(
    id         UUID UNIQUE DEFAULT uuid_generate_v4() NOT NULL,
    name       VARCHAR(255)                           NOT NULL,
    url        VARCHAR(255)                           NOT NULL,
    sort_order INTEGER                                NOT NULL,
    active     BOOLEAN     DEFAULT TRUE               NOT NULL,

    created_at TIMESTAMP   DEFAULT NOW()              NOT NULL,
    created_by UUID        DEFAULT '0e95efda-f9e2-4fac-8184-3ce2e8b7e0e1' REFERENCES public.user (id),
    updated_at TIMESTAMP   DEFAULT NOW()              NOT NULL,
    updated_by UUID        DEFAULT '0e95efda-f9e2-4fac-8184-3ce2e8b7e0e1' REFERENCES public.user (id),

    CONSTRAINT product_status_pkey PRIMARY KEY (id)
);

-- ownership, index and comment
ALTER TABLE public.product_status OWNER TO postgres;
CREATE INDEX product_status_id ON public.product_status (id);
CREATE INDEX product_status_url ON public.product_status (url);
CREATE INDEX product_status_sort_order ON public.product_status (sort_order);
CREATE INDEX product_status_updated_at ON public.product_status (updated_at);

-- add comment to table
COMMENT ON TABLE public.product_status IS 'Product Statuses';
COMMENT ON COLUMN public.product_status.id IS 'Unique identifier';
COMMENT ON COLUMN public.product_status.name IS 'Name of the status';
COMMENT ON COLUMN public.product_status.url IS 'Url of the status';
COMMENT ON COLUMN public.product_status.sort_order IS 'Sort order of the status';
COMMENT ON COLUMN public.product_status.active IS 'Status is active or not';

-- auto update updated_at
CREATE TRIGGER product_status_updated_at
    BEFORE UPDATE
    ON public.user
    FOR EACH ROW
EXECUTE FUNCTION update_update_at_column();

-- demo data
INSERT INTO public.product_status (id, name, url, sort_order)
VALUES ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11', 'Public', 'public', 0),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a12', 'Private', 'private', 1),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a13', 'Out of stock', 'out-of-stock', 2),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a14', 'Discontinued', 'discontinued', 3),
       ('a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a15', 'Archived', 'archived', 4);

-- get data
select * from public.product_status;