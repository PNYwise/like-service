DROP INDEX idx_user_post_uuid;

CREATE UNIQUE INDEX IF NOT EXISTS idx_unique_user_post_uuid ON likes (user_uuid,post_uuid);