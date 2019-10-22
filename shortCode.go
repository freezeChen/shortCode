package shortCode

import (
	"errors"
	"strconv"
	"strings"
)

type shortCode struct {
	o *Options
}

func (s *shortCode) Encode(id int64) (string, error) {

	if err := s.checkId(id); err != nil {
		return "", err
	}
	var (
		seedNum int64
		seed    string //种子
		result  string
	)

	seedNum = id % 10
	seed = s.o.Seeds[seedNum]
	for id > 0 {
		yu := id % s.o.seedsLength
		id = id / s.o.seedsLength
		result = string(seed[yu]) + result
	}

	return result[:s.o.SeedsIndex] + strconv.FormatInt(seedNum, 10) + result[s.o.SeedsIndex:], nil
}

func (s *shortCode) Decode(code string) (int64, error) {
	var (
		seedNum int64
		seed    string
		result  int64
		err     error
	)

	if s.o.Len != int64(len(code)) {
		return 0, errors.New("code length is error")
	}

	seedNum, err = s.getSeed(code)
	if err != nil {
		return 0, err
	}
	seed = s.o.Seeds[seedNum]

	var i int64 = 1
	for j := 0; j < int(s.o.Len)-2; j++ {
		i = i * s.o.seedsLength

	}
	str := code[:s.o.SeedsIndex] + code[s.o.SeedsIndex:]
	for codeIndex, ch := range str {
		if int64(codeIndex) == s.o.SeedsIndex {
			continue
		}
		index := strings.Index(seed, string(ch))
		result = (int64(index) * i) + result
		i = i / s.o.seedsLength
	}

	return result, nil
}

//获取种子
func (s *shortCode) getSeed(code string) (int64, error) {
	seedNum, err := strconv.ParseInt(string(code[s.o.SeedsIndex]), 10, 64)
	if err != nil {
		return 0, err
	}
	return seedNum, nil
}

func (s *shortCode) checkId(id int64) error {
	var i, min, max int64 = 1, 1, 1
	for i = 0; i < s.o.Len-1; i++ {
		max = s.o.seedsLength * max
	}
	min = max / s.o.seedsLength

	if id < min || id > max {
		return errors.New("id not in range")
	}

	return nil
}

func NewShortCode(opts ...Option) (*shortCode, error) {
	o := newOptions(opts...)

	if len(o.Seeds) != 10 {
		return nil, errors.New("only support ten seeds")
	}

	for _, seed := range o.Seeds {
		if o.seedsLength == 0 {
			o.seedsLength = int64(len(seed))
		}

		if o.seedsLength != int64(len(seed)) {
			return nil, errors.New("seeds must be same length")
		}
	}

	return &shortCode{
		o: o,
	}, nil

}
