package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// BaseQuestion struct: holds the common fields in the collection
type BaseQuestion struct {
	ID        primitive.ObjectID  `bson:"_id" json:"_id"`
	Type      string              `bson:"type" json:"type"`
	Title     string              `bson:"title" json:"title"`
	SiblingID *primitive.ObjectID `bson:"siblingId" json:"siblingId"`
}

// ContentOnly Questions
type ContentOnly struct {
	ID        primitive.ObjectID  `bson:"_id" json:"_id"`
	Type      string              `bson:"type" json:"type"`
	Title     string              `bson:"title" json:"title"`
	SiblingID *primitive.ObjectID `bson:"siblingId" json:"siblingId"`
}

// Conversation Questions
type Conversation struct {
	ID        primitive.ObjectID  `bson:"_id" json:"_id"`
	Type      string              `bson:"type" json:"type"`
	Title     string              `bson:"title" json:"title"`
	SiblingID *primitive.ObjectID `bson:"siblingId" json:"siblingId"`
}

// Anagram Questions
type Anagram struct {
	ID        primitive.ObjectID  `bson:"_id" json:"_id"`
	Type      string              `bson:"type" json:"type"`
	Title     string              `bson:"title" json:"title"`
	SiblingID *primitive.ObjectID `bson:"siblingId" json:"siblingId"`

	AnagramType string  `bson:"anagramType" json:"anagramType"`
	Blocks      []Block `bson:"blocks" json:"blocks"`
	Solution    string  `bson:"solution" json:"solution"`
}

// MCQ Questions
type MCQ struct {
	ID        primitive.ObjectID  `bson:"_id" json:"_id"`
	Type      string              `bson:"type" json:"type"`
	Title     string              `bson:"title" json:"title"`
	SiblingID *primitive.ObjectID `bson:"siblingId" json:"siblingId"`

	Options []Option `bson:"options" json:"options"`
}

// ReadAlong Questions
type ReadAlong struct {
	ID        primitive.ObjectID  `bson:"_id" json:"_id"`
	Type      string              `bson:"type" json:"type"`
	Title     string              `bson:"title" json:"title"`
	SiblingID *primitive.ObjectID `bson:"siblingId" json:"siblingId"`
}

// substructs definitions

// Block struct for AnagramType Questions
type Block struct {
	Text         string `bson:"text" json:"text"`
	ShowInOption bool   `bson:"showInOption" json:"showInOption"`
	IsAnswer     bool   `bson:"isAnswer" json:"isAnswer"`
}

// Option struct for MCQ Questions
type Option struct {
	Text            string `bson:"text" json:"text"`
	IsCorrectAnswer bool   `bson:"isCorrectAnswer" json:"isCorrectAnswer"`
}
