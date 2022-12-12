# Cookie Clicker TASBot
Cookies with rutheless efficiency

[![Watch the Demo](https://img.youtube.com/vi/6NtXC0P3Oz0/default.jpg)](https://youtu.be/6NtXC0P3Oz0)

## How
If you're on most modern Unix-like systems, this will work out of the box.

```bash
$ go run .
# ... or ...
$ go run github.com/niko-dunixi/cookie-clicker-tasbot
```

Alternatively you should run with the utility script if:
* If you're on WSL because you'll need special handling for chromedp to work
* Or you like fun colors and cool countdowns

```bash
$ ./run-me.sh
```

## Dependencies
* lolcat
* figlet
* GoLang
* Chromedp compatible browser
  * Vanila Google-Chrome
  * Chromium

## Roadmap

- [X] Basic gameplay
- [ ] Save creation/restoration
- [ ] Intelligent upgrade purchases
  - [ ] Fastest upgrade route to [Grandmapocalypse](https://cookieclicker.fandom.com/wiki/Grandmapocalypse)
