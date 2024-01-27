ALTER TABLE public.likes 
DROP COLUMN created_at,
DROP COLUMN updated_at,
DROP COLUMN deleted_at;

DROP INDEX IF EXISTS idx_deleted_at;
DROP INDEX IF EXISTS idx_created_at;