=localsdata

Sometimes I know I have a curl command in a certain slack room in a certain slack team.

But trying to use the slack ui to find it is not ideal.

I want the equivilent of `grep -R curl .` so I can easily see a list of the 27 or so curls in that room and then visually I can find the one I need in a few seconds.

Or maybe I want to pipe that information to another process.

Enter lsd:

1. you type ./lsd with no command line parameter to get a list of your slack teams:

  ```
  $ ./lsd
  team1
  team2
  anotherteam
  yetanotherteam
  wowstillmoreteams
  ```

2. you type ./lsd with a team name to get a list of rooms:

  $ ./lsd anotherteam

  C093KBABZ test
  C19P7P6CB test-emails
  G2AS2JEG2 project1
  G0A3YDZFD newyork
  G09NX8C22 nbc
  G2AS5UZN3 zoo
  D19P5FBD7 markfreely
  D19NWL7TZ clarkbenson
  D29QLPRE9 timinaccounting

3. you type ./lsd <team> <room>

   $ ./lsd anotherteam G2AS5UZN3

to download every message, every attachment, every file, and every little peice of information about that room in that slack team from the begining of time to now and save it locally somewhere on your hard drive.

4. Now you have all the data locally and you can:

  4a. Do that `grep -R curl .` command now and search just that one room's data.
  4b. Open the files in Finder and visually scan each jpg to find just the one you know is there someplace.
  4c. Come up with other creative greps to narrow your search and find the info you need.
  4d. Avoid having the main slack UI program be in a weird state of scrolled to a message far away in time.
  4e. Avoid having to learn the main slack UI program's search syntax when grep piped to greps is way faster/better.

= INSTALL

1. email author for instructions
