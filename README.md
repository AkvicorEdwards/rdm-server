# RDM(Reward Mechanism)

这个项目起因与我想记录自己每天主要干了什么事情，以及每件/类事情干了多久。
于是开发了一个只有记录功能的web程序，奈何自己总是忘记记录导致需要过后补充。
于是又引入了硬币奖励和消耗的机制，每件事情都会获得或消耗一定数量的硬币，
有的是按照次数来计算，有的则是按照时间来计算，亦或者是在长时间没有做某项任务
之后自动计算。

# Usage

```shell script
go mod download
# 编译
go build
# 初始化数据库
./rdm init
# 运行（端口8080）
./rdm
```

## 初始化数据库

数据库使用sqlite3，文件名为`rdm.db`

`./rdm init`

## 用户/权限

只能使用注册码注册，每个注册码都包含着使用此注册码注册用户所能获得的权限。  
默认管理员注册码 `admin`，权限值为 `11747296`  
推荐普通用户权限值为 `21807135`

相关信息和操作定义在`def/permission.go`

```text
使用二进制的 与 运算将两个权限值合并，高权限覆盖低权限

    1 // 使用Forest
    2 // 使用充值
    4 // 创建自己的任务
    8 // 查看自己的任务
    16 // 修改自己的任务
    32 // 创建所有人的任务
    64 // 查看所有人的任务
    128 // 修改所有人的任务
    256 // 创建注册码
    512 // 查看注册码
    1024 // 修改注册码
    2048 // 创建用户
    4096 // 查看用户
    8192 // 修改用户
    16384 // 查看自己的记录
    32768 // 修改自己的记录
    65536 // 查看所有人的记录
    131072 // 修改所有人的记录
    262144 // 查看自己的交易记录
    524288 // 修改自己的交易记录
    1048576 // 查看所有人的交易记录
    2097152 // 修改所有人的交易记录
    4194304 // 查看自己的分析
    8388608 // 查看所有人的分析
    16777216 // 使用TODO
```

## 添加任务

### 任务类型

 - **自动计算**： **每经过固定的时间**就计算一次收益，如果**关联任务id不为0**，
   则每次关联任务完成后就**重置**自动任务的剩余时间
 - **充值/提现**： 充值固定比例为**1:100**，提现比例为**1:10000**。
 - **计次任务**： 每次完成任务后获得对应硬币
 - **计时任务**： 按照**开始任务-结束任务**的时长计算应获得的硬币
 
### 一组包含的单位数量

 - **计次/计时/充值/提现任务** 填写 **1**
 - **自动任务** 填写间隔时间（单位为**秒**）

### 完成一组后获得的硬币数量

硬币数量为**正**代表获得硬币  
硬币数量为**负**代表消耗硬币

- **充值/提现** 填写 100 和 10000，代表每1元可以兑换的硬币数量
- **自动任务/计次任务** 填写每次完成后获得的硬币数量
- **计时任务** 填写每分钟获得的硬币数量

### 关联的自动执行任务的id

**自动任务**和其他任务的关联，每次关联的任务触发时会重置此自动任务的剩余时间（只能关联一个其他任务）

## Forest

在此界面点击**已选**来选择任务。  
可以添加标签来标注此任务的详细信息。  
点击**开始/结束**按钮后，计次任务会直接完成，计时任务则会开始计时（再次点击则完成该任务）。

## 删除/修改 操作

选中记录后，被选中的记录会变红，此时点击删除或修改操作的按钮按照提示进行操作。  
删除可以多选，修改只能单选