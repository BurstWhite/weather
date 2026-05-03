package service

import (
	"weather-app/backend/internal/database"
	"weather-app/backend/internal/model"
)

func ListCities() ([]model.City, error) {
	rows, err := database.DB.Query("SELECT id, location_id, name, adm1, adm2, lat, lon FROM cities ORDER BY name")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var cities []model.City
	for rows.Next() {
		var c model.City
		if err := rows.Scan(&c.ID, &c.LocationID, &c.Name, &c.Adm1, &c.Adm2, &c.Lat, &c.Lon); err != nil {
			return nil, err
		}
		cities = append(cities, c)
	}
	return cities, nil
}

func GetCityByLocationID(locationID string) (*model.City, error) {
	row := database.DB.QueryRow("SELECT id, location_id, name, adm1, adm2, lat, lon FROM cities WHERE location_id = ?", locationID)
	var c model.City
	err := row.Scan(&c.ID, &c.LocationID, &c.Name, &c.Adm1, &c.Adm2, &c.Lat, &c.Lon)
	if err != nil {
		return nil, err
	}
	return &c, nil
}

func AddCity(name, locationID, adm1, adm2 string, lat, lon float64) (*model.City, error) {
	_, err := database.DB.Exec(
		"INSERT OR IGNORE INTO cities(location_id, name, adm1, adm2, lat, lon) VALUES(?, ?, ?, ?, ?, ?)",
		locationID, name, adm1, adm2, lat, lon,
	)
	if err != nil {
		return nil, err
	}
	return GetCityByLocationID(locationID)
}

func DeleteCity(id int64) error {
	_, err := database.DB.Exec("DELETE FROM cities WHERE id = ?", id)
	return err
}
