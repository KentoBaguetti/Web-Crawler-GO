# Web-Crawler-GO

## Web cralwer

A webscrawler that takes in an initial url and and array of keywords, and searches for other links on the page with matching keywords. The crawler then searches
matching links with BFS until certain end conditions are met.

## How it works

![UML Diagram](./images/webcrawler%20v0.png)

1. net/http - fetch pages
2. golang.org/x/net/html - parse html
3. sync - concurrency

### Why I built this and what I learned

I built this project because I was previously only building fullstack webapps, which was starting to get a bit repetitive.
So I decided to build this simple project to get my foot into backend/systems. <br>

I chose to build the project in GO as I found the its goroutines to be a an easy way to start building programs that require concurrency.
It is also lightweight for crawling, as I was able to build the project with simple packages without the requirement of anything heavy like Colly.
Building it with lightweight packages allowed me to understand more of smaller components of crawling. Because of its simplicity, I can also scale it in the future quite easily. <br>

The crawler uses the Breadth-First Search algorithm, that way links that are "closer" to the initial link are searched first. After I'd like to tr out different search algorithms, such as combining BFS with DFS, changing algorithm based on certain conditions. <br>

I also will implement a database to store the scraped links.
