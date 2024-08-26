package types

import (
	compose "github.com/compose-spec/compose-go/v2/types"
	"github.com/stretchr/testify/require"
	"reflect"
	"testing"
)

func TestParseComposeContainerPorts(t *testing.T) {

	var validConfiguration1 = compose.ServicePortConfig{
		Name:        "http",
		Mode:        "ingress",
		HostIP:      "",
		Target:      uint32(5672),
		Published:   "5672",
		Protocol:    "tcp",
		AppProtocol: "",
		Extensions:  nil,
	}

	var validConfiguration2 = compose.ServicePortConfig{
		Name:        "a-port",
		Mode:        "ingress",
		HostIP:      "0.0.0.0",
		Target:      uint32(80),
		Published:   "8080",
		Protocol:    "udp",
		AppProtocol: "udp",
		Extensions:  nil,
	}

	var validConfiguration3 = compose.ServicePortConfig{
		Name:        "http-metrics",
		Mode:        "ingress",
		HostIP:      "1.2.3.4",
		Target:      uint32(1234),
		Published:   "1234",
		Protocol:    "",
		AppProtocol: "",
		Extensions:  nil,
	}

	var invalidConfiguration1 = compose.ServicePortConfig{
		Name:        "",
		Mode:        "outgress",
		HostIP:      "ABCD",
		Target:      0,
		Published:   "",
		Protocol:    "",
		AppProtocol: "",
		Extensions:  nil,
	}

	var invalidConfiguration2 = compose.ServicePortConfig{
		Name:        "invalid",
		Mode:        "",
		HostIP:      "",
		Target:      uint32(12345),
		Published:   "JKML",
		Protocol:    "SOAP",
		AppProtocol: "",
		Extensions:  nil,
	}
	expectedName := "http"
	expectedContainerPort := float64(5672)
	expectedPublished := float64(5672)
	expectedProtocol := ProtocolTCP

	var expectedConfiguration1 = Port{
		Name:          &expectedName,
		ContainerPort: &expectedContainerPort,
		PublishedPort: &expectedPublished,
		Protocol:      &expectedProtocol,
	}

	t.Run("test valid port configuration", func(t *testing.T) {
		portConfig := []compose.ServicePortConfig{validConfiguration1}
		ports, err := ParseComposeContainerPorts(portConfig)
		require.NoError(t, err)
		assertPortsEqual(t, ports[0], &expectedConfiguration1)
	})

	t.Run("test invalid port configuration", func(t *testing.T) {
		portConfig1 := []compose.ServicePortConfig{invalidConfiguration1}
		_, err := ParseComposeContainerPorts(portConfig1)
		require.Error(t, err)

		portConfig2 := []compose.ServicePortConfig{invalidConfiguration2}
		_, err = ParseComposeContainerPorts(portConfig2)
		require.Error(t, err)
	})

	t.Run("test multiple port configurations", func(t *testing.T) {
		portConfigs := []compose.ServicePortConfig{validConfiguration1, validConfiguration2, validConfiguration3}
		ports, err := ParseComposeContainerPorts(portConfigs)
		require.NoError(t, err)
		require.Len(t, ports, len(portConfigs))
	})

	t.Run("test mixed validity port configurations", func(t *testing.T) {
		portConfigs := []compose.ServicePortConfig{validConfiguration1, validConfiguration2, invalidConfiguration1}
		ports, err := ParseComposeContainerPorts(portConfigs)
		require.Error(t, err)

		if len(ports) != 2 {
			t.Errorf("expected 2 ports, got %d", len(ports))
		}
	})
}

func TestConvertToK8sContainerPorts(t *testing.T) {
	t.Run("test single port configuration", func(t *testing.T) {

	})

	t.Run("test multiple port configuration", func(t *testing.T) {

	})
}

func assertPortsEqual(t testing.TB, got, want *Port) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got: %v, want: %v", got, want)
	}
}
