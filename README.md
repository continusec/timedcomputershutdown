# `timedcomputershutdown`

This utility is designed to be used remotely to trigger the shutdown of a computer, with audio messages to the operator (if running Mac) as the shutdown time gets closer.

e.g.

```bash
timedcomputershutdown 20m
```

will result in the following announcements:

```
Computer will turn off in 20 minutes
15 minute warning
10 minute warning
5 minute warning
1 minute warning
50 seconds
40 seconds
30 seconds
25 seconds
20 seconds
15 seconds
10
9
8
7
6
5
4
3
2
1
0
Shutting down
```

## Why?

To get my son to the dinner table. He will hate this.

## How to run?

Concourse job running on a Kubernetes in the Cloud, that uses Tailscale to be able to SSh to my Mac.

I can use Tailscale on my phone to access the Concourse UI and easily trigger the job while I'm in the kitchen preparing dinner.

```bash
# install
go install github.com/continusec/timedcomputershutdown/bin/{timedcomputershutdown,shutdownnow}

# Make the second binary (shutdownnow) owned by root, and setuid - that way the whole thing doesn't need to be root
# and we can still "killall timecomputer shutdown" as non-root.
sudo chown root $HOME/go/bin/shutdownnow
sudo chmod +s $HOME/go/bin/shutdownnow
```