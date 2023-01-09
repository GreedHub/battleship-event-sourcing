import { assertEquals } from "../../deps.ts";
import { GetWsPort } from "./env.ts";

Deno.test("Get Ws Port Without Env", () => {
  Deno.env.set("WS_PORT", "NaN");
  const WS_PORT = GetWsPort();
  assertEquals(WS_PORT, 8000);
});

Deno.test("Get Ws Port With Env", () => {
  Deno.env.set("WS_PORT", "5000");
  const WS_PORT = GetWsPort();
  assertEquals(WS_PORT, 5000);
});
