ALTER TABLE public.likes
ADD COLUMN created_at timestamp,
ADD COLUMN updated_at timestamp,
ADD COLUMN deleted_at timestamp;

CREATE INDEX idx_deleted_at ON public.likes (deleted_at) WHERE deleted_at IS NOT NULL;
CREATE INDEX idx_created_at ON public.likes (created_at) WHERE created_at IS NOT NULL;