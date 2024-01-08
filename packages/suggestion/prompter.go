package suggestion

type Prompter interface {
	Build() (string, error)
	Update(...interface{}) error
}
