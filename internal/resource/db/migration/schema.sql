--
-- Name: config_tournament; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.config_tournament (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
);
--
-- Name: difficult; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.difficult (
    id bigint NOT NULL,
    code varchar(3) NOT NULL,
    name varchar(50) NOT NULL,
    value_add int NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
);
--
-- Name: difficult_positions; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.difficult_positions (
    id bigint NOT NULL,
    difficult_id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
);
--
-- Name: positions; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.positions (
    id bigint NOT NULL,
    name varchar(50) NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
);
--
-- Name: tournament; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.tournaments (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    city text,
    microzone_name text,
    microzone_id bigint,
    max_distance_meters bigint
);

ALTER TABLE public.tournaments
    ADD CONSTRAINT tournaments_conf_fk FOREIGN KEY (order_id) REFERENCES public.possible_cross_selling (order_id);