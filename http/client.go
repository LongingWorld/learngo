package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

func main() {
	request, err := http.NewRequest(http.MethodGet, "http://imooc.com", nil)
	request.Header.Add("User-Agent", "Mozilla/5.0 (Linux; U; Android 4.3; en-us; SM-N900T Build/JSS15J) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Mobile Safari/534.30")

	client := http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			fmt.Println("Redirect:", req)
			return nil
		},
	}

	resp, err := client.Do(request)
	//resp, err := http.DefaultClient.Do(request)
	//resp, err := http.Get("http://www.baidu.com")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	bytes, e := httputil.DumpResponse(resp, true)
	if e != nil {
		panic(e)
	}
	fmt.Printf("%s\n", bytes)
}
