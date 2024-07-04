# Site Checker Telegram Bot

This checker is used to check the status of your sites for their availability!

A common situation is that your company services dozens or even hundreds of sites at once, and you need to constantly monitor their condition

## Setting:
TL;DR; All you need to do is:
1. Generate your bot using @BotFather
2. Add the bot to the group in which you want it to send a notification about unavailable services
3. In the **config/service.yaml** file, enter the following parameters
   * checker_bot_token - the token of your bot created through BotFather
   * chat_id - The ID of your chat where you added your bot which will notify about unavailable services (You can find the group ID by googling or [by following the link](https://docs.b2core.b2broker.com/en/how-to-get-telegram-chat-groups-and-channel-identifiers.html))
   * sites - Fill in the array all the URLs that need to be checked
4. DONE, next do sh ```$ make & make run```
5. You can configure the cron tab to run the generated script whenever you need it.