[TOC]

本周review分享Linux日志相关的内容


### 开胃菜
在介绍Linux日志相关的内容之前，先来点开胃菜，比如有时候我们希望查看同时在线的用户，可以用 who命令,可以看到如下输出

```bash
root     pts/0        2019-06-14 09:42 (183.14.18.123)
root     pts/1        2019-06-14 09:42 (183.14.18.321)
root     pts/2        2019-06-15 09:39 (183.14.18.652)
root     pts/3        2019-06-15 09:41 (183.14.18.132)
```
> 在centos中或debian中，who命令是从 /var/run/utmp日志文件中获取登陆用户信息。

而如果你想查看root用户登陆的历史，可以用last命令，你会看到类似的输出：
```bash
root     pts/3        183.14.18.123     Sat Jun 15 09:41   still logged in   
root     pts/2        183.14.18.149     Sat Jun 15 09:39   still logged in   
root     pts/4        183.14.18.151     Fri Jun 14 16:20 - 19:00  (02:40)    
root     pts/4        183.14.18.152     Fri Jun 14 14:49 - 16:20  (01:30)    
```

> 这里顺便介绍一个小技巧，比如你想查看系统重启的记录，可以使用这条命令 last reboot

所有用户最近一次的登陆时间，可以用lastlog命令

```
Username         Port     From             Latest
root             pts/0    183.66.77.88    Tue Jun 18 09:40:20 +0800 2019
bin                                        **Never logged in**
daemon                                     **Never logged in**
adm                                        **Never logged in**
lp                                         **Never logged in**
sync                                       **Never logged in**
shutdown                                   **Never logged in**
halt                                       **Never logged in**
mail                                       **Never logged in**
operator                                   **Never logged in**
games                                      **Never logged in**
ftp                                        **Never logged in**
nobody                                     **Never logged in**
systemd-network                            **Never logged in**
dbus                                       **Never logged in**
polkitd                                    **Never logged in**
ntp                                        **Never logged in**
sshd                                       **Never logged in**
postfix                                    **Never logged in**
chrony                                     **Never logged in**
tss                                        **Never logged in**
jack       pts/0                     Tue Jun 18 09:49:04 +0800 2019

```


开胃菜讲完了，下面来进入主题，介绍一下如系统日志相关的内容，主要包含如下方面：

- 系统日志是怎么产生的
- 产生的系统日志都去了哪里
- 日志的组成部分有哪些，如何查看
- 介绍日志的规则，即日志的来源和去处(来源意味着需要记录哪些进程产生的记录，去处是指这些记录需要重定向到终端还是写入日志文件中
)
- 最后通过一个完整的例子，教你如何自定义规则，并产生一条日志，然后可在文件中查看(包含完整的命令)

### 系统日志是怎么产生的

系统日志的是由一个叫rsyslog的守护进程，它会监听系统的各个模块所生成的消息，比如用户通过ssh登陆的时候，操作系统的sshd模块会生产一条类似这样的记录

    pam_unix(sshd:session): session opened for user root by (uid=0)

然后rsyslog会将这条记录存入到系统特定的目录中

### 产生的日志都去了哪里

你们会不会好奇，系统生产的日志都去了什么地方，它其实存放在/var/log中(centOS),如果你用 ls /var/log，就会看到如下输出：

```
anaconda           btmp           dmesg       grubby              maillog            myTestInfo.log  secure            tallylog
boot.log           chrony         dmesg.old   grubby_prune_debug  maillog-20190616   ntpstats        secure-20190616   tuned
boot.log-20190614  cron           ethcon.log  imageboot.log       messages           qemu-ga         spooler           wtmp
boot.log-20190617  cron-20190616  firewalld   lastlog             messages-20190616  rhsm            spooler-20190616  yum.log
```
可以看到里面包含了许多日志文件

### 日志的结构

我们先来看一下具体系统日志的长什么样：

```
Jun 17 10:12:13 vultr sshd[21546]: pam_succeed_if(sshd:auth): requirement "uid >= 1000" not met by user "root"
Jun 17 10:12:14 vultr sshd[21546]: Failed password for root from 218.92.0.206 port 33909 ssh2
```
先说一下，这是一条很有意思的记录，这条记录的意思是218.92.0.206这个ip的人，尝试用root用户登陆我的主机，我不知道他是登陆的时候输错IP，还是想暴力破解我的服务器密码，反正我看到的时候吓一跳。

好了，接着说日志的格式，日志由四个部分组成，分别是：日期+主机名+模块+日志信息



|   格式   |                            实例                            |
|----------|------------------------------------------------------------|
| 日期     | Jun 17 10:12:14                                            |
| 主机名   | vultr                                                      |
| 模块     | sshd[21546]                                                |
| 日志信息 | Failed password for root from 218.92.0.206 port 33909 ssh2 |


每条日志都会标明时间，主机，和产生日志的模块，所以你可以很清晰地区分这条日志是什么时候产生的，由谁产生的。

### 日志规则

前面介绍了日志是怎么来的，也介绍了日志都去了哪里，即日志的来源和去处都介绍了。你会不会有那么一点冲动，自己定义日志的来源和去处，你可以指使日志从哪里来，到哪里去。

下面我给你介绍一下，如何定义日志的来源和去处。

其实系统定义日志的规则，是放在/etc/rsyslog.conf文件种，我摘抄其中一段给你介绍一下日志的规则

```
cron.info                                               /var/log/cron
auth.*                                                  @prep.ai.mit.edu
auth.*                                                  root,amrood
```

可以看到日志分两个部分，左边的部分就是日志的来源，来源分中也分两个部分，格式为：设备+日志级别

比如cron.info ，cron是指产生日志的设备(facility),info是指日志的级别

去处就简单了，就一个部分，你可以指定本地的一个文件，也可以指定远程的一台服务器，比如 @prep.ai.mit.edu 或/var/log/cron

好了，规则你知道了。

### 例子

下面通过一个具体的例子，向你介绍如何自定义日志规则，并产生一条日志消息。

**第一步，定义日志规则**

在 /etc/rsyslog.d/下创建一个文件，名字你随便取，在这里我用myLogRules.conf，你可以直接复制以下命令(centOS)

```bash
vim /etc/rsyslog.d/myLogRules.conf
```

然后输入以下内容

    *.info  /var/log/myTestInfo.log

退出保存(:wq)

**第二步，重启rsyslog守护进程**

```bash
service rsyslog restart
```

**第三步，生成一条日志**

    logger -p local4.info " This is a info message from local 4"

好了，你已经生成了一条日志消息了，输入以下命令：

     tail /var/log/myTestInfo.log

你可以在/var/log/myTestInfo.log文件了看到如下日志
```
Jun 19 10:31:40 vultr root: This is a info message from local 4

```

到目前为止，我们已经介绍了日志是怎么产生的，产生的日志都去了哪里，日志由哪几部分组成，还介绍了如何自定义日志的来源和去处，并通过了一个具体的例子向你介绍如何实现这个过程，虽然我认为还有一些内容没有在这个review里介绍到(因为我自己都没搞懂，包括我查了一些文档也没有清晰地介绍这个，就是自定义日志的action中，如果不填用户是不是默认是所有用户都能查看，如果指定用户又会怎样)，但我感觉已经大致框架还是很清晰的，下次有空再详细介绍跟日志相关的内容，上周ARTS总算补全了。

本周review来源于这篇[文章](https://www.digitalocean.com/community/tutorials/how-to-view-and-configure-linux-logs-on-ubuntu-and-centos)并加以整理，你可以跳转到原文查看详细的介绍
