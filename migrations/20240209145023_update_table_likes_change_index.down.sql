DROP INDEX idx_unique_user_post_uuid;

CREATE INDEX IF NOT EXISTS idx_user_post_uuid ON likes (user_uuid,post_uuid);
