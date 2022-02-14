package checker

import (
	"context"
	"github.com/chromedp/chromedp"
	"github.com/fatih/color"
	"github.com/pkg/errors"
	"main/internal/chrome"
	"main/internal/chrome/tasks"
	"main/internal/structures"
	"os"
	"path/filepath"
)

func Verify(i structures.AccInfo) {
	ctx, cancel := chrome.ChromeConfiguration()
	defer cancel()
	filePrefix, _ := filepath.Abs("/var/www/investments-auto-registration-rambler/captcha/")

	var args = structures.Args{
		I:          i,
		Prefix:     GenerateString(20),
		FilePrefix: filePrefix,
	}

	FirstStep(ctx, args)
}

func FirstStep(ctx context.Context, args structures.Args) {
	var b []byte
	if err := chromedp.Run(ctx, tasks.RamblerFirstStep(os.Getenv("RamblerLoginUrl"), args.I, &b)); err != nil {
		color.New(color.FgRed).Add(color.Underline).Println(errors.Wrap(err, "Couldn't launch chrome browser"))
	}
}
