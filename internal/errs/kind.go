package errs

type Kind int

const (
	KindOther Kind = iota
	KindReagentNotFound
	KindInvalidService
)

func (k Kind) String() string {
	switch k {
	case KindReagentNotFound:
		return "reagent not found"
	case KindInvalidService:
		return "invalid service"
	default:
		return "unknown error kind"
	}
}
