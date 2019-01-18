README

Created by: Lucas Marzocco

Email: lucasmarzocco@gmail.com

Took (Approx.): 3 hours

Language: Go

Testing framework: Vegeta

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


Tests done using Vegeta (https://thisdata.com/blog/load-testing-api-interfaces-with-go-and-vegeta/):
Performance drops a little when there are 70 requests per second but isn't significant until you hit around 100 requests per second.

MacBookPro:~ lucasmarzocco$ echo "GET http://127.0.0.1:51115/joke" | vegeta attack -rate=10 -duration=60s | vegeta report
Requests      [total, rate]            600, 10.02
Duration      [total, attack, wait]    1m0.146191708s, 59.902157s, 244.034708ms
Latencies     [mean, 50, 95, 99, max]  240.251204ms, 233.68879ms, 289.555741ms, 402.515057ms, 455.388708ms
Bytes In      [total, mean]            48297, 80.50
Bytes Out     [total, mean]            0, 0.00
Success       [ratio]                  100.00%
Status Codes  [code:count]             200:600  
Error Set:


MacBookPro:~ lucasmarzocco$ echo "GET http://127.0.0.1:51115/joke" | vegeta attack -rate=30 -duration=60s | vegeta report
Requests      [total, rate]            1800, 30.02
Duration      [total, attack, wait]    1m0.200323843s, 59.967921s, 232.402843ms
Latencies     [mean, 50, 95, 99, max]  257.50625ms, 240.460866ms, 393.682055ms, 485.785413ms, 805.155631ms
Bytes In      [total, mean]            141845, 78.80
Bytes Out     [total, mean]            0, 0.00
Success       [ratio]                  100.00%
Status Codes  [code:count]             200:1800  
Error Set:


MacBookPro:~ lucasmarzocco$ echo "GET http://127.0.0.1:51115/joke" | vegeta attack -rate=50 -duration=60s | vegeta report
Requests      [total, rate]            3000, 50.02
Duration      [total, attack, wait]    1m0.321606387s, 59.981031s, 340.575387ms
Latencies     [mean, 50, 95, 99, max]  276.819189ms, 245.740302ms, 448.513858ms, 565.601161ms, 1.196804465s
Bytes In      [total, mean]            238912, 79.64
Bytes Out     [total, mean]            0, 0.00
Success       [ratio]                  100.00%
Status Codes  [code:count]             200:3000  
Error Set:

MacBookPro:~ lucasmarzocco$ echo "GET http://127.0.0.1:51115/joke" | vegeta attack -rate=70 -duration=60s | vegeta report
Requests      [total, rate]            4200, 70.01
Duration      [total, attack, wait]    1m0.575433067s, 59.989055s, 586.378067ms
Latencies     [mean, 50, 95, 99, max]  481.725965ms, 476.782401ms, 707.818153ms, 1.067711782s, 3.784086889s
Bytes In      [total, mean]            321617, 76.58
Bytes Out     [total, mean]            0, 0.00
Success       [ratio]                  97.10%
Status Codes  [code:count]             0:122  200:4078  
Error Set:


MacBookPro:~ lucasmarzocco$ echo "GET http://127.0.0.1:51115/joke" | vegeta attack -rate=100 -duration=60s | vegeta report
Requests      [total, rate]            6000, 100.01
Duration      [total, attack, wait]    1m0.476695603s, 59.993022s, 483.673603ms
Latencies     [mean, 50, 95, 99, max]  356.012491ms, 471.781285ms, 704.15232ms, 810.312741ms, 1.254137372s
Bytes In      [total, mean]            311886, 51.98
Bytes Out     [total, mean]            0, 0.00
Success       [ratio]                  66.32%
Status Codes  [code:count]             0:2021  200:3979  
Error Set:
Get http://127.0.0.1:51115/joke: EOF





