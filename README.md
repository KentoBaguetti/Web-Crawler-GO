# Web-Crawler-GO

## Web cralwer made with Go and Colly

Given a starting link, the web crawler will search the given page for other links with specific keywords. Then search those links with Breath-First Search

![UML Diagram](./images/webcrawler%20v0.png)

1. net/http - fetch pages
2. net/url - normalize links
3. golang.org/x/net/html - parse html
4. sync + time - concurrency
5. context - cancel on deadline
6. bufio - parse robots.txt
