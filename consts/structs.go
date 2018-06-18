package consts

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

// MySQLConf is the MySQL Configuration
type MySQLConf struct {
	Hostname string
	Port     int
	Username string
	Password string
	Database string
}

// REDISConf used for Redis(later)
type REDISConf struct {
	Host     string
	Port     int16
	Password string
}

// User is a User Object
type User struct {
	ID                    int32
	UserName              string
	UserNameSafe          string
	EMail                 string
	BCryptPassword        string
	Privileges            int32
	Status                Status
	Relax                 bool
	Achievements          int32
	AchievementsDisplayed int32
}

// Status status of current User
type Status struct {
	Country        int16
	Lat            float64
	Lon            float64
	Banned         bool
	BannedUntil    time.Time
	BannedReason   string
	Silenced       bool
	SilencedUntil  time.Time
	SilencedReason string
	Verified       bool
}

// Leaderboard is ofc our Leaderboard! where our "PeppyPoints" is located at.
type Leaderboard struct {
	UserID      uint32
	RankedScore int64
	TotalScore  int64
	Count300    int64
	Count100    int64
	Count50     int64
	CountMiss   int64
	Playcount   int64
	PeppyPoints float64
	Position    int32
}

// GeoIP Is to get where the User is located. E.G ("Germany")
type GeoIP struct {
	City          string `json:"city"`
	Continent     string `json:"continent"`
	ContinentFull string `json:"continent_full"`
	Country       string `json:"country"`
	CountryFull   string `json:"country_full"`
	IP            string `json:"ip"`
	LocRaw        string `json:"loc"`
	Location      struct {
		Lon float64
		Lat float64
	}
}

// CheckPassword if valid.
func (u *User) CheckPassword(passwordMD5 string) bool {
	if bcrypt.CompareHashAndPassword([]byte(u.BCryptPassword), []byte(passwordMD5)) != nil {
		return false
	}
	return true
}
