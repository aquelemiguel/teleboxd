package locales

const (
	TrackSuccess         = "&#9989; Started tracking <b><a href=\"https://letterboxd.com/%s\">%s</a></b>'s diary!"
	TrackInvalidUser     = "&#10060; <b><a href=\"https://letterboxd.com/%s\">%s</a></b> is not a valid Letterboxd user!"
	AlreadyTracking      = "&#10060; I'm already tracking <b><a href=\"https://letterboxd.com/%s\">%s</a></b>'s diary!"
	InvalidTrackUsage    = "&#10060; Invalid usage of /track command! Usage: <code>/track &lt;handle&gt;</code>."
	UntrackSuccess       = "&#9989; Stopped tracking <b><a href=\"https://letterboxd.com/%s\">%s</a></b>'s diary!"
	NotTracking          = "&#10060; I'm not tracking <b><a href=\"https://letterboxd.com/%s\">%s</a></b>'s diary!"
	InvalidUntrackUsage  = "&#10060; Invalid usage of /untrack command! Usage: <code>/untrack &lt;handle&gt;</code>."
	NewFilmEntry         = "&#127871; <b><a href=\"%v\">%v</a></b> just watched <b><a href=\"%v\">%v (%v)</a></b> and rated it %v."
	NewFilmEntryNoRating = "&#127871; <b><a href=\"%v\">%v</a></b> just watched <b><a href=\"%v\">%v (%v)</a></b>."
	ListHeader           = "&#128466;&#65039; I'm tracking these Letterboxd users:\n"
	ListHeaderEmpty      = "&#128466;&#65039; I'm not tracking any users!\nUsage: <code>/track &lt;handle&gt;</code>"
	ListEntry            = "• <b><a href=\"https://letterboxd.com/%s\">%s</a></b>\n"
	SomethingWentWrong   = "&#10060; Oops, something went wrong! Is Letterboxd down?"
)
