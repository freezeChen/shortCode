package shortCode

import (
	"errors"
	"strconv"
	"strings"
)

type shortCode struct {
	o *Options
}

var _seeds = []string{
	"hLGBDC3IXiT7Qpntu5r4kZRvzc9sFbVMASjdyql01fHNeKWUwx82gmEYP6aJ",
	"3KYLSXtADP25ZaqVHNRkFg68x4jcfwB7zshn0UTIvWdmp1EQ9iGrMybeClJu",
	"F4DMKB2S8g6ryCexpXtTwl73AYIZVcHh1vjiz0PsqUaGkbNRQnJ5W9uELmdf",
	"QzuC6fhbeSlyRMA1LT9iUc4nBVGWYw2HkI8jsrJqtDN537PZKXmFgEdp0vxa",
	"eQ4TJmchW1vC2XPdyFfqR0zYEtkx8KjDV6wnNgaBiuILb7Z3p9Us5lrGHMSA",
	"rR8eCfVPZty3TFUWNczxGaqLMm1j6JshiDbgA5ndHv0XEp9uQk4lYBKI2w7S",
	"cpxi12IayQ6rqWSn5hTGw84VCgMXPb7sDklYUz3AfBFvEjJ0tHNLZedmRu9K",
	"U1Qkve80GClx3dbKzcfwEqInPuMiYaAmShr9pZVy5RTWJXB6N7H4L2FjsgtD",
	"sXURIcVCg5ZTktxKEMFyBJn6Pp0lYr2f9diwAhGvSeW3b4jaD1zL7NmuQH8q",
	"LlC0SkTrbJP7aI2mgM5j8ywzvW6sHfDRnGteXVU1QxB3EhKYpu4cZqA9NFid",
}

const (
	_minNum = 777600000
	_maxNum = 46656000000
)

func (s *shortCode) Encode(id int64) (string, error) {
	if id < _minNum || id > _maxNum {
		return "", errors.New("id not in range")
	}

	var (
		seedNum int64
		seed    string //种子
		result  string
	)

	seedNum = id % 10
	seed = _seeds[seedNum]
	for id > 0 {
		yu := id % 60
		id = id / 60
		result = string(seed[yu]) + result
	}

	return result[:3] + strconv.FormatInt(seedNum, 10) + result[3:], nil
}

func (s *shortCode) Decode(code string) (int64, error) {
	var (
		seedNum int64
		seed    string
		result  int64
		err     error
	)

	seedNum, err = getSeed(code)
	if err != nil {
		return 0, err
	}
	seed = _seeds[seedNum]
	var i = 1 * 60 * 60 * 60 * 60 * 60

	str := code[:3] + code[4:]

	for _, ch := range str {
		index := strings.Index(seed, string(ch))
		result = int64(index*i) + result
		i = i / 60
	}

	return result, nil
}

func getSeed(code string) (int64, error) {
	seedNum, err := strconv.ParseInt(string(code[3]), 10, 64)
	if err != nil {
		return 0, err
	}

	return seedNum, nil
}

func NewShortCode(opts ...Option) *shortCode {
	o := newOptions(opts...)
	return &shortCode{
		o: o,
	}

}
