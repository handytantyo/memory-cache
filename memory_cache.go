package memorycache

// Cache interface for memory-cache library
type Cache interface {
	Get(key string) (result interface{})
	Set(key string, value interface{}) (err error)
	Count() int
}
