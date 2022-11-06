--
-- PostgreSQL database dump
--

-- Dumped from database version 14.5 (Ubuntu 14.5-2.pgdg22.04+2)
-- Dumped by pg_dump version 15.0 (Ubuntu 15.0-1.pgdg22.04+1)

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
-- Name: public; Type: SCHEMA; Schema: -; Owner: postgres
--

-- *not* creating schema, since initdb creates it


ALTER SCHEMA public OWNER TO postgres;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: delivery; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.delivery (
    order_uid character varying(50),
    name character varying(50),
    phone character varying(50),
    zip character varying(50),
    city character varying(50),
    address character varying(50),
    region character varying(50),
    email character varying(50)
);


ALTER TABLE public.delivery OWNER TO postgres;

--
-- Name: items; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.items (
    order_uid character varying(50),
    chrt_id bigint,
    track_number character varying(50),
    price bigint,
    rid character varying(50),
    name character varying(50),
    sale smallint,
    size character varying(50),
    total_price bigint,
    nm_id bigint,
    brand character varying(50),
    status smallint
);


ALTER TABLE public.items OWNER TO postgres;

--
-- Name: orders; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.orders (
    order_uid character varying(50) NOT NULL,
    track_number character varying(50),
    entry character varying(50),
    locale character varying(50),
    internal_signature character varying(50),
    customer_id character varying(50),
    delivery_service character varying(50),
    shardkey character varying(50),
    sm_id bigint,
    date_created character varying(50),
    oof_shred character varying(50)
);


ALTER TABLE public.orders OWNER TO postgres;

--
-- Name: payment; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.payment (
    order_uid character varying(50),
    transaction character varying(50),
    request_id character varying(50),
    currency character varying(50),
    provider character varying(50),
    amount bigint,
    payment_dt bigint,
    bank character varying(50),
    delivery_cost bigint,
    goods_total bigint,
    custom_fee smallint
);


ALTER TABLE public.payment OWNER TO postgres;

--
-- Data for Name: delivery; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.delivery (order_uid, name, phone, zip, city, address, region, email) FROM stdin;
1	Test Testov	+9720000000	2639809	Kiryat Mozkin	Ploshad Mira 11	Kraiot	test@gmail.com
2	Test Testov	+9720000000	2639809	Kiryat Mozkin	Ploshad Mira 12	Kraiot	test@gmail.com
3	Test Testov	+9720000000	2639809	Kiryat Mozkin	Ploshad Mira 13	Kraiot	test@gmail.com
4	Test Testov	+9720000000	2639809	Kiryat Mozkin	Ploshad Mira 14	Kraiot	test@gmail.com
5	Test Testov	+9720000000	2639809	Kiryat Mozkin	Ploshad Mira 15	Kraiot	test@gmail.com
6	Test Testov	+9720000000	2639809	Kiryat Mozkin	Ploshad Mira 16	Kraiot	test@gmail.com
b563feb7b2b84b6test	Test Testov	+9720000000	2639809	Kiryat Mozkin	Ploshad Mira	Kraiot	test@gmail.com
\.


--
-- Data for Name: items; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.items (order_uid, chrt_id, track_number, price, rid, name, sale, size, total_price, nm_id, brand, status) FROM stdin;
1	9934930	WBILMTESTTRACK	1001	ab4219087a764ae0btest	Mascaras	30	0	317	2389212	Vivienne Sabo	202
2	9934930	WBILMTESTTRACK	1002	ab4219087a764ae0btest	Mascaras	30	0	317	2389212	Vivienne Sabo	202
2	9934930	WBILMTESTTRACK	453	ab4219087a764ae0btest	Mascaras	30	0	317	2389212	Vivienne Sabo	202
2	9934930	WBILMTESTTRACK	453	ab4219087a764ae0btest	Mascaras	30	0	317	2389212	Vivienne Sabo	202
3	9934930	WBILMTESTTRACK	1003	ab4219087a764ae0btest	Mascaras	30	0	317	2389212	Vivienne Sabo	202
3	9934930	WBILMTESTTRACK	453	ab4219087a764ae0btest	Mascaras	30	0	317	2389212	Vivienne Sabo	202
3	9934930	WBILMTESTTRACK	453	ab4219087a764ae0btest	Mascaras	30	0	317	2389212	Vivienne Sabo	202
4	9934930	WBILMTESTTRACK	1004	ab4219087a764ae0btest	Mascaras	30	0	317	2389212	Vivienne Sabo	202
4	9934930	WBILMTESTTRACK	453	ab4219087a764ae0btest	Mascaras	30	0	317	2389212	Vivienne Sabo	202
4	9934930	WBILMTESTTRACK	453	ab4219087a764ae0btest	Mascaras	30	0	317	2389212	Vivienne Sabo	202
5	9934930	WBILMTESTTRACK	1005	ab4219087a764ae0btest	Mascaras	30	0	317	2389212	Vivienne Sabo	202
5	9934930	WBILMTESTTRACK	453	ab4219087a764ae0btest	Mascaras	30	0	317	2389212	Vivienne Sabo	202
5	9934930	WBILMTESTTRACK	453	ab4219087a764ae0btest	Mascaras	30	0	317	2389212	Vivienne Sabo	202
6	9934930	WBILMTESTTRACK	1006	ab4219087a764ae0btest	Mascaras	30	0	317	2389212	Vivienne Sabo	202
b563feb7b2b84b6test	9934930	WBILMTESTTRACK	453	ab4219087a764ae0btest	Mascaras	30	0	317	2389212	Vivienne Sabo	202
\.


--
-- Data for Name: orders; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.orders (order_uid, track_number, entry, locale, internal_signature, customer_id, delivery_service, shardkey, sm_id, date_created, oof_shred) FROM stdin;
1	WBILMTESTTRACK	WBIL	en		test	meest	9	99	2021-11-26T06:22:19Z	1
2	WBILMTESTTRACK	WBIL	en		test	meest	9	99	2021-11-26T06:22:19Z	1
3	WBILMTESTTRACK	WBIL	en		test	meest	9	99	2021-11-26T06:22:19Z	1
4	WBILMTESTTRACK	WBIL	en		test	meest	9	99	2021-11-26T06:22:19Z	1
5	WBILMTESTTRACK	WBIL	en		test	meest	9	99	2021-11-26T06:22:19Z	1
6	WBILMTESTTRACK	WBIL	en		test	meest	9	99	2021-11-26T06:22:19Z	1
b563feb7b2b84b6test	WBILMTESTTRACK	WBIL	en		test	meest	9	99	2021-11-26T06:22:19Z	1
\.


--
-- Data for Name: payment; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.payment (order_uid, transaction, request_id, currency, provider, amount, payment_dt, bank, delivery_cost, goods_total, custom_fee) FROM stdin;
1	b563feb7b2b84b6test		USD	wbpay	100	1817	alpha	1500	317	0
2	b563feb7b2b84b6test		USD	wbpay	400	1817	alpha	1500	317	0
3	b563feb7b2b84b6test		USD	wbpay	900	1817	alpha	1500	317	0
4	b563feb7b2b84b6test		USD	wbpay	1600	1817	alpha	1500	317	0
5	b563feb7b2b84b6test		USD	wbpay	2500	1817	alpha	1500	317	0
6	b563feb7b2b84b6test		USD	wbpay	3600	1817	alpha	1500	317	0
b563feb7b2b84b6test	b563feb7b2b84b6test		USD	wbpay	1637907727	1817	alpha	1500	317	0
\.


--
-- Name: orders orders_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.orders
    ADD CONSTRAINT orders_pkey PRIMARY KEY (order_uid);


--
-- Name: delivery delivery_order_uid_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.delivery
    ADD CONSTRAINT delivery_order_uid_fkey FOREIGN KEY (order_uid) REFERENCES public.orders(order_uid) ON DELETE CASCADE;


--
-- Name: items items_order_uid_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.items
    ADD CONSTRAINT items_order_uid_fkey FOREIGN KEY (order_uid) REFERENCES public.orders(order_uid) ON DELETE CASCADE;


--
-- Name: payment payment_order_uid_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.payment
    ADD CONSTRAINT payment_order_uid_fkey FOREIGN KEY (order_uid) REFERENCES public.orders(order_uid) ON DELETE CASCADE;


--
-- Name: SCHEMA public; Type: ACL; Schema: -; Owner: postgres
--

REVOKE USAGE ON SCHEMA public FROM PUBLIC;
GRANT ALL ON SCHEMA public TO PUBLIC;


--
-- PostgreSQL database dump complete
--

