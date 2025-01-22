// Original file: src/api/grpc/proto/question.proto

import type { Question as _questions_Question, Question__Output as _questions_Question__Output } from '../questions/Question';

export interface GetQuestionByIdResponse {
  'question'?: (_questions_Question | null);
}

export interface GetQuestionByIdResponse__Output {
  'question': (_questions_Question__Output | null);
}
