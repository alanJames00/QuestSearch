// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.3
// 	protoc        v5.29.3
// source: questions.proto

package question

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Block structure for Anagram Type questions
type Block struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Text          string                 `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	ShowInOption  bool                   `protobuf:"varint,2,opt,name=show_in_option,json=showInOption,proto3" json:"show_in_option,omitempty"`
	IsAnswer      bool                   `protobuf:"varint,3,opt,name=is_answer,json=isAnswer,proto3" json:"is_answer,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Block) Reset() {
	*x = Block{}
	mi := &file_questions_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Block) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Block) ProtoMessage() {}

func (x *Block) ProtoReflect() protoreflect.Message {
	mi := &file_questions_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Block.ProtoReflect.Descriptor instead.
func (*Block) Descriptor() ([]byte, []int) {
	return file_questions_proto_rawDescGZIP(), []int{0}
}

func (x *Block) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *Block) GetShowInOption() bool {
	if x != nil {
		return x.ShowInOption
	}
	return false
}

func (x *Block) GetIsAnswer() bool {
	if x != nil {
		return x.IsAnswer
	}
	return false
}

// Option structure for MCQ Type questions
type Option struct {
	state           protoimpl.MessageState `protogen:"open.v1"`
	Text            string                 `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	IsCorrectAnswer bool                   `protobuf:"varint,2,opt,name=is_correct_answer,json=isCorrectAnswer,proto3" json:"is_correct_answer,omitempty"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *Option) Reset() {
	*x = Option{}
	mi := &file_questions_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Option) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Option) ProtoMessage() {}

func (x *Option) ProtoReflect() protoreflect.Message {
	mi := &file_questions_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Option.ProtoReflect.Descriptor instead.
func (*Option) Descriptor() ([]byte, []int) {
	return file_questions_proto_rawDescGZIP(), []int{1}
}

func (x *Option) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *Option) GetIsCorrectAnswer() bool {
	if x != nil {
		return x.IsCorrectAnswer
	}
	return false
}

// ContentOnly Question structure
type ContentOnlyQuestion struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ContentOnlyQuestion) Reset() {
	*x = ContentOnlyQuestion{}
	mi := &file_questions_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ContentOnlyQuestion) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContentOnlyQuestion) ProtoMessage() {}

func (x *ContentOnlyQuestion) ProtoReflect() protoreflect.Message {
	mi := &file_questions_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContentOnlyQuestion.ProtoReflect.Descriptor instead.
func (*ContentOnlyQuestion) Descriptor() ([]byte, []int) {
	return file_questions_proto_rawDescGZIP(), []int{2}
}

type ConversationQuestion struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ConversationQuestion) Reset() {
	*x = ConversationQuestion{}
	mi := &file_questions_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ConversationQuestion) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConversationQuestion) ProtoMessage() {}

func (x *ConversationQuestion) ProtoReflect() protoreflect.Message {
	mi := &file_questions_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConversationQuestion.ProtoReflect.Descriptor instead.
func (*ConversationQuestion) Descriptor() ([]byte, []int) {
	return file_questions_proto_rawDescGZIP(), []int{3}
}

type AnagramQuestion struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	AnagramType   string                 `protobuf:"bytes,1,opt,name=anagram_type,json=anagramType,proto3" json:"anagram_type,omitempty"`
	Blocks        []*Block               `protobuf:"bytes,2,rep,name=blocks,proto3" json:"blocks,omitempty"`
	Solution      string                 `protobuf:"bytes,3,opt,name=solution,proto3" json:"solution,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AnagramQuestion) Reset() {
	*x = AnagramQuestion{}
	mi := &file_questions_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AnagramQuestion) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnagramQuestion) ProtoMessage() {}

func (x *AnagramQuestion) ProtoReflect() protoreflect.Message {
	mi := &file_questions_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnagramQuestion.ProtoReflect.Descriptor instead.
func (*AnagramQuestion) Descriptor() ([]byte, []int) {
	return file_questions_proto_rawDescGZIP(), []int{4}
}

func (x *AnagramQuestion) GetAnagramType() string {
	if x != nil {
		return x.AnagramType
	}
	return ""
}

func (x *AnagramQuestion) GetBlocks() []*Block {
	if x != nil {
		return x.Blocks
	}
	return nil
}

func (x *AnagramQuestion) GetSolution() string {
	if x != nil {
		return x.Solution
	}
	return ""
}

type MCQQuestion struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Options       []*Option              `protobuf:"bytes,1,rep,name=options,proto3" json:"options,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MCQQuestion) Reset() {
	*x = MCQQuestion{}
	mi := &file_questions_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MCQQuestion) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MCQQuestion) ProtoMessage() {}

func (x *MCQQuestion) ProtoReflect() protoreflect.Message {
	mi := &file_questions_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MCQQuestion.ProtoReflect.Descriptor instead.
func (*MCQQuestion) Descriptor() ([]byte, []int) {
	return file_questions_proto_rawDescGZIP(), []int{5}
}

func (x *MCQQuestion) GetOptions() []*Option {
	if x != nil {
		return x.Options
	}
	return nil
}

type ReadAlongQuestion struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ReadAlongQuestion) Reset() {
	*x = ReadAlongQuestion{}
	mi := &file_questions_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ReadAlongQuestion) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadAlongQuestion) ProtoMessage() {}

func (x *ReadAlongQuestion) ProtoReflect() protoreflect.Message {
	mi := &file_questions_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadAlongQuestion.ProtoReflect.Descriptor instead.
func (*ReadAlongQuestion) Descriptor() ([]byte, []int) {
	return file_questions_proto_rawDescGZIP(), []int{6}
}

type Question struct {
	state     protoimpl.MessageState `protogen:"open.v1"`
	Id        string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Type      string                 `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	Title     string                 `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	SiblingId string                 `protobuf:"bytes,4,opt,name=sibling_id,json=siblingId,proto3" json:"sibling_id,omitempty"`
	// Types that are valid to be assigned to QuestionDetails:
	//
	//	*Question_ContentOnly
	//	*Question_Conversation
	//	*Question_Anagram
	//	*Question_Mcq
	//	*Question_ReadAlong
	QuestionDetails isQuestion_QuestionDetails `protobuf_oneof:"question_details"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *Question) Reset() {
	*x = Question{}
	mi := &file_questions_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Question) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Question) ProtoMessage() {}

func (x *Question) ProtoReflect() protoreflect.Message {
	mi := &file_questions_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Question.ProtoReflect.Descriptor instead.
func (*Question) Descriptor() ([]byte, []int) {
	return file_questions_proto_rawDescGZIP(), []int{7}
}

func (x *Question) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Question) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Question) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Question) GetSiblingId() string {
	if x != nil {
		return x.SiblingId
	}
	return ""
}

func (x *Question) GetQuestionDetails() isQuestion_QuestionDetails {
	if x != nil {
		return x.QuestionDetails
	}
	return nil
}

func (x *Question) GetContentOnly() *ContentOnlyQuestion {
	if x != nil {
		if x, ok := x.QuestionDetails.(*Question_ContentOnly); ok {
			return x.ContentOnly
		}
	}
	return nil
}

func (x *Question) GetConversation() *ConversationQuestion {
	if x != nil {
		if x, ok := x.QuestionDetails.(*Question_Conversation); ok {
			return x.Conversation
		}
	}
	return nil
}

func (x *Question) GetAnagram() *AnagramQuestion {
	if x != nil {
		if x, ok := x.QuestionDetails.(*Question_Anagram); ok {
			return x.Anagram
		}
	}
	return nil
}

func (x *Question) GetMcq() *MCQQuestion {
	if x != nil {
		if x, ok := x.QuestionDetails.(*Question_Mcq); ok {
			return x.Mcq
		}
	}
	return nil
}

func (x *Question) GetReadAlong() *ReadAlongQuestion {
	if x != nil {
		if x, ok := x.QuestionDetails.(*Question_ReadAlong); ok {
			return x.ReadAlong
		}
	}
	return nil
}

type isQuestion_QuestionDetails interface {
	isQuestion_QuestionDetails()
}

type Question_ContentOnly struct {
	ContentOnly *ContentOnlyQuestion `protobuf:"bytes,5,opt,name=content_only,json=contentOnly,proto3,oneof"`
}

type Question_Conversation struct {
	Conversation *ConversationQuestion `protobuf:"bytes,6,opt,name=conversation,proto3,oneof"`
}

type Question_Anagram struct {
	Anagram *AnagramQuestion `protobuf:"bytes,7,opt,name=anagram,proto3,oneof"`
}

type Question_Mcq struct {
	Mcq *MCQQuestion `protobuf:"bytes,8,opt,name=mcq,proto3,oneof"`
}

type Question_ReadAlong struct {
	ReadAlong *ReadAlongQuestion `protobuf:"bytes,9,opt,name=read_along,json=readAlong,proto3,oneof"`
}

func (*Question_ContentOnly) isQuestion_QuestionDetails() {}

func (*Question_Conversation) isQuestion_QuestionDetails() {}

func (*Question_Anagram) isQuestion_QuestionDetails() {}

func (*Question_Mcq) isQuestion_QuestionDetails() {}

func (*Question_ReadAlong) isQuestion_QuestionDetails() {}

// request to get questions
type GetQuestionsRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Page          int32                  `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	Limit         int32                  `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetQuestionsRequest) Reset() {
	*x = GetQuestionsRequest{}
	mi := &file_questions_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetQuestionsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetQuestionsRequest) ProtoMessage() {}

func (x *GetQuestionsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_questions_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetQuestionsRequest.ProtoReflect.Descriptor instead.
func (*GetQuestionsRequest) Descriptor() ([]byte, []int) {
	return file_questions_proto_rawDescGZIP(), []int{8}
}

func (x *GetQuestionsRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *GetQuestionsRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type GetQuestionsResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Questions     []*Question            `protobuf:"bytes,1,rep,name=questions,proto3" json:"questions,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetQuestionsResponse) Reset() {
	*x = GetQuestionsResponse{}
	mi := &file_questions_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetQuestionsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetQuestionsResponse) ProtoMessage() {}

func (x *GetQuestionsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_questions_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetQuestionsResponse.ProtoReflect.Descriptor instead.
func (*GetQuestionsResponse) Descriptor() ([]byte, []int) {
	return file_questions_proto_rawDescGZIP(), []int{9}
}

func (x *GetQuestionsResponse) GetQuestions() []*Question {
	if x != nil {
		return x.Questions
	}
	return nil
}

type GetQuestionByIdRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetQuestionByIdRequest) Reset() {
	*x = GetQuestionByIdRequest{}
	mi := &file_questions_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetQuestionByIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetQuestionByIdRequest) ProtoMessage() {}

func (x *GetQuestionByIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_questions_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetQuestionByIdRequest.ProtoReflect.Descriptor instead.
func (*GetQuestionByIdRequest) Descriptor() ([]byte, []int) {
	return file_questions_proto_rawDescGZIP(), []int{10}
}

func (x *GetQuestionByIdRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetQuestionByIdResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Question      *Question              `protobuf:"bytes,1,opt,name=question,proto3" json:"question,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetQuestionByIdResponse) Reset() {
	*x = GetQuestionByIdResponse{}
	mi := &file_questions_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetQuestionByIdResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetQuestionByIdResponse) ProtoMessage() {}

func (x *GetQuestionByIdResponse) ProtoReflect() protoreflect.Message {
	mi := &file_questions_proto_msgTypes[11]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetQuestionByIdResponse.ProtoReflect.Descriptor instead.
func (*GetQuestionByIdResponse) Descriptor() ([]byte, []int) {
	return file_questions_proto_rawDescGZIP(), []int{11}
}

func (x *GetQuestionByIdResponse) GetQuestion() *Question {
	if x != nil {
		return x.Question
	}
	return nil
}

type SearchQuestionRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	SearchTerm    string                 `protobuf:"bytes,1,opt,name=search_term,json=searchTerm,proto3" json:"search_term,omitempty"`
	Page          int32                  `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty"`
	Limit         int32                  `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SearchQuestionRequest) Reset() {
	*x = SearchQuestionRequest{}
	mi := &file_questions_proto_msgTypes[12]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SearchQuestionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchQuestionRequest) ProtoMessage() {}

func (x *SearchQuestionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_questions_proto_msgTypes[12]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchQuestionRequest.ProtoReflect.Descriptor instead.
func (*SearchQuestionRequest) Descriptor() ([]byte, []int) {
	return file_questions_proto_rawDescGZIP(), []int{12}
}

func (x *SearchQuestionRequest) GetSearchTerm() string {
	if x != nil {
		return x.SearchTerm
	}
	return ""
}

func (x *SearchQuestionRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *SearchQuestionRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

var File_questions_proto protoreflect.FileDescriptor

var file_questions_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x09, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x5e, 0x0a, 0x05,
	0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x12, 0x24, 0x0a, 0x0e, 0x73, 0x68, 0x6f,
	0x77, 0x5f, 0x69, 0x6e, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x0c, 0x73, 0x68, 0x6f, 0x77, 0x49, 0x6e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x1b, 0x0a, 0x09, 0x69, 0x73, 0x5f, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x41, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x22, 0x48, 0x0a, 0x06,
	0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x12, 0x2a, 0x0a, 0x11, 0x69, 0x73,
	0x5f, 0x63, 0x6f, 0x72, 0x72, 0x65, 0x63, 0x74, 0x5f, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0f, 0x69, 0x73, 0x43, 0x6f, 0x72, 0x72, 0x65, 0x63, 0x74,
	0x41, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x22, 0x15, 0x0a, 0x13, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x4f, 0x6e, 0x6c, 0x79, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x16, 0x0a,
	0x14, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x51, 0x75, 0x65,
	0x73, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x7a, 0x0a, 0x0f, 0x41, 0x6e, 0x61, 0x67, 0x72, 0x61, 0x6d,
	0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x6e, 0x61, 0x67,
	0x72, 0x61, 0x6d, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x61, 0x6e, 0x61, 0x67, 0x72, 0x61, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x12, 0x28, 0x0a, 0x06, 0x62,
	0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x06, 0x62,
	0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f,
	0x6e, 0x22, 0x3a, 0x0a, 0x0b, 0x4d, 0x43, 0x51, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x2b, 0x0a, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x11, 0x2e, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x4f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x13, 0x0a,
	0x11, 0x52, 0x65, 0x61, 0x64, 0x41, 0x6c, 0x6f, 0x6e, 0x67, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69,
	0x6f, 0x6e, 0x22, 0xa6, 0x03, 0x0a, 0x08, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x69, 0x62,
	0x6c, 0x69, 0x6e, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73,
	0x69, 0x62, 0x6c, 0x69, 0x6e, 0x67, 0x49, 0x64, 0x12, 0x43, 0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x5f, 0x6f, 0x6e, 0x6c, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e,
	0x2e, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x4f, 0x6e, 0x6c, 0x79, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x00,
	0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x4f, 0x6e, 0x6c, 0x79, 0x12, 0x45, 0x0a,
	0x0c, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x51, 0x75, 0x65, 0x73,
	0x74, 0x69, 0x6f, 0x6e, 0x48, 0x00, 0x52, 0x0c, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x36, 0x0a, 0x07, 0x61, 0x6e, 0x61, 0x67, 0x72, 0x61, 0x6d, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x41, 0x6e, 0x61, 0x67, 0x72, 0x61, 0x6d, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f,
	0x6e, 0x48, 0x00, 0x52, 0x07, 0x61, 0x6e, 0x61, 0x67, 0x72, 0x61, 0x6d, 0x12, 0x2a, 0x0a, 0x03,
	0x6d, 0x63, 0x71, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x4d, 0x43, 0x51, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f,
	0x6e, 0x48, 0x00, 0x52, 0x03, 0x6d, 0x63, 0x71, 0x12, 0x3d, 0x0a, 0x0a, 0x72, 0x65, 0x61, 0x64,
	0x5f, 0x61, 0x6c, 0x6f, 0x6e, 0x67, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x41, 0x6c, 0x6f,
	0x6e, 0x67, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x00, 0x52, 0x09, 0x72, 0x65,
	0x61, 0x64, 0x41, 0x6c, 0x6f, 0x6e, 0x67, 0x42, 0x12, 0x0a, 0x10, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x22, 0x3f, 0x0a, 0x13, 0x47,
	0x65, 0x74, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0x49, 0x0a, 0x14,
	0x47, 0x65, 0x74, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x31, 0x0a, 0x09, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x28, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x51, 0x75,
	0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x22, 0x4a, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e,
	0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2f, 0x0a, 0x08,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13,
	0x2e, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x51, 0x75, 0x65, 0x73, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x08, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x62, 0x0a,
	0x15, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68,
	0x5f, 0x74, 0x65, 0x72, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x65, 0x61,
	0x72, 0x63, 0x68, 0x54, 0x65, 0x72, 0x6d, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c,
	0x69, 0x6d, 0x69, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69,
	0x74, 0x32, 0x91, 0x02, 0x0a, 0x0f, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4f, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x51, 0x75, 0x65, 0x73,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1e, 0x2e, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x47, 0x65, 0x74, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x47, 0x65, 0x74, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x58, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x51, 0x75, 0x65,
	0x73, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x79, 0x49, 0x64, 0x12, 0x21, 0x2e, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f,
	0x6e, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x51, 0x75, 0x65, 0x73,
	0x74, 0x69, 0x6f, 0x6e, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x53, 0x0a, 0x0e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x20, 0x2e, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x53,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x47, 0x65, 0x74, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x18, 0x5a, 0x16, 0x2e, 0x2f, 0x70, 0x62, 0x2f, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x3b, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_questions_proto_rawDescOnce sync.Once
	file_questions_proto_rawDescData = file_questions_proto_rawDesc
)

func file_questions_proto_rawDescGZIP() []byte {
	file_questions_proto_rawDescOnce.Do(func() {
		file_questions_proto_rawDescData = protoimpl.X.CompressGZIP(file_questions_proto_rawDescData)
	})
	return file_questions_proto_rawDescData
}

var file_questions_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_questions_proto_goTypes = []any{
	(*Block)(nil),                   // 0: questions.Block
	(*Option)(nil),                  // 1: questions.Option
	(*ContentOnlyQuestion)(nil),     // 2: questions.ContentOnlyQuestion
	(*ConversationQuestion)(nil),    // 3: questions.ConversationQuestion
	(*AnagramQuestion)(nil),         // 4: questions.AnagramQuestion
	(*MCQQuestion)(nil),             // 5: questions.MCQQuestion
	(*ReadAlongQuestion)(nil),       // 6: questions.ReadAlongQuestion
	(*Question)(nil),                // 7: questions.Question
	(*GetQuestionsRequest)(nil),     // 8: questions.GetQuestionsRequest
	(*GetQuestionsResponse)(nil),    // 9: questions.GetQuestionsResponse
	(*GetQuestionByIdRequest)(nil),  // 10: questions.GetQuestionByIdRequest
	(*GetQuestionByIdResponse)(nil), // 11: questions.GetQuestionByIdResponse
	(*SearchQuestionRequest)(nil),   // 12: questions.SearchQuestionRequest
}
var file_questions_proto_depIdxs = []int32{
	0,  // 0: questions.AnagramQuestion.blocks:type_name -> questions.Block
	1,  // 1: questions.MCQQuestion.options:type_name -> questions.Option
	2,  // 2: questions.Question.content_only:type_name -> questions.ContentOnlyQuestion
	3,  // 3: questions.Question.conversation:type_name -> questions.ConversationQuestion
	4,  // 4: questions.Question.anagram:type_name -> questions.AnagramQuestion
	5,  // 5: questions.Question.mcq:type_name -> questions.MCQQuestion
	6,  // 6: questions.Question.read_along:type_name -> questions.ReadAlongQuestion
	7,  // 7: questions.GetQuestionsResponse.questions:type_name -> questions.Question
	7,  // 8: questions.GetQuestionByIdResponse.question:type_name -> questions.Question
	8,  // 9: questions.QuestionService.GetQuestions:input_type -> questions.GetQuestionsRequest
	10, // 10: questions.QuestionService.GetQuestionById:input_type -> questions.GetQuestionByIdRequest
	12, // 11: questions.QuestionService.SearchQuestion:input_type -> questions.SearchQuestionRequest
	9,  // 12: questions.QuestionService.GetQuestions:output_type -> questions.GetQuestionsResponse
	11, // 13: questions.QuestionService.GetQuestionById:output_type -> questions.GetQuestionByIdResponse
	9,  // 14: questions.QuestionService.SearchQuestion:output_type -> questions.GetQuestionsResponse
	12, // [12:15] is the sub-list for method output_type
	9,  // [9:12] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_questions_proto_init() }
func file_questions_proto_init() {
	if File_questions_proto != nil {
		return
	}
	file_questions_proto_msgTypes[7].OneofWrappers = []any{
		(*Question_ContentOnly)(nil),
		(*Question_Conversation)(nil),
		(*Question_Anagram)(nil),
		(*Question_Mcq)(nil),
		(*Question_ReadAlong)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_questions_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_questions_proto_goTypes,
		DependencyIndexes: file_questions_proto_depIdxs,
		MessageInfos:      file_questions_proto_msgTypes,
	}.Build()
	File_questions_proto = out.File
	file_questions_proto_rawDesc = nil
	file_questions_proto_goTypes = nil
	file_questions_proto_depIdxs = nil
}
