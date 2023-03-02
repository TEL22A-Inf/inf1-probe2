package aufgabe6

/*
AUFGABENSTELLUNG:
Hier ist ein Struct Person vorgegeben.
Vervollständigen Sie die Funktion CommonFriends.
PUNKTE: 15
BEWERTUNG:
*/

// Person repräsentiert eine Person in einem sozialen Netzwerk.
type Person struct {
	Name    string
	Friends []Person
}

// CommonFriends erwartet zwei Personen und liefert eine Liste
// mit allen Personen, die mit beiden befreundet sind.
func CommonFriends(p1, p2 Person) []Person {
	// BEGIN-SOLUTION
	result := []Person{}
	for _, f1 := range p1.Friends {
		for _, f2 := range p2.Friends {
			if f1.Name == f2.Name {
				result = append(result, f1)
			}
		}
	}
	return result
	// END-SOLUTION
}
