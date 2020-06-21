package errs

type Kind int

const (
	KindOther Kind = iota
)

func (k Kind) String() string {
	switch k {
	default:
		return "unknown error kind"
	}
}
