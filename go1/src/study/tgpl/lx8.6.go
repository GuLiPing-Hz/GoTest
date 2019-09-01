package main

import (
	"net/http"
	"io/ioutil"
	"bytes"
	"golang.org/x/net/html"
	"fmt"
	"sync"
)

func detactAUrl(node *html.Node) string {
	if node.Type == html.ElementNode && node.Data == "a" {
		for _, a := range node.Attr {
			if a.Key == "href" {
				return a.Val
			}
		}
	}

	return ""
}

func detactA(links []string, node *html.Node) []string {
	if node == nil {
		return nil
	}

	link := detactAUrl(node)
	if link != "" {
		links = append(links, link)
	}

	for i := node.FirstChild; i != nil; i = i.NextSibling {
		//fmt.Printf("label=%v\n", i)
		links = detactA(links, i)
	}

	return links
}

func detactALabel(data []byte) ([]string, error) {
	rootNode, err := html.Parse(bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	links := detactA(nil, rootNode)
	return links, nil
}

func detactUrl(rootUrl string) ([]string, error) {
	resp, err := http.Get(rootUrl)
	if err != nil {
		return nil, err
	}

	buf, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, err
	}

	//fmt.Printf("url=%s\n", string(buf))
	links, err := detactALabel(buf)
	return links, nil
}

type LinkDepths struct {
	links []string
	depth int
}

type LinkDepth struct {
	link  string
	depth int
}

func main() {
	const thread = 3
	wg := sync.WaitGroup{}
	chanLink := make(chan LinkDepths, thread)
	climbLink := make(chan LinkDepth, thread)
	for i := 0; i < thread; i++ {
		wg.Add(1)
		go func(index int) {
			defer func() {
				wg.Done()
				fmt.Printf("goroutine finish %d\n", index)
			}()

			for curLink := range climbLink {
				links, err := detactUrl(curLink.link)
				if err != nil {
					fmt.Printf("err=%s\n", err.Error())
					break
				}

				fmt.Println("###############################################")
				var rightLinks []string
				for i := range links {
					fmt.Printf("link[%d]=%s\n", i, links[i])
					data := links[i]
					if data == "#" {
						continue
					}
					if data[0] == '/' {
						if len(data) == 1 {
							continue
						}
						rightLinks = append(rightLinks, curLink.link+data)
					}
				}

				go func() {
					//这里并不完善，会由于信道过早关闭而panic，不过出于练习的目的已经达到了。
					chanLink <- LinkDepths{rightLinks, curLink.depth + 1}
				}()
			}
		}(i)
	}

	go func() {
		chanLink <- LinkDepths{[]string{"http://www.json.cn"}, 1}
	}()

	go func() {
		seenLinks := make(map[string]bool)
		for links := range chanLink {
			if links.depth >= 3 {
				//close(climbLink)
				continue
			}

			for i := range links.links {
				if !seenLinks[links.links[i]] {
					seenLinks[links.links[i]] = true
					climbLink <- LinkDepth{links.links[i], links.depth}
				}
			}

		}
	}()

	wg.Wait()
	close(chanLink)

	fmt.Println("Finish...")
}
