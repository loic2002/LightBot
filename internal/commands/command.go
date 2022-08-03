package commands

type Command interface {
	Invokes() []string
	Description() string
	Version() string
	AdminRequired() bool
	Exec(ctx *Context) error
}