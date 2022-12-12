package main

import (
	"context"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/chromedp/chromedp"
)

func main() {
	dir, err := ioutil.TempDir("", "chromedp-example")
	if err != nil {
		log.Fatal(err)
	}
	defer os.RemoveAll(dir)

	opts := append(
		// All defaults except for headless
		append(
			chromedp.DefaultExecAllocatorOptions[:2],
			chromedp.DefaultExecAllocatorOptions[3:]...,
		),
		// chromedp.DisableGPU,
		chromedp.UserDataDir(dir),
	)

	// timeoutDuration := time.Second * (60*3 + 55)
	// // timeoutDuration := time.Second * 15
	// timeoutCtx, timeoutCtxCancel := context.WithTimeout(context.Background(), timeoutDuration)
	// defer timeoutCtxCancel()
	parentCtx, cancelParentCtx := context.WithCancel(context.Background())
	defer cancelParentCtx()

	allocCtx, cancel := chromedp.NewExecAllocator(parentCtx, opts...)
	defer cancel()

	// also set up a custom logger
	taskCtx, cancel := chromedp.NewContext(allocCtx, chromedp.WithLogf(log.Printf))
	defer cancel()

	// Setup closer
	go func() {
		defer cancelParentCtx()
		signalChannel := make(chan os.Signal, 1)
		signal.Notify(signalChannel, syscall.SIGINT, syscall.SIGTERM)
		<-signalChannel
		log.Println("CTRL-C captured, canceling parent context")
	}()

	// Setup cookie clicker environment
	if err := chromedp.Run(taskCtx,
		chromedp.Navigate(`https://orteil.dashnet.org/cookieclicker/`),
		chromedp.Click(`#langSelect-EN`, chromedp.ByID, chromedp.NodeVisible),
	); err != nil {
		log.Fatal(err)
	}
	wg := &sync.WaitGroup{}
	go handleDismissAdNotice(taskCtx, wg)
	go handleClickBigCookie(taskCtx, wg)
	go handleClickShimmer(taskCtx, wg)
	go handleDismissNotificationNotes(taskCtx, wg)
	go handleStoreUpgrades(taskCtx, wg)
	go handleStoreProducts(taskCtx, wg)

	<-parentCtx.Done()
	wg.Wait()
	log.Println("END OF NIKO PWNAGE")
}

func handleDismissAdNotice(ctx context.Context, wg *sync.WaitGroup) {
	wg.Add(1)
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			return
		default:
			chromedp.Run(ctx, chromedp.Click(".cc_btn_accept_all", chromedp.BySearch, chromedp.NodeVisible))
			log.Println(`Accepting all "cookies"`)
			// chromedp.Run(ctx, chromedp.Sleep(time.Millisecond*500))
			time.Sleep(time.Millisecond * 500)
		}
	}
}

func handleClickBigCookie(ctx context.Context, wg *sync.WaitGroup) {
	wg.Add(1)
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			return
		default:
			chromedp.Run(ctx, chromedp.Click("#bigCookie", chromedp.NodeVisible))
		}
	}
}

func handleClickShimmer(ctx context.Context, wg *sync.WaitGroup) {
	wg.Add(1)
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			return
		default:
			chromedp.Run(ctx, chromedp.Click(".shimmer", chromedp.NodeVisible))
			log.Println("clicked shimmer cookie")
			// chromedp.Run(ctx, chromedp.Sleep(time.Millisecond*500))
			time.Sleep(time.Millisecond * 500)
		}
	}
}

func handleDismissNotificationNotes(ctx context.Context, wg *sync.WaitGroup) {
	wg.Add(1)
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			return
		default:
			chromedp.Run(ctx, chromedp.Click("#notes .close", chromedp.BySearch, chromedp.NodeVisible))
			log.Println("closing notification")
			// chromedp.Run(ctx, chromedp.Sleep(time.Millisecond*500))
			time.Sleep(time.Millisecond * 500)
		}
	}
}

func handleStoreProducts(ctx context.Context, wg *sync.WaitGroup) {
	wg.Add(1)
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			return
		default:
			chromedp.Run(ctx, chromedp.Click(".product.unlocked.enabled", chromedp.BySearch))
			log.Println("bought product")
			// chromedp.Run(ctx, chromedp.Sleep(time.Millisecond*500))
			time.Sleep(time.Millisecond * 500)
		}
	}
}

func handleStoreUpgrades(ctx context.Context, wg *sync.WaitGroup) {
	wg.Add(1)
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			return
		default:
			// This could be refined, but will take time.
			// Naive behavior will be to always select an upgrade
			// and then a product
			chromedp.Run(ctx, chromedp.Click(".crate.upgrade.enabled", chromedp.BySearch))
			log.Println("bought upgrade")
			// chromedp.Run(ctx, chromedp.Sleep(time.Millisecond*500))
			time.Sleep(time.Millisecond * 500)
		}
	}
}

func handleStoreAI(ctx context.Context, wg *sync.WaitGroup) {
	wg.Add(1)
	defer wg.Done()
	// Buy your first item, and then we can actually begin dynamic
	// behavior alternating between upgrades and products
	// - https://cookieclicker.fandom.com/wiki/Upgrades#Tiered_upgrades
	chromedp.Run(ctx, chromedp.Click(".product.unlocked.enabled", chromedp.BySearch))
	for {
		select {
		case <-ctx.Done():
			return
		default:
			// This could be refined, but will take time.
			// Naive behavior will be to always select an upgrade
			// and then a product
			chromedp.Run(ctx, chromedp.Click(".crate.upgrade.enabled", chromedp.BySearch))
			chromedp.Run(ctx, chromedp.Click(".product.unlocked.enabled", chromedp.BySearch))
		}
	}
}
