package model

type SensorData struct {
	ID    int     `json:"id" gorm:"primary_key"`
	water float64 `json:"water"`
	wind  float64 `json:"wind"`
}
