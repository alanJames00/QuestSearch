// Original file: src/api/grpc/proto/question.proto

import type { Question as _questions_Question, Question__Output as _questions_Question__Output } from '../questions/Question';

export interface GetQuestionsResponse {
  'questions'?: (_questions_Question)[];
  'totalQuestions'?: (number);
  'totalPages'?: (number);
}

export interface GetQuestionsResponse__Output {
  'questions': (_questions_Question__Output)[];
  'totalQuestions': (number);
  'totalPages': (number);
}
