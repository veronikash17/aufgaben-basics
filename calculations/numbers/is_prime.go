package numbers

// Erwartet eine Zahl n und prüft, ob n eine Primzahl ist.
func IsPrime(n int) bool {
	for i:=2; i < n; i++{
		if n%i == 0{
			return false
		}
	}
	return n>1
}
