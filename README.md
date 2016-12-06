=localsdata

Sometimes I know I have a curl command in a certain slack room in a certain slack team.

But trying to use the slack ui to find it is not ideal.

I want the equivilent of `grep -R curl .` so I can easily see a list of the 27 or so curls in that room and then visually I can find the one I need in a few seconds.

Or maybe I want to pipe that information to another process.

Enter lsd:

1. you type ./lsd with no command line parameter to get a list of your slack teams:

  $ ./lsd
  team1
  team2
  anotherteam
  yetanotherteam
  wowstillmoreteams

2. you type ./lsd with a team name to get a list of rooms:

  $ ./lsd
  team1
  team2
