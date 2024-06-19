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

type home struct {
	baseURL *url.URL
	values  map[string]interface{}
	client  *http.Client
}

func New(baseURL string) (Home, error) {
	var err error

	h := &home{
		values: make(map[string]interface{}),
	}

	h.baseURL, err = url.Parse(baseURL)
	if err != nil {
		return nil, err
	}

	h.client = http.DefaultClient

	return h, nil
}

func (h *home) Poll() error {
	response, err := h.client.Get(h.baseURL.String() + "/poll")
	if err != nil {
		return err
	}

	reader := bufio.NewReader(response.Body)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}

		keyAndValue := strings.Split(line, " ")
		if len(keyAndValue) != 2 {
			return fmt.Errorf("invalid line in response: %s", line)
		}
		key := keyAndValue[0]
		stringValue := strings.TrimSpace(keyAndValue[1])

		if strings.HasPrefix(key, "rso/") || strings.HasPrefix(key, "bgs/") || strings.HasPrefix(key, "temp/") ||
			strings.HasPrefix(key, "tsp/") || strings.HasPrefix(key, "otemp") || strings.HasPrefix(key, "rhm") ||
			strings.HasPrefix(key, "wdsp") || strings.HasPrefix(key, "lat") || strings.HasPrefix(key, "long") {
			value, err := strconv.ParseFloat(stringValue, 64)
			if err != nil {
				return err
			}
			h.values[key] = value
		} else if strings.HasPrefix(key, "year") || strings.HasPrefix(key, "month") || strings.HasPrefix(key, "day") ||
			strings.HasPrefix(key, "hour") || strings.HasPrefix(key, "minute") || strings.HasPrefix(key, "second") {
			value, err := strconv.ParseInt(stringValue, 10, 64)
			if err != nil {
				return err
			}
			h.values[key] = value
		} else {
			value, err := strconv.ParseBool(stringValue)
			if err != nil {
				return err
			}
			h.values[key] = value
		}
	}

	return nil
}
