function WebSocketTest() {
	if ("WebSocket" in window) {
		alert("WebSocket is supported by your Browser!");

		// Let us open a web socket
		var ws = new WebSocket("ws://127.0.0.1:8080/echo");

		ws.onopen = function() {
			// Web Socket is connected, send data using send()
			ws.send("Message to send");
			alert("Message is sent...");
		};

		ws.onmessage = function (evt) { 
			var received_msg = evt.data;
			alert("Message is received...");
			$('#message-stream').append("<div class='message border'><p>Author</p><p>Message</p><p>Today</p></div>")
		};

		ws.onclose = function() { 
			// websocket is closed.
			alert("Connection is closed..."); 
		};
	} else {
		// The browser doesn't support WebSocket
		alert("WebSocket NOT supported by your Browser!");
	}
}

WebSocketTest()