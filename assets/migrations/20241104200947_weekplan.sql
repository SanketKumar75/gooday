-- +goose Up
-- +goose StatementBegin
CREATE TABLE weekplan (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  day VARCHAR(50) NOT NULL,
  workout VARCHAR(50) NOT NULL
);
INSERT INTO weekplan (day, workout) VALUES
('Monday', 'Running'),
('Tuesday', 'Yoga'),
('Wednesday', 'Weightlifting'),
('Thursday', 'Cycling'),
('Friday', 'Swimming'),
('Saturday', 'Pilates'),
('Sunday', 'Rest Day'),
('Monday', 'HIIT'),
('Tuesday', 'Boxing'),
('Wednesday', 'Hiking'),
('Thursday', 'Stretching'),
('Friday', 'CrossFit'),
('Saturday', 'Dance'),
('Sunday', 'Meditation'),
('Monday', 'Rock Climbing'),
('Tuesday', 'Jogging'),
('Wednesday', 'Kickboxing'),
('Thursday', 'Rowing'),
('Friday', 'Circuit Training'),
('Saturday', 'Swimming');
-- +goose StatementEnd


-- +goose Down
-- +goose StatementBegin
DROP TABLE weekplan;
-- +goose StatementEnd
