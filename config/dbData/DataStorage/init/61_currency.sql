-- create table
CREATE TABLE IF NOT EXISTS public.currency
(
    id          UUID UNIQUE DEFAULT uuid_generate_v4() NOT NULL,
    name        VARCHAR(255)                           NOT NULL,
    symbol      VARCHAR(10)                            NOT NULL,
    url         VARCHAR(255)                           NOT NULL,
    sort_order  INTEGER                                NOT NULL,
    active      BOOLEAN     DEFAULT TRUE               NOT NULL,

    created_at TIMESTAMP   DEFAULT NOW()              NOT NULL,
    created_by UUID        DEFAULT '0e95efda-f9e2-4fac-8184-3ce2e8b7e0e1' REFERENCES public.user (id),
    updated_at TIMESTAMP   DEFAULT NOW()              NOT NULL,
    updated_by UUID        DEFAULT '0e95efda-f9e2-4fac-8184-3ce2e8b7e0e1' REFERENCES public.user (id),

    CONSTRAINT currency_pkey PRIMARY KEY (id)
);

-- ownership and index
ALTER TABLE public.currency OWNER TO postgres;
CREATE INDEX currency_name ON public.currency (name);
CREATE INDEX currency_url ON public.currency (url);
CREATE INDEX currency_sort_order ON public.currency (sort_order);
CREATE INDEX currency_updated_at ON public.currency (updated_at);

-- add comments
COMMENT ON TABLE public.currency IS 'Currency table';
COMMENT ON COLUMN public.currency.id IS 'Unique identifier';
COMMENT ON COLUMN public.currency.name IS 'Currency name';
COMMENT ON COLUMN public.currency.symbol IS 'Currency symbol';
COMMENT ON COLUMN public.currency.url IS 'Currency URL';
COMMENT ON COLUMN public.currency.sort_order IS 'Sort order';
COMMENT ON COLUMN public.currency.active IS 'Active status';
COMMENT ON COLUMN public.currency.created_at IS 'Record created date';
COMMENT ON COLUMN public.currency.created_by IS 'Record created by';
COMMENT ON COLUMN public.currency.updated_at IS 'Record updated date';
COMMENT ON COLUMN public.currency.updated_by IS 'Record updated by';

-- auto set sort_order column
CREATE TRIGGER currency_order
    BEFORE INSERT
    ON public.currency
    FOR EACH ROW
EXECUTE FUNCTION set_order_column_universal();

-- auto update updated_at
CREATE TRIGGER currency_updated_at
    BEFORE UPDATE
    ON public.currency
    FOR EACH ROW
EXECUTE FUNCTION update_update_at_column();

-- default data
INSERT INTO public.currency (id, name, symbol, url)
VALUES ('c475a6f3-55ad-4641-8caa-a76bfae13fb0', 'USD', '$', 'USD'),
       ('c475a6f3-55ad-4641-8caa-a76bfae13fb1', 'Euro', '€', 'EUR');
