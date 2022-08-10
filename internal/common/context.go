package common

type Context struct {
	Config Config
}

func newContext() (*Context, error) {
	return &Context{}, nil
}

type Config struct {
}
