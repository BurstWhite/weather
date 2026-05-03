package database

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite"
)

var DB *sql.DB

func Init(dbPath string) error {
	var err error
	DB, err = sql.Open("sqlite", dbPath)
	if err != nil {
		return err
	}
	DB.SetMaxOpenConns(1)
	if err = migrate(); err != nil {
		return err
	}
	return nil
}

func migrate() error {
	_, err := DB.Exec(`
		CREATE TABLE IF NOT EXISTS cities (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			location_id TEXT UNIQUE NOT NULL,
			name TEXT NOT NULL,
			adm1 TEXT,
			adm2 TEXT,
			lat REAL,
			lon REAL
		);
		CREATE TABLE IF NOT EXISTS weather_cache (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			city_id INTEGER REFERENCES cities(id),
			fx_date TEXT NOT NULL,
			temp_max INTEGER,
			temp_min INTEGER,
			text_day TEXT,
			text_night TEXT,
			icon_day TEXT,
			icon_night TEXT,
			humidity INTEGER,
			wind_dir_day TEXT,
			wind_scale_day TEXT,
			precip REAL,
			uv_index INTEGER,
			pressure REAL,
			sunrise TEXT,
			sunset TEXT,
			updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			UNIQUE(city_id, fx_date)
		);
		CREATE TABLE IF NOT EXISTS air_cache (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			city_id INTEGER REFERENCES cities(id),
			aqi INTEGER,
			category TEXT,
			level TEXT,
			primary_pollutant TEXT,
			pm2p5 REAL,
			pm10 REAL,
			no2 REAL,
			o3 REAL,
			co REAL,
			so2 REAL,
			health_effect TEXT,
			health_advice TEXT,
			updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		);
	`)
	if err != nil {
		return err
	}
	if err = seed(); err != nil {
		log.Printf("seed warning: %v", err)
	}
	return nil
}

func Close() {
	if DB != nil {
		DB.Close()
	}
}
