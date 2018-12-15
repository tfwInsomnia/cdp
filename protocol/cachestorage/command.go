// Code generated by cdpgen. DO NOT EDIT.

package cachestorage

// DeleteCacheArgs represents the arguments for DeleteCache in the CacheStorage domain.
type DeleteCacheArgs struct {
	CacheID CacheID `json:"cacheId"` // Id of cache for deletion.
}

// NewDeleteCacheArgs initializes DeleteCacheArgs with the required arguments.
func NewDeleteCacheArgs(cacheID CacheID) *DeleteCacheArgs {
	args := new(DeleteCacheArgs)
	args.CacheID = cacheID
	return args
}

// DeleteEntryArgs represents the arguments for DeleteEntry in the CacheStorage domain.
type DeleteEntryArgs struct {
	CacheID CacheID `json:"cacheId"` // Id of cache where the entry will be deleted.
	Request string  `json:"request"` // URL spec of the request.
}

// NewDeleteEntryArgs initializes DeleteEntryArgs with the required arguments.
func NewDeleteEntryArgs(cacheID CacheID, request string) *DeleteEntryArgs {
	args := new(DeleteEntryArgs)
	args.CacheID = cacheID
	args.Request = request
	return args
}

// RequestCacheNamesArgs represents the arguments for RequestCacheNames in the CacheStorage domain.
type RequestCacheNamesArgs struct {
	SecurityOrigin string `json:"securityOrigin"` // Security origin.
}

// NewRequestCacheNamesArgs initializes RequestCacheNamesArgs with the required arguments.
func NewRequestCacheNamesArgs(securityOrigin string) *RequestCacheNamesArgs {
	args := new(RequestCacheNamesArgs)
	args.SecurityOrigin = securityOrigin
	return args
}

// RequestCacheNamesReply represents the return values for RequestCacheNames in the CacheStorage domain.
type RequestCacheNamesReply struct {
	Caches []Cache `json:"caches"` // Caches for the security origin.
}

// RequestCachedResponseArgs represents the arguments for RequestCachedResponse in the CacheStorage domain.
type RequestCachedResponseArgs struct {
	CacheID    CacheID `json:"cacheId"`    // Id of cache that contains the entry.
	RequestURL string  `json:"requestURL"` // URL spec of the request.
}

// NewRequestCachedResponseArgs initializes RequestCachedResponseArgs with the required arguments.
func NewRequestCachedResponseArgs(cacheID CacheID, requestURL string) *RequestCachedResponseArgs {
	args := new(RequestCachedResponseArgs)
	args.CacheID = cacheID
	args.RequestURL = requestURL
	return args
}

// RequestCachedResponseReply represents the return values for RequestCachedResponse in the CacheStorage domain.
type RequestCachedResponseReply struct {
	Response CachedResponse `json:"response"` // Response read from the cache.
}

// RequestEntriesArgs represents the arguments for RequestEntries in the CacheStorage domain.
type RequestEntriesArgs struct {
	CacheID    CacheID `json:"cacheId"`              // ID of cache to get entries from.
	SkipCount  int     `json:"skipCount"`            // Number of records to skip.
	PageSize   int     `json:"pageSize"`             // Number of records to fetch.
	PathFilter *string `json:"pathFilter,omitempty"` // If present, only return the entries containing this substring in the path
}

// NewRequestEntriesArgs initializes RequestEntriesArgs with the required arguments.
func NewRequestEntriesArgs(cacheID CacheID, skipCount int, pageSize int) *RequestEntriesArgs {
	args := new(RequestEntriesArgs)
	args.CacheID = cacheID
	args.SkipCount = skipCount
	args.PageSize = pageSize
	return args
}

// SetPathFilter sets the PathFilter optional argument. If present,
// only return the entries containing this substring in the path
func (a *RequestEntriesArgs) SetPathFilter(pathFilter string) *RequestEntriesArgs {
	a.PathFilter = &pathFilter
	return a
}

// RequestEntriesReply represents the return values for RequestEntries in the CacheStorage domain.
type RequestEntriesReply struct {
	CacheDataEntries []DataEntry `json:"cacheDataEntries"` // Array of object store data entries.
	HasMore          bool        `json:"hasMore"`          // If true, there are more entries to fetch in the given range.
}
