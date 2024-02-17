package irdata

type CarClass struct {
	CarClassID    int          `json:"car_class_id"`
	CarsInClass   []CarInClass `json:"cars_in_class"`
	CustID        int          `json:"cust_id"`
	Name          string       `json:"name"`
	ShortName     string       `json:"short_name"`
	RelativeSpeed int          `json:"relative_speed"`
}

type CarInClass struct {
	CarDirPath string `json:"car_dirpath"`
	CarID      int    `json:"car_id"`
	Retired    bool   `json:"retired"`
}
