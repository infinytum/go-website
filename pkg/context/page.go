package context

type PageContext struct {
	Custom   interface{}
	Title    string
	Language map[string]interface{}
}
