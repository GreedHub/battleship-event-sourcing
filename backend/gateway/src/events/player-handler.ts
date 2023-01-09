export default function handle(event) {
  switch (event.type) {
    case "":
      break;
    default:
      console.error(`Unknown Player event ${event.type}`);
  }
}
