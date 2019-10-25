package shortCode

var (
	defaultSeeds = []string{
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
	defaultSeedsIndex int64 = 3
	defaultLength     int64 = 7
)

type Options struct {
	Seeds       []string
	seedsLength int64 //种子长度 等同于多少进制
	SeedsIndex  int64 //种子所在下标
	Len         int64 //字符串总长度
}

type Option func(o *Options)

//	generate code length
//	default value is 7
func CodeLength(len int64) Option {
	return func(o *Options) {
		o.Len = len
	}
}

//	base seeds
//	require: all strings must have same length and not repeated
func Seeds(seeds []string) Option {
	return func(o *Options) {
		o.Seeds = seeds
	}
}

//	seeds index
func SeedsIndex(index int64) Option {
	return func(o *Options) {
		o.SeedsIndex = index
	}
}

func newOptions(opts ...Option) *Options {
	o := &Options{
		Seeds:      defaultSeeds,
		SeedsIndex: defaultSeedsIndex,
		Len:        defaultLength,
	}

	for _, opt := range opts {
		opt(o)
	}

	return o
}
