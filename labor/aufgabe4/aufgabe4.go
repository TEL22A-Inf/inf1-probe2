package aufgabe4

/*
AUFGABENSTELLUNG: Vervollständigen Sie die Funktion Greater.

Zustzbedingungen:
- Die Funktion muss rekursiv sein.
- Verwenden Sie ausschließlich Tests der Form x == 0 oder y != 0.
- Gehen Sie davon aus, das x,y >= 0 gilt.
PUNKTE: 10
BEWERTUNG:
*/

// Greater erwartet zwei int-Parameter x und y und liefert true, falls x < y.
func Greater(x, y int) bool {
	// BEGIN-SOLUTION
	if x == 0 {
		return false
	}
	if y == 0 {
		return x != 0
	}
	return Greater(x-1, y-1)
	// END-SOLUTION
}
