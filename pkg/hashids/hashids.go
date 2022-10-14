package hashids

import (
	"github.com/AH-dark/apollo/config"
	"github.com/speps/go-hashids/v2"
)

var hd *hashids.HashIDData

func Init() {
	hd = hashids.NewData()
	hd.Salt = config.System.HashIDSalt
	hd.MinLength = 4
	hd.Alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
}

func HashEncode(n []int) (string, error) {
	h, err := hashids.NewWithData(hd)
	if err != nil {
		return "", err
	}

	e, err := h.Encode(n)
	if err != nil {
		return "", err
	}

	return e, nil
}

func HashDecode(s string) ([]int, error) {
	h, err := hashids.NewWithData(hd)
	if err != nil {
		return nil, err
	}

	d, err := h.DecodeWithError(s)
	if err != nil {
		return nil, err
	}

	return d, nil
}
