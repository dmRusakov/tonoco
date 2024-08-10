-- drop
DROP TABLE IF EXISTS public.price;

-- create table
CREATE TABLE IF NOT EXISTS public.price
(
    id            UUID UNIQUE    DEFAULT uuid_generate_v4() NOT NULL,

    product_id    UUID           DEFAULT NULL,
    price_type_id UUID           DEFAULT NULL,
    currency_id   UUID           DEFAULT NULL,
    warehouse_id  UUID           DEFAULT NULL,
    store_id      UUID           DEFAULT NULL,

    price         DECIMAL(10, 2) DEFAULT NULL,

    sort_order    INTEGER,
    active        BOOLEAN        DEFAULT TRUE,

    start_date    TIMESTAMP      DEFAULT NULL,
    end_date      TIMESTAMP      DEFAULT NULL,

    created_at    TIMESTAMP      DEFAULT NOW()              NOT NULL,
    created_by    UUID           DEFAULT NULL,
    updated_at    TIMESTAMP      DEFAULT NOW()              NOT NULL,
    updated_by    UUID           DEFAULT NULL,

    CONSTRAINT price_pkey PRIMARY KEY (id)
);

-- ownership and index
ALTER TABLE public.price OWNER TO postgres;
CREATE INDEX price_id ON public.price USING btree (id);
CREATE INDEX price_product_id ON public.price USING btree (product_id);
CREATE INDEX price_currency_id ON public.price USING btree (currency_id);
CREATE INDEX price_updated_at ON public.price USING btree (updated_at);

-- comment on table
COMMENT ON TABLE public.price IS 'Price of product';
COMMENT ON COLUMN public.price.id IS 'Unique identifier for price';
COMMENT ON COLUMN public.price.product_id IS 'Product identifier';
COMMENT ON COLUMN public.price.price_type_id IS 'Price type identifier';
COMMENT ON COLUMN public.price.currency_id IS 'Currency identifier';
COMMENT ON COLUMN public.price.price IS 'Price of product';
COMMENT ON COLUMN public.price.created_at IS 'Creation time of price';
COMMENT ON COLUMN public.price.created_by IS 'Creator of price';
COMMENT ON COLUMN public.price.updated_at IS 'Update time of price';
COMMENT ON COLUMN public.price.updated_by IS 'Updater of price';

-- auto update updated_at
CREATE OR REPLACE TRIGGER price_updated_at
    BEFORE UPDATE
    ON public.price
    FOR EACH ROW
    EXECUTE FUNCTION update_update_at_column();

-- auto set sort_order column by product_id
CREATE OR REPLACE FUNCTION set_sort_order()
    RETURNS TRIGGER AS
$BODY$
DECLARE
    max_sort_order BIGINT ;
BEGIN
    EXECUTE format('SELECT COALESCE(MAX(sort_order), 0) + 1 FROM %I WHERE product_id = $1', TG_TABLE_NAME)
        USING NEW.product_id INTO max_sort_order;
    IF NEW.sort_order IS NULL or NEW.sort_order = 0 THEN
        NEW.sort_order = max_sort_order;
    ELSE
        NEW.sort_order = NEW.sort_order + max_sort_order;
    END IF;
    RETURN NEW;
END;
$BODY$
    LANGUAGE plpgsql;

CREATE OR REPLACE TRIGGER set_sort_order
    BEFORE INSERT
    ON public.price
    FOR EACH ROW
    EXECUTE FUNCTION set_sort_order();

-- demo data
INSERT INTO public.price (product_id, price_type_id, currency_id, price) VALUES
    ((select id from public.product_info where sku = 'IS48LUXOR'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '2695'),
    ((select id from public.product_info where sku = 'WL36MELARA'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '1395'),
    ((select id from public.product_info where sku = 'IS48POSITANO'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '2295'),
    ((select id from public.product_info where sku = 'IS48STREAMLINE-WHT'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '2495'),
    ((select id from public.product_info where sku = 'IS42ACQUA-GLS'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '2295'),
    ((select id from public.product_info where sku = 'IS36LUXOR'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '2595'),
    ((select id from public.product_info where sku = 'IS36MOONCRYS'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '1895'),
    ((select id from public.product_info where sku = 'IS36MOONINOX'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '1795'),
    ((select id from public.product_info where sku = 'IS36VENICE'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '1795'),
    ((select id from public.product_info where sku = 'IS36LOMBARDY-WHT'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '2995'),
    ((select id from public.product_info where sku = 'IS36LOMBARDY-BLK'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '2995'),
    ((select id from public.product_info where sku = 'IS36JUPITER-GLS'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '2295'),
    ((select id from public.product_info where sku = 'IS36STREAMLINE-WHT'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '2195'),
    ((select id from public.product_info where sku = 'IS36POSITANO'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '1995'),
    ((select id from public.product_info where sku = 'IS36POSITANO-FS'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '1995'),
    ((select id from public.product_info where sku = 'IS48EUROPE'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '2995'),
    ((select id from public.product_info where sku = 'IS36LOFT'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '2095'),
    ((select id from public.product_info where sku = 'IS34MUR-ALFA'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '2795'),
    ((select id from public.product_info where sku = 'IS34MUR-FORTUNA'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '2795'),
    ((select id from public.product_info where sku = 'IS34MUR-SNOW'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '2795'),
    ((select id from public.product_info where sku = 'IS36CONNECTICUT'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '1795'),
    ((select id from public.product_info where sku = 'IS28ELLIPSO'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '3195'),
    ((select id from public.product_info where sku = 'IS28MODERNO'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '2395'),
    ((select id from public.product_info where sku = 'IS27MUR-AUTUMN'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '2295'),
    ((select id from public.product_info where sku = 'IS24LOMBARDY-BLK'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '2495'),
    ((select id from public.product_info where sku = 'IS24LOMBARDY-WHT'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '2495'),
    ((select id from public.product_info where sku = 'IS24SPIRIT'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '1895'),
    ((select id from public.product_info where sku = 'IS24MOONINOX'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '1695'),
    ((select id from public.product_info where sku = 'IS16LOFTCTM'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '1895'),
    ((select id from public.product_info where sku = 'IS16LOFT'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '1695'),
    ((select id from public.product_info where sku = 'IS27MUR-ZEBRA'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '2295'),
    ((select id from public.product_info where sku = 'IS27MUR-ORION'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '2295'),
    ((select id from public.product_info where sku = 'IS27MUR-NEWYORK'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '2295'),
    ((select id from public.product_info where sku = 'IS27MUR-MOTION'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '2295'),
    ((select id from public.product_info where sku = 'IS27MUR-FROST'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '2295'),
    ((select id from public.product_info where sku = 'IS27MUR-SNOW'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '2295'),
    ((select id from public.product_info where sku = 'IS54SKYLIGHT'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '3595'),
    ((select id from public.product_info where sku = 'IS84EUROPE-STN'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '3995'),
    ((select id from public.product_info where sku = 'WL48STREAMLINE-WHT'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '2295'),
    ((select id from public.product_info where sku = 'WL48POSITANO'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '2095'),
    ((select id from public.product_info where sku = 'WL48MAGNUS'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '1795'),
    ((select id from public.product_info where sku = 'WL48LUXOR'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '2495'),
    ((select id from public.product_info where sku = 'WL48QUEST-BLK'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '1995'),
    ((select id from public.product_info where sku = 'WL48QUEST-WHT'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '1995'),
    ((select id from public.product_info where sku = 'WL48CAPRI'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '1695'),
    ((select id from public.product_info where sku = 'WL36SOLARIS'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '2395'),
    ((select id from public.product_info where sku = 'WL36ACQUA-GLS'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '1695'),
    ((select id from public.product_info where sku = 'WL36EDGE-WHT'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '1695'),
    ((select id from public.product_info where sku = 'WL36LOMBARDY-BLK'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '2495'),
    ((select id from public.product_info where sku = 'WL36LOMBARDY-WHT'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '2495'),
    ((select id from public.product_info where sku = 'WL36ACQUA-INOX'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '1695'),
    ((select id from public.product_info where sku = 'WL36INTEGRA-WHT'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '1795'),
    ((select id from public.product_info where sku = 'WL36MOONCRYS'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '1495'),
    ((select id from public.product_info where sku = 'WL36MOONINOX'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '1395'),
    ((select id from public.product_info where sku = 'WL36VENICE'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '1395'),
    ((select id from public.product_info where sku = 'WL36MYSTIC-GLS'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '1595'),
    ((select id from public.product_info where sku = 'WL36MYSTIC-INOX'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '1595'),
    ((select id from public.product_info where sku = 'WL36STREAMLINE-WHT'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '1895'),
    ((select id from public.product_info where sku = 'WL36JUPITER-GLS'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '1995'),
    ((select id from public.product_info where sku = 'WL36LINEARE'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '1495'),
    ((select id from public.product_info where sku = 'WL36RAZOR'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '1695'),
    ((select id from public.product_info where sku = 'WL36WAVEBLK'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '1795'),
    ((select id from public.product_info where sku = 'WL36DIAMOND'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '1695'),
    ((select id from public.product_info where sku = 'WL36BLKDIAM'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '1295'),
    ((select id from public.product_info where sku = 'WL36QUEST-BLK'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '1795'),
    ((select id from public.product_info where sku = 'WL36RAINBOW'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '1495'),
    ((select id from public.product_info where sku = 'WL36CAPRI'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '1495'),
    ((select id from public.product_info where sku = 'WL36IDEA'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '1595'),
    ((select id from public.product_info where sku = 'WL36MAGNUS'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '1495'),
    ((select id from public.product_info where sku = 'WL36PORTLAND'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '1395'),
    ((select id from public.product_info where sku = 'WL36PORTLAND-PLUS'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '1395'),
    ((select id from public.product_info where sku = 'WL36CONNECTICUT'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '1395'),
    ((select id from public.product_info where sku = 'WL36BRIDGEPORT'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '1395'),
    ((select id from public.product_info where sku = 'WL36ORCHID'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '1395'),
    ((select id from public.product_info where sku = 'WL24MODERNO'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '1895'),
    ((select id from public.product_info where sku = 'WL24AMALFI'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '1595'),
    ((select id from public.product_info where sku = 'WL24LOMBARDY-WHT'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '1995'),
    ((select id from public.product_info where sku = 'WL24LOMBARDY-BLK'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '1995'),
    ((select id from public.product_info where sku = 'WL24LINEARE'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '1395'),
    ((select id from public.product_info where sku = 'WL14JUPITER-LT'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '2195'),
    ((select id from public.product_info where sku = 'WL14JUPITER'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '1795'),
    ((select id from public.product_info where sku = 'WL24MOONINOX'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '1395'),
    ((select id from public.product_info where sku = 'WL24VENICE'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '1395'),
    ((select id from public.product_info where sku = 'WL24STREAMLINE-WHT'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '1395'),
    ((select id from public.product_info where sku = 'WL24CAPRI'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '1395'),
    ((select id from public.product_info where sku = 'WL24SPIRIT'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '1395'),
    ((select id from public.product_info where sku = 'WL22RAVENNA'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '2595'),
    ((select id from public.product_info where sku = 'WL30ACQUA-GLS'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '1595'),
    ((select id from public.product_info where sku = 'WL30MOONCRYS'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '1495'),
    ((select id from public.product_info where sku = 'WL22INSERT-BAF'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '1295'),
    ((select id from public.product_info where sku = 'WL32INSERT'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '1295'),
    ((select id from public.product_info where sku = 'WL27MUR-FORTUNALED'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '2095'),
    ((select id from public.product_info where sku = 'WL27MUR-ALFALED'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '2095'),
    ((select id from public.product_info where sku = 'WL27MUR-FROSTLED'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '2095'),
    ((select id from public.product_info where sku = 'WL27MUR-SNOWLED'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '2095'),
    ((select id from public.product_info where sku = 'WL27MUR-EMPIRELED'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '2095'),
    ((select id from public.product_info where sku = 'WL27MUR-MOTIONLED'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '2095'),
    ((select id from public.product_info where sku = 'WL27MUR-ORIONLED'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '2095'),
    ((select id from public.product_info where sku = 'WL39MUR-SNOWLED'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '2395'),
    ((select id from public.product_info where sku = 'AC-RBK'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '345'),
    ((select id from public.product_info where sku = 'AC-WIRELESS'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '325'),
    ((select id from public.product_info where sku = 'AC-CARBR6'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '50'),
    ((select id from public.product_info where sku = 'AC-CARBR8'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '50'),
    ((select id from public.product_info where sku = 'AC-METALFILTER'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '85'),
    ((select id from public.product_info where sku = 'AC-EXTIS'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '565'),
    ((select id from public.product_info where sku = 'AC-EXT-MURPEARL'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '565'),
    ((select id from public.product_info where sku = 'AC-EXTWS'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '245'),
    ((select id from public.product_info where sku = 'AC-EXTWR'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '245'),
    ((select id from public.product_info where sku = 'AC-HALOBULB'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '25'),
    ((select id from public.product_info where sku = 'AC-HALOFXRND'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '75'),
    ((select id from public.product_info where sku = 'AC-HALOFXSQ'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '75'),
    ((select id from public.product_info where sku = 'AC-MINIHALO'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '45'),
    ((select id from public.product_info where sku = 'AC-FLRBULB'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '35'),
    ((select id from public.product_info where sku = 'AC-BACKSPL'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '265'),
    ((select id from public.product_info where sku = 'AC-GLSJUPITERIS'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '400'),
    ((select id from public.product_info where sku = 'AC-GLSJUPITERWL'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '400'),
    ((select id from public.product_info where sku = 'AC-36PANELLOFT'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '350'),
    ((select id from public.product_info where sku = 'AC-EXT-CON850'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '245'),
    ((select id from public.product_info where sku = 'AC-EXT-FAB960'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '245'),
    ((select id from public.product_info where sku = 'AC-EXT-MIM850'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '245'),
    ((select id from public.product_info where sku = 'AC-EXT-STP960'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '245'),
    ((select id from public.product_info where sku = 'WL36AMALFI'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '1995'),
    ((select id from public.product_info where sku = 'IS36GULLWING-WHT'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '3395'),
    ((select id from public.product_info where sku = 'IS36GULLWING-BLK'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '3395'),
    ((select id from public.product_info where sku = 'WL36WAVEWHT'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '1795'),
    ((select id from public.product_info where sku = 'WL36SHADE'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '1295'),
    ((select id from public.product_info where sku = 'WL26PEARL-BLK'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '3295'),
    ((select id from public.product_info where sku = 'WL26PEARL-WHT'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '3295'),
    ((select id from public.product_info where sku = 'IS24AMALFI'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '1895'),
    ((select id from public.product_info where sku = 'AC-CARBSQ'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '50'),
    ((select id from public.product_info where sku = 'WL27MUR-AUTUMNLED'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '2095'),
    ((select id from public.product_info where sku = 'IS30PEARL-WHT'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '3595'),
    ((select id from public.product_info where sku = 'IS30PEARL-BLK'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '3595'),
    ((select id from public.product_info where sku = 'WL36POSITANO'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '1795'),
    ((select id from public.product_info where sku = 'WL48SHADE'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '1395'),
    ((select id from public.product_info where sku = 'IS42MOONCRYS'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '1995'),
    ((select id from public.product_info where sku = 'IS36ACQUA-GLS'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '2195'),
    ((select id from public.product_info where sku = 'WL27MUR-SPACE'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '1895'),
    ((select id from public.product_info where sku = 'WL22INSERT'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '1195'),
    ((select id from public.product_info where sku = 'IS36EUROPE'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '2395'),
    ((select id from public.product_info where sku = 'IS72EUROPE-STN'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '3695'),
    ((select id from public.product_info where sku = 'WL27MUR-GLOWLED'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '2095'),
    ((select id from public.product_info where sku = 'WL27MUR-MOONLIGHTLED'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '2095'),
    ((select id from public.product_info where sku = 'WL27MUR-NEWYORKLED'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '2095'),
    ((select id from public.product_info where sku = 'WL27MUR-ZEBRALED'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '2095'),
    ((select id from public.product_info where sku = 'WL39MUR-ALFALED'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '2395'),
    ((select id from public.product_info where sku = 'WL39MUR-FORTUNALED'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '2395'),
    ((select id from public.product_info where sku = 'AC-EXT-LUXWALL'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '345'),
    ((select id from public.product_info where sku = 'AC-EXTIR'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '565'),
    ((select id from public.product_info where sku = 'IS34MUR-EMPIRELED'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '2995'),
    ((select id from public.product_info where sku = 'IS34MUR-GLOWLED'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '2995'),
    ((select id from public.product_info where sku = 'IS34MUR-MOONLIGHTLED'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '2995'),
    ((select id from public.product_info where sku = 'IS34MUR-MOTIONLED'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '2995'),
    ((select id from public.product_info where sku = 'IS34MUR-NEWYORKLED'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '2995'),
    ((select id from public.product_info where sku = 'IS34MUR-ORIONLED'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '2995'),
    ((select id from public.product_info where sku = 'IS34MUR-ZEBRALED'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '2995'),
    ((select id from public.product_info where sku = 'IS27MUR-AUTUMNLED'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '2495'),
    ((select id from public.product_info where sku = 'IS27MUR-GLOWLED'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '2495'),
    ((select id from public.product_info where sku = 'IS27MUR-MOTIONLED'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '2495'),
    ((select id from public.product_info where sku = 'IS27MUR-NEWYORKLED'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '2495'),
    ((select id from public.product_info where sku = 'IS27MUR-ORIONLED'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '2495'),
    ((select id from public.product_info where sku = 'IS27MUR-ZEBRALED'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '2495'),
    ((select id from public.product_info where sku = 'IS34MUR-ALFALED'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '2995'),
    ((select id from public.product_info where sku = 'IS34MUR-FORTUNALED'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '2995'),
    ((select id from public.product_info where sku = 'IS34MUR-SNOWLED'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '2995'),
    ((select id from public.product_info where sku = 'IS38SKYLIGHT'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '2895'),
    ((select id from public.product_info where sku = 'WL42INSERT'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '1595'),
    ((select id from public.product_info where sku = 'AC-BLOWER4'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '330'),
    ((select id from public.product_info where sku = 'AC-EXT-LUXISLAND'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '595'),
    ((select id from public.product_info where sku = 'WL36MASSACHUSETTS'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '1395'),
    ((select id from public.product_info where sku = 'WL36VERMONT'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '1395'),
    ((select id from public.product_info where sku = 'IS27MUR-MOONLIGHT'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '2295'),
    ((select id from public.product_info where sku = 'IS27MUR-MOONLIGHTLED'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '2495'),
    ((select id from public.product_info where sku = 'WL36LUXOREQUO'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '2395'),
    ((select id from public.product_info where sku = 'WL27MUR-ECHOLED'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '2095'),
    ((select id from public.product_info where sku = 'IS27MUR-ECHOLED'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '2495'),
    ((select id from public.product_info where sku = 'IS34MUR-ECHOLED'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '2995'),
    ((select id from public.product_info where sku = 'WL36FLORARED'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '1695'),
    ((select id from public.product_info where sku = 'WL36FLORAGREEN'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '1695'),
    ((select id from public.product_info where sku = 'AC-GLEURO48'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '350'),
    ((select id from public.product_info where sku = 'AC-GLEURO36'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '350'),
    ((select id from public.product_info where sku = 'AC-GLSEUROSHELF'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '350'),
    ((select id from public.product_info where sku = 'WL16LOFT-CTM'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '1795'),
    ((select id from public.product_info where sku = 'WL36BOSTON'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '1395'),
    ((select id from public.product_info where sku = 'WL36CAMBRIDGE'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '1395'),
    ((select id from public.product_info where sku = 'WL27MUR-METROLED'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '2095'),
    ((select id from public.product_info where sku = 'IS34MUR-METROLED'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '2995'),
    ((select id from public.product_info where sku = 'IS27MUR-METROLED'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '2495'),
    ((select id from public.product_info where sku = 'IS36LIVORNO'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '2395'),
    ((select id from public.product_info where sku = 'WL36LIVORNO'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '1995'),
    ((select id from public.product_info where sku = 'IS36MINIMALWHT'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '1795'),
    ((select id from public.product_info where sku = 'WL36MINIMAL-WHT'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '1495'),
    ((select id from public.product_info where sku = 'IS48LINEARE'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '2495'),
    ((select id from public.product_info where sku = 'IS27MUR-EMPIRELED'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '2495'),
    ((select id from public.product_info where sku = 'WL48MARINO'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '2195'),
    ((select id from public.product_info where sku = 'WL24SHADE'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '1195'),
    ((select id from public.product_info where sku = 'IS27MUR-MAYFLOWERLED'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '2495'),
    ((select id from public.product_info where sku = 'WL27MUR-MAYFLOWERLED'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '2095'),
    ((select id from public.product_info where sku = 'IS27MUR-SERENITYLED'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '2495'),
    ((select id from public.product_info where sku = 'IS34MUR-SERENITYLED'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '2995'),
    ((select id from public.product_info where sku = 'WL27MUR-SERENITYLED'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '2095'),
    ((select id from public.product_info where sku = 'IS24MOONCRYS'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '1595'),
    ((select id from public.product_info where sku = 'WL32INSERT-BAF'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '1395'),
    ((select id from public.product_info where sku = 'IS34MUR-MAYFLOWERLED'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '2995'),
    ((select id from public.product_info where sku = 'WL24STREAMLINE-BLU'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '1395'),
    ((select id from public.product_info where sku = 'WL36MARINO'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '1995'),
    ((select id from public.product_info where sku = 'IS27MUR-ALFALED'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '2495'),
    ((select id from public.product_info where sku = 'AC-RCRS'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '439'),
    ((select id from public.product_info where sku = 'AC-RRCF'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '149'),
    ((select id from public.product_info where sku = 'AC-LED-INCBULB-CW'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '45'),
    ((select id from public.product_info where sku = 'AC-LED-INCBULB-WW'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '45'),
    ((select id from public.product_info where sku = 'IS27MUR-SNOWLED'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '2495'),
    ((select id from public.product_info where sku = 'IS27MUR-FORTUNALED'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '2495'),
    ((select id from public.product_info where sku = 'IS69STREAMLINE'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '3295'),
    ((select id from public.product_info where sku = 'WL30SHADOW-BLK'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '3295'),
    ((select id from public.product_info where sku = 'WL30SHADOW-WHT'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '3295'),
    ((select id from public.product_info where sku = 'WL36VITTORIA'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '3195'),
    ((select id from public.product_info where sku = 'WL34FOLIO-BLK'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '1495'),
    ((select id from public.product_info where sku = 'WL34FOLIO-WHT'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '1495'),
    ((select id from public.product_info where sku = 'AC-FLRBULB48'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '35'),
    ((select id from public.product_info where sku = 'IS40VITTORIA'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '3495'),
    ((select id from public.product_info where sku = 'IS36SAVONA-INOX'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '2595'),
    ((select id from public.product_info where sku = 'IS36SAVONA-WHT'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '2595'),
    ((select id from public.product_info where sku = 'WL24RACCOLTA'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '1395'),
    ((select id from public.product_info where sku = 'WL30RACCOLTA'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '1495'),
    ((select id from public.product_info where sku = 'WL36DECORSA'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '1695'),
    ((select id from public.product_info where sku = 'IS36VIALE-BLK'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '2495'),
    ((select id from public.product_info where sku = 'IS48VIALE-BLK'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '2795'),
    ((select id from public.product_info where sku = 'IS36VIALE-WHT'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '2495'),
    ((select id from public.product_info where sku = 'IS48VIALE-WHT'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '2795'),
    ((select id from public.product_info where sku = 'IS60SILVANA'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '2995'),
    ((select id from public.product_info where sku = 'IS60RECANO'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '3495'),
    ((select id from public.product_info where sku = 'IS40RECANO'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '3195'),
    ((select id from public.product_info where sku = 'IS19NATALIE'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '2795'),
    ((select id from public.product_info where sku = 'IS36TACTIO'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '2995'),
    ((select id from public.product_info where sku = 'WL36TRANS-BLK'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '2295'),
    ((select id from public.product_info where sku = 'AC-SILENT-DUCTLESS'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '295'),
    ((select id from public.product_info where sku = 'IS24ARENA-INOX'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '2995'),
    ((select id from public.product_info where sku = 'IS24ARENA-BLK'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '2995'),
    ((select id from public.product_info where sku = 'IS24ARENA-CR'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '3295'),
    ((select id from public.product_info where sku = 'IS21DOME-BLK'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '2595'),
    ((select id from public.product_info where sku = 'IS21DOME-IRON'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '2995'),
    ((select id from public.product_info where sku = 'IS21DOME-WHT'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '2595'),
    ((select id from public.product_info where sku = 'IS40MONOLITH'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '2795'),
    ((select id from public.product_info where sku = 'IS21DOME-BRS'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '2995'),
    ((select id from public.product_info where sku = 'IS21DOME-CR'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '2995'),
    ((select id from public.product_info where sku = 'IS23HALO'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '3295'),
    ((select id from public.product_info where sku = 'IS30ORBIT-BLK'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '3095'),
    ((select id from public.product_info where sku = 'IS30ORBIT-INOX'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '3095'),
    ((select id from public.product_info where sku = 'IS48PERIMETER-WHT'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '3495'),
    ((select id from public.product_info where sku = 'IS48PERIMETER-INOX'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '3495'),
    ((select id from public.product_info where sku = 'WL36TRANS-CEM'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '2295'),
    ((select id from public.product_info where sku = 'AC-CARB-A'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '50'),
    ((select id from public.product_info where sku = 'IS14JUPITER-WHT'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '1995'),
    ((select id from public.product_info where sku = 'IS14JUPITER-LT'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '2395'),
    ((select id from public.product_info where sku = 'IS14JUPITER'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '1895'),
    ((select id from public.product_info where sku = 'WL36EVO'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '1595'),
    ((select id from public.product_info where sku = 'AC-STRINGS-6'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '75'),
    ((select id from public.product_info where sku = 'AC-STRINGS-4'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '75'),
    ((select id from public.product_info where sku = 'WL36SLIDE'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '1495'),
    ((select id from public.product_info where sku = 'WL46SLIDE'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '1695'),
    ((select id from public.product_info where sku = 'AC-SHELF-EUROPE'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '1395'),
    ((select id from public.product_info where sku = 'IS14JUPITER-BLK'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '1995'),
    ((select id from public.product_info where sku = 'WL36POSITANO-BLK'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '1795'),
    ((select id from public.product_info where sku = 'AC-EXT-MURWL'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '245.00'),
    ((select id from public.product_info where sku = 'IS36POSITANO-BLK'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '1995'),
    ((select id from public.product_info where sku = 'IS16LOFT-BLK'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '1695'),
    ((select id from public.product_info where sku = 'IS69STREAMLINE-LH'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '3495'),
    ((select id from public.product_info where sku = 'IS69STREAMLINE-RH'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '3495'),
    ((select id from public.product_info where sku = 'AC-ADAPTER'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '65'),
    ((select id from public.product_info where sku = 'AC-MAGICSTEEL'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '29'),
    ((select id from public.product_info where sku = 'AC-STRINGS-2'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '75'),
    ((select id from public.product_info where sku = 'AC-SHELF-L'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '750'),
    ((select id from public.product_info where sku = 'WL30KELVIN-BLK'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '1495'),
    ((select id from public.product_info where sku = 'WL30KELVIN-INOX'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '1495'),
    ((select id from public.product_info where sku = 'WL30SLIDE'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '1495'),
    ((select id from public.product_info where sku = 'WL36GABI'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '1795'),
    ((select id from public.product_info where sku = 'IS36BALANCE-CHR'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '2995'),
    ((select id from public.product_info where sku = 'IS36BALANCE-WHT'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '2995'),
    ((select id from public.product_info where sku = 'IS36BALANCE-BLK'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '2995'),
    ((select id from public.product_info where sku = 'AC-STRINGS-8'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '75'),
    ((select id from public.product_info where sku = 'WL36QUEST-WHT'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '1795'),
    ((select id from public.product_info where sku = 'IS69TURO-BLK'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '3995'),
    ((select id from public.product_info where sku = 'IS36TURO-BLK'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '2995'),
    ((select id from public.product_info where sku = 'IS36TURO-INOX'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '2995'),
    ((select id from public.product_info where sku = 'IS36TRIM-WHT'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '2995'),
    ((select id from public.product_info where sku = 'WL36FANO-GM'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '2695'),
    ((select id from public.product_info where sku = 'WL48FANO-GM'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '2995'),
    ((select id from public.product_info where sku = 'WL36LORENZO-BLK'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '2995'),
    ((select id from public.product_info where sku = 'IS14JUPITER-CPR'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '1995'),
    ((select id from public.product_info where sku = 'IS14JUPITER-GLD'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '1995'),
    ((select id from public.product_info where sku = 'WL36CAMINO-GM'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '1895'),
    ((select id from public.product_info where sku = 'WL36CASCADE'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '1995'),
    ((select id from public.product_info where sku = 'WL36KNOX-GM'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '2195'),
    ((select id from public.product_info where sku = 'WL36COUNTER-GM'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '1995'),
    ((select id from public.product_info where sku = 'WL48COUNTER-GM'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '2195'),
    ((select id from public.product_info where sku = 'WL36SPHINX'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '1795'),
    ((select id from public.product_info where sku = 'WL48SPHINX'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '2095'),
    ((select id from public.product_info where sku = 'WL36PYRAMID'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '1795'),
    ((select id from public.product_info where sku = 'IS48ALINA-WHT'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '3495'),
    ((select id from public.product_info where sku = 'WL36CASTLE'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '1795'),
    ((select id from public.product_info where sku = 'WL48CASCADE'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '2295'),
    ((select id from public.product_info where sku = 'WL48CASTLE'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '2095'),
    ((select id from public.product_info where sku = 'AC-CARB-ART'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '50'),
    ((select id from public.product_info where sku = 'IS48KNOX-BLK'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '3995'),
    ((select id from public.product_info where sku = 'IS50STEALTH-BLK'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '3095'),
    ((select id from public.product_info where sku = 'WL36COUNTRY'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '1795'),
    ((select id from public.product_info where sku = 'WL48COUNTRY'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '2095'),
    ((select id from public.product_info where sku = 'WL36SPLASH-BLK'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '3595'),
    ((select id from public.product_info where sku = 'WL36SPLASH-WHT'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '3595'),
    ((select id from public.product_info where sku = 'WL36RAVEN-BLK'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '2495'),
    ((select id from public.product_info where sku = 'IS40MAGNOLIA'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '3995'),
    ((select id from public.product_info where sku = 'WL20CAMPANA-CPR'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '2195'),
    ((select id from public.product_info where sku = 'WL36ARTTEMPO-WHT'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '1795'),
    ((select id from public.product_info where sku = 'WL48ARTTEMPO-WHT'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '2095'),
    ((select id from public.product_info where sku = 'WL36NOVA-GM'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '2195'),
    ((select id from public.product_info where sku = 'AC-MCIF'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '375'),
    ((select id from public.product_info where sku = 'AC-24NOVA-SHELF'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '1195'),
    ((select id from public.product_info where sku = 'IS48TURO-INOX'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '3295'),
    ((select id from public.product_info where sku = 'IS48HORIZON-BLK'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '3695'),
    ((select id from public.product_info where sku = 'IS63HORIZON-BAR'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '3995'),
    ((select id from public.product_info where sku = 'IS63HORIZON-SHELF'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '4195'),
    ((select id from public.product_info where sku = 'IS48TURO-BLK'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '3295'),
    ((select id from public.product_info where sku = 'WL36FLIPPER'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '1995'),
    ((select id from public.product_info where sku = 'WL48FLIPPER'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '2095'),
    ((select id from public.product_info where sku = 'WL36LINEA'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '1795'),
    ((select id from public.product_info where sku = 'WL36AESTHETICS-BLK'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '2995'),
    ((select id from public.product_info where sku = 'AC-CARB-G'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '50'),
    ((select id from public.product_info where sku = 'AC-CARB-HORIZON'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '149'),
    ((select id from public.product_info where sku = 'AC-CARB-FLIPPER'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '149'),
    ((select id from public.product_info where sku = 'AC-CARB-MAGNOLIA'), (select id from public.price_type where url = 'sale'), (select id from public.currency where url = 'usd'), '149');

-- get data
select * from public.price;

-- get data from woocommerce Database
-- SELECT
--     CONCAT ("(select id from public.product_info where sku = '", T.sku, "')") as product_id,
--     CONCAT ("(select id from public.price_type where url = '", T.price_type, "')") as price_type_id,
--     CONCAT ("(select id from public.currency where url = '", T.currency, "')") as currency_id,
--     T.price
-- FROM (SELECT
--           pm1.meta_value as sku,
--           CASE
--               WHEN pm3.meta_value > 0 AND pm3.meta_value IS NOT NULL THEN 'sale'
--               ELSE 'regular'
--               END as price_type,
--           CASE
--               WHEN pm3.meta_value > 0 AND pm3.meta_value IS NOT NULL THEN pm3.meta_value
--               ELSE pm2.meta_value
--               END as price,
--           (SELECT option_value FROM wp_options WHERE option_name = 'woocommerce_currency') as currency
--       FROM
--           wp_posts p
--               LEFT JOIN
--           wp_postmeta pm1 ON p.ID = pm1.post_id AND pm1.meta_key = '_sku'
--               LEFT JOIN
--           wp_postmeta pm2 ON p.ID = pm2.post_id AND pm2.meta_key = '_regular_price'
--               LEFT JOIN
--           wp_postmeta pm3 ON p.ID = pm3.post_id AND pm3.meta_key = '_sale_price'
--       WHERE
--           p.post_type = 'product') as T;