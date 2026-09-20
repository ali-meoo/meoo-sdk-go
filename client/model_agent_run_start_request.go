package client

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the AgentRunStartRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AgentRunStartRequest{}

// AgentRunStartRequest struct for AgentRunStartRequest
type AgentRunStartRequest struct {
	Message        string  `json:"message" validate:"regexp=\\\\S"`
	ConversationId *string `json:"conversation_id,omitempty" validate:"regexp=^[A-Za-z0-9_-]+$"`
	// 附件列表；图片、视频和普通文件合计最多 10 个。
	Attachments []AgentRunAttachment `json:"attachments,omitempty"`
	// 本次 Run 强制加载的 Skill 引用。不传时，有 conversation_id 则继承该会话 已有用户 Skill，否则继承项目最近会话已有用户 Skill；传 [] 时显式清除继承 用户 Skill。显式选择会保存到会话配置，供 blocking/resume 和后续未传 skills 的 Run 继承。解析后的 Skill 名称必须互不重复且不能与运行时内置 Skill 重名。服务端内置 Skill 不受 [] 影响。
	Skills []AgentRunSkillSelection `json:"skills,omitempty"`
	// 传 true 会为本次 Run 开启 YOLO，并将开启状态保存到 Agent 状态快照， 供同项目后续 Run 继承。false 或不传不会主动开启，也不能清除已经持久化的 true。YOLO 仅应用服务端预置自动决策；需要真人完成的 OAuth、Secret 或 Input 场景仍可能按现有策略取消或关闭。本字段不会改变 HTTP 响应或 SSE 事件 Schema。
	Yolo *bool `json:"yolo,omitempty"`
	// 本次 Run 使用的 canonical 模型 ID，必须精确来自当前项目的 agent/capabilities 响应；未知、内部别名或已下线模型返回 400。 省略时使用 capabilities.defaults.model，包括继续会话的请求。
	Model *string `json:"model,omitempty"`
	// 每次 Run 独立选择的性能档位：fast 优先较低延迟，standard 为均衡默认值， deep 使用更多推理能力且可能增加延迟和消耗。省略时使用 standard， 包括继续会话的请求。
	SpeedTier *string `json:"speed_tier,omitempty"`
	// 仅本次 Run 禁用全部 meoo-cli cloud 能力（包含读取、帮助、数据库、鉴权、 存储及云函数），不产生云能力确认卡。工具恢复、上下文压缩及子 Agent 继承本次策略，YOLO 不能绕过。下一次 Run（含同会话继续）不传则恢复 false。 外部 API 不受限制；缺少可用后端时使用 mock，并在交付时说明模拟部分。 不关闭或删除已有云资源，不改变 HTTP 响应或 SSE 事件 Schema。
	DisableCloud *bool `json:"disable_cloud,omitempty"`
}

type _AgentRunStartRequest AgentRunStartRequest

// NewAgentRunStartRequest instantiates a new AgentRunStartRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAgentRunStartRequest(message string) *AgentRunStartRequest {
	this := AgentRunStartRequest{}
	this.Message = message
	var disableCloud bool = false
	this.DisableCloud = &disableCloud
	return &this
}

// NewAgentRunStartRequestWithDefaults instantiates a new AgentRunStartRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAgentRunStartRequestWithDefaults() *AgentRunStartRequest {
	this := AgentRunStartRequest{}
	var disableCloud bool = false
	this.DisableCloud = &disableCloud
	return &this
}

// GetMessage returns the Message field value
func (o *AgentRunStartRequest) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *AgentRunStartRequest) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *AgentRunStartRequest) SetMessage(v string) {
	o.Message = v
}

// GetConversationId returns the ConversationId field value if set, zero value otherwise.
func (o *AgentRunStartRequest) GetConversationId() string {
	if o == nil || IsNil(o.ConversationId) {
		var ret string
		return ret
	}
	return *o.ConversationId
}

// GetConversationIdOk returns a tuple with the ConversationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentRunStartRequest) GetConversationIdOk() (*string, bool) {
	if o == nil || IsNil(o.ConversationId) {
		return nil, false
	}
	return o.ConversationId, true
}

// HasConversationId returns a boolean if a field has been set.
func (o *AgentRunStartRequest) HasConversationId() bool {
	if o != nil && !IsNil(o.ConversationId) {
		return true
	}

	return false
}

// SetConversationId gets a reference to the given string and assigns it to the ConversationId field.
func (o *AgentRunStartRequest) SetConversationId(v string) {
	o.ConversationId = &v
}

// GetAttachments returns the Attachments field value if set, zero value otherwise.
func (o *AgentRunStartRequest) GetAttachments() []AgentRunAttachment {
	if o == nil || IsNil(o.Attachments) {
		var ret []AgentRunAttachment
		return ret
	}
	return o.Attachments
}

// GetAttachmentsOk returns a tuple with the Attachments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentRunStartRequest) GetAttachmentsOk() ([]AgentRunAttachment, bool) {
	if o == nil || IsNil(o.Attachments) {
		return nil, false
	}
	return o.Attachments, true
}

// HasAttachments returns a boolean if a field has been set.
func (o *AgentRunStartRequest) HasAttachments() bool {
	if o != nil && !IsNil(o.Attachments) {
		return true
	}

	return false
}

// SetAttachments gets a reference to the given []AgentRunAttachment and assigns it to the Attachments field.
func (o *AgentRunStartRequest) SetAttachments(v []AgentRunAttachment) {
	o.Attachments = v
}

// GetSkills returns the Skills field value if set, zero value otherwise.
func (o *AgentRunStartRequest) GetSkills() []AgentRunSkillSelection {
	if o == nil || IsNil(o.Skills) {
		var ret []AgentRunSkillSelection
		return ret
	}
	return o.Skills
}

// GetSkillsOk returns a tuple with the Skills field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentRunStartRequest) GetSkillsOk() ([]AgentRunSkillSelection, bool) {
	if o == nil || IsNil(o.Skills) {
		return nil, false
	}
	return o.Skills, true
}

// HasSkills returns a boolean if a field has been set.
func (o *AgentRunStartRequest) HasSkills() bool {
	if o != nil && !IsNil(o.Skills) {
		return true
	}

	return false
}

// SetSkills gets a reference to the given []AgentRunSkillSelection and assigns it to the Skills field.
func (o *AgentRunStartRequest) SetSkills(v []AgentRunSkillSelection) {
	o.Skills = v
}

// GetYolo returns the Yolo field value if set, zero value otherwise.
func (o *AgentRunStartRequest) GetYolo() bool {
	if o == nil || IsNil(o.Yolo) {
		var ret bool
		return ret
	}
	return *o.Yolo
}

// GetYoloOk returns a tuple with the Yolo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentRunStartRequest) GetYoloOk() (*bool, bool) {
	if o == nil || IsNil(o.Yolo) {
		return nil, false
	}
	return o.Yolo, true
}

// HasYolo returns a boolean if a field has been set.
func (o *AgentRunStartRequest) HasYolo() bool {
	if o != nil && !IsNil(o.Yolo) {
		return true
	}

	return false
}

// SetYolo gets a reference to the given bool and assigns it to the Yolo field.
func (o *AgentRunStartRequest) SetYolo(v bool) {
	o.Yolo = &v
}

// GetModel returns the Model field value if set, zero value otherwise.
func (o *AgentRunStartRequest) GetModel() string {
	if o == nil || IsNil(o.Model) {
		var ret string
		return ret
	}
	return *o.Model
}

// GetModelOk returns a tuple with the Model field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentRunStartRequest) GetModelOk() (*string, bool) {
	if o == nil || IsNil(o.Model) {
		return nil, false
	}
	return o.Model, true
}

// HasModel returns a boolean if a field has been set.
func (o *AgentRunStartRequest) HasModel() bool {
	if o != nil && !IsNil(o.Model) {
		return true
	}

	return false
}

// SetModel gets a reference to the given string and assigns it to the Model field.
func (o *AgentRunStartRequest) SetModel(v string) {
	o.Model = &v
}

// GetSpeedTier returns the SpeedTier field value if set, zero value otherwise.
func (o *AgentRunStartRequest) GetSpeedTier() string {
	if o == nil || IsNil(o.SpeedTier) {
		var ret string
		return ret
	}
	return *o.SpeedTier
}

// GetSpeedTierOk returns a tuple with the SpeedTier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentRunStartRequest) GetSpeedTierOk() (*string, bool) {
	if o == nil || IsNil(o.SpeedTier) {
		return nil, false
	}
	return o.SpeedTier, true
}

// HasSpeedTier returns a boolean if a field has been set.
func (o *AgentRunStartRequest) HasSpeedTier() bool {
	if o != nil && !IsNil(o.SpeedTier) {
		return true
	}

	return false
}

// SetSpeedTier gets a reference to the given string and assigns it to the SpeedTier field.
func (o *AgentRunStartRequest) SetSpeedTier(v string) {
	o.SpeedTier = &v
}

// GetDisableCloud returns the DisableCloud field value if set, zero value otherwise.
func (o *AgentRunStartRequest) GetDisableCloud() bool {
	if o == nil || IsNil(o.DisableCloud) {
		var ret bool
		return ret
	}
	return *o.DisableCloud
}

// GetDisableCloudOk returns a tuple with the DisableCloud field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentRunStartRequest) GetDisableCloudOk() (*bool, bool) {
	if o == nil || IsNil(o.DisableCloud) {
		return nil, false
	}
	return o.DisableCloud, true
}

// HasDisableCloud returns a boolean if a field has been set.
func (o *AgentRunStartRequest) HasDisableCloud() bool {
	if o != nil && !IsNil(o.DisableCloud) {
		return true
	}

	return false
}

// SetDisableCloud gets a reference to the given bool and assigns it to the DisableCloud field.
func (o *AgentRunStartRequest) SetDisableCloud(v bool) {
	o.DisableCloud = &v
}

func (o AgentRunStartRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AgentRunStartRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["message"] = o.Message
	if !IsNil(o.ConversationId) {
		toSerialize["conversation_id"] = o.ConversationId
	}
	if !IsNil(o.Attachments) {
		toSerialize["attachments"] = o.Attachments
	}
	if !IsNil(o.Skills) {
		toSerialize["skills"] = o.Skills
	}
	if !IsNil(o.Yolo) {
		toSerialize["yolo"] = o.Yolo
	}
	if !IsNil(o.Model) {
		toSerialize["model"] = o.Model
	}
	if !IsNil(o.SpeedTier) {
		toSerialize["speed_tier"] = o.SpeedTier
	}
	if !IsNil(o.DisableCloud) {
		toSerialize["disable_cloud"] = o.DisableCloud
	}
	return toSerialize, nil
}

func (o *AgentRunStartRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"message",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varAgentRunStartRequest := _AgentRunStartRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAgentRunStartRequest)

	if err != nil {
		return err
	}

	*o = AgentRunStartRequest(varAgentRunStartRequest)

	return err
}

type NullableAgentRunStartRequest struct {
	value *AgentRunStartRequest
	isSet bool
}

func (v NullableAgentRunStartRequest) Get() *AgentRunStartRequest {
	return v.value
}

func (v *NullableAgentRunStartRequest) Set(val *AgentRunStartRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAgentRunStartRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAgentRunStartRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAgentRunStartRequest(val *AgentRunStartRequest) *NullableAgentRunStartRequest {
	return &NullableAgentRunStartRequest{value: val, isSet: true}
}

func (v NullableAgentRunStartRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAgentRunStartRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
