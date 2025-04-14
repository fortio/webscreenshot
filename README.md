# webscreenshot
Take a screenshot of a Web page after giving it a delay for rendering

## Install
```bash
go install github.com/fortio/webscreenshot@latest
```

## Usage
```bash
webscreenshot [-delay duration] URL > screenshot.png
```

e.g
```bash
webscreenshot -delay 5s \
    "https://demo.fortio.org/browse?url=qps_max-s1_to_s2-0.7.1-2018-04-05-22-04.json" \
    > screenshot.png
open screenshot.png
```
