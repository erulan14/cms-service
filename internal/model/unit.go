package model

type Unit struct {
	UUID                string   `json:"uuid"`
	Name                string   `json:"name"`
	UID                 string   `json:"uid"`
	HW                  string   `json:"hw"`
	Phone               string   `json:"phone"`
	Phone2              string   `json:"phone2"`
	Tariff              string   `json:"tariff"`
	Icon                string   `json:"icon"`
	Creator             []string `json:"creator"`
	Sensors             []Sensor `json:"sensors"`
	ActiveNotifyGeozone bool     `json:"active_notify_geozone"`
	ActiveNotifySos     bool     `json:"active_notify_sos"`
	ActiveNotifyTrip    bool     `json:"active_notify_trip"`
	MinSpeed            int      `json:"min_speed"`
	MinTimeParking      int      `json:"min_time_parking"`
	ActiveGpsCorrection bool     `json:"active_gps_correction"`
	MinSat              int      `json:"min_sat"`
	MaxDisMessage       int      `json:"max_dis_message"`
	MinTimeMove         int      `json:"min_time_move"`
	MinDisTrip          int      `json:"min_dis_trip"`
	FuelAlg             bool     `json:"fuel_alg"`
	FuelConsNormal      bool     `json:"fuel_cons_normal"`
	FuelConsGaz         float32  `json:"fuel_cons_gaz"`
	FuelConsBenz        float32  `json:"fuel_cons_benz"`
	CreatedAt           string   `json:"createdAt"`
	UpdatedAt           string   `json:"updatedAt"`
}

type CreateUnitDTO struct {
	Name   string `json:"name"`
	UID    string `json:"uid"`
	HW     string `json:"hw"`
	Phone  string `json:"phone"`
	Phone2 string `json:"phone2"`
	Tariff string `json:"tariff"`
	Icon   string `json:"icon"`
	//Creator             []string `json:"creator"`
	Sensors             []Sensor `json:"sensors"`
	ActiveNotifyGeozone bool     `json:"active_notify_geozone"`
	ActiveNotifySos     bool     `json:"active_notify_sos"`
	ActiveNotifyTrip    bool     `json:"active_notify_trip"`
	MinSpeed            int      `json:"min_speed"`
	MinTimeParking      int      `json:"min_time_parking"`
	ActiveGpsCorrection bool     `json:"active_gps_correction"`
	MinSat              int      `json:"min_sat"`
	MaxDisMessage       int      `json:"max_dis_message"`
	MinTimeMove         int      `json:"min_time_move"`
	MinDisTrip          int      `json:"min_dis_trip"`
	FuelAlg             bool     `json:"fuel_alg"`
	FuelConsNormal      bool     `json:"fuel_cons_normal"`
	FuelConsGaz         float32  `json:"fuel_cons_gaz"`
	FuelConsBenz        float32  `json:"fuel_cons_benz"`
	CreatedAt           string   `json:"createdAt"`
	UpdatedAt           string   `json:"updatedAt"`
}

type UpdateUnitDTO struct {
	Name                string   `json:"name"`
	UID                 string   `json:"uid"`
	HW                  string   `json:"hw"`
	Phone               string   `json:"phone"`
	Phone2              string   `json:"phone2"`
	Tariff              string   `json:"tariff"`
	Icon                string   `json:"icon"`
	Creator             []string `json:"creator"`
	Sensors             []Sensor `json:"sensors"`
	ActiveNotifyGeozone bool     `json:"active_notify_geozone"`
	ActiveNotifySos     bool     `json:"active_notify_sos"`
	ActiveNotifyTrip    bool     `json:"active_notify_trip"`
	MinSpeed            int      `json:"min_speed"`
	MinTimeParking      int      `json:"min_time_parking"`
	ActiveGpsCorrection bool     `json:"active_gps_correction"`
	MinSat              int      `json:"min_sat"`
	MaxDisMessage       int      `json:"max_dis_message"`
	MinTimeMove         int      `json:"min_time_move"`
	MinDisTrip          int      `json:"min_dis_trip"`
	FuelAlg             bool     `json:"fuel_alg"`
	FuelConsNormal      bool     `json:"fuel_cons_normal"`
	FuelConsGaz         float32  `json:"fuel_cons_gaz"`
	FuelConsBenz        float32  `json:"fuel_cons_benz"`
	CreatedAt           string   `json:"createdAt"`
	UpdatedAt           string   `json:"updatedAt"`
}
