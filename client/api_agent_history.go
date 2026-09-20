package client

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// AgentHistoryAPIService AgentHistoryAPI service
type AgentHistoryAPIService service

type ApiListAgentConversationMessagesRequest struct {
	ctx            context.Context
	ApiService     *AgentHistoryAPIService
	projectId      string
	conversationId string
	pageSize       *int32
	pageToken      *string
	order          *string
}

// 本页最多返回数量。0 或不传使用默认值 20；超过 100 时按 100 处理。
func (r ApiListAgentConversationMessagesRequest) PageSize(pageSize int32) ApiListAgentConversationMessagesRequest {
	r.pageSize = &pageSize
	return r
}

// 上一页返回的 next_page_token；调用方不得解析或修改。
func (r ApiListAgentConversationMessagesRequest) PageToken(pageToken string) ApiListAgentConversationMessagesRequest {
	r.pageToken = &pageToken
	return r
}

// 返回顺序。page_token 与请求时的 order 绑定； 跨 order 复用 page_token 会返回 invalid_page_token。
func (r ApiListAgentConversationMessagesRequest) Order(order string) ApiListAgentConversationMessagesRequest {
	r.order = &order
	return r
}

func (r ApiListAgentConversationMessagesRequest) Execute() (*AgentMessageListResponse, *http.Response, error) {
	return r.ApiService.ListAgentConversationMessagesExecute(r)
}

/*
ListAgentConversationMessages 分页查询指定 Agent 会话的公开消息

按创建时间返回指定会话中已落库的 user 和 assistant 文本消息；
order=asc（默认）为最旧到最新，order=desc 为最新到最旧，翻页方向随之向新或向旧。
服务端会移除 system-reminder、隐藏消息、工具调用、工具结果、metadata、
token 用量和内部诊断字段；该接口不会触发 Agent 运行。

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param projectId 项目的公开 URL ID。
	@param conversationId start/current 或会话列表返回的 conversation_id。
	@return ApiListAgentConversationMessagesRequest
*/
func (a *AgentHistoryAPIService) ListAgentConversationMessages(ctx context.Context, projectId string, conversationId string) ApiListAgentConversationMessagesRequest {
	return ApiListAgentConversationMessagesRequest{
		ApiService:     a,
		ctx:            ctx,
		projectId:      projectId,
		conversationId: conversationId,
	}
}

// Execute executes the request
//
//	@return AgentMessageListResponse
func (a *AgentHistoryAPIService) ListAgentConversationMessagesExecute(r ApiListAgentConversationMessagesRequest) (*AgentMessageListResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AgentMessageListResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AgentHistoryAPIService.ListAgentConversationMessages")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/open/v1/projects/{project_id}/agent/conversations/{conversation_id}/messages"
	localVarPath = strings.Replace(localVarPath, "{"+"project_id"+"}", url.PathEscape(parameterValueToString(r.projectId, "projectId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"conversation_id"+"}", url.PathEscape(parameterValueToString(r.conversationId, "conversationId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.projectId) < 1 {
		return localVarReturnValue, nil, reportError("projectId must have at least 1 elements")
	}
	if strlen(r.projectId) > 50 {
		return localVarReturnValue, nil, reportError("projectId must have less than 50 elements")
	}
	if strlen(r.conversationId) < 1 {
		return localVarReturnValue, nil, reportError("conversationId must have at least 1 elements")
	}
	if strlen(r.conversationId) > 64 {
		return localVarReturnValue, nil, reportError("conversationId must have less than 64 elements")
	}

	if r.pageSize != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page_size", r.pageSize, "form", "")
	} else {
		var defaultValue int32 = 20
		r.pageSize = &defaultValue
	}
	if r.pageToken != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page_token", r.pageToken, "form", "")
	}
	if r.order != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "order", r.order, "form", "")
	} else {
		var defaultValue string = "asc"
		r.order = &defaultValue
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "application/problem+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v Problem
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v Problem
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v Problem
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v Problem
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v Problem
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v Problem
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiListAgentConversationsRequest struct {
	ctx        context.Context
	ApiService *AgentHistoryAPIService
	projectId  string
	pageSize   *int32
	pageToken  *string
}

// 本页最多返回数量。0 或不传使用默认值 20；超过 100 时按 100 处理。
func (r ApiListAgentConversationsRequest) PageSize(pageSize int32) ApiListAgentConversationsRequest {
	r.pageSize = &pageSize
	return r
}

// 上一页返回的 next_page_token；调用方不得解析或修改。
func (r ApiListAgentConversationsRequest) PageToken(pageToken string) ApiListAgentConversationsRequest {
	r.pageToken = &pageToken
	return r
}

func (r ApiListAgentConversationsRequest) Execute() (*AgentConversationListResponse, *http.Response, error) {
	return r.ApiService.ListAgentConversationsExecute(r)
}

/*
ListAgentConversations 分页查询项目的 Agent 会话

按会话创建时间倒序返回项目已有的 Agent conversation_id。
查询不会创建新会话，也不会触发 Agent 运行。

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param projectId 项目的公开 URL ID。
	@return ApiListAgentConversationsRequest
*/
func (a *AgentHistoryAPIService) ListAgentConversations(ctx context.Context, projectId string) ApiListAgentConversationsRequest {
	return ApiListAgentConversationsRequest{
		ApiService: a,
		ctx:        ctx,
		projectId:  projectId,
	}
}

// Execute executes the request
//
//	@return AgentConversationListResponse
func (a *AgentHistoryAPIService) ListAgentConversationsExecute(r ApiListAgentConversationsRequest) (*AgentConversationListResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AgentConversationListResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AgentHistoryAPIService.ListAgentConversations")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/open/v1/projects/{project_id}/agent/conversations"
	localVarPath = strings.Replace(localVarPath, "{"+"project_id"+"}", url.PathEscape(parameterValueToString(r.projectId, "projectId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.projectId) < 1 {
		return localVarReturnValue, nil, reportError("projectId must have at least 1 elements")
	}
	if strlen(r.projectId) > 50 {
		return localVarReturnValue, nil, reportError("projectId must have less than 50 elements")
	}

	if r.pageSize != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page_size", r.pageSize, "form", "")
	} else {
		var defaultValue int32 = 20
		r.pageSize = &defaultValue
	}
	if r.pageToken != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page_token", r.pageToken, "form", "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "application/problem+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v Problem
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v Problem
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v Problem
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v Problem
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v Problem
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v Problem
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
