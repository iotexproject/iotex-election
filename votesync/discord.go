package votesync

import "github.com/bwmarrin/discordgo"

type discord struct {
	botToken    string
	channelID   string
	newCycleMsg string
	reminderMsg string
	reminded    bool
}

func (d *discord) sendMessage(msg string) error {
	if d.botToken == "" || msg == "" {
		return nil
	}

	dg, err := discordgo.New("Bot " + d.botToken)
	if err != nil {
		return err
	}
	if err := dg.Open(); err != nil {
		return err
	}
	defer dg.Close()

	_, err = dg.ChannelMessageSend(d.channelID, msg)
	return err
}

func (d *discord) SendNewCycleMessage() error {
	d.reminded = false
	return d.sendMessage(d.newCycleMsg)
}

func (d *discord) SendReminder() error {
	d.reminded = true
	return d.sendMessage(d.reminderMsg)
}

func (d *discord) Reminded() bool {
	return d.reminded
}
