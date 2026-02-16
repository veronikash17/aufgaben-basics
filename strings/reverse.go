package strings

// Erwartet einen String s und liefert s rückwärts zurück.
func Reverse(s string) string {
	// TODO
	result := ""
	
	for i:=len(s)-1; i>=0 ; i--{
		result += string(s[i])
	}
	return result
}

// Erwartet zwei Strings s1 und s2 und prüft, ob der eine der andere umgedreht ist.
func IsReverse(s1, s2 string) bool {
	if s1 == Reverse(s2){
		return true
	}


	return false

}

// Erwartet einen String s und prüft, ob s ein Palindrom ist.
// Ein Palindrom ist ein String, der vorwärts und rückwärts gelesen gleich ist.
func IsPalindrome(s string) bool {
	if Reverse(s) == s{
		return true
	}

	return false

}
