package main

type User struct {
	UserID   int                    `json:"user_id"`
	Username string                 `json:"username"`
	Map      map[string]interface{} `json:"map"`
	Books    []Book                 `json:"books"`
	Book     Book                   `json:"book"`
	BookStar *Book                  `json:"book_star"`
}
type Book struct {
	Name string `json:"name"`
}
