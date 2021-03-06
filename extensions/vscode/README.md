
<img src="https://vignette.wikia.nocookie.net/looneytunesshow/images/1/1b/The_Road_Runner.png/revision/latest/scale-to-width-down/340?cb=20121013144306" style="width:50%" align="right" />

# TADA RUNNER GENERATOR

Generator for TADA Runner, simplify you life and avoiding `copas` from existing runner.

## 🚀 Installation

Using AUR for ArchLinux user
```bash
yay -S tada-runner-generator
```

Using Brew for macOS user
```bash
brew tap ayatmaulana/pkg
brew install tada-runner-generator
```

From Golang Source

```bash
go get -u https://github.com/ayatmaulana/tada-runner-generator
```

## 👀 Usage
```bash
aymln@MBP $ tada-runner-generator --help

Usage:
  tada-runner-generator [command]

Available Commands:
  add         Add new runner
  help        Help about any command
  interactive Enter to interactive mode

Flags:
  -h, --help   help for tada-runner-generator

Use "tada-runner-generator [command] --help" for more information about a command.
```


### Manual Mode (Normal CLI)


```bash
cd {$YOUR_RUNNER_PATH}
tada-runner-generator add "birthday-notification" -i -c
```

### Interactive Mode
```bash
tada-runner-generator interactive
```

<img src="https://raw.githubusercontent.com/ayatmaulana/tada-runner-generator/master/img/interactive.gif" />

<!-- ### VS Code Extension -->


## 👨🏻‍💻Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

--



**Requirements**
- Golang 1.11 or above


**Development Tutorial.**

```bash
git clone http://github.com/ayatmaulana/tada-runner-generator
cd tada-runner-generator
make dep
make build
```

## ❤️ License
- MIT - [https://choosealicense.com/licenses/mit](https://choosealicense.com/licenses/mit/)
- Image - [https://looneytunesshow.fandom.com/wiki/Road_Runner](https://looneytunesshow.fandom.com/wiki/Road_Runner)
