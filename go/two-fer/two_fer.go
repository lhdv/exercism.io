package twofer

// ShareWith do it's job.
func ShareWith(s string) string {

	if len(s) > 0 {
		return "One for " + s + ", one for me."
	}

	return "One for you, one for me."
}
