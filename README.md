# mangadex

A lightweight Go client and HTTP‐handler wrappers for the MangaDex API.

## Installation

```bash
go get github.com/Seann-Moser/mangadex
```

```go
client, err := mangadex.NewClient("https://api.mangadex.org")
if err != nil {
    return nil, err
}
```


### List Chapter Pages and Constructing Page Urls

```go
u, err := uuid.Parse("UUID")
if err != nil {
  http.Error(w, "invalid UUID", http.StatusBadRequest)
  return
}
resp, err := client.GetAtHomeServerChapterId(ctx, u, nil)
if err != nil {
  slog.Error("failed searching for books mangadex", "err", err)
  return
}
defer resp.Body.Close()

chapterPages := client.GetAtHomeServerChapterIdResponse{}
err = json.NewDecoder(resp.Body).Decode(&chapterPages)
if err != nil {
  slog.Error("failed searching for books mangadex", "err", err)
  return
}

for i, p := range *chapterPages.JSON200.Chapter.Data {
  (*chapterPages.JSON200.Chapter.Data)[i], err = url.JoinPath(*chapterPages.JSON200.BaseUrl, "data", *chapterPages.JSON200.Chapter.Hash, p)
  if err != nil {
    return
  }
}

for i, p := range *chapterPages.JSON200.Chapter.DataSaver {
  (*chapterPages.JSON200.Chapter.DataSaver)[i], err = url.JoinPath(*chapterPages.JSON200.BaseUrl, "data", *chapterPages.JSON200.Chapter.Hash, p)
  if err != nil {
    return
  }
}
```
