CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS public.likes (
    uuid UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_uuid UUID NOT NULL,
    post_uuid UUID NOT NULL
);
CREATE INDEX IF NOT EXISTS idx_user_post_uuid ON likes (user_uuid,post_uuid);

ALTER TABLE public.likes OWNER TO "user";
GRANT ALL ON TABLE public.likes TO "user";