## Session status

| Status                       | By Event                 | Description                                                                      |
|------------------------------|--------------------------|----------------------------------------------------------------------------------|
| WAITING_FOR_GUEST            | SessionCreated           | Create a new game and start waiting for guest player                             |
| WAITING_TO_START             | GuestConnected           | Guest player connected, waiting for owner to start the game                      |
| WAITING_FOR_PLACEMENTS       | MatchStarted             | Game started, waiting for player to place their ships                            |
| WAITING_FOR_OWNER_PLACEMENTS | GuestReady               | Guest placed their ships and is waiting for the owner to finish their placements |
| WAITING_FOR_GUEST_PLACEMENTS | OwnerReady               | Owner placed their ships and is waiting for the guest to finish their placements |
| IN_GAME                      | OwnerReady \| GuestReady | Both players placed their ships and the game is in progress                      |
| FINISHED_GUEST_WON           | GuestWon                 | Game has ended, guest player won                                                 |
| FINISHED_OWNER_WON           | OwnerWon                 | Game has ended, owner player won                                                 |