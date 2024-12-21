package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/net/html"
)

/*
	TODO:
	Утилита wget
	Реализовать утилиту wget с возможностью скачивать сайты целиком.

	Теория:
	Это утилита командной строки для загрузки файлов из интернета. Она поддерживает протоколы HTTP, HTTPS и FTP, что делает её очень полезной для скачивания содержимого веб-страниц, файлов и даже целых сайтов.
*/

func download(url string, domain string, isVisited map[string]bool) error {
	if isVisited[url] {
		return nil
	}
	isVisited[url] = true
	resp, err := http.Get(url)

	if err != nil {
		return fmt.Errorf("error during GET Method: %s", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("error reading url %s: %s", url, resp.Status)
	}

	fileName := filepath.Base(url)
	out, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer out.Close()

	tee := io.TeeReader(resp.Body, out)

	domTree, err := html.Parse(tee)
	if err != nil {
		return fmt.Errorf("error by %s html parsing: %s", url, err)
	}

	links := getAllLinks(nil, url, isVisited, domTree)
	fmt.Println("Downloading:", url)
	for _, link := range links {
		if isSameDomain(link, domain) {
			err := download(link, domain, isVisited)
			if err != nil {
				fmt.Fprintf(os.Stderr, "download error for %s: %v\n", link, err)
			}
		}
	}

	return nil
}

func isSameDomain(link string, domain string) bool {
	return strings.HasPrefix(link, domain) || strings.HasPrefix(link, "http://"+domain) || strings.HasPrefix(link, "https://"+domain)
}

func getAllLinks(links []string, absUrl string, isVisited map[string]bool, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				link := a.Val

				// Преобразуем относительные ссылки в абсолютные
				parsedAbsURL, err := url.Parse(absUrl)
				if err != nil {
					continue // Пропускаем, если не удалось разобрать базовый URL
				}

				parsedLink, err := url.Parse(link)
				if err != nil {
					continue // Пропускаем, если не удалось разобрать ссылку
				}

				// Если ссылка относительная, то разрешаем её относительно базового URL
				if !parsedLink.IsAbs() {
					parsedLink = parsedAbsURL.ResolveReference(parsedLink)
				}

				// Проверяем, совпадает ли домен
				if parsedLink.Host != parsedAbsURL.Host {
					continue // Пропускаем, если домены не совпадают
				}
				if !isVisited[parsedLink.String()] {
					links = append(links, parsedLink.String())
				}
			}
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = getAllLinks(links, absUrl, isVisited, c)
	}

	return links
}

func main() {
	isVisited := make(map[string]bool)

	for _, url := range os.Args[1:] {
		domain := strings.Split(url, "/")[2]
		err := download(url, domain, isVisited)
		if err != nil {
			fmt.Fprintf(os.Stderr, "download error: %v\n", err)
		}
	}
}
