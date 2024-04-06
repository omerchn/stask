watch-compile-json:
	- npx nodemon -q -e go --signal SIGTERM --exec go run . compile example.tasks.md

