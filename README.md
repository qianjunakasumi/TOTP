# TOTP 基于时间的一次性密码算法

![LICENSE](https://img.shields.io/github/license/qianjunakasumi/TOTP?style=for-the-badge)
[![Go Report Card](https://goreportcard.com/badge/github.com/qianjunakasumi/TOTP?style=for-the-badge)](https://goreportcard.com/report/github.com/qianjunakasumi/TOTP)

![Doc](https://pkg.go.dev/badge/github.com/qianjunakasumi/totp?style=for-the-badge)

适用于TOTP(RFC 6238)中默认和推荐的计数、哈希方法和密钥、密码长度实现的身份验证器中验证用户的算法

## 支持的验证器

- Google身份验证器
- Microsoft Authenticator
- 其他以时间为基准的默认协商参数的验证器

## 参数信息

| 参数 | 默认值 |
| --- |  ---  |
| X 时间步长（间隔） | 30（秒） |
| T0 纪元 | Unix（0） |
| 哈希算法 | HMAC-SHA-1 |
| 密码长度 | 6（位） |

## 如何使用

```
go get -u github.com/qianjunakasumi/totp
```

### 代码示例

```
ok, err := totp.Authenticate("KL67YBZBQJSU6FR7L4HZUYQ5OKKWNR2N", time.Unix(1614556800, 0), "721396")
if err != nil {
    // 错误处理
}
if !ok {
    // oops! 验证失败
}

// do 逻辑代码
```

## 基准测试

```
goos: windows
goarch: amd64
pkg: github.com/qianjunakasumi/totp
cpu: Intel(R) Core(TM) i5-9300H CPU @ 2.40GHz
BenchmarkGenerateTOTP
BenchmarkGenerateTOTP-8   	 1000000	      1052 ns/op	     520 B/op	       9 allocs/op
BenchmarkAuthenticate
BenchmarkAuthenticate-8   	 1000000	      1054 ns/op	     520 B/op	       9 allocs/op
PASS
```

## 参考文献及鸣谢

[基于时间的一次性密码算法 - 维基百科](https://zh.wikipedia.org/wiki/%E5%9F%BA%E4%BA%8E%E6%97%B6%E9%97%B4%E7%9A%84%E4%B8%80%E6%AC%A1%E6%80%A7%E5%AF%86%E7%A0%81%E7%AE%97%E6%B3%95)

[RFC 6238 - TOTP: Time-Based One-Time Password Algorithm](https://tools.ietf.org/html/rfc6238)

## 许可证

```
MIT License

Copyright (c) 2021 千橘 雫霞

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
```
