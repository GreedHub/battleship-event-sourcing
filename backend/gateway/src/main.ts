import { Application } from "../deps.ts";
import wsServer from "./websockets/server.ts";
import HealthRouter from "./routes/health.ts";
const PORT = Deno.env.HTTP_PORT || 5000

wsServer();

const app = new Application();

app.use(HealthRouter.routes());

console.log("API listening on http://localhost:5000/");
await app.listen({port: PORT});
