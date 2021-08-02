
function connectToWSS() {
    console.log("connecting the websocket")

    const ws = new WebSocket("ws://127.0.0.1:8080/message-stream")

    // ws.onclose = function(event) {
    //     ws.terminate()
    //     clearTimeout(ws.pingTimeout)
    //     setTimeout(() => {
    //         ws.removeAllListeners()
    //         ws = connectToWSS()
    //     }, autoReconnectDelay)
    // }

    // ws.onerror =  function(err) {
    //     if (err.code === 'ECONNREFUSED') {
    //         ws.removeAllListeners()
    //         ws = (connectToWSS()).ws
    //     }

    //     ws.terminate()
    // }
    
    ws.onmessage = function(message) {
		AddMessage(JSON.parse(message.data))
    }

    return ws
}

function AddMessage(message) {
	var author = message.Author
	var message = message.Message
	var date = message.Date

	console.log("Message is received...");
	console.log(message);
	$('#message-stream').append(`<div class='message border'><p>${author}</p><p>${message}</p><p>${date}</p></div>`)	
}

function SendMessage() {
    var email = document.getElementById("email");
    var message = document.getElementById("message");

    if (!conn) {
        return false;
    }
    if (!msg.value) {
        return false;
    }
    conn.send(msg.value);
    msg.value = "";
    return false;
}