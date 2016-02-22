package minietcd

import (
	"os"
	"testing"
)

const Endpoint = "http://127.0.0.1:4001"

func TestComplete(t *testing.T) {
	conn := New()
	conn.SetLoggingOutput(os.Stderr)

	if err := conn.Dial(Endpoint); err != nil {
		t.Fatalf("failed to connect to endpoint %q, error %q\n", Endpoint, err)
	}

	type pair struct {
		key   string
		value string
	}

	var tests = []struct {
		DirName string
		NumKeys int
		Keys    []pair
	}{
		{
			DirName: "foo",
			NumKeys: 3,
			Keys:    []pair{{"foo", "Hello"}, {"bar", "2"}, {"qux", "4.5"}},
		},
		{
			DirName: "bar/qux",
			NumKeys: 1,
			Keys:    []pair{{"hello", "World"}},
		},
	}

	for idx, tt := range tests {
		keys, err := conn.Keys(tt.DirName)
		if err != nil {
			t.Fatalf("failed to get keys for dir %q, error %q\n", tt.DirName, err)
		}

		numKeys := len(keys)
		if tt.NumKeys != numKeys {
			t.Fatalf("expected %d keys, got %d\n", tt.NumKeys, numKeys)
		}

		for _, pair := range tt.Keys {
			output, ok := keys[pair.key]

			if !ok {
				t.Fatalf("%d. key %q doesn't exist\n", idx+1, pair.key)
			}

			if pair.value != output {
				t.Fatalf("%d. %s.%s expected %q, got %q\n",
					idx+1, tt.DirName, pair.key, pair.value, output)
			}
		}
	}
}
