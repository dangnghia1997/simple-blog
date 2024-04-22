package main

import "github.com/dangnghia1997/simple-blog/internal/routes"

func main() {
	r := routes.NewRouter()
	r.Run() // listen and serve on 0.0.0.0:8080
}
