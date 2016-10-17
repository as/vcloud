package login

import (
	"crypto/tls"
	"encoding/base64"
	"fmt"
	"net/http"
)

const (
	xml55       = "application/*+xml;version=5.5" // vcloud 5.5 Content-Type
	VCTokenName = "X-Vcloud-Authorization"        // HTTP session header
	LoginURI    = "https://%s/api/sessions"       // The login URI
)

// Do connects to the vCloud server specified in the socket argument
// and logs in with the provided org, user, and password. Do returns 
// either a vCloud session ID, or an empty string with an error set
func Do(socket, org, user, pass string) (string, error) {
	uri := fmt.Sprintf(LoginURI, socket)
	b64auth := mkLogin(org, user, pass)
	client := mkClient()

	rq, _ := http.NewRequest("POST", uri, nil)
	rq.Header.Add("Accept", xml55)
	rq.Header.Add("Authorization", "Basic " + b64auth)

	resp, err := client.Do(rq)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if httpStat := resp.StatusCode; httpStat != 200 {
		httpErr := http.StatusText(httpStat)
		return "", fmt.Errorf("login: HTTP %v: %s", httpStat, httpErr)
	}

	token := resp.Header.Get(VCTokenName)
	if token == "" {
		return "", fmt.Errorf("login: no token recieved")
	}

	return token, nil
}

// mkLogin combines the input strings into a Base64
// vCloud login string. 
func mkLogin(org, user, pass string) string {
	login := []byte(user + "@" + org + ":" + pass)
	return base64.StdEncoding.EncodeToString(login)
}

// mkClient creates an http client. The default behavior
// is to not validate TLS Certificates. TODO: Change this
// for security reasons.
func mkClient() *http.Client {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	return &http.Client{Transport: tr}
}
