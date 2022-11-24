CREATE TABLE public.accounts
(
    id         SERIAL PRIMARY KEY,
    username   character varying                      NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL
);