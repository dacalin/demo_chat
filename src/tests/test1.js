import ws from 'k6/ws';
import { randomIntBetween } from 'https://jslib.k6.io/k6-utils/1.2.0/index.js';
import  { check } from 'k6';


export const options = {
    // Key configurations for avg load test in this section
    stages: [
        { duration: '10s', target: 100 }, // stay at 100 users for 30 minutes
        { duration: '10s', target: 200 }, // stay at 100 users for 30 minutes
        { duration: '20s', target: 500 }, // stay at 100 users for 30 minutes

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
    const url = 'ws://host.docker.internal:8080/connect'+'?cid='+player;
    const params = { tags: { my_tag: 'hello' } };

    return ws.connect(url, params, function (socket) {
        socket.on('open', function open(){
            //console.log('connected');
            socket.setInterval(function timeout() {
                //socket.send("hi" + randomIntBetween(1, 1000));
                socket.send("hello");

            }, randomIntBetween(10000, 10000));
        });
        //socket.on('close', () => console.log(player + ' disconnected'));

        socket.on('message', (data) => {
            //console.log("a msg is received", data)
        });
    });
}
