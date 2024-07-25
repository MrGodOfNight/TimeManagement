--
-- PostgreSQL database cluster dump
--

-- Started on 2024-07-26 00:29:27

SET default_transaction_read_only = off;

SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;

--
-- Roles
--

CREATE ROLE postgres;
ALTER ROLE postgres WITH SUPERUSER INHERIT CREATEROLE CREATEDB LOGIN REPLICATION BYPASSRLS;

--
-- User Configurations
--








--
-- Databases
--

--
-- Database "template1" dump
--

\connect template1

--
-- PostgreSQL database dump
--

-- Dumped from database version 16.0
-- Dumped by pg_dump version 16.0

-- Started on 2024-07-26 00:29:27

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

-- Completed on 2024-07-26 00:29:28

--
-- PostgreSQL database dump complete
--

--
-- Database "time" dump
--

--
-- PostgreSQL database dump
--

-- Dumped from database version 16.0
-- Dumped by pg_dump version 16.0

-- Started on 2024-07-26 00:29:28

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

--
-- TOC entry 4810 (class 1262 OID 25565)
-- Name: time; Type: DATABASE; Schema: -; Owner: postgres
--

CREATE DATABASE "time" WITH TEMPLATE = template0 ENCODING = 'UTF8' LOCALE_PROVIDER = libc LOCALE = 'Russian_Russia.1251';


ALTER DATABASE "time" OWNER TO postgres;

\connect "time"

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- TOC entry 220 (class 1259 OID 25589)
-- Name: report; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.report (
    id integer NOT NULL,
    text text NOT NULL,
    time_id integer NOT NULL
);


ALTER TABLE public.report OWNER TO postgres;

--
-- TOC entry 4811 (class 0 OID 0)
-- Dependencies: 220
-- Name: TABLE report; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON TABLE public.report IS 'table that stores all reports in markdown format';


--
-- TOC entry 4812 (class 0 OID 0)
-- Dependencies: 220
-- Name: COLUMN report.text; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.report.text IS 'is the report text itself in markdown format';


--
-- TOC entry 219 (class 1259 OID 25588)
-- Name: report_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.report_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.report_id_seq OWNER TO postgres;

--
-- TOC entry 4813 (class 0 OID 0)
-- Dependencies: 219
-- Name: report_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.report_id_seq OWNED BY public.report.id;


--
-- TOC entry 216 (class 1259 OID 25567)
-- Name: user; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public."user" (
    id integer NOT NULL,
    login text NOT NULL,
    password text NOT NULL,
    reg_date date DEFAULT CURRENT_DATE NOT NULL,
    end_date date NOT NULL,
    last_time timestamp without time zone NOT NULL,
    admin integer
);


ALTER TABLE public."user" OWNER TO postgres;

--
-- TOC entry 4814 (class 0 OID 0)
-- Dependencies: 216
-- Name: TABLE "user"; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON TABLE public."user" IS 'this table of user accounts';


--
-- TOC entry 4815 (class 0 OID 0)
-- Dependencies: 216
-- Name: COLUMN "user".reg_date; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public."user".reg_date IS 'registration date';


--
-- TOC entry 4816 (class 0 OID 0)
-- Dependencies: 216
-- Name: COLUMN "user".end_date; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public."user".end_date IS 'account expiration date';


--
-- TOC entry 4817 (class 0 OID 0)
-- Dependencies: 216
-- Name: COLUMN "user".last_time; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public."user".last_time IS 'last login time';


--
-- TOC entry 4818 (class 0 OID 0)
-- Dependencies: 216
-- Name: COLUMN "user".admin; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public."user".admin IS 'admin lvl if you need. If the level is null, it is a normal user';


--
-- TOC entry 215 (class 1259 OID 25566)
-- Name: users_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.users_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.users_id_seq OWNER TO postgres;

--
-- TOC entry 4819 (class 0 OID 0)
-- Dependencies: 215
-- Name: users_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.users_id_seq OWNED BY public."user".id;


--
-- TOC entry 218 (class 1259 OID 25577)
-- Name: work_time; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.work_time (
    id integer NOT NULL,
    date date NOT NULL,
    start_time time without time zone,
    end_time time without time zone,
    start_lunch time without time zone,
    end_lunch time without time zone,
    user_id integer NOT NULL
);


ALTER TABLE public.work_time OWNER TO postgres;

--
-- TOC entry 217 (class 1259 OID 25576)
-- Name: work_time_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.work_time_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.work_time_id_seq OWNER TO postgres;

--
-- TOC entry 4820 (class 0 OID 0)
-- Dependencies: 217
-- Name: work_time_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.work_time_id_seq OWNED BY public.work_time.id;


--
-- TOC entry 4647 (class 2604 OID 25592)
-- Name: report id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.report ALTER COLUMN id SET DEFAULT nextval('public.report_id_seq'::regclass);


--
-- TOC entry 4644 (class 2604 OID 25570)
-- Name: user id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."user" ALTER COLUMN id SET DEFAULT nextval('public.users_id_seq'::regclass);


--
-- TOC entry 4646 (class 2604 OID 25580)
-- Name: work_time id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.work_time ALTER COLUMN id SET DEFAULT nextval('public.work_time_id_seq'::regclass);


--
-- TOC entry 4804 (class 0 OID 25589)
-- Dependencies: 220
-- Data for Name: report; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.report (id, text, time_id) FROM stdin;
\.


--
-- TOC entry 4800 (class 0 OID 25567)
-- Dependencies: 216
-- Data for Name: user; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public."user" (id, login, password, reg_date, end_date, last_time, admin) FROM stdin;
\.


--
-- TOC entry 4802 (class 0 OID 25577)
-- Dependencies: 218
-- Data for Name: work_time; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.work_time (id, date, start_time, end_time, start_lunch, end_lunch, user_id) FROM stdin;
\.


--
-- TOC entry 4821 (class 0 OID 0)
-- Dependencies: 219
-- Name: report_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.report_id_seq', 1, false);


--
-- TOC entry 4822 (class 0 OID 0)
-- Dependencies: 215
-- Name: users_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.users_id_seq', 1, false);


--
-- TOC entry 4823 (class 0 OID 0)
-- Dependencies: 217
-- Name: work_time_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.work_time_id_seq', 1, false);


--
-- TOC entry 4653 (class 2606 OID 25596)
-- Name: report report_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.report
    ADD CONSTRAINT report_pk PRIMARY KEY (id);


--
-- TOC entry 4649 (class 2606 OID 25575)
-- Name: user users_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."user"
    ADD CONSTRAINT users_pk PRIMARY KEY (id);


--
-- TOC entry 4651 (class 2606 OID 25582)
-- Name: work_time work_time_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.work_time
    ADD CONSTRAINT work_time_pk PRIMARY KEY (id);


--
-- TOC entry 4655 (class 2606 OID 25597)
-- Name: report report_work_time_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.report
    ADD CONSTRAINT report_work_time_fk FOREIGN KEY (time_id) REFERENCES public.work_time(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- TOC entry 4654 (class 2606 OID 25583)
-- Name: work_time work_time_users_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.work_time
    ADD CONSTRAINT work_time_users_fk FOREIGN KEY (user_id) REFERENCES public."user"(id) ON UPDATE CASCADE ON DELETE CASCADE;


-- Completed on 2024-07-26 00:29:28

--
-- PostgreSQL database dump complete
--

-- Completed on 2024-07-26 00:29:28

--
-- PostgreSQL database cluster dump complete
--

