package client

import (
	"context"
	"net/http"
	"net/url"
)

// UploadAssetWithAppId uploads an asset file to the Busy Bar API with the specified application ID.
func (c *Client) UploadAssetWithAppId(ctx context.Context, query url.Values, payload any) (*SuccessResponse, error) {
	return doJSON[SuccessResponse](c, ctx, http.MethodPost, "/api/assets/upload", query, payload)
}

// UploadAsset uploads an asset file to the Busy Bar API.
func (c *Client) DeleteAppAssets(ctx context.Context, query url.Values) (*SuccessResponse, error) {
	return doJSON[SuccessResponse](c, ctx, http.MethodDelete, "/api/assets/upload", query, nil)
}

// DrawOnDisplayWithAppId draws on the Busy Bar display with the specified application ID.
func (c *Client) DrawOnDisplay(ctx context.Context, query url.Values, payload any) (*SuccessResponse, error) {
	return doJSON[SuccessResponse](c, ctx, http.MethodPost, "/api/display/draw", query, payload)
}

// ClearDisplay clears the Busy Bar display.
func (c *Client) ClearDisplay(ctx context.Context, query url.Values) (*SuccessResponse, error) {
	return doJSON[SuccessResponse](c, ctx, http.MethodDelete, "/api/display/draw", query, nil)
}

// PlayAudio plays audio on the Busy Bar with the specified application ID.
func (c *Client) PlayAudio(ctx context.Context, query url.Values, payload any) (*SuccessResponse, error) {
	return doJSON[SuccessResponse](c, ctx, http.MethodPost, "/api/audio/play", query, payload)
}

// StopAudio stops audio playback on the Busy Bar.
func (c *Client) StopAudio(ctx context.Context, query url.Values) (*SuccessResponse, error) {
	return doJSON[SuccessResponse](c, ctx, http.MethodDelete, "/api/audio/play", query, nil)
}
