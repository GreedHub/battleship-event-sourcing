export default function handle(event: any) {
  switch (event.type) {
    case "":
      break;
    default:
      console.error(`Unknown Player event ${event.type}`);
  }
}
