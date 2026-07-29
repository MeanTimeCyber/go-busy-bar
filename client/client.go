package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

const apiTokenHeader = "X-API-Token"

type Client struct {
	BaseURL    string
	APIKey     string
	HTTPClient *http.Client
}

func NewClient(baseURL string) *Client {
	return &Client{
		BaseURL:    strings.TrimRight(baseURL, "/"),
		HTTPClient: http.DefaultClient,
	}
}

func NewClientWithAPIKey(baseURL, apiKey string) *Client {
	client := NewClient(baseURL)
	client.APIKey = apiKey
	return client
}

func (c *Client) do(ctx context.Context, method, path string, query url.Values, payload any) ([]byte, error) {
	client := c.HTTPClient
	if client == nil {
		client = http.DefaultClient
	}

	var body io.Reader
	if payload != nil {
		encodedPayload, err := json.Marshal(payload)
		if err != nil {
			return nil, err
		}
		body = bytes.NewReader(encodedPayload)
	}

	req, err := http.NewRequestWithContext(ctx, method, c.BaseURL+path, body)
	if err != nil {
		return nil, err
	}
	if query != nil {
		req.URL.RawQuery = query.Encode()
	}

	req.Header.Set("accept", "application/json")
	if payload != nil {
		req.Header.Set("content-type", "application/json")
	}
	if c.APIKey != "" {
		req.Header.Set(apiTokenHeader, c.APIKey)
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusMultipleChoices {
		return nil, fmt.Errorf("%s %s returned status %d: %s", method, req.URL.String(), resp.StatusCode, string(responseBody))
	}

	return responseBody, nil
}

func (c *Client) UnlinkAccount(ctx context.Context, query url.Values) ([]byte, error) {
	return c.do(ctx, http.MethodDelete, "/api/account", query, nil)
}

func (c *Client) LinkAccount(ctx context.Context, query url.Values, payload any) ([]byte, error) {
	return c.do(ctx, http.MethodPost, "/api/account/link", query, payload)
}

func (c *Client) GetAccountInfo(ctx context.Context, query url.Values) ([]byte, error) {
	return c.do(ctx, http.MethodGet, "/api/account/info", query, nil)
}

func (c *Client) GetAccountStatus(ctx context.Context, query url.Values) ([]byte, error) {
	return c.do(ctx, http.MethodGet, "/api/account/status", query, nil)
}

func (c *Client) GetAccountBackend(ctx context.Context, query url.Values) ([]byte, error) {
	return c.do(ctx, http.MethodGet, "/api/account/backend", query, nil)
}

func (c *Client) SetAccountBackend(ctx context.Context, query url.Values, payload any) ([]byte, error) {
	return c.do(ctx, http.MethodPut, "/api/account/backend", query, payload)
}

func (c *Client) UploadAssetWithAppId(ctx context.Context, query url.Values, payload any) ([]byte, error) {
	return c.do(ctx, http.MethodPost, "/api/assets/upload", query, payload)
}

func (c *Client) DeleteAppAssets(ctx context.Context, query url.Values) ([]byte, error) {
	return c.do(ctx, http.MethodDelete, "/api/assets/upload", query, nil)
}

func (c *Client) DrawOnDisplay(ctx context.Context, query url.Values, payload any) ([]byte, error) {
	return c.do(ctx, http.MethodPost, "/api/display/draw", query, payload)
}

func (c *Client) ClearDisplay(ctx context.Context, query url.Values) ([]byte, error) {
	return c.do(ctx, http.MethodDelete, "/api/display/draw", query, nil)
}

func (c *Client) PlayAudio(ctx context.Context, query url.Values, payload any) ([]byte, error) {
	return c.do(ctx, http.MethodPost, "/api/audio/play", query, payload)
}

func (c *Client) StopAudio(ctx context.Context, query url.Values) ([]byte, error) {
	return c.do(ctx, http.MethodDelete, "/api/audio/play", query, nil)
}

func (c *Client) PostBleEnable(ctx context.Context, query url.Values, payload any) ([]byte, error) {
	return c.do(ctx, http.MethodPost, "/api/ble/enable", query, payload)
}

func (c *Client) PostBleDisable(ctx context.Context, query url.Values, payload any) ([]byte, error) {
	return c.do(ctx, http.MethodPost, "/api/ble/disable", query, payload)
}

func (c *Client) DeleteBlePairing(ctx context.Context, query url.Values) ([]byte, error) {
	return c.do(ctx, http.MethodDelete, "/api/ble/pairing", query, nil)
}

func (c *Client) GetBleStatus(ctx context.Context, query url.Values) ([]byte, error) {
	return c.do(ctx, http.MethodGet, "/api/ble/status", query, nil)
}

func (c *Client) GetBusySnapshot(ctx context.Context, query url.Values) ([]byte, error) {
	return c.do(ctx, http.MethodGet, "/api/busy/snapshot", query, nil)
}

func (c *Client) SetBusySnapshot(ctx context.Context, query url.Values, payload any) ([]byte, error) {
	return c.do(ctx, http.MethodPut, "/api/busy/snapshot", query, payload)
}

func (c *Client) GetBusyProfile(ctx context.Context, slot string, query url.Values) ([]byte, error) {
	return c.do(ctx, http.MethodGet, fmt.Sprintf("/api/busy/profiles/%s", url.PathEscape(slot)), query, nil)
}

func (c *Client) SetBusyProfile(ctx context.Context, slot string, query url.Values, payload any) ([]byte, error) {
	return c.do(ctx, http.MethodPut, fmt.Sprintf("/api/busy/profiles/%s", url.PathEscape(slot)), query, payload)
}

func (c *Client) SetInputKey(ctx context.Context, query url.Values, payload any) ([]byte, error) {
	return c.do(ctx, http.MethodPost, "/api/input", query, payload)
}

func (c *Client) GetHttpAccess(ctx context.Context, query url.Values) ([]byte, error) {
	return c.do(ctx, http.MethodGet, "/api/access", query, nil)
}

func (c *Client) SetHttpAccess(ctx context.Context, query url.Values, payload any) ([]byte, error) {
	return c.do(ctx, http.MethodPost, "/api/access", query, payload)
}

func (c *Client) GetName(ctx context.Context, query url.Values) ([]byte, error) {
	return c.do(ctx, http.MethodGet, "/api/name", query, nil)
}

func (c *Client) PostName(ctx context.Context, query url.Values, payload any) ([]byte, error) {
	return c.do(ctx, http.MethodPost, "/api/name", query, payload)
}

func (c *Client) GetDisplayBrightness(ctx context.Context, query url.Values) ([]byte, error) {
	return c.do(ctx, http.MethodGet, "/api/display/brightness", query, nil)
}

func (c *Client) SetDisplayBrightness(ctx context.Context, query url.Values, payload any) ([]byte, error) {
	return c.do(ctx, http.MethodPost, "/api/display/brightness", query, payload)
}

func (c *Client) GetAudioVolume(ctx context.Context, query url.Values) ([]byte, error) {
	return c.do(ctx, http.MethodGet, "/api/audio/volume", query, nil)
}

func (c *Client) SetAudioVolume(ctx context.Context, query url.Values, payload any) ([]byte, error) {
	return c.do(ctx, http.MethodPost, "/api/audio/volume", query, payload)
}

func (c *Client) GetSmartHomeCommissioningStatus(ctx context.Context, query url.Values) ([]byte, error) {
	return c.do(ctx, http.MethodGet, "/api/smart_home/pairing", query, nil)
}

func (c *Client) StartSmartHomePairing(ctx context.Context, query url.Values, payload any) ([]byte, error) {
	return c.do(ctx, http.MethodPost, "/api/smart_home/pairing", query, payload)
}

func (c *Client) DeleteSmartHomePairing(ctx context.Context, query url.Values) ([]byte, error) {
	return c.do(ctx, http.MethodDelete, "/api/smart_home/pairing", query, nil)
}

func (c *Client) GetSmartHomeSwitch(ctx context.Context, query url.Values) ([]byte, error) {
	return c.do(ctx, http.MethodGet, "/api/smart_home/switch", query, nil)
}

func (c *Client) PostSmartHomeSwitch(ctx context.Context, query url.Values, payload any) ([]byte, error) {
	return c.do(ctx, http.MethodPost, "/api/smart_home/switch", query, payload)
}

func (c *Client) WriteStorageFile(ctx context.Context, query url.Values, payload any) ([]byte, error) {
	return c.do(ctx, http.MethodPost, "/api/storage/write", query, payload)
}

func (c *Client) ReadStorageFile(ctx context.Context, query url.Values) ([]byte, error) {
	return c.do(ctx, http.MethodGet, "/api/storage/read", query, nil)
}

func (c *Client) ListStorageFiles(ctx context.Context, query url.Values) ([]byte, error) {
	return c.do(ctx, http.MethodGet, "/api/storage/list", query, nil)
}

func (c *Client) RemoveStorageFile(ctx context.Context, query url.Values) ([]byte, error) {
	return c.do(ctx, http.MethodDelete, "/api/storage/remove", query, nil)
}

func (c *Client) CreateStorageDir(ctx context.Context, query url.Values, payload any) ([]byte, error) {
	return c.do(ctx, http.MethodPost, "/api/storage/mkdir", query, payload)
}

func (c *Client) RenameStorageFile(ctx context.Context, query url.Values, payload any) ([]byte, error) {
	return c.do(ctx, http.MethodPost, "/api/storage/rename", query, payload)
}

func (c *Client) GetStorageStatus(ctx context.Context, query url.Values) ([]byte, error) {
	return c.do(ctx, http.MethodGet, "/api/storage/status", query, nil)
}

func (c *Client) ConnectWebSocket(ctx context.Context, query url.Values) ([]byte, error) {
	return c.do(ctx, http.MethodGet, "/api/status/ws", query, nil)
}

func (c *Client) GetScreen(ctx context.Context, query url.Values) ([]byte, error) {
	return c.do(ctx, http.MethodGet, "/api/screen", query, nil)
}

func (c *Client) GetVersion(ctx context.Context, query url.Values) ([]byte, error) {
	return c.do(ctx, http.MethodGet, "/api/version", query, nil)
}

func (c *Client) GetTransport(ctx context.Context, query url.Values) ([]byte, error) {
	return c.do(ctx, http.MethodGet, "/api/transport", query, nil)
}

func (c *Client) GetStatus(ctx context.Context, query url.Values) (*Status, error) {
	resp, err := c.do(ctx, http.MethodGet, "/api/status", query, nil)
	
	if err != nil {
		return nil, err
	}
	var status Status
	
	if err := json.Unmarshal(resp, &status); err != nil {
		return nil, err
	}
	
	return &status, nil
}

func (c *Client) GetStatusDevice(ctx context.Context, query url.Values) ([]byte, error) {
	return c.do(ctx, http.MethodGet, "/api/status/device", query, nil)
}

func (c *Client) GetStatusFirmware(ctx context.Context, query url.Values) ([]byte, error) {
	return c.do(ctx, http.MethodGet, "/api/status/firmware", query, nil)
}

func (c *Client) GetStatusSystem(ctx context.Context, query url.Values) ([]byte, error) {
	return c.do(ctx, http.MethodGet, "/api/status/system", query, nil)
}

func (c *Client) GetStatusPower(ctx context.Context, query url.Values) ([]byte, error) {
	return c.do(ctx, http.MethodGet, "/api/status/power", query, nil)
}

func (c *Client) DumpLog(ctx context.Context, query url.Values, payload any) ([]byte, error) {
	return c.do(ctx, http.MethodPost, "/api/log_dump", query, payload)
}

func (c *Client) GetTime(ctx context.Context, query url.Values) ([]byte, error) {
	return c.do(ctx, http.MethodGet, "/api/time", query, nil)
}

func (c *Client) SetTimeTimestamp(ctx context.Context, query url.Values, payload any) ([]byte, error) {
	return c.do(ctx, http.MethodPost, "/api/time/timestamp", query, payload)
}

func (c *Client) GetTimeTimezone(ctx context.Context, query url.Values) ([]byte, error) {
	return c.do(ctx, http.MethodGet, "/api/time/timezone", query, nil)
}

func (c *Client) SetTimeTimezone(ctx context.Context, query url.Values, payload any) ([]byte, error) {
	return c.do(ctx, http.MethodPost, "/api/time/timezone", query, payload)
}

func (c *Client) GetTimeTzlist(ctx context.Context, query url.Values) ([]byte, error) {
	return c.do(ctx, http.MethodGet, "/api/time/tzlist", query, nil)
}

func (c *Client) UpdateFirmware(ctx context.Context, query url.Values, payload any) ([]byte, error) {
	return c.do(ctx, http.MethodPost, "/api/update", query, payload)
}

func (c *Client) CheckFirmwareUpdate(ctx context.Context, query url.Values, payload any) ([]byte, error) {
	return c.do(ctx, http.MethodPost, "/api/update/check", query, payload)
}

func (c *Client) GetFirmwareUpdateStatus(ctx context.Context, query url.Values) ([]byte, error) {
	return c.do(ctx, http.MethodGet, "/api/update/status", query, nil)
}

func (c *Client) GetUpdateChangelog(ctx context.Context, query url.Values) ([]byte, error) {
	return c.do(ctx, http.MethodGet, "/api/update/changelog", query, nil)
}

func (c *Client) InstallFirmwareUpdate(ctx context.Context, query url.Values, payload any) ([]byte, error) {
	return c.do(ctx, http.MethodPost, "/api/update/install", query, payload)
}

func (c *Client) AbortFirmwareDownload(ctx context.Context, query url.Values, payload any) ([]byte, error) {
	return c.do(ctx, http.MethodPost, "/api/update/abort_download", query, payload)
}

func (c *Client) GetAutoupdateSettings(ctx context.Context, query url.Values) ([]byte, error) {
	return c.do(ctx, http.MethodGet, "/api/update/autoupdate", query, nil)
}

func (c *Client) SetAutoupdateSettings(ctx context.Context, query url.Values, payload any) ([]byte, error) {
	return c.do(ctx, http.MethodPost, "/api/update/autoupdate", query, payload)
}

func (c *Client) GetWifiStatus(ctx context.Context, query url.Values) ([]byte, error) {
	return c.do(ctx, http.MethodGet, "/api/wifi/status", query, nil)
}

func (c *Client) PostWifiConnect(ctx context.Context, query url.Values, payload any) ([]byte, error) {
	return c.do(ctx, http.MethodPost, "/api/wifi/connect", query, payload)
}

func (c *Client) PostWifiDisconnect(ctx context.Context, query url.Values, payload any) ([]byte, error) {
	return c.do(ctx, http.MethodPost, "/api/wifi/disconnect", query, payload)
}

func (c *Client) GetWifiNetworks(ctx context.Context, query url.Values) ([]byte, error) {
	return c.do(ctx, http.MethodGet, "/api/wifi/networks", query, nil)
}
