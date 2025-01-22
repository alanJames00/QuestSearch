// Original file: src/api/grpc/proto/question.proto

import type { Block as _questions_Block, Block__Output as _questions_Block__Output } from '../questions/Block';

export interface AnagramQuestion {
  'anagramType'?: (string);
  'blocks'?: (_questions_Block)[];
  'solution'?: (string);
}

export interface AnagramQuestion__Output {
  'anagramType': (string);
  'blocks': (_questions_Block__Output)[];
  'solution': (string);
}
