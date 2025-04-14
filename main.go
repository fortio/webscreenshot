package main

import (
	"context"
	"flag"
	"os"
	"time"

	"fortio.org/cli"
	"fortio.org/log"
	"github.com/chromedp/chromedp"
)

func main() {
	os.Exit(Main())
}

func Main() int {
	delayFlag := flag.Duration("delay", 1*time.Second, "Delay before the screen shot")
	var width, height int64
	flag.Int64Var(&width, "width", 1280, "Width of the screenshot")
	flag.Int64Var(&height, "height", 720, "Height of the screenshot")
	cli.MinArgs = 1
	cli.ArgsHelp = "URL > screenshot.png\n" +
		"to take a screenshot of a web page using Chrome after a given delay"
	cli.Main()
	url := flag.Arg(0)
	delay := *delayFlag
	var res []byte
	log.Printf("Taking ~ %dx%d screenshot of %s after %s", width, height, url, delay)
	ctx, cancel := chromedp.NewContext(
		context.Background(),
		// chromedp.WithDebugf(log.Printf),
	)
	defer cancel()
	// Handle ^C gracefully
	go func() {
		cli.UntilInterrupted()
		cancel()
	}()
	err := chromedp.Run(ctx, chromedp.Tasks{
		chromedp.EmulateViewport(width, height),
		chromedp.Navigate(url),
		chromedp.ActionFunc(func(ctx context.Context) error {
			select {
			case <-time.After(delay):
				return nil
			case <-ctx.Done():
				return ctx.Err()
			}
		}),
		chromedp.FullScreenshot(&res, 100),
	})
	if err != nil {
		return log.FErrf("Failed to take screenshot: %v", err)
	}
	os.Stdout.Write(res)
	return 0
}
