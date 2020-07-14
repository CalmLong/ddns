# DDNS

基于 DNSPod 动态域名解析实现

## 运行

[点击这里](https://github.com/CalmLong/ddns/releases) 下载最新发布版本

解压并添加执行权限

```
# apt install unzip
unzip linux-amd64-ddns.zip
chmod 777 ddns
```

简单运行(需要设置[环境变量](#环境变量))

```
nohup ./ddns &
```

或者使用 [systemd](#systemd) 服务运行

## 环境变量

配置以下环境变量

* `DNSPOD_DOMAIN`
* `DNSPOD_KEY`

### 示例

添加临时环境变量

**注意：程序运行前该子域名必须存在**

```
export DNSPOD_DOMAIN=baidu.com,www
export DNSPOD_KEY=123456,abcdefghijklmn
```

或者编辑 `/etc/profile` 添加永久环境变量

## systemd

创建配置并编辑

```
vi /etc/systemd/system/ddns.service
```

写入以下内容，注意更改 `[Service]` 中的配置为实际内容

```
[Unit]
Description=DDNS
After=network.target
Wants=network.target

[Service]
# 运行本程序的用户
User=root
Group=root
# 环境变量
# 注意: 程序运行前该子域名必须存在
Environment="DNSPOD_DOMAIN=baidu.com,www"
Environment="DNSPOD_KEY=123456,abcdefghijklmn"
# 程序位置
ExecStart=/usr/bin/ddns
Restart=on-failure
RestartSec=60s

[Install]
WantedBy=multi-user.target
```

启用/禁用开机自启

```
systemctl enable|disable ddns
```

启动/停止/重启程序

```
systemctl start|stop|restart ddns
```

## 自定义返回公网 IP

请移步 [Wiki](https://github.com/CalmLong/ddns/wiki)

## 引用以下项目

* `http://icanhazip.com`

