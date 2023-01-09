import { Application } from "../deps.ts";
import wsServer from "./websockets/server.ts";
import HealthRouter from "./routes/health.ts";

wsServer();

const app = new Application();

app.use(HealthRouter.routes());

await app.listen(`localhost:5000`);
