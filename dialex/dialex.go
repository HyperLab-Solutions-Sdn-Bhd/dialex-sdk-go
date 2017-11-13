package dialex

import (
        "net/http"
        "io/ioutil"
        "time"
        "encoding/json"
)

type Dial struct {
        apikey string
}

// Initialise the Dialex interface. Takes the API key as a parameter.
func New(key string) (d *Dial) {
        return &Dial {
                apikey: key,
        }
}

// Transform a given string. Takes the string to be transformed
// and language.
func (d *Dial) Transform(data, lang string) (res string, err error) {
        var client = &http.Client{
                Timeout: time.Second * 10,
        }

        req, _ := http.NewRequest("GET", "https://dialexherok.herokuapp.com/api/v1/process", nil)

        q := req.URL.Query()
        q.Add("apikey", d.apikey)
        q.Add("data", data)
        q.Add("lang", lang)

        req.URL.RawQuery = q.Encode()
        resp, err := client.Do(req)

        if err != nil {
                return "Request Error", err
        }

        defer resp.Body.Close()
        body, parse_err := ioutil.ReadAll(resp.Body)

        if parse_err != nil {
                return "Parsing Error", parse_err
        }

        api_resp := make(map[string]interface{})
        json_parse_err := json.Unmarshal(body, &api_resp)

        if json_parse_err != nil {
                return "JSON Parse Error", json_parse_err
        }

        transformed := api_resp["output"].(map[string]interface{})

        if out, ok := transformed["result"].(string); ok {
                return out, nil
        }

        return
}
