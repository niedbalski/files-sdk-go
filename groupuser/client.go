package group_user

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

func (i *Iter) GroupUser() files_sdk.GroupUser {
	return i.Current().(files_sdk.GroupUser)
}

func (c *Client) List(params files_sdk.GroupUserListParams) *Iter {
	params.ListParams.Set(params.Page, params.PerPage, params.Cursor, params.MaxPages)
	i := &Iter{Iter: &lib.Iter{}}
	path := "/group_users"

	i.Query = func() (*[]interface{}, string, error) {
		data, res, err := files_sdk.Call("GET", c.Config, path, i.ExportParams())
		defaultValue := make([]interface{}, 0)
		if err != nil {
			return &defaultValue, "", err
		}
		list := files_sdk.GroupUserCollection{}
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

func List(params files_sdk.GroupUserListParams) *Iter {
	return (&Client{}).List(params)
}

func (c *Client) Update(params files_sdk.GroupUserUpdateParams) (files_sdk.GroupUser, error) {
	groupUser := files_sdk.GroupUser{}
	path := "/group_users/" + lib.QueryEscape(strconv.FormatInt(params.Id, 10)) + ""
	data, res, err := files_sdk.Call("PATCH", c.Config, path, lib.ExportParams(params))
	if err != nil {
		return groupUser, err
	}
	if res.StatusCode == 204 {
		return groupUser, nil
	}
	if err := groupUser.UnmarshalJSON(*data); err != nil {
		return groupUser, err
	}

	return groupUser, nil
}

func Update(params files_sdk.GroupUserUpdateParams) (files_sdk.GroupUser, error) {
	return (&Client{}).Update(params)
}

func (c *Client) Delete(params files_sdk.GroupUserDeleteParams) (files_sdk.GroupUser, error) {
	groupUser := files_sdk.GroupUser{}
	path := "/group_users/" + lib.QueryEscape(strconv.FormatInt(params.Id, 10)) + ""
	data, res, err := files_sdk.Call("DELETE", c.Config, path, lib.ExportParams(params))
	if err != nil {
		return groupUser, err
	}
	if res.StatusCode == 204 {
		return groupUser, nil
	}
	if err := groupUser.UnmarshalJSON(*data); err != nil {
		return groupUser, err
	}

	return groupUser, nil
}

func Delete(params files_sdk.GroupUserDeleteParams) (files_sdk.GroupUser, error) {
	return (&Client{}).Delete(params)
}
