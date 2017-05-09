package mixmax

import (
	"github.com/ghmeier/go-mixmax/appointmentlinks"
	"github.com/ghmeier/go-mixmax/availability"
	"github.com/ghmeier/go-mixmax/client"
	"github.com/ghmeier/go-mixmax/codesnippet"
	"github.com/ghmeier/go-mixmax/contacts"
)

type Client struct {
	AppointmentLinks *appointmentlinks.Client
	Availability     *availability.Client
	CodeSnippet      *codesnippet.Client
	Contacts         *contacts.Client
}

func New(key string) *Client {
	c := new(Client)
	client := client.New(key)
	c.AppointmentLinks = appointmentlinks.New(client)
	c.Availability = availability.New(client)
	c.CodeSnippet = codesnippet.New(client)
	c.Contacts = contacts.New(client)
	return c
}
