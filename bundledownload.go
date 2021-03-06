package files_sdk

import (
	"encoding/json"
	"time"

	lib "github.com/Files-com/files-sdk-go/lib"
)

type BundleDownload struct {
	BundleRegistration string    `json:"bundle_registration,omitempty"`
	DownloadMethod     string    `json:"download_method,omitempty"`
	Path               string    `json:"path,omitempty"`
	CreatedAt          time.Time `json:"created_at,omitempty"`
}

type BundleDownloadCollection []BundleDownload

type BundleDownloadListParams struct {
	Cursor               string `url:"cursor,omitempty" required:"false"`
	PerPage              int    `url:"per_page,omitempty" required:"false"`
	BundleId             int64  `url:"bundle_id,omitempty" required:"false"`
	BundleRegistrationId int64  `url:"bundle_registration_id,omitempty" required:"false"`
	lib.ListParams
}

func (b *BundleDownload) UnmarshalJSON(data []byte) error {
	type bundleDownload BundleDownload
	var v bundleDownload
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*b = BundleDownload(v)
	return nil
}

func (b *BundleDownloadCollection) UnmarshalJSON(data []byte) error {
	type bundleDownloads []BundleDownload
	var v bundleDownloads
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*b = BundleDownloadCollection(v)
	return nil
}
