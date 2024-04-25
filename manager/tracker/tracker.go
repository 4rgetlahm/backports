package tracker

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"time"

	"cloud.google.com/go/pubsub"
	"github.com/4rgetlahm/backports/api/service"
	"github.com/4rgetlahm/backports/types"
	"github.com/emicklei/go-restful/v3/log"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/api/option"
)

type runnerStatusMessage struct {
	Reference string                 `json:"reference"`
	Stage     string                 `json:"stage"`
	Status    string                 `json:"status"`
	Payload   map[string]interface{} `json:"payload"`
}

func StartStatusTracker(pubsubProject string, pubsubTopic string, base64pubsubCredentials string) {
	pubsubCredentials, err := base64.StdEncoding.DecodeString(base64pubsubCredentials)

	if err != nil {
		log.Printf("Error decoding pubsub credentials: %v", err)
		return
	}

	client, err := pubsub.NewClient(context.Background(), pubsubProject, option.WithCredentialsJSON([]byte(pubsubCredentials)))

	if err != nil {
		log.Printf("Error creating pubsub client: %v", err)
		return
	}

	sub := client.Subscription(pubsubTopic + "-sub")
	err = sub.Receive(context.Background(), func(ctx context.Context, msg *pubsub.Message) {
		err = processMessage(msg)
		if err != nil {
			log.Printf("Error processing message: %s, err: %v", msg.Data, err)
		}
		msg.Ack()
	})

	if err != nil {
		log.Printf("Error receiving messages: %s", err.Error())
		return
	}

	log.Printf("Status tracker exited")
}

func processMessage(msg *pubsub.Message) error {

	var statusMessage runnerStatusMessage
	err := json.Unmarshal(msg.Data, &statusMessage)

	if err != nil {
		return err
	}

	log.Printf("Received valid status message: %s", msg.Data)

	backportEvent, err := statusMessage.toBackportEvent()

	if err != nil {
		return err
	}

	objectId, err := primitive.ObjectIDFromHex(statusMessage.Reference)

	if err != nil {
		return err
	}

	service.AddBackportEvent(objectId, backportEvent)

	return nil
}

func (statusMessage *runnerStatusMessage) toBackportEvent() (*types.BackportEvent, error) {
	var eventType string
	if statusMessage.Status == "success" {
		switch statusMessage.Stage {
		case "fetch":
			eventType = types.ActionGitFetchSuccess
		case "checkout":
			eventType = types.ActionGitCheckoutSuccess
		case "pull":
			eventType = types.ActionGitPullSuccess
		case "checkout_new_branch":
			eventType = types.ActionGitCheckoutNewBranchSuccess
		case "cherry_pick":
			eventType = types.ActionGitCherryPickSuccess
		case "push":
			eventType = types.ActionGitPushSuccess
		case "exit":
			eventType = types.ActionRunnerExited
		default:
			log.Printf("Unknown stage: %v", statusMessage.Stage)
			return nil, errors.New("unknown stage")
		}
	} else if statusMessage.Status == "failure" {
		switch statusMessage.Stage {
		case "fetch":
			eventType = types.ActionGitFetchFailure
		case "checkout":
			eventType = types.ActionGitCheckoutFailure
		case "pull":
			eventType = types.ActionGitPullFailure
		case "checkout_new_branch":
			eventType = types.ActionGitCheckoutNewBranchFailure
		case "cherry_pick":
			eventType = types.ActionGitCherryPickFailure
		case "push":
			eventType = types.ActionGitPushFailure
		default:
			log.Printf("Unknown stage: %v", statusMessage.Stage)
			return nil, errors.New("unknown stage")
		}
	} else if statusMessage.Status == "start" {
		switch statusMessage.Stage {
		case "fetch":
			eventType = types.ActionGitFetchStart
		case "checkout":
			eventType = types.ActionGitCheckoutStart
		case "pull":
			eventType = types.ActionGitPullStart
		case "checkout_new_branch":
			eventType = types.ActionGitCheckoutNewBranchStart
		case "cherry_pick":
			eventType = types.ActionGitCherryPickStart
		case "push":
			eventType = types.ActionGitPushStart
		default:
			log.Printf("Unknown stage: %v", statusMessage.Stage)
			return nil, errors.New("unknown stage")
		}
	} else {
		log.Printf("Unknown status: %v", statusMessage.Status)
		return nil, errors.New("unknown status")
	}

	return &types.BackportEvent{
		Action:      eventType,
		Content:     statusMessage.Payload,
		DateCreated: time.Now().UTC(),
	}, nil
}
