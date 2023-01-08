import { config } from "https://deno.land/std/dotenv/mod.ts";

const configData = await config();

export const GetWsPort = (): number => {
  const DEFAULT_PORT = 8000;
  const ENV_PORT = configData.PORT;
  let PORT = parseInt(ENV_PORT) || DEFAULT_PORT;
  return PORT;
};

export default {
  GetWsPort,
};
