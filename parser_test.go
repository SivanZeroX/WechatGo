package wechatgo

import (
	"testing"
)

func TestParseMessage_EmptyData(t *testing.T) {
	_, err := ParseMessage([]byte{})
	if err == nil {
		t.Fatalf("Expected error for empty data, got nil")
	}
	if err.Error() != "empty message data" {
		t.Fatalf("Expected error message 'empty message data', got '%s'", err.Error())
	}
}

func TestParseMessage_NilData(t *testing.T) {
	_, err := ParseMessage(nil)
	if err == nil {
		t.Fatalf("Expected error for nil data, got nil")
	}
	if err.Error() != "empty message data" {
		t.Fatalf("Expected error message 'empty message data', got '%s'", err.Error())
	}
}

func TestParseMessage_UnknownType(t *testing.T) {
	xmlData := []byte(`
		<xml>
			<ToUserName>toUser</ToUserName>
			<FromUserName>fromUser</FromUserName>
			<CreateTime>1234567890</CreateTime>
			<MsgType>unknown</MsgType>
			<MsgId>1234567890</MsgId>
		</xml>
	`)

	result, err := ParseMessage(xmlData)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	msg, ok := result.(*BaseMessage)
	if !ok {
		t.Fatalf("Expected BaseMessage, got %T", result)
	}

	if msg.ToUserName != "toUser" {
		t.Fatalf("Expected ToUserName 'toUser', got '%s'", msg.ToUserName)
	}
	if msg.FromUserName != "fromUser" {
		t.Fatalf("Expected FromUserName 'fromUser', got '%s'", msg.FromUserName)
	}
	if msg.MsgType != "unknown" {
		t.Fatalf("Expected MsgType 'unknown', got '%s'", msg.MsgType)
	}
}

func TestParseMessage_TextMessage(t *testing.T) {
	xmlData := []byte(`
		<xml>
			<ToUserName>toUser</ToUserName>
			<FromUserName>fromUser</FromUserName>
			<CreateTime>1234567890</CreateTime>
			<MsgType>text</MsgType>
			<MsgId>1234567890</MsgId>
			<Content>Hello World</Content>
		</xml>
	`)

	result, err := ParseMessage(xmlData)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	msg, ok := result.(*TextMessage)
	if !ok {
		t.Fatalf("Expected TextMessage, got %T", result)
	}

	if msg.ToUserName != "toUser" {
		t.Fatalf("Expected ToUserName 'toUser', got '%s'", msg.ToUserName)
	}
	if msg.FromUserName != "fromUser" {
		t.Fatalf("Expected FromUserName 'fromUser', got '%s'", msg.FromUserName)
	}
	if msg.MsgType != "text" {
		t.Fatalf("Expected MsgType 'text', got '%s'", msg.MsgType)
	}
}

func TestParseMessage_SubscribeEvent(t *testing.T) {
	xmlData := []byte(`
		<xml>
			<ToUserName>toUser</ToUserName>
			<FromUserName>fromUser</FromUserName>
			<CreateTime>1234567890</CreateTime>
			<MsgType>event</MsgType>
			<Event>subscribe</Event>
			<EventKey>EVENTKEY</EventKey>
			<Ticket>TICKET</Ticket>
		</xml>
	`)

	result, err := ParseMessage(xmlData)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	event, ok := result.(*SubscribeEvent)
	if !ok {
		t.Fatalf("Expected SubscribeEvent, got %T", result)
	}

	if event.ToUserName != "toUser" {
		t.Fatalf("Expected ToUserName 'toUser', got '%s'", event.ToUserName)
	}
	if event.FromUserName != "fromUser" {
		t.Fatalf("Expected FromUserName 'fromUser', got '%s'", event.FromUserName)
	}
	if event.Event != "subscribe" {
		t.Fatalf("Expected Event 'subscribe', got '%s'", event.Event)
	}
	if event.EventKey != "EVENTKEY" {
		t.Fatalf("Expected EventKey 'EVENTKEY', got '%s'", event.EventKey)
	}
	if event.Ticket != "TICKET" {
		t.Fatalf("Expected Ticket 'TICKET', got '%s'", event.Ticket)
	}
}

func TestParseMessage_ClickEvent(t *testing.T) {
	xmlData := []byte(`
		<xml>
			<ToUserName>toUser</ToUserName>
			<FromUserName>fromUser</FromUserName>
			<CreateTime>1234567890</CreateTime>
			<MsgType>event</MsgType>
			<Event>CLICK</Event>
			<EventKey>MENU_KEY</EventKey>
		</xml>
	`)

	result, err := ParseMessage(xmlData)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	event, ok := result.(*ClickEvent)
	if !ok {
		t.Fatalf("Expected ClickEvent, got %T", result)
	}

	if event.EventKey != "MENU_KEY" {
		t.Fatalf("Expected EventKey 'MENU_KEY', got '%s'", event.EventKey)
	}
}
