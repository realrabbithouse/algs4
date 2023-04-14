package xmlrpc

import (
	"github.com/stretchr/testify/require"
	"testing"
)

const faultResponseXML = `
<?xml version="1.0" encoding="UTF-8"?>
<methodResponse>
  <fault>
    <value>
      <struct>
        <member>
          <name>faultString</name>
          <value>
            <string>You must log in before using this part of Bugzilla.</string>
          </value>
        </member>
        <member>
          <name>faultCode</name>
          <value>
            <int>401</int>
          </value>
        </member>
      </struct>
    </value>
  </fault>
</methodResponse>`

func TestFaultError_Error(t *testing.T) {
	resp := rawResponse(faultResponseXML)
	err := resp.Fault()
	require.Error(t, err)

	fault, ok := resp.Fault().(*FaultError)
	require.True(t, ok)
	require.Equal(t, 401, fault.Code)
	require.Equal(t, "You must log in before using this part of Bugzilla.", fault.String)
}
