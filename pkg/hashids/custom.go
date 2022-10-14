package hashids

import "fmt"

type HashIDType int

const (
	UserHash HashIDType = iota
	CommentHash
)

var (
	ErrTypeNotMatch = fmt.Errorf("type not match")
)

func Encode(id uint, t HashIDType) string {
	s, _ := HashEncode([]int{int(id), int(t)})
	return s
}

func Decode(s string, t HashIDType) (uint, error) {
	d, err := HashDecode(s)
	if err != nil {
		return 0, err
	}

	if len(d) != 2 || d[1] != int(t) {
		return 0, ErrTypeNotMatch
	}

	return uint(d[0]), nil
}
