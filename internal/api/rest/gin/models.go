package gin

// Go file used as backend for data
// May be replaced by any storage like filesystem, minio, ...

// album represents data about a record album.
type postAlbumBody struct {
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}
