var socket = new WebSocket("ws://localhost:8080/ws");

let connect = () => {
    console.log("Attempting Connection...");

    socket.onopen = () => {
        console.log("Succesfuly Connected");
    };

    socket.onmessage = (event: MessageEvent) => {
        console.log("msg", event.data);
    };

    socket.onclose = event => {
        console.log("Connection Closed", event);
    };

    socket.onerror = error => {
        console.log("Socket Error: ", error);
    };

};

// Send a message
let sendMsg = (msg: string) => {
    console.log("Sending message", msg);
    socket.send(msg);
};

export { connect, sendMsg };