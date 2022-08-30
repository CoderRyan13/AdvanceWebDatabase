-- Filename: 000002_add_column_type_to_messages_table.up.sql

ALTER TABLE school
ADD COLUMN type integer NOT NULL DEFAULT 1;