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
