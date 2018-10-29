# What is it? 

* minimalistic example of using SBE in Go
* although these are already provided in SBE docs, this was a good starting example to learn about Go project setup, defining/exporting interfaces and playing  

# Generating SBE model files 

SBE schema is defined in [message.xml](api/message.xml) and [model](src/model/Message.go) was generated using [SBE tool](https://jar-download.com/artifacts/uk.co.real-logic/sbe-all/1.8.6/source-code):

```
java -Dsbe.target.language=Golang -jar /path_to/sbe-all-1.8.7.jar message.xml
```