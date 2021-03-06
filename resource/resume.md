### 个人信息：


**姓名：巫凌健**

**微信：wlj1995222**

**求职意向：全栈开发工程师**

**Email：wulingjian123@gmail.com**

**Github: https://github.com/wljgithub** 

**教育经历：本科--吉林大学珠海学院（19届毕业生-工作大半年）** 

**技术栈：Golang，Linux Shell编程， Vue.js ，微信小程序**



### 个人经历


  我并非计算机专业的，大二的时候突然对编程很感兴趣，然后通过Python入了门。但是身边没什么人交流，只能到社区里“找食物吃，吃的东西很杂”。期间摸过很多东西，比如Python爬虫，Sql注入，Arp欺骗，但都是浅尝辄止，基本属于入门乱来阶段。

后来在大学期间分别有过两段Golang后台的实习，两家公司分别是做IM（即时通讯软件）和区块链的，我在那里主要写RESTFUL API。

毕业后，在一家做WebIDE的公司，做微信小程序开发。

算起来，从写"Hello,World"到现在，也有三年时间了。(大学两年+工作大半年)



## 个人项目


[**演示地址**](https://xiulu.xyz/)  [**代码**](https://github.com/wljgithub/my-blog)

描述：这是我从零开始写的一个博客，包括前后端的设计（技术选型，表的设计，API制定，UI布局），开发 （独立解决各种问题 ) ，调试（Postman，Chrome开发者工具，断点调试），部署（AWS EC2，Nginx 代理，Let'encrypt签证书）



功能：支持多用户，Markdown预览，分页等。

特性：实现前后端分离，响应式布局(Bootstrap实现)

技术选型：前端用Bootstrap + Vue，后端用Gin（Golang的一个Web框架）



写完后用Godaddy签了个域名，然后部署到 AWS的EC2上，主要干了几件事：

- 用Let'Encrypt 签了个证书，配置Nginx 支持HTTPS 传输

- 配置DNS CNAME 使得  [www.xiulu.xyz](https://www.xiulu.xyz)     [xiulu.xyz](https://xiulu.xyz) 都跳到同一个服务器

- 配置Webpack生成Gzip静态文件，配置Nginx支持Gzip传输，节省网络带宽（4M压缩到700K）



然后嫌每次部署都太麻烦，还写了一键Shell脚本，[manage.sh](https://github.com/wljgithub/my-blog/blob/master/manage.sh)

主要干的就是自动拉取Repo编译，安装配置Nginx，Mysql，Iptables之类的事。

只要新创一个Centos6的服务器，运行这个脚本之后，就可以从浏览器正常访问博客了，你也可以新建一个服务器来玩一下。



Bootstrap 和 Vue 是为了做这个博客才学的，算是现学现做把。

**温馨提示：演示地址和代码见左上角 ↑**



### 专业技能


**后端相关：**

- 熟悉Golang，及Golang的一个Web框架Gin，并能用Gin来快速搭建Restful API 服务器

- 有过Mysql，Sqllite 使用经历，并能独立设计数据表，对数据表设计做优化

- 遇到过浏览器跨域问题（Preflight），并尝试用过不同方案解决过该问题

- 有了解过Https相关原理（TLS 握手过程，证书校验，Wireshark的数据报文分析），并用let's encrypt签过证书

- 对TCP/IP 协议栈中ICMP有过深入的了解，曾用Wireshark抓包分析过，并根据其报文格式实现过一个Golang版的Ping

  


**前端相关：**


- 有过微信小程序开发的经历

- 熟悉Vue，可以基于Vue CLI快速搭建并开发项目

- 熟悉bootstrap，并能用来实现响应式布局

- 能对Webpack，Babel和Eslint 进行基本的配置

  

**运维相关**


- 有过 Godaddy域名注册，Cloudflare CDN 使用的经历
- 有过多个云平台的使用经历，Vultr，AWS，星光互联
- 有过不同Linux发行版使用经历，Centos，Ubuntu，Amazon Linux等
- 扎实的网络基础，一定的网络排查能力，掌握常用网络排查用具（Tcpdump，Wireshark，Traceroute，Netstat等）
- 能独立编写Shell脚本，并能兼容不同的发行版（Ubuntu,Centos等）



### 英语能力


大学的时候考了一次四级没过，后来接触了编程之后深刻地认识到了英语的重要性，踩了很多坑，花了很大的心思去学习。

现在阅读和写作方面，能无障碍逛一些技术社区(Stackoverflow，Medium)，也可以用英语做一些简单的邮件交流。

听力和口语的方面上还不行，只能听懂 [Vue JS Crash Course](https://www.youtube.com/watch?v=Wy9q22isx3U) 这种水平的视频。



### 感兴趣的东西


对网络协议相关的东西很感兴趣，有折腾过路由表，和各种代理软件。

之前也根据ICMP报文格式实现过Golang版的Ping，后面又基于Ping，实现过Golang版的Tracetoure，并写了一篇文章记录了实现的过程。[点这里](https://segmentfault.com/a/1190000020048492)



### 期望


**找团队上：**

从我描述的经历您或许可以看得出来，我对技术是不挑剔的，我只是对人挑剔而已。

我前后端和运维都可以做一点，而且我对我的学习能力怀着充足的信心，你把我放到程序员中任何一个领域我都能存活下来。

最后，希望能找到一群热爱技术、对技术有追求的人，一个关心员工成长的团队。

当然，如果能接触得到 Elastic Search，Kubernets之类的技术就更好了。

**期望薪资:** 

15k~20k

### 简化版简历

为了尽可能地介绍自己，这份简历可能写得有些啰嗦。
如果需要打印简历，或其他用途。我这里也写了一份[简化版](https://github.com/wljgithub/for_interview/blob/master/resource/briefResume.md) 的。
