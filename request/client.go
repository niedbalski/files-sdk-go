package request

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

func (i *Iter) Request() files_sdk.Request {
	return i.Current().(files_sdk.Request)
}

func (c *Client) List(params files_sdk.RequestListParams) *Iter {
	params.ListParams.Set(params.Page, params.PerPage, params.Cursor, params.MaxPages)
	i := &Iter{Iter: &lib.Iter{}}
	path := "/requests"

	i.Query = func() (*[]interface{}, string, error) {
		data, res, err := files_sdk.Call("GET", c.Config, path, i.ExportParams())
		defaultValue := make([]interface{}, 0)
		if err != nil {
			return &defaultValue, "", err
		}
		list := files_sdk.RequestCollection{}
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

func List(params files_sdk.RequestListParams) *Iter {
	return (&Client{}).List(params)
}

func (c *Client) GetFolder(params files_sdk.RequestGetFolderParams) (files_sdk.RequestCollection, error) {
	requestCollection := files_sdk.RequestCollection{}
	path := "/requests/folders/" + lib.QueryEscape(params.Path) + ""
	data, res, err := files_sdk.Call("GET", c.Config, path, lib.ExportParams(params))
	if err != nil {
		return requestCollection, err
	}
	if res.StatusCode == 204 {
		return requestCollection, nil
	}
	if err := requestCollection.UnmarshalJSON(*data); err != nil {
		return requestCollection, err
	}

	return requestCollection, nil
}

func GetFolder(params files_sdk.RequestGetFolderParams) (files_sdk.RequestCollection, error) {
	return (&Client{}).GetFolder(params)
}

func (c *Client) Create(params files_sdk.RequestCreateParams) (files_sdk.Request, error) {
	request := files_sdk.Request{}
	path := "/requests"
	data, res, err := files_sdk.Call("POST", c.Config, path, lib.ExportParams(params))
	if err != nil {
		return request, err
	}
	if res.StatusCode == 204 {
		return request, nil
	}
	if err := request.UnmarshalJSON(*data); err != nil {
		return request, err
	}

	return request, nil
}

func Create(params files_sdk.RequestCreateParams) (files_sdk.Request, error) {
	return (&Client{}).Create(params)
}

func (c *Client) Delete(params files_sdk.RequestDeleteParams) (files_sdk.Request, error) {
	request := files_sdk.Request{}
	path := "/requests/" + lib.QueryEscape(strconv.FormatInt(params.Id, 10)) + ""
	data, res, err := files_sdk.Call("DELETE", c.Config, path, lib.ExportParams(params))
	if err != nil {
		return request, err
	}
	if res.StatusCode == 204 {
		return request, nil
	}
	if err := request.UnmarshalJSON(*data); err != nil {
		return request, err
	}

	return request, nil
}

func Delete(params files_sdk.RequestDeleteParams) (files_sdk.Request, error) {
	return (&Client{}).Delete(params)
}
