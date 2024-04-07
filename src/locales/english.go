package locales

const (
	TrackSuccess         = "&#9989; Started tracking <b><a href=\"https://letterboxd.com/%s\">%s</a></b>'s diary!"
	TrackInvalidUser     = "&#10060; <b><a href=\"https://letterboxd.com/%s\">%s</a></b> is not a valid Letterboxd user!"
	AlreadyTracking      = "&#10060; I'm already tracking <b><a href=\"https://letterboxd.com/%s\">%s</a></b>'s diary!"
	InvalidTrackUsage    = "&#10060; Invalid usage of /track command! Usage: <code>/track &lt;handle&gt;</code>."
	UntrackSuccess       = "&#9989; Stopped tracking <b><a href=\"https://letterboxd.com/%s\">%s</a></b>'s diary!"
	NotTracking          = "&#10060; I'm not tracking <b><a href=\"https://letterboxd.com/%s\">%s</a></b>'s diary!"
	InvalidUntrackUsage  = "&#10060; Invalid usage of /untrack command! Usage: <code>/untrack &lt;handle&gt;</code>."
	NewFilmWatch         = "&#127871; <b><a href=\"%v\">%v</a></b> watched <b><a href=\"%v\">%v (%v)</a></b>."
	NewFilmWatchRating   = "&#127871; <b><a href=\"%v\">%v</a></b> watched and rated <b><a href=\"%v\">%v (%v)</a></b> %v."
	NewFilmRewatch       = "&#127871; <b><a href=\"%v\">%v</a></b> rewatched <b><a href=\"%v\">%v (%v)</a></b>."
	NewFilmRewatchRating = "&#127871; <b><a href=\"%v\">%v</a></b> rewatched and rated <b><a href=\"%v\">%v (%v)</a></b> %v."
	ListHeader           = "&#128466;&#65039; I'm tracking these Letterboxd users:\n"
	ListHeaderEmpty      = "&#128466;&#65039; I'm not tracking any users!\nUsage: <code>/track &lt;handle&gt;</code>"
	ListEntry            = "• <b><a href=\"https://letterboxd.com/%s\">%s</a></b>\n"
	SomethingWentWrong   = "&#10060; Oops, something went wrong! Is Letterboxd down?"
)
