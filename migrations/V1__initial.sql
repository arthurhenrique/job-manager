-- DROP TABLE public.job_execution;
CREATE TYPE job_status AS ENUM ('PROCESSING', 'SUCCESS', 'FAILED', 'CANCELLED');

CREATE TABLE job_execution (
    id bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    object_id bigint NOT NULL,
    sleep integer NULL,
    status job_status NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone NOT NULL,
    UNIQUE(object_id)
);