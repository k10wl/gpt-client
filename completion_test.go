package gpt_client

import (
	"reflect"
	"testing"
)

func TestBuildHistory(t *testing.T) {
	testCases := []struct {
		name      string
		prevMsgs  []Message
		expected  []Message
		expectErr bool
		tokens    int
	}{
		{
			name: "Messages within tokens limit",
			prevMsgs: []Message{
				{Role: System, Content: "System Message"},
				{Role: User, Content: "User message"},
				{Role: Assistant, Content: "Assistant message"},
			},
			expected: []Message{
				{Role: System, Content: "System Message"},
				{Role: User, Content: "User message"},
				{Role: Assistant, Content: "Assistant message"},
			},
			expectErr: false,
			tokens:    MaxRequestTokens,
		},
		{
			name: "System message with last message within tokens limit",
			prevMsgs: []Message{
				{Role: System, Content: "System Message"},
				{Role: User, Content: "User message1"},
				{Role: Assistant, Content: "Assistant message1"},
				{Role: User, Content: "User message2"},
				{Role: Assistant, Content: "Assistant message2"},
				{Role: User, Content: "User message3"},
				{Role: Assistant, Content: "Assistant message3"},
			},
			expected: []Message{
				{Role: System, Content: "System Message"},
				{Role: Assistant, Content: "Assistant message3"},
			},
			expectErr: false,
			tokens:    5,
		},
	}

	client := NewClient("")

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			client.maxRequestTokens = tc.tokens
			res, err := client.BuildHistory(&tc.prevMsgs)
			if err != nil {
				t.Fatalf("Unexpected error in %v test: %+v\n", tc.name, err)
			}

			equal := reflect.DeepEqual(res, &tc.expected)
			if !equal {
				t.Fatalf("Failed %v. Expected: %+v\nReceived: %+v", tc.name, &tc.expected, res)
			}
		})
	}
}
