package usertools

import (
	"strings"
	"time"

	"github.com/Gigamons/common/consts"
	"github.com/Gigamons/common/helpers"
	"github.com/Gigamons/common/logger"
)

// GetUserID Returns a UserID of the given UserName
func GetUserID(username string) *uint32 {
	var UserID uint32
	db := helpers.DB
	safe := strings.ToLower(strings.Replace(username, " ", "_", -1))
	err := db.Ping()
	if err != nil {
		logger.Errorln(err)
		return nil
	}
	rows, err := db.Query("SELECT id FROM users WHERE username_safe = ?", safe)
	defer rows.Close()
	if err != nil {
		logger.Errorln(err)
		return nil
	}
	for rows.Next() {
		err := rows.Scan(&UserID)
		if err != nil {
			logger.Errorln(err)
			return nil
		}
	}
	return &UserID
}

// GetUser returns a User based on the Defined UserID
func GetUser(userid *uint32) *consts.User {
	// To prevent some errors.
	if userid == nil {
		return nil
	}
	db := helpers.DB
	u := consts.User{}

	rows, err := db.Query("SELECT id, username, username_safe, email, password, privileges, achievements, achievements_displayed FROM users WHERE id = ?", userid)
	defer rows.Close()
	if err != nil {
		logger.Errorln(err)
	}
	for rows.Next() {
		err := rows.Scan(&u.ID, &u.UserName, &u.UserNameSafe, &u.EMail, &u.BCryptPassword, &u.Privileges, &u.Achievements, &u.AchievementsDisplayed)
		if err != nil {
			logger.Errorln(err)
		}
	}

	u.Status = getUserStatus(userid)

	return &u
}

func getUserStatus(userid *uint32) *consts.Status {
	// To prevent some errors.
	if userid == nil {
		return nil
	}
	var banneduntil string
	var silenceduntil string
	var u consts.Status

	db := helpers.DB

	row := db.QueryRow("SELECT country, lat, lon, banned, banned_until, banned_reason, silenced, silenced_until, silenced_reason, verified FROM users_status WHERE id = ?", userid)
	if row == nil {
		return nil
	}

	err := row.Scan(&u.Country, &u.Lat, &u.Lon, &u.Banned, &banneduntil, &u.BannedReason, &u.Silenced, &silenceduntil, &u.SilencedReason, &u.Verified)
	if err != nil {
		logger.Errorln(err)
		return nil
	}

	if banneduntil == "0000-00-00 00:00:00" || banneduntil == "" {
		banneduntil = "0001-01-01 00:00:00"
	}
	if silenceduntil == "0000-00-00 00:00:00" || silenceduntil == "" {
		silenceduntil = "0001-01-01 00:00:00"
	}

	buntil, err := time.Parse("2006-01-02 15:04:05", banneduntil)
	if err != nil {
		logger.Errorln(err)
		buntil = time.Now()
	}
	u.BannedUntil = buntil

	suntil, err := time.Parse("2006-01-02 15:04:05", silenceduntil)
	if err != nil {
		logger.Errorln(err)
		suntil = time.Now()
	}
	u.SilencedUntil = suntil

	return &u
}

// GetLeaderboard is getting the Leaderboard of the Database!
func GetLeaderboard(u *consts.User, playMode byte) *consts.Leaderboard {
	if u == nil {
		return nil
	}
	var lb consts.Leaderboard
	var m string
	if u.Relax {
		m = "_rx"
	} else {
		m = ""
	}

	db := helpers.DB
	pm := consts.ToPlaymodeString(playMode)
	if pm != "" {
		pm = "_" + pm
	}

	row := db.QueryRow("SELECT rankedscore"+pm+", totalscore"+pm+", count_300"+pm+", count_100"+pm+", count_50"+pm+", count_miss"+pm+", playcount"+pm+", pp"+pm+" FROM leaderboard"+m+" WHERE id = ?", u.ID)
	if row == nil {
		return nil
	}
	err := row.Scan(&lb.RankedScore, &lb.TotalScore, &lb.Count300, &lb.Count100, &lb.Count50, &lb.CountMiss, &lb.Playcount, &lb.PeppyPoints)
	if err != nil {
		logger.Errorln(err)
		return nil
	}

	lb.Position = GetLeaderboardPosition(u, playMode)
	lb.UserID = u.ID

	return &lb
}

// GetLeaderboardPosition is used to get the leaderboard position.
func GetLeaderboardPosition(u *consts.User, playMode byte) uint32 {
	if u == nil {
		return 0
	}

	var Pos uint32
	var m string
	if u.Relax {
		m = "_rx"
	} else {
		m = ""
	}

	db := helpers.DB
	pm := consts.ToPlaymodeString(playMode)
	if pm != "" {
		pm = "_" + pm
	}

	rows, err := db.Query("SELECT id FROM leaderboard" + m + " ORDER BY pp" + pm + " DESC, rankedscore" + pm + " DESC")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var userid uint32
		err := rows.Scan(&userid)
		if err != nil {
			logger.Errorln(err)
		}
		if userid == u.ID {
			return Pos + 1
		}
		Pos++
	}

	return 0
}

// GetFriends return a array of friend userids!
func GetFriends(u *consts.User) []uint32 {
	var userids []uint32
	db := helpers.DB

	rows, err := db.Query("SELECT friendid FROM friends WHERE userid = ?", u.ID)
	if err != nil {
		logger.Debugln(err)
		return nil
	}

	defer rows.Close()
	for rows.Next() {
		var useridstr uint32
		err := rows.Scan(&useridstr)
		if err != nil {
			logger.Errorln(err)
		}
		userids = append(userids, useridstr)
	}

	return userids
}
