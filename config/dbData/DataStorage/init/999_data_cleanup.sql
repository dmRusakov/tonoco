-- image
update public.image set origin_path = substring(origin_path from '^[^?]+') where origin_path like '%?%';
update public.image set extension = substring(origin_path from '[^.]+$') where 1=1;
update public.image set filename =  substring(origin_path from '[^/]+(?=\.[^.]+$)') where 1=1;
delete from public.image where extension not in ('jpg', 'png', 'gif', 'webp', 'svg');
update public.image set copyright = '©FuturoFuturo.com' where 1=1;
update public.image set creator = 'FuturoFuturo.com' where 1=1;
update public.image set rating = 5 where 1=1;
select * from public.image;
DO $$
    DECLARE
        image_file_name text := 'Modern_Black-kitchen-hood-36-inch-Positano-Black-island-range-hood-Futuro-Futuro-range-hoods-image-05_wm.jpg';
    BEGIN
        delete from product_image where image_id = (select id from image where origin_path like '%' || image_file_name || '%');
        delete from image where origin_path like '%' || image_file_name || '%';
    END $$;

DO $$
    DECLARE
        image_file_name text := 'Black-kitchen-hood-36-inch-Positano-Black-wall-range-hood-Futuro-Futuro-range-hoods-image-13_wm.jpg';
    BEGIN
        delete from product_image where image_id = (select id from image where origin_path like '%' || image_file_name || '%');
        delete from image where origin_path like '%' || image_file_name || '%';
    END $$;

-- product image
delete from public.product_image where product_id is null or image_id is null;