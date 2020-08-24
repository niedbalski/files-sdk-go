package user

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

func (i *Iter) User() files_sdk.User {
	return i.Current().(files_sdk.User)
}

func (c *Client) List(params files_sdk.UserListParams) *Iter {
	params.ListParams.Set(params.Page, params.PerPage, params.Cursor, params.MaxPages)
	i := &Iter{Iter: &lib.Iter{}}
	path := "/users"

	i.Query = func() (*[]interface{}, string, error) {
		data, res, err := files_sdk.Call("GET", c.Config, path, i.ExportParams())
		defaultValue := make([]interface{}, 0)
		if err != nil {
			return &defaultValue, "", err
		}
		list := files_sdk.UserCollection{}
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

func List(params files_sdk.UserListParams) *Iter {
	return (&Client{}).List(params)
}

func (c *Client) Find(params files_sdk.UserFindParams) (files_sdk.User, error) {
	user := files_sdk.User{}
	path := "/users/" + lib.QueryEscape(strconv.FormatInt(params.Id, 10)) + ""
	data, res, err := files_sdk.Call("GET", c.Config, path, lib.ExportParams(params))
	if err != nil {
		return user, err
	}
	if res.StatusCode == 204 {
		return user, nil
	}
	if err := user.UnmarshalJSON(*data); err != nil {
		return user, err
	}

	return user, nil
}

func Find(params files_sdk.UserFindParams) (files_sdk.User, error) {
	return (&Client{}).Find(params)
}

func (c *Client) Create(params files_sdk.UserCreateParams) (files_sdk.User, error) {
	user := files_sdk.User{}
	path := "/users"
	data, res, err := files_sdk.Call("POST", c.Config, path, lib.ExportParams(params))
	if err != nil {
		return user, err
	}
	if res.StatusCode == 204 {
		return user, nil
	}
	if err := user.UnmarshalJSON(*data); err != nil {
		return user, err
	}

	return user, nil
}

func Create(params files_sdk.UserCreateParams) (files_sdk.User, error) {
	return (&Client{}).Create(params)
}

func (c *Client) Unlock(params files_sdk.UserUnlockParams) (files_sdk.User, error) {
	user := files_sdk.User{}
	path := "/users/" + lib.QueryEscape(strconv.FormatInt(params.Id, 10)) + "/unlock"
	data, res, err := files_sdk.Call("POST", c.Config, path, lib.ExportParams(params))
	if err != nil {
		return user, err
	}
	if res.StatusCode == 204 {
		return user, nil
	}
	if err := user.UnmarshalJSON(*data); err != nil {
		return user, err
	}

	return user, nil
}

func Unlock(params files_sdk.UserUnlockParams) (files_sdk.User, error) {
	return (&Client{}).Unlock(params)
}

func (c *Client) ResendWelcomeEmail(params files_sdk.UserResendWelcomeEmailParams) (files_sdk.User, error) {
	user := files_sdk.User{}
	path := "/users/" + lib.QueryEscape(strconv.FormatInt(params.Id, 10)) + "/resend_welcome_email"
	data, res, err := files_sdk.Call("POST", c.Config, path, lib.ExportParams(params))
	if err != nil {
		return user, err
	}
	if res.StatusCode == 204 {
		return user, nil
	}
	if err := user.UnmarshalJSON(*data); err != nil {
		return user, err
	}

	return user, nil
}

func ResendWelcomeEmail(params files_sdk.UserResendWelcomeEmailParams) (files_sdk.User, error) {
	return (&Client{}).ResendWelcomeEmail(params)
}

func (c *Client) User2faReset(params files_sdk.UserUser2faResetParams) (files_sdk.User, error) {
	user := files_sdk.User{}
	path := "/users/" + lib.QueryEscape(strconv.FormatInt(params.Id, 10)) + "/2fa/reset"
	data, res, err := files_sdk.Call("POST", c.Config, path, lib.ExportParams(params))
	if err != nil {
		return user, err
	}
	if res.StatusCode == 204 {
		return user, nil
	}
	if err := user.UnmarshalJSON(*data); err != nil {
		return user, err
	}

	return user, nil
}

func User2faReset(params files_sdk.UserUser2faResetParams) (files_sdk.User, error) {
	return (&Client{}).User2faReset(params)
}

func (c *Client) Update(params files_sdk.UserUpdateParams) (files_sdk.User, error) {
	user := files_sdk.User{}
	path := "/users/" + lib.QueryEscape(strconv.FormatInt(params.Id, 10)) + ""
	data, res, err := files_sdk.Call("PATCH", c.Config, path, lib.ExportParams(params))
	if err != nil {
		return user, err
	}
	if res.StatusCode == 204 {
		return user, nil
	}
	if err := user.UnmarshalJSON(*data); err != nil {
		return user, err
	}

	return user, nil
}

func Update(params files_sdk.UserUpdateParams) (files_sdk.User, error) {
	return (&Client{}).Update(params)
}

func (c *Client) Delete(params files_sdk.UserDeleteParams) (files_sdk.User, error) {
	user := files_sdk.User{}
	path := "/users/" + lib.QueryEscape(strconv.FormatInt(params.Id, 10)) + ""
	data, res, err := files_sdk.Call("DELETE", c.Config, path, lib.ExportParams(params))
	if err != nil {
		return user, err
	}
	if res.StatusCode == 204 {
		return user, nil
	}
	if err := user.UnmarshalJSON(*data); err != nil {
		return user, err
	}

	return user, nil
}

func Delete(params files_sdk.UserDeleteParams) (files_sdk.User, error) {
	return (&Client{}).Delete(params)
}
