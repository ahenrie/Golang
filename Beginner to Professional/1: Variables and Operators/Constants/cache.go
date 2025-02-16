package constants

const GlobalLimit = 100
const MaxCacheSize = 10 * GlobalLimit

const (
	CacheKeyBook = "book_"
	CackeKeyCD   = "cd_"
)

type Cache map[string]string

func (c Cache) cacheSet(key, val string) {
	//Check size so we do not go over our cache size
	if len(c)+1 >= MaxCacheSize {
		return
	}
	c[key] = val
}

func (c Cache) SetBook(isbn string, name string) {
	c.cacheSet(CacheKeyBook+isbn, name)
}

// Get item from the cache
func (c Cache) cacheGet(key string) string {
	return c[key]
}

// Get a book from the map (cache)
func (c Cache) GetBook(isbn string) string {
	return c.cacheGet(CacheKeyBook + isbn)
}

func (c Cache) GetCD(sku string) string {
	return c.cacheGet(CackeKeyCD + sku)
}

func (c Cache) SetCD(sku string, title string) {
	c.cacheSet(CackeKeyCD+sku, title)
}
