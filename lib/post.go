package lib

import (
	"net/http"
	"net/url"
)

func PostFormTo(apiUrl string, formKey string, formValue string) (*http.Response, error) {
	form := url.Values{formKey: {formValue}}
	res, err := http.PostForm(apiUrl, form)
	if err != nil {
		return nil, err
	}

	return res, nil
}
