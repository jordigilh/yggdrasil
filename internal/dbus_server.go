package yggdrasil

import (
	"fmt"
	"io/ioutil"

	"github.com/godbus/dbus/v5"
	"github.com/godbus/dbus/v5/introspect"
	yggdrasil "github.com/redhatinsights/yggdrasil/pkg"
)

// DBusServer serves yggdrasil functionality over the system D-Bus.
type DBusServer struct {
	client  *yggdrasil.HTTPClient
	conn    *dbus.Conn
	xmlData []byte
}

// NewDBusServer creates a new server. The provided HTTP client will be used for
// all HTTP requests.
func NewDBusServer(client *yggdrasil.HTTPClient, dbusInterfaceFile string) (*DBusServer, error) {
	xmlData, err := ioutil.ReadFile(dbusInterfaceFile)
	if err != nil {
		return nil, err
	}
	return &DBusServer{
		client:  client,
		xmlData: xmlData,
	}, nil
}

// Connect opens a connection to the system bus, exports the server as an object
// on the bus, and requests the well-known name "com.redhat.yggdrasil".
func (s *DBusServer) Connect() error {
	conn, err := dbus.SystemBus()
	if err != nil {
		return err
	}
	s.conn = conn

	s.conn.Export(s, "/com/redhat/yggdrasil", "com.redhat.yggdrasil1")
	s.conn.Export(introspect.Introspectable(s.xmlData),
		"/com/redhat/yggdrasil",
		"org.freedesktop.DBus.Introspectable")

	reply, err := s.conn.RequestName("com.redhat.yggdrasil1", dbus.NameFlagDoNotQueue)
	if err != nil {
		return err
	}
	if reply != dbus.RequestNameReplyPrimaryOwner {
		return fmt.Errorf("failed to request name '%v'", "com.redhat.yggdrasil1")
	}
	return nil
}

// Close closes the connection.
func (s *DBusServer) Close() error {
	return s.conn.Close()
}

// Upload calls yggdrasil.Upload and returns the result.
func (s *DBusServer) Upload(file string, collector string, metadata map[string]interface{}) (string, *dbus.Error) {
	var facts *yggdrasil.CanonicalFacts
	if len(metadata) > 0 {
		var err error
		facts, err = yggdrasil.CanonicalFactsFromMap(metadata)
		if err != nil {
			return "", &dbus.Error{
				Name: err.Error(),
			}
		}
	}

	requestID, err := yggdrasil.Upload(s.client, file, collector, facts)
	if err != nil {
		return "", &dbus.Error{
			Name: err.Error(),
		}
	}
	return requestID, nil
}
