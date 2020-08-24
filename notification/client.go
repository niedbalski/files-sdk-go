package notification

import (
	"strconv"

	files_sdk "github.com/Files-com/files-sdk-go"
	lib "github.com/Files-com/files-sdk-go/lib"
)

type Client struct {
	files_sdk.Config
}

type Iter struct {
	*lib.Iter
}

func (i *Iter) Notification() files_sdk.Notification {
	return i.Current().(files_sdk.Notification)
}

func (c *Client) List(params files_sdk.NotificationListParams) *Iter {
	params.ListParams.Set(params.Page, params.PerPage, params.Cursor, params.MaxPages)
	i := &Iter{Iter: &lib.Iter{}}
	path := "/notifications"

	i.Query = func() (*[]interface{}, string, error) {
		data, res, err := files_sdk.Call("GET", c.Config, path, i.ExportParams())
		defaultValue := make([]interface{}, 0)
		if err != nil {
			return &defaultValue, "", err
		}
		list := files_sdk.NotificationCollection{}
		if err := list.UnmarshalJSON(*data); err != nil {
			return &defaultValue, "", err
		}

		ret := make([]interface{}, len(list))
		for i, v := range list {
			ret[i] = v
		}
		cursor := res.Header.Get("X-Files-Cursor")
		return &ret, cursor, nil
	}
	i.ListParams = &params
	return i
}

func List(params files_sdk.NotificationListParams) *Iter {
	return (&Client{}).List(params)
}

func (c *Client) Find(params files_sdk.NotificationFindParams) (files_sdk.Notification, error) {
	notification := files_sdk.Notification{}
	path := "/notifications/" + lib.QueryEscape(strconv.FormatInt(params.Id, 10)) + ""
	data, res, err := files_sdk.Call("GET", c.Config, path, lib.ExportParams(params))
	if err != nil {
		return notification, err
	}
	if res.StatusCode == 204 {
		return notification, nil
	}
	if err := notification.UnmarshalJSON(*data); err != nil {
		return notification, err
	}

	return notification, nil
}

func Find(params files_sdk.NotificationFindParams) (files_sdk.Notification, error) {
	return (&Client{}).Find(params)
}

func (c *Client) Create(params files_sdk.NotificationCreateParams) (files_sdk.Notification, error) {
	notification := files_sdk.Notification{}
	path := "/notifications"
	data, res, err := files_sdk.Call("POST", c.Config, path, lib.ExportParams(params))
	if err != nil {
		return notification, err
	}
	if res.StatusCode == 204 {
		return notification, nil
	}
	if err := notification.UnmarshalJSON(*data); err != nil {
		return notification, err
	}

	return notification, nil
}

func Create(params files_sdk.NotificationCreateParams) (files_sdk.Notification, error) {
	return (&Client{}).Create(params)
}

func (c *Client) Update(params files_sdk.NotificationUpdateParams) (files_sdk.Notification, error) {
	notification := files_sdk.Notification{}
	path := "/notifications/" + lib.QueryEscape(strconv.FormatInt(params.Id, 10)) + ""
	data, res, err := files_sdk.Call("PATCH", c.Config, path, lib.ExportParams(params))
	if err != nil {
		return notification, err
	}
	if res.StatusCode == 204 {
		return notification, nil
	}
	if err := notification.UnmarshalJSON(*data); err != nil {
		return notification, err
	}

	return notification, nil
}

func Update(params files_sdk.NotificationUpdateParams) (files_sdk.Notification, error) {
	return (&Client{}).Update(params)
}

func (c *Client) Delete(params files_sdk.NotificationDeleteParams) (files_sdk.Notification, error) {
	notification := files_sdk.Notification{}
	path := "/notifications/" + lib.QueryEscape(strconv.FormatInt(params.Id, 10)) + ""
	data, res, err := files_sdk.Call("DELETE", c.Config, path, lib.ExportParams(params))
	if err != nil {
		return notification, err
	}
	if res.StatusCode == 204 {
		return notification, nil
	}
	if err := notification.UnmarshalJSON(*data); err != nil {
		return notification, err
	}

	return notification, nil
}

func Delete(params files_sdk.NotificationDeleteParams) (files_sdk.Notification, error) {
	return (&Client{}).Delete(params)
}
