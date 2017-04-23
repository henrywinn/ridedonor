Okay so this is how you get the server started:
Run command: `docker-compose up --build`

That will build the docker image which compiles the source and starts the server. Once it's done, you can go to `localhost:8080/ping` and if it sends something back you're golden. If not, well idk.

Unfortunately, having to rebuild the docker image each time takes a ton of time which I don't like. I'm gonna try to figure out a better way to make this all work so it doesn't take a year and a day to test a code change.
