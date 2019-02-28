README

Created by: Lucas Marzocco

Email: lucasmarzocco@gmail.com

Language: Go

API Framework: Swagger

Make sure to have Swagger installed as well as all of the packages required and the proj folder in the SRC folder of your Go installation.

To run:

Be in the main directory (webscraper).
Start the server (go run cmd/simple-web-scrape-server/main.go)
Take the local server output and add "/get_jobs"
Send a POST request and make sure the body is a list of JSON-formatted URLs.

Example:

POST on http://127.0.0.1:52312/get_jobs

With body:

[
“http://www.indeed.com/viewjob?jk=8cfd54301d909668”,
“http://www.indeed.com/viewjob?jk=b17c354e3cabe4f1”,
“http://www.indeed.com/viewjob?jk=38123d02e67210d9”
]

Output will be:

[
    {
        "company": "Healthify.us",
        "location": "Remote",
        "title": "Backend Engineer",
        "url": "http://www.indeed.com/viewjob?jk=b17c354e3cabe4f1"
    },
    {
        "company": "MuleSoft",
        "location": "San Francisco, CA",
        "title": "Software Engineer",
        "url": "http://www.indeed.com/viewjob?jk=8cfd54301d909668"
    },
    {
        "company": "Intechriti",
        "location": "San Francisco, CA",
        "title": "Java Developer (Consumer Web)",
        "url": "http://www.indeed.com/viewjob?jk=38123d02e67210d9"
    }
]

ASSUMPTIONS: All links are valid links to Indeed jobs. Please don't pass in empty sets and weird data. :) 

Thanks!

Performance:

List of 100 Links: ~2-3 seconds.
