package inits

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/loic2002/LightBot/internal/commands"
	"github.com/loic2002/LightBot/internal/config"
	"github.com/loic2002/LightBot/internal/events"
)

var dgBot *discordgo.Session

func InitDiscordBot() {

	// Create a new Discord session using the provided bot token.
	dgBot, err := discordgo.New("Bot " + config.Token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	dgBot.Identify.Intents = discordgo.MakeIntent(
		discordgo.IntentGuildMembers | discordgo.IntentGuildMessages | discordgo.IntentGuilds)

	registerEvents(dgBot)
	registerCommnds(dgBot)

	// Open a websocket connection to Discord and begin listening.
	err = dgBot.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	dgBot.Close()

}

func registerEvents(dg *discordgo.Session){

	joinLeaveHandler := events.NewJoinLeaveHandler()
	dg.AddHandler(joinLeaveHandler.HandlerJoin)
	dg.AddHandler(joinLeaveHandler.HandlerLeave)

	guildCreateDeleteHandler := events.NewGuildCreateDelete()
	dg.AddHandler(guildCreateDeleteHandler.HandlerCreateGuild)
	dg.AddHandler(guildCreateDeleteHandler.HandlerDeleteGuild)

	dg.AddHandler(events.NewReadyHandler().Handler)
	dg.AddHandler(events.NewMessageHandler().Handler)


}

func registerCommnds(dg *discordgo.Session){

	cmdHandler := commands.NewCommandHandler(config.Prefix)
	cmdHandler.OnError = func(err error, ctx *commands.Context) {
		ctx.Session.ChannelMessageSend(ctx.Message.ChannelID,
			fmt.Sprintf("Command Execution failed: %s", err.Error()))
	}

	cmdHandler.RegisterCommand(&commands.CmdPing{})
	cmdHandler.RegisterMiddleware(&commands.MwPermissions{})

	dg.AddHandler(cmdHandler.HandleMessage)
	
}