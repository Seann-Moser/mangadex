package mangadex

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	openapi_types "github.com/oapi-codegen/runtime/types"
	"io"
)

// CachedClientWithResponsesInterface wraps ClientWithResponsesInterface adding caching layer.
type CachedClientWithResponsesInterface struct {
	client ClientWithResponsesInterface
	cache  Cache
}

// NewCachedClientWithResponsesInterface constructs a new CachedClientWithResponsesInterface.
func NewCachedClientWithResponsesInterface(client ClientWithResponsesInterface, cache Cache) *CachedClientWithResponsesInterface {
	return &CachedClientWithResponsesInterface{client: client, cache: cache}
}

// generateKey creates a cache key from the method name and argument values.
func generateKey(method string, args ...interface{}) string {
	b, err := json.Marshal(args)
	if err != nil {
		// Handle error: perhaps log it and return a non-cacheable key or panic
		return fmt.Sprintf("%s:json_marshal_error:%s", method, err.Error())
	}

	hasher := sha256.New()
	hasher.Write(b)
	hash := hex.EncodeToString(hasher.Sum(nil))

	return fmt.Sprintf("%s:%s", method, hash)
}

// GetAtHomeServerChapterIdWithResponse applies caching before delegating to the underlying client.
func (c *CachedClientWithResponsesInterface) GetAtHomeServerChapterIdWithResponse(ctx context.Context, chapterId openapi_types.UUID, params *GetAtHomeServerChapterIdParams, reqEditors ...RequestEditorFn) (*GetAtHomeServerChapterIdResponse, error) {
	// Build cache key
	key := generateKey("GetAtHomeServerChapterIdWithResponse", ctx, chapterId, params, reqEditors)
	if v, ok := c.cache.Get(key); ok {
		output := GetAtHomeServerChapterIdResponse{}
		err := json.Unmarshal(v, &output)
		if err == nil {
			return &output, nil
		}
	}

	// Cache miss: delegate to underlying client
	resp, err := c.client.GetAtHomeServerChapterIdWithResponse(ctx, chapterId, params, reqEditors...)
	if err != nil {
		return resp, err
	}
	c.cache.Set(key, resp)
	return resp, nil
}

// GetAuthCheckWithResponse applies caching before delegating to the underlying client.
func (c *CachedClientWithResponsesInterface) GetAuthCheckWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetAuthCheckResponse, error) {
	// Build cache key
	key := generateKey("GetAuthCheckWithResponse", ctx, reqEditors)
	if v, ok := c.cache.Get(key); ok {
		output := GetAuthCheckResponse{}
		err := json.Unmarshal(v, &output)
		if err == nil {
			return &output, nil
		}
	}

	// Cache miss: delegate to underlying client
	resp, err := c.client.GetAuthCheckWithResponse(ctx, reqEditors...)
	if err != nil {
		return resp, err
	}
	c.cache.Set(key, resp)
	return resp, nil
}

// PostAuthLoginWithBodyWithResponse applies caching before delegating to the underlying client.
func (c *CachedClientWithResponsesInterface) PostAuthLoginWithBodyWithResponse(ctx context.Context, params *PostAuthLoginParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostAuthLoginResponse, error) {
	// Build cache key
	key := generateKey("PostAuthLoginWithBodyWithResponse", ctx, params, contentType, body, reqEditors)
	if v, ok := c.cache.Get(key); ok {
		output := PostAuthLoginResponse{}
		err := json.Unmarshal(v, &output)
		if err == nil {
			return &output, nil
		}
	}

	// Cache miss: delegate to underlying client
	resp, err := c.client.PostAuthLoginWithBodyWithResponse(ctx, params, contentType, body, reqEditors...)
	if err != nil {
		return resp, err
	}
	c.cache.Set(key, resp)
	return resp, nil
}

// PostAuthLoginWithResponse applies caching before delegating to the underlying client.
func (c *CachedClientWithResponsesInterface) PostAuthLoginWithResponse(ctx context.Context, params *PostAuthLoginParams, body PostAuthLoginJSONRequestBody, reqEditors ...RequestEditorFn) (*PostAuthLoginResponse, error) {
	// Build cache key
	key := generateKey("PostAuthLoginWithResponse", ctx, params, body, reqEditors)
	if v, ok := c.cache.Get(key); ok {
		output := PostAuthLoginResponse{}
		err := json.Unmarshal(v, &output)
		if err == nil {
			return &output, nil
		}
	}

	// Cache miss: delegate to underlying client
	resp, err := c.client.PostAuthLoginWithResponse(ctx, params, body, reqEditors...)
	if err != nil {
		return resp, err
	}
	c.cache.Set(key, resp)
	return resp, nil
}

// PostAuthLogoutWithResponse applies caching before delegating to the underlying client.
func (c *CachedClientWithResponsesInterface) PostAuthLogoutWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*PostAuthLogoutResponse, error) {
	// Build cache key
	key := generateKey("PostAuthLogoutWithResponse", ctx, reqEditors)
	if v, ok := c.cache.Get(key); ok {
		output := PostAuthLogoutResponse{}
		err := json.Unmarshal(v, &output)
		if err == nil {
			return &output, nil
		}
	}

	// Cache miss: delegate to underlying client
	resp, err := c.client.PostAuthLogoutWithResponse(ctx, reqEditors...)
	if err != nil {
		return resp, err
	}
	c.cache.Set(key, resp)
	return resp, nil
}

// PostAuthRefreshWithBodyWithResponse applies caching before delegating to the underlying client.
func (c *CachedClientWithResponsesInterface) PostAuthRefreshWithBodyWithResponse(ctx context.Context, params *PostAuthRefreshParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostAuthRefreshResponse, error) {
	// Build cache key
	key := generateKey("PostAuthRefreshWithBodyWithResponse", ctx, params, contentType, body, reqEditors)
	if v, ok := c.cache.Get(key); ok {
		output := PostAuthRefreshResponse{}
		err := json.Unmarshal(v, &output)
		if err == nil {
			return &output, nil
		}
	}

	// Cache miss: delegate to underlying client
	resp, err := c.client.PostAuthRefreshWithBodyWithResponse(ctx, params, contentType, body, reqEditors...)
	if err != nil {
		return resp, err
	}
	c.cache.Set(key, resp)
	return resp, nil
}

// PostAuthRefreshWithResponse applies caching before delegating to the underlying client.
func (c *CachedClientWithResponsesInterface) PostAuthRefreshWithResponse(ctx context.Context, params *PostAuthRefreshParams, body PostAuthRefreshJSONRequestBody, reqEditors ...RequestEditorFn) (*PostAuthRefreshResponse, error) {
	// Build cache key
	key := generateKey("PostAuthRefreshWithResponse", ctx, params, body, reqEditors)
	if v, ok := c.cache.Get(key); ok {
		output := PostAuthRefreshResponse{}
		err := json.Unmarshal(v, &output)
		if err == nil {
			return &output, nil
		}
	}

	// Cache miss: delegate to underlying client
	resp, err := c.client.PostAuthRefreshWithResponse(ctx, params, body, reqEditors...)
	if err != nil {
		return resp, err
	}
	c.cache.Set(key, resp)
	return resp, nil
}

// GetAuthorWithResponse applies caching before delegating to the underlying client.
func (c *CachedClientWithResponsesInterface) GetAuthorWithResponse(ctx context.Context, params *GetAuthorParams, reqEditors ...RequestEditorFn) (*GetAuthorResponse, error) {
	// Build cache key
	key := generateKey("GetAuthorWithResponse", ctx, params, reqEditors)
	if v, ok := c.cache.Get(key); ok {
		output := GetAuthorResponse{}
		err := json.Unmarshal(v, &output)
		if err == nil {
			return &output, nil
		}
	}

	// Cache miss: delegate to underlying client
	resp, err := c.client.GetAuthorWithResponse(ctx, params, reqEditors...)
	if err != nil {
		return resp, err
	}
	c.cache.Set(key, resp)
	return resp, nil
}

// PostAuthorWithBodyWithResponse applies caching before delegating to the underlying client.
func (c *CachedClientWithResponsesInterface) PostAuthorWithBodyWithResponse(ctx context.Context, params *PostAuthorParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostAuthorResponse, error) {
	// Build cache key
	key := generateKey("PostAuthorWithBodyWithResponse", ctx, params, contentType, body, reqEditors)
	if v, ok := c.cache.Get(key); ok {
		output := PostAuthorResponse{}
		err := json.Unmarshal(v, &output)
		if err == nil {
			return &output, nil
		}
	}

	// Cache miss: delegate to underlying client
	resp, err := c.client.PostAuthorWithBodyWithResponse(ctx, params, contentType, body, reqEditors...)
	if err != nil {
		return resp, err
	}
	c.cache.Set(key, resp)
	return resp, nil
}

// PostAuthorWithResponse applies caching before delegating to the underlying client.
func (c *CachedClientWithResponsesInterface) PostAuthorWithResponse(ctx context.Context, params *PostAuthorParams, body PostAuthorJSONRequestBody, reqEditors ...RequestEditorFn) (*PostAuthorResponse, error) {
	// Build cache key
	key := generateKey("PostAuthorWithResponse", ctx, params, body, reqEditors)
	if v, ok := c.cache.Get(key); ok {
		output := PostAuthorResponse{}
		err := json.Unmarshal(v, &output)
		if err == nil {
			return &output, nil
		}
	}

	// Cache miss: delegate to underlying client
	resp, err := c.client.PostAuthorWithResponse(ctx, params, body, reqEditors...)
	if err != nil {
		return resp, err
	}
	c.cache.Set(key, resp)
	return resp, nil
}

// DeleteAuthorIdWithResponse applies caching before delegating to the underlying client.
func (c *CachedClientWithResponsesInterface) DeleteAuthorIdWithResponse(ctx context.Context, id openapi_types.UUID, reqEditors ...RequestEditorFn) (*DeleteAuthorIdResponse, error) {
	// Build cache key
	key := generateKey("DeleteAuthorIdWithResponse", ctx, id, reqEditors)
	if v, ok := c.cache.Get(key); ok {
		output := DeleteAuthorIdResponse{}
		err := json.Unmarshal(v, &output)
		if err == nil {
			return &output, nil
		}
	}

	// Cache miss: delegate to underlying client
	resp, err := c.client.DeleteAuthorIdWithResponse(ctx, id, reqEditors...)
	if err != nil {
		return resp, err
	}
	c.cache.Set(key, resp)
	return resp, nil
}

// GetAuthorIdWithResponse applies caching before delegating to the underlying client.
func (c *CachedClientWithResponsesInterface) GetAuthorIdWithResponse(ctx context.Context, id openapi_types.UUID, params *GetAuthorIdParams, reqEditors ...RequestEditorFn) (*GetAuthorIdResponse, error) {
	// Build cache key
	key := generateKey("GetAuthorIdWithResponse", ctx, id, params, reqEditors)
	if v, ok := c.cache.Get(key); ok {
		output := GetAuthorIdResponse{}
		err := json.Unmarshal(v, &output)
		if err == nil {
			return &output, nil
		}
	}

	// Cache miss: delegate to underlying client
	resp, err := c.client.GetAuthorIdWithResponse(ctx, id, params, reqEditors...)
	if err != nil {
		return resp, err
	}
	c.cache.Set(key, resp)
	return resp, nil
}

// PutAuthorIdWithBodyWithResponse applies caching before delegating to the underlying client.
func (c *CachedClientWithResponsesInterface) PutAuthorIdWithBodyWithResponse(ctx context.Context, id openapi_types.UUID, params *PutAuthorIdParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutAuthorIdResponse, error) {
	// Build cache key
	key := generateKey("PutAuthorIdWithBodyWithResponse", ctx, id, params, contentType, body, reqEditors)
	if v, ok := c.cache.Get(key); ok {
		output := PutAuthorIdResponse{}
		err := json.Unmarshal(v, &output)
		if err == nil {
			return &output, nil
		}
	}

	// Cache miss: delegate to underlying client
	resp, err := c.client.PutAuthorIdWithBodyWithResponse(ctx, id, params, contentType, body, reqEditors...)
	if err != nil {
		return resp, err
	}
	c.cache.Set(key, resp)
	return resp, nil
}

// PutAuthorIdWithResponse applies caching before delegating to the underlying client.
func (c *CachedClientWithResponsesInterface) PutAuthorIdWithResponse(ctx context.Context, id openapi_types.UUID, params *PutAuthorIdParams, body PutAuthorIdJSONRequestBody, reqEditors ...RequestEditorFn) (*PutAuthorIdResponse, error) {
	// Build cache key
	key := generateKey("PutAuthorIdWithResponse", ctx, id, params, body, reqEditors)
	if v, ok := c.cache.Get(key); ok {
		output := PutAuthorIdResponse{}
		err := json.Unmarshal(v, &output)
		if err == nil {
			return &output, nil
		}
	}

	// Cache miss: delegate to underlying client
	resp, err := c.client.PutAuthorIdWithResponse(ctx, id, params, body, reqEditors...)
	if err != nil {
		return resp, err
	}
	c.cache.Set(key, resp)
	return resp, nil
}

// PostCaptchaSolveWithBodyWithResponse applies caching before delegating to the underlying client.
func (c *CachedClientWithResponsesInterface) PostCaptchaSolveWithBodyWithResponse(ctx context.Context, params *PostCaptchaSolveParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostCaptchaSolveResponse, error) {
	// Build cache key
	key := generateKey("PostCaptchaSolveWithBodyWithResponse", ctx, params, contentType, body, reqEditors)
	if v, ok := c.cache.Get(key); ok {
		output := PostCaptchaSolveResponse{}
		err := json.Unmarshal(v, &output)
		if err == nil {
			return &output, nil
		}
	}

	// Cache miss: delegate to underlying client
	resp, err := c.client.PostCaptchaSolveWithBodyWithResponse(ctx, params, contentType, body, reqEditors...)
	if err != nil {
		return resp, err
	}
	c.cache.Set(key, resp)
	return resp, nil
}

// PostCaptchaSolveWithResponse applies caching before delegating to the underlying client.
func (c *CachedClientWithResponsesInterface) PostCaptchaSolveWithResponse(ctx context.Context, params *PostCaptchaSolveParams, body PostCaptchaSolveJSONRequestBody, reqEditors ...RequestEditorFn) (*PostCaptchaSolveResponse, error) {
	// Build cache key
	key := generateKey("PostCaptchaSolveWithResponse", ctx, params, body, reqEditors)
	if v, ok := c.cache.Get(key); ok {
		output := PostCaptchaSolveResponse{}
		err := json.Unmarshal(v, &output)
		if err == nil {
			return &output, nil
		}
	}

	// Cache miss: delegate to underlying client
	resp, err := c.client.PostCaptchaSolveWithResponse(ctx, params, body, reqEditors...)
	if err != nil {
		return resp, err
	}
	c.cache.Set(key, resp)
	return resp, nil
}

// GetChapterWithResponse applies caching before delegating to the underlying client.
func (c *CachedClientWithResponsesInterface) GetChapterWithResponse(ctx context.Context, params *GetChapterParams, reqEditors ...RequestEditorFn) (*GetChapterResponse, error) {
	// Build cache key
	key := generateKey("GetChapterWithResponse", ctx, params, reqEditors)
	if v, ok := c.cache.Get(key); ok {
		output := GetChapterResponse{}
		err := json.Unmarshal(v, &output)
		if err == nil {
			return &output, nil
		}
	}

	// Cache miss: delegate to underlying client
	resp, err := c.client.GetChapterWithResponse(ctx, params, reqEditors...)
	if err != nil {
		return resp, err
	}
	c.cache.Set(key, resp)
	return resp, nil
}

// DeleteChapterIdWithResponse applies caching before delegating to the underlying client.
func (c *CachedClientWithResponsesInterface) DeleteChapterIdWithResponse(ctx context.Context, id openapi_types.UUID, reqEditors ...RequestEditorFn) (*DeleteChapterIdResponse, error) {
	// Build cache key
	key := generateKey("DeleteChapterIdWithResponse", ctx, id, reqEditors)
	if v, ok := c.cache.Get(key); ok {
		output := DeleteChapterIdResponse{}
		err := json.Unmarshal(v, &output)
		if err == nil {
			return &output, nil
		}
	}

	// Cache miss: delegate to underlying client
	resp, err := c.client.DeleteChapterIdWithResponse(ctx, id, reqEditors...)
	if err != nil {
		return resp, err
	}
	c.cache.Set(key, resp)
	return resp, nil
}

// GetChapterIdWithResponse applies caching before delegating to the underlying client.
func (c *CachedClientWithResponsesInterface) GetChapterIdWithResponse(ctx context.Context, id openapi_types.UUID, params *GetChapterIdParams, reqEditors ...RequestEditorFn) (*GetChapterIdResponse, error) {
	// Build cache key
	key := generateKey("GetChapterIdWithResponse", ctx, id, params, reqEditors)
	if v, ok := c.cache.Get(key); ok {
		output := GetChapterIdResponse{}
		err := json.Unmarshal(v, &output)
		if err == nil {
			return &output, nil
		}
	}

	// Cache miss: delegate to underlying client
	resp, err := c.client.GetChapterIdWithResponse(ctx, id, params, reqEditors...)
	if err != nil {
		return resp, err
	}
	c.cache.Set(key, resp)
	return resp, nil
}

// PutChapterIdWithBodyWithResponse applies caching before delegating to the underlying client.
func (c *CachedClientWithResponsesInterface) PutChapterIdWithBodyWithResponse(ctx context.Context, id openapi_types.UUID, params *PutChapterIdParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutChapterIdResponse, error) {
	// Build cache key
	key := generateKey("PutChapterIdWithBodyWithResponse", ctx, id, params, contentType, body, reqEditors)
	if v, ok := c.cache.Get(key); ok {
		output := PutChapterIdResponse{}
		err := json.Unmarshal(v, &output)
		if err == nil {
			return &output, nil
		}
	}

	// Cache miss: delegate to underlying client
	resp, err := c.client.PutChapterIdWithBodyWithResponse(ctx, id, params, contentType, body, reqEditors...)
	if err != nil {
		return resp, err
	}
	c.cache.Set(key, resp)
	return resp, nil
}

// PutChapterIdWithResponse applies caching before delegating to the underlying client.
func (c *CachedClientWithResponsesInterface) PutChapterIdWithResponse(ctx context.Context, id openapi_types.UUID, params *PutChapterIdParams, body PutChapterIdJSONRequestBody, reqEditors ...RequestEditorFn) (*PutChapterIdResponse, error) {
	// Build cache key
	key := generateKey("PutChapterIdWithResponse", ctx, id, params, body, reqEditors)
	if v, ok := c.cache.Get(key); ok {
		output := PutChapterIdResponse{}
		err := json.Unmarshal(v, &output)
		if err == nil {
			return &output, nil
		}
	}

	// Cache miss: delegate to underlying client
	resp, err := c.client.PutChapterIdWithResponse(ctx, id, params, body, reqEditors...)
	if err != nil {
		return resp, err
	}
	c.cache.Set(key, resp)
	return resp, nil
}

// GetListApiclientsWithResponse applies caching before delegating to the underlying client.
func (c *CachedClientWithResponsesInterface) GetListApiclientsWithResponse(ctx context.Context, params *GetListApiclientsParams, reqEditors ...RequestEditorFn) (*GetListApiclientsResponse, error) {
	// Build cache key
	key := generateKey("GetListApiclientsWithResponse", ctx, params, reqEditors)
	if v, ok := c.cache.Get(key); ok {
		output := GetListApiclientsResponse{}
		err := json.Unmarshal(v, &output)
		if err == nil {
			return &output, nil
		}
	}

	// Cache miss: delegate to underlying client
	resp, err := c.client.GetListApiclientsWithResponse(ctx, params, reqEditors...)
	if err != nil {
		return resp, err
	}
	c.cache.Set(key, resp)
	return resp, nil
}

// PostCreateApiclientWithBodyWithResponse applies caching before delegating to the underlying client.
func (c *CachedClientWithResponsesInterface) PostCreateApiclientWithBodyWithResponse(ctx context.Context, params *PostCreateApiclientParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostCreateApiclientResponse, error) {
	// Build cache key
	key := generateKey("PostCreateApiclientWithBodyWithResponse", ctx, params, contentType, body, reqEditors)
	if v, ok := c.cache.Get(key); ok {
		output := PostCreateApiclientResponse{}
		err := json.Unmarshal(v, &output)
		if err == nil {
			return &output, nil
		}
	}

	// Cache miss: delegate to underlying client
	resp, err := c.client.PostCreateApiclientWithBodyWithResponse(ctx, params, contentType, body, reqEditors...)
	if err != nil {
		return resp, err
	}
	c.cache.Set(key, resp)
	return resp, nil
}

// PostCreateApiclientWithResponse applies caching before delegating to the underlying client.
func (c *CachedClientWithResponsesInterface) PostCreateApiclientWithResponse(ctx context.Context, params *PostCreateApiclientParams, body PostCreateApiclientJSONRequestBody, reqEditors ...RequestEditorFn) (*PostCreateApiclientResponse, error) {
	// Build cache key
	key := generateKey("PostCreateApiclientWithResponse", ctx, params, body, reqEditors)
	if v, ok := c.cache.Get(key); ok {
		output := PostCreateApiclientResponse{}
		err := json.Unmarshal(v, &output)
		if err == nil {
			return &output, nil
		}
	}

	// Cache miss: delegate to underlying client
	resp, err := c.client.PostCreateApiclientWithResponse(ctx, params, body, reqEditors...)
	if err != nil {
		return resp, err
	}
	c.cache.Set(key, resp)
	return resp, nil
}

// DeleteApiclientWithResponse applies caching before delegating to the underlying client.
func (c *CachedClientWithResponsesInterface) DeleteApiclientWithResponse(ctx context.Context, id openapi_types.UUID, params *DeleteApiclientParams, reqEditors ...RequestEditorFn) (*DeleteApiclientResponse, error) {
	// Build cache key
	key := generateKey("DeleteApiclientWithResponse", ctx, id, params, reqEditors)
	if v, ok := c.cache.Get(key); ok {
		output := DeleteApiclientResponse{}
		err := json.Unmarshal(v, &output)
		if err == nil {
			return &output, nil
		}
	}

	// Cache miss: delegate to underlying client
	resp, err := c.client.DeleteApiclientWithResponse(ctx, id, params, reqEditors...)
	if err != nil {
		return resp, err
	}
	c.cache.Set(key, resp)
	return resp, nil
}

// GetApiclientWithResponse applies caching before delegating to the underlying client.
func (c *CachedClientWithResponsesInterface) GetApiclientWithResponse(ctx context.Context, id openapi_types.UUID, params *GetApiclientParams, reqEditors ...RequestEditorFn) (*GetApiclientResponse, error) {
	// Build cache key
	key := generateKey("GetApiclientWithResponse", ctx, id, params, reqEditors)
	if v, ok := c.cache.Get(key); ok {
		output := GetApiclientResponse{}
		err := json.Unmarshal(v, &output)
		if err == nil {
			return &output, nil
		}
	}

	// Cache miss: delegate to underlying client
	resp, err := c.client.GetApiclientWithResponse(ctx, id, params, reqEditors...)
	if err != nil {
		return resp, err
	}
	c.cache.Set(key, resp)
	return resp, nil
}

// PostEditApiclientWithBodyWithResponse applies caching before delegating to the underlying client.
func (c *CachedClientWithResponsesInterface) PostEditApiclientWithBodyWithResponse(ctx context.Context, id openapi_types.UUID, params *PostEditApiclientParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostEditApiclientResponse, error) {
	// Build cache key
	key := generateKey("PostEditApiclientWithBodyWithResponse", ctx, id, params, contentType, body, reqEditors)
	if v, ok := c.cache.Get(key); ok {
		output := PostEditApiclientResponse{}
		err := json.Unmarshal(v, &output)
		if err == nil {
			return &output, nil
		}
	}

	// Cache miss: delegate to underlying client
	resp, err := c.client.PostEditApiclientWithBodyWithResponse(ctx, id, params, contentType, body, reqEditors...)
	if err != nil {
		return resp, err
	}
	c.cache.Set(key, resp)
	return resp, nil
}

// PostEditApiclientWithResponse applies caching before delegating to the underlying client.
func (c *CachedClientWithResponsesInterface) PostEditApiclientWithResponse(ctx context.Context, id openapi_types.UUID, params *PostEditApiclientParams, body PostEditApiclientJSONRequestBody, reqEditors ...RequestEditorFn) (*PostEditApiclientResponse, error) {
	// Build cache key
	key := generateKey("PostEditApiclientWithResponse", ctx, id, params, body, reqEditors)
	if v, ok := c.cache.Get(key); ok {
		output := PostEditApiclientResponse{}
		err := json.Unmarshal(v, &output)
		if err == nil {
			return &output, nil
		}
	}

	// Cache miss: delegate to underlying client
	resp, err := c.client.PostEditApiclientWithResponse(ctx, id, params, body, reqEditors...)
	if err != nil {
		return resp, err
	}
	c.cache.Set(key, resp)
	return resp, nil
}

// GetApiclientSecretWithResponse applies caching before delegating to the underlying client.
func (c *CachedClientWithResponsesInterface) GetApiclientSecretWithResponse(ctx context.Context, id openapi_types.UUID, reqEditors ...RequestEditorFn) (*GetApiclientSecretResponse, error) {
	// Build cache key
	key := generateKey("GetApiclientSecretWithResponse", ctx, id, reqEditors)
	if v, ok := c.cache.Get(key); ok {
		output := GetApiclientSecretResponse{}
		err := json.Unmarshal(v, &output)
		if err == nil {
			return &output, nil
		}
	}

	// Cache miss: delegate to underlying client
	resp, err := c.client.GetApiclientSecretWithResponse(ctx, id, reqEditors...)
	if err != nil {
		return resp, err
	}
	c.cache.Set(key, resp)
	return resp, nil
}

// PostRegenerateApiclientSecretWithBodyWithResponse applies caching before delegating to the underlying client.
func (c *CachedClientWithResponsesInterface) PostRegenerateApiclientSecretWithBodyWithResponse(ctx context.Context, id openapi_types.UUID, params *PostRegenerateApiclientSecretParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostRegenerateApiclientSecretResponse, error) {
	// Build cache key
	key := generateKey("PostRegenerateApiclientSecretWithBodyWithResponse", ctx, id, params, contentType, body, reqEditors)
	if v, ok := c.cache.Get(key); ok {
		output := PostRegenerateApiclientSecretResponse{}
		err := json.Unmarshal(v, &output)
		if err == nil {
			return &output, nil
		}
	}

	// Cache miss: delegate to underlying client
	resp, err := c.client.PostRegenerateApiclientSecretWithBodyWithResponse(ctx, id, params, contentType, body, reqEditors...)
	if err != nil {
		return resp, err
	}
	c.cache.Set(key, resp)
	return resp, nil
}

// PostRegenerateApiclientSecretWithResponse applies caching before delegating to the underlying client.
func (c *CachedClientWithResponsesInterface) PostRegenerateApiclientSecretWithResponse(ctx context.Context, id openapi_types.UUID, params *PostRegenerateApiclientSecretParams, body PostRegenerateApiclientSecretJSONRequestBody, reqEditors ...RequestEditorFn) (*PostRegenerateApiclientSecretResponse, error) {
	// Build cache key
	key := generateKey("PostRegenerateApiclientSecretWithResponse", ctx, id, params, body, reqEditors)
	if v, ok := c.cache.Get(key); ok {
		output := PostRegenerateApiclientSecretResponse{}
		err := json.Unmarshal(v, &output)
		if err == nil {
			return &output, nil
		}
	}

	// Cache miss: delegate to underlying client
	resp, err := c.client.PostRegenerateApiclientSecretWithResponse(ctx, id, params, body, reqEditors...)
	if err != nil {
		return resp, err
	}
	c.cache.Set(key, resp)
	return resp, nil
}

// GetCoverWithResponse applies caching before delegating to the underlying client.
func (c *CachedClientWithResponsesInterface) GetCoverWithResponse(ctx context.Context, params *GetCoverParams, reqEditors ...RequestEditorFn) (*GetCoverResponse, error) {
	// Build cache key
	key := generateKey("GetCoverWithResponse", ctx, params, reqEditors)
	if v, ok := c.cache.Get(key); ok {
		output := GetCoverResponse{}
		err := json.Unmarshal(v, &output)
		if err == nil {
			return &output, nil
		}
	}

	// Cache miss: delegate to underlying client
	resp, err := c.client.GetCoverWithResponse(ctx, params, reqEditors...)
	if err != nil {
		return resp, err
	}
	c.cache.Set(key, resp)
	return resp, nil
}

// DeleteCoverWithResponse applies caching before delegating to the underlying client.
func (c *CachedClientWithResponsesInterface) DeleteCoverWithResponse(ctx context.Context, mangaOrCoverId openapi_types.UUID, reqEditors ...RequestEditorFn) (*DeleteCoverResponse, error) {
	// Build cache key
	key := generateKey("DeleteCoverWithResponse", ctx, mangaOrCoverId, reqEditors)
	if v, ok := c.cache.Get(key); ok {
		output := DeleteCoverResponse{}
		err := json.Unmarshal(v, &output)
		if err == nil {
			return &output, nil
		}
	}

	// Cache miss: delegate to underlying client
	resp, err := c.client.DeleteCoverWithResponse(ctx, mangaOrCoverId, reqEditors...)
	if err != nil {
		return resp, err
	}
	c.cache.Set(key, resp)
	return resp, nil
}

// GetCoverIdWithResponse applies caching before delegating to the underlying client.
func (c *CachedClientWithResponsesInterface) GetCoverIdWithResponse(ctx context.Context, mangaOrCoverId openapi_types.UUID, params *GetCoverIdParams, reqEditors ...RequestEditorFn) (*GetCoverIdResponse, error) {
	// Build cache key
	key := generateKey("GetCoverIdWithResponse", ctx, mangaOrCoverId, params, reqEditors)
	if v, ok := c.cache.Get(key); ok {
		output := GetCoverIdResponse{}
		err := json.Unmarshal(v, &output)
		if err == nil {
			return &output, nil
		}
	}

	// Cache miss: delegate to underlying client
	resp, err := c.client.GetCoverIdWithResponse(ctx, mangaOrCoverId, params, reqEditors...)
	if err != nil {
		return resp, err
	}
	c.cache.Set(key, resp)
	return resp, nil
}

// UploadCoverWithBodyWithResponse applies caching before delegating to the underlying client.
func (c *CachedClientWithResponsesInterface) UploadCoverWithBodyWithResponse(ctx context.Context, mangaOrCoverId openapi_types.UUID, params *UploadCoverParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*UploadCoverResponse, error) {
	// Build cache key
	key := generateKey("UploadCoverWithBodyWithResponse", ctx, mangaOrCoverId, params, contentType, body, reqEditors)
	if v, ok := c.cache.Get(key); ok {
		output := UploadCoverResponse{}
		err := json.Unmarshal(v, &output)
		if err == nil {
			return &output, nil
		}
	}

	// Cache miss: delegate to underlying client
	resp, err := c.client.UploadCoverWithBodyWithResponse(ctx, mangaOrCoverId, params, contentType, body, reqEditors...)
	if err != nil {
		return resp, err
	}
	c.cache.Set(key, resp)
	return resp, nil
}

// EditCoverWithBodyWithResponse applies caching before delegating to the underlying client.
func (c *CachedClientWithResponsesInterface) EditCoverWithBodyWithResponse(ctx context.Context, mangaOrCoverId openapi_types.UUID, params *EditCoverParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*EditCoverResponse, error) {
	// Build cache key
	key := generateKey("EditCoverWithBodyWithResponse", ctx, mangaOrCoverId, params, contentType, body, reqEditors)
	if v, ok := c.cache.Get(key); ok {
		output := EditCoverResponse{}
		err := json.Unmarshal(v, &output)
		if err == nil {
			return &output, nil
		}
	}

	// Cache miss: delegate to underlying client
	resp, err := c.client.EditCoverWithBodyWithResponse(ctx, mangaOrCoverId, params, contentType, body, reqEditors...)
	if err != nil {
		return resp, err
	}
	c.cache.Set(key, resp)
	return resp, nil
}

// EditCoverWithResponse applies caching before delegating to the underlying client.
func (c *CachedClientWithResponsesInterface) EditCoverWithResponse(ctx context.Context, mangaOrCoverId openapi_types.UUID, params *EditCoverParams, body EditCoverJSONRequestBody, reqEditors ...RequestEditorFn) (*EditCoverResponse, error) {
	// Build cache key
	key := generateKey("EditCoverWithResponse", ctx, mangaOrCoverId, params, body, reqEditors)
	if v, ok := c.cache.Get(key); ok {
		output := EditCoverResponse{}
		err := json.Unmarshal(v, &output)
		if err == nil {
			return &output, nil
		}
	}

	// Cache miss: delegate to underlying client
	resp, err := c.client.EditCoverWithResponse(ctx, mangaOrCoverId, params, body, reqEditors...)
	if err != nil {
		return resp, err
	}
	c.cache.Set(key, resp)
	return resp, nil
}

// ForumsThreadCreateWithBodyWithResponse applies caching before delegating to the underlying client.
func (c *CachedClientWithResponsesInterface) ForumsThreadCreateWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*ForumsThreadCreateResponse, error) {
	// Build cache key
	key := generateKey("ForumsThreadCreateWithBodyWithResponse", ctx, contentType, body, reqEditors)
	if v, ok := c.cache.Get(key); ok {
		output := ForumsThreadCreateResponse{}
		err := json.Unmarshal(v, &output)
		if err == nil {
			return &output, nil
		}
	}

	// Cache miss: delegate to underlying client
	resp, err := c.client.ForumsThreadCreateWithBodyWithResponse(ctx, contentType, body, reqEditors...)
	if err != nil {
		return resp, err
	}
	c.cache.Set(key, resp)
	return resp, nil
}

// ForumsThreadCreateWithResponse applies caching before delegating to the underlying client.
func (c *CachedClientWithResponsesInterface) ForumsThreadCreateWithResponse(ctx context.Context, body ForumsThreadCreateJSONRequestBody, reqEditors ...RequestEditorFn) (*ForumsThreadCreateResponse, error) {
	// Build cache key
	key := generateKey("ForumsThreadCreateWithResponse", ctx, body, reqEditors)
	if v, ok := c.cache.Get(key); ok {
		output := ForumsThreadCreateResponse{}
		err := json.Unmarshal(v, &output)
		if err == nil {
			return &output, nil
		}
	}

	// Cache miss: delegate to underlying client
	resp, err := c.client.ForumsThreadCreateWithResponse(ctx, body, reqEditors...)
	if err != nil {
		return resp, err
	}
	c.cache.Set(key, resp)
	return resp, nil
}

// GetSearchGroupWithResponse applies caching before delegating to the underlying client.
func (c *CachedClientWithResponsesInterface) GetSearchGroupWithResponse(ctx context.Context, params *GetSearchGroupParams, reqEditors ...RequestEditorFn) (*GetSearchGroupResponse, error) {
	// Build cache key
	key := generateKey("GetSearchGroupWithResponse", ctx, params, reqEditors)
	if v, ok := c.cache.Get(key); ok {
		output := GetSearchGroupResponse{}
		err := json.Unmarshal(v, &output)
		if err == nil {
			return &output, nil
		}
	}

	// Cache miss: delegate to underlying client
	resp, err := c.client.GetSearchGroupWithResponse(ctx, params, reqEditors...)
	if err != nil {
		return resp, err
	}
	c.cache.Set(key, resp)
	return resp, nil
}

// PostGroupWithBodyWithResponse applies caching before delegating to the underlying client.
func (c *CachedClientWithResponsesInterface) PostGroupWithBodyWithResponse(ctx context.Context, params *PostGroupParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostGroupResponse, error) {
	// Build cache key
	key := generateKey("PostGroupWithBodyWithResponse", ctx, params, contentType, body, reqEditors)
	if v, ok := c.cache.Get(key); ok {
		output := PostGroupResponse{}
		err := json.Unmarshal(v, &output)
		if err == nil {
			return &output, nil
		}
	}

	// Cache miss: delegate to underlying client
	resp, err := c.client.PostGroupWithBodyWithResponse(ctx, params, contentType, body, reqEditors...)
	if err != nil {
		return resp, err
	}
	c.cache.Set(key, resp)
	return resp, nil
}

// PostGroupWithResponse applies caching before delegating to the underlying client.
func (c *CachedClientWithResponsesInterface) PostGroupWithResponse(ctx context.Context, params *PostGroupParams, body PostGroupJSONRequestBody, reqEditors ...RequestEditorFn) (*PostGroupResponse, error) {
	// Build cache key
	key := generateKey("PostGroupWithResponse", ctx, params, body, reqEditors)
	if v, ok := c.cache.Get(key); ok {
		output := PostGroupResponse{}
		err := json.Unmarshal(v, &output)
		if err == nil {
			return &output, nil
		}
	}

	// Cache miss: delegate to underlying client
	resp, err := c.client.PostGroupWithResponse(ctx, params, body, reqEditors...)
	if err != nil {
		return resp, err
	}
	c.cache.Set(key, resp)
	return resp, nil
}

// DeleteGroupIdWithResponse applies caching before delegating to the underlying client.
func (c *CachedClientWithResponsesInterface) DeleteGroupIdWithResponse(ctx context.Context, id openapi_types.UUID, reqEditors ...RequestEditorFn) (*DeleteGroupIdResponse, error) {
	// Build cache key
	key := generateKey("DeleteGroupIdWithResponse", ctx, id, reqEditors)
	if v, ok := c.cache.Get(key); ok {
		output := DeleteGroupIdResponse{}
		err := json.Unmarshal(v, &output)
		if err == nil {
			return &output, nil
		}
	}

	// Cache miss: delegate to underlying client
	resp, err := c.client.DeleteGroupIdWithResponse(ctx, id, reqEditors...)
	if err != nil {
		return resp, err
	}
	c.cache.Set(key, resp)
	return resp, nil
}

// GetGroupIdWithResponse applies caching before delegating to the underlying client.
func (c *CachedClientWithResponsesInterface) GetGroupIdWithResponse(ctx context.Context, id openapi_types.UUID, params *GetGroupIdParams, reqEditors ...RequestEditorFn) (*GetGroupIdResponse, error) {
	// Build cache key
	key := generateKey("GetGroupIdWithResponse", ctx, id, params, reqEditors)
	if v, ok := c.cache.Get(key); ok {
		output := GetGroupIdResponse{}
		err := json.Unmarshal(v, &output)
		if err == nil {
			return &output, nil
		}
	}

	// Cache miss: delegate to underlying client
	resp, err := c.client.GetGroupIdWithResponse(ctx, id, params, reqEditors...)
	if err != nil {
		return resp, err
	}
	c.cache.Set(key, resp)
	return resp, nil
}

// PutGroupIdWithBodyWithResponse applies caching before delegating to the underlying client.
func (c *CachedClientWithResponsesInterface) PutGroupIdWithBodyWithResponse(ctx context.Context, id openapi_types.UUID, params *PutGroupIdParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutGroupIdResponse, error) {
	// Build cache key
	key := generateKey("PutGroupIdWithBodyWithResponse", ctx, id, params, contentType, body, reqEditors)
	if v, ok := c.cache.Get(key); ok {
		output := PutGroupIdResponse{}
		err := json.Unmarshal(v, &output)
		if err == nil {
			return &output, nil
		}
	}

	// Cache miss: delegate to underlying client
	resp, err := c.client.PutGroupIdWithBodyWithResponse(ctx, id, params, contentType, body, reqEditors...)
	if err != nil {
		return resp, err
	}
	c.cache.Set(key, resp)
	return resp, nil
}

// PutGroupIdWithResponse applies caching before delegating to the underlying client.
func (c *CachedClientWithResponsesInterface) PutGroupIdWithResponse(ctx context.Context, id openapi_types.UUID, params *PutGroupIdParams, body PutGroupIdJSONRequestBody, reqEditors ...RequestEditorFn) (*PutGroupIdResponse, error) {
	// Build cache key
	key := generateKey("PutGroupIdWithResponse", ctx, id, params, body, reqEditors)
	if v, ok := c.cache.Get(key); ok {
		output := PutGroupIdResponse{}
		err := json.Unmarshal(v, &output)
		if err == nil {
			return &output, nil
		}
	}

	// Cache miss: delegate to underlying client
	resp, err := c.client.PutGroupIdWithResponse(ctx, id, params, body, reqEditors...)
	if err != nil {
		return resp, err
	}
	c.cache.Set(key, resp)
	return resp, nil
}

// DeleteGroupIdFollowWithResponse applies caching before delegating to the underlying client.
func (c *CachedClientWithResponsesInterface) DeleteGroupIdFollowWithResponse(ctx context.Context, id openapi_types.UUID, reqEditors ...RequestEditorFn) (*DeleteGroupIdFollowResponse, error) {
	// Build cache key
	key := generateKey("DeleteGroupIdFollowWithResponse", ctx, id, reqEditors)
	if v, ok := c.cache.Get(key); ok {
		output := DeleteGroupIdFollowResponse{}
		err := json.Unmarshal(v, &output)
		if err == nil {
			return &output, nil
		}
	}

	// Cache miss: delegate to underlying client
	resp, err := c.client.DeleteGroupIdFollowWithResponse(ctx, id, reqEditors...)
	if err != nil {
		return resp, err
	}
	c.cache.Set(key, resp)
	return resp, nil
}

// PostGroupIdFollowWithResponse applies caching before delegating to the underlying client.
func (c *CachedClientWithResponsesInterface) PostGroupIdFollowWithResponse(ctx context.Context, id openapi_types.UUID, reqEditors ...RequestEditorFn) (*PostGroupIdFollowResponse, error) {
	// Build cache key
	key := generateKey("PostGroupIdFollowWithResponse", ctx, id, reqEditors)
	if v, ok := c.cache.Get(key); ok {
		output := PostGroupIdFollowResponse{}
		err := json.Unmarshal(v, &output)
		if err == nil {
			return &output, nil
		}
	}

	// Cache miss: delegate to underlying client
	resp, err := c.client.PostGroupIdFollowWithResponse(ctx, id, reqEditors...)
	if err != nil {
		return resp, err
	}
	c.cache.Set(key, resp)
	return resp, nil
}

// PostLegacyMappingWithBodyWithResponse applies caching before delegating to the underlying client.
func (c *CachedClientWithResponsesInterface) PostLegacyMappingWithBodyWithResponse(ctx context.Context, params *PostLegacyMappingParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostLegacyMappingResponse, error) {
	// Build cache key
	key := generateKey("PostLegacyMappingWithBodyWithResponse", ctx, params, contentType, body, reqEditors)
	if v, ok := c.cache.Get(key); ok {
		output := PostLegacyMappingResponse{}
		err := json.Unmarshal(v, &output)
		if err == nil {
			return &output, nil
		}
	}

	// Cache miss: delegate to underlying client
	resp, err := c.client.PostLegacyMappingWithBodyWithResponse(ctx, params, contentType, body, reqEditors...)
	if err != nil {
		return resp, err
	}
	c.cache.Set(key, resp)
	return resp, nil
}

// PostLegacyMappingWithResponse applies caching before delegating to the underlying client.
func (c *CachedClientWithResponsesInterface) PostLegacyMappingWithResponse(ctx context.Context, params *PostLegacyMappingParams, body PostLegacyMappingJSONRequestBody, reqEditors ...RequestEditorFn) (*PostLegacyMappingResponse, error) {
	// Build cache key
	key := generateKey("PostLegacyMappingWithResponse", ctx, params, body, reqEditors)
	if v, ok := c.cache.Get(key); ok {
		output := PostLegacyMappingResponse{}
		err := json.Unmarshal(v, &output)
		if err == nil {
			return &output, nil
		}
	}

	// Cache miss: delegate to underlying client
	resp, err := c.client.PostLegacyMappingWithResponse(ctx, params, body, reqEditors...)
	if err != nil {
		return resp, err
	}
	c.cache.Set(key, resp)
	return resp, nil
}

// PostListWithBodyWithResponse applies caching before delegating to the underlying client.
func (c *CachedClientWithResponsesInterface) PostListWithBodyWithResponse(ctx context.Context, params *PostListParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostListResponse, error) {
	// Build cache key
	key := generateKey("PostListWithBodyWithResponse", ctx, params, contentType, body, reqEditors)
	if v, ok := c.cache.Get(key); ok {
		output := PostListResponse{}
		err := json.Unmarshal(v, &output)
		if err == nil {
			return &output, nil
		}
	}

	// Cache miss: delegate to underlying client
	resp, err := c.client.PostListWithBodyWithResponse(ctx, params, contentType, body, reqEditors...)
	if err != nil {
		return resp, err
	}
	c.cache.Set(key, resp)
	return resp, nil
}

// PostListWithResponse applies caching before delegating to the underlying client.
func (c *CachedClientWithResponsesInterface) PostListWithResponse(ctx context.Context, params *PostListParams, body PostListJSONRequestBody, reqEditors ...RequestEditorFn) (*PostListResponse, error) {
	// Build cache key
	key := generateKey("PostListWithResponse", ctx, params, body, reqEditors)
	if v, ok := c.cache.Get(key); ok {
		output := PostListResponse{}
		err := json.Unmarshal(v, &output)
		if err == nil {
			return &output, nil
		}
	}

	// Cache miss: delegate to underlying client
	resp, err := c.client.PostListWithResponse(ctx, params, body, reqEditors...)
	if err != nil {
		return resp, err
	}
	c.cache.Set(key, resp)
	return resp, nil
}

// DeleteListIdWithResponse applies caching before delegating to the underlying client.
func (c *CachedClientWithResponsesInterface) DeleteListIdWithResponse(ctx context.Context, id openapi_types.UUID, reqEditors ...RequestEditorFn) (*DeleteListIdResponse, error) {
	// Build cache key
	key := generateKey("DeleteListIdWithResponse", ctx, id, reqEditors)
	if v, ok := c.cache.Get(key); ok {
		output := DeleteListIdResponse{}
		err := json.Unmarshal(v, &output)
		if err == nil {
			return &output, nil
		}
	}

	// Cache miss: delegate to underlying client
	resp, err := c.client.DeleteListIdWithResponse(ctx, id, reqEditors...)
	if err != nil {
		return resp, err
	}
	c.cache.Set(key, resp)
	return resp, nil
}

// GetListIdWithResponse applies caching before delegating to the underlying client.
func (c *CachedClientWithResponsesInterface) GetListIdWithResponse(ctx context.Context, id openapi_types.UUID, reqEditors ...RequestEditorFn) (*GetListIdResponse, error) {
	// Build cache key
	key := generateKey("GetListIdWithResponse", ctx, id, reqEditors)
	if v, ok := c.cache.Get(key); ok {
		output := GetListIdResponse{}
		err := json.Unmarshal(v, &output)
		if err == nil {
			return &output, nil
		}
	}

	// Cache miss: delegate to underlying client
	resp, err := c.client.GetListIdWithResponse(ctx, id, reqEditors...)
	if err != nil {
		return resp, err
	}
	c.cache.Set(key, resp)
	return resp, nil
}

// PutListIdWithBodyWithResponse applies caching before delegating to the underlying client.
func (c *CachedClientWithResponsesInterface) PutListIdWithBodyWithResponse(ctx context.Context, id openapi_types.UUID, params *PutListIdParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutListIdResponse, error) {
	// Build cache key
	key := generateKey("PutListIdWithBodyWithResponse", ctx, id, params, contentType, body, reqEditors)
	if v, ok := c.cache.Get(key); ok {
		output := PutListIdResponse{}
		err := json.Unmarshal(v, &output)
		if err == nil {
			return &output, nil
		}
	}

	// Cache miss: delegate to underlying client
	resp, err := c.client.PutListIdWithBodyWithResponse(ctx, id, params, contentType, body, reqEditors...)
	if err != nil {
		return resp, err
	}
	c.cache.Set(key, resp)
	return resp, nil
}

// PutListIdWithResponse applies caching before delegating to the underlying client.
func (c *CachedClientWithResponsesInterface) PutListIdWithResponse(ctx context.Context, id openapi_types.UUID, params *PutListIdParams, body PutListIdJSONRequestBody, reqEditors ...RequestEditorFn) (*PutListIdResponse, error) {
	// Build cache key
	key := generateKey("PutListIdWithResponse", ctx, id, params, body, reqEditors)
	if v, ok := c.cache.Get(key); ok {
		output := PutListIdResponse{}
		err := json.Unmarshal(v, &output)
		if err == nil {
			return &output, nil
		}
	}

	// Cache miss: delegate to underlying client
	resp, err := c.client.PutListIdWithResponse(ctx, id, params, body, reqEditors...)
	if err != nil {
		return resp, err
	}
	c.cache.Set(key, resp)
	return resp, nil
}

// GetListIdFeedWithResponse applies caching before delegating to the underlying client.
func (c *CachedClientWithResponsesInterface) GetListIdFeedWithResponse(ctx context.Context, id openapi_types.UUID, params *GetListIdFeedParams, reqEditors ...RequestEditorFn) (*GetListIdFeedResponse, error) {
	// Build cache key
	key := generateKey("GetListIdFeedWithResponse", ctx, id, params, reqEditors)
	if v, ok := c.cache.Get(key); ok {
		output := GetListIdFeedResponse{}
		err := json.Unmarshal(v, &output)
		if err == nil {
			return &output, nil
		}
	}

	// Cache miss: delegate to underlying client
	resp, err := c.client.GetListIdFeedWithResponse(ctx, id, params, reqEditors...)
	if err != nil {
		return resp, err
	}
	c.cache.Set(key, resp)
	return resp, nil
}

// UnfollowListIdWithBodyWithResponse applies caching before delegating to the underlying client.
func (c *CachedClientWithResponsesInterface) UnfollowListIdWithBodyWithResponse(ctx context.Context, id openapi_types.UUID, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*UnfollowListIdResponse, error) {
	// Build cache key
	key := generateKey("UnfollowListIdWithBodyWithResponse", ctx, id, contentType, body, reqEditors)
	if v, ok := c.cache.Get(key); ok {
		output := UnfollowListIdResponse{}
		err := json.Unmarshal(v, &output)
		if err == nil {
			return &output, nil
		}
	}

	// Cache miss: delegate to underlying client
	resp, err := c.client.UnfollowListIdWithBodyWithResponse(ctx, id, contentType, body, reqEditors...)
	if err != nil {
		return resp, err
	}
	c.cache.Set(key, resp)
	return resp, nil
}

// UnfollowListIdWithResponse applies caching before delegating to the underlying client.
func (c *CachedClientWithResponsesInterface) UnfollowListIdWithResponse(ctx context.Context, id openapi_types.UUID, body UnfollowListIdJSONRequestBody, reqEditors ...RequestEditorFn) (*UnfollowListIdResponse, error) {
	// Build cache key
	key := generateKey("UnfollowListIdWithResponse", ctx, id, body, reqEditors)
	if v, ok := c.cache.Get(key); ok {
		output := UnfollowListIdResponse{}
		err := json.Unmarshal(v, &output)
		if err == nil {
			return &output, nil
		}
	}

	// Cache miss: delegate to underlying client
	resp, err := c.client.UnfollowListIdWithResponse(ctx, id, body, reqEditors...)
	if err != nil {
		return resp, err
	}
	c.cache.Set(key, resp)
	return resp, nil
}

// FollowListIdWithBodyWithResponse applies caching before delegating to the underlying client.
func (c *CachedClientWithResponsesInterface) FollowListIdWithBodyWithResponse(ctx context.Context, id openapi_types.UUID, params *FollowListIdParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*FollowListIdResponse, error) {
	// Build cache key
	key := generateKey("FollowListIdWithBodyWithResponse", ctx, id, params, contentType, body, reqEditors)
	if v, ok := c.cache.Get(key); ok {
		output := FollowListIdResponse{}
		err := json.Unmarshal(v, &output)
		if err == nil {
			return &output, nil
		}
	}

	// Cache miss: delegate to underlying client
	resp, err := c.client.FollowListIdWithBodyWithResponse(ctx, id, params, contentType, body, reqEditors...)
	if err != nil {
		return resp, err
	}
	c.cache.Set(key, resp)
	return resp, nil
}

// FollowListIdWithResponse applies caching before delegating to the underlying client.
func (c *CachedClientWithResponsesInterface) FollowListIdWithResponse(ctx context.Context, id openapi_types.UUID, params *FollowListIdParams, body FollowListIdJSONRequestBody, reqEditors ...RequestEditorFn) (*FollowListIdResponse, error) {
	// Build cache key
	key := generateKey("FollowListIdWithResponse", ctx, id, params, body, reqEditors)
	if v, ok := c.cache.Get(key); ok {
		output := FollowListIdResponse{}
		err := json.Unmarshal(v, &output)
		if err == nil {
			return &output, nil
		}
	}

	// Cache miss: delegate to underlying client
	resp, err := c.client.FollowListIdWithResponse(ctx, id, params, body, reqEditors...)
	if err != nil {
		return resp, err
	}
	c.cache.Set(key, resp)
	return resp, nil
}

// GetSearchMangaWithResponse applies caching before delegating to the underlying client.
func (c *CachedClientWithResponsesInterface) GetSearchMangaWithResponse(ctx context.Context, params *GetSearchMangaParams, reqEditors ...RequestEditorFn) (*GetSearchMangaResponse, error) {
	// Build cache key
	key := generateKey("GetSearchMangaWithResponse", ctx, params, reqEditors)
	if v, ok := c.cache.Get(key); ok {
		output := GetSearchMangaResponse{}
		err := json.Unmarshal(v, &output)
		if err == nil {
			return &output, nil
		}
	}

	// Cache miss: delegate to underlying client
	resp, err := c.client.GetSearchMangaWithResponse(ctx, params, reqEditors...)
	if err != nil {
		return resp, err
	}
	c.cache.Set(key, resp)
	return resp, nil
}

// PostMangaWithBodyWithResponse applies caching before delegating to the underlying client.
func (c *CachedClientWithResponsesInterface) PostMangaWithBodyWithResponse(ctx context.Context, params *PostMangaParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostMangaResponse, error) {
	// Build cache key
	key := generateKey("PostMangaWithBodyWithResponse", ctx, params, contentType, body, reqEditors)
	if v, ok := c.cache.Get(key); ok {
		output := PostMangaResponse{}
		err := json.Unmarshal(v, &output)
		if err == nil {
			return &output, nil
		}
	}

	// Cache miss: delegate to underlying client
	resp, err := c.client.PostMangaWithBodyWithResponse(ctx, params, contentType, body, reqEditors...)
	if err != nil {
		return resp, err
	}
	c.cache.Set(key, resp)
	return resp, nil
}

// PostMangaWithResponse applies caching before delegating to the underlying client.
func (c *CachedClientWithResponsesInterface) PostMangaWithResponse(ctx context.Context, params *PostMangaParams, body PostMangaJSONRequestBody, reqEditors ...RequestEditorFn) (*PostMangaResponse, error) {
	// Build cache key
	key := generateKey("PostMangaWithResponse", ctx, params, body, reqEditors)
	if v, ok := c.cache.Get(key); ok {
		output := PostMangaResponse{}
		err := json.Unmarshal(v, &output)
		if err == nil {
			return &output, nil
		}
	}

	// Cache miss: delegate to underlying client
	resp, err := c.client.PostMangaWithResponse(ctx, params, body, reqEditors...)
	if err != nil {
		return resp, err
	}
	c.cache.Set(key, resp)
	return resp, nil
}

// GetMangaDraftsWithResponse applies caching before delegating to the underlying client.
func (c *CachedClientWithResponsesInterface) GetMangaDraftsWithResponse(ctx context.Context, params *GetMangaDraftsParams, reqEditors ...RequestEditorFn) (*GetMangaDraftsResponse, error) {
	// Build cache key
	key := generateKey("GetMangaDraftsWithResponse", ctx, params, reqEditors)
	if v, ok := c.cache.Get(key); ok {
		output := GetMangaDraftsResponse{}
		err := json.Unmarshal(v, &output)
		if err == nil {
			return &output, nil
		}
	}

	// Cache miss: delegate to underlying client
	resp, err := c.client.GetMangaDraftsWithResponse(ctx, params, reqEditors...)
	if err != nil {
		return resp, err
	}
	c.cache.Set(key, resp)
	return resp, nil
}

// GetMangaIdDraftWithResponse applies caching before delegating to the underlying client.
func (c *CachedClientWithResponsesInterface) GetMangaIdDraftWithResponse(ctx context.Context, id openapi_types.UUID, params *GetMangaIdDraftParams, reqEditors ...RequestEditorFn) (*GetMangaIdDraftResponse, error) {
	// Build cache key
	key := generateKey("GetMangaIdDraftWithResponse", ctx, id, params, reqEditors)
	if v, ok := c.cache.Get(key); ok {
		output := GetMangaIdDraftResponse{}
		err := json.Unmarshal(v, &output)
		if err == nil {
			return &output, nil
		}
	}

	// Cache miss: delegate to underlying client
	resp, err := c.client.GetMangaIdDraftWithResponse(ctx, id, params, reqEditors...)
	if err != nil {
		return resp, err
	}
	c.cache.Set(key, resp)
	return resp, nil
}

// CommitMangaDraftWithBodyWithResponse applies caching before delegating to the underlying client.
func (c *CachedClientWithResponsesInterface) CommitMangaDraftWithBodyWithResponse(ctx context.Context, id openapi_types.UUID, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*CommitMangaDraftResponse, error) {
	// Build cache key
	key := generateKey("CommitMangaDraftWithBodyWithResponse", ctx, id, contentType, body, reqEditors)
	if v, ok := c.cache.Get(key); ok {
		output := CommitMangaDraftResponse{}
		err := json.Unmarshal(v, &output)
		if err == nil {
			return &output, nil
		}
	}

	// Cache miss: delegate to underlying client
	resp, err := c.client.CommitMangaDraftWithBodyWithResponse(ctx, id, contentType, body, reqEditors...)
	if err != nil {
		return resp, err
	}
	c.cache.Set(key, resp)
	return resp, nil
}

// CommitMangaDraftWithResponse applies caching before delegating to the underlying client.
func (c *CachedClientWithResponsesInterface) CommitMangaDraftWithResponse(ctx context.Context, id openapi_types.UUID, body CommitMangaDraftJSONRequestBody, reqEditors ...RequestEditorFn) (*CommitMangaDraftResponse, error) {
	// Build cache key
	key := generateKey("CommitMangaDraftWithResponse", ctx, id, body, reqEditors)
	if v, ok := c.cache.Get(key); ok {
		output := CommitMangaDraftResponse{}
		err := json.Unmarshal(v, &output)
		if err == nil {
			return &output, nil
		}
	}

	// Cache miss: delegate to underlying client
	resp, err := c.client.CommitMangaDraftWithResponse(ctx, id, body, reqEditors...)
	if err != nil {
		return resp, err
	}
	c.cache.Set(key, resp)
	return resp, nil
}

// GetMangaRandomWithResponse applies caching before delegating to the underlying client.
func (c *CachedClientWithResponsesInterface) GetMangaRandomWithResponse(ctx context.Context, params *GetMangaRandomParams, reqEditors ...RequestEditorFn) (*GetMangaRandomResponse, error) {
	// Build cache key
	key := generateKey("GetMangaRandomWithResponse", ctx, params, reqEditors)
	if v, ok := c.cache.Get(key); ok {
		output := GetMangaRandomResponse{}
		err := json.Unmarshal(v, &output)
		if err == nil {
			return &output, nil
		}
	}

	// Cache miss: delegate to underlying client
	resp, err := c.client.GetMangaRandomWithResponse(ctx, params, reqEditors...)
	if err != nil {
		return resp, err
	}
	c.cache.Set(key, resp)
	return resp, nil
}

// GetMangaChapterReadmarkers2WithResponse applies caching before delegating to the underlying client.
func (c *CachedClientWithResponsesInterface) GetMangaChapterReadmarkers2WithResponse(ctx context.Context, params *GetMangaChapterReadmarkers2Params, reqEditors ...RequestEditorFn) (*GetMangaChapterReadmarkers2Response, error) {
	// Build cache key
	key := generateKey("GetMangaChapterReadmarkers2WithResponse", ctx, params, reqEditors)
	if v, ok := c.cache.Get(key); ok {
		output := GetMangaChapterReadmarkers2Response{}
		err := json.Unmarshal(v, &output)
		if err == nil {
			return &output, nil
		}
	}

	// Cache miss: delegate to underlying client
	resp, err := c.client.GetMangaChapterReadmarkers2WithResponse(ctx, params, reqEditors...)
	if err != nil {
		return resp, err
	}
	c.cache.Set(key, resp)
	return resp, nil
}

// GetMangaStatusWithResponse applies caching before delegating to the underlying client.
func (c *CachedClientWithResponsesInterface) GetMangaStatusWithResponse(ctx context.Context, params *GetMangaStatusParams, reqEditors ...RequestEditorFn) (*GetMangaStatusResponse, error) {
	// Build cache key
	key := generateKey("GetMangaStatusWithResponse", ctx, params, reqEditors)
	if v, ok := c.cache.Get(key); ok {
		output := GetMangaStatusResponse{}
		err := json.Unmarshal(v, &output)
		if err == nil {
			return &output, nil
		}
	}

	// Cache miss: delegate to underlying client
	resp, err := c.client.GetMangaStatusWithResponse(ctx, params, reqEditors...)
	if err != nil {
		return resp, err
	}
	c.cache.Set(key, resp)
	return resp, nil
}

// GetMangaTagWithResponse applies caching before delegating to the underlying client.
func (c *CachedClientWithResponsesInterface) GetMangaTagWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetMangaTagResponse, error) {
	// Build cache key
	key := generateKey("GetMangaTagWithResponse", ctx, reqEditors)
	if v, ok := c.cache.Get(key); ok {
		output := GetMangaTagResponse{}
		err := json.Unmarshal(v, &output)
		if err == nil {
			return &output, nil
		}
	}

	// Cache miss: delegate to underlying client
	resp, err := c.client.GetMangaTagWithResponse(ctx, reqEditors...)
	if err != nil {
		return resp, err
	}
	c.cache.Set(key, resp)
	return resp, nil
}

// DeleteMangaIdWithResponse applies caching before delegating to the underlying client.
func (c *CachedClientWithResponsesInterface) DeleteMangaIdWithResponse(ctx context.Context, id openapi_types.UUID, reqEditors ...RequestEditorFn) (*DeleteMangaIdResponse, error) {
	// Build cache key
	key := generateKey("DeleteMangaIdWithResponse", ctx, id, reqEditors)
	if v, ok := c.cache.Get(key); ok {
		output := DeleteMangaIdResponse{}
		err := json.Unmarshal(v, &output)
		if err == nil {
			return &output, nil
		}
	}

	// Cache miss: delegate to underlying client
	resp, err := c.client.DeleteMangaIdWithResponse(ctx, id, reqEditors...)
	if err != nil {
		return resp, err
	}
	c.cache.Set(key, resp)
	return resp, nil
}

// GetMangaIdWithResponse applies caching before delegating to the underlying client.
func (c *CachedClientWithResponsesInterface) GetMangaIdWithResponse(ctx context.Context, id openapi_types.UUID, params *GetMangaIdParams, reqEditors ...RequestEditorFn) (*GetMangaIdResponse, error) {
	// Build cache key
	key := generateKey("GetMangaIdWithResponse", ctx, id, params, reqEditors)
	if v, ok := c.cache.Get(key); ok {
		output := GetMangaIdResponse{}
		err := json.Unmarshal(v, &output)
		if err == nil {
			return &output, nil
		}
	}

	// Cache miss: delegate to underlying client
	resp, err := c.client.GetMangaIdWithResponse(ctx, id, params, reqEditors...)
	if err != nil {
		return resp, err
	}
	c.cache.Set(key, resp)
	return resp, nil
}

// PutMangaIdWithBodyWithResponse applies caching before delegating to the underlying client.
func (c *CachedClientWithResponsesInterface) PutMangaIdWithBodyWithResponse(ctx context.Context, id openapi_types.UUID, params *PutMangaIdParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutMangaIdResponse, error) {
	// Build cache key
	key := generateKey("PutMangaIdWithBodyWithResponse", ctx, id, params, contentType, body, reqEditors)
	if v, ok := c.cache.Get(key); ok {
		output := PutMangaIdResponse{}
		err := json.Unmarshal(v, &output)
		if err == nil {
			return &output, nil
		}
	}

	// Cache miss: delegate to underlying client
	resp, err := c.client.PutMangaIdWithBodyWithResponse(ctx, id, params, contentType, body, reqEditors...)
	if err != nil {
		return resp, err
	}
	c.cache.Set(key, resp)
	return resp, nil
}

// PutMangaIdWithResponse applies caching before delegating to the underlying client.
func (c *CachedClientWithResponsesInterface) PutMangaIdWithResponse(ctx context.Context, id openapi_types.UUID, params *PutMangaIdParams, body PutMangaIdJSONRequestBody, reqEditors ...RequestEditorFn) (*PutMangaIdResponse, error) {
	// Build cache key
	key := generateKey("PutMangaIdWithResponse", ctx, id, params, body, reqEditors)
	if v, ok := c.cache.Get(key); ok {
		output := PutMangaIdResponse{}
		err := json.Unmarshal(v, &output)
		if err == nil {
			return &output, nil
		}
	}

	// Cache miss: delegate to underlying client
	resp, err := c.client.PutMangaIdWithResponse(ctx, id, params, body, reqEditors...)
	if err != nil {
		return resp, err
	}
	c.cache.Set(key, resp)
	return resp, nil
}

// GetMangaAggregateWithResponse applies caching before delegating to the underlying client.
func (c *CachedClientWithResponsesInterface) GetMangaAggregateWithResponse(ctx context.Context, id openapi_types.UUID, params *GetMangaAggregateParams, reqEditors ...RequestEditorFn) (*GetMangaAggregateResponse, error) {
	// Build cache key
	key := generateKey("GetMangaAggregateWithResponse", ctx, id, params, reqEditors)
	if v, ok := c.cache.Get(key); ok {
		output := GetMangaAggregateResponse{}
		err := json.Unmarshal(v, &output)
		if err == nil {
			return &output, nil
		}
	}

	// Cache miss: delegate to underlying client
	resp, err := c.client.GetMangaAggregateWithResponse(ctx, id, params, reqEditors...)
	if err != nil {
		return resp, err
	}
	c.cache.Set(key, resp)
	return resp, nil
}

// GetMangaIdFeedWithResponse applies caching before delegating to the underlying client.
func (c *CachedClientWithResponsesInterface) GetMangaIdFeedWithResponse(ctx context.Context, id openapi_types.UUID, params *GetMangaIdFeedParams, reqEditors ...RequestEditorFn) (*GetMangaIdFeedResponse, error) {
	// Build cache key
	key := generateKey("GetMangaIdFeedWithResponse", ctx, id, params, reqEditors)
	if v, ok := c.cache.Get(key); ok {
		output := GetMangaIdFeedResponse{}
		err := json.Unmarshal(v, &output)
		if err == nil {
			return &output, nil
		}
	}

	// Cache miss: delegate to underlying client
	resp, err := c.client.GetMangaIdFeedWithResponse(ctx, id, params, reqEditors...)
	if err != nil {
		return resp, err
	}
	c.cache.Set(key, resp)
	return resp, nil
}

// DeleteMangaIdFollowWithResponse applies caching before delegating to the underlying client.
func (c *CachedClientWithResponsesInterface) DeleteMangaIdFollowWithResponse(ctx context.Context, id openapi_types.UUID, reqEditors ...RequestEditorFn) (*DeleteMangaIdFollowResponse, error) {
	// Build cache key
	key := generateKey("DeleteMangaIdFollowWithResponse", ctx, id, reqEditors)
	if v, ok := c.cache.Get(key); ok {
		output := DeleteMangaIdFollowResponse{}
		err := json.Unmarshal(v, &output)
		if err == nil {
			return &output, nil
		}
	}

	// Cache miss: delegate to underlying client
	resp, err := c.client.DeleteMangaIdFollowWithResponse(ctx, id, reqEditors...)
	if err != nil {
		return resp, err
	}
	c.cache.Set(key, resp)
	return resp, nil
}

// PostMangaIdFollowWithResponse applies caching before delegating to the underlying client.
func (c *CachedClientWithResponsesInterface) PostMangaIdFollowWithResponse(ctx context.Context, id openapi_types.UUID, reqEditors ...RequestEditorFn) (*PostMangaIdFollowResponse, error) {
	// Build cache key
	key := generateKey("PostMangaIdFollowWithResponse", ctx, id, reqEditors)
	if v, ok := c.cache.Get(key); ok {
		output := PostMangaIdFollowResponse{}
		err := json.Unmarshal(v, &output)
		if err == nil {
			return &output, nil
		}
	}

	// Cache miss: delegate to underlying client
	resp, err := c.client.PostMangaIdFollowWithResponse(ctx, id, reqEditors...)
	if err != nil {
		return resp, err
	}
	c.cache.Set(key, resp)
	return resp, nil
}

// DeleteMangaIdListListIdWithResponse applies caching before delegating to the underlying client.
func (c *CachedClientWithResponsesInterface) DeleteMangaIdListListIdWithResponse(ctx context.Context, id openapi_types.UUID, listId openapi_types.UUID, reqEditors ...RequestEditorFn) (*DeleteMangaIdListListIdResponse, error) {
	// Build cache key
	key := generateKey("DeleteMangaIdListListIdWithResponse", ctx, id, listId, reqEditors)
	if v, ok := c.cache.Get(key); ok {
		output := DeleteMangaIdListListIdResponse{}
		err := json.Unmarshal(v, &output)
		if err == nil {
			return &output, nil
		}
	}

	// Cache miss: delegate to underlying client
	resp, err := c.client.DeleteMangaIdListListIdWithResponse(ctx, id, listId, reqEditors...)
	if err != nil {
		return resp, err
	}
	c.cache.Set(key, resp)
	return resp, nil
}

// PostMangaIdListListIdWithResponse applies caching before delegating to the underlying client.
func (c *CachedClientWithResponsesInterface) PostMangaIdListListIdWithResponse(ctx context.Context, id openapi_types.UUID, listId openapi_types.UUID, reqEditors ...RequestEditorFn) (*PostMangaIdListListIdResponse, error) {
	// Build cache key
	key := generateKey("PostMangaIdListListIdWithResponse", ctx, id, listId, reqEditors)
	if v, ok := c.cache.Get(key); ok {
		output := PostMangaIdListListIdResponse{}
		err := json.Unmarshal(v, &output)
		if err == nil {
			return &output, nil
		}
	}

	// Cache miss: delegate to underlying client
	resp, err := c.client.PostMangaIdListListIdWithResponse(ctx, id, listId, reqEditors...)
	if err != nil {
		return resp, err
	}
	c.cache.Set(key, resp)
	return resp, nil
}

// GetMangaChapterReadmarkersWithResponse applies caching before delegating to the underlying client.
func (c *CachedClientWithResponsesInterface) GetMangaChapterReadmarkersWithResponse(ctx context.Context, id openapi_types.UUID, reqEditors ...RequestEditorFn) (*GetMangaChapterReadmarkersResponse, error) {
	// Build cache key
	key := generateKey("GetMangaChapterReadmarkersWithResponse", ctx, id, reqEditors)
	if v, ok := c.cache.Get(key); ok {
		output := GetMangaChapterReadmarkersResponse{}
		err := json.Unmarshal(v, &output)
		if err == nil {
			return &output, nil
		}
	}

	// Cache miss: delegate to underlying client
	resp, err := c.client.GetMangaChapterReadmarkersWithResponse(ctx, id, reqEditors...)
	if err != nil {
		return resp, err
	}
	c.cache.Set(key, resp)
	return resp, nil
}

// PostMangaChapterReadmarkersWithBodyWithResponse applies caching before delegating to the underlying client.
func (c *CachedClientWithResponsesInterface) PostMangaChapterReadmarkersWithBodyWithResponse(ctx context.Context, id openapi_types.UUID, params *PostMangaChapterReadmarkersParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostMangaChapterReadmarkersResponse, error) {
	// Build cache key
	key := generateKey("PostMangaChapterReadmarkersWithBodyWithResponse", ctx, id, params, contentType, body, reqEditors)
	if v, ok := c.cache.Get(key); ok {
		output := PostMangaChapterReadmarkersResponse{}
		err := json.Unmarshal(v, &output)
		if err == nil {
			return &output, nil
		}
	}

	// Cache miss: delegate to underlying client
	resp, err := c.client.PostMangaChapterReadmarkersWithBodyWithResponse(ctx, id, params, contentType, body, reqEditors...)
	if err != nil {
		return resp, err
	}
	c.cache.Set(key, resp)
	return resp, nil
}

// PostMangaChapterReadmarkersWithResponse applies caching before delegating to the underlying client.
func (c *CachedClientWithResponsesInterface) PostMangaChapterReadmarkersWithResponse(ctx context.Context, id openapi_types.UUID, params *PostMangaChapterReadmarkersParams, body PostMangaChapterReadmarkersJSONRequestBody, reqEditors ...RequestEditorFn) (*PostMangaChapterReadmarkersResponse, error) {
	// Build cache key
	key := generateKey("PostMangaChapterReadmarkersWithResponse", ctx, id, params, body, reqEditors)
	if v, ok := c.cache.Get(key); ok {
		output := PostMangaChapterReadmarkersResponse{}
		err := json.Unmarshal(v, &output)
		if err == nil {
			return &output, nil
		}
	}

	// Cache miss: delegate to underlying client
	resp, err := c.client.PostMangaChapterReadmarkersWithResponse(ctx, id, params, body, reqEditors...)
	if err != nil {
		return resp, err
	}
	c.cache.Set(key, resp)
	return resp, nil
}

// GetMangaIdStatusWithResponse applies caching before delegating to the underlying client.
func (c *CachedClientWithResponsesInterface) GetMangaIdStatusWithResponse(ctx context.Context, id openapi_types.UUID, reqEditors ...RequestEditorFn) (*GetMangaIdStatusResponse, error) {
	// Build cache key
	key := generateKey("GetMangaIdStatusWithResponse", ctx, id, reqEditors)
	if v, ok := c.cache.Get(key); ok {
		output := GetMangaIdStatusResponse{}
		err := json.Unmarshal(v, &output)
		if err == nil {
			return &output, nil
		}
	}

	// Cache miss: delegate to underlying client
	resp, err := c.client.GetMangaIdStatusWithResponse(ctx, id, reqEditors...)
	if err != nil {
		return resp, err
	}
	c.cache.Set(key, resp)
	return resp, nil
}

// PostMangaIdStatusWithBodyWithResponse applies caching before delegating to the underlying client.
func (c *CachedClientWithResponsesInterface) PostMangaIdStatusWithBodyWithResponse(ctx context.Context, id openapi_types.UUID, params *PostMangaIdStatusParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostMangaIdStatusResponse, error) {
	// Build cache key
	key := generateKey("PostMangaIdStatusWithBodyWithResponse", ctx, id, params, contentType, body, reqEditors)
	if v, ok := c.cache.Get(key); ok {
		output := PostMangaIdStatusResponse{}
		err := json.Unmarshal(v, &output)
		if err == nil {
			return &output, nil
		}
	}

	// Cache miss: delegate to underlying client
	resp, err := c.client.PostMangaIdStatusWithBodyWithResponse(ctx, id, params, contentType, body, reqEditors...)
	if err != nil {
		return resp, err
	}
	c.cache.Set(key, resp)
	return resp, nil
}

// PostMangaIdStatusWithResponse applies caching before delegating to the underlying client.
func (c *CachedClientWithResponsesInterface) PostMangaIdStatusWithResponse(ctx context.Context, id openapi_types.UUID, params *PostMangaIdStatusParams, body PostMangaIdStatusJSONRequestBody, reqEditors ...RequestEditorFn) (*PostMangaIdStatusResponse, error) {
	// Build cache key
	key := generateKey("PostMangaIdStatusWithResponse", ctx, id, params, body, reqEditors)
	if v, ok := c.cache.Get(key); ok {
		output := PostMangaIdStatusResponse{}
		err := json.Unmarshal(v, &output)
		if err == nil {
			return &output, nil
		}
	}

	// Cache miss: delegate to underlying client
	resp, err := c.client.PostMangaIdStatusWithResponse(ctx, id, params, body, reqEditors...)
	if err != nil {
		return resp, err
	}
	c.cache.Set(key, resp)
	return resp, nil
}

// GetMangaRelationWithResponse applies caching before delegating to the underlying client.
func (c *CachedClientWithResponsesInterface) GetMangaRelationWithResponse(ctx context.Context, mangaId openapi_types.UUID, params *GetMangaRelationParams, reqEditors ...RequestEditorFn) (*GetMangaRelationResponse, error) {
	// Build cache key
	key := generateKey("GetMangaRelationWithResponse", ctx, mangaId, params, reqEditors)
	if v, ok := c.cache.Get(key); ok {
		output := GetMangaRelationResponse{}
		err := json.Unmarshal(v, &output)
		if err == nil {
			return &output, nil
		}
	}

	// Cache miss: delegate to underlying client
	resp, err := c.client.GetMangaRelationWithResponse(ctx, mangaId, params, reqEditors...)
	if err != nil {
		return resp, err
	}
	c.cache.Set(key, resp)
	return resp, nil
}

// PostMangaRelationWithBodyWithResponse applies caching before delegating to the underlying client.
func (c *CachedClientWithResponsesInterface) PostMangaRelationWithBodyWithResponse(ctx context.Context, mangaId openapi_types.UUID, params *PostMangaRelationParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostMangaRelationResponse, error) {
	// Build cache key
	key := generateKey("PostMangaRelationWithBodyWithResponse", ctx, mangaId, params, contentType, body, reqEditors)
	if v, ok := c.cache.Get(key); ok {
		output := PostMangaRelationResponse{}
		err := json.Unmarshal(v, &output)
		if err == nil {
			return &output, nil
		}
	}

	// Cache miss: delegate to underlying client
	resp, err := c.client.PostMangaRelationWithBodyWithResponse(ctx, mangaId, params, contentType, body, reqEditors...)
	if err != nil {
		return resp, err
	}
	c.cache.Set(key, resp)
	return resp, nil
}

// PostMangaRelationWithResponse applies caching before delegating to the underlying client.
func (c *CachedClientWithResponsesInterface) PostMangaRelationWithResponse(ctx context.Context, mangaId openapi_types.UUID, params *PostMangaRelationParams, body PostMangaRelationJSONRequestBody, reqEditors ...RequestEditorFn) (*PostMangaRelationResponse, error) {
	// Build cache key
	key := generateKey("PostMangaRelationWithResponse", ctx, mangaId, params, body, reqEditors)
	if v, ok := c.cache.Get(key); ok {
		output := PostMangaRelationResponse{}
		err := json.Unmarshal(v, &output)
		if err == nil {
			return &output, nil
		}
	}

	// Cache miss: delegate to underlying client
	resp, err := c.client.PostMangaRelationWithResponse(ctx, mangaId, params, body, reqEditors...)
	if err != nil {
		return resp, err
	}
	c.cache.Set(key, resp)
	return resp, nil
}

// DeleteMangaRelationIdWithResponse applies caching before delegating to the underlying client.
func (c *CachedClientWithResponsesInterface) DeleteMangaRelationIdWithResponse(ctx context.Context, mangaId openapi_types.UUID, id openapi_types.UUID, reqEditors ...RequestEditorFn) (*DeleteMangaRelationIdResponse, error) {
	// Build cache key
	key := generateKey("DeleteMangaRelationIdWithResponse", ctx, mangaId, id, reqEditors)
	if v, ok := c.cache.Get(key); ok {
		output := DeleteMangaRelationIdResponse{}
		err := json.Unmarshal(v, &output)
		if err == nil {
			return &output, nil
		}
	}

	// Cache miss: delegate to underlying client
	resp, err := c.client.DeleteMangaRelationIdWithResponse(ctx, mangaId, id, reqEditors...)
	if err != nil {
		return resp, err
	}
	c.cache.Set(key, resp)
	return resp, nil
}

// GetPingWithResponse applies caching before delegating to the underlying client.
func (c *CachedClientWithResponsesInterface) GetPingWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetPingResponse, error) {
	// Build cache key
	key := generateKey("GetPingWithResponse", ctx, reqEditors)
	if v, ok := c.cache.Get(key); ok {
		output := GetPingResponse{}
		err := json.Unmarshal(v, &output)
		if err == nil {
			return &output, nil
		}
	}

	// Cache miss: delegate to underlying client
	resp, err := c.client.GetPingWithResponse(ctx, reqEditors...)
	if err != nil {
		return resp, err
	}
	c.cache.Set(key, resp)
	return resp, nil
}

// GetRatingWithResponse applies caching before delegating to the underlying client.
func (c *CachedClientWithResponsesInterface) GetRatingWithResponse(ctx context.Context, params *GetRatingParams, reqEditors ...RequestEditorFn) (*GetRatingResponse, error) {
	// Build cache key
	key := generateKey("GetRatingWithResponse", ctx, params, reqEditors)
	if v, ok := c.cache.Get(key); ok {
		output := GetRatingResponse{}
		err := json.Unmarshal(v, &output)
		if err == nil {
			return &output, nil
		}
	}

	// Cache miss: delegate to underlying client
	resp, err := c.client.GetRatingWithResponse(ctx, params, reqEditors...)
	if err != nil {
		return resp, err
	}
	c.cache.Set(key, resp)
	return resp, nil
}

// DeleteRatingMangaIdWithResponse applies caching before delegating to the underlying client.
func (c *CachedClientWithResponsesInterface) DeleteRatingMangaIdWithResponse(ctx context.Context, mangaId openapi_types.UUID, reqEditors ...RequestEditorFn) (*DeleteRatingMangaIdResponse, error) {
	// Build cache key
	key := generateKey("DeleteRatingMangaIdWithResponse", ctx, mangaId, reqEditors)
	if v, ok := c.cache.Get(key); ok {
		output := DeleteRatingMangaIdResponse{}
		err := json.Unmarshal(v, &output)
		if err == nil {
			return &output, nil
		}
	}

	// Cache miss: delegate to underlying client
	resp, err := c.client.DeleteRatingMangaIdWithResponse(ctx, mangaId, reqEditors...)
	if err != nil {
		return resp, err
	}
	c.cache.Set(key, resp)
	return resp, nil
}

// PostRatingMangaIdWithBodyWithResponse applies caching before delegating to the underlying client.
func (c *CachedClientWithResponsesInterface) PostRatingMangaIdWithBodyWithResponse(ctx context.Context, mangaId openapi_types.UUID, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostRatingMangaIdResponse, error) {
	// Build cache key
	key := generateKey("PostRatingMangaIdWithBodyWithResponse", ctx, mangaId, contentType, body, reqEditors)
	if v, ok := c.cache.Get(key); ok {
		output := PostRatingMangaIdResponse{}
		err := json.Unmarshal(v, &output)
		if err == nil {
			return &output, nil
		}
	}

	// Cache miss: delegate to underlying client
	resp, err := c.client.PostRatingMangaIdWithBodyWithResponse(ctx, mangaId, contentType, body, reqEditors...)
	if err != nil {
		return resp, err
	}
	c.cache.Set(key, resp)
	return resp, nil
}

// PostRatingMangaIdWithResponse applies caching before delegating to the underlying client.
func (c *CachedClientWithResponsesInterface) PostRatingMangaIdWithResponse(ctx context.Context, mangaId openapi_types.UUID, body PostRatingMangaIdJSONRequestBody, reqEditors ...RequestEditorFn) (*PostRatingMangaIdResponse, error) {
	// Build cache key
	key := generateKey("PostRatingMangaIdWithResponse", ctx, mangaId, body, reqEditors)
	if v, ok := c.cache.Get(key); ok {
		output := PostRatingMangaIdResponse{}
		err := json.Unmarshal(v, &output)
		if err == nil {
			return &output, nil
		}
	}

	// Cache miss: delegate to underlying client
	resp, err := c.client.PostRatingMangaIdWithResponse(ctx, mangaId, body, reqEditors...)
	if err != nil {
		return resp, err
	}
	c.cache.Set(key, resp)
	return resp, nil
}

// GetReportsWithResponse applies caching before delegating to the underlying client.
func (c *CachedClientWithResponsesInterface) GetReportsWithResponse(ctx context.Context, params *GetReportsParams, reqEditors ...RequestEditorFn) (*GetReportsResponse, error) {
	// Build cache key
	key := generateKey("GetReportsWithResponse", ctx, params, reqEditors)
	if v, ok := c.cache.Get(key); ok {
		output := GetReportsResponse{}
		err := json.Unmarshal(v, &output)
		if err == nil {
			return &output, nil
		}
	}

	// Cache miss: delegate to underlying client
	resp, err := c.client.GetReportsWithResponse(ctx, params, reqEditors...)
	if err != nil {
		return resp, err
	}
	c.cache.Set(key, resp)
	return resp, nil
}

// PostReportWithBodyWithResponse applies caching before delegating to the underlying client.
func (c *CachedClientWithResponsesInterface) PostReportWithBodyWithResponse(ctx context.Context, params *PostReportParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostReportResponse, error) {
	// Build cache key
	key := generateKey("PostReportWithBodyWithResponse", ctx, params, contentType, body, reqEditors)
	if v, ok := c.cache.Get(key); ok {
		output := PostReportResponse{}
		err := json.Unmarshal(v, &output)
		if err == nil {
			return &output, nil
		}
	}

	// Cache miss: delegate to underlying client
	resp, err := c.client.PostReportWithBodyWithResponse(ctx, params, contentType, body, reqEditors...)
	if err != nil {
		return resp, err
	}
	c.cache.Set(key, resp)
	return resp, nil
}

// PostReportWithResponse applies caching before delegating to the underlying client.
func (c *CachedClientWithResponsesInterface) PostReportWithResponse(ctx context.Context, params *PostReportParams, body PostReportJSONRequestBody, reqEditors ...RequestEditorFn) (*PostReportResponse, error) {
	// Build cache key
	key := generateKey("PostReportWithResponse", ctx, params, body, reqEditors)
	if v, ok := c.cache.Get(key); ok {
		output := PostReportResponse{}
		err := json.Unmarshal(v, &output)
		if err == nil {
			return &output, nil
		}
	}

	// Cache miss: delegate to underlying client
	resp, err := c.client.PostReportWithResponse(ctx, params, body, reqEditors...)
	if err != nil {
		return resp, err
	}
	c.cache.Set(key, resp)
	return resp, nil
}

// GetReportReasonsByCategoryWithResponse applies caching before delegating to the underlying client.
func (c *CachedClientWithResponsesInterface) GetReportReasonsByCategoryWithResponse(ctx context.Context, category string, reqEditors ...RequestEditorFn) (*GetReportReasonsByCategoryResponse, error) {
	// Build cache key
	key := generateKey("GetReportReasonsByCategoryWithResponse", ctx, category, reqEditors)
	if v, ok := c.cache.Get(key); ok {
		output := GetReportReasonsByCategoryResponse{}
		err := json.Unmarshal(v, &output)
		if err == nil {
			return &output, nil
		}
	}

	// Cache miss: delegate to underlying client
	resp, err := c.client.GetReportReasonsByCategoryWithResponse(ctx, category, reqEditors...)
	if err != nil {
		return resp, err
	}
	c.cache.Set(key, resp)
	return resp, nil
}

// GetSettingsWithResponse applies caching before delegating to the underlying client.
func (c *CachedClientWithResponsesInterface) GetSettingsWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetSettingsResponse, error) {
	// Build cache key
	key := generateKey("GetSettingsWithResponse", ctx, reqEditors)
	if v, ok := c.cache.Get(key); ok {
		output := GetSettingsResponse{}
		err := json.Unmarshal(v, &output)
		if err == nil {
			return &output, nil
		}
	}

	// Cache miss: delegate to underlying client
	resp, err := c.client.GetSettingsWithResponse(ctx, reqEditors...)
	if err != nil {
		return resp, err
	}
	c.cache.Set(key, resp)
	return resp, nil
}

// PostSettingsWithBodyWithResponse applies caching before delegating to the underlying client.
func (c *CachedClientWithResponsesInterface) PostSettingsWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostSettingsResponse, error) {
	// Build cache key
	key := generateKey("PostSettingsWithBodyWithResponse", ctx, contentType, body, reqEditors)
	if v, ok := c.cache.Get(key); ok {
		output := PostSettingsResponse{}
		err := json.Unmarshal(v, &output)
		if err == nil {
			return &output, nil
		}
	}

	// Cache miss: delegate to underlying client
	resp, err := c.client.PostSettingsWithBodyWithResponse(ctx, contentType, body, reqEditors...)
	if err != nil {
		return resp, err
	}
	c.cache.Set(key, resp)
	return resp, nil
}

// PostSettingsWithResponse applies caching before delegating to the underlying client.
func (c *CachedClientWithResponsesInterface) PostSettingsWithResponse(ctx context.Context, body PostSettingsJSONRequestBody, reqEditors ...RequestEditorFn) (*PostSettingsResponse, error) {
	// Build cache key
	key := generateKey("PostSettingsWithResponse", ctx, body, reqEditors)
	if v, ok := c.cache.Get(key); ok {
		output := PostSettingsResponse{}
		err := json.Unmarshal(v, &output)
		if err == nil {
			return &output, nil
		}
	}

	// Cache miss: delegate to underlying client
	resp, err := c.client.PostSettingsWithResponse(ctx, body, reqEditors...)
	if err != nil {
		return resp, err
	}
	c.cache.Set(key, resp)
	return resp, nil
}

// GetSettingsTemplateWithResponse applies caching before delegating to the underlying client.
func (c *CachedClientWithResponsesInterface) GetSettingsTemplateWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetSettingsTemplateResponse, error) {
	// Build cache key
	key := generateKey("GetSettingsTemplateWithResponse", ctx, reqEditors)
	if v, ok := c.cache.Get(key); ok {
		output := GetSettingsTemplateResponse{}
		err := json.Unmarshal(v, &output)
		if err == nil {
			return &output, nil
		}
	}

	// Cache miss: delegate to underlying client
	resp, err := c.client.GetSettingsTemplateWithResponse(ctx, reqEditors...)
	if err != nil {
		return resp, err
	}
	c.cache.Set(key, resp)
	return resp, nil
}

// PostSettingsTemplateWithBodyWithResponse applies caching before delegating to the underlying client.
func (c *CachedClientWithResponsesInterface) PostSettingsTemplateWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostSettingsTemplateResponse, error) {
	// Build cache key
	key := generateKey("PostSettingsTemplateWithBodyWithResponse", ctx, contentType, body, reqEditors)
	if v, ok := c.cache.Get(key); ok {
		output := PostSettingsTemplateResponse{}
		err := json.Unmarshal(v, &output)
		if err == nil {
			return &output, nil
		}
	}

	// Cache miss: delegate to underlying client
	resp, err := c.client.PostSettingsTemplateWithBodyWithResponse(ctx, contentType, body, reqEditors...)
	if err != nil {
		return resp, err
	}
	c.cache.Set(key, resp)
	return resp, nil
}

// PostSettingsTemplateWithResponse applies caching before delegating to the underlying client.
func (c *CachedClientWithResponsesInterface) PostSettingsTemplateWithResponse(ctx context.Context, body PostSettingsTemplateJSONRequestBody, reqEditors ...RequestEditorFn) (*PostSettingsTemplateResponse, error) {
	// Build cache key
	key := generateKey("PostSettingsTemplateWithResponse", ctx, body, reqEditors)
	if v, ok := c.cache.Get(key); ok {
		output := PostSettingsTemplateResponse{}
		err := json.Unmarshal(v, &output)
		if err == nil {
			return &output, nil
		}
	}

	// Cache miss: delegate to underlying client
	resp, err := c.client.PostSettingsTemplateWithResponse(ctx, body, reqEditors...)
	if err != nil {
		return resp, err
	}
	c.cache.Set(key, resp)
	return resp, nil
}

// GetSettingsTemplateVersionWithResponse applies caching before delegating to the underlying client.
func (c *CachedClientWithResponsesInterface) GetSettingsTemplateVersionWithResponse(ctx context.Context, version openapi_types.UUID, reqEditors ...RequestEditorFn) (*GetSettingsTemplateVersionResponse, error) {
	// Build cache key
	key := generateKey("GetSettingsTemplateVersionWithResponse", ctx, version, reqEditors)
	if v, ok := c.cache.Get(key); ok {
		output := GetSettingsTemplateVersionResponse{}
		err := json.Unmarshal(v, &output)
		if err == nil {
			return &output, nil
		}
	}

	// Cache miss: delegate to underlying client
	resp, err := c.client.GetSettingsTemplateVersionWithResponse(ctx, version, reqEditors...)
	if err != nil {
		return resp, err
	}
	c.cache.Set(key, resp)
	return resp, nil
}

// GetStatisticsChaptersWithResponse applies caching before delegating to the underlying client.
func (c *CachedClientWithResponsesInterface) GetStatisticsChaptersWithResponse(ctx context.Context, params *GetStatisticsChaptersParams, reqEditors ...RequestEditorFn) (*GetStatisticsChaptersResponse, error) {
	// Build cache key
	key := generateKey("GetStatisticsChaptersWithResponse", ctx, params, reqEditors)
	if v, ok := c.cache.Get(key); ok {
		output := GetStatisticsChaptersResponse{}
		err := json.Unmarshal(v, &output)
		if err == nil {
			return &output, nil
		}
	}

	// Cache miss: delegate to underlying client
	resp, err := c.client.GetStatisticsChaptersWithResponse(ctx, params, reqEditors...)
	if err != nil {
		return resp, err
	}
	c.cache.Set(key, resp)
	return resp, nil
}

// GetStatisticsChapterUuidWithResponse applies caching before delegating to the underlying client.
func (c *CachedClientWithResponsesInterface) GetStatisticsChapterUuidWithResponse(ctx context.Context, uuid openapi_types.UUID, reqEditors ...RequestEditorFn) (*GetStatisticsChapterUuidResponse, error) {
	// Build cache key
	key := generateKey("GetStatisticsChapterUuidWithResponse", ctx, uuid, reqEditors)
	if v, ok := c.cache.Get(key); ok {
		output := GetStatisticsChapterUuidResponse{}
		err := json.Unmarshal(v, &output)
		if err == nil {
			return &output, nil
		}
	}

	// Cache miss: delegate to underlying client
	resp, err := c.client.GetStatisticsChapterUuidWithResponse(ctx, uuid, reqEditors...)
	if err != nil {
		return resp, err
	}
	c.cache.Set(key, resp)
	return resp, nil
}

// GetStatisticsGroupsWithResponse applies caching before delegating to the underlying client.
func (c *CachedClientWithResponsesInterface) GetStatisticsGroupsWithResponse(ctx context.Context, params *GetStatisticsGroupsParams, reqEditors ...RequestEditorFn) (*GetStatisticsGroupsResponse, error) {
	// Build cache key
	key := generateKey("GetStatisticsGroupsWithResponse", ctx, params, reqEditors)
	if v, ok := c.cache.Get(key); ok {
		output := GetStatisticsGroupsResponse{}
		err := json.Unmarshal(v, &output)
		if err == nil {
			return &output, nil
		}
	}

	// Cache miss: delegate to underlying client
	resp, err := c.client.GetStatisticsGroupsWithResponse(ctx, params, reqEditors...)
	if err != nil {
		return resp, err
	}
	c.cache.Set(key, resp)
	return resp, nil
}

// GetStatisticsGroupUuidWithResponse applies caching before delegating to the underlying client.
func (c *CachedClientWithResponsesInterface) GetStatisticsGroupUuidWithResponse(ctx context.Context, uuid openapi_types.UUID, reqEditors ...RequestEditorFn) (*GetStatisticsGroupUuidResponse, error) {
	// Build cache key
	key := generateKey("GetStatisticsGroupUuidWithResponse", ctx, uuid, reqEditors)
	if v, ok := c.cache.Get(key); ok {
		output := GetStatisticsGroupUuidResponse{}
		err := json.Unmarshal(v, &output)
		if err == nil {
			return &output, nil
		}
	}

	// Cache miss: delegate to underlying client
	resp, err := c.client.GetStatisticsGroupUuidWithResponse(ctx, uuid, reqEditors...)
	if err != nil {
		return resp, err
	}
	c.cache.Set(key, resp)
	return resp, nil
}

// GetStatisticsMangaWithResponse applies caching before delegating to the underlying client.
func (c *CachedClientWithResponsesInterface) GetStatisticsMangaWithResponse(ctx context.Context, params *GetStatisticsMangaParams, reqEditors ...RequestEditorFn) (*GetStatisticsMangaResponse, error) {
	// Build cache key
	key := generateKey("GetStatisticsMangaWithResponse", ctx, params, reqEditors)
	if v, ok := c.cache.Get(key); ok {
		output := GetStatisticsMangaResponse{}
		err := json.Unmarshal(v, &output)
		if err == nil {
			return &output, nil
		}
	}

	// Cache miss: delegate to underlying client
	resp, err := c.client.GetStatisticsMangaWithResponse(ctx, params, reqEditors...)
	if err != nil {
		return resp, err
	}
	c.cache.Set(key, resp)
	return resp, nil
}

// GetStatisticsMangaUuidWithResponse applies caching before delegating to the underlying client.
func (c *CachedClientWithResponsesInterface) GetStatisticsMangaUuidWithResponse(ctx context.Context, uuid openapi_types.UUID, reqEditors ...RequestEditorFn) (*GetStatisticsMangaUuidResponse, error) {
	// Build cache key
	key := generateKey("GetStatisticsMangaUuidWithResponse", ctx, uuid, reqEditors)
	if v, ok := c.cache.Get(key); ok {
		output := GetStatisticsMangaUuidResponse{}
		err := json.Unmarshal(v, &output)
		if err == nil {
			return &output, nil
		}
	}

	// Cache miss: delegate to underlying client
	resp, err := c.client.GetStatisticsMangaUuidWithResponse(ctx, uuid, reqEditors...)
	if err != nil {
		return resp, err
	}
	c.cache.Set(key, resp)
	return resp, nil
}

// GetUploadSessionWithResponse applies caching before delegating to the underlying client.
func (c *CachedClientWithResponsesInterface) GetUploadSessionWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetUploadSessionResponse, error) {
	// Build cache key
	key := generateKey("GetUploadSessionWithResponse", ctx, reqEditors)
	if v, ok := c.cache.Get(key); ok {
		output := GetUploadSessionResponse{}
		err := json.Unmarshal(v, &output)
		if err == nil {
			return &output, nil
		}
	}

	// Cache miss: delegate to underlying client
	resp, err := c.client.GetUploadSessionWithResponse(ctx, reqEditors...)
	if err != nil {
		return resp, err
	}
	c.cache.Set(key, resp)
	return resp, nil
}

// BeginUploadSessionWithBodyWithResponse applies caching before delegating to the underlying client.
func (c *CachedClientWithResponsesInterface) BeginUploadSessionWithBodyWithResponse(ctx context.Context, params *BeginUploadSessionParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*BeginUploadSessionResponse, error) {
	// Build cache key
	key := generateKey("BeginUploadSessionWithBodyWithResponse", ctx, params, contentType, body, reqEditors)
	if v, ok := c.cache.Get(key); ok {
		output := BeginUploadSessionResponse{}
		err := json.Unmarshal(v, &output)
		if err == nil {
			return &output, nil
		}
	}

	// Cache miss: delegate to underlying client
	resp, err := c.client.BeginUploadSessionWithBodyWithResponse(ctx, params, contentType, body, reqEditors...)
	if err != nil {
		return resp, err
	}
	c.cache.Set(key, resp)
	return resp, nil
}

// BeginUploadSessionWithResponse applies caching before delegating to the underlying client.
func (c *CachedClientWithResponsesInterface) BeginUploadSessionWithResponse(ctx context.Context, params *BeginUploadSessionParams, body BeginUploadSessionJSONRequestBody, reqEditors ...RequestEditorFn) (*BeginUploadSessionResponse, error) {
	// Build cache key
	key := generateKey("BeginUploadSessionWithResponse", ctx, params, body, reqEditors)
	if v, ok := c.cache.Get(key); ok {
		output := BeginUploadSessionResponse{}
		err := json.Unmarshal(v, &output)
		if err == nil {
			return &output, nil
		}
	}

	// Cache miss: delegate to underlying client
	resp, err := c.client.BeginUploadSessionWithResponse(ctx, params, body, reqEditors...)
	if err != nil {
		return resp, err
	}
	c.cache.Set(key, resp)
	return resp, nil
}

// BeginEditSessionWithBodyWithResponse applies caching before delegating to the underlying client.
func (c *CachedClientWithResponsesInterface) BeginEditSessionWithBodyWithResponse(ctx context.Context, chapterId openapi_types.UUID, params *BeginEditSessionParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*BeginEditSessionResponse, error) {
	// Build cache key
	key := generateKey("BeginEditSessionWithBodyWithResponse", ctx, chapterId, params, contentType, body, reqEditors)
	if v, ok := c.cache.Get(key); ok {
		output := BeginEditSessionResponse{}
		err := json.Unmarshal(v, &output)
		if err == nil {
			return &output, nil
		}
	}

	// Cache miss: delegate to underlying client
	resp, err := c.client.BeginEditSessionWithBodyWithResponse(ctx, chapterId, params, contentType, body, reqEditors...)
	if err != nil {
		return resp, err
	}
	c.cache.Set(key, resp)
	return resp, nil
}

// BeginEditSessionWithResponse applies caching before delegating to the underlying client.
func (c *CachedClientWithResponsesInterface) BeginEditSessionWithResponse(ctx context.Context, chapterId openapi_types.UUID, params *BeginEditSessionParams, body BeginEditSessionJSONRequestBody, reqEditors ...RequestEditorFn) (*BeginEditSessionResponse, error) {
	// Build cache key
	key := generateKey("BeginEditSessionWithResponse", ctx, chapterId, params, body, reqEditors)
	if v, ok := c.cache.Get(key); ok {
		output := BeginEditSessionResponse{}
		err := json.Unmarshal(v, &output)
		if err == nil {
			return &output, nil
		}
	}

	// Cache miss: delegate to underlying client
	resp, err := c.client.BeginEditSessionWithResponse(ctx, chapterId, params, body, reqEditors...)
	if err != nil {
		return resp, err
	}
	c.cache.Set(key, resp)
	return resp, nil
}

// UploadCheckApprovalRequiredWithBodyWithResponse applies caching before delegating to the underlying client.
func (c *CachedClientWithResponsesInterface) UploadCheckApprovalRequiredWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*UploadCheckApprovalRequiredResponse, error) {
	// Build cache key
	key := generateKey("UploadCheckApprovalRequiredWithBodyWithResponse", ctx, contentType, body, reqEditors)
	if v, ok := c.cache.Get(key); ok {
		output := UploadCheckApprovalRequiredResponse{}
		err := json.Unmarshal(v, &output)
		if err == nil {
			return &output, nil
		}
	}

	// Cache miss: delegate to underlying client
	resp, err := c.client.UploadCheckApprovalRequiredWithBodyWithResponse(ctx, contentType, body, reqEditors...)
	if err != nil {
		return resp, err
	}
	c.cache.Set(key, resp)
	return resp, nil
}

// UploadCheckApprovalRequiredWithResponse applies caching before delegating to the underlying client.
func (c *CachedClientWithResponsesInterface) UploadCheckApprovalRequiredWithResponse(ctx context.Context, body UploadCheckApprovalRequiredJSONRequestBody, reqEditors ...RequestEditorFn) (*UploadCheckApprovalRequiredResponse, error) {
	// Build cache key
	key := generateKey("UploadCheckApprovalRequiredWithResponse", ctx, body, reqEditors)
	if v, ok := c.cache.Get(key); ok {
		output := UploadCheckApprovalRequiredResponse{}
		err := json.Unmarshal(v, &output)
		if err == nil {
			return &output, nil
		}
	}

	// Cache miss: delegate to underlying client
	resp, err := c.client.UploadCheckApprovalRequiredWithResponse(ctx, body, reqEditors...)
	if err != nil {
		return resp, err
	}
	c.cache.Set(key, resp)
	return resp, nil
}

// AbandonUploadSessionWithResponse applies caching before delegating to the underlying client.
func (c *CachedClientWithResponsesInterface) AbandonUploadSessionWithResponse(ctx context.Context, uploadSessionId openapi_types.UUID, reqEditors ...RequestEditorFn) (*AbandonUploadSessionResponse, error) {
	// Build cache key
	key := generateKey("AbandonUploadSessionWithResponse", ctx, uploadSessionId, reqEditors)
	if v, ok := c.cache.Get(key); ok {
		output := AbandonUploadSessionResponse{}
		err := json.Unmarshal(v, &output)
		if err == nil {
			return &output, nil
		}
	}

	// Cache miss: delegate to underlying client
	resp, err := c.client.AbandonUploadSessionWithResponse(ctx, uploadSessionId, reqEditors...)
	if err != nil {
		return resp, err
	}
	c.cache.Set(key, resp)
	return resp, nil
}

// PutUploadSessionFileWithBodyWithResponse applies caching before delegating to the underlying client.
func (c *CachedClientWithResponsesInterface) PutUploadSessionFileWithBodyWithResponse(ctx context.Context, uploadSessionId openapi_types.UUID, params *PutUploadSessionFileParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutUploadSessionFileResponse, error) {
	// Build cache key
	key := generateKey("PutUploadSessionFileWithBodyWithResponse", ctx, uploadSessionId, params, contentType, body, reqEditors)
	if v, ok := c.cache.Get(key); ok {
		output := PutUploadSessionFileResponse{}
		err := json.Unmarshal(v, &output)
		if err == nil {
			return &output, nil
		}
	}

	// Cache miss: delegate to underlying client
	resp, err := c.client.PutUploadSessionFileWithBodyWithResponse(ctx, uploadSessionId, params, contentType, body, reqEditors...)
	if err != nil {
		return resp, err
	}
	c.cache.Set(key, resp)
	return resp, nil
}

// DeleteUploadedSessionFilesWithBodyWithResponse applies caching before delegating to the underlying client.
func (c *CachedClientWithResponsesInterface) DeleteUploadedSessionFilesWithBodyWithResponse(ctx context.Context, uploadSessionId openapi_types.UUID, params *DeleteUploadedSessionFilesParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*DeleteUploadedSessionFilesResponse, error) {
	// Build cache key
	key := generateKey("DeleteUploadedSessionFilesWithBodyWithResponse", ctx, uploadSessionId, params, contentType, body, reqEditors)
	if v, ok := c.cache.Get(key); ok {
		output := DeleteUploadedSessionFilesResponse{}
		err := json.Unmarshal(v, &output)
		if err == nil {
			return &output, nil
		}
	}

	// Cache miss: delegate to underlying client
	resp, err := c.client.DeleteUploadedSessionFilesWithBodyWithResponse(ctx, uploadSessionId, params, contentType, body, reqEditors...)
	if err != nil {
		return resp, err
	}
	c.cache.Set(key, resp)
	return resp, nil
}

// DeleteUploadedSessionFilesWithResponse applies caching before delegating to the underlying client.
func (c *CachedClientWithResponsesInterface) DeleteUploadedSessionFilesWithResponse(ctx context.Context, uploadSessionId openapi_types.UUID, params *DeleteUploadedSessionFilesParams, body DeleteUploadedSessionFilesJSONRequestBody, reqEditors ...RequestEditorFn) (*DeleteUploadedSessionFilesResponse, error) {
	// Build cache key
	key := generateKey("DeleteUploadedSessionFilesWithResponse", ctx, uploadSessionId, params, body, reqEditors)
	if v, ok := c.cache.Get(key); ok {
		output := DeleteUploadedSessionFilesResponse{}
		err := json.Unmarshal(v, &output)
		if err == nil {
			return &output, nil
		}
	}

	// Cache miss: delegate to underlying client
	resp, err := c.client.DeleteUploadedSessionFilesWithResponse(ctx, uploadSessionId, params, body, reqEditors...)
	if err != nil {
		return resp, err
	}
	c.cache.Set(key, resp)
	return resp, nil
}

// CommitUploadSessionWithBodyWithResponse applies caching before delegating to the underlying client.
func (c *CachedClientWithResponsesInterface) CommitUploadSessionWithBodyWithResponse(ctx context.Context, uploadSessionId openapi_types.UUID, params *CommitUploadSessionParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*CommitUploadSessionResponse, error) {
	// Build cache key
	key := generateKey("CommitUploadSessionWithBodyWithResponse", ctx, uploadSessionId, params, contentType, body, reqEditors)
	if v, ok := c.cache.Get(key); ok {
		output := CommitUploadSessionResponse{}
		err := json.Unmarshal(v, &output)
		if err == nil {
			return &output, nil
		}
	}

	// Cache miss: delegate to underlying client
	resp, err := c.client.CommitUploadSessionWithBodyWithResponse(ctx, uploadSessionId, params, contentType, body, reqEditors...)
	if err != nil {
		return resp, err
	}
	c.cache.Set(key, resp)
	return resp, nil
}

// CommitUploadSessionWithResponse applies caching before delegating to the underlying client.
func (c *CachedClientWithResponsesInterface) CommitUploadSessionWithResponse(ctx context.Context, uploadSessionId openapi_types.UUID, params *CommitUploadSessionParams, body CommitUploadSessionJSONRequestBody, reqEditors ...RequestEditorFn) (*CommitUploadSessionResponse, error) {
	// Build cache key
	key := generateKey("CommitUploadSessionWithResponse", ctx, uploadSessionId, params, body, reqEditors)
	if v, ok := c.cache.Get(key); ok {
		output := CommitUploadSessionResponse{}
		err := json.Unmarshal(v, &output)
		if err == nil {
			return &output, nil
		}
	}

	// Cache miss: delegate to underlying client
	resp, err := c.client.CommitUploadSessionWithResponse(ctx, uploadSessionId, params, body, reqEditors...)
	if err != nil {
		return resp, err
	}
	c.cache.Set(key, resp)
	return resp, nil
}

// DeleteUploadedSessionFileWithResponse applies caching before delegating to the underlying client.
func (c *CachedClientWithResponsesInterface) DeleteUploadedSessionFileWithResponse(ctx context.Context, uploadSessionId openapi_types.UUID, uploadSessionFileId openapi_types.UUID, reqEditors ...RequestEditorFn) (*DeleteUploadedSessionFileResponse, error) {
	// Build cache key
	key := generateKey("DeleteUploadedSessionFileWithResponse", ctx, uploadSessionId, uploadSessionFileId, reqEditors)
	if v, ok := c.cache.Get(key); ok {
		output := DeleteUploadedSessionFileResponse{}
		err := json.Unmarshal(v, &output)
		if err == nil {
			return &output, nil
		}
	}

	// Cache miss: delegate to underlying client
	resp, err := c.client.DeleteUploadedSessionFileWithResponse(ctx, uploadSessionId, uploadSessionFileId, reqEditors...)
	if err != nil {
		return resp, err
	}
	c.cache.Set(key, resp)
	return resp, nil
}

// GetUserWithResponse applies caching before delegating to the underlying client.
func (c *CachedClientWithResponsesInterface) GetUserWithResponse(ctx context.Context, params *GetUserParams, reqEditors ...RequestEditorFn) (*GetUserResponse, error) {
	// Build cache key
	key := generateKey("GetUserWithResponse", ctx, params, reqEditors)
	if v, ok := c.cache.Get(key); ok {
		output := GetUserResponse{}
		err := json.Unmarshal(v, &output)
		if err == nil {
			return &output, nil
		}
	}

	// Cache miss: delegate to underlying client
	resp, err := c.client.GetUserWithResponse(ctx, params, reqEditors...)
	if err != nil {
		return resp, err
	}
	c.cache.Set(key, resp)
	return resp, nil
}

// PostUserDeleteCodeWithResponse applies caching before delegating to the underlying client.
func (c *CachedClientWithResponsesInterface) PostUserDeleteCodeWithResponse(ctx context.Context, code openapi_types.UUID, reqEditors ...RequestEditorFn) (*PostUserDeleteCodeResponse, error) {
	// Build cache key
	key := generateKey("PostUserDeleteCodeWithResponse", ctx, code, reqEditors)
	if v, ok := c.cache.Get(key); ok {
		output := PostUserDeleteCodeResponse{}
		err := json.Unmarshal(v, &output)
		if err == nil {
			return &output, nil
		}
	}

	// Cache miss: delegate to underlying client
	resp, err := c.client.PostUserDeleteCodeWithResponse(ctx, code, reqEditors...)
	if err != nil {
		return resp, err
	}
	c.cache.Set(key, resp)
	return resp, nil
}

// GetUserFollowsGroupWithResponse applies caching before delegating to the underlying client.
func (c *CachedClientWithResponsesInterface) GetUserFollowsGroupWithResponse(ctx context.Context, params *GetUserFollowsGroupParams, reqEditors ...RequestEditorFn) (*GetUserFollowsGroupResponse, error) {
	// Build cache key
	key := generateKey("GetUserFollowsGroupWithResponse", ctx, params, reqEditors)
	if v, ok := c.cache.Get(key); ok {
		output := GetUserFollowsGroupResponse{}
		err := json.Unmarshal(v, &output)
		if err == nil {
			return &output, nil
		}
	}

	// Cache miss: delegate to underlying client
	resp, err := c.client.GetUserFollowsGroupWithResponse(ctx, params, reqEditors...)
	if err != nil {
		return resp, err
	}
	c.cache.Set(key, resp)
	return resp, nil
}

// GetUserFollowsGroupIdWithResponse applies caching before delegating to the underlying client.
func (c *CachedClientWithResponsesInterface) GetUserFollowsGroupIdWithResponse(ctx context.Context, id openapi_types.UUID, reqEditors ...RequestEditorFn) (*GetUserFollowsGroupIdResponse, error) {
	// Build cache key
	key := generateKey("GetUserFollowsGroupIdWithResponse", ctx, id, reqEditors)
	if v, ok := c.cache.Get(key); ok {
		output := GetUserFollowsGroupIdResponse{}
		err := json.Unmarshal(v, &output)
		if err == nil {
			return &output, nil
		}
	}

	// Cache miss: delegate to underlying client
	resp, err := c.client.GetUserFollowsGroupIdWithResponse(ctx, id, reqEditors...)
	if err != nil {
		return resp, err
	}
	c.cache.Set(key, resp)
	return resp, nil
}

// GetUserFollowsListWithResponse applies caching before delegating to the underlying client.
func (c *CachedClientWithResponsesInterface) GetUserFollowsListWithResponse(ctx context.Context, params *GetUserFollowsListParams, reqEditors ...RequestEditorFn) (*GetUserFollowsListResponse, error) {
	// Build cache key
	key := generateKey("GetUserFollowsListWithResponse", ctx, params, reqEditors)
	if v, ok := c.cache.Get(key); ok {
		output := GetUserFollowsListResponse{}
		err := json.Unmarshal(v, &output)
		if err == nil {
			return &output, nil
		}
	}

	// Cache miss: delegate to underlying client
	resp, err := c.client.GetUserFollowsListWithResponse(ctx, params, reqEditors...)
	if err != nil {
		return resp, err
	}
	c.cache.Set(key, resp)
	return resp, nil
}

// GetUserFollowsListIdWithResponse applies caching before delegating to the underlying client.
func (c *CachedClientWithResponsesInterface) GetUserFollowsListIdWithResponse(ctx context.Context, id openapi_types.UUID, reqEditors ...RequestEditorFn) (*GetUserFollowsListIdResponse, error) {
	// Build cache key
	key := generateKey("GetUserFollowsListIdWithResponse", ctx, id, reqEditors)
	if v, ok := c.cache.Get(key); ok {
		output := GetUserFollowsListIdResponse{}
		err := json.Unmarshal(v, &output)
		if err == nil {
			return &output, nil
		}
	}

	// Cache miss: delegate to underlying client
	resp, err := c.client.GetUserFollowsListIdWithResponse(ctx, id, reqEditors...)
	if err != nil {
		return resp, err
	}
	c.cache.Set(key, resp)
	return resp, nil
}

// GetUserFollowsMangaWithResponse applies caching before delegating to the underlying client.
func (c *CachedClientWithResponsesInterface) GetUserFollowsMangaWithResponse(ctx context.Context, params *GetUserFollowsMangaParams, reqEditors ...RequestEditorFn) (*GetUserFollowsMangaResponse, error) {
	// Build cache key
	key := generateKey("GetUserFollowsMangaWithResponse", ctx, params, reqEditors)
	if v, ok := c.cache.Get(key); ok {
		output := GetUserFollowsMangaResponse{}
		err := json.Unmarshal(v, &output)
		if err == nil {
			return &output, nil
		}
	}

	// Cache miss: delegate to underlying client
	resp, err := c.client.GetUserFollowsMangaWithResponse(ctx, params, reqEditors...)
	if err != nil {
		return resp, err
	}
	c.cache.Set(key, resp)
	return resp, nil
}

// GetUserFollowsMangaFeedWithResponse applies caching before delegating to the underlying client.
func (c *CachedClientWithResponsesInterface) GetUserFollowsMangaFeedWithResponse(ctx context.Context, params *GetUserFollowsMangaFeedParams, reqEditors ...RequestEditorFn) (*GetUserFollowsMangaFeedResponse, error) {
	// Build cache key
	key := generateKey("GetUserFollowsMangaFeedWithResponse", ctx, params, reqEditors)
	if v, ok := c.cache.Get(key); ok {
		output := GetUserFollowsMangaFeedResponse{}
		err := json.Unmarshal(v, &output)
		if err == nil {
			return &output, nil
		}
	}

	// Cache miss: delegate to underlying client
	resp, err := c.client.GetUserFollowsMangaFeedWithResponse(ctx, params, reqEditors...)
	if err != nil {
		return resp, err
	}
	c.cache.Set(key, resp)
	return resp, nil
}

// GetUserFollowsMangaIdWithResponse applies caching before delegating to the underlying client.
func (c *CachedClientWithResponsesInterface) GetUserFollowsMangaIdWithResponse(ctx context.Context, id openapi_types.UUID, reqEditors ...RequestEditorFn) (*GetUserFollowsMangaIdResponse, error) {
	// Build cache key
	key := generateKey("GetUserFollowsMangaIdWithResponse", ctx, id, reqEditors)
	if v, ok := c.cache.Get(key); ok {
		output := GetUserFollowsMangaIdResponse{}
		err := json.Unmarshal(v, &output)
		if err == nil {
			return &output, nil
		}
	}

	// Cache miss: delegate to underlying client
	resp, err := c.client.GetUserFollowsMangaIdWithResponse(ctx, id, reqEditors...)
	if err != nil {
		return resp, err
	}
	c.cache.Set(key, resp)
	return resp, nil
}

// GetUserFollowsUserWithResponse applies caching before delegating to the underlying client.
func (c *CachedClientWithResponsesInterface) GetUserFollowsUserWithResponse(ctx context.Context, params *GetUserFollowsUserParams, reqEditors ...RequestEditorFn) (*GetUserFollowsUserResponse, error) {
	// Build cache key
	key := generateKey("GetUserFollowsUserWithResponse", ctx, params, reqEditors)
	if v, ok := c.cache.Get(key); ok {
		output := GetUserFollowsUserResponse{}
		err := json.Unmarshal(v, &output)
		if err == nil {
			return &output, nil
		}
	}

	// Cache miss: delegate to underlying client
	resp, err := c.client.GetUserFollowsUserWithResponse(ctx, params, reqEditors...)
	if err != nil {
		return resp, err
	}
	c.cache.Set(key, resp)
	return resp, nil
}

// GetUserFollowsUserIdWithResponse applies caching before delegating to the underlying client.
func (c *CachedClientWithResponsesInterface) GetUserFollowsUserIdWithResponse(ctx context.Context, id openapi_types.UUID, reqEditors ...RequestEditorFn) (*GetUserFollowsUserIdResponse, error) {
	// Build cache key
	key := generateKey("GetUserFollowsUserIdWithResponse", ctx, id, reqEditors)
	if v, ok := c.cache.Get(key); ok {
		output := GetUserFollowsUserIdResponse{}
		err := json.Unmarshal(v, &output)
		if err == nil {
			return &output, nil
		}
	}

	// Cache miss: delegate to underlying client
	resp, err := c.client.GetUserFollowsUserIdWithResponse(ctx, id, reqEditors...)
	if err != nil {
		return resp, err
	}
	c.cache.Set(key, resp)
	return resp, nil
}

// GetReadingHistoryWithResponse applies caching before delegating to the underlying client.
func (c *CachedClientWithResponsesInterface) GetReadingHistoryWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetReadingHistoryResponse, error) {
	// Build cache key
	key := generateKey("GetReadingHistoryWithResponse", ctx, reqEditors)
	if v, ok := c.cache.Get(key); ok {
		output := GetReadingHistoryResponse{}
		err := json.Unmarshal(v, &output)
		if err == nil {
			return &output, nil
		}
	}

	// Cache miss: delegate to underlying client
	resp, err := c.client.GetReadingHistoryWithResponse(ctx, reqEditors...)
	if err != nil {
		return resp, err
	}
	c.cache.Set(key, resp)
	return resp, nil
}

// GetUserListWithResponse applies caching before delegating to the underlying client.
func (c *CachedClientWithResponsesInterface) GetUserListWithResponse(ctx context.Context, params *GetUserListParams, reqEditors ...RequestEditorFn) (*GetUserListResponse, error) {
	// Build cache key
	key := generateKey("GetUserListWithResponse", ctx, params, reqEditors)
	if v, ok := c.cache.Get(key); ok {
		output := GetUserListResponse{}
		err := json.Unmarshal(v, &output)
		if err == nil {
			return &output, nil
		}
	}

	// Cache miss: delegate to underlying client
	resp, err := c.client.GetUserListWithResponse(ctx, params, reqEditors...)
	if err != nil {
		return resp, err
	}
	c.cache.Set(key, resp)
	return resp, nil
}

// GetUserMeWithResponse applies caching before delegating to the underlying client.
func (c *CachedClientWithResponsesInterface) GetUserMeWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetUserMeResponse, error) {
	// Build cache key
	key := generateKey("GetUserMeWithResponse", ctx, reqEditors)
	if v, ok := c.cache.Get(key); ok {
		output := GetUserMeResponse{}
		err := json.Unmarshal(v, &output)
		if err == nil {
			return &output, nil
		}
	}

	// Cache miss: delegate to underlying client
	resp, err := c.client.GetUserMeWithResponse(ctx, reqEditors...)
	if err != nil {
		return resp, err
	}
	c.cache.Set(key, resp)
	return resp, nil
}

// DeleteUserIdWithResponse applies caching before delegating to the underlying client.
func (c *CachedClientWithResponsesInterface) DeleteUserIdWithResponse(ctx context.Context, id openapi_types.UUID, reqEditors ...RequestEditorFn) (*DeleteUserIdResponse, error) {
	// Build cache key
	key := generateKey("DeleteUserIdWithResponse", ctx, id, reqEditors)
	if v, ok := c.cache.Get(key); ok {
		output := DeleteUserIdResponse{}
		err := json.Unmarshal(v, &output)
		if err == nil {
			return &output, nil
		}
	}

	// Cache miss: delegate to underlying client
	resp, err := c.client.DeleteUserIdWithResponse(ctx, id, reqEditors...)
	if err != nil {
		return resp, err
	}
	c.cache.Set(key, resp)
	return resp, nil
}

// GetUserIdWithResponse applies caching before delegating to the underlying client.
func (c *CachedClientWithResponsesInterface) GetUserIdWithResponse(ctx context.Context, id openapi_types.UUID, reqEditors ...RequestEditorFn) (*GetUserIdResponse, error) {
	// Build cache key
	key := generateKey("GetUserIdWithResponse", ctx, id, reqEditors)
	if v, ok := c.cache.Get(key); ok {
		output := GetUserIdResponse{}
		err := json.Unmarshal(v, &output)
		if err == nil {
			return &output, nil
		}
	}

	// Cache miss: delegate to underlying client
	resp, err := c.client.GetUserIdWithResponse(ctx, id, reqEditors...)
	if err != nil {
		return resp, err
	}
	c.cache.Set(key, resp)
	return resp, nil
}

// GetUserIdListWithResponse applies caching before delegating to the underlying client.
func (c *CachedClientWithResponsesInterface) GetUserIdListWithResponse(ctx context.Context, id openapi_types.UUID, params *GetUserIdListParams, reqEditors ...RequestEditorFn) (*GetUserIdListResponse, error) {
	// Build cache key
	key := generateKey("GetUserIdListWithResponse", ctx, id, params, reqEditors)
	if v, ok := c.cache.Get(key); ok {
		output := GetUserIdListResponse{}
		err := json.Unmarshal(v, &output)
		if err == nil {
			return &output, nil
		}
	}

	// Cache miss: delegate to underlying client
	resp, err := c.client.GetUserIdListWithResponse(ctx, id, params, reqEditors...)
	if err != nil {
		return resp, err
	}
	c.cache.Set(key, resp)
	return resp, nil
}
