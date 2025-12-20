# wxapkg

🚀 微信小程序一键解密和解包工具

## 📥 下载

在 [Release](https://github.com/zhuweiyou/wxapkg/releases/) 页面下载对应平台的二进制文件：

- **Windows**: `wxapkg_windows_amd64.exe`
- **MacOS**: `wxapkg_macos_amd64`

## 🚀 使用方法

### Windows 平台

1. 下载 `wxapkg_windows_amd64.exe` 文件
2. 找到微信小程序文件：
   - 微信 4.0 之前：`C:\Users\用户名\Documents\WeChat Files\Applet\{wxid}\{n}\__APP__.wxapkg`
   - 微信 4.0 之后：`C:\Users\用户名\AppData\Roaming\Tencent\xwechat\radium\Applet\packages\{wxid}\{n}\__APP__.wxapkg`
3. 将 `__APP__.wxapkg` 文件拖拽到 `wxapkg_windows_amd64.exe` 上即可完成解包

![演示GIF](https://github.com/zhuweiyou/wxapkg/assets/8413791/07a5cfa5-00c9-47b5-aaa3-ee42b878495f)

### MacOS 平台

1. 下载 `wxapkg_macos_amd64` 文件
2. 打开终端，赋予执行权限：
   ```bash
   chmod +x ./wxapkg_macos_amd64
   ```
3. 找到微信小程序文件：
   - 微信 4.0 之后：`/Users/用户名/Library/Containers/com.tencent.xinWeChat/Data/Documents/app_data/radium/Applet/packages/{wxid}/{n}/__APP__.wxapkg`
4. 运行解包：
   ```bash
   ./wxapkg_macos_amd64 /path/to/__APP__.wxapkg
   ```
