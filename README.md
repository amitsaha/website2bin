# Website2Bin - Static website to a Go binary

This command line application, given a directory containing HTML files and other directories (a static
website for example), will create a `go.mod` and `server.go` file (or overwrite. 

Once you build the server, you can deploy it without having to copy the files.

See the files, `server.go` and `go.mod` in this [repository](https://github.com/practicalgo/website/tree/main/public)
for example.

## Usage

```
$ go install github.com/amitsaha/website2bin

$ ~/go/bin/website2bin --help
Usage of website2bin:
  -listen-address string
    	Address for the server to listen on (default ":8080")
  -website-path string
    	Directory containing the website files
2021/12/10 07:58:49 flag: help requested
```

## Learn more

Checkout the [presentation](./presentation) at GopherCon 2021 lightning talks Day 3. I tried
to make the slides accessible using Microsoft Powerpoint's suggestions (.PPTX file).

A [video](https://youtu.be/XnPHI6cCL7E?t=10222) of the talk is also available now.

I have described how I took the idea that this tool implements to write a server which serves the 
content on https://practicalgobook.net in a [blog post](https://practicalgobook.net/posts/practicalgobook_net/).

## Limitations

- No sanity checking - it will embed all the files and directories at the specified path

## Contributions

Please file a issue and/or create a pull request after you create a fork.

