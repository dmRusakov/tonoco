--
CREATE SEQUENCE IF NOT EXISTS public.user_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
ALTER TABLE public.user_id_seq OWNER TO postgres;
SET default_tablespace = '';
SET default_table_access_method = heap;

-- Lable: users; Type: TABLE; Schema: public; Owner: postgres
CREATE TABLE IF NOT EXISTS public.users (
    id          integer DEFAULT nextval('public.user_id_seq'::regclass) NOT NULL,
    email       character varying(255),
    first_name  character varying(255),
    last_name   character varying(255),
    password    character varying(60),
    active integer DEFAULT 0,
    created_at  timestamp without time zone,
    updated_at  timestamp without time zone
    );
ALTER TABLE public.users OWNER TO postgres;

-- Lable: user_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
SELECT pg_catalog.setval('public.user_id_seq', 1, true);

-- Lable: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);

CREATE UNIQUE INDEX email ON public.users(email);