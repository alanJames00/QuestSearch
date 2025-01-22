// Original file: src/api/grpc/proto/question.proto


export interface SearchQuestionRequest {
  'searchTerm'?: (string);
  'page'?: (number);
  'limit'?: (number);
}

export interface SearchQuestionRequest__Output {
  'searchTerm': (string);
  'page': (number);
  'limit': (number);
}
