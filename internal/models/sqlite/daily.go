package sqlite

import (
	"database/sql"

	"github.com/sanketkumar75/goodDay/internal/models"
)

type DailyModel struct {
	DB *sql.DB
}

func (m *DailyModel) All() ([]models.Daily, error) {
	stmt := `SELECT id, workout, date from daily order by id`

	rows, err := m.DB.Query(stmt)

	if err != nil {
		return nil, err
	}

	daily := []models.Daily{}

	for rows.Next() {
		p := models.Daily{}
		err := rows.Scan(&p.Id, &p.Workout, &p.Date)
		if err != nil {
			return nil, err
		}
		daily = append(daily, p)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return daily, nil
}
