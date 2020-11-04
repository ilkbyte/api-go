package ilkbyte

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/joho/godotenv"
)

type QueryBuilder struct {
	Endpoint string
	Query    map[string]string
}

type Options struct {
	Access  string
	Secret  string
	BaseURL string
}

func (options *Options) new() *Options {
	godotenv.Load()

	options.Access = os.Getenv("ILKBYTE_ACCESSKEY")
	options.Secret = os.Getenv("ILKBYTE_SECRETKEY")
	options.BaseURL = "https://api.ilkbyte.com"

	return options
}

func Request(queryBuilder *QueryBuilder) string {
	options := &Options{}
	options = options.new()

	baseURL, err := url.Parse(options.BaseURL + queryBuilder.Endpoint)

	params := url.Values{}
	params.Add("secret", options.Secret)
	params.Add("access", options.Access)

	if len(queryBuilder.Query) > 0 {
		for k, v := range queryBuilder.Query {
			params.Add(k, v)
		}
	}

	baseURL.RawQuery = params.Encode()

	resp, err := http.Get(baseURL.String())
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	return string(body)
}
