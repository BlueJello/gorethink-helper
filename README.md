GoRethink-Helper - GoRethink helper for Go
======
*** Note: Lots of the code so far is based on github.com/jinzhu/gorm ***
### Basic Usage

```go
import (
    rh "github.com/MattAitchison/gorethink-helper"
    )

type Post struct {
    ID      string
    Title   string
    Content string
}
func GetAllPosts() ([]*Post, error) {
    // This will basically create the
    // following query. Then convert
    // it to a struct.
    // r.Table("post").Run(sess)
    var posts []*Post
    err := rh.All(&posts)
    return posts, err

}
func main() {

}

```
