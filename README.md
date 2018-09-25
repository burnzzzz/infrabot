### Infrabot - OpenShift cluster automation


#### Build & Install

```bash
git clone git@github.com:burnzzzz/infrabot.git
make build
sudo make install
```

 ----
 
 #### Usage
 
- CLI:

```bash
$ infrabot
NAME:
   infrabot - easily tune your infrastructure

USAGE:
   infrabot [global options] command [command options] [arguments...]


COMMANDS:
     namespace, ns  operations with namespaces
     help, h        Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help
   --version, -v  print the version

```

```bash
$ infrabot ns create dev
Creating namespace: dev
```

- SlackBot

\#TODO