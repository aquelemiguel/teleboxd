package locales

const (
	TrackSuccess        = "Successfully started tracking @%v!"
	AlreadyTracking     = "I'm already tracking @%v!"
	InvalidTrackUsage   = "Invalid usage of /track command. Usage: /track <username>."
	UntrackSuccess      = "Successfully stopped tracking @%v!"
	NotTracking         = "I'm not tracking @%v!"
	InvalidUntrackUsage = "Invalid usage of /untrack command. Usage: /untrack <username>."
	NewFilmEntry        = "🎬 <b><a href=\"%v\">%v</a></b> just watched <b><a href=\"%v\">%v (%v)</a></b> and rated it %v."
	ListHeader          = "I'm tracking these Letterboxd users:\n"
	ListHeaderEmpty     = "Oops, I'm not tracking any users!\nUse <code>/track &lt;handle&gt;</code> to start tracking someone."
	ListEntry           = "• <b><a href=\"https://letterboxd.com/%s\">%s</a></b>\n"
)
