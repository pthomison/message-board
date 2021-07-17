package main

func bones() MessageBoard {
	return MessageBoard{
		Messages: []Message{
			Message{
				Author:  "John Doe",
				Message: "Hello World This Is A Test",
				Date:    "12-12-12",
			},
		},
	}
}
