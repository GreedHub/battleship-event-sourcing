import { serve } from "https://deno.land/std/http/mod.ts";
import { handleConnected, handleError, handleMessage } from "./handlers.ts";

function logError(msg: string) {
  console.log(msg);
}

async function reqHandler(req: Request) {
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

export default function Connect() {
  console.log("Waiting for client ...");
  serve(reqHandler, { port: 8000 });
}
