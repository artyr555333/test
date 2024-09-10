package telegram

const msgHelp = `I can save and keep you pages. Also I can offer you then to read.

In order to save the page, just send me al link to it.

In order to get a random page from your list, send me command /rnd.
Caution! After that, this page will be removed from your list`

const msgHello = "Hi there!✋ \n\n" + msgHelp

const (
	msgUnknownCommand = "Неизвестная команда! 😕"
	msgNoSavedPages   = "У вас нет сохраненных ссылок! 🧐"
	msgSaved          = "Сохранено!👌"
	msgAlreadyExists  = "Эта ссылка уже итак сохранена в вашем листе!🙂"
)
