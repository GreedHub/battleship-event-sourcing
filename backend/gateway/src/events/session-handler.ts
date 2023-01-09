import SessionEvents from "../consts/session-events.ts";

const { CREATE_SESSION } = SessionEvents;

export default function handle(event) {
  switch (event.type) {
    case CREATE_SESSION:
      onSessionCreated(event.payload);
      break;
    default:
      console.error(`Unknown Session event ${event.type}`);
  }
}

function onSessionCreated(event) {
  console.log({ event });
}

/* 

{
    "entity": "session",
    "event": {
        "type": "SessionCreated",
        "payload": {
            "owner": 1234
        }
    }
}

*/
