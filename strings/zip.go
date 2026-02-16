package strings

// Erwartet zwei strings s1 und s2.
// Liefert einen String, der abwechselnd aus den Buchstaben des einen und des anderen
// Strings besteht.
func Zip(s1, s2 string) string {
	// TODO
	result := ""
	min := len(s1)
	if len(s2)< min{
		min = len(s2)
	}

	for i:=0;i<min;i++{
	result += string(s1[i])+string(s2[i])
	}
	result += s1[min:]
	result += s2[min:]
	return result
}
