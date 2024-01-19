# interface

type Notebook struct {
content []byte
}

func (nb *Notebook) Write(p []byte) (n int, err error) {
nb.content = append(nb.content, p...)
return len(p), nil
}
func (nb*Notebook) String() string {
return string(nb.content)
}
func main() {
nb := Notebook{}
fmt.Fprintln(&nb, "hello world")
fmt.Println(nb.String())
}
