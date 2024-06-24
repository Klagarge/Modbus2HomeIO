package homeio

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// home implements the Home interface by connecting to the REST server disposed by the Home I/O game.
type home struct {
	baseURL *url.URL
	inputs  map[string]interface{}
	outputs map[string]interface{}
	client  *http.Client
}

// New creates a new Home instance that can be used to interact with the Home I/O game.
func New(baseURL string) (Home, error) {
	var err error

	// Create the home instance.
	h := &home{
		inputs:  make(map[string]interface{}),
		outputs: make(map[string]interface{}),
	}

	// Parse the base URL.
	h.baseURL, err = url.Parse(baseURL)
	if err != nil {
		return nil, err
	}

	// Create the HTTP client.
	h.client = http.DefaultClient

	return h, nil
}

func (h *home) Poll() error {

	// Send a GET request to the REST server.
	response, err := h.client.Get(h.baseURL.String() + "/poll")
	if err != nil {
		return err
	}
	defer response.Body.Close()

	reader := bufio.NewReader(response.Body)
	for {

		// Read the response line by line.
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}

		// Parse the line. The format is "key value".
		keyAndValue := strings.Split(line, " ")
		if len(keyAndValue) != 2 {
			return fmt.Errorf("invalid line in response: %s", line)
		}
		key := keyAndValue[0]
		stringValue := strings.TrimSpace(keyAndValue[1])

		// Convert the value to the appropriate data type and save it to the inputs map.
		if strings.HasPrefix(key, "rso/") || strings.HasPrefix(key, "bgs/") || strings.HasPrefix(key, "temp/") ||
			strings.HasPrefix(key, "tsp/") || strings.HasPrefix(key, "otemp") || strings.HasPrefix(key, "rhm") ||
			strings.HasPrefix(key, "wdsp") || strings.HasPrefix(key, "lat") || strings.HasPrefix(key, "long") {
			value, err := strconv.ParseFloat(stringValue, 64)
			if err != nil {
				return err
			}
			h.inputs[key] = value
		} else if strings.HasPrefix(key, "year") || strings.HasPrefix(key, "month") || strings.HasPrefix(key, "day") ||
			strings.HasPrefix(key, "hour") || strings.HasPrefix(key, "minute") || strings.HasPrefix(key, "second") {
			value, err := strconv.ParseInt(stringValue, 10, 64)
			if err != nil {
				return err
			}
			h.inputs[key] = value
		} else {
			value, err := strconv.ParseBool(stringValue)
			if err != nil {
				return err
			}
			h.inputs[key] = value
		}
	}

	return nil
}
