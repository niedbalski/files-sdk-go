package sso_strategy

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

func (i *Iter) SsoStrategy() files_sdk.SsoStrategy {
	return i.Current().(files_sdk.SsoStrategy)
}

func (c *Client) List(params files_sdk.SsoStrategyListParams) (*Iter, error) {
	params.ListParams.Set(params.Page, params.PerPage, params.Cursor, params.MaxPages)
	i := &Iter{Iter: &lib.Iter{}}
	path := "/sso_strategies"
	i.ListParams = &params
	exportParams, err := i.ExportParams()
	if err != nil {
		return i, err
	}
	i.Query = func() (*[]interface{}, string, error) {
		data, res, err := files_sdk.Call("GET", c.Config, path, exportParams)
		defaultValue := make([]interface{}, 0)
		if err != nil {
			return &defaultValue, "", err
		}
		list := files_sdk.SsoStrategyCollection{}
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
	return i, nil
}

func List(params files_sdk.SsoStrategyListParams) (*Iter, error) {
	return (&Client{}).List(params)
}

func (c *Client) Find(params files_sdk.SsoStrategyFindParams) (files_sdk.SsoStrategy, error) {
	ssoStrategy := files_sdk.SsoStrategy{}
	if params.Id == 0 {
		return ssoStrategy, lib.CreateError(params, "Id")
	}
	path := "/sso_strategies/" + strconv.FormatInt(params.Id, 10) + ""
	exportedParams, err := lib.ExportParams(params)
	if err != nil {
		return ssoStrategy, err
	}
	data, res, err := files_sdk.Call("GET", c.Config, path, exportedParams)
	if err != nil {
		return ssoStrategy, err
	}
	if res.StatusCode == 204 {
		return ssoStrategy, nil
	}
	if err := ssoStrategy.UnmarshalJSON(*data); err != nil {
		return ssoStrategy, err
	}

	return ssoStrategy, nil
}

func Find(params files_sdk.SsoStrategyFindParams) (files_sdk.SsoStrategy, error) {
	return (&Client{}).Find(params)
}
