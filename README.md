<p align="center">
<img src="https://i.loli.net/2021/03/26/EvQy3whrHxd8Cq6.gif" alt="LikeTriple" width="300">
</p>
<h1 align="center">LikeTriple 一键三连</h1>

>LikeTriple 是一个B站视频下载工具，他能很方便的下载你B站投币视频。

由于未知原因，你的收藏夹的视频看着看着就没了。

![](doc/404.png)
失效的视频太难受了

## 快速上手

```
docker run -d \
--name liketriple \
--env mid=xxxxxxx \
-v `pwd`:/download \
--restart always \
mirrors2/liketriple
```
* --env mid= 请替换为你自己的，https://space.bilibili.com/ 后面的那串数字

* -v 把你的下载目录挂载到 /download 

* 请确保你的空间设置为 公开
  ![](doc/coinvideo.png)

## 成果展示
![](doc/start_succ.png)



妈妈再也不用担心我收藏夹的视频不见了。闲置的NAS，软路由可以用起来了
