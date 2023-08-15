<p align="center">
 <img src="https://github.com/jeronimo3875br/rikka/blob/master/assets/rikka_main.gif" alt="rikka_main" width="300"/>
</p>

# Introduction
Rikka is a CLI developed for monitoring and automatic reloading of applications, similar to solutions like [nodemon](https://github.com/remy/nodemon). However, Rikka has been designed to support any programming language, eliminating the need to manually restart your applications or install different packages for various technologies.

# Installation
To install Rikka, you need to have GO installed on your machine. You can find installation instructions for the latest version at: [https://go.dev/dl/](https://go.dev/dl/)

After installing GO, you can use [GIT](https://git-scm.com/) to download Rikka to your machine by running the following command in your terminal:

```sh 
git clone jeronimo3875br/rikka
```

Make sure to adjust the path according to the desired location on your machine.

## Compile 
After having Rikka on your machine, you can compile it into an executable using the command:

```sh
go build -o rikka main.go
```

After compilation, an executable file will be generated, which you can move to your preferred folder.

It's recommended to add the path to the executable and register it in the **PATH** environment variables of your system for easy access.

Make sure to adjust the path of the executable according to the location of the compiled file on your machine.

## How to Use
You can check all available commands and flags by executing and passing the **"--help"** or **"-h"** flag. This flag works globally, and you can use it anywhere in the CLI. For example:

```sh
./rikka --help
```

Or, if configured in the environment variables, simply:
 
 ```sh
 rikka --help
```

## Monitoring Files and Auto Reload
To start monitoring a project, checking for changes, and performing automatic reloads of the execution, you can use the subcommand **watch**. This command accepts some flags for configuration:

- **"--path"** or **"-p"**: Set the path to your project, with the default value of **"./"**, which corresponds to the current directory.
- **"--reflect"** or **"-r"**: Specifies the action (command) to be executed after identifying a change in the project **(required)**.
- **"--ignore"** or **"-i"**: Allows you to define folders and files to be ignored; in other words, changes to these won't trigger any action.

The structure to use this command is: `rikka watch --path /your/program/path --reflect "reload command"`

Here are some usage examples:

```sh
# Python
rikka watch --path C:\Users\Jeronimo\projects\web_scraping --reflect "python main.py"

# JavaScript (ignoring "node_modules" folder)
rikka watch --path C:\Users\Jeronimo\projects\api --reflect "node bin/server.js" --ignore node_modules

# Golang (ignoring "go.mod" and "README.md")
rikka watch --path C:\Users\Jeronimo\projects\cli --reflect "go run cmd/main.go" --ignore go.mod,README.md
```

# License
Rikka is released under the MIT license. See: <a href="https://github.com/jeronimo3875br/rikka/blob/master/LICENSE">License</a>
