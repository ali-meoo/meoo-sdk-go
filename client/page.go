/*
 * Meoo Open API Go SDK —— 手写高层 Runtime（package client）。
 *
 * 分页结果封装，语义对齐 Java com.meoo.runtime.Page 与 TypeScript Page。泛型参数 T 通常是
 * 生成模型（例如 generated.Project）。
 */

package client

// Page 是一页结果；NextPageToken 为空表示没有下一页。
type Page[T any] struct {
	// Items 是本页条目，永不为 nil（无数据时为空切片）。
	Items []T
	// NextPageToken 是下一页游标，空串表示没有下一页。
	NextPageToken string
}

// HasNextPage 报告是否还有下一页。
func (p Page[T]) HasNextPage() bool { return p.NextPageToken != "" }
