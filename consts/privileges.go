package consts

// User Privileges.
const (
	Supporter = 2 << iota
	BAT
	TournamentStaff
	AdminPannelAccess
	AdminManageUsers
	AdminBanUsers
	AdminSilenceUsers
	AdminWipeUsers
	AdminBeatmaps
	AdminDeveloper
	AdminSettings
	AdminReports
	AdminPrivileges
	AdminSendAnnouncements
	AdminChatMod
	AdminKickUsers
)
