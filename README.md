# ROS-Web for MyTeleop

![预览](https://github.com/Nomango/ROSWeb-MyTeleop/raw/master/preview.png)

一个可以在网页中预览 Gazebo 仿真环境并通过按键发送 geometry_msgs::Twist 消息的可视化程序。

### 环境

**安装环境**
- Ubuntu 16.04
- Golang 1.11
- ROS kinetic
- Gazebo 7.15

**库环境**
- 后端 GoGin
- 前端 ROS3d.js
- 仿真 TurtleBot3
- ROS包 rosbridge tf2_web_republisher

### 安装

1. 安装 rosbridge

```bash
sudo apt-get install ros-kinetic-rosbridge-suite
sudo apt-get install ros-kinetic-tf2-ros
```

2. 克隆仓库

```bash
go get github.com/Nomango/ROSWeb-MyTeleop
```

3. 复制模型

将 Gazebo 模型文件夹复制到 /static/models 目录下，为服务器提供模型资源

4. 编译并启动

```bash
go build main.go
./main
```
