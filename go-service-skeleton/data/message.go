package data

import (
	"errors"
	"sre.qlik.com/palindrome/dbconnection"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// ErrMessageNotFound is the error that is returned when there is not matching message
var ErrMessageNotFound = errors.New("no message found with the given ID")

// Message is the structure for a message
type Message struct {
	ID     int    `json:"id"`
	Text   string `json:"text"`
	Sender string `json:"sender"`
	Time   string `json:"time"`
	IsPalindrome bool `json:"ispalindrome"`
}

// Messages is the databaseCollection of all messages
type Messages []*Message

var messages Messages

// GetMessages returns all the messages
func GetMessages() Messages {
	return messages
}

// GetMessagesFromDB returns all the messages from DB
func GetMessagesFromDB() ([]Message, error) {
	// Define filterQuery query for fetching particular record/document from databaseCollection
	// bson.D{{}} gets 'all documents'. Bson is a format used to store documents in MongoDB
	filterQuery := bson.D{{}} 
	messages := []Message{}

	// Create connection
	client, err := connectionhelper.GetMongoClient()
	if err != nil {
		return messages, err
	}

	// Handle for the databaseCollection
	databaseCollection := client.Database(connectionhelper.DB).Collection(connectionhelper.MESSAGE)
	
	// Find operation
	cur, findError := databaseCollection.Find(context.TODO(), filterQuery)
	if findError != nil {
		return messages, findError
	}

	// Prepare messages
	for cur.Next(context.TODO()) {
		t := Message{}
		err := cur.Decode(&t)
		if err != nil {
			return messages, err
		}
		messages = append(messages, t)
	}
	// once exhausted, close the cursor
	cur.Close(context.TODO())
	if len(messages) == 0 {
		return messages, mongo.ErrNoDocuments
	}
	return messages, nil
}

// AddMessage adds the message to the list of messages
func AddMessage(msg *Message) {
	var lastMessageID int
	// there is at least one message
	if len(messages) > 0 {
		lastMessageID = messages[len(messages)-1].ID
	}
	msg.ID = lastMessageID + 1
	messages = append(messages, msg)
}

// AddMessageToDB adds the message to the list of messages in DB
func AddMessageToDB(msg *Message) error {
	// Create connection
	client, err := connectionhelper.GetMongoClient()
	if err != nil {
		return err
	}

	// Handle to the database.
	databaseCollection := client.Database(connectionhelper.DB).Collection(connectionhelper.MESSAGE)

	// Insert operation
	_, err = databaseCollection.InsertOne(context.TODO(), msg)
	if err != nil {
		return err
	}

	return nil
}

// GetMessageByID returns a message for a given messageID
func GetMessageByID(messageID int) (*Message, error) {
	for i := range messages {
		if messages[i].ID == messageID {
			return messages[i], nil
		}
	}
	return nil, ErrMessageNotFound
}

// GetMessageFromDBByID returns a message for a given messageID
func GetMessageFromDBByID(messageID int) (Message, error) {
	resultToReturn := Message{}

	// Define filterQuery query for fetching particular record/document from databaseCollection
	filterQuery := bson.D{primitive.E{Key: "id", Value: messageID}}

	// Create connection
	client, err := connectionhelper.GetMongoClient()
	if err != nil {
		return resultToReturn, err
	}
	 
	// Handle to the database
	databaseCollection := client.Database(connectionhelper.DB).Collection(connectionhelper.MESSAGE)

	// FindOne operation
	err = databaseCollection.FindOne(context.TODO(), filterQuery).Decode(&resultToReturn)
	if err != nil {
		return resultToReturn, err
	}
	return resultToReturn, nil
}

// DeleteMessageWithID deletes a message with the given ID
func DeleteMessageWithID(messageID int) error {
	var indexToDelete = -1
	for i := range messages {
		if messageID == messages[i].ID {
			indexToDelete = i
			break
		}
	}
	// no message with the given ID was found
	if indexToDelete == -1 {
		return ErrMessageNotFound
	}
	messages = append(messages[:indexToDelete], messages[indexToDelete+1:]...)
	return nil
}

// DeleteMessageFromDBWithID deletes a message with the given ID
func DeleteMessageFromDBWithID(messageID int) error {
	// Define filterQuery query for fetching particular record/document from databaseCollection
	filterQuery := bson.D{primitive.E{Key: "id", Value: messageID}}

	// Create connection
	client, err := connectionhelper.GetMongoClient()
	if err != nil {
		return err
	}

	// Handle to the database.
	databaseCollection := client.Database(connectionhelper.DB).Collection(connectionhelper.MESSAGE)

	// DeleteOne operation
	_, err = databaseCollection.DeleteOne(context.TODO(), filterQuery)
	if err != nil {
		return err
	}
	return nil
}
