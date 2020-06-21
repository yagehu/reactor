package errs

type Kind int

const (
	KindOther Kind = iota
	KindReagentNotFound
)

func (k Kind) String() string {
	switch k {
	case KindReagentNotFound:
		return "reagent not found"
	default:
		return "unknown error kind"
	}
}
