DROP TABLE IF EXISTS public.text CASCADE;

-- create table
CREATE TABLE IF NOT EXISTS public.text
(
    id         UUID UNIQUE  DEFAULT uuid_generate_v4() NOT NULL,
    language   VARCHAR(2)   DEFAULT 'en',
    source     VARCHAR(100) DEFAULT NULL,
    source_id  UUID         DEFAULT NULL,

    text       TEXT         DEFAULT NULL,
    vector     TSVECTOR GENERATED ALWAYS AS (to_tsvector('english', text)) STORED,

    active     BOOLEAN      DEFAULT TRUE,

    created_at TIMESTAMP    DEFAULT NOW(),
    created_by UUID         DEFAULT NULL,
    updated_at TIMESTAMP    DEFAULT NOW(),
    updated_by UUID         DEFAULT NULL,

    CONSTRAINT texts_pkey PRIMARY KEY (id)
);

-- ownership and index
ALTER TABLE public.text OWNER TO postgres;
CREATE INDEX text_id ON public.text (id);
CREATE INDEX text_language ON public.text (language);
CREATE INDEX text_source ON public.text (source);
CREATE INDEX text_source_id ON public.text (source_id);
CREATE INDEX text_active ON public.text (active);
CREATE INDEX text_updated_at ON public.text (updated_at);
CREATE INDEX text_vector ON public.text USING gin (vector);

-- add comments
COMMENT ON TABLE public.text IS 'Text table';
COMMENT ON COLUMN public.text.id IS 'Unique identifier';
COMMENT ON COLUMN public.text.language IS 'Language of the text';
COMMENT ON COLUMN public.text.source IS 'Source of the text';
COMMENT ON COLUMN public.text.source_id IS 'Source identifier';
COMMENT ON COLUMN public.text.text IS 'Text';
COMMENT ON COLUMN public.text.vector IS 'Text vector';
COMMENT ON COLUMN public.text.active IS 'Active status';
COMMENT ON COLUMN public.text.created_at IS 'Creation date';
COMMENT ON COLUMN public.text.created_by IS 'Creator';
COMMENT ON COLUMN public.text.updated_at IS 'Last update date';
COMMENT ON COLUMN public.text.updated_by IS 'Last updater';

-- auto update updated_at
CREATE OR REPLACE TRIGGER text_updated_at
    BEFORE UPDATE
    ON public.text
    FOR EACH ROW
EXECUTE FUNCTION update_update_at_column();

-- auto set created_by
CREATE OR REPLACE TRIGGER text_created_by
    BEFORE INSERT
    ON public.text
    FOR EACH ROW
EXECUTE FUNCTION set_created_by_if_null();

-- auto set updated_by
CREATE OR REPLACE TRIGGER text_updated_by
    BEFORE INSERT
    ON public.text
    FOR EACH ROW
EXECUTE FUNCTION set_updated_by_if_null();

-- get data
select *
from public.text;
