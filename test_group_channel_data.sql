-- Test data for group_channels
INSERT INTO public.group_channels (id, created_at, updated_at, deleted_at, name, guild_id) VALUES
(1, NOW(), NOW(), NULL, 'General', 1),
(2, NOW(), NOW(), NULL, 'Voice', 1),
(3, NOW(), NOW(), NULL, 'Announcements', 2),
(4, NOW(), NOW(), NULL, 'Test Channel', 3);