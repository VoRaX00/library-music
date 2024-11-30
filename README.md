<h1 align="center">library-music</h1>

## About

This is a test task

## How install

1. We execute the command: 
```sh
   git clone https://github.com/VoRaX00/library-music.git
```
2. Create a file.env with this text: DB_PASSWORD=your password
3. Add the CONFIG_PATH variable to env, which specifies the path to the config.yml file. 

    Or use the command to run: 
```sh 
go run main.go --config=path_do_config.yml
```
4. We execute the command:
```sh
    docker compose build
    docker compose up
```

