CREATE TABLE
  public.users (
    id serial NOT NULL,
    created_at timestamp without time zone NOT NULL DEFAULT now(),
    email character varying(255) NOT NULL,
    first_name character varying(255) NOT NULL,
    last_name character varying(255) NOT NULL,
    password character varying(60) NOT NULL,
    updated_at timestamp without time zone NOT NULL DEFAULT now(),
    user_active integer NOT NULL DEFAULT 0
  );

ALTER TABLE
  public.users
ADD
  CONSTRAINT users_pkey PRIMARY KEY (id)

CREATE TABLE
  public.tokens (
    id serial NOT NULL,
    user_id integer NULL,
    email character varying(255) NOT NULL,
    token character varying(255) NOT NULL,
    created_at timestamp without time zone NOT NULL DEFAULT now(),
    token_hash bytea NOT NULL,
    updated_at timestamp without time zone NOT NULL,
    expiry timestamp
    with
      time zone NOT NULL
  );

ALTER TABLE
  public.tokens
ADD
  CONSTRAINT tokens_pkey PRIMARY KEY (id)