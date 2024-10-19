-- image
update public.image set origin_path = substring(origin_path from '^[^?]+') where origin_path like '%?%';
update public.image set extension = substring(origin_path from '[^.]+$') where 1=1;
update public.image set filename =  substring(origin_path from '[^/]+(?=\.[^.]+$)') where 1=1;
delete from public.image where extension not in ('jpg', 'png', 'gif', 'webp', 'svg');
update public.image set copyright = '©FuturoFuturo.com' where 1=1;
update public.image set creator = 'FuturoFuturo.com' where 1=1;
update public.image set rating = 5 where 1=1;
select * from public.image;

-- product image
delete from public.product_image where product_id is null or image_id is null;