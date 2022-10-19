1. 运行模式配置
    ```
    优先使用环境变量
   设置环境变量的方法：export ginMode="local"
   当环境变量不存在的时候，使用配置文件 config/config.ini 中的 mode 配置
   
   首次运行，需将 local.ini.example 重命名成 local.ini
   
   支持四种运行环境配置
   local：本地
   dev：开发
   test：测试
   prod：线上
    ```
2. 配置参数的读取
    ```
   读取通用配置参数：Config.Section("").Key("这里填写配置的key").String()
   读取不同环境的配置参数：Config.Section(Env).Key("这里填写配置的key").String()
   ```
3. 运行方式
   ```
   air server
   可自动监听文件变化，重启服务
   ```
4. swagger 访问路径
   ```
   http://127.0.0.1:8080/swagger/index.html
   ```
