package aufgabe7

/*
AUFGABENSTELLUNG: Vervollständigen Sie die Funktion Difference.
PUNKTE: 10
BEWERTUNG:
*/

// Difference erwartet zwei int-Listen l1 und l2.
// Die Funktion liefert eine int-Liste mit den Elementen aus l1,
// die nicht in l2 vorhanden sind.
func Difference(l1, l2 []int) []int {
	// BEGIN-SOLUTION
	result := []int{}
	for _, el1 := range l1 {
		found := false
		for _, el2 := range l2 {
			if el1 == el2 {
				found = true
			}
		}
		if !found {
			result = append(result, el1)
		}
	}
	return result
	// END-SOLUTION
}
