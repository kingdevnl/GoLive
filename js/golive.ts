import {getLiveComponents} from "./discovery";
import morphdom from "morphdom";

export let socket: WebSocket;


function _debounce(func:any, timeout = 300){
    console.log('aa')
    let timer:any;
    return (...args:any) => {
        clearTimeout(timer);
        timer = setTimeout(() => { func.apply(this, args); }, timeout);
    };
}

function debounce(func:any, target: any) {
    let timeout = target.getAttribute("live-debouce")
    if(!timeout) {
        timeout = 100
    }
    return _debounce(func, timeout)
}


export function send(data: any) {
    socket.send(JSON.stringify(data));
}

function connect() {
    socket = new WebSocket("ws://" + window.location.host + "/ws");
    socket.onopen = function () {
        console.log("Connected to server");


        getLiveComponents().forEach(e => {
            send({
                event: "register",
                type: "register",
                id: e.getAttribute("live-id"),
                component: e.getAttribute("live-component")
            });

            e.querySelectorAll("[live-click]").forEach((clicker: any) => {
                clicker.onclick = debounce(()=> {
                    send({
                        event: clicker.getAttribute("live-click"),
                        type: "click",
                        id: e.getAttribute("live-id"),
                        component: e.getAttribute("live-component")
                    })
                }, clicker)
            })

            e.querySelectorAll("[live-bind]").forEach((input: any) => {
                input.oninput = debounce(()=> {
                    console.log('debounce input')
                    send({
                        event: input.getAttribute("live-bind"),
                        type: "bind",
                        value: (input as any).value,
                        id: e.getAttribute("live-id"),
                        component: e.getAttribute("live-component")
                    })
                }, input)
            })
        })

    };
    socket.onclose = function (event) {
        console.log("Disconnected from server, trying to reconnect");
        setTimeout(connect, 1000);
    };

    socket.onmessage = function (event) {
        let data = JSON.parse(event.data);

        const kind = data.kind;

        if (kind === "action") {
            switch (data.action) {
                case "reload": {
                    window.location.reload();
                    break;
                }
            }
        }
        if (kind === "rerender") {
            const html = data.html;
            morphdom(document.querySelector('[live-id="' + data.id + '"]'), html);
        }

        console.log(data)
    };
}


connect();


