import { serve } from "../../deps.ts";
import { handleConnected, handleError, handleMessage } from "./handlers.ts";

import { GetWsPort } from "../helpers/env.ts";

const PORT = GetWsPort();

function logError(msg: string) {
  console.log(msg);
}

function reqHandler(req: Request) {
  if (req.headers.get("upgrade") != "websocket") {
    return new Response(null, { status: 501 });
  }

  const { socket: ws, response } = Deno.upgradeWebSocket(req);

  ws.onopen = () => handleConnected();
  ws.onmessage = (m) => handleMessage(ws, m.data);
  ws.onclose = () => logError("Disconnected from client ...");
  ws.onerror = (e) => handleError(e);

  return response;
}

export default () => {
  serve(reqHandler, { port: PORT });
};
