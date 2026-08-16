# go-closer-chain

串联 Close 中断

internal/closekit/chain.go 的 CloseAll：遇第一个非 nil 错误立即 return，未继续关闭后续
