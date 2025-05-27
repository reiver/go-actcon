package actcon_test

import (
	"testing"

	"bytes"

	"github.com/reiver/go-json"

	"github.com/reiver/go-actcon"
)

func TestAnnounceConference_MarshalJSON(t *testing.T) {

	tests := []struct{
		Object actcon.AnnounceConference
		Expected []byte
	}{
		{
			Expected: []byte(`{"type":"Announce"}`),
		},
		{
			Object: actcon.AnnounceConference{},
			Expected: []byte(`{"type":"Announce"}`),
		},



		{
			Object: actcon.AnnounceConference{
				Object: actcon.Conference{
					Actor:"apple",
				},
			},
			Expected: []byte(`{`+
				`"type":"Announce"`+
				`,`+
				`"object":{`+
					`"type":"Conference"`+
					`,`+
					`"actor":"apple"`+
				`}`+
			`}`),
		},
		{
			Object: actcon.AnnounceConference{
				Object: actcon.Conference{
					Actor:"BANANA",
				},
			},
			Expected: []byte(`{`+
				`"type":"Announce"`+
				`,`+
				`"object":{`+
					`"type":"Conference"`+
					`,`+
					`"actor":"BANANA"`+
				`}`+
			`}`),
		},
		{
			Object: actcon.AnnounceConference{
				Object: actcon.Conference{
					Actor:"Cherry",
				},
			},
			Expected: []byte(`{`+
				`"type":"Announce"`+
				`,`+
				`"object":{`+
					`"type":"Conference"`+
					`,`+
					`"actor":"Cherry"`+
				`}`+
			`}`),
		},
		{
			Object: actcon.AnnounceConference{
				Object: actcon.Conference{
					Actor:"dAtE",
				},
			},
			Expected: []byte(`{`+
				`"type":"Announce"`+
				`,`+
				`"object":{`+
					`"type":"Conference"`+
					`,`+
					`"actor":"dAtE"`+
				`}`+
			`}`),
		},



		{
			Object: actcon.AnnounceConference{
				Object: actcon.Conference{
					Actor:"acct:reiver@mastodon.social",
				},
			},
			Expected: []byte(`{`+
				`"type":"Announce"`+
				`,`+
				`"object":{`+
					`"type":"Conference"`+
					`,`+
					`"actor":"acct:reiver@mastodon.social"`+
				`}`+
			`}`),
		},
		{
			Object: actcon.AnnounceConference{
				Object: actcon.Conference{
					Actor:"https://mastodon.social/@reiver",
				},
			},
			Expected: []byte(`{`+
				`"type":"Announce"`+
				`,`+
				`"object":{`+
					`"type":"Conference"`+
					`,`+
					`"actor":"https://mastodon.social/@reiver"`+
				`}`+
			`}`),
		},
		{
			Object: actcon.AnnounceConference{
				Object: actcon.Conference{
					Actor:"https://mastodon.social/users/reiver",
				},
			},
			Expected: []byte(`{`+
				`"type":"Announce"`+
				`,`+
				`"object":{`+
					`"type":"Conference"`+
					`,`+
					`"actor":"https://mastodon.social/users/reiver"`+
				`}`+
			`}`),
		},



		{
			Object: actcon.AnnounceConference{
				Object: actcon.Conference{
					ID:"https://conference.example/@reiver@mastodon.social/conf/1667517459",
				},
			},
			Expected: []byte(`{`+
				`"type":"Announce"`+
				`,`+
				`"object":{`+
					`"type":"Conference"`+
					`,`+
					`"id":"https://conference.example/@reiver@mastodon.social/conf/1667517459"`+
				`}`+
			`}`),
		},



		{
			Object: actcon.AnnounceConference{
				Object: actcon.Conference{
					Name: "@reiver@mastodon.social — Conference",
				},
			},
			Expected: []byte(`{`+
				`"type":"Announce"`+
				`,`+
				`"object":{`+
					`"type":"Conference"`+
					`,`+
					`"name":"@reiver@mastodon.social — Conference"`+
				`}`+
			`}`),
		},



		{
			Object: actcon.AnnounceConference{
				Object: actcon.Conference{
					Origin: []string{
						"https://mastodon.social/users/john_idol",
						"https://mastodon.social/users/reiver",
						"https://mastodon.social/users/soel79",
						"https://onepicaday.com/users/massoud",
					},
				},
			},
			Expected: []byte(`{`+
				`"type":"Announce"`+
				`,`+
				`"object":{`+
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
				`}`+
			`}`),
		},



		{
			Object: actcon.AnnounceConference{
				Object: actcon.Conference{
					StartTime: "2022-11-03T16:17:39.759Z",
				},
			},
			Expected: []byte(`{`+
				`"type":"Announce"`+
				`,`+
				`"object":{`+
					`"type":"Conference"`+
					`,`+
					`"startTime":"2022-11-03T16:17:39.759Z"`+
				`}`+
			`}`),
		},



		{
			Object: actcon.AnnounceConference{
				Object: actcon.Conference{
					Summary: "Hello world!",
				},
			},
			Expected: []byte(`{`+
				`"type":"Announce"`+
				`,`+
				`"object":{`+
					`"type":"Conference"`+
					`,`+
					`"summary":"Hello world!"`+
				`}`+
			`}`),
		},



		{
			Object: actcon.AnnounceConference{
				Object: actcon.Conference{
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
			},
			Expected: []byte(`{`+
				`"type":"Announce"`+
				`,`+
				`"object":{`+
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
				`}`+
			`}`),
		},



		{
			Object: actcon.AnnounceConference{
				Object: actcon.Conference{
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
			},
			Expected: []byte(`{`+
				`"type":"Announce"`+
				`,`+
				`"object":{`+
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
				`}`+
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
