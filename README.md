# rettungsring
backup all your gitlab repos to disk

## requirements

create a personal api token in gitlab profile settings and grant api permissions

## getting started

### Docker

1. `docker build -t rettungsring:latest .`
2. `docker run rettungsring:latest -token <your-token> -url https://gitlab.com/api/v4/ -user <your-username> -listprojects`
3. if you want to download the projects you can omit the `-listprojects` it will create a *repos* folder inside the CWD and stores the repos in there.
4. if you want to have them on your local disk you have to mount a volume into the container with -v /path-local:/data

### build it from source

1. `go build`
2. `./rettungsring -token <your-token> -url https://gitlab.com/api/v4/ -user <your-username> -listprojects`
3. if you want to store them locally omit the `-listprojects` flag
4. it will create a folder named *repos* in your current folder if it does not exist.

## Contributing

All coontributions are welcome

1. fork
2. create a feature branch
3. update the code
4. create a pull request
