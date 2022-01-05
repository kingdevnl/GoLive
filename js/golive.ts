import {getLiveComponents} from "./discovery";
import morphdom from "morphdom";

export let socket: WebSocket;


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
                clicker.onclick = function () {
                    send({
                        event: clicker.getAttribute("live-click"),
                        type: "click",
                        id: e.getAttribute("live-id"),
                        component: e.getAttribute("live-component")
                    })
                }
            })

            e.querySelectorAll("[live-bind]").forEach((input: any) => {
                input.oninput = function () {
                    console.log((event.target as any).value);
                    send({
                        event: input.getAttribute("live-bind"),
                        type: "bind",
                        value: (event.target as any).value,
                        id: e.getAttribute("live-id"),
                        component: e.getAttribute("live-component")
                    })
                    // send({
                    //     event: "event",
                    //     type: "input",
                    //     value: (event.target as any).value,
                    //     name: input.getAttribute("live-input"),
                    //     id: e.getAttribute("live-id"),
                    //     component: e.getAttribute("live-component")
                    // });
                }

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


