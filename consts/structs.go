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
	ID                    uint32
	UserName              string
	UserNameSafe          string
	EMail                 string
	BCryptPassword        string
	Privileges            uint64
	Status                *Status
	Relax                 bool
	Achievements          uint32
	AchievementsDisplayed uint32
}

// Status status of current User
type Status struct {
	Country        byte
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
	RankedScore uint64
	TotalScore  uint64
	Count300    uint64
	Count100    uint64
	Count50     uint64
	CountMiss   uint64
	Playcount   uint64
	PeppyPoints float64
	Position    uint32
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
