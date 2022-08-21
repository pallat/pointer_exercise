package pointer

// n := 2
// double(&n)
// n == 4
func double(n *int) {
	*n *= 2
}

// name := "Bob"
// appendGreeting(&name)
// name == "Hi, Bob"
func appendGreeting(s *string) {
	*s = "Hi, " + *s
}
