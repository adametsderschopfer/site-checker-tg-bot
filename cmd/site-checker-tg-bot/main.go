package main

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	checker "portal-service-checker/internal/checkers"
	"portal-service-checker/internal/config"
)

type SiteState struct {
	url         string
	isAvailable bool
	hasError    bool
}

func main() {
	cfg := config.MustLoad()

	bot, err := tgbotapi.NewBotAPI(cfg.CheckerBotToken)
	if err != nil {
		log.Panic(err)
	}

	var notAvailableSites []SiteState

	for _, site := range cfg.Sites {
		availability, err := checker.CheckSiteAvailability(site)
		if err != nil {
			notAvailableSites = append(notAvailableSites, SiteState{
				url:         site,
				isAvailable: availability,
				hasError:    true,
			})
			continue
		}

		if !availability {
			notAvailableSites = append(notAvailableSites, SiteState{
				url:         site,
				isAvailable: availability,
				hasError:    false,
			})
		}
	}

	if len(notAvailableSites) != 0 {
		var botMessage string = "*Unavailable services detected:*\n"

		for index, checkedSite := range notAvailableSites {
			status := `*Unavailable*`
			messageAboutError := ""
			count := index + 1

			if checkedSite.hasError {
				messageAboutError = ` _[Something went wrong on site check]_`
			}

			botMessage += fmt.Sprintf("\n%d. %s => %s%s", count, checkedSite.url, status, messageAboutError)
		}

		msg := tgbotapi.NewMessage(cfg.ChatId, botMessage)
		msg.ParseMode = tgbotapi.ModeMarkdown
		_, err = bot.Send(msg)
		if err != nil {
			log.Panic(err)
		}
	}
}
