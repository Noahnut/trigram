# Trigram
Implement simple Trigram for search word in the file

# Install

```
 go get github.com/Noahnut/trigram
```

# Usage 
```golang
tri := NewTrigram()
err = tri.Add("testTwo") //read the file context to the Trigram index
result = tri.Find("Cod") //Search the word exist in which file
```

# Benchmark

# Reference 
* http://www.evanlin.com/trigram-study-note/
* https://swtch.com/~rsc/regexp/regexp4.html
* https://github.com/kkdai/trigram
* https://github.com/dgryski/go-trigram