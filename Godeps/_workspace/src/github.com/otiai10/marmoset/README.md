marmoset
========

less than "web framework", just make your code a bit DRY.

```go
func main() {

	r := marmoset.NewRouter()

	r.GET("/foo", your.FooHttpHandlerFunc)
	r.POST("/bar", your.BarHttpHandlerFunc)

	r.Static("/public", "/your/assets/path")

	s := marmoset.NewFilter(r).Add(&your.Filter{}).Server()

	http.ListenAndServe(":8080", server)
}
```
