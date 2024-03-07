package domain

type Driver struct {
	DriverID        string `json:"driverId"`
	PermanentNumber string `json:"permanentNumber"`
	Code            string `json:"code"`
	URL             string `json:"url"`
	GivenName       string `json:"givenName"`
	FamilyName      string `json:"familyName"`
	DateOfBirth     string `json:"dateOfBirth"`
	Nationality     string `json:"nationality"`
}

type Constructor struct {
	ConstructorID string `json:"constructorId"`
	URL           string `json:"url"`
	Name          string `json:"name"`
	Nationality   string `json:"nationality"`
}

type DriverStanding struct {
	Position     string        `json:"position"`
	PositionText string        `json:"positionText"`
	Points       string        `json:"points"`
	Wins         string        `json:"wins"`
	Driver       Driver        `json:"Driver"`
	Constructors []Constructor `json:"Constructors"`
}

type StandingList struct {
	Season          string           `json:"season"`
	Round           string           `json:"round"`
	DriverStandings []DriverStanding `json:"DriverStandings"`
}

type StandingsTable struct {
	Season         string         `json:"season"`
	StandingsLists []StandingList `json:"StandingsLists"`
}
type MRData struct {
	Xmlns          string         `json:"xmlns"`
	Series         string         `json:"series"`
	Url            string         `json:"url"`
	Limit          string         `json:"limit"`
	Offset         string         `json:"offset"`
	Total          string         `json:"total"`
	StandingsTable StandingsTable `json:"StandingsTable"`
}

type F1Drivers struct {
	Data MRData `json:"MRData"`
}
