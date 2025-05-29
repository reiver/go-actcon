package actsock_test

import (
	"testing"

	"bytes"

	"github.com/reiver/go-json"

	"github.com/reiver/go-actsock"
)

func TestConference_MarshalJSON(t *testing.T) {

	tests := []struct{
		Object actsock.Conference
		Expected []byte
	}{
		{
			Expected: []byte(`{"type":"Conference"}`),
		},
		{
			Object: actsock.Conference{},
			Expected: []byte(`{"type":"Conference"}`),
		},



		{
			Object: actsock.Conference{
				Actor:"apple",
			},
			Expected: []byte(`{`+
				`"type":"Conference"`+
				`,`+
				`"actor":"apple"`+
			`}`),
		},
		{
			Object: actsock.Conference{
				Actor:"BANANA",
			},
			Expected: []byte(`{`+
				`"type":"Conference"`+
				`,`+
				`"actor":"BANANA"`+
			`}`),
		},
		{
			Object: actsock.Conference{
				Actor:"Cherry",
			},
			Expected: []byte(`{`+
				`"type":"Conference"`+
				`,`+
				`"actor":"Cherry"`+
			`}`),
		},
		{
			Object: actsock.Conference{
				Actor:"dAtE",
			},
			Expected: []byte(`{`+
				`"type":"Conference"`+
				`,`+
				`"actor":"dAtE"`+
			`}`),
		},



		{
			Object: actsock.Conference{
				Actor:"acct:reiver@mastodon.social",
			},
			Expected: []byte(`{`+
				`"type":"Conference"`+
				`,`+
				`"actor":"acct:reiver@mastodon.social"`+
			`}`),
		},
		{
			Object: actsock.Conference{
				Actor:"https://mastodon.social/@reiver",
			},
			Expected: []byte(`{`+
				`"type":"Conference"`+
				`,`+
				`"actor":"https://mastodon.social/@reiver"`+
			`}`),
		},
		{
			Object: actsock.Conference{
				Actor:"https://mastodon.social/users/reiver",
			},
			Expected: []byte(`{`+
				`"type":"Conference"`+
				`,`+
				`"actor":"https://mastodon.social/users/reiver"`+
			`}`),
		},



		{
			Object: actsock.Conference{
				EndPoints:nil,
			},
			Expected: []byte(`{`+
				`"type":"Conference"`+
			`}`),
		},
		{
			Object: actsock.Conference{
				EndPoints:map[string]string{},
			},
			Expected: []byte(`{`+
				`"type":"Conference"`+
			`}`),
		},
		{
			Object: actsock.Conference{
				EndPoints:map[string]string{
					"inoutbox":"wss://conference.example/@reiver@mastodon.social/conf/1667517459",
				},
			},
			Expected: []byte(`{`+
				`"type":"Conference"`+
				`,`+
				`"endpoints":{`+
					`"inoutbox":"wss://conference.example/@reiver@mastodon.social/conf/1667517459"`+
				`}`+
			`}`),
		},



		{
			Object: actsock.Conference{
				ID:"https://conference.example/@reiver@mastodon.social/conf/1667517459",
			},
			Expected: []byte(`{`+
				`"type":"Conference"`+
				`,`+
				`"id":"https://conference.example/@reiver@mastodon.social/conf/1667517459"`+
			`}`),
		},



		{
			Object: actsock.Conference{
				Name: "@reiver@mastodon.social — Conference",
			},
			Expected: []byte(`{`+
				`"type":"Conference"`+
				`,`+
				`"name":"@reiver@mastodon.social — Conference"`+
			`}`),
		},



		{
			Object: actsock.Conference{
				Origin: []string{
					"https://mastodon.social/users/john_idol",
					"https://mastodon.social/users/reiver",
					"https://mastodon.social/users/soel79",
					"https://onepicaday.com/users/massoud",
				},
			},
			Expected: []byte(`{`+
				`"type":"Conference"`+
				`,`+
				`"origin":[`+
					`"https://mastodon.social/users/john_idol"`+
					`,`+
					`"https://mastodon.social/users/reiver"`+
					`,`+
					`"https://mastodon.social/users/soel79"`+
					`,`+
					`"https://onepicaday.com/users/massoud"`+
				`]`+
			`}`),
		},



		{
			Object: actsock.Conference{
				StartTime: "2022-11-03T16:17:39.759Z",
			},
			Expected: []byte(`{`+
				`"type":"Conference"`+
				`,`+
				`"startTime":"2022-11-03T16:17:39.759Z"`+

			`}`),
		},



		{
			Object: actsock.Conference{
				Summary: "Hello world!",
			},
			Expected: []byte(`{`+
				`"type":"Conference"`+
				`,`+
				`"summary":"Hello world!"`+
			`}`),
		},



		{
			Object: actsock.Conference{
				To: []string{
					"https://example.com/~joeblow",
					"https://example.com/~janedoe",
					"https://mastodon.social/users/john_idol",
					"https://mastodon.social/users/reiver",
					"https://mastodon.social/users/soel79",
					"https://onepicaday.com/users/massoud",
					"https://social.example/~dariush",
					"https://social.example/~malekeh",
				},
			},
			Expected: []byte(`{`+
				`"type":"Conference"`+
				`,`+
				`"to":[`+
					`"https://example.com/~joeblow"`+
					`,`+
					`"https://example.com/~janedoe"`+
					`,`+
					`"https://mastodon.social/users/john_idol"`+
					`,`+
					`"https://mastodon.social/users/reiver"`+
					`,`+
					`"https://mastodon.social/users/soel79"`+
					`,`+
					`"https://onepicaday.com/users/massoud"`+
					`,`+
					`"https://social.example/~dariush"`+
					`,`+
					`"https://social.example/~malekeh"`+
				`]`+
			`}`),
		},



		{
			Object: actsock.Conference{
				Actor:"https://mastodon.social/users/reiver",
				ID:"https://conference.example/@reiver@mastodon.social/conf/1667517459",
				Name: "@reiver@mastodon.social — Conference",
				Origin: []string{
					"https://mastodon.social/users/john_idol",
					"https://mastodon.social/users/reiver",
					"https://mastodon.social/users/soel79",
					"https://onepicaday.com/users/massoud",
				},
				StartTime: "2022-11-03T16:17:39.759Z",
				Summary: "Hello world!",
				To: []string{
					"https://example.com/~joeblow",
					"https://example.com/~janedoe",
					"https://mastodon.social/users/john_idol",
					"https://mastodon.social/users/reiver",
					"https://mastodon.social/users/soel79",
					"https://onepicaday.com/users/massoud",
					"https://social.example/~dariush",
					"https://social.example/~malekeh",
				},
			},
			Expected: []byte(`{`+
				`"type":"Conference"`+
				`,`+
				`"actor":"https://mastodon.social/users/reiver"`+
				`,`+
				`"id":"https://conference.example/@reiver@mastodon.social/conf/1667517459"`+
				`,`+
				`"name":"@reiver@mastodon.social — Conference"`+
				`,`+
				`"origin":[`+
					`"https://mastodon.social/users/john_idol"`+
					`,`+
					`"https://mastodon.social/users/reiver"`+
					`,`+
					`"https://mastodon.social/users/soel79"`+
					`,`+
					`"https://onepicaday.com/users/massoud"`+
				`]`+
				`,`+
				`"startTime":"2022-11-03T16:17:39.759Z"`+
				`,`+
				`"summary":"Hello world!"`+
				`,`+
				`"to":[`+
					`"https://example.com/~joeblow"`+
					`,`+
					`"https://example.com/~janedoe"`+
					`,`+
					`"https://mastodon.social/users/john_idol"`+
					`,`+
					`"https://mastodon.social/users/reiver"`+
					`,`+
					`"https://mastodon.social/users/soel79"`+
					`,`+
					`"https://onepicaday.com/users/massoud"`+
					`,`+
					`"https://social.example/~dariush"`+
					`,`+
					`"https://social.example/~malekeh"`+
				`]`+
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
