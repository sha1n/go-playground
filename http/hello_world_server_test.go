package hellohttp

import "testing"
import "net/http"
import "io/ioutil"

func check(msg string, err error, t *testing.T) {
	if err != nil {
		t.Errorf(msg, err)
	}
}

func TestSimpleGetResponse(t *testing.T) {
	StartHTTPServer(":8080")

	resp, err := http.Get("http://localhost:8080")
	check("Request failed!", err, t)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	check("Failed to read from response", err, t)

	bodyString := string(body)
	t.Log("Got body:", bodyString)

	StopHTTPServer()
}
