package locales

const (
	TrackSuccess        = "Started tracking <b><a href=\"https://letterboxd.com/%s\">@%s</a></b>'s diary!"
	TrackInvalidUser    = "<b><a href=\"https://letterboxd.com/%s\">@%s</a></b> is not a valid Letterboxd user!"
	AlreadyTracking     = "I'm already tracking <b><a href=\"https://letterboxd.com/%s\">@%s</a></b>'s diary!"
	InvalidTrackUsage   = "Invalid usage of /track command!\nUsage: <code>/track &lt;handle&gt;</code>"
	UntrackSuccess      = "Stopped tracking <b><a href=\"https://letterboxd.com/%s\">@%s</a></b>'s diary!"
	NotTracking         = "I'm not tracking <b><a href=\"https://letterboxd.com/%s\">@%s</a></b>'s diary!"
	InvalidUntrackUsage = "Invalid usage of /untrack command!\nUsage: <code>/untrack &lt;handle&gt;</code>"
	NewFilmEntry        = "🎬 <b><a href=\"%v\">%v</a></b> just watched <b><a href=\"%v\">%v (%v)</a></b> and rated it %v."
	ListHeader          = "I'm tracking these Letterboxd users:\n"
	ListHeaderEmpty     = "I'm not tracking any users!\nUsage: <code>/track &lt;handle&gt;</code>"
	ListEntry           = "• <b><a href=\"https://letterboxd.com/%s\">%s</a></b>\n"
)
