/*drop if needed */
DROP TABLE IF EXISTS public.specification;

CREATE TABLE IF NOT EXISTS public.specification
(
    specification_id  uuid unique            DEFAULT uuid_generate_v4(),
    name              character varying(255) COLLATE pg_catalog."default" NOT NULL,
    specification_key               character varying(255) COLLATE pg_catalog."default" NOT NULL,

    created_at        timestamp              default now(),
    created_by        uuid                   default null REFERENCES public.user (user_id),
    updated_at        timestamp              default now(),
    updated_by        uuid                   default null REFERENCES public.user (user_id)
);

ALTER TABLE ONLY public.specification ADD CONSTRAINT specification_pkey PRIMARY KEY (specification_id);
CREATE UNIQUE INDEX specification_id ON public.specification(specification_id);
CREATE UNIQUE INDEX specification_key ON public.specification(specification_key);
ALTER TABLE public.specification OWNER TO postgres;

-- insert data
INSERT INTO public.specification (name, specification_key)
VALUES ('specification 1', 'specification_1'),
       ('specification 2', 'specification_2'),
       ('specification 3', 'specification_3'),
       ('specification 4', 'specification_4'),
       ('specification 5', 'specification_5'),
       ('specification 6', 'specification_6'),
       ('specification 7', 'specification_7'),
       ('specification 8', 'specification_8'),
       ('specification 9', 'specification_9'),
       ('specification 10', 'specification_10'),
       ('specification 11', 'specification_11'),
       ('specification 12', 'specification_12'),
       ('specification 13', 'specification_13'),
       ('specification 14', 'specification_14'),
       ('specification 15', 'specification_15'),
       ('specification 16', 'specification_16'),
       ('specification 17', 'specification_17'),
       ('specification 18', 'specification_18'),
       ('specification 19', 'specification_19'),
       ('specification 20', 'specification_20');

