# `timedcomputershutdown`

This utility is designed to be used remotely to trigger the shutdown of a Mac computer, with BOOMING audio messages to the operator as the shutdown time gets closer.

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

SIGTERM and friends will cancel the shutdown sequence.

## Why?

To get my son to the dinner table. He will hate this.

## How to run?

Concourse job running on a Kubernetes in the Cloud, that uses Tailscale to be able to SSH to my Mac.

See the `pipeline` directory in this repo for a `ytt` template that will generate a nice pipeline.

I can use Tailscale on my phone to access the Concourse UI and easily trigger the job while I'm in the kitchen preparing dinner.

```bash
# install
go install github.com/continusec/timedcomputershutdown/bin/{timedcomputershutdown,shutdownnow}

# Make the second binary (shutdownnow) owned by root, and setuid - that way the whole thing doesn't need to be root
# and we can still "killall timecomputer shutdown" as non-root.
sudo chown root $HOME/go/bin/shutdownnow
sudo chmod +s $HOME/go/bin/shutdownnow
```

## Screenshots

Find the job for computer and duration you want:

![Screenshot_20220626-202613_2](https://user-images.githubusercontent.com/5984070/175810271-a10527a2-cf3e-4fd6-a3f8-5ca96c30c6e3.png)

Kick off a new "build":

![Screenshot_20220626-202436_2](https://user-images.githubusercontent.com/5984070/175810267-f3813610-7eab-4009-a4c0-b4d01e8aab90.png)

Cancelling concourse build will stop the shutdown:

![Screenshot_20220626-202446_2](https://user-images.githubusercontent.com/5984070/175810308-b089d58e-515f-4585-819c-76256d6083d4.png)

Hard to tell if it succeeds or not, sometimes it shuts down too fast:

![Screenshot_20220626-202702_3](https://user-images.githubusercontent.com/5984070/175810273-d9fe1193-8ab7-49b7-8860-b5ebcb5df696.png)

## Acknowledgements

Thank you TailScale for having an awesomely fun product to allow fun projects like this, and for accepting my PRs on a weekend. :)
