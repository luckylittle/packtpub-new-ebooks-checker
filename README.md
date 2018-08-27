# PacktPub new e-books checker

Golang application that lists latest e-books on the Packt Publishing website.

## run

`./packtpub-new-ebooks-checker`

## output

```bash
New Packt Publishing E-books on 2018-08-27 21:39:51.611530196 +0200 CEST m=+0.531558086
"Learning Salesforce Lightning Application Development"
"Qt5 Python GUI Programming Cookbook"
"R Deep Learning Essentials - Second Edition"
"Remote Usability Testing"
"Java EE 8 Design Patterns and Best Practices"
"Blockchain Quick Reference"
"Hands-On Intelligent Agents with OpenAI Gym"
"Bash Cookbook"
"Hands-On Deep Learning for Images with TensorFlow"
"Building Machine Learning Systems with Python - Third Edition"
"Python Artificial Intelligence Projects for Beginners"
"Hands-On Full Stack Web Development with Angular 6 and Laravel 5"
"Learning Salesforce Lightning Application Development"
```

## build

1. `git clone` this repo to your `$GOPATH/src/github.com/luckylittle/packtpub-new-ebooks-checker`.

2. `go build` inside the above folder.

3. `sudo cp packtpub-new-ebooks-checker /usr/local/bin`.

4. `packtpub-new-ebooks-checker` from anywhere.

---

_Tested on go version go1.10.2 linux/amd64_

_Mon Aug 27 21:47:38 CEST 2018_
