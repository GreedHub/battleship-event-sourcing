import { Router } from "../../deps.ts";

const HealthRouter = new Router();

HealthRouter.get("/health", ({ response }) => {
  response.status = 200;
  response.body = { msg: "healthy" };
  return;
});

export default HealthRouter;
