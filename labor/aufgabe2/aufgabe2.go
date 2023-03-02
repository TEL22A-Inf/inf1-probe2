package aufgabe2

/*
AUFGABENSTELLUNG: Vervollständigen Sie die Funktion GetStringsBetween.
PUNKTE: 15
BEWERTUNG:
*/

// GetStringsBetween eine Liste und zwei Strings first und last.
// Die Funktion liefert die Teilliste, die zwischen first und last liegt.
// first sollen nicht zum Ergebnis gehören.
// Falls die Liste first oder last nicht enthält, oder falls last vor first vorkommt,
// soll die leere Liste geliefert werden.
func GetStringsBetween(list []string, first, last string) []string {
	result := []string{}
	firstfound := false
	// BEGIN-SOLUTION
	for _, s := range list {
		if s == first {
			firstfound = true
		} else if firstfound && s != last {
			result = append(result, s)
		}
		if s == last {
			return result
		}
	}
	// END-SOLUTION
	return result
}
