const BOARD_DIM = 750;
const NUM_SQUARES = 8;
const SQUARE_DIM = BOARD_DIM / NUM_SQUARES;
const SQUARE_COLOR = "green";
const SQUARE_COLOR_ALT = "darkgreen";
const PLAYER_A_COLOR = "white";
const PLAYER_B_COLOR = "black";
const NO_PLAYER_COLOR = false;
let playerName = null;
let playerId = null;
let conn = null;
while (playerName == null) {
    playerName = prompt("Please enter your name", "Harry Potter");
}
postJson("/register", {PlayerName: playerName}, (response) => {
    response.json().then((data) => {
        playerId = data["PlayerId"];
        conn = new WebSocket("ws://" + document.location.host + "/ws/game/" + playerId);
        conn.onclose = (evt) => {
            alert("Disconnected from server, page will now reload");
            window.location.reload();
        };
        conn.onmessage = (evt) => {
            console.log(evt);
        };
        console.log("Registered with id: " + playerId);
    });
});

updatePlayerList();

window.addEventListener('load', () => {
    const canvas = document.getElementById("game");
    canvas.width = BOARD_DIM;
    canvas.height = BOARD_DIM;
    let context = canvas.getContext("2d");
    let board = new Array(NUM_SQUARES);

    for (let x = 0; x < NUM_SQUARES; x++) {
        board[x] = new Array(NUM_SQUARES);
        for (let y = 0; y < NUM_SQUARES; y++) {
            board[x][y] = NO_PLAYER_COLOR;
            const isEven = ((x + y) % 2) < 0.5;
            drawGameboardSquare(x, y, isEven ? SQUARE_COLOR : SQUARE_COLOR_ALT);
        }
    }
    board[3][3] = PLAYER_A_COLOR;
    board[3][4] = PLAYER_B_COLOR;
    board[4][3] = PLAYER_B_COLOR;
    board[4][4] = PLAYER_A_COLOR;
    drawGameboardPeice(3, 3, board[3][3]);
    drawGameboardPeice(3, 4, board[3][4]);
    drawGameboardPeice(4, 3, board[4][3]);
    drawGameboardPeice(4, 4, board[4][4]);

    function drawGameboardSquare(x, y, color) {
        context.fillStyle = color;
        context.fillRect(x * SQUARE_DIM, y * SQUARE_DIM, SQUARE_DIM, SQUARE_DIM);
    }

    function drawGameboardPeice(x, y, color) {
        context.fillStyle = color;
        context.beginPath();
        context.ellipse((x + 0.5) * SQUARE_DIM, (y + 0.5) * SQUARE_DIM, 0.4 * SQUARE_DIM, 0.4 * SQUARE_DIM, Math.PI / 4, 0, 2 * Math.PI, false);
        context.fill();
    }

    canvas.addEventListener("click", function (event) {
        const x = Math.floor(event.x / SQUARE_DIM);
        const y = Math.floor(event.y / SQUARE_DIM);
        if (board[x][y] === false) {
            postJson("/playerMove", {x: x, y: y}, (result) => {
                if (result.ok) {
                    board[x][y] = PLAYER_A_COLOR;
                    drawGameboardPeice(x, y, PLAYER_A_COLOR);
                }
            });
        } else {
            alert("Cannot move there!");
        }
    }, false);
});

function postJson(url, body, handler) {
    fetch(url, {method: "POST", body: JSON.stringify(body)}).then(handler);
}

function updatePlayerList() {
    fetch("/playerList", {method:"GET"}).then((response) => {
        response.json().then((data) => {
            const opponentList = document.getElementById("opponentList");
            opponentList.innerHTML = "";
            data.forEach((player) => {
                if (player.PlayerId !== playerId) {
                    opponentList.innerHTML += "<option>" + player.PlayerName + "</option>"
                }
            })
        })
    })
}