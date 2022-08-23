package constant

type StatusCode struct {
	Code int32
	Msg  string
}

const (
	SUCCESS int32 = iota
	FAILED
)
