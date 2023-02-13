import { config } from "../../deps.ts";

const configData = await config();

export const GetWsPort = (): number => {
  const DEFAULT_PORT = 8000;
  const ENV_PORT = configData.WS_PORT;
  const PORT = parseInt(ENV_PORT) || DEFAULT_PORT;
  return PORT;
};

export default {
  GetWsPort,
};
