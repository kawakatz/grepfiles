# grepfiles🔍
<p align="center">
<a href="https://opensource.org/licenses/MIT"><img src="https://img.shields.io/badge/license-MIT-_red.svg"></a>
<a href="https://github.com/kawakatz/grepfiles/issues"><img src="https://img.shields.io/badge/contributions-welcome-brightgreen.svg?style=flat"></a>
<a href="https://goreportcard.com/badge/github.com/kawakatz/grepfiles"><img src="https://goreportcard.com/badge/github.com/kawakatz/grepfiles"></a>
<a href="https://www.codefactor.io/repository/github/kawakatz/grepfiles/badge"><img src="https://www.codefactor.io/repository/github/kawakatz/grepfiles/badge"></a>
<a href="https://twitter.com/kawakatz"><img src="https://img.shields.io/twitter/follow/kawakatz.svg?logo=twitter"></a>
</p>

<p align="center">
  <a href="#installation">Installation</a> •
  <a href="#usage">Usage</a>  •
  <a href="#-todo">Todo</a>  •
  <a href="#acknowledgement">Acknowledgement</a>
</p>

Grep tool for some file formats like:
- Image (.png, .jpg, etc.)
- PDF
- Microsoft Word, Excel, PowerPoint
- SQLite

and other text formats.

## Installation
```sh
# install dependencies
➜  ~ brew install wv
➜  ~ brew install tesseract
➜  ~ wget https://github.com/tesseract-ocr/tessdata/raw/main/jpn.traineddata -O /usr/local/share/tessdata/jpn.traineddata # add japanese data
# install grepfiles
➜  ~ go install -v github.com/kawakatz/grepfiles/cmd/grepfiles@latest
```

## Usage
```sh
➜  ~ grepfiles <path> <keyword>
```

`<path>` can be a path of a file or directory.<br>
If a directory path is specified, this tool will **recursively** grep all files in the directory.<br>
<br>
This tool is case **insensitive**.<br>
If you need a case sensitive grep, you can grep as below:
```sh
➜  ~ grepfiles <dir> <keyword> | grep <keyword>
```

## 📋 Todo
- Support more file formats (feature requests and pull requests are welcome)
- Optimize performance

## Acknowledgement
This README.md format is inspired by  [@projectdiscovery](https://github.com/projectdiscovery/)🙇‍♂️