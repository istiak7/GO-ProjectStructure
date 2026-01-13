-- +migrate Up
ALTER TABLE suppliers
ADD COLUMN address TEXT;
