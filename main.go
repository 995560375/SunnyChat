package main

import (
	"SunnyChat/myRouter"
)

func main() {
	r := myRouter.Router()
	r.Run(":8081") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
