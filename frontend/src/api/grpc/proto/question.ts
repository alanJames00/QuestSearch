import type * as grpc from '@grpc/grpc-js';
import type { MessageTypeDefinition } from '@grpc/proto-loader';

import type { QuestionServiceClient as _questions_QuestionServiceClient, QuestionServiceDefinition as _questions_QuestionServiceDefinition } from './questions/QuestionService';

type SubtypeConstructor<Constructor extends new (...args: any) => any, Subtype> = {
  new(...args: ConstructorParameters<Constructor>): Subtype;
};

export interface ProtoGrpcType {
  questions: {
    AnagramQuestion: MessageTypeDefinition
    Block: MessageTypeDefinition
    ContentOnlyQuestion: MessageTypeDefinition
    ConversationQuestion: MessageTypeDefinition
    GetQuestionByIdRequest: MessageTypeDefinition
    GetQuestionByIdResponse: MessageTypeDefinition
    GetQuestionsRequest: MessageTypeDefinition
    GetQuestionsResponse: MessageTypeDefinition
    MCQQuestion: MessageTypeDefinition
    Option: MessageTypeDefinition
    Question: MessageTypeDefinition
    QuestionService: SubtypeConstructor<typeof grpc.Client, _questions_QuestionServiceClient> & { service: _questions_QuestionServiceDefinition }
    ReadAlongQuestion: MessageTypeDefinition
    SearchQuestionRequest: MessageTypeDefinition
  }
}

