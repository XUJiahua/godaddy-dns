package domain

import (
	"context"
	"errors"
	"fmt"
	"github.com/xujiahua/godaddy-dns/gen/out/godaddy"
	"strings"
)

const defaultTTL = 3600

type Secret struct {
	APIKey    string
	APISecret string
}

type Client struct {
	apiClient *godaddy.APIClient
}

func New(secret Secret) *Client {
	apiClient := godaddy.NewAPIClient(&godaddy.Configuration{
		BasePath: "https://api.godaddy.com",
		DefaultHeader: map[string]string{
			"Authorization": fmt.Sprintf("sso-key %s:%s", secret.APIKey, secret.APISecret),
		},
		UserAgent: "Swagger-Codegen/1.0.0/go",
	})

	return &Client{
		apiClient: apiClient,
	}
}

func (c Client) EnsureRecordSimplified(domainAddr string, type_ string, value string) error {
	parts := strings.Split(domainAddr, ".")
	if len(parts) < 2 {
		return errors.New("invalid domain")
	}
	var name, domain string
	// example.com
	if len(parts) == 2 {
		name = "@"
		domain = domainAddr
	} else {
		// abc.example.com
		name = strings.Join(parts[:len(parts)-2], ".")
		domain = strings.Join(parts[len(parts)-2:], ".")
	}
	return c.EnsureRecord(domain, type_, name, value)
}

func (c Client) EnsureRecord(domain string, type_ string, name string, value string) error {
	record, err := c.GetRecord(domain, type_, name)
	if err != nil {
		return err
	}
	if record == nil {
		return c.AddRecord(domain, type_, name, value)
	}
	if record.Data == value {
		fmt.Printf("keeping: %s.%s => %s\n", name, domain, value)
		return nil
	}
	return c.UpdateRecord(domain, type_, name, value)
}

func (c Client) GetRecord(domain string, type_ string, name string) (*godaddy.DnsRecord, error) {
	records, _, err := c.apiClient.V1Api.RecordGet(context.Background(), domain, type_, name, nil)
	if err != nil {
		return nil, err
	}

	// not exist
	if len(records) == 0 {
		return nil, nil
	}

	return &records[0], nil
}

func (c Client) UpdateRecord(domain string, type_ string, name string, value string) error {
	fmt.Printf("updating: %s.%s => %s\n", name, domain, value)
	body := []godaddy.DnsRecordCreateTypeName{
		{
			Data: value,
			Ttl:  defaultTTL,
		},
	}
	_, err := c.apiClient.V1Api.RecordReplaceTypeName(context.Background(), body, domain, type_, name, nil)
	return err
}

func (c Client) AddRecord(domain string, type_ string, name string, value string) error {
	fmt.Printf("creating: %s.%s => %s\n", name, domain, value)
	body := []godaddy.DnsRecord{
		{
			Name:  name,
			Type_: type_,
			Data:  value,
			Ttl:   defaultTTL,
		},
	}
	_, err := c.apiClient.V1Api.RecordAdd(context.Background(), body, domain, nil)
	return err
}
