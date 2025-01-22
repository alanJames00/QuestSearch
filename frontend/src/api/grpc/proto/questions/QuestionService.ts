// Original file: src/api/grpc/proto/question.proto

import type * as grpc from '@grpc/grpc-js'
import type { MethodDefinition } from '@grpc/proto-loader'
import type { GetQuestionByIdRequest as _questions_GetQuestionByIdRequest, GetQuestionByIdRequest__Output as _questions_GetQuestionByIdRequest__Output } from '../questions/GetQuestionByIdRequest';
import type { GetQuestionByIdResponse as _questions_GetQuestionByIdResponse, GetQuestionByIdResponse__Output as _questions_GetQuestionByIdResponse__Output } from '../questions/GetQuestionByIdResponse';
import type { GetQuestionsRequest as _questions_GetQuestionsRequest, GetQuestionsRequest__Output as _questions_GetQuestionsRequest__Output } from '../questions/GetQuestionsRequest';
import type { GetQuestionsResponse as _questions_GetQuestionsResponse, GetQuestionsResponse__Output as _questions_GetQuestionsResponse__Output } from '../questions/GetQuestionsResponse';
import type { SearchQuestionRequest as _questions_SearchQuestionRequest, SearchQuestionRequest__Output as _questions_SearchQuestionRequest__Output } from '../questions/SearchQuestionRequest';

export interface QuestionServiceClient extends grpc.Client {
  GetQuestionById(argument: _questions_GetQuestionByIdRequest, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_questions_GetQuestionByIdResponse__Output>): grpc.ClientUnaryCall;
  GetQuestionById(argument: _questions_GetQuestionByIdRequest, metadata: grpc.Metadata, callback: grpc.requestCallback<_questions_GetQuestionByIdResponse__Output>): grpc.ClientUnaryCall;
  GetQuestionById(argument: _questions_GetQuestionByIdRequest, options: grpc.CallOptions, callback: grpc.requestCallback<_questions_GetQuestionByIdResponse__Output>): grpc.ClientUnaryCall;
  GetQuestionById(argument: _questions_GetQuestionByIdRequest, callback: grpc.requestCallback<_questions_GetQuestionByIdResponse__Output>): grpc.ClientUnaryCall;
  getQuestionById(argument: _questions_GetQuestionByIdRequest, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_questions_GetQuestionByIdResponse__Output>): grpc.ClientUnaryCall;
  getQuestionById(argument: _questions_GetQuestionByIdRequest, metadata: grpc.Metadata, callback: grpc.requestCallback<_questions_GetQuestionByIdResponse__Output>): grpc.ClientUnaryCall;
  getQuestionById(argument: _questions_GetQuestionByIdRequest, options: grpc.CallOptions, callback: grpc.requestCallback<_questions_GetQuestionByIdResponse__Output>): grpc.ClientUnaryCall;
  getQuestionById(argument: _questions_GetQuestionByIdRequest, callback: grpc.requestCallback<_questions_GetQuestionByIdResponse__Output>): grpc.ClientUnaryCall;
  
  GetQuestions(argument: _questions_GetQuestionsRequest, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_questions_GetQuestionsResponse__Output>): grpc.ClientUnaryCall;
  GetQuestions(argument: _questions_GetQuestionsRequest, metadata: grpc.Metadata, callback: grpc.requestCallback<_questions_GetQuestionsResponse__Output>): grpc.ClientUnaryCall;
  GetQuestions(argument: _questions_GetQuestionsRequest, options: grpc.CallOptions, callback: grpc.requestCallback<_questions_GetQuestionsResponse__Output>): grpc.ClientUnaryCall;
  GetQuestions(argument: _questions_GetQuestionsRequest, callback: grpc.requestCallback<_questions_GetQuestionsResponse__Output>): grpc.ClientUnaryCall;
  getQuestions(argument: _questions_GetQuestionsRequest, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_questions_GetQuestionsResponse__Output>): grpc.ClientUnaryCall;
  getQuestions(argument: _questions_GetQuestionsRequest, metadata: grpc.Metadata, callback: grpc.requestCallback<_questions_GetQuestionsResponse__Output>): grpc.ClientUnaryCall;
  getQuestions(argument: _questions_GetQuestionsRequest, options: grpc.CallOptions, callback: grpc.requestCallback<_questions_GetQuestionsResponse__Output>): grpc.ClientUnaryCall;
  getQuestions(argument: _questions_GetQuestionsRequest, callback: grpc.requestCallback<_questions_GetQuestionsResponse__Output>): grpc.ClientUnaryCall;
  
  SearchQuestion(argument: _questions_SearchQuestionRequest, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_questions_GetQuestionsResponse__Output>): grpc.ClientUnaryCall;
  SearchQuestion(argument: _questions_SearchQuestionRequest, metadata: grpc.Metadata, callback: grpc.requestCallback<_questions_GetQuestionsResponse__Output>): grpc.ClientUnaryCall;
  SearchQuestion(argument: _questions_SearchQuestionRequest, options: grpc.CallOptions, callback: grpc.requestCallback<_questions_GetQuestionsResponse__Output>): grpc.ClientUnaryCall;
  SearchQuestion(argument: _questions_SearchQuestionRequest, callback: grpc.requestCallback<_questions_GetQuestionsResponse__Output>): grpc.ClientUnaryCall;
  searchQuestion(argument: _questions_SearchQuestionRequest, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_questions_GetQuestionsResponse__Output>): grpc.ClientUnaryCall;
  searchQuestion(argument: _questions_SearchQuestionRequest, metadata: grpc.Metadata, callback: grpc.requestCallback<_questions_GetQuestionsResponse__Output>): grpc.ClientUnaryCall;
  searchQuestion(argument: _questions_SearchQuestionRequest, options: grpc.CallOptions, callback: grpc.requestCallback<_questions_GetQuestionsResponse__Output>): grpc.ClientUnaryCall;
  searchQuestion(argument: _questions_SearchQuestionRequest, callback: grpc.requestCallback<_questions_GetQuestionsResponse__Output>): grpc.ClientUnaryCall;
  
}

export interface QuestionServiceHandlers extends grpc.UntypedServiceImplementation {
  GetQuestionById: grpc.handleUnaryCall<_questions_GetQuestionByIdRequest__Output, _questions_GetQuestionByIdResponse>;
  
  GetQuestions: grpc.handleUnaryCall<_questions_GetQuestionsRequest__Output, _questions_GetQuestionsResponse>;
  
  SearchQuestion: grpc.handleUnaryCall<_questions_SearchQuestionRequest__Output, _questions_GetQuestionsResponse>;
  
}

export interface QuestionServiceDefinition extends grpc.ServiceDefinition {
  GetQuestionById: MethodDefinition<_questions_GetQuestionByIdRequest, _questions_GetQuestionByIdResponse, _questions_GetQuestionByIdRequest__Output, _questions_GetQuestionByIdResponse__Output>
  GetQuestions: MethodDefinition<_questions_GetQuestionsRequest, _questions_GetQuestionsResponse, _questions_GetQuestionsRequest__Output, _questions_GetQuestionsResponse__Output>
  SearchQuestion: MethodDefinition<_questions_SearchQuestionRequest, _questions_GetQuestionsResponse, _questions_SearchQuestionRequest__Output, _questions_GetQuestionsResponse__Output>
}
