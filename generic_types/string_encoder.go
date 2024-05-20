package generic_types

type StringEncoder interface {
	Encode() ([]byte, error)
	Length() int
}

type StringType struct {
	msg string
}

func (st *StringType) Encode() ([]byte, error) {
	return []byte(st.msg), nil
}

func (st *StringType) Length() int {
	return len(st.msg)
}
