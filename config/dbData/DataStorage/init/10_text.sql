DROP TABLE IF EXISTS public.text CASCADE;

-- create table
CREATE TABLE IF NOT EXISTS public.text
(
    id         UUID UNIQUE  DEFAULT uuid_generate_v4() NOT NULL,
    language   VARCHAR(2)   DEFAULT 'en',
    source     VARCHAR(100) DEFAULT NULL,
    sub_source VARCHAR(100) DEFAULT NULL,
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
CREATE INDEX text_sub_source ON public.text (sub_source);
CREATE INDEX text_source_id ON public.text (source_id);
CREATE INDEX text_active ON public.text (active);
CREATE INDEX text_updated_at ON public.text (updated_at);
CREATE INDEX text_vector ON public.text USING gin (vector);

-- add comments
COMMENT ON TABLE public.text IS 'Text table';
COMMENT ON COLUMN public.text.id IS 'Unique identifier';
COMMENT ON COLUMN public.text.language IS 'Language of the text';
COMMENT ON COLUMN public.text.source IS 'Source of the text';
COMMENT ON COLUMN public.text.sub_source IS 'Sub source of the text';
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
select * from public.text;

-- test data insert
INSERT INTO public.text (id, language, source, sub_source, source_id, text, active, created_at, created_by, updated_at, updated_by) VALUES ('3184cf72-a92c-463e-9517-939712abfeac', 'en', 'shop', 'seo_title', '79997faf-fd52-4bc9-bda2-696b20d29973', 'Premium Range Hoods by Futuro Futuro | Modern Kitchen Ventilation', true, '2024-11-27 03:09:34.742775', '0e95efda-f9e2-4fac-8184-3ce2e8b7e0e1', '2024-11-27 03:12:15.630133', '0e95efda-f9e2-4fac-8184-3ce2e8b7e0e1');
INSERT INTO public.text (id, language, source, sub_source, source_id, text, active, created_at, created_by, updated_at, updated_by) VALUES ('03fe5a5f-d40c-4543-8e36-26bcf29cd563', 'en', 'shop', 'short_description', '79997faf-fd52-4bc9-bda2-696b20d29973', 'Upgrade your kitchen with Futuro Futuro’s stylish and efficient range hoods. Discover Italian craftsmanship, whisper-quiet operation, and exceptional performance, tailored to fit every kitchen style.', true, '2024-11-27 03:09:34.742775', '0e95efda-f9e2-4fac-8184-3ce2e8b7e0e1', '2024-11-27 03:12:15.630133', '0e95efda-f9e2-4fac-8184-3ce2e8b7e0e1');
INSERT INTO public.text (id, language, source, sub_source, source_id, text, active, created_at, created_by, updated_at, updated_by) VALUES ('ba30af8e-ad87-4f75-b81c-829396b739e3', 'en', 'shop', 'name', '79997faf-fd52-4bc9-bda2-696b20d29973', 'Discover the Perfect Range Hood for Your Kitchen – Futuro Futuro', true, '2024-11-27 03:09:34.742775', '0e95efda-f9e2-4fac-8184-3ce2e8b7e0e1', '2024-11-27 03:16:10.235540', '0e95efda-f9e2-4fac-8184-3ce2e8b7e0e1');
INSERT INTO public.text (id, language, source, sub_source, source_id, text, active, created_at, created_by, updated_at, updated_by) VALUES ('fe8d1d54-2286-4f9d-9c65-bbc5728d30fc', 'en', 'shop', 'description', '79997faf-fd52-4bc9-bda2-696b20d29973', '<p>Futuro Futuro brings together decades of Italian craftsmanship and state-of-the-art engineering to offer a range hood collection that redefines kitchen ventilation. Designed with both performance and aesthetics in mind, our range hoods combine powerful airflow systems with elegant designs to elevate any kitchen. Whether you’re cooking a quiet family dinner or hosting a lively gathering, Futuro Futuro range hoods ensure clean, fresh air and a stunning centerpiece for your space.</p>
<h2>Stylish Designs for Every Kitchen</h2>
<p>At Futuro Futuro, we understand that the kitchen is more than just a place to cook—it’s the heart of your home. That’s why our range hoods are available in a variety of styles, including sleek stainless steel, contemporary tempered glass, and minimalist designs. Choose from wall-mounted, island, or ceiling-mounted models to match your kitchen layout. Our range hoods not only complement modern and traditional kitchens but also create a focal point that showcases your style and attention to detail.</p>
<h2>Superior Performance and Quiet Operation</h2>
<p>Cooking should be a joyful experience, not a noisy chore. Futuro Futuro range hoods are engineered for whisper-quiet operation without compromising on power. Featuring advanced filtration systems, our hoods effectively capture grease, smoke, and odors, ensuring excellent air quality. With energy-efficient LED lighting and durable construction, these hoods are built to last while enhancing your cooking experience. Whether you need heavy-duty ventilation or a more subtle airflow solution, we have a range hood designed to meet your needs.</p>
<h2>Trusted Quality – Made in Italy</h2>
<p>Every Futuro Futuro range hood is designed and manufactured in Italy, reflecting a heritage of quality and innovation. Our products are trusted by homeowners, interior designers, and chefs across the U.S. We focus on combining beauty with functionality, ensuring each range hood delivers exceptional performance and seamlessly integrates into your kitchen. With a commitment to sustainability, we also prioritize eco-friendly materials and energy-efficient technologies in every design.</p>', true, '2024-11-27 03:09:34.742775', '0e95efda-f9e2-4fac-8184-3ce2e8b7e0e1', '2024-11-27 03:17:06.110262', '0e95efda-f9e2-4fac-8184-3ce2e8b7e0e1');
