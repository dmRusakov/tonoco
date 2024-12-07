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

-- test data insert
INSERT INTO public.text (id, language, source, sub_source, source_id, text, active, created_at, created_by, updated_at, updated_by) VALUES ('3184cf72-a92c-463e-9517-939712abfeac', 'en', 'shop', 'seo_title', '79997faf-fd52-4bc9-bda2-696b20d29973', 'Premium Range Hoods by Futuro Futuro | Modern Kitchen Ventilation', true, '2024-11-27 03:09:34.742775', '0e95efda-f9e2-4fac-8184-3ce2e8b7e0e1', '2024-11-27 03:12:15.630133', '0e95efda-f9e2-4fac-8184-3ce2e8b7e0e1');
INSERT INTO public.text (id, language, source, sub_source, source_id, text, active, created_at, created_by, updated_at, updated_by) VALUES ('03fe5a5f-d40c-4543-8e36-26bcf29cd563', 'en', 'shop', 'short_description', '79997faf-fd52-4bc9-bda2-696b20d29973', 'Upgrade your kitchen with Futuro Futuro’s stylish and efficient range hoods. Discover Italian craftsmanship, whisper-quiet operation, and exceptional performance, tailored to fit every kitchen style.', true, '2024-11-27 03:09:34.742775', '0e95efda-f9e2-4fac-8184-3ce2e8b7e0e1', '2024-11-27 03:12:15.630133', '0e95efda-f9e2-4fac-8184-3ce2e8b7e0e1');
INSERT INTO public.text (id, language, source, sub_source, source_id, text, active, created_at, created_by, updated_at, updated_by) VALUES ('ba30af8e-ad87-4f75-b81c-829396b739e3', 'en', 'shop', 'name', '79997faf-fd52-4bc9-bda2-696b20d29973', 'Discover the Perfect Range Hood for Your Kitchen.', true, '2024-11-27 03:09:34.742775', '0e95efda-f9e2-4fac-8184-3ce2e8b7e0e1', '2024-11-27 03:16:10.235540', '0e95efda-f9e2-4fac-8184-3ce2e8b7e0e1');
INSERT INTO public.text (id, language, source, sub_source, source_id, text, active, created_at, created_by, updated_at, updated_by) VALUES ('fe8d1d54-2286-4f9d-9c65-bbc5728d30fc', 'en', 'shop', 'description', '79997faf-fd52-4bc9-bda2-696b20d29973', '<section><h2>Transform Your Kitchen with Futuro Futuro Range Hoods</h2><p>Futuro Futuro combines the elegance of Italian design with advanced engineering to create range hoods that elevate your kitchen experience. Whether you’re an avid home chef or someone who enjoys occasional cooking, our range hoods offer powerful performance, exceptional durability, and timeless style. Built to ensure superior ventilation, they remove smoke, odors, and grease while enhancing the overall aesthetics of your space. With a commitment to quality and innovation, Futuro Futuro has become a trusted name in kitchens across the U.S.</p></section><section><h2>Stylish Designs That Suit Every Home</h2><p>Your kitchen is more than just a cooking area—it’s a central space for family, friends, and memories. Futuro Futuro range hoods are designed to complement all styles, from ultra-modern to classic traditional. Explore our wide selection of wall-mounted, island, and ceiling-mounted range hoods, each crafted from premium materials like polished stainless steel, elegant tempered glass, and bold architectural forms. Whether you prefer a sleek minimalist look or a dramatic statement piece, Futuro Futuro offers designs that blend seamlessly with your kitchen while standing out as a sophisticated focal point.</p><p>Beyond aesthetics, these range hoods are meticulously engineered to provide functionality. The intuitive controls, customizable settings, and modern finishes make them a joy to use. Our models are designed to fit kitchens of all sizes, ensuring you find the perfect match for your space and cooking needs.</p></section><section><h2>Advanced Performance for Cleaner, Fresher Air</h2><p>A great kitchen deserves great air quality. Futuro Futuro range hoods are equipped with powerful motors that provide efficient airflow to eliminate smoke, grease, and lingering odors. Our innovative filtration systems, including dishwasher-safe grease filters and activated carbon filters, ensure that even the busiest kitchens stay fresh and clean. Engineered for whisper-quiet operation, you’ll hardly notice the ventilation even at higher settings, allowing you to focus on cooking and enjoying your time in the kitchen.</p><p>Energy efficiency is another hallmark of Futuro Futuro range hoods. With low-energy LED lighting and optimized airflow systems, our products are designed to minimize energy consumption while maximizing performance. Whether you need heavy-duty ventilation for elaborate meals or a subtle air-clearing solution for light cooking, we have a model tailored to your needs.</p></section><section><h2>Durability That Stands the Test of Time</h2><p>Investing in a range hood means investing in your home’s future. Futuro Futuro uses premium-grade materials, including food-grade stainless steel and durable tempered glass, ensuring that your range hood remains beautiful and functional for years. Our hoods are resistant to corrosion, easy to clean, and require minimal maintenance. The combination of high-quality construction and attention to detail means your range hood will retain its elegance and efficiency, even in the most demanding conditions.</p><p>Our range hoods are rigorously tested to meet high performance and safety standards. From the durable motors to the carefully designed control panels, every component is crafted with precision. This makes Futuro Futuro range hoods not just a stylish addition to your kitchen but a long-term solution for clean air and improved cooking experiences.</p></section><section><h2>Italian Craftsmanship, Trusted in the U.S.</h2><p>Futuro Futuro’s range hoods are a testament to the art of Italian manufacturing. With a legacy of combining traditional craftsmanship and modern technology, we deliver products that are as functional as they are beautiful. Trusted by homeowners, chefs, and interior designers, our range hoods set the standard for kitchen ventilation systems across the U.S.</p><p>Our commitment to sustainability also means you can enjoy premium quality while being environmentally conscious. We prioritize eco-friendly materials, energy-efficient designs, and minimal environmental impact throughout the manufacturing process. This ensures you’re not just investing in your kitchen but also in the planet’s future.</p></section><section><p>Discover the difference a Futuro Futuro range hood can make in your kitchen. Browse our extensive collection today and experience the perfect balance of style, performance, and reliability.</p></section>', true, '2024-11-27 03:09:34.742775', '0e95efda-f9e2-4fac-8184-3ce2e8b7e0e1', '2024-11-27 03:17:06.110262', '0e95efda-f9e2-4fac-8184-3ce2e8b7e0e1');

-- get data
select * from public.text;