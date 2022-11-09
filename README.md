# ATTENTION： Waiting To Be Modified

## 原Whistle ReadMe

Whistle is a collection of Titanium projects.

How to build:
1. mkdir build && cd build
2. cmake ..
3. make

### 1. Clone the repository

```shell
git clone ssh://git@code.corp.bcollie.net/source/whistle.git
```

![image-20220322121931571](https://luochengyu.oss-cn-beijing.aliyuncs.com/img/image-20220322121931571.png)

### 2. Install go-swagger

```shell
go get github.com/go-swagger/go-swagger
```

If you can't download it because of the network error， try：

```
go env -w GOPROXY=https://goproxy.cn
```

Go to the directory of go-swagger and install go-swagger:

```shell
go install ./cmd/swagger
```

![image-20220322124337390](https://luochengyu.oss-cn-beijing.aliyuncs.com/img/image-20220322124337390.png)

### 3. Update submodule

```shell
cd whistle
git submodule init
git submodule update
```

![image-20220322124805944](https://luochengyu.oss-cn-beijing.aliyuncs.com/img/image-20220322124805944.png)

### 4. Build the project

```shell
mkdir build
cd build
cmake ..
make
```

![image-20220322125032956](https://luochengyu.oss-cn-beijing.aliyuncs.com/img/image-20220322125032956.png)

![image-20220322125257933](https://luochengyu.oss-cn-beijing.aliyuncs.com/img/image-20220322125257933.png)
