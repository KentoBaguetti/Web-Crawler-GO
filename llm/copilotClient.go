package llm

import (
	"context"
	"errors"
	"fmt"
	"sync"

	copilot "github.com/github/copilot-sdk/go"
)

// temp function for testing
// send data to the llm and recieve a parsed response

type ClientSession struct {
	Client *copilot.Client
}

type AgentSession struct {
	Session *copilot.Session
}

var options = copilot.ClientOptions{
	LogLevel: "debug",
}

func CreateClient() (*ClientSession, error) {

	client := copilot.NewClient(&options)
	if err := client.Start(context.Background()); err != nil {
		return nil, fmt.Errorf("start copilot client: %w", err)
	}

	return &ClientSession{client}, nil
}

func StopClient(client *ClientSession) error {
	if client == nil || client.Client == nil {
		return nil
	}

	if err := client.Client.Stop(); err != nil {
		return fmt.Errorf("stop copilot client: %w", err)
	}

	return nil
}

var copilotSessionConfigs = copilot.SessionConfig{
	Model:               "gpt-5",
	OnPermissionRequest: copilot.PermissionHandler.ApproveAll,
}

func CreateSessions(numberOfSessions int, client *ClientSession) ([]*AgentSession, error) {

	if client == nil {
		return nil, errors.New("client is nil")
	}

	clientInstance := client.Client

	sessions := []*AgentSession{}

	for range numberOfSessions {
		tempSession, err := clientInstance.CreateSession(context.Background(), &copilotSessionConfigs)
		if err != nil {
			return nil, fmt.Errorf("create session: %w", err)
		}

		tempAgentSession := AgentSession{tempSession}

		sessions = append(sessions, &tempAgentSession)
	}

	return sessions, nil
}

func SendMessageToAgent(msg string, agentSession *AgentSession) error {
	fmt.Println("Start parsing...")
	defer fmt.Println("End parsing.")

	if agentSession == nil || agentSession.Session == nil {
		return errors.New("agent session is nil")
	}

	session := agentSession.Session

	done := make(chan bool)
	var once sync.Once
	unsubscribe := agentSession.Session.On(func(event copilot.SessionEvent) {
		if event.Type == "assistant.message" {
			if event.Data.Content != nil {
				fmt.Println(*event.Data.Content)
			}
		}
		if event.Type == "session.idle" {
			once.Do(func() {
				close(done)
			})
		}
	})
	defer unsubscribe()

	_, err := session.Send(context.Background(), copilot.MessageOptions{
		Prompt: msg,
	})
	if err != nil {
		return fmt.Errorf("send prompt: %w", err)
	}

	<-done
	return nil
}

// this is a test function that runs everything i want but only this needs to be called in main
func TestFunction() error {

	clientSession, err := CreateClient()
	if err != nil {
		return err
	}
	defer func() {
		if stopErr := StopClient(clientSession); stopErr != nil {
			fmt.Printf("copilot shutdown error: %v\n", stopErr)
		}
	}()

	agentSessions, err := CreateSessions(3, clientSession)
	if err != nil {
		return err
	}

	for i := range 3 {
		prompt := fmt.Sprintf("Return the following number multiplied by 10: %d", i)
		if err := SendMessageToAgent(prompt, agentSessions[i]); err != nil {
			return fmt.Errorf("session %d: %w", i, err)
		}
	}

	return nil
}
