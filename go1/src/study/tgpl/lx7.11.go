package main

import (
	"fmt"
	"net/http"
	"strconv"
)

type RMB float32

func (p RMB) String() string {
	return fmt.Sprintf("￥%.2f元", p)
}

type database map[string]RMB

var DB = database{"衬衫": 200, "长裤": 100}

func (DB *database) list(w http.ResponseWriter, r *http.Request) {
	for k := range *DB {
		fmt.Fprintf(w, "%s:%s\n", k, (*DB)[k])
	}
}

func (DB *database) price(w http.ResponseWriter, r *http.Request) {
	item := r.URL.Query().Get("item")
	rmb, ok := (*DB)[item];
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		if item == "" {
			fmt.Fprintf(w, "对不起，我们不知道你想要什么")
		} else {
			fmt.Fprintf(w, "对不起，我们这里没有您要的 %s", item)
		}
		return
	}

	fmt.Fprintf(w, "%s 的价格是 %s", item, rmb)
}

func (DB *database) update(w http.ResponseWriter, r *http.Request) {
	item := r.URL.Query().Get("item")
	priceStr := r.URL.Query().Get("price")
	price, err := strconv.ParseFloat(priceStr, 32)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "对不起，您输入的价格有问题 %s", priceStr)
		return
	}

	_, ok := (*DB)[item];
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "对不起，我们这里没有您要的 %s", item)
		return
	}

	(*DB)[item] = RMB(price)
	fmt.Fprintf(w, "商品价格更新成功! %s 的价格是 %s", item, RMB(price))
}

func main() {
	server := http.NewServeMux()
	server.Handle("/list", http.HandlerFunc(DB.list))
	server.HandleFunc("/price", DB.price) //简化写法
	server.HandleFunc("/update", DB.update)
	fmt.Println(http.ListenAndServe("localhost:8000", server))
}
