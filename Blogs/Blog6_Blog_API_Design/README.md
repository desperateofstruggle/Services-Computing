# 开发 web 服务程序cloudgo

<script async src="//busuanzi.ibruce.info/busuanzi/2.3/busuanzi.pure.mini.js"></script>
<span id="busuanzi_container_page_pv">本文总阅读量<span id="busuanzi_value_page_pv"></span>次</span>

## 前言

本次实验是服务计算课程的第五次实验，是开发简单 web 服务程序 cloudgo，并从中了解 web 服务器工作原理的一次实验。

---

文章目录

- [前言](#前言)
    - [1. 概述](#1.概述)   
    - [2. 博客API设计](#2.博客API设计)    
    - [3. 实验总结](#3.实验总结)

---

## 1.概述

  - REST
    
    - REST，表现层状态转换(`Representational State Transfer`），是[Roy Thomas Fielding博士](https://zh.wikipedia.org/w/index.php?title=Roy_Thomas_Fielding&action=edit&redlink=1)于2000年在他的博士[论文](https://www.ics.uci.edu/~fielding/pubs/dissertation/top.htm)中提出来的一种万维网软件架构风格，它使在不同软件不同程序在网络中互相传递信息变得更为方便。

    - 它是基于HTTP协议之上实现的一组约束和属性，是一种设计提供万维网络服务的软件构建风格。

    - 主要约束有，**参考[wikipedia](https://zh.wikipedia.org/wiki/%E8%A1%A8%E7%8E%B0%E5%B1%82%E7%8A%B6%E6%80%81%E8%BD%AC%E6%8D%A2)**：

      - Client-Server，二者的关注点分离。

      - Stateless， 每一次从客户端发送的请求中, 要包含所有的必须的状态信息, 会话信息由客户端保存, 服务器端根据这些状态信息来处理请求。

      - Uniform Interface：

        统一接口是 RESTful 系统设计的基本出发点. 它简化了系统的架构, 减少了耦合性, 可以让所有模块各自独立的进行改进. 对于统一接口的四个约束是:

        1. 请求中包含资源的 ID (Resource identification in requests )

           **在本实验中博客服务端的相关urls便是其中的资源ID，且服务器将以JSON 的方式发送给客户端**

        2. 资源通过标识来操作(Resource manipulation through representations)

        **本实验中将会根据HTTP的方法定义进行操作**

        3. 消息的自我描述性(Self-descriptive messages)

        4. 用超媒体驱动应用状态 ( Hypermedia as the engine of application state (HATEOAS))

## 2.博客API设计

  - 假定博客的web站点为`api.myBlog.com`, 且收发数据皆为JSON。

  - Blog API 相关功能接口

    - 登录: 

    ```
        'login' : curl -u username:password https://api.myBlog.com
    ```

    - 获取当前用户信息(`GET`)

    ```
        curl -i https://api.myBlog.com/user1/info
    ```

    - 获取当前用户的所用博客文章信息(`GET`),以way的排序方式，如热度(hot)、最后修改时间(time)等

    ```
        curl -i https://api.myBlog.com/user1/articles_info/{ways}
    ```

    - 获取文章的分类(GET)

    ```
        curl -i https://api.myBlog.com/user1/articles_category
    ```

    - 获取指定文章(`GET`)

    ```
        curl -i https://api.myBlog.com/user1/articles/{id}
        或
        curl -i https://api.myBlog.com/user1/articles/{title}
    ```

    - 查看指定文章下的评论(`GET`)

    ```
    curl -i https://api.myBlog.com/user1/articles/{id}/comments
    或
    curl -i https://api.myBlog.com/user1/articles/{title}/comments
    ```

    - 创建评论(`POST`)

    ```
    curl -i https://api.myBlog.com/user1/articles/{title}/comments -d {"content":"xxx", "length": "xxx"}
    ```

    - 发布文章(`POST`)

    ```
    curl -i https://api.myBlog.com/user1/articles/issue -d {"content":"xxx", "length": "xxx"
    ...}
    ```

    - 删除文章(`DELETE`)

    ```
    curl -i https://api.myBlog.com/user1/articles/delete -d {"title":"xxx","content":"xxx", 
    ...}
    ```

    - 收藏某博客(`POST`)

    ```
    curl -i https://api.myBlog.com/user1/articles/loves -d {"title":"xxx",
    ...}
    ```

    - 关注、取关、拉黑、取消拉黑某用户(`POST`)

    ```
    curl -i https://api.myBlog.com/user1/articles/follows -d {"username":"xxx",
    ...}
    和
    curl -i https://api.myBlog.com/user1/articles/unfollows -d {"username":"xxx",
    ...}
    和
    curl -i https://api.myBlog.com/user1/articles/block -d {"username":"xxx",
    ...}
    和
    curl -i https://api.myBlog.com/user1/articles/unblock -d {"username":"xxx",
    ...}
    ```

    - 等...

  - 错误信息相关(`状态码相关`)

    - 状态码由三个十进制数字组成，第一个十进制数字定义了状态码的类型，后两个数字没有分类的作用。HTML状态码共分为5种类型，此处采用其中的4种:

      - 1** : 信息，服务器收到请求，需要请求者继续执行操作

      - 2** : 成功，操作被成功接收并处理

      - 4** : 客户端错误，请求包含语法错误或无法完成请求

      - 5** : 服务器错误，服务器在处理请求的过程中发生了错误

    - 详细：

      - 200 : 请求成功。一般用于GET与POST请求(OK)

      - 201 : 已创建。成功请求并创建了新的资源(Created)

      - 202 : 已接受。已经接受请求，但未处理完成(Accepted)

      - 400 : 客户端请求的语法错误，服务器无法理解(Bad Request)

      - 401 : 请求要求用户的身份认证(Unauthorized)
      
      - 403 : 服务器理解请求客户端的请求，但是拒绝执行此请求(Forbidden)

      - 404 : 服务器无法根据客户端的请求找到资源（网页）。通过此代码，网站设计人员可设置"您所请求的资源无法找到"的个性页面(Not Found)

      - 500 : 服务器内部错误，无法完成请求(Internal Server Error)

      - 501 : 服务器不支持请求的功能，无法完成请求(Not Implemented)

      - 503 : 由于超载或系统维护，服务器暂时的无法处理客户端的请求。延时的长度可包含在服务器的Retry-After头信息中(Service Unavailable)

      - 505 : 服务器不支持请求的HTTP协议的版本，无法完成处理(HTTP Version not supported)


## 3.实验总结
  
  本次实验是基于REST API设计的一次实验，涉及内容丰富，所获颇丰。

