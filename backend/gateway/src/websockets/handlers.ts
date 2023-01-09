import HandleSessionEvent from "../events/session-handler.ts";
import HandlePlayerEvent from "../events/player-handler.ts";
import HandleShipEvent from "../events/ship-handler.ts";

export function handleConnected() {
  console.log("Connected to client ...");
}

export function handleMessage(ws: WebSocket, data: string) {
  let msg;
  try {
    msg = JSON.parse(data);
  } catch {
    console.error(`Cannot parse message ${data} to json`);
  }

  switch (msg.entity) {
    case "ship":
      HandleShipEvent(msg.event);
      break;
    case "session":
      HandleSessionEvent(msg.event);
      break;
    case "player":
      HandlePlayerEvent(msg.event);
      break;
    default:
      console.error(`Unknown entity ${msg.entity}`);
  }
}

export function handleError(e: Event | ErrorEvent) {
  console.log(e instanceof ErrorEvent ? e.message : e.type);
}

export default {
  handleConnected,
  handleError,
  handleMessage,
};
