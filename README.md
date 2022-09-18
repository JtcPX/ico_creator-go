一個go版本的ico圖標生成器，用於修改ico的hash值
### 用法
- go run ico_creator.go -h

- usage:ico_creator.go  [-f F] [-n N]
- -f File Input file directory
- -n Numbers Please enter a num


### go run ico_creator.go -f /home/test.ico -n 10


<!--運行邏輯
生成一份--》打開文件，插入1，生成文件

生成兩份--》打開文件，插入1生成文件，插入2生成文件
使用for 循環
```
i=0，i<2,i++{
	打開文件，插入i，
	保存文件
}
```
-->