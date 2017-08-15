# Gofind

> A tool help you find anything you try to find. Developed by Jin Tian.

## Synoposis

Gofind is a command line utils to find any file under current dir, it searchs not only dirs but also file inside under this dir.

```
# searching dir names and content under current dir
gofind 金庸
# searching only dir and file names under current dir
gofind -d -f 金庸
# you can also using like this
gofind -df 金庸
# searching only content under current dir
gofind -c 金庸
# searching specific directory
gofind /any/path/you/want/search  金庸
```
That is all usage, you will locate all the things under your directory.



## Showtime

gofind is a marvelous helper you will use all the time. find anything in a snap time.

<img src="http://opbocoyb4.bkt.clouddn.com/WechatIMG6680.jpeg" >

<img src="http://opbocoyb4.bkt.clouddn.com/WechatIMG6678.jpeg">

<img src="http://opbocoyb4.bkt.clouddn.com/WechatIMG6681.jpeg">



## Tips

gofind has three mode to search things, search folder name, search file name, and search all content. gofind will automatically enable all mode by default, so if you want search only file and directory you can do this:

```
gofind -df 'you want search'
```



## Copyright

gofind implement by Jin Tian using golang. you should using it under **Apache License**.