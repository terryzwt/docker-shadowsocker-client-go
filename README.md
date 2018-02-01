# 介绍
  pgfast的代理，基于[shadowsocks-go](https://github.com/shadowsocks/shadowsocks-go)。
  
  可以在多个shadowsocks-server之间自动选择可用server.
  
  每小时从pgfast网站订阅一次地址并更新。
  
  推荐使用chrome浏览器，并使用SwitchyOmega插件进行浏览器代理上网。

# 用法
```bash
git clone git@github.com:terryzwt/docker-shadowsocker-client-go.git
cp env-example .env #里面可以配置自定义暴漏的端口。默认是2080.
cd pgfast
cp pgconfig.toml.example pgconfig.toml
vi pgconfig.toml #里面有两个参数Subscribe_url和Password,可以在https://www.pgfastss.net上找到。需要的是付费用户才行。
docker-compose up -d
```
