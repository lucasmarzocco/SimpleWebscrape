README

Created by: Lucas Marzocco

Email: lucasmarzocco@gmail.com

Took (Approx.): 3 hours

Language: Go

Webserver tool: Swagger

Make sure to have Swagger installed as well as all of the packages required and the apple_proj folder in the SRC folder of your Go installation.

To run:

1) Be in the main directory (apple_proj).
2) Start the server (go run cmd/apple-take-home-project-server/main.go)
	Should see something like: 2019/01/16 19:34:48 Serving apple take home project at http://127.0.0.1:59987
3) Take the local server output and add "/joke" (I just decided to have a specific endpoint).
4) In another window use the curl command:
	
	curl http://127.0.0.1:59987/joke [EXAMPLE]

	It will output a new joke as per the requirements.

If any more explanation or help is needed to get this running, please contact me.




