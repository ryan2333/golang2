package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func fetchUrl(url string) ([]string, error) {
	var urls []string
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(fmt.Sprintf("%d", resp.StatusCode))
	}
	doc, err := goquery.NewDocumentFromResponse(resp)
	if err != nil {
		return nil, err
	}
	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		link, _ := s.Attr("href")
		if strings.HasPrefix(link, "javascript") {

		} else if !strings.HasPrefix(link, "http") {
			urls = append(urls, url+link)
		} else {
			urls = append(urls, link)
		}
	})
	return urls, nil
}

func fetchImgUrl(urls []string) ([]string, error) {
	var imgUrls []string
	for _, url := range urls {
		resp, err := http.Get(url)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()
		if resp.StatusCode != http.StatusOK {
			return nil, errors.New(fmt.Sprintf("%d", resp.StatusCode))
		}
		doc, err := goquery.NewDocumentFromResponse(resp)
		if err != nil {
			return nil, err
		}
		doc.Find("img").Each(func(i int, s *goquery.Selection) {
			link, _ := s.Attr("src")
			if strings.HasPrefix(link, "javascript") {

			} else if !strings.HasPrefix(link, "http") {
				imgUrls = append(imgUrls, url+link)
			} else {
				imgUrls = append(imgUrls, link)
			}
		})
	}
	return imgUrls, nil

}

func getImage(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return "", errors.New(fmt.Sprintf("%d", resp.StatusCode))
	}
	imgSplit := strings.Split(url, "/")
	imgName := imgSplit[len(imgSplit)-1]
	ifh, err := os.Create("imgs/" + imgName)
	if err != nil {
		return "", err
	}
	n, err := io.Copy(ifh, resp.Body)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%d bytes download!", n), nil
}

func main() {
	url := "http://daily.zhihu.com"
	// url := os.Args[1]
	// url := "http://www.douban.com/group/haixiuzu"
	urls, err := fetchUrl(url)
	if err != nil {
		log.Fatal(err)
	}
	imgUrls, err1 := fetchImgUrl(urls)
	if err1 != nil {
		log.Fatal(err1)
	}
	for _, link := range imgUrls {
		num, err := getImage(link)
		if err != nil {
			log.Println(err)
			continue
		}
		fmt.Println(num)
	}

}
