# goget-rquote
a go program that uses a public random quote api (api.forismatic.com) to get a quote and its author and print it to stdout 

## Installation instructions
* Clone the repo
```
git clone https://github.com/EssaAlshammri/goget-rquote.git
```

* Move the file to /opt/rquote
```
sudo mv goget-rquote/rquote /opt/rquote
```

* Put this line into .bashrc so every time you open the terminal a new quote will be printed
```
echo /opt/rquote >> ~/.bashrc
source ~/.bashrc
```

and now you are good to go :+1:

this is my very first program using golang


