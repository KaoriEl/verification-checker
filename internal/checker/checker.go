package checker

import (
	"context"
	"fmt"
	"github.com/chromedp/chromedp"
	"github.com/fatih/color"
	"github.com/pkg/errors"
	"main/internal/chrome"
	"main/internal/chrome/tasks"
	"main/internal/command/api"
	"main/internal/structures"
	"os"
	"path/filepath"
)

func Verify(i structures.CoinlistAccs) {
	ctx, cancel := chrome.ChromeConfiguration()
	defer cancel()
	filePrefix, _ := filepath.Abs("/var/www/investments-auto-registration-rambler/captcha/")

	var args = structures.Args{
		I:          i,
		Prefix:     GenerateString(20),
		FilePrefix: filePrefix,
	}

	FirstStep(ctx, args)
	SecondStep(ctx, args)

}

func FirstStep(ctx context.Context, args structures.Args) {
	var b []byte
	if err := chromedp.Run(ctx, tasks.RamblerFirstStep(os.Getenv("RamblerLoginUrl"), args.I, &b)); err != nil {
		color.New(color.FgRed).Add(color.Underline).Println(errors.Wrap(err, "Couldn't launch chrome browser"))
	}
}

func SecondStep(ctx context.Context, args structures.Args) {
	var res string
	var b []byte
	if err := chromedp.Run(ctx, tasks.RamblerSecondStep(&res, &b)); err != nil {
		color.New(color.FgRed).Add(color.Underline).Println(errors.Wrap(err, "Couldn't launch chrome browser"))
	}
	fmt.Println(res)
	if res == "found" {
		api.Status("Поздравляем! Ваша верификация одобрена и ваши бонусные средства за регистрацию в размере 500 руб уже зачислены на ваш счет\nВы можете запросить вывод в Профиле пользователя", args.I)
		api.ChangeVerificationStatus(args.I)
	}
}
