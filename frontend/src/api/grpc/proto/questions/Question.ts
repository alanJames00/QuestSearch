// Original file: src/api/grpc/proto/question.proto

import type { ContentOnlyQuestion as _questions_ContentOnlyQuestion, ContentOnlyQuestion__Output as _questions_ContentOnlyQuestion__Output } from '../questions/ContentOnlyQuestion';
import type { ConversationQuestion as _questions_ConversationQuestion, ConversationQuestion__Output as _questions_ConversationQuestion__Output } from '../questions/ConversationQuestion';
import type { AnagramQuestion as _questions_AnagramQuestion, AnagramQuestion__Output as _questions_AnagramQuestion__Output } from '../questions/AnagramQuestion';
import type { MCQQuestion as _questions_MCQQuestion, MCQQuestion__Output as _questions_MCQQuestion__Output } from '../questions/MCQQuestion';
import type { ReadAlongQuestion as _questions_ReadAlongQuestion, ReadAlongQuestion__Output as _questions_ReadAlongQuestion__Output } from '../questions/ReadAlongQuestion';

export interface Question {
  'id'?: (string);
  'type'?: (string);
  'title'?: (string);
  'siblingId'?: (string);
  'contentOnly'?: (_questions_ContentOnlyQuestion | null);
  'conversation'?: (_questions_ConversationQuestion | null);
  'anagram'?: (_questions_AnagramQuestion | null);
  'mcq'?: (_questions_MCQQuestion | null);
  'readAlong'?: (_questions_ReadAlongQuestion | null);
  'questionDetails'?: "contentOnly"|"conversation"|"anagram"|"mcq"|"readAlong";
}

export interface Question__Output {
  'id': (string);
  'type': (string);
  'title': (string);
  'siblingId': (string);
  'contentOnly'?: (_questions_ContentOnlyQuestion__Output | null);
  'conversation'?: (_questions_ConversationQuestion__Output | null);
  'anagram'?: (_questions_AnagramQuestion__Output | null);
  'mcq'?: (_questions_MCQQuestion__Output | null);
  'readAlong'?: (_questions_ReadAlongQuestion__Output | null);
  'questionDetails': "contentOnly"|"conversation"|"anagram"|"mcq"|"readAlong";
}
