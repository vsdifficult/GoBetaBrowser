package models

var Bookmarks []string

func AddBookmark(url string) {
    Bookmarks = append(Bookmarks, url)
}

func RemoveBookmark(url string) {
    for i, bookmark := range Bookmarks {
        if bookmark == url {
            Bookmarks = append(Bookmarks[:i], Bookmarks[i+1:]...)
            break
        }
    }
}

func GetBookmarks() []string {
    return Bookmarks
}