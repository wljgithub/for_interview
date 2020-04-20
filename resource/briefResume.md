### 个人信息：

**姓名：巫凌健**

**微信：wlj1995222**

**求职意向：全栈开发工程师**

**Email：wulingjian123@gmail.com**

**教育经历：本科--吉林大学珠海学院（19届毕业生-工作大半年）** 

**技术栈：Golang，Linux Shell编程， Vue.js ，微信小程序**



### 个人经历

- 自学编程，大二的时候通过Python爬虫入门

- 后来在大学期间有过两段Golang后台的实习，他们分别是做区块链，IM（即时通讯软件）

- 毕业后在一个做WebIDE的公司，做微信小程序开发。

  

## 个人项目

[**演示地址 **](https://xiulu.xyz/)  [**代码**](https://github.com/wljgithub/my-blog)

**描述：**这是我从零开始写的一个博客，包括

- 前后端的设计（技术选型，表的设计，API制定，UI布局）
- 开发 （独立解决各种问题 ) 
- 调试（Postman，Chrome开发者工具，断点调试）
- 部署（AWS EC2，Nginx 代理，Let'encrypt签证书）



**支持： **

- 功能：支持多用户，Markdown预览，分页等
- 特性：实现前后端分离，响应式布局(Bootstrap实现)
- 技术选型：前端用Bootstrap + Vue，后端用Gin（Golang的一个Web框架）



**部署: **

写完后用Godaddy签了个域名，然后部署到 AWS的EC2上，主要干了几件事：

- 用Let'Encrypt 签了个证书，配置Nginx 支持HTTPS 传输
- 配置DNS服务器的 CNAME 使得  [www.xiulu.xyz](https://www.xiulu.xyz)     [xiulu.xyz](https://xiulu.xyz) 都跳到同一个Ip
- 配置Webpack生成Gzip静态文件，配置Nginx支持Gzip传输，减少传输时间（4M压缩到700K）



**脚本自动化：**

然后嫌每次部署都太麻烦，还写了一键Shell脚本，[manage.sh](https://github.com/wljgithub/my-blog/blob/master/manage.sh)

自动拉取Repo编译，安装配置Nginx，Mysql，Iptables等。

只要新创一个Centos6的服务器，运行这个脚本之后，就可以从浏览器正常访问博客了



### 专业技能

**后端相关：**

- 熟悉Golang，常用标准库,并能运用一个Web框架 Gin，快速开发

- 了解HTTP协议,能独立设计Restful API ，并有过相关实践

- 有过Mysql，Sqllite 使用经历，并能独立设计数据表，对数据表设计做优化

- 遇到过浏览器跨域问题（Preflight），并尝试用过不同方案解决过该问题

- 对TCP/IP 协议栈中ICMP有过深入的了解，曾用Wireshark抓包分析过，并根据其报文格式实现过一个Golang版的Ping

  

**前端相关：**

- 有过参与过微信小程序开发，熟悉开发，上线流程

- 能熟练使用Vue开发常用组件，基于Vue CLI快速开发项目

- 能基于bootstrap，实现响应式布局，兼容不同设备

- 能配置Webpack，Babel和Eslint 常用选项

- 接触过Typescript，接触过Sass，

- 一定的网站性能，分析，和优化能力，并有相关实践( 懒加载，Code-Splitting)

  

**运维相关： **

- 了解HTTPS相关的原理，也用Let's Encrypt签过证书
- 掌握一定的安全常识，SSH安全配置，防火墙规则配置
- 有过不同Linux发行版使用经历，Centos，Ubuntu，Amazon Linux等，
- 熟悉Bash语法，能独立编写Shell脚本，可以写出兼容不同发行版的Shell脚本
- 一定的网络排查能力，DNS 缓存，Traceroute追踪数据包，Wireshark分析数据报文
- 有过 Godaddy域名注册，Cloudflare CDN 使用，和Vultr，AWS，星光互联多个云平台的使用经历。

### 英语能力

大学考了次四级没过，后来接触了编程后花费了很大一番心思去学习

现在阅读和写作方面，能无障碍逛一些技术社区(Stackoverflow，Medium)，也可以用英语做一些简单的邮件交流，大概是这个[README](https://github.com/wljgithub/my-blog)的水平

听力和口语方面，能听懂 [Vue JS Crash Course](https://www.youtube.com/watch?v=Wy9q22isx3U) 这种水平的视频。



### 感兴趣的东西

对网络协议相关的东西很感兴趣，有折腾过路由表，和各种代理软件。

之前也根据ICMP报文格式实现过Golang版的Ping，后面又基于Ping，实现过Golang版的Tracetoure，并写了一篇文章记录了实现的过程。[点这里](https://segmentfault.com/a/1190000020048492)



