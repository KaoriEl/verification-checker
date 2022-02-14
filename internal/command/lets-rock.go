package command

import (
	"context"
	"fmt"
	"github.com/fatih/color"
	"github.com/gammazero/workerpool"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/joho/godotenv"
	"github.com/pkg/errors"
	"main/internal/checker"
	"main/internal/structures"
	"path/filepath"
)

func initEnv() {
	filePrefix, _ := filepath.Abs("/var/www/investments-verification-checker/configs") // path from the working directory
	err := godotenv.Load(filePrefix + "/.env")
	if err != nil {
		color.New(color.FgRed).Add(color.Underline).Println(errors.Wrap(err, "ENV was not loaded correctly"))
	}
}

func RockNRoll() {
	initEnv()
	conn := DBConnection()
	defer conn.Close(context.Background())
	var accs []structures.CoinlistAccs
	err := pgxscan.Select(
		ctx,
		conn,
		&accs,
		`
		select coinlist_accs.id, cc.cid, status_verify, rm.email, rm.password from coinlist_accs 
		LEFT JOIN (select id,cid from clients) as cc 
		ON cc.id = coinlist_accs.client_id
		LEFT JOIN (select id, email, password from rambler_mails) as rm 
		ON coinlist_accs.rambler_mail_id = rm.id where "status_verify" = $1;
		`,
		"no",
	)
	if err != nil {
		fmt.Println(err)
	}
	var maxWorkers int

	if len(accs) < 10 {
		maxWorkers = len(accs)
	} else {
		maxWorkers = 10
	}
	fmt.Println(maxWorkers)
	wp := workerpool.New(maxWorkers)
	for _, i := range accs {
		r := i
		wp.Submit(func() {
			fmt.Println(r)
			checker.Verify(r)
		})
	}

	wp.StopWait()

}
