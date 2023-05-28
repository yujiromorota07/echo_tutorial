
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE `todo_status` (
    `code` int unsigned NOT NULL comment 'ステータスコード',
    `status` varchar(255) NOT NULL comment 'ステータス',
    `created_at` timestamp  default current_timestamp,
    `updated_at` timestamp default current_timestamp on update current_timestamp,
    PRIMARY KEY (`code`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4
comment='Todoステータス';

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE `todo_status`;