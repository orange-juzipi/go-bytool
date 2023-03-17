# go-bytool
## 🥤简介

go-bytool是一个简便的工具库，它能让你在Go里降低学习API的成本，使开发者无需把时间花在API如何实现，开发者仅需关注业务即可，剩下的go-bytool全部帮你实现。
##  🍉预实现的工具包

- 日期格式化
    - [x] 自定义格式化暂时只支持常见的：
        - yyyy-MM-dd HH:mm:ss
        - yyyy/MM/dd HH:mm:ss
        - yyyy.MM.dd HH:mm:ss
        - yyyy年MM月dd日 HH:mm:ss

        <br/>不支持不规则的格式，例如：yyyy-MM月dd HH:mm:ss，最终以第一个分隔符为准。

- map编排
    - [x] map排序
    - 更多的map花里胡哨操作

- 信息脱敏
  - [x] 手机号
  - [x] 密码
  - [x] 邮箱
  - [x] 身份证
  - [x] 姓名

- 类型转换工具包
    - [x] Slice转map类型 

- utils工具
  - 身份证工具
      - [x] 获取年龄
      - [x] 获取生日
      - [x] 获取省份

- string类型的操作
- 生成工具
- cron定时
- 第三方：邮件、二维码、图形验证码（captcha）、Emoji等
- 日志
- jwt
- poi：针对Excel、world操作
- io
- http工具包

## 🍑包分类
```
.
├── core                    // 核心包
│   ├── convert             // 类型转换工具包
│   ├── dates               // 日期时间工具包
│   ├── desensitization     // 信息脱敏工具包
│   ├── maps                // map工具包
│   └── utils               // 各种工具包
└── pkg                     // 公共包
...
``` 


## 🍊要求

Go版本需要：1.20+ <br/>

## 🍓issues
欢迎提 **issues** | **fork**，为开源贡献一份力量。