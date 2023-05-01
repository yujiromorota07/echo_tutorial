
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE  `todos` (
    `id` int unsigned AUTO_INCREMENT NOT NULL comment 'todoID',
    `title` varchar(255) NOT NULL comment 'タイトル',
    `content` varchar(255) NOT NULL comment 'コンテンツ',
    `created_at` timestamp DEFAULT current_timestamp,
    `updated_at` timestamp DEFAULT current_timestamp on update current_timestamp,
	PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4
comment='Todo用テーブル';


-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE `todos`;
