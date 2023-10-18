-- drop if needed
DROP TABLE IF EXISTS public.image;

-- create table
CREATE TABLE IF NOT EXISTS public.image
(
    image_id  uuid unique            DEFAULT uuid_generate_v4(),
    file_name         varchar(255)          not null,
    file_type         varchar(5)          not null,
    is_support_webp   boolean               default false,
    is_support_avif   boolean               default false,

    folder_id         uuid                   default null REFERENCES public.folder (folder_id),

    thumbnail         varchar(255)                 default null,
    small             varchar(255)                 default null,
    medium            varchar(255)                 default null,
    large             varchar(255)                 default null,
    original          varchar(255)                 default null,

    created_at        timestamp              default now(),
    created_by        uuid                   default null REFERENCES public.user (user_id),
    updated_at        timestamp              default now(),
    updated_by        uuid                   default null REFERENCES public.user (user_id)
);

-- index, constraint and ownership
ALTER TABLE ONLY public.image ADD CONSTRAINT specification_pkey PRIMARY KEY (image_id);
CREATE UNIQUE INDEX specification_id ON public.image(image_id);
ALTER TABLE public.image OWNER TO postgres;

-- insert data
INSERT INTO public.image

