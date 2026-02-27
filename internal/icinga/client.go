package icinga

import (
	"context"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"time"
)

const (
	endpointApiListener             = "/status/ApiListener"
	endpointApplication             = "/status/IcingaApplication"
	endpointCIB                     = "/status/CIB"
	endpointCheckerComponent        = "/status/CheckerComponent"
	endpointCompatLogger            = "/status/CompatLogger"
	endpointElasticsearchWriter     = "/status/ElasticsearchWriter"
	endpointExternalCommandListener = "/status/ExternalCommandListener"
	endpointFileLogger              = "/status/FileLogger"
	endpointGelfWriter              = "/status/GelfWriter"
	endpointGraphiteWriter          = "/status/GraphiteWriter"
	endpointIcingaApplication       = "/status/IcingaApplication"
	endpointIdoMysqlConnection      = "/status/IdoMysqlConnection"
	endpointIdoPgsqlConnection      = "/status/IdoPgsqlConnection"
	endpointInfluxdb2Writer         = "/status/Influxdb2Writer"
	endpointInfluxdbWriter          = "/status/InfluxdbWriter"
	endpointJournaldLogger          = "/status/JournaldLogger"
	endpointLivestatusListener      = "/status/LivestatusListener"
	endpointNotificationComponent   = "/status/NotificationComponent"
	endpointOpenTsdbWriter          = "/status/OpenTsdbWriter"
	endpointPerfdataWriter          = "/status/PerfdataWriter"
	endpointSyslogLogger            = "/status/SyslogLogger"
)

type Config struct {
	BasicAuthUsername string
	BasicAuthPassword string
	CAFile            string
	CertFile          string
	KeyFile           string
	Insecure          bool
	IcingaAPIURI      url.URL
}

type Client struct {
	Client http.Client
	URL    url.URL
}

func NewClient(c Config) (*Client, error) {
	// Create TLS configuration for default RoundTripper
	tlsConfig, err := newTLSConfig(&TLSConfig{
		InsecureSkipVerify: c.Insecure,
		CAFile:             c.CAFile,
		KeyFile:            c.KeyFile,
		CertFile:           c.CertFile,
	})

	if err != nil {
		return nil, err
	}

	var rt http.RoundTripper = &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   10 * time.Second,
			KeepAlive: 10 * time.Second,
		}).DialContext,
		TLSHandshakeTimeout: 10 * time.Second,
		TLSClientConfig:     tlsConfig,
	}

	// Using a BasicAuth for authentication
	if c.BasicAuthUsername != "" {
		rt = newBasicAuthRoundTripper(c.BasicAuthUsername, c.BasicAuthPassword, rt)
	}

	cli := &Client{
		URL: c.IcingaAPIURI,
		Client: http.Client{
			Transport: rt,
		},
	}

	return cli, nil
}

func (icinga *Client) GetApiListenerMetrics() (APIResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	u := icinga.URL.JoinPath(endpointApiListener)

	req, errReq := http.NewRequestWithContext(ctx, http.MethodGet, u.String(), nil)

	var result APIResult

	if errReq != nil {
		return result, fmt.Errorf("error creating request: %w", errReq)
	}

	resp, errDo := icinga.Client.Do(req)

	if errDo != nil {
		return result, fmt.Errorf("error performing request: %w", errDo)
	}

	if resp.StatusCode != http.StatusOK {
		return result, fmt.Errorf("request failed: %s", resp.Status)
	}

	defer resp.Body.Close()

	errDecode := json.NewDecoder(resp.Body).Decode(&result)

	if errDecode != nil {
		return result, fmt.Errorf("error parsing response: %w", errDecode)
	}

	return result, nil
}

func (icinga *Client) GetCIBMetrics() (CIBResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	u := icinga.URL.JoinPath(endpointCIB)

	req, errReq := http.NewRequestWithContext(ctx, http.MethodGet, u.String(), nil)

	var result CIBResult

	if errReq != nil {
		return result, fmt.Errorf("error creating request: %w", errReq)
	}

	resp, errDo := icinga.Client.Do(req)

	if errDo != nil {
		return result, fmt.Errorf("error performing request: %w", errDo)
	}

	if resp.StatusCode != http.StatusOK {
		return result, fmt.Errorf("request failed: %s", resp.Status)
	}

	defer resp.Body.Close()

	errDecode := json.NewDecoder(resp.Body).Decode(&result)

	if errDecode != nil {
		return result, fmt.Errorf("error parsing response: %w", errDecode)
	}

	return result, nil
}

func (icinga *Client) GetApplicationMetrics() (ApplicationResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	u := icinga.URL.JoinPath(endpointApplication)

	req, errReq := http.NewRequestWithContext(ctx, http.MethodGet, u.String(), nil)

	var result ApplicationResult

	if errReq != nil {
		return result, fmt.Errorf("error creating request: %w", errReq)
	}

	resp, errDo := icinga.Client.Do(req)

	if errDo != nil {
		return result, fmt.Errorf("error performing request: %w", errDo)
	}

	if resp.StatusCode != http.StatusOK {
		return result, fmt.Errorf("request failed: %s", resp.Status)
	}

	defer resp.Body.Close()

	errDecode := json.NewDecoder(resp.Body).Decode(&result)

	if errDecode != nil {
		return result, fmt.Errorf("error parsing response: %w", errDecode)
	}

	return result, nil
}
