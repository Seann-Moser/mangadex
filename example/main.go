package main

import (
	"context"
	"encoding/json"

	"github.com/Seann-Moser/mangadex"
)

func main() {
	ctx := context.Background()
	c, err := mangadex.NewClientWithResponses("https://api.mangadex.org")
	if err != nil {
		return
	}
	r, err := c.GetSearchMangaWithResponse(ctx, &mangadex.GetSearchMangaParams{Title: mangadex.Ptr("Fly Me to The Moon")})
	if err != nil {
		return
	}
	d, _ := json.MarshalIndent(r.JSON200.Data, "", "  ")
	println(string(d))

}
