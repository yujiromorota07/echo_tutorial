
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
ALTER TABLE `todos` ADD COLUMN `status_code` int unsigned NOT NULL AFTER `content`;
ALTER TABLE `todos` ADD FOREIGN KEY (`status_code`) REFERENCES `todo_status` (`code`);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
ALTER TABLE `todos` DROP FOREIGN KEY `todos_ibfk_1`;
ALTER TABLE `todos` DROP COLUMN `status_code`;
