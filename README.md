# Simple Terminal Hacker News

A quick and dirty implementation of a Hacker News terminal tool. It returns a list of stories according to the provided 'sort' flag (default top) along with a link to their corresponding Hacker News threads.

Example output:

```
# ./hn
01 - ACE: Apple Type-C Port Controller Secrets (https://news.ycombinator.com/item?id=25579286)
02 - Visa Advertising Solutions (VAS) Opt Out (https://news.ycombinator.com/item?id=25576113)
03 - C Template Library (https://news.ycombinator.com/item?id=25576466)
04 - Oxford-AstraZeneca coronavirus vaccine approved for use in UK (https://news.ycombinator.com/item?id=25579393)
05 - Pierre Cardin has died (https://news.ycombinator.com/item?id=25569487)
06 - Running BSDs on AMD Ryzen 5000 Series – FreeBSD/Linux Benchmarks (https://news.ycombinator.com/item?id=25580298)
07 - Show HN: Candymail – Email Automation for Node.js (https://news.ycombinator.com/item?id=25578834)
08 - Graph Toy, an interactive graph visualizer using mathematical functions (https://news.ycombinator.com/item?id=25574661)
09 - Messengers of hope: two mRNA Covid-19 vaccines herald a new era for vaccinology (https://news.ycombinator.com/item?id=25577215)
10 - A Generative Grammar for Jazz Chord Sequences (1984) (https://news.ycombinator.com/item?id=25569056)
```

## Build

To build, use the following makefile command:

```
make build
```

## Usage

```
# ./hn
```

Optional flags:
* `-n` - number of stories to return (max 500) (default 10)
* `-sort` - sort to apply to stories (default "top")

Sort options:
* `top` - top stories
* `best` - best stories
* `new` - new stories

