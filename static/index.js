const connectBtn = document.querySelector('#connect');
const disconnectBtn = document.querySelector('#disconnect');
const sendBtn = document.querySelector('#send');
const connectionStatus = document.querySelector('#status');
const messageInput = document.querySelector('#message');
const messagesList = document.querySelector('#messages');

let socket

function connect() {
    console.log("Attempting Connection...");
    socket = new WebSocket("ws://localhost:8080/ws?name=native");

    socket.onopen = () => {
        console.log("Successfully Connected");
        connectionStatus.innerHTML = "Connected"
    };

    socket.onmessage = msg => {
        console.log(msg);
        const li = document.createElement("li");
        li.innerHTML = msg.data;
        messagesList.appendChild(li);
    };

    socket.onclose = event => {
        console.log("Socket Closed Connection: ", event);
        connectionStatus.innerHTML = "Not connected"
    };

    socket.onerror = error => {
        console.log("Socket Error: ", error);
        connectionStatus.innerHTML = "Not connected. Error: " + error
    };
};

function disconnect() {
    if (socket.bufferedAmount !== 0) {
        alert("Still transfering data");
        return;
    }
    socket.close()
    // socket = undefined;
}

function sendMsg(msg) {
    console.log("sending msg: ", msg);
    socket.send(msg);
};

connectBtn.addEventListener("click", () => {
    connect();
});

disconnectBtn.addEventListener("click", () => {
    disconnect();
});

sendBtn.addEventListener("click", () => {
    const msg = messageInput.value;
    sendMsg(msg);
    messageInput.value = "";
});



