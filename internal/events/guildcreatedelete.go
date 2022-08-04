package events

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

type GuildCreateDelete struct{}

func NewGuildCreateDelete() *GuildCreateDelete {
	return &GuildCreateDelete{}
}

func (h *GuildCreateDelete) HandlerCreateGuild(s *discordgo.Session, g *discordgo.GuildCreate) {

	fmt.Printf("Guild create: %s", g.Name)
}

func (h *GuildCreateDelete) HandlerDeleteGuild(s *discordgo.Session, g *discordgo.GuildDelete) {

	fmt.Printf("Guild delete: %s", g.Name)
}