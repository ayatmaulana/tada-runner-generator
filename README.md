
<img src="https://vignette.wikia.nocookie.net/looneytunesshow/images/1/1b/The_Road_Runner.png/revision/latest/scale-to-width-down/340?cb=20121013144306" style="width:50%" align="right" />

# TADA RUNNER GENERATOR

Foobar is a Python library for dealing with word pluralization.

## 🚀 Installation

Use the package manager [pip](https://pip.pypa.io/en/stable/) to install foobar.

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

### VS Code Extension


## 👨🏻‍💻Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## ❤️License
- MIT - [https://choosealicense.com/licenses/mit](https://choosealicense.com/licenses/mit/)
- Image - [https://looneytunesshow.fandom.com/wiki/Road_Runner](https://looneytunesshow.fandom.com/wiki/Road_Runner)
