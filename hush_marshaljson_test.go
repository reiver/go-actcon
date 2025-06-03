package actsock_test

import (
	"testing"

	"bytes"

	"github.com/reiver/go-json"

	"github.com/reiver/go-actsock"
)

func TestHush_MarshalJSON(t *testing.T) {

	tests := []struct{
		Object actsock.Hush
		Expected []byte
	}{
		{
			Expected: []byte(`{"type":"Hush","@value":false}`),
		},
		{
			Object: actsock.Hush{},
			Expected: []byte(`{"type":"Hush","@value":false}`),
		},



		{
			Object: actsock.Hush{
				Value:false,
			},
			Expected: []byte(`{`+
				`"type":"Hush"`+
				`,`+
				`"@value":false`+
			`}`),
		},
		{
			Object: actsock.Hush{
				Value:true,
			},
			Expected: []byte(`{`+
				`"type":"Hush"`+
				`,`+
				`"@value":true`+
			`}`),
		},
	}

	for testNumber, test := range tests {

		actual, err := json.Marshal(test.Object)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
			t.Logf("ERROR: %s", err)
			t.Logf("OBJECT: (%T) %#v", test.Object, test.Object)
			continue
		}

		expected := test.Expected

		if !bytes.Equal(expected, actual) {
			t.Errorf("For test #%d, the actual 'object' (of type %T) is not what was expected.", testNumber, test.Object)
			t.Logf("EXPECTED: (%d)\n%s", len(expected), expected)
			t.Logf("ACTUAL:   (%d)\n%s", len(actual), actual)
			t.Logf("OBJECT: (%T) %#v", test.Object, test.Object)
			continue
		}
	}
}
