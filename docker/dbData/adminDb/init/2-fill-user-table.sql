INSERT INTO "public"."users" (email, first_name, last_name, password, active, created_at, updated_at)
VALUES ('mike@yaronia.com','Mike','Ross','$2a$12$1zGLuYDDNvATh4RA4avbKuheAMpb1svexSzrQm7up.bnpwQHs0jNe',1,'2022-03-14 00:00:00','2022-03-14 00:00:00')
ON CONFLICT (email) DO UPDATE
           SET
                first_name = EXCLUDED.first_name,
                last_name = EXCLUDED.last_name,
                password = EXCLUDED.password,
                active = EXCLUDED.active,
                created_at = EXCLUDED.created_at;