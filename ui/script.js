function WebSocketTest() {
	if ("WebSocket" in window) {
		alert("WebSocket is supported by your Browser!");

		// Let us open a web socket
		var ws = new WebSocket("ws://127.0.0.1:8080/ping");

		// ws.onopen = function() {
		// 	// Web Socket is connected, send data using send()
		// 	ws.send("Message to send");
		// 	alert("Message is sent...");
		// };

		ws.onmessage = function (evt) { 
			var received_msg = JSON.parse(evt.data);

			var author = received_msg.Author
			var message = received_msg.Message
			var date = received_msg.Date

			console.log("Message is received...");
			console.log(received_msg);
			$('#message-stream').append(`<div class='message border'><p>${author}</p><p>${message}</p><p>${date}</p></div>`)
		};

		// ws.onclose = function() { 
		// 	// websocket is closed.
		// 	alert("Connection is closed..."); 
		// };
	} else {
		// The browser doesn't support WebSocket
		alert("WebSocket NOT supported by your Browser!");
	}
}

WebSocketTest()