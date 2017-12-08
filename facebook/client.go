package facebook

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/kusold/facebreak/config"
)

const facebookAPI = "https://graph.facebook.com/v2.11/me"

// Client is the facebook client
type Client struct {
	config *config.Config
}

// NewClient returns a new instance of the facebook client
func NewClient(config *config.Config) *Client {
	return &Client{config: config}
}

// Run stuff
func (c *Client) Run() error {
	c.downloadTaggedPhotos()
	return nil
}

func (c *Client) fetch(fields string) (*Response, error) {
	v := url.Values{}
	v.Set("fields", fields)
	v.Set("access_token", *c.config.AccessToken)
	url := facebookAPI + "?" + v.Encode()
	fmt.Println(url)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	response := Response{}
	err = json.NewDecoder(resp.Body).Decode(&response)
	return &response, err
}

func (c *Client) downloadTaggedPhotos() error {
	result, err := c.fetch("id,name,photos.limit(1){tags,images,name,id,created_time}")
	if err != nil {
		return err
	}

	printJSON(result)
	c.downloadPhotos(result.Photos.Data)
	return nil
}

func (c *Client) downloadPhotos(p []*photo) error {

	for _, photo := range p {
		// TODO: Find Max photo
		img := photo.Images[0].Source
		fmt.Println(img)
		name := formatPhotoName(photo.ID, photo.CreatedTime)
		fetchImage(name, img)
		// TODO: Write exif data
	}
	return nil
}

func formatPhotoName(id string, date time.Time) string {
	// TODO: Detect file type here
	return date.UTC().Format(time.RFC3339) + "_" + id
}

func fetchImage(name, source string) {
	resp, err := http.Get(source)
	if err != nil {
		log.Panic(err)
	}
	defer resp.Body.Close()

	file, err := os.Create("/tmp/" + name + ".jpg")
	if err != nil {
		log.Panic(err)
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		log.Panic(err)
	}
}

func printJSON(object interface{}) error {
	marshalled, err := json.Marshal(object)
	if err != nil {
		return err
	}
	var buffer bytes.Buffer
	buffer.Write(marshalled)
	fmt.Println(buffer.String())
	return nil
}
