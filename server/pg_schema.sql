--
-- PostgreSQL database dump
--

\restrict ltTjVhsYi6LM05tZYQHcvr9wim62pCyZYGQicM9E9iRwj7036FGYQpFjnYM4yaV

-- Dumped from database version 18.1 (Debian 18.1-1.pgdg13+2)
-- Dumped by pg_dump version 18.1

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET transaction_timeout = 0;
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
-- Name: channels; Type: TABLE; Schema: public; Owner: root
--

CREATE TABLE public.channels (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    name text,
    group_channel_id bigint
);


ALTER TABLE public.channels OWNER TO root;

--
-- Name: channels_id_seq; Type: SEQUENCE; Schema: public; Owner: root
--

CREATE SEQUENCE public.channels_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.channels_id_seq OWNER TO root;

--
-- Name: channels_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: root
--

ALTER SEQUENCE public.channels_id_seq OWNED BY public.channels.id;


--
-- Name: group_channels; Type: TABLE; Schema: public; Owner: root
--

CREATE TABLE public.group_channels (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    name text,
    guild_id bigint
);


ALTER TABLE public.group_channels OWNER TO root;

--
-- Name: group_channels_id_seq; Type: SEQUENCE; Schema: public; Owner: root
--

CREATE SEQUENCE public.group_channels_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.group_channels_id_seq OWNER TO root;

--
-- Name: group_channels_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: root
--

ALTER SEQUENCE public.group_channels_id_seq OWNED BY public.group_channels.id;


--
-- Name: guild_members; Type: TABLE; Schema: public; Owner: root
--

CREATE TABLE public.guild_members (
    guild_id bigint NOT NULL,
    user_id bigint NOT NULL,
    nickname text,
    joined_at timestamp with time zone
);


ALTER TABLE public.guild_members OWNER TO root;

--
-- Name: guilds; Type: TABLE; Schema: public; Owner: root
--

CREATE TABLE public.guilds (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    name text,
    profile text
);


ALTER TABLE public.guilds OWNER TO root;

--
-- Name: guilds_id_seq; Type: SEQUENCE; Schema: public; Owner: root
--

CREATE SEQUENCE public.guilds_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.guilds_id_seq OWNER TO root;

--
-- Name: guilds_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: root
--

ALTER SEQUENCE public.guilds_id_seq OWNED BY public.guilds.id;


--
-- Name: member_roles; Type: TABLE; Schema: public; Owner: root
--

CREATE TABLE public.member_roles (
    guild_member_guild_id bigint NOT NULL,
    guild_member_user_id bigint NOT NULL,
    role_id bigint NOT NULL
);


ALTER TABLE public.member_roles OWNER TO root;

--
-- Name: messages; Type: TABLE; Schema: public; Owner: root
--

CREATE TABLE public.messages (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    channel_id bigint,
    user_id bigint,
    content text
);


ALTER TABLE public.messages OWNER TO root;

--
-- Name: messages_id_seq; Type: SEQUENCE; Schema: public; Owner: root
--

CREATE SEQUENCE public.messages_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.messages_id_seq OWNER TO root;

--
-- Name: messages_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: root
--

ALTER SEQUENCE public.messages_id_seq OWNED BY public.messages.id;


--
-- Name: roles; Type: TABLE; Schema: public; Owner: root
--

CREATE TABLE public.roles (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    name text,
    guild_id bigint
);


ALTER TABLE public.roles OWNER TO root;

--
-- Name: roles_id_seq; Type: SEQUENCE; Schema: public; Owner: root
--

CREATE SEQUENCE public.roles_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.roles_id_seq OWNER TO root;

--
-- Name: roles_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: root
--

ALTER SEQUENCE public.roles_id_seq OWNED BY public.roles.id;


--
-- Name: users; Type: TABLE; Schema: public; Owner: root
--

CREATE TABLE public.users (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    username text,
    password_hash text,
    email text,
    status text,
    online_status bigint,
    forgot_token text,
    expire_forgot_token timestamp with time zone,
    profile text
);


ALTER TABLE public.users OWNER TO root;

--
-- Name: users_id_seq; Type: SEQUENCE; Schema: public; Owner: root
--

CREATE SEQUENCE public.users_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.users_id_seq OWNER TO root;

--
-- Name: users_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: root
--

ALTER SEQUENCE public.users_id_seq OWNED BY public.users.id;


--
-- Name: channels id; Type: DEFAULT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.channels ALTER COLUMN id SET DEFAULT nextval('public.channels_id_seq'::regclass);


--
-- Name: group_channels id; Type: DEFAULT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.group_channels ALTER COLUMN id SET DEFAULT nextval('public.group_channels_id_seq'::regclass);


--
-- Name: guilds id; Type: DEFAULT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.guilds ALTER COLUMN id SET DEFAULT nextval('public.guilds_id_seq'::regclass);


--
-- Name: messages id; Type: DEFAULT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.messages ALTER COLUMN id SET DEFAULT nextval('public.messages_id_seq'::regclass);


--
-- Name: roles id; Type: DEFAULT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.roles ALTER COLUMN id SET DEFAULT nextval('public.roles_id_seq'::regclass);


--
-- Name: users id; Type: DEFAULT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.users ALTER COLUMN id SET DEFAULT nextval('public.users_id_seq'::regclass);


--
-- Name: channels channels_pkey; Type: CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.channels
    ADD CONSTRAINT channels_pkey PRIMARY KEY (id);


--
-- Name: group_channels group_channels_pkey; Type: CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.group_channels
    ADD CONSTRAINT group_channels_pkey PRIMARY KEY (id);


--
-- Name: guild_members guild_members_pkey; Type: CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.guild_members
    ADD CONSTRAINT guild_members_pkey PRIMARY KEY (guild_id, user_id);


--
-- Name: guilds guilds_pkey; Type: CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.guilds
    ADD CONSTRAINT guilds_pkey PRIMARY KEY (id);


--
-- Name: member_roles member_roles_pkey; Type: CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.member_roles
    ADD CONSTRAINT member_roles_pkey PRIMARY KEY (guild_member_guild_id, guild_member_user_id, role_id);


--
-- Name: messages messages_pkey; Type: CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.messages
    ADD CONSTRAINT messages_pkey PRIMARY KEY (id);


--
-- Name: roles roles_pkey; Type: CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.roles
    ADD CONSTRAINT roles_pkey PRIMARY KEY (id);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- Name: idx_channels_deleted_at; Type: INDEX; Schema: public; Owner: root
--

CREATE INDEX idx_channels_deleted_at ON public.channels USING btree (deleted_at);


--
-- Name: idx_group_channels_deleted_at; Type: INDEX; Schema: public; Owner: root
--

CREATE INDEX idx_group_channels_deleted_at ON public.group_channels USING btree (deleted_at);


--
-- Name: idx_guilds_deleted_at; Type: INDEX; Schema: public; Owner: root
--

CREATE INDEX idx_guilds_deleted_at ON public.guilds USING btree (deleted_at);


--
-- Name: idx_messages_deleted_at; Type: INDEX; Schema: public; Owner: root
--

CREATE INDEX idx_messages_deleted_at ON public.messages USING btree (deleted_at);


--
-- Name: idx_roles_deleted_at; Type: INDEX; Schema: public; Owner: root
--

CREATE INDEX idx_roles_deleted_at ON public.roles USING btree (deleted_at);


--
-- Name: idx_users_deleted_at; Type: INDEX; Schema: public; Owner: root
--

CREATE INDEX idx_users_deleted_at ON public.users USING btree (deleted_at);


--
-- Name: messages fk_channels_messages; Type: FK CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.messages
    ADD CONSTRAINT fk_channels_messages FOREIGN KEY (channel_id) REFERENCES public.channels(id);


--
-- Name: channels fk_group_channels_channels; Type: FK CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.channels
    ADD CONSTRAINT fk_group_channels_channels FOREIGN KEY (group_channel_id) REFERENCES public.group_channels(id);


--
-- Name: group_channels fk_guilds_group_channels; Type: FK CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.group_channels
    ADD CONSTRAINT fk_guilds_group_channels FOREIGN KEY (guild_id) REFERENCES public.guilds(id);


--
-- Name: guild_members fk_guilds_members; Type: FK CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.guild_members
    ADD CONSTRAINT fk_guilds_members FOREIGN KEY (guild_id) REFERENCES public.guilds(id);


--
-- Name: roles fk_guilds_role; Type: FK CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.roles
    ADD CONSTRAINT fk_guilds_role FOREIGN KEY (guild_id) REFERENCES public.guilds(id);


--
-- Name: member_roles fk_member_roles_guild_member; Type: FK CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.member_roles
    ADD CONSTRAINT fk_member_roles_guild_member FOREIGN KEY (guild_member_guild_id, guild_member_user_id) REFERENCES public.guild_members(guild_id, user_id);


--
-- Name: member_roles fk_member_roles_role; Type: FK CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.member_roles
    ADD CONSTRAINT fk_member_roles_role FOREIGN KEY (role_id) REFERENCES public.roles(id);


--
-- Name: guild_members fk_users_memberships; Type: FK CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.guild_members
    ADD CONSTRAINT fk_users_memberships FOREIGN KEY (user_id) REFERENCES public.users(id);


--
-- Name: messages fk_users_messages; Type: FK CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.messages
    ADD CONSTRAINT fk_users_messages FOREIGN KEY (user_id) REFERENCES public.users(id);


--
-- PostgreSQL database dump complete
--

\unrestrict ltTjVhsYi6LM05tZYQHcvr9wim62pCyZYGQicM9E9iRwj7036FGYQpFjnYM4yaV

