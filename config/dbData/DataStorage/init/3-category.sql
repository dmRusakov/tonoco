/*drop if needed */
DROP TABLE IF EXISTS public.category;

CREATE TABLE IF NOT EXISTS public.category
(
    category_id       uuid unique             DEFAULT uuid_generate_v4(),
    name              character varying(255) COLLATE pg_catalog."default" NOT NULL,
    category_key      character varying(255) COLLATE pg_catalog."default" NOT NULL,
    short_description character varying(255)  default null,
    description       character varying(4000) default null,
    "order"           integer                 default null,

    created_at        timestamp               default now(),
    created_by        uuid                    default null REFERENCES public.user (user_id),
    updated_at        timestamp               default now(),
    updated_by        uuid                    default null REFERENCES public.user (user_id)
);

ALTER TABLE ONLY public.category ADD CONSTRAINT category_pkey PRIMARY KEY (category_id);
CREATE UNIQUE INDEX category_id ON public.category(category_id);
CREATE UNIQUE INDEX category_key ON public.category(category_key);
ALTER TABLE public.category OWNER TO postgres;

-- insert data
INSERT INTO public.category (category_id, category_key, name, short_description, description, "order")
VALUES   ('1f484cda-c00e-4ed8-a325-9c5e035f992f', 'cat_1', 'Category 1', 'Short description 1', 'Description 1', 1),
         ('d338d096-1ee7-4a57-b334-798b32a01fda', 'cat_2', 'Category 2', 'Short description 2', 'Description 2', 2),
         ('d48eb997-6882-4ff1-b7e7-d8c05aab2d29', 'cat_3', 'Category 3', 'Short description 3', 'Description 3', 3),
         ('154e1bcd-4f51-433d-b2a5-bd41f84209a8', 'cat_4', 'Category 4', 'Short description 4', 'Description 4', 4),
         ('95bfde1c-a06e-45d9-89af-0f467905c941', 'cat_5', 'Category 5', 'Short description 5', 'Description 5', 5);

