package telegramauth

import (
	"log"
	"strconv"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var Tbot *Bot

func SetTbot(b *Bot) {
	Tbot = b
}

// link is the link with which the hash is merged (https://somesite/)
type Bot struct {
	*tgbotapi.BotAPI
	ustore        *userTokenStore
	usertokensize int
	link          string
}

type BotOptions struct {
	TokenBot, Link string
	UserTokenSize  int
	TTLusertoken   time.Duration
}

func NewAuthBot(opt BotOptions) (*Bot, error) {
	bot, err := tgbotapi.NewBotAPI(opt.TokenBot)
	if err != nil {
		return &Bot{}, err
	}

	ustore := newTokenStore(opt.TTLusertoken)

	return &Bot{
		BotAPI:        bot,
		link:          opt.Link,
		usertokensize: opt.UserTokenSize,
		ustore:        ustore,
	}, nil
}

func (b *Bot) Start() {
	log.Printf("Authorized on account %s", b.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := b.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		if update.Message.IsCommand() {
			switch update.Message.Command() {
			case "login":
				timenow := strconv.FormatInt(time.Now().Unix(), 10)
				str := strconv.FormatInt(update.Message.Chat.ID, 10) + timenow

				utoken := getUserTokenstring(str, b.usertokensize)
				b.ustore.SaveUserToken(utoken, strconv.FormatInt(update.Message.Chat.ID, 10))

				loginlink := b.link + utoken

				msg := tgbotapi.NewMessage(update.Message.Chat.ID, loginlink)
				b.Send(msg)

			case "down":
				// b.GetAvatarById(385305675)
			}
		}
	}
}

func (b *Bot) IsUsertokenExists(usertoken string) (exists bool, telegramID string) {
	return b.ustore.ValidateUserToken(usertoken)
}

// to do
// func (b *Bot) GetAvatarById(id int64) (url string) {
// 	photos, err := b.GetUserProfilePhotos(
// 		tgbotapi.UserProfilePhotosConfig{
// 			UserID: id,
// 			Limit:  1,
// 		},
// 	)
// 	if err != nil {
// 		fmt.Println("Error on getting profile photo")
// 		return ""
// 	}

// 	if photos.TotalCount < 1 {
// 		fmt.Println("User has no photos")
// 		return ""
// 	}

// 	fileID := photos.Photos[0][0].FileID

// 	file, err := b.GetFile(
// 		tgbotapi.FileConfig{
// 			FileID: fileID,
// 		},
// 	)
// 	if err != nil {
// 		fmt.Println("Error on getting file photo")
// 		return ""
// 	}

// 	url = file.Link(b.Token)

// 	err = downloadFile("avatar.jpg", url)
//         if err != nil {
//             fmt.Println("Ошибка при загрузке файла:", err)
//             return
//         }

//     fmt.Println("Файл успешно загружен и сохранен как avatar.jpg")

// 	return url
// }

// func downloadFile(filepath string, url string) error {
// 	// Отправляем GET-запрос
// 	response, err := http.Get(url)
// 	if err != nil {
// 		return err
// 	}
// 	defer response.Body.Close()

// 	// Создаем файл
// 	file, err := os.Create(filepath)
// 	if err != nil {
// 		return err
// 	}
// 	defer file.Close()

// 	// Копируем данные из ответа в файл
// 	_, err = io.Copy(file, response.Body)
// 	return err
// }
