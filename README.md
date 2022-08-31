# grepfiles🔍
<p align="center">
<a href="https://opensource.org/licenses/MIT"><img src="https://img.shields.io/badge/license-MIT-_red.svg"></a>
<a href="https://github.com/kawakatz/grepfiles/issues"><img src="https://img.shields.io/badge/contributions-welcome-brightgreen.svg?style=flat"></a>
<a href="https://goreportcard.com/badge/github.com/kawakatz/grepfiles"><img src="https://goreportcard.com/badge/github.com/kawakatz/grepfiles"></a>
<a href="https://github.com/kawakatz/grepfiles/blob/master/go.mod"><img src="https://img.shields.io/github/go-mod/go-version/kawakatz/grepfiles"></a>
<a href="https://twitter.com/kawakatz"><img src="https://img.shields.io/twitter/follow/kawakatz.svg?logo=twitter"></a>
</p>

<p align="center">
  <a href="#installation">Installation</a> •
  <a href="#usage">Usage</a>  •
  <a href="#-todo">Todo</a>  •
</p>

Grep tool for some file formats like:
- Image (.png, .jpg, etc.)
- PDF
- Microsoft Word, Excel, PowerPoint
- SQLite

and other text formats.

## Installation
```sh
# install dependencies on macOS
➜  ~ brew install wv
➜  ~ brew install tesseract
➜  ~ wget https://github.com/tesseract-ocr/tessdata/raw/main/jpn.traineddata -O /usr/local/share/tessdata/jpn.traineddata # add japanese data

# install dependencies on Kali Linux
$ sudo apt install libpng-dev libjpeg-dev libtiff-dev zlib1g-dev gcc g++ autoconf automake libtool checkinstall
$ cd /opt
$ wget https://github.com/DanBloomberg/leptonica/releases/download/1.82.0/leptonica-1.82.0.tar.gz # download the latest release from GitHub (https://github.com/DanBloomberg/leptonica/)
$ tar -zxvf leptonica-1.82.0.tar.gz
$ cd leptonica-1.82.0/
$ ./configure
$ make
$ sudo checkinstall
$ sudo ldconfig
$ sudo apt install libtesseract-dev tesseract-ocr
$ wget https://github.com/tesseract-ocr/tessdata/raw/main/jpn.traineddata -O /usr/local/share/tessdata/jpn.traineddata # add japanese data

# install grepfiles
➜  ~ go install -v github.com/kawakatz/grepfiles/cmd/grepfiles@latest
or
➜  ~ go install -v github.com/kawakatz/grepfiles/cmd/grepfiles@v1.0.0
```

## Usage
```sh
➜  ~ grepfiles <path> <keyword> 2>/dev/null
➜  ~ grepfiles . 'secret' 2>/dev/null
```

`<path>` can be a path of a file or directory.<br>
If a directory path is specified, *grepfiles* will **recursively** grep all files in the directory.<br>
<br>
*grepfiles* is case **insensitive**.<br>
If you need a case sensitive grep, you can grep as below:
```sh
➜  ~ grepfiles <dir> <keyword> | grep <keyword>
```

## 📋 Todo
- Optimize performance
