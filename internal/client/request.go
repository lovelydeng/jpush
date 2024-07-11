package client

import "net/http"

func (j *Jpush) DoRequest(req *http.Request) ([]byte, error) {
	resp, err := j.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

}
