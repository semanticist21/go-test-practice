package sort_example

type Student struct {
	Name string
	Age  int
}

type Students []Student

func (s Students) Len() int           { return len(s) }
func (s Students) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s Students) Less(i, j int) bool { return s[i].Age < s[j].Age }
