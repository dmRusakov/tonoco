-- drop table if exists
DROP TABLE IF EXISTS public.error CASCADE;

-- create table
CREATE TABLE IF NOT EXISTS public.error
(
    id          UUID UNIQUE DEFAULT uuid_generate_v4() NOT NULL,
    type        varchar(11) NOT NULL,
    message     varchar(4000) NOT NULL,
    dev_message varchar(4000) NOT NULL,
    field       varchar(255) NOT NULL,
    code        varchar(11) NOT NULL,

    created_at TIMESTAMP   DEFAULT NOW()               NOT NULL,
    created_by UUID        DEFAULT '0e95efda-f9e2-4fac-8184-3ce2e8b7e0e1' NOT NULL,
    updated_at TIMESTAMP   DEFAULT NOW()               NOT NULL,
    updated_by UUID        DEFAULT '0e95efda-f9e2-4fac-8184-3ce2e8b7e0e1' NOT NULL,

    CONSTRAINT error_pkey PRIMARY KEY (id)
);

-- ownership and index
ALTER TABLE public.error OWNER TO postgres;
CREATE INDEX IF NOT EXISTS error_id ON public.error (id);
CREATE INDEX IF NOT EXISTS error_code ON public.error (code);

-- add comments
COMMENT ON TABLE public.error IS 'Error table';
COMMENT ON COLUMN public.error.id IS 'Unique identifier';
COMMENT ON COLUMN public.error.type IS 'Error type(error/validation)';
COMMENT ON COLUMN public.error.message IS 'Error message';
COMMENT ON COLUMN public.error.dev_message IS 'Developer error message';
COMMENT ON COLUMN public.error.field IS 'Field that caused the error';
COMMENT ON COLUMN public.error.code IS 'Error code';

CREATE OR REPLACE TRIGGER update_update_at_column
BEFORE UPDATE ON public.error
FOR EACH ROW
EXECUTE FUNCTION set_order_column_universal();
