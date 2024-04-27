package foo

type Foo interface {
	Bar(n int) int
}

func SUT(f Foo) int {
	return f.Bar(2)
}
