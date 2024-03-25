--
-- PostgreSQL database dump
--

-- Dumped from database version 12.3 (Debian 12.3-1.pgdg100+1)
-- Dumped by pg_dump version 12.16 (Ubuntu 12.16-0ubuntu0.20.04.1)

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
-- Name: aggFt; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public."aggFt" (
    tx_hash character varying(255) NOT NULL,
    from_addr character varying(255) NOT NULL,
    to_addr character varying(255) NOT NULL,
    contract character varying(255) NOT NULL,
    value numeric(64,0) NOT NULL,
    start_ts bigint NOT NULL,
    chain_id bigint NOT NULL,
    end_ts bigint NOT NULL,
    start_block bigint NOT NULL,
    end_block bigint NOT NULL
);


ALTER TABLE public."aggFt" OWNER TO postgres;

--
-- Name: aggNft; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public."aggNft" (
    tx_hash character varying(255) NOT NULL,
    from_addr character varying(255) NOT NULL,
    to_addr character varying(255) NOT NULL,
    contract character varying(255) NOT NULL,
    value bigint NOT NULL,
    start_ts bigint NOT NULL,
    chain_id bigint NOT NULL,
    end_ts bigint NOT NULL,
    start_block bigint NOT NULL,
    end_block bigint NOT NULL
);


ALTER TABLE public."aggNft" OWNER TO postgres;

--
-- Name: chain_data; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.chain_data (
    chain_id bigint NOT NULL,
    height bigint NOT NULL,
    block_hash character varying(255) NOT NULL,
    ts bigint NOT NULL,
    tx_hash character varying(255) NOT NULL,
    tx_idx integer NOT NULL,
    log_idx integer NOT NULL,
    from_addr character varying(255) NOT NULL,
    to_addr character varying(255) NOT NULL,
    contract character varying(255) NOT NULL,
    value character varying(255) NOT NULL,
    gas character varying(255) NOT NULL,
    gas_price character varying(255) NOT NULL,
    nonce bigint
);


ALTER TABLE public.chain_data OWNER TO postgres;

--
-- Name: chainfromcontract; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX chainfromcontract ON public."aggFt" USING btree (from_addr, contract, chain_id);


--
-- Name: fromtscontractid; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX fromtscontractid ON public.chain_data USING btree (ts, from_addr, contract, chain_id);


--
-- Name: hashtxidxlogidx; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX hashtxidxlogidx ON public.chain_data USING btree (chain_id, tx_hash, tx_idx, log_idx);


--
-- Name: nftuniqe; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX nftuniqe ON public."aggNft" USING btree (from_addr, contract, chain_id);


--
-- Name: totscontractid; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX totscontractid ON public.chain_data USING btree (ts, to_addr, contract, chain_id);


--
-- Name: tscontractid; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX tscontractid ON public.chain_data USING btree (chain_id, ts, contract);


--
-- PostgreSQL database dump complete
--

