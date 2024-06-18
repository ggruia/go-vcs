package object

type Reference struct {
	Type string
	Hash string
	Name string
}

type Tree struct {
	Hash       string
	Size       int
	References []Reference
}
