package object

type Commit struct {
	Pointer   string
	Parents   []string
	Author    string
	Committer string
	Message   string
}
