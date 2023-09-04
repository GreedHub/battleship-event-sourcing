import React, { ReactElement } from "react";
import useWebSocket, { ReadyState } from "react-use-websocket";
import Board from "./board";

import "./styles.scss";
const WS_URL = "wss://battleshipgw.lglab.com.ar";

export default function Home(): ReactElement {
  const { sendMessage, readyState } = useWebSocket(WS_URL, {
    onOpen: () => {
      console.log("WebSocket connection established.");
    },
    shouldReconnect: () => true,
  });

  const onTextSend = () => {
    const sessionEvent = {
      entity: "session",
      event: {
        type: "SessionCreated",
        payload: {
          owner: 1234,
        },
      },
    };

    sendMessage(JSON.stringify(sessionEvent));
  };

  return (
    <div className="home">
      <Board className="top" />
      <Board className="bottom" />
      <button onClick={onTextSend} disabled={readyState !== ReadyState.OPEN}>
        Create Session
      </button>
    </div>
  );
}
