-- Test data for guilds
INSERT INTO public.guilds (id, created_at, updated_at, deleted_at, name, profile) VALUES
(1, NOW(), NOW(), NULL, 'Test Guild 1', '{"description":"Test guild one"}'),
(2, NOW(), NOW(), NULL, 'Test Guild 2', '{"description":"Test guild two"}'),
(3, NOW(), NOW(), NULL, 'Community Hub', '{"description":"Community hub for testing"}');