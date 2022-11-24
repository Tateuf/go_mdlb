package main

import (
	b64 "encoding/base64"
	"fmt"
	"html"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"sync"
)

func main() {
	colormap := give_color()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.Error(w, "404 not found.", http.StatusNotFound)
			return
		}

		if r.Method != "GET" {
			http.Error(w, "Method is not supported.", http.StatusNotFound)
			return
		}

		max_iteration, err_max_iteration := strconv.ParseFloat(r.URL.Query().Get("max_iteration"), 64)
		fmt.Println(max_iteration)
		resolution, err_resolution := strconv.Atoi(r.URL.Query().Get("resolution"))
		fmt.Println(resolution)
		start_position_x, err_start_position_x := strconv.ParseFloat(r.URL.Query().Get("start_position_x"), 64)
		fmt.Println(start_position_x)
		start_position_y, err_start_position_y := strconv.ParseFloat(r.URL.Query().Get("start_position_y"), 64)
		fmt.Println(start_position_y)
		quantize_length, err_quantize_length := strconv.ParseFloat(r.URL.Query().Get("quantize_length"), 64)
		fmt.Println(quantize_length)
		if err_max_iteration == nil {
			if err_resolution == nil {
				if err_start_position_x == nil {
					if err_start_position_y == nil {
						if err_quantize_length == nil {
							wg := new(sync.WaitGroup)
							wg.Add(1)
							go png_generator(resolution, resolution, start_position_x, start_position_y, quantize_length, max_iteration, colormap, wg)
							wg.Wait()
							bytes, err := ioutil.ReadFile("./image.png")
							if err != nil {
								log.Fatal(err)
							}
							var base64Encoding string
							base64Encoding += "data:image/png;base64,"
							base64Encoding += b64.StdEncoding.EncodeToString(bytes)
							fmt.Fprintf(w, base64Encoding, html.EscapeString(r.URL.Path))
						} else {
							fmt.Println(err_quantize_length)
						}
					} else {
						fmt.Println(err_start_position_y)
					}
				} else {
					fmt.Println(err_start_position_x)
				}
			} else {
				fmt.Println(err_resolution)
			}
		} else {
			fmt.Println(err_max_iteration)
		}
	})

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
