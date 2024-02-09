DROP INDEX IF EXISTS idx_unique_user_post_uuid;
ALTER TABLE public.likes DROP CONSTRAINT unique_user_and_post_uuid;
CREATE INDEX IF NOT EXISTS idx_user_post_uuid ON likes (user_uuid,post_uuid);
