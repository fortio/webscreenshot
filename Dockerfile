FROM chromedp/headless-shell:latest
COPY webscreenshot /usr/bin/webscreenshot
VOLUME /tmp
ENTRYPOINT ["/usr/bin/webscreenshot"]
CMD ["-h"]
