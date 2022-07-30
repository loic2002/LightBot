package commands

import (
	"strings"

	"github.com/bwmarrin/discordgo"
)

type CommandHandler struct {
	prefix string

	cmdInstances []Command
	cmdMap       map[string]Command
	middlewares  []Middleware

	OnError func(err error, ctx *Context)
}

func NewCommandHandler(prefix string) *CommandHandler {
	return &CommandHandler{
		prefix:       prefix,
		cmdInstances: make([]Command, 0),
		cmdMap:       make(map[string]Command),
		middlewares:  make([]Middleware, 0),
		OnError:      func(error, *Context) {},
	}
}

func (c *CommandHandler) RegisterCommand(cmd Command) {
	c.cmdInstances = append(c.cmdInstances, cmd)

	for _, invoke := range cmd.Invokes() {
		c.cmdMap[invoke] = cmd
	}
}

func (c *CommandHandler) RegisterMiddleware(mv Middleware) {
	c.middlewares = append(c.middlewares, mv)
}

func (c *CommandHandler) HandleMessage(s *discordgo.Session, e *discordgo.MessageCreate) {
	if e.Author.ID == s.State.User.ID || e.Author.Bot || !strings.HasPrefix(e.Content, c.prefix) {
	return
	}

	split := strings.Split(e.Content[len(c.prefix):], " ")
	if len(split) < 1 {
		return
	}
	invoke := strings.ToLower(split[0])
	args := split[1:]

	cmd, ok := c.cmdMap[invoke]
	if !ok || cmd ==nil {
		return
	}

	ctx := &Context{
		Session: s,
		Args: args,
		Handler: c,
		Message: e.Message,
	}

	for _, mv := range c.middlewares{
		next, err := mv.Exec(ctx, cmd)
		if err != nil{
			c.OnError(err, ctx)
			return
		}
		if !next {
			return
		}
	}

	if err := cmd.Exec(ctx); err != nil {
			c.OnError(err,ctx)
	}
}