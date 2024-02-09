DROP INDEX IF EXISTS idx_user_post_uuid;
ALTER TABLE public.likes ADD CONSTRAINT unique_user_and_post_uuid UNIQUE (user_uuid,post_uuid);
CREATE UNIQUE INDEX IF NOT EXISTS idx_unique_user_post_uuid ON public.likes (user_uuid,post_uuid);