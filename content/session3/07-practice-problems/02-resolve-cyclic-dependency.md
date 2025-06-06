### Problem 2: Resolve Cyclic Dependency
Fix this cyclic dependency:
```go
// blog/post.go
package blog
import "github.com/myapp/user"
type Post struct {
    Author user.User
}

// user/profile.go
package user
import "github.com/myapp/blog"
type User struct {
    Posts []blog.Post
}
```
---
*Copyright © 2025 Mike Schinkel and NewClarity Consulting, LLC. All rights reserved.*
