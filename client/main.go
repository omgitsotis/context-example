package main

import(
	"context"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	ctx := context.Background()

	req, err := http.NewRequest(http.MethodGet, "http://localhost:4040", nil)
	if err != nil {
		log.Fatal(err)
	}

	ctx = context.WithValue(ctx, "foo", "bar")

	req = req.WithContext(ctx)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	if res.StatusCode != 200 {
		log.Fatal(res.Status)
	}

	defer res.Body.Close()

	io.Copy(os.Stdout, res.Body)
}