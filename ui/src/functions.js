import Sarus from '@anephenix/sarus';

var $ = require('jquery');



export function ConnectToWSS() {
    console.log("connecting the websocket")

    const sarus = new Sarus({
      url: "ws://127.0.0.1:8080/message-stream",
    });

    sarus.on('message', event => {
        AddMessage(JSON.parse(event.data));
    });
}

export function RegisterFormEvents() {
    $( "#message-form" ).submit(function( event ) {
      event.preventDefault();
      SendMessage()
    });
}

function AddMessage(message) {
	var author = message.Author
	var message = message.Message
	var date = message.Date

	$('#message-stream-messages').append(`<div class='message border'><p>${author}</p><p>${message}</p><p>${date}</p></div>`)	
}

function SendMessage() {
    var email = $( "#email" ).val();
    var message = $( "#message" ).val();

    var obj = {"Author": email, "Message": message};

    // Connect to a WebSocket server
    const ws = new Sarus({
      url: "ws://127.0.0.1:8080/message-receive",
    });

    ws.on('open', () => {
        console.log("connection opened!")
        ws.send(JSON.stringify(obj))
        ws.disconnect()
    });

    return false
 
}