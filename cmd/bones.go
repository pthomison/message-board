package cmd

func bones() MessageBoard {
	return MessageBoard{
		Messages: []Message{
			Message{
				Author:  "John Doe",
				Message: "Hello World This Is A Test",
				Date:    "12-12-12",
			},
			Message{
				Author:  "John Doe",
				Message: "Hello World This Is B Test",
				Date:    "12-12-12",
			},
			Message{
				Author:  "John Doe",
				Message: "Hello World This Is C Test",
				Date:    "12-12-12",
			},
		},
	}
}
