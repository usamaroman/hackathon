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
-- Name: priority; Type: TYPE; Schema: public; Owner: chechyotka
--

CREATE TYPE public.priority AS ENUM (
    'низкий',
    'средний',
    'высокий'
);


ALTER TYPE public.priority OWNER TO chechyotka;

--
-- Name: role; Type: TYPE; Schema: public; Owner: chechyotka
--

CREATE TYPE public.role AS ENUM (
    'user',
    'admin'
);


ALTER TYPE public.role OWNER TO chechyotka;

--
-- Name: status; Type: TYPE; Schema: public; Owner: chechyotka
--

CREATE TYPE public.status AS ENUM (
    'надо сделать',
    'в процессе',
    'выполнено'
);


ALTER TYPE public.status OWNER TO chechyotka;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: projects; Type: TABLE; Schema: public; Owner: chechyotka
--

CREATE TABLE public.projects (
    id uuid NOT NULL,
    title text NOT NULL,
    description text NOT NULL,
    start timestamp without time zone NOT NULL,
    "end" timestamp without time zone NOT NULL
);


ALTER TABLE public.projects OWNER TO chechyotka;

--
-- Name: projects_tasks; Type: TABLE; Schema: public; Owner: chechyotka
--

CREATE TABLE public.projects_tasks (
    project_id uuid,
    task_id uuid
);


ALTER TABLE public.projects_tasks OWNER TO chechyotka;

--
-- Name: tasks; Type: TABLE; Schema: public; Owner: chechyotka
--

CREATE TABLE public.tasks (
    id uuid NOT NULL,
    title text NOT NULL,
    description text NOT NULL,
    start timestamp without time zone NOT NULL,
    "end" timestamp without time zone NOT NULL,
    difficulty integer NOT NULL,
    priority public.priority NOT NULL,
    status public.status NOT NULL,
    CONSTRAINT tasks_difficulty_check CHECK (((difficulty > 0) AND (difficulty <= 100)))
);


ALTER TABLE public.tasks OWNER TO chechyotka;

--
-- Name: users; Type: TABLE; Schema: public; Owner: chechyotka
--

CREATE TABLE public.users (
    id uuid NOT NULL,
    email text NOT NULL,
    password text NOT NULL,
    role public.role DEFAULT 'user'::public.role
);


ALTER TABLE public.users OWNER TO chechyotka;

--
-- Data for Name: projects; Type: TABLE DATA; Schema: public; Owner: chechyotka
--

COPY public.projects (id, title, description, start, "end") FROM stdin;
98130782-9811-4ec2-b60f-ac14ddf4ad3e	123	123	2023-12-13 00:00:00	2023-12-14 00:00:00
a2f4bb07-1268-4b06-ae10-9ab48d80764f	finish hackathon	qwerty	2023-10-27 00:00:00	2023-10-28 00:00:00
\.


--
-- Data for Name: projects_tasks; Type: TABLE DATA; Schema: public; Owner: chechyotka
--

COPY public.projects_tasks (project_id, task_id) FROM stdin;
98130782-9811-4ec2-b60f-ac14ddf4ad3e	95184e07-5b89-40ec-a759-0559dcae4b34
\.


--
-- Data for Name: tasks; Type: TABLE DATA; Schema: public; Owner: chechyotka
--

COPY public.tasks (id, title, description, start, "end", difficulty, priority, status) FROM stdin;
a14d34e1-a755-42bd-8ad8-f8d6be6b7def	куптьь коду	хочу колу сильно	2023-10-27 00:00:00	2023-10-28 00:00:00	100	высокий	надо сделать
d234a264-e6f3-414a-b4a8-f161a8bab13b	куптьь коду	хочу колу сильно	2023-10-27 00:00:00	2023-10-28 00:00:00	100	высокий	надо сделать
95184e07-5b89-40ec-a759-0559dcae4b34	куптьь коду2	хочу колу сильно	2023-10-27 00:00:00	2023-10-28 00:00:00	100	высокий	надо сделать
\.


--
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: chechyotka
--

COPY public.users (id, email, password, role) FROM stdin;
be19d517-7011-42df-aba7-9f4435c80aee	a@gmail.com	$2a$10$4H99fEJ/mA70K3cynLQ8/e6gckOx1Gd/50wHkoDMXROJDjcM.7yPO	user
\.


--
-- Name: projects projects_pkey; Type: CONSTRAINT; Schema: public; Owner: chechyotka
--

ALTER TABLE ONLY public.projects
    ADD CONSTRAINT projects_pkey PRIMARY KEY (id);


--
-- Name: tasks tasks_pkey; Type: CONSTRAINT; Schema: public; Owner: chechyotka
--

ALTER TABLE ONLY public.tasks
    ADD CONSTRAINT tasks_pkey PRIMARY KEY (id);


--
-- Name: users users_email_key; Type: CONSTRAINT; Schema: public; Owner: chechyotka
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_email_key UNIQUE (email);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: chechyotka
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- Name: projects_tasks projects_tasks_project_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: chechyotka
--

ALTER TABLE ONLY public.projects_tasks
    ADD CONSTRAINT projects_tasks_project_id_fkey FOREIGN KEY (project_id) REFERENCES public.projects(id);


--
-- Name: projects_tasks projects_tasks_task_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: chechyotka
--

ALTER TABLE ONLY public.projects_tasks
    ADD CONSTRAINT projects_tasks_task_id_fkey FOREIGN KEY (task_id) REFERENCES public.tasks(id);


--
-- PostgreSQL database dump complete
--

