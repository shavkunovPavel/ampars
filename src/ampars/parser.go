package main

import (
	"fmt"
	"sync"
)

func showError(u string, err error) {
	if err != nil {
		fmt.Println(u)
		fmt.Println(err)
	}
}

// обработчик урла
func doUrl(url_in string) (object *ResultType, err error) {
	object, err = Create_object(new(AmazonParser), url_in)
	return
}

// ассинхронно обрабатываем массив урлов
func process_urls(url_list *UrlsType) *[]interface{} {
	var wg sync.WaitGroup

	var ret []interface{}

	for _, urli := range *url_list.Urls {
		wg.Add(1)
		go func(urlgo string) {
			defer wg.Done()
			obj, err := doUrl(urlgo)
			if err != nil {
				showError(urlgo, err)
			}
			ret = append(ret, obj)
		}(urli.UrlString)
	}
	wg.Wait()

	return &ret
}
