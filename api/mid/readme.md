# Multi Interface Description
MID is a extendable **IDL** and description language to declare your interfaces. It is just a simple JSON or YAML file ;-) .
Write your MIDs centrally and create protos, schemas and more from them.

## under the hood
1. uses golang template engine
2. the *funcmap* enriching the template engine is based on [Masterminds/sprig](https://github.com/Masterminds/sprig), and contains type-manipulation, iteration and language-specific helpers
3. the `ast` is JSON or YAML 
4. uses [simple-generator](https://github.com/veith/simple-generator/blob/master/main.go) which is inspired by [moul's protoc-gen-gotemplate](https://github.com/moul/protoc-gen-gotemplate)

## usage
just execute
```
scripts/midDemo.sh


simple-generator --h
Usage of simple-generator:
  -d string
        Path to data file which contains YAML or JSON
  -t string
        Path to tpl file

```

## writing your own templates (documentation) 
 - [Build web application with Golang, 7.4 Templates](https://astaxie.gitbooks.io/build-web-application-with-golang/en/07.4.html)
 - [**sprig**, Useful template functions for Go templates](http://masterminds.github.io/sprig/)



# Install for simple-generator

If you have Go installed:

```
go get github.com/veith/simple-generator
```

