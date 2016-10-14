package vcloud

import (
	"compress/gzip"
	"crypto/tls"
	"encoding/base64"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

const (
	xml15             string = "application/*+xml;version=1.5"    // vcloud 1.5 Content-Type
	xml51             string = "application/*+xml;version=5.1"    // vcloud 5.1 Content-Type
	xml55             string = "application/*+xml;version=5.5"    // vcloud 5.5 Content-Type
	VcloudTokenHeader string = "X-Vcloud-Authorization"           // HTTP session header identifier
	loginUriFmt       string = "https://%s/api/sessions"       // The login URI format in the form (host, port)
	orglistUriFmt     string = "https://%s/api/org/"           // Request URL for an OrgList
	queryUriFmt       string = "https://%s/api/query/?type=%s" // Request URL for a Query
)

type Status int

const (
	UNRESOLVED Status = iota - 1
	RESOLVED
	DEPLOYED
	SUSPENDED
	POWERED_ON
	BLOCKED
	UNKNOWN
	UNRECOGNIZED
	POWERED_OFF
	INCONSISTENT
	BAD_CHILDREN
	UPLOAD_INIT
	UPLOAD_COPY
	UPLOAD_DISK
	QUARANTINED
	UNQUARANTINED
)

type Element struct {
	Type string `xml:"type,attr"`
	Name string `xml:"name,attr"`
	Href string `xml:"href,attr"`
}

type Org struct {
	XMLName xml.Name `xml:"Org"`
	Element
}

type OrgList struct {
	XMLName xml.Name `xml:"OrgList"`
	Element
	Orgs []Org `xml:"Org"`
}

type Stringer interface {
	String() string
}

func dateconv(s string) (r string) {
	t, err := time.Parse(time.RFC3339, s)
	t = t.Local()
	if err != nil {
		r = fmt.Sprintf("ERROR_PARSING")
	}
	r = fmt.Sprintf("%04v-%02d-%02v_%02v:%02v:%02v", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
	return r
}

func valid(args ...string) bool {
	for _, v := range args {
		if len(v) == 0 {
			return false
		}
	}
	return true
}

func (s *Session) BytesRx() int64 {
	return s.rx
}

func (s *Session) BytesTx() int64 {
	return s.tx
}

func NewSession(server string, user string) *Session {
	x := strings.Split(user, "@")
	if len(x) < 2 {
		return nil
	}
	y := strings.Split(x[1], ":")
	if len(y) < 1 {
		return nil
	}
	org := y[0]
	return &Session{Server: server, User: user, Org: org}
}

type Session struct {
	Server string
	User   string
	Org string
	Token  string

	client *http.Client
	rx     int64
	tx     int64
}

func check(s *Session) error {
	if s.Server == "" {
		return fmt.Errorf("no server name")
	}
	if s.User == "" {
		return fmt.Errorf("no user info")
	}
	if strings.IndexAny(s.Server, ":") < 0 {
		s.User += ":443"
	}
	return nil
}

func (s *Session) Init() error {
	if check(s) != nil {
		return nil //error
	}

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	s.client = &http.Client{Transport: tr}

	return s.Login()
}

func (s *Session) IsLoggedIn() bool {
	// Local check
	if s.Token == "" {
		return false
	}

	//TODO: Contact vCloud and see if the token is valid
	return true
}

// Function DoRequest is a wrapper for http.NewRequest() and http.Client.Do(). It adds the vCloud Token
// header and the Accept header for vCloud XML content to the request and then runs the request.
func (s *Session) DoRequest(method, urlStr string, body io.Reader) (*http.Response, error) {
	rq, err := http.NewRequest(method, urlStr, body)
	if err != nil {
		return nil, err
	}
	rq.Header.Add(VcloudTokenHeader, s.Token)
	rq.Header.Add("Accept", xml55)
	//request.Header.Add("Accept-Encoding", "gzip, deflate")
	resp, err := s.client.Do(rq)
	if err != nil {
		return nil, err
	}
	//fmt.Println(resp)

	return resp, err
}

func (s *Session) DoRequestGetBody(method, url string, body io.Reader) ([]byte, error) {
	resp, err := s.DoRequest(method, url, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var b []byte

	if resp.Header.Get("Content-Encoding") == "gzip" {
		gz, err := gzip.NewReader(resp.Body)
		if err != nil {
		}
		defer gz.Close()
		b, err = ioutil.ReadAll(gz)
	} else {
		b, err = ioutil.ReadAll(resp.Body)
	}

	if err != nil {
		return nil, err
	}

	return b, nil
}

func (s *Session) LoginParamsOk() (bool, error) {
	if s.IsLoggedIn() {
		return true, nil
	}
	if s.Server == "" {
		return false, errors.New("No host/server provided")
	}
	if s.User == "" {
		return false, errors.New("No port provided")
	}
	return true, nil
}

func (s *Session) Login() error {
	if ok, err := s.LoginParamsOk(); !ok {
		return err
	}
	uri := fmt.Sprintf(loginUriFmt, s.Server)
	request, _ := http.NewRequest("POST", uri, nil)
	request.Header.Add("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte(s.User)))
	request.Header.Add("Accept", xml55)
	resp, err := s.client.Do(request)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return errors.New("Login: HTTP " + string(resp.StatusCode) + "(" + http.StatusText(resp.StatusCode) + ")")
	}

	s.Token = resp.Header.Get(VcloudTokenHeader)
	if !s.IsLoggedIn() {
		return fmt.Errorf("Login: vCloud didn't return a session token")
	}
	return nil
}

func (s *Session) OrgList() (*OrgList, error) {
	uri := fmt.Sprintf(orglistUriFmt, s.Server)
	body, err := s.DoRequestGetBody("GET", uri, nil)
	if err != nil {
		return nil, err
	}

	var o OrgList
	err = xml.Unmarshal(body, &o)
	if err != nil {
		return nil, err
	}
	return &o, nil
}
