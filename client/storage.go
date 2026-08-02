package client

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"os"

	"github.com/markkurossi/tabulate"
)

type StorageList struct {
	// List contains storage file and directory entries.
	List []StorageListElement `json:"list"`
}

type StorageListElement struct {
	// Type is the element type (file or dir).
	Type string `json:"type"`
	// Name is the file or directory name.
	Name string `json:"name"`
	// Size is the file size in bytes and is present for file entries.
	Size int `json:"size,omitempty"`
}

type StorageStatus struct {
	// UsedBytes is used storage capacity in bytes.
	UsedBytes int `json:"used_bytes"`
	// FreeBytes is remaining storage capacity in bytes.
	FreeBytes int `json:"free_bytes"`
	// TotalBytes is total storage partition size in bytes.
	TotalBytes int `json:"total_bytes"`
}

func storagePathQuery(path string) url.Values {
	return requiredQuery(map[string]string{"path": path})
}

func storageRenameQuery(path, newPath string) url.Values {
	return requiredQuery(map[string]string{
		"path":     path,
		"new_path": newPath,
	})
}

// WriteStorageFile writes a file to the device storage.
func (c *Client) WriteStorageFile(ctx context.Context, path string, payload any) (*SuccessResponse, error) {
	return doJSON[SuccessResponse](c, ctx, http.MethodPost, "/api/storage/write", storagePathQuery(path), payload)
}

// ReadStorageFile reads a file from the device storage.
func (c *Client) ReadStorageFile(ctx context.Context, path string) ([]byte, error) {
	return c.do(ctx, http.MethodGet, "/api/storage/read", storagePathQuery(path), nil)
}

// ListStorageFiles retrieves the list of files and directories from the device storage.
func (c *Client) ListStorageFiles(ctx context.Context, path string) (*StorageList, error) {
	return doJSON[StorageList](c, ctx, http.MethodGet, "/api/storage/list", storagePathQuery(path), nil)
}

// RemoveStorageFile removes a file from the device storage.
func (c *Client) RemoveStorageFile(ctx context.Context, path string) (*SuccessResponse, error) {
	return doJSON[SuccessResponse](c, ctx, http.MethodDelete, "/api/storage/remove", storagePathQuery(path), nil)
}

// CreateStorageDir creates a directory in the device storage.
func (c *Client) CreateStorageDir(ctx context.Context, path string, payload any) (*SuccessResponse, error) {
	return doJSON[SuccessResponse](c, ctx, http.MethodPost, "/api/storage/mkdir", storagePathQuery(path), payload)
}

// RenameStorageFile renames a file in the device storage.
func (c *Client) RenameStorageFile(ctx context.Context, path, newPath string, payload any) (*SuccessResponse, error) {
	return doJSON[SuccessResponse](c, ctx, http.MethodPost, "/api/storage/rename", storageRenameQuery(path, newPath), payload)
}

// GetStorageStatus retrieves the storage status information from the Busy Bar API.
func (c *Client) GetStorageStatus(ctx context.Context, query url.Values) (*StorageStatus, error) {
	return doJSON[StorageStatus](c, ctx, http.MethodGet, "/api/storage/status", query, nil)
}

func (s *StorageList) PrettyPrint() {
	fmt.Printf("\nStorage Files\n")

	tab := tabulate.New(tabulate.Unicode)
	tab.Header("Type").SetAlign(tabulate.ML)
	tab.Header("Name").SetAlign(tabulate.ML)
	tab.Header("Size").SetAlign(tabulate.ML)

	for _, item := range s.List {
		row := tab.Row()
		row.Column(item.Type)
		row.Column(item.Name)
		if item.Size != 0 {
			row.Column(fmt.Sprintf("%d", item.Size))
		} else {
			row.Column("-")
		}
	}

	tab.Print(os.Stdout)
	fmt.Println()
}

func (s *StorageStatus) PrettyPrint() {
	fmt.Printf("\nStorage Status\n")

	tab := tabulate.New(tabulate.Unicode)
	tab.Header("Field").SetAlign(tabulate.ML)
	tab.Header("Value").SetAlign(tabulate.ML)

	row := tab.Row()
	row.Column("Used Bytes")
	row.Column(fmt.Sprintf("%d", s.UsedBytes))

	row = tab.Row()
	row.Column("Free Bytes")
	row.Column(fmt.Sprintf("%d", s.FreeBytes))

	row = tab.Row()
	row.Column("Total Bytes")
	row.Column(fmt.Sprintf("%d", s.TotalBytes))

	tab.Print(os.Stdout)
	fmt.Println()
}
