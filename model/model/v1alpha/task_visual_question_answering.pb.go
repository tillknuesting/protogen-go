// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: model/model/v1alpha/task_visual_question_answering.proto

package modelv1alpha

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	structpb "google.golang.org/protobuf/types/known/structpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// VisualQuestionAnsweringInput represents the input of visaul question answering task
type VisualQuestionAnsweringInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The prompt text
	Prompt string `protobuf:"bytes,1,opt,name=prompt,proto3" json:"prompt,omitempty"`
	// The prompt images
	PromptImages []*PromptImage `protobuf:"bytes,2,rep,name=prompt_images,json=promptImages,proto3" json:"prompt_images,omitempty"`
	// The chat history
	ChatHistory []*Message `protobuf:"bytes,3,rep,name=chat_history,json=chatHistory,proto3" json:"chat_history,omitempty"`
	// The system message
	SystemMessage *string `protobuf:"bytes,4,opt,name=system_message,json=systemMessage,proto3,oneof" json:"system_message,omitempty"`
	// The maximum number of tokens for model to generate
	MaxNewTokens *int32 `protobuf:"varint,5,opt,name=max_new_tokens,json=maxNewTokens,proto3,oneof" json:"max_new_tokens,omitempty"`
	// The temperature for sampling
	Temperature *float32 `protobuf:"fixed32,6,opt,name=temperature,proto3,oneof" json:"temperature,omitempty"`
	// Top k for sampling
	TopK *int32 `protobuf:"varint,7,opt,name=top_k,json=topK,proto3,oneof" json:"top_k,omitempty"`
	// The seed
	Seed *int32 `protobuf:"varint,8,opt,name=seed,proto3,oneof" json:"seed,omitempty"`
	// The extra parameters
	ExtraParams *structpb.Struct `protobuf:"bytes,9,opt,name=extra_params,json=extraParams,proto3" json:"extra_params,omitempty"`
}

func (x *VisualQuestionAnsweringInput) Reset() {
	*x = VisualQuestionAnsweringInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_model_model_v1alpha_task_visual_question_answering_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VisualQuestionAnsweringInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VisualQuestionAnsweringInput) ProtoMessage() {}

func (x *VisualQuestionAnsweringInput) ProtoReflect() protoreflect.Message {
	mi := &file_model_model_v1alpha_task_visual_question_answering_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VisualQuestionAnsweringInput.ProtoReflect.Descriptor instead.
func (*VisualQuestionAnsweringInput) Descriptor() ([]byte, []int) {
	return file_model_model_v1alpha_task_visual_question_answering_proto_rawDescGZIP(), []int{0}
}

func (x *VisualQuestionAnsweringInput) GetPrompt() string {
	if x != nil {
		return x.Prompt
	}
	return ""
}

func (x *VisualQuestionAnsweringInput) GetPromptImages() []*PromptImage {
	if x != nil {
		return x.PromptImages
	}
	return nil
}

func (x *VisualQuestionAnsweringInput) GetChatHistory() []*Message {
	if x != nil {
		return x.ChatHistory
	}
	return nil
}

func (x *VisualQuestionAnsweringInput) GetSystemMessage() string {
	if x != nil && x.SystemMessage != nil {
		return *x.SystemMessage
	}
	return ""
}

func (x *VisualQuestionAnsweringInput) GetMaxNewTokens() int32 {
	if x != nil && x.MaxNewTokens != nil {
		return *x.MaxNewTokens
	}
	return 0
}

func (x *VisualQuestionAnsweringInput) GetTemperature() float32 {
	if x != nil && x.Temperature != nil {
		return *x.Temperature
	}
	return 0
}

func (x *VisualQuestionAnsweringInput) GetTopK() int32 {
	if x != nil && x.TopK != nil {
		return *x.TopK
	}
	return 0
}

func (x *VisualQuestionAnsweringInput) GetSeed() int32 {
	if x != nil && x.Seed != nil {
		return *x.Seed
	}
	return 0
}

func (x *VisualQuestionAnsweringInput) GetExtraParams() *structpb.Struct {
	if x != nil {
		return x.ExtraParams
	}
	return nil
}

// VisualQuestionAnsweringOutput represents the output of visaul question answering task
type VisualQuestionAnsweringOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The text generated by the model
	Text string `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
}

func (x *VisualQuestionAnsweringOutput) Reset() {
	*x = VisualQuestionAnsweringOutput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_model_model_v1alpha_task_visual_question_answering_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VisualQuestionAnsweringOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VisualQuestionAnsweringOutput) ProtoMessage() {}

func (x *VisualQuestionAnsweringOutput) ProtoReflect() protoreflect.Message {
	mi := &file_model_model_v1alpha_task_visual_question_answering_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VisualQuestionAnsweringOutput.ProtoReflect.Descriptor instead.
func (*VisualQuestionAnsweringOutput) Descriptor() ([]byte, []int) {
	return file_model_model_v1alpha_task_visual_question_answering_proto_rawDescGZIP(), []int{1}
}

func (x *VisualQuestionAnsweringOutput) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

var File_model_model_v1alpha_task_visual_question_answering_proto protoreflect.FileDescriptor

var file_model_model_v1alpha_task_visual_question_answering_proto_rawDesc = []byte{
	0x0a, 0x38, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x76, 0x69, 0x73, 0x75, 0x61,
	0x6c, 0x5f, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x61, 0x6e, 0x73, 0x77, 0x65,
	0x72, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x1a,
	0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c,
	0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xa1, 0x04, 0x0a, 0x1c, 0x56, 0x69, 0x73, 0x75, 0x61, 0x6c, 0x51, 0x75, 0x65, 0x73, 0x74,
	0x69, 0x6f, 0x6e, 0x41, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x69, 0x6e, 0x67, 0x49, 0x6e, 0x70, 0x75,
	0x74, 0x12, 0x1b, 0x0a, 0x06, 0x70, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x06, 0x70, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x12, 0x4a,
	0x0a, 0x0d, 0x70, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x50, 0x72, 0x6f, 0x6d,
	0x70, 0x74, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x0c, 0x70, 0x72,
	0x6f, 0x6d, 0x70, 0x74, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x12, 0x44, 0x0a, 0x0c, 0x63, 0x68,
	0x61, 0x74, 0x5f, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x1c, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x42, 0x03,
	0xe0, 0x41, 0x01, 0x52, 0x0b, 0x63, 0x68, 0x61, 0x74, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79,
	0x12, 0x2f, 0x0a, 0x0e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x48, 0x00, 0x52,
	0x0d, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x88, 0x01,
	0x01, 0x12, 0x2e, 0x0a, 0x0e, 0x6d, 0x61, 0x78, 0x5f, 0x6e, 0x65, 0x77, 0x5f, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x48, 0x01,
	0x52, 0x0c, 0x6d, 0x61, 0x78, 0x4e, 0x65, 0x77, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x88, 0x01,
	0x01, 0x12, 0x2a, 0x0a, 0x0b, 0x74, 0x65, 0x6d, 0x70, 0x65, 0x72, 0x61, 0x74, 0x75, 0x72, 0x65,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x02, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x48, 0x02, 0x52, 0x0b, 0x74,
	0x65, 0x6d, 0x70, 0x65, 0x72, 0x61, 0x74, 0x75, 0x72, 0x65, 0x88, 0x01, 0x01, 0x12, 0x1d, 0x0a,
	0x05, 0x74, 0x6f, 0x70, 0x5f, 0x6b, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x42, 0x03, 0xe0, 0x41,
	0x01, 0x48, 0x03, 0x52, 0x04, 0x74, 0x6f, 0x70, 0x4b, 0x88, 0x01, 0x01, 0x12, 0x1c, 0x0a, 0x04,
	0x73, 0x65, 0x65, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x48,
	0x04, 0x52, 0x04, 0x73, 0x65, 0x65, 0x64, 0x88, 0x01, 0x01, 0x12, 0x3f, 0x0a, 0x0c, 0x65, 0x78,
	0x74, 0x72, 0x61, 0x5f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x0b,
	0x65, 0x78, 0x74, 0x72, 0x61, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x42, 0x11, 0x0a, 0x0f, 0x5f,
	0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x42, 0x11,
	0x0a, 0x0f, 0x5f, 0x6d, 0x61, 0x78, 0x5f, 0x6e, 0x65, 0x77, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x73, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x74, 0x65, 0x6d, 0x70, 0x65, 0x72, 0x61, 0x74, 0x75, 0x72,
	0x65, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x74, 0x6f, 0x70, 0x5f, 0x6b, 0x42, 0x07, 0x0a, 0x05, 0x5f,
	0x73, 0x65, 0x65, 0x64, 0x22, 0x38, 0x0a, 0x1d, 0x56, 0x69, 0x73, 0x75, 0x61, 0x6c, 0x51, 0x75,
	0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x69, 0x6e, 0x67, 0x4f,
	0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x17, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x42, 0xed,
	0x01, 0x0a, 0x17, 0x63, 0x6f, 0x6d, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x42, 0x20, 0x54, 0x61, 0x73, 0x6b,
	0x56, 0x69, 0x73, 0x75, 0x61, 0x6c, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x6e,
	0x73, 0x77, 0x65, 0x72, 0x69, 0x6e, 0x67, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x42,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6e, 0x73, 0x74, 0x69,
	0x6c, 0x6c, 0x2d, 0x61, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x67, 0x65, 0x6e, 0x2d, 0x67,
	0x6f, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x3b, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0xa2, 0x02, 0x03, 0x4d, 0x4d, 0x58, 0xaa, 0x02, 0x13, 0x4d, 0x6f, 0x64, 0x65, 0x6c,
	0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0xca, 0x02,
	0x13, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x5c, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x5c, 0x56, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0xe2, 0x02, 0x1f, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x5c, 0x4d, 0x6f, 0x64,
	0x65, 0x6c, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x15, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x3a, 0x3a,
	0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x3a, 0x3a, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_model_model_v1alpha_task_visual_question_answering_proto_rawDescOnce sync.Once
	file_model_model_v1alpha_task_visual_question_answering_proto_rawDescData = file_model_model_v1alpha_task_visual_question_answering_proto_rawDesc
)

func file_model_model_v1alpha_task_visual_question_answering_proto_rawDescGZIP() []byte {
	file_model_model_v1alpha_task_visual_question_answering_proto_rawDescOnce.Do(func() {
		file_model_model_v1alpha_task_visual_question_answering_proto_rawDescData = protoimpl.X.CompressGZIP(file_model_model_v1alpha_task_visual_question_answering_proto_rawDescData)
	})
	return file_model_model_v1alpha_task_visual_question_answering_proto_rawDescData
}

var file_model_model_v1alpha_task_visual_question_answering_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_model_model_v1alpha_task_visual_question_answering_proto_goTypes = []interface{}{
	(*VisualQuestionAnsweringInput)(nil),  // 0: model.model.v1alpha.VisualQuestionAnsweringInput
	(*VisualQuestionAnsweringOutput)(nil), // 1: model.model.v1alpha.VisualQuestionAnsweringOutput
	(*PromptImage)(nil),                   // 2: model.model.v1alpha.PromptImage
	(*Message)(nil),                       // 3: model.model.v1alpha.Message
	(*structpb.Struct)(nil),               // 4: google.protobuf.Struct
}
var file_model_model_v1alpha_task_visual_question_answering_proto_depIdxs = []int32{
	2, // 0: model.model.v1alpha.VisualQuestionAnsweringInput.prompt_images:type_name -> model.model.v1alpha.PromptImage
	3, // 1: model.model.v1alpha.VisualQuestionAnsweringInput.chat_history:type_name -> model.model.v1alpha.Message
	4, // 2: model.model.v1alpha.VisualQuestionAnsweringInput.extra_params:type_name -> google.protobuf.Struct
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_model_model_v1alpha_task_visual_question_answering_proto_init() }
func file_model_model_v1alpha_task_visual_question_answering_proto_init() {
	if File_model_model_v1alpha_task_visual_question_answering_proto != nil {
		return
	}
	file_model_model_v1alpha_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_model_model_v1alpha_task_visual_question_answering_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VisualQuestionAnsweringInput); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_model_model_v1alpha_task_visual_question_answering_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VisualQuestionAnsweringOutput); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_model_model_v1alpha_task_visual_question_answering_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_model_model_v1alpha_task_visual_question_answering_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_model_model_v1alpha_task_visual_question_answering_proto_goTypes,
		DependencyIndexes: file_model_model_v1alpha_task_visual_question_answering_proto_depIdxs,
		MessageInfos:      file_model_model_v1alpha_task_visual_question_answering_proto_msgTypes,
	}.Build()
	File_model_model_v1alpha_task_visual_question_answering_proto = out.File
	file_model_model_v1alpha_task_visual_question_answering_proto_rawDesc = nil
	file_model_model_v1alpha_task_visual_question_answering_proto_goTypes = nil
	file_model_model_v1alpha_task_visual_question_answering_proto_depIdxs = nil
}
