package internal

type User struct {
	ID            string `gorm:"primaryKey"`
	Givenname     string `form:"givenname"`
	Familyname    string `form:"familyname"`
	Streetaddress string `form:"street-address"`
	Addresslevel1 string `form:"address-level1"`
	Addresslevel2 string `form:"address-level2"`
	Postalcode    string `form:"postal-code"`
	Country       string `form:"country"`
	Countryname   string `form:"country-name"`

	Locations []Location
}

type Location struct {
	ID        string `gorm:"primaryKey"`
	UserID    string
	Timestamp uint `json:"timestamp"`
	Coords    struct {
		Accuracy         float32 `json:"accuracy"`
		Latitude         float32 `json:"latitude"`
		Longitude        float32 `json:"longitude"`
		Altitude         float32 `json:"altitude"`
		AltitudeAccuracy float32 `json:"altitudeAccuracy"`
		Heading          float32 `json:"heading"`
		Speed            float32 `json:"speed"`
	} `json:"coords" gorm:"embedded"`
}

// TODO: do some validation
func (l Location) SeemsLegit() bool {
	return true
}

type Device struct {
}
