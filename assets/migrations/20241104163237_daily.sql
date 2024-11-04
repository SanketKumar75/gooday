-- +goose Up
-- +goose StatementBegin
CREATE TABLE daily (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  workout VARCHAR(50) NOT NULL,
  date DATETIME NOT NULL
);
INSERT INTO daily (workout, date) VALUES
('Running', '2024-11-01 07:00:00'),
('Yoga', '2024-11-02 08:30:00'),
('Weightlifting', '2024-11-03 18:00:00'),
('Cycling', '2024-11-04 07:30:00'),
('Swimming', '2024-11-05 09:00:00'),
('Pilates', '2024-11-06 17:00:00'),
('HIIT', '2024-11-07 06:30:00'),
('Boxing', '2024-11-08 19:00:00'),
('Hiking', '2024-11-09 10:00:00'),
('Stretching', '2024-11-10 08:00:00');
-- +goose StatementEnd


-- +goose Down
-- +goose StatementBegin
DROP TABLE daily;
-- +goose StatementEnd
