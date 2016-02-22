package minietcd

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"
	"strings"
	"time"
)

type versionResponse struct {
	EtcdCluster string `json:"etcdcluster"`
	EtcdServer  string `json:"etcdserver"`
}

type readResponse struct {
	Action string `json:"action"`
	Node   struct {
		Node
		Nodes []Node `json:"nodes"`
	} `json:"node"`
}

type Node struct {
	Dir           bool   `json:"dir,omitempty"`
	Key           string `json:"key"`
	Value         string `json:"value,omitempty"`
	CreatedIndex  int    `json:"createdIndex"`
	ModifiedIndex int    `json:"modifiedIndex"`
}

var ErrSupportedVersion = errors.New("minietcd only works with version 2")

type Conn struct {
	_url   string
	client *http.Client
	log    *log.Logger
}

func New() (conn *Conn) {
	conn = new(Conn)
	conn.client = &http.Client{Timeout: 5 * time.Second}
	conn.log = log.New(os.Stdout, "[minietcd] ", log.LstdFlags)

	return conn
}

func (c *Conn) SetLoggingOutput(w io.Writer) {
	c.log.SetOutput(w)
}

func (c *Conn) Dial(_url string) error {
	c._url = _url

	resp, err := c.do("/version")
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var etcdResponse versionResponse
	err = json.NewDecoder(resp.Body).Decode(&etcdResponse)
	if err != nil {
		return err
	}

	serverVersionOk := strings.HasPrefix(etcdResponse.EtcdServer, "2.")
	clusterVersionOk := strings.HasPrefix(etcdResponse.EtcdCluster, "2.")
	if !serverVersionOk || !clusterVersionOk {
		return ErrSupportedVersion
	}

	return nil
}

func (c *Conn) Keys(name string) (kv map[string]string, err error) {
	resp, err := c.do(path.Join("/v2", "keys", name))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// var buf bytes.Buffer
	var etcdResponse readResponse
	// err = json.NewDecoder(io.TeeReader(resp.Body, &buf)).Decode(&etcdResponse)
	err = json.NewDecoder(resp.Body).Decode(&etcdResponse)
	if err != nil {
		// c.log.Println("buffer contents", buf.String())
		// defer buf.Reset()
		return nil, err
	}

	kv = make(map[string]string)
	for _, node := range etcdResponse.Node.Nodes {
		key := strings.TrimPrefix(node.Key, "/"+name+"/")
		kv[key] = node.Value
	}

	return kv, nil
}

func (c *Conn) do(path string) (*http.Response, error) {
	req, err := newRequest(c._url, path)
	if err != nil {
		return nil, err
	}

	c.log.Println("Conn.do GET", req.URL)

	return c.client.Do(req)
}

func newRequest(rawurl, path string) (*http.Request, error) {
	u, err := url.Parse(rawurl)
	if err != nil {
		return nil, err
	}

	u.Path = path
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}
