package tasks

import (
	"github.com/chromedp/chromedp"
	"main/internal/structures"
	"time"
)

const (
	// LetLink Find CoinList
	LetLink                          = `let link = document.querySelectorAll('span');`
	ConvertToArrayAndFilterByText    = `link = Array.from( link ).filter( e => (/Your account is ready - Start trading now/i).test( e.textContent ) );`
	ClickOnTheFirstElementOfTheArray = `if (link.length > 0){ link = "found" }else{ link = "Not found" }`
)

func RamblerFirstStep(url string, i structures.CoinlistAccs, buffer *[]byte) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(url),
		chromedp.WaitVisible(`//*[@id="login"]`),
		chromedp.Click(`//*[@id="login"]`, chromedp.NodeVisible),
		chromedp.SendKeys(`//*[@id="login"]`, i.Email),
		chromedp.WaitVisible(`//*[@id="password"]`),
		chromedp.Click(`//*[@id="password"]`, chromedp.NodeVisible),
		chromedp.SendKeys(`//*[@id="password"]`, i.Password),
		chromedp.Sleep(1 * time.Second),
		chromedp.Click(`//*[@id="__next"]/div/div/div[2]/div/div/div/div[1]/form/button/span`, chromedp.NodeVisible),
		chromedp.Sleep(2 * time.Second),
	}
}

func RamblerSecondStep(res *string, buffer *[]byte) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.WaitVisible(`//*[@id="app"]/div[2]/div[3]/div[2]/div[2]/div[1]/div/div/div[1]/div/div/div/div`),
		chromedp.Sleep(1 * time.Second),
		chromedp.Evaluate(LetLink+ConvertToArrayAndFilterByText+ClickOnTheFirstElementOfTheArray, &res),
	}
}
