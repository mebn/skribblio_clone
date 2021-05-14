const socket = new WebSocket("ws://" + document.location.host + "/ws")

/**
 * Send on a channel name to server. For example:
 * 
 *     sendOn("say hello", "hello")
 * 
 * will send ```hello``` to the receiver ```say hello```
 * on the server.
 * @param {*} name The channel name to send on.
 * @param {*} msg The message/data to send.
 */
const sendOn = (name, msg) => {
    const obj = {
        "code": name,
        "data": msg
    }
    const data = JSON.stringify(obj)

    socket.send(data)
}

/**
 * Receives on a channel name, sent from server.
 * For example:
 * 
 *     receiveOn("say hello", msg, data => {
 *         console.log("A message from the server: " + data)
 *     })
 * 
 * @param {*} name The channel name to send to.
 * @param {*} msg The message/data to send,
 * @param {*} cb A callback function to do something with received data from server.
 */
const receiveOn = (name, msg, cb) => {
    const obj = JSON.parse(msg["data"])
    const code = obj["code"]
    const data = obj["data"]

    if (name == code) {
        cb(data)
    }
}



// Draw.js
function _(selector) {
    return document.querySelector(selector);
}

function setup() {
    let canvas = createCanvas(750, 700);
    canvas.parent("canvas-wrapper");
    background(255);
    strokeWeight(10);
}

let size = 10;
function resize(clicked_id) {
    size = 0;
    if (clicked_id == "one") size = 10;
    else if (clicked_id == "two") size = 20;
    else if (clicked_id == "three") size = 30;
    else size = 40;
}

function mouseDragged() {
    let type = _("#pen-brush").checked ? "brush" : "eraser";
    let color = _("#color-picker").value;

    // Send data needed to draw and then do the actual
    // drawing in socket.onmessage. OBS! there might be some delay
    // on poor network connection.
    let obj = { type, color, size, pmouseX, pmouseY, mouseX, mouseY };
    let data = JSON.stringify(obj);
    sendOn("drawing", data);
}

_("#reset-canvas").addEventListener("click", function () {
    sendOn("clear canvas", true);
});




socket.onopen = () => {
    console.log("Connected to server!")
}

socket.onmessage = msg => {
    receiveOn("username", msg, data => {
        document.getElementById("username").innerHTML = `Username: ${data}`
    })

    receiveOn("roomnumber", msg, data => {
        document.getElementById("roomnumber").innerHTML = `Roomnumber: ${data}`
    })

    receiveOn("sendToRoom", msg, data => {
        document.getElementById("msgs").innerHTML += `${data} <br />`
    })

    receiveOn("drawing", msg, data => {
        const drawData = JSON.parse(data)

        stroke(drawData.color);
        strokeWeight(drawData.size);

        if (drawData.type == "brush") {
            line(drawData.pmouseX, drawData.pmouseY, drawData.mouseX, drawData.mouseY);
        } else {
            stroke(255);
            line(drawData.pmouseX, drawData.pmouseY, drawData.mouseX, drawData.mouseY);
        }
    })

    receiveOn("clear canvas", msg, data => {
        background(255);
    })
}

socket.onclose = err => {
    console.log("Disconnected from server: ", err)
}

socket.onerror = err => {
    console.log("socket error: ", err)
}


document.getElementById("send").addEventListener("click", e => {
    const text = document.getElementById("text")

    if (text.value != "")
        sendOn("sendToRoom", text.value)

    text.value = ""
})
