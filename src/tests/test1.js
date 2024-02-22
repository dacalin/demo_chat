import ws from 'k6/ws';
import { randomIntBetween } from 'https://jslib.k6.io/k6-utils/1.2.0/index.js';
import  { check } from 'k6';


export const options = {
    // Key configurations for avg load test in this section
    stages: [
        { duration: '30s', target: 1000 }, // stay at 100 users for 30 minutes
    ],
};

export default function () {
    let player, socket;
    let i = randomIntBetween(1, 1000000000)
    player = "player" + i;
    socket = runWebSocket(player);
    check(socket, { 'Connected successfully': (r) => r && r.status === 101 });
}

function runWebSocket(player){
    //console.log("connecting", player)
    const url = 'ws://demo-chat:8080/connect'+'?cid='+player;
    const params = { tags: { my_tag: 'hello' } };

    return ws.connect(url, params, function (socket) {
        socket.on('open', function open(){
            //console.log('connected');
            socket.setInterval(function timeout() {
                //socket.send("hi" + randomIntBetween(1, 1000));
                socket.send("ping");

            }, randomIntBetween(50, 55));
        });
        socket.on('ping', () => console.log('PING!'));
        socket.on('pong', () => console.log('PONG!'));
        socket.on('close', () => console.log(player + 'disconnected'));
        socket.on('error', (e) => {
            if (e.error() != 'websocket: close sent') {
                console.log('An unexpected error occurred: ', e.error());
            }
        });
        socket.on('message', (data) => {
            //console.log("a msg is received", data)
        });
    });
}
