# PDF TO PNG NODE
本程序用于在含有fitz的环境里使用，将pdf生成9宫格缩略图
需要搭配easyoa主服务，rocketmq消费

## build镜像教程
拉取最新代码
然后docker build . 构建出镜像
然后docker tag出pdf2png_node:1.0.x版本号
然后通过pnd2png.sh快速启动即可


## 直接运行
```shell
docker pull abbhb/pdf2png_node:latest
docker run -itd abbhb/pdf2png_node:latest即可！
```

## 生产业务
快速启动，通过脚本配合镜像
```shell
./pdf2png.sh start
```