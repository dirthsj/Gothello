window.addEventListener('load', () => {
    let conn;
    const msg = document.getElementById("msg");
    const log = document.getElementById("log");

    function appendLog(item) {
        const doScroll = log.scrollTop > log.scrollHeight - log.clientHeight - 1;
        log.appendChild(item);
        if (doScroll) {
            log.scrollTop = log.scrollHeight - log.clientHeight;
        }
    }

    const form = document.getElementById("chat");
    function handleSubmit(event) {
        event.preventDefault();
        if (!conn) {
            return false;
        }
        if (!msg.value) {
            return false;
        }
        conn.send(playerName + ": " + msg.value);
        msg.value = "";
        return false;
    }
    form.addEventListener('submit', handleSubmit);

    if (window["WebSocket"]) {
        let websocketPrefix;
        if (location.protocol === "https:" )  {
            websocketPrefix = "wss://"
        } else {
            websocketPrefix = "ws://"
        }
        conn = new WebSocket(websocketPrefix + document.location.host + "/ws/chat");
        conn.onclose = function (evt) {
            var item = document.createElement("div");
            item.innerHTML = "<b>Connection closed.</b>";
            appendLog(item);
        };
        conn.onmessage = function (evt) {
            const messages = evt.data.split('\n');
            for (let i = 0; i < messages.length; i++) {
                const item = document.createElement("div");
                item.innerText = messages[i];
                appendLog(item);
            }
        };
    } else {
        const item = document.createElement("div");
        item.innerHTML = "<b>Your browser does not support WebSockets.</b>";
        appendLog(item);
    }
});