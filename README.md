# T.LY Go Client

This Go package provides a client for the T.LY URL Shortener API. It includes methods for all API endpoints such as managing pixels, short links, stats, and tags.

## Installation

**Obtain an API Token**: Sign up or log in to [T.LY](https://t.ly/settings#/api) and retrieve your API token from the T.LY dashboard.

To install the package, run:

```bash
go get github.com/timleland/t.ly-go-url-shortener-api
```

Then import it in your Go code:

```go
import "github.com/timleland/t.ly-go-url-shortener-api"
```

## Usage

Create a new client by providing your API token:

```go
client := tly.NewClient("YOUR_API_TOKEN")
```

### Pixel Management

#### Create a Pixel

```go
pixelReq := tly.PixelCreateRequest{
    Name:      "GTMPixel",
    PixelID:   "GTM-xxxx",
    PixelType: "googleTagManager",
}
pixel, err := client.CreatePixel(pixelReq)
if err != nil {
    // handle error
}
fmt.Println("Created Pixel:", pixel)
```

#### List Pixels

```go
pixels, err := client.ListPixels()
if err != nil {
    // handle error
}
fmt.Println("Pixels:", pixels)
```

#### Get a Pixel

```go
pixel, err := client.GetPixel(12345)
if err != nil {
    // handle error
}
fmt.Println("Pixel:", pixel)
```

#### Update a Pixel

```go
updateReq := tly.PixelUpdateRequest{
    ID:        12345,
    Name:      "UpdatedPixel",
    PixelID:   "GTM-xxxx",
    PixelType: "googleTagManager",
}
updatedPixel, err := client.UpdatePixel(updateReq)
if err != nil {
    // handle error
}
fmt.Println("Updated Pixel:", updatedPixel)
```

#### Delete a Pixel

```go
err = client.DeletePixel(12345)
if err != nil {
    // handle error
}
fmt.Println("Pixel deleted")
```

### Short Link Management

#### Create a Short Link

```go
shortLinkReq := tly.ShortLinkCreateRequest{
    LongURL: "http://example.com/",
    Domain:  "https://t.ly/",
}
shortLink, err := client.CreateShortLink(shortLinkReq)
if err != nil {
    // handle error
}
fmt.Println("Created Short Link:", shortLink)
```

#### Get a Short Link

```go
link, err := client.GetShortLink("https://t.ly/c55j")
if err != nil {
    // handle error
}
fmt.Println("Short Link:", link)
```

#### Update a Short Link

```go
updateLinkReq := tly.ShortLinkUpdateRequest{
    ShortURL: "https://t.ly/c55j",
    LongURL:  "http://updated-example.com/",
}
updatedLink, err := client.UpdateShortLink(updateLinkReq)
if err != nil {
    // handle error
}
fmt.Println("Updated Short Link:", updatedLink)
```

#### Delete a Short Link

```go
err = client.DeleteShortLink("https://t.ly/c55j")
if err != nil {
    // handle error
}
fmt.Println("Short Link deleted")
```

#### Expand a Short Link

```go
expandReq := tly.ExpandRequest{
    ShortURL: "https://t.ly/OYXL",
}
expanded, err := client.ExpandShortLink(expandReq)
if err != nil {
    // handle error
}
fmt.Println("Expanded URL:", expanded.LongURL)
```

#### List Short Links

```go
params := map[string]string{
    "search": "amazon",
}
links, err := client.ListShortLinks(params)
if err != nil {
    // handle error
}
fmt.Println("Short Links List:", links)
```

#### Bulk Shorten Links

```go
bulkReq := tly.BulkShortenRequest{
    Domain: "https://t.ly/",
    Links:  []string{"http://example1.com", "http://example2.com"},
}
result, err := client.BulkShortenLinks(bulkReq)
if err != nil {
    // handle error
}
fmt.Println("Bulk Shorten Result:", result)
```

### Stats Management

#### Get Stats for a Short Link

```go
stats, err := client.GetStats("https://t.ly/OYXL")
if err != nil {
    // handle error
}
fmt.Println("Stats:", stats)
```

### Tag Management

#### List Tags

```go
tags, err := client.ListTags()
if err != nil {
    // handle error
}
fmt.Println("Tags:", tags)
```

#### Create a Tag

```go
tag, err := client.CreateTag("fall2024")
if err != nil {
    // handle error
}
fmt.Println("Created Tag:", tag)
```

#### Get a Tag

```go
tag, err := client.GetTag(12345)
if err != nil {
    // handle error
}
fmt.Println("Tag:", tag)
```

#### Update a Tag

```go
updatedTag, err := client.UpdateTag(12345, "fall2025")
if err != nil {
    // handle error
}
fmt.Println("Updated Tag:", updatedTag)
```

#### Delete a Tag

```go
err = client.DeleteTag(12345)
if err != nil {
    // handle error
}
fmt.Println("Tag deleted")
```

## License

This project is licensed under the MIT License.
