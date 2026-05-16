# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Commands

```bash
go run main.go          # run the program (prompts for a URL interactively)
go build ./...          # compile all packages
go vet ./...            # static analysis
```

There are no tests yet. The `llm` package requires a running GitHub Copilot agent accessible via the Copilot SDK — `CreateClient()` will fail without it.

## Architecture

The project is a BFS web crawler with keyword-scored prioritization and an in-progress LLM integration layer.

### Packages

**`scraper`** — two working crawlers plus a WIP third:
- `Crawl` (`scraper.go`): single-threaded BFS using a plain `Queue`. Processes pages sequentially (~25s for 100 pages).
- `ParallelCrawl` (`parallelScraper.go`): worker pool crawl. N workers receive URLs from a buffered `jobs` channel. A dedicated feeder goroutine pops from the `PriorityQueue` and pushes to `jobs`; workers only push new URLs back to the queue (never to `jobs` directly — this separation was the fix for a send/receive deadlock). An `inFlight` counter (protected by `qMux`) acts as a semaphore: the feeder closes `jobs` only when `pq` is empty and `inFlight == 0`.
- `KeywordPriorityCrawler` (`parallelScraper2.go`): stub, not yet implemented.
- `CalculateKeywordScore` scores a URL by counting keyword substring matches (case-insensitive). Used to populate the max-heap so high-relevance URLs surface first.

**`datastructures`** — all three types embed a `sync.Mutex` and are safe for concurrent use:
- `Queue`: slice-backed FIFO for BFS and for tracking crawled URLs.
- `PriorityQueue`: binary max-heap (or min-heap via `CreatePriorityQueue(true)`). Elements are `ScoreValue{Score int, Value string}`. `Pop` does a heap-swap-then-heapifyDown; `Append` does heapifyUp.
- `Set`: `map[string]bool` used as a visited/seen set for O(1) deduplication.

**`llm`** — wraps the GitHub Copilot SDK:
- `ClientSession` / `AgentSession` wrap the SDK's `Client` and `Session` types.
- `CreateClient` → `CreateSessions` → `SendMessageToAgent` is the flow. Events come back asynchronously via a subscription; `session.idle` signals completion.
- `TestFunction` is the current entry point called from `main.go` — it spins up 3 sessions and sends a simple arithmetic prompt to each.
- The TODOs in `ParallelCrawl` outline the planned integration: create a client, instantiate sessions, and attach one session per worker so scraped content can be analyzed by the LLM.

### Data flow (parallel crawl)

```
initialUrl → PriorityQueue (scored)
                ↓ feeder goroutine (pops when jobs channel has space)
             jobs channel (buffered, cap 100)
                ↓ worker goroutines (N workers)
             HTTP fetch → HTML tokenize → extract <a href> links
                ↓ (under qMux lock)
             seen Set (dedup) → PriorityQueue (re-scored new URLs)
```

### Current state of `main.go`

The crawl invocations are commented out. `main` currently only exercises `llm.TestFunction()`. To switch back to crawling, uncomment one of the `scraper.Crawl` or `scraper.ParallelCrawl` calls and comment out the LLM block.
