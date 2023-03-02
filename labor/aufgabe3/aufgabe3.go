package aufgabe3

/*
AUFGABENSTELLUNG: Vervollständigen Sie die Funktion CountEven.
Die Funktion muss rekursiv sein.
PUNKTE: 10
BEWERTUNG:
*/

// CountEven erwartet eine Liste von Zahlen und liefert die Anzahl der geraden Zahlen darin.
func CountEven(list []int) int {
	// BEGIN-SOLUTION
	if len(list) == 0 {
		return 0
	}
	head, result := list[0], CountEven(list[1:])
	if head%2 == 0 {
		result++
	}
	return result
	// END-SOLUTION
}
