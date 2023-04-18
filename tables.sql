CREATE TABLE public.call_history
(
    id text NOT NULL,
    duration numeric(11,6) NOT NULL,
    code integer NOT NULL,
    url_request text NOT NULL,
    created_at timestamp
    without time zone NOT NULL,
    CONSTRAINT call_history_pkey PRIMARY KEY
    (id)
);


    CREATE TABLE public.currency
    (
        id text NOT NULL,
        code character varying(5) NOT NULL,
        value double precision NOT NULL,
        updated_at timestamp
        without time zone NOT NULL,
    CONSTRAINT currency_pkey PRIMARY KEY
        (id)
);

