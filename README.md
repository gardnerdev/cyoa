# chooseadventure-webapp
Web application which recreates experience from "choose your own adventure" book series.
During reading you have options about how you want to proceed.


Written with the use of https://github.com/mholt/json-to-go tool.


For simplicty, all stories will have a story arc named "intro" that is where the story starts.
That is, every `JSON` file will have a key with the value `intro` and this is where story should start.


Usage:
```
go run cmd/cyoa/main.go
```

Go to: `localhost:3000/story`