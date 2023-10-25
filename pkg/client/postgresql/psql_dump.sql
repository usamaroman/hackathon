--
-- PostgreSQL database dump
--

-- Dumped from database version 15.3
-- Dumped by pg_dump version 15.3

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
-- Name: role; Type: TYPE; Schema: public; Owner: postgres
--

CREATE TYPE public.role AS ENUM (
    'user',
    'admin'
);


ALTER TYPE public.role OWNER TO postgres;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: projects; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.projects (
    id uuid NOT NULL,
    title text NOT NULL,
    description text NOT NULL,
    start timestamp without time zone NOT NULL,
    "end" timestamp without time zone NOT NULL
);


ALTER TABLE public.projects OWNER TO postgres;

--
-- Name: users; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.users (
    id uuid NOT NULL,
    email text NOT NULL,
    password text NOT NULL,
    role public.role DEFAULT 'user'::public.role
);


ALTER TABLE public.users OWNER TO postgres;

--
-- Data for Name: projects; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.projects (id, title, description, start, "end") FROM stdin;
98130782-9811-4ec2-b60f-ac14ddf4ad3e	123	123	2023-12-13 00:00:00	2023-12-14 00:00:00
\.


--
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.users (id, email, password, role) FROM stdin;
be19d517-7011-42df-aba7-9f4435c80aee	a@gmail.com	$2a$10$4H99fEJ/mA70K3cynLQ8/e6gckOx1Gd/50wHkoDMXROJDjcM.7yPO	user
\.


--
-- Name: projects projects_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.projects
    ADD CONSTRAINT projects_pkey PRIMARY KEY (id);


--
-- Name: users users_email_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_email_key UNIQUE (email);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- PostgreSQL database dump complete
--

