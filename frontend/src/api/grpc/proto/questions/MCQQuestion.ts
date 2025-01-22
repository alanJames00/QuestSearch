// Original file: src/api/grpc/proto/question.proto

import type { Option as _questions_Option, Option__Output as _questions_Option__Output } from '../questions/Option';

export interface MCQQuestion {
  'options'?: (_questions_Option)[];
}

export interface MCQQuestion__Output {
  'options': (_questions_Option__Output)[];
}
