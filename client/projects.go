/*
 * Meoo Open API Go SDK — 项目资源。
 *
 * Projects() 封装最高频的项目 operation；其余 operation 可用 Client.Generated() 的完整 API
 * 客户端，或 Client.Transport() 直接发请求。
 *
 * 列表信封用局部结构 + 显式 json tag 解码，只依赖响应的 JSON 字段名，把底层模型字段命名
 * 变化的编译期风险降到最低。
 */

package client

import (
	"context"
	"encoding/json"
	"strconv"
)

// projectsCollectionPath 是项目集合路径。
const projectsCollectionPath = "/open/v1/projects"

// ProjectsResource 封装项目相关的高频 operation。
type ProjectsResource struct {
	transport *Transport
}

func newProjectsResource(transport *Transport) *ProjectsResource {
	return &ProjectsResource{transport: transport}
}

// CreateProjectParams 是创建项目的入参，对应 ProjectCreateRequest。
type CreateProjectParams struct {
	// Name 项目名（长度 1..100 且含非空白字符）。
	Name string
	// Type 项目类型：web、app 或 miniprogram；留空则不发送该字段，由服务端按默认 web 处理。
	Type string
}

// Create 创建项目；opts 携带 IdempotencyKey 时允许自动重试。
func (r *ProjectsResource) Create(ctx context.Context, params CreateProjectParams, opts *RequestOptions) (*Project, error) {
	body := map[string]interface{}{"name": params.Name}
	if params.Type != "" {
		body["type"] = params.Type
	}
	raw, err := r.transport.Request(ctx, "POST", projectsCollectionPath, body, opts)
	if err != nil {
		return nil, err
	}
	return decodeProject(raw)
}

// ListProjectsParams 是列出项目的可选查询参数；零值字段不会被拼进 query。
type ListProjectsParams struct {
	// PageSize 对应 page_size，nil 表示不传。
	PageSize *int
	// PageToken 对应 page_token，空串表示不传。
	PageToken string
	// Query 对应 query（名称模糊匹配），空串表示不传。
	Query string
}

// List 列出项目的一页。
func (r *ProjectsResource) List(ctx context.Context, params ListProjectsParams, opts *RequestOptions) (*Page[Project], error) {
	var query []queryParam
	if params.PageSize != nil {
		query = append(query, queryParam{key: "page_size", value: strconv.Itoa(*params.PageSize)})
	}
	if params.PageToken != "" {
		query = append(query, queryParam{key: "page_token", value: params.PageToken})
	}
	if params.Query != "" {
		query = append(query, queryParam{key: "query", value: params.Query})
	}

	raw, err := r.transport.Request(ctx, "GET", projectsCollectionPath+queryString(query...), nil, opts)
	if err != nil {
		return nil, err
	}

	// 局部信封结构：只依赖响应的 JSON 字段名。
	var envelope struct {
		Projects      []Project `json:"projects"`
		NextPageToken *string   `json:"next_page_token"`
	}
	if len(raw) > 0 {
		if err := json.Unmarshal(raw, &envelope); err != nil {
			return nil, newMeooError("failed to map response to ProjectListResponse", err)
		}
	}

	items := envelope.Projects
	if items == nil {
		items = []Project{}
	}
	page := &Page[Project]{Items: items}
	if envelope.NextPageToken != nil {
		page.NextPageToken = *envelope.NextPageToken
	}
	return page, nil
}

// Iter 返回惰性翻页迭代器；PageToken 由迭代器内部管理，params.PageToken 会被忽略。
// 用法对齐 sql.Rows：
//
//	it := client.Projects().Iter(ctx, meoo.ListProjectsParams{}, nil)
//	for it.Next() {
//	    project := it.Item()
//	    ...
//	}
//	if err := it.Err(); err != nil { ... }
func (r *ProjectsResource) Iter(ctx context.Context, params ListProjectsParams, opts *RequestOptions) *ProjectIterator {
	managed := params
	managed.PageToken = ""
	return &ProjectIterator{resource: r, ctx: ctx, params: managed, opts: opts}
}

// ProjectIterator 是项目的惰性翻页迭代器。非并发安全，应由单个 goroutine 使用。
type ProjectIterator struct {
	resource *ProjectsResource
	ctx      context.Context
	params   ListProjectsParams
	opts     *RequestOptions

	items     []Project
	pos       int
	exhausted bool
	item      *Project
	err       error
}

// Next 前进一步；有下一个项目返回 true。当前页耗尽时自动拉取下一页，服务端不再返回
// next_page_token 时结束。出错时返回 false，错误经 Err 取得。
func (it *ProjectIterator) Next() bool {
	if it.err != nil {
		return false
	}
	for {
		if it.pos < len(it.items) {
			it.item = &it.items[it.pos]
			it.pos++
			return true
		}
		if it.exhausted {
			return false
		}
		page, err := it.resource.List(it.ctx, it.params, it.opts)
		if err != nil {
			it.err = err
			return false
		}
		it.items = page.Items
		it.pos = 0
		it.params.PageToken = page.NextPageToken
		it.exhausted = !page.HasNextPage()
		// 继续循环以产出新页的第一个项目；空页且已耗尽时在上方返回 false。
	}
}

// Item 返回 Next 最近一次定位到的项目；仅在 Next 返回 true 后有效。
func (it *ProjectIterator) Item() *Project { return it.item }

// Err 返回迭代过程中出现的第一个错误（无错误时为 nil）。
func (it *ProjectIterator) Err() error { return it.err }

// decodeProject 把响应 JSON 映射为 Project；raw 为空（例如 204）时返回 (nil, nil)。
func decodeProject(raw json.RawMessage) (*Project, error) {
	if len(raw) == 0 {
		return nil, nil
	}
	var project Project
	if err := json.Unmarshal(raw, &project); err != nil {
		return nil, newMeooError("failed to map response to Project", err)
	}
	return &project, nil
}
