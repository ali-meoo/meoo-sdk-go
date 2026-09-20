/*
 * Meoo Open API Go SDK —— 手写高层 Runtime（package client）。
 *
 * 公开模型别名：把门面签名与文档用到的生成模型 re-export 到 meoo 包，使用者写
 * meoo.Project / meoo.AgentRun 即可，无需直接依赖内部 generated 包（对齐 Python 顶层
 * re-export 与 go-github 的单包公开面）。Go type alias 与 generated 中的类型完全等同，
 * meoo.Project 与 generated.Project 可互换、零转换开销。generated 仍是覆盖全部 operation
 * 的低层逃生舱（见 Client.Generated()）；此处只收敛门面直接暴露的核心模型，不逐一搬运
 * 全部生成类型。
 */

package client

import "gitlab.alibaba-inc.com/oneday/meoo-sdk-go/internal/generated"

type (
	// Project 是项目资源（re-export 自内部 generated.Project）。
	Project = generated.Project
	// AgentRun 是一次 Agent Run（re-export 自内部 generated.AgentRun）。
	AgentRun = generated.AgentRun
	// AgentRunStartRequest 是启动 Run 的请求体，可作为 Agent().Start 的 input
	// （re-export 自内部 generated.AgentRunStartRequest）。
	AgentRunStartRequest = generated.AgentRunStartRequest
	// AgentRunTerminalEvent 是 Run 终态 SSE 事件，EventStream + AgentEvent.DataAs 的常用目标
	// （re-export 自内部 generated.AgentRunTerminalEvent）。
	AgentRunTerminalEvent = generated.AgentRunTerminalEvent
)
