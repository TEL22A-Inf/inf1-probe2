package aufgabe5

/*
AUFGABENSTELLUNG: Vervollständigen Sie die Funktion ArrayMax.
PUNKTE: 5
BEWERTUNG:
*/

// ArrayMax erwartet eine int-Slice und liefert das größte Element aus dieser Liste.
// Ist die liste leer, soll 0 geliefert werden.
func ArrayMax(list []int) int {
	// BEGIN-SOLUTION
	result := 0
	for _, el := range list {
		if el > result {
			result = el
		}
	}
	return result
	// END-SOLUTION
}
