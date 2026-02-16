package strings

// Erwartet einen String s und liefert einen neuen String,
// in dem jeder Buchstabe aus s zweimal hintereinander vorkommt.
func DuplicateChars(s string) string {
	// TODO
	result := ""
	for i := range s{
		result+= string(s[i])
		result+= string(s[i])
	}

	return result 
}
