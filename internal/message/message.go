package message

import (
	"net/http"
	"regexp"
	"sync"

	"github.com/mvdan/xurls"

	"code_exercise/internal/html"
	"code_exercise/internal/gorountine"
)

var httpClient *http.Client = &http.Client{
	Transport: &http.Transport{
		MaxIdleConnsPerHost: maxIdleConnection,
	},
	Timeout: timeout,
}

type Message struct {
	Mentions  []string `json:"mentions,omitempty"`
	Emoticons []string `json:"emoticons,omitempty"`
	Links     []Link   `json:"links,omitempty"`
}

func Parse(rawMessage string) Message {
	if len(rawMessage) == 0 {
		return Message{}
	}

	mentions := parseMentions(rawMessage)

	emoticons := parseEmoticons(rawMessage)

	links := parseLinks(rawMessage)

	return Message{Mentions: mentions, Emoticons: emoticons, Links: links}
}

func parseMentions(str string) []string {
	m := regexp.MustCompile(regexMention).FindAllStringSubmatch(str, -1)

	mentions := make([]string, 0)
	for _, value := range m {
		mentions = append(mentions, value[1])
	}

	return mentions
}

func parseEmoticons(str string) []string {

	e := regexp.MustCompile(regexEmoticons).FindAllStringSubmatch(str, -1)

	emoticons := make([]string, 0)
	for _, value := range e {
		emoticons = append(emoticons, value[1])
	}

	return emoticons
}

func parseLinks(str string) []Link {
	urls := getUrls(str)
	if len(urls) == 0 {
		return make([]Link, 0)
	}


	var wg sync.WaitGroup

	links := make([]Link, len(urls))

	for index, url := range urls {
		gorountine.Sem <- true
		wg.Add(1)
		go func(i int, u string) {
			defer wg.Done()
			title, _ := getTitle(u, httpClient)
			links[i] = Link{u, title}
			<- gorountine.Sem
		}(index, url)
	}

	wg.Wait()

	return links
}

func getUrls(str string) []string {
	return xurls.Strict.FindAllString(str, -1)
}

func getTitle(url string, client *http.Client) (string, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return "", err
	}

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	return html.GetTitle(resp.Body)
}
