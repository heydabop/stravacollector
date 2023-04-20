-- Code generated by 'make coderd/database/generate'. DO NOT EDIT.

CREATE TABLE athlete_efforts (
    id uuid NOT NULL,
    athlete_id integer NOT NULL,
    segment_id integer NOT NULL,
    last_checked timestamp without time zone NOT NULL,
    best_effort_id integer NOT NULL,
    kom_rank integer NOT NULL
);

CREATE TABLE athletes (
    id bigint NOT NULL,
    premium boolean NOT NULL,
    username text NOT NULL,
    firstname text NOT NULL,
    lastname text NOT NULL,
    sex text NOT NULL,
    provider_id text NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    oauth_access_token text NOT NULL,
    oauth_refresh_token text NOT NULL,
    oauth_expiry timestamp with time zone NOT NULL,
    raw text NOT NULL
);

COMMENT ON COLUMN athletes.provider_id IS 'Oauth app client ID';

CREATE TABLE segments (
    id integer NOT NULL,
    name text NOT NULL
);

CREATE TABLE webhook_dump (
    id uuid NOT NULL,
    recorded_at timestamp without time zone NOT NULL,
    raw text NOT NULL
);

ALTER TABLE ONLY athlete_efforts
    ADD CONSTRAINT athlete_efforts_pkey PRIMARY KEY (id);

ALTER TABLE ONLY athletes
    ADD CONSTRAINT athletes_pkey PRIMARY KEY (id);

ALTER TABLE ONLY segments
    ADD CONSTRAINT segments_pkey PRIMARY KEY (id);

ALTER TABLE ONLY webhook_dump
    ADD CONSTRAINT webhook_dump_pkey PRIMARY KEY (id);

ALTER TABLE ONLY athlete_efforts
    ADD CONSTRAINT athlete_efforts_athlete_id_fkey FOREIGN KEY (athlete_id) REFERENCES athletes(id);

ALTER TABLE ONLY athlete_efforts
    ADD CONSTRAINT athlete_efforts_segment_id_fkey FOREIGN KEY (segment_id) REFERENCES segments(id);

