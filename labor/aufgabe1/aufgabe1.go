package aufgabe1

/*
AUFGABENSTELLUNG: Vervollständigen Sie die Funktion LongestAbc.
PUNKTE: 10
BEWERTUNG:
*/

// LongestAbc erwartet eine Liste von Strings und liefert
// das längste Element, das mit der Buchstabenfolge "abc" beginnt.
// Liefert den leeren String, falls es kein solches Element gibt.
func LongestAbc(list []string) string {
	longestLen := -1
	longestPos := -1
	// BEGIN-SOLUTION
	for pos, val := range list {
		currentLen := len(val)
		if currentLen >= 3 && val[:3] == "abc" {
			if currentLen > longestLen {
				longestLen = currentLen
				longestPos = pos
			}
		}
	}
	if longestPos != -1 {
		return list[longestPos]
	}
	// END-SOLUTION
	return ""
}
