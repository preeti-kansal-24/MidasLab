package kafka

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/IBM/sarama"
	log "github.com/sirupsen/logrus"
	"preeti-kansal-24/MidasLab.git/clients"
	"preeti-kansal-24/MidasLab.git/domain/repository"
	"preeti-kansal-24/MidasLab.git/schema"
	"time"
)

type Consumer struct {
	otpStore repository.OTPStore
}

func (c *Consumer) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		phoneNumber := string(msg.Key)
		//generate otp
		userPrfile := &schema.UserProfile{}
		err := json.Unmarshal(msg.Value, userPrfile)
		if err != nil {
			log.Printf("failed to unmarshal user profile: %v", err)
			continue
		}
		fmt.Printf("Generating otp for user %v\n", userPrfile.Name)
		m, otp, _ := clients.GenerateOTP(phoneNumber)

		log.Printf("mesasage is %v\n", m)
		//Save otp to db
		c.otpStore.Save(&schema.Otps{UserProfileId: userPrfile.Id, Otp: otp})
		session.MarkMessage(msg, "")
	}
	return nil
}

func (c *Consumer) Setup(sarama.ConsumerGroupSession) error {
	return nil
}

func (c *Consumer) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

func initializeConsumerGroup() (sarama.ConsumerGroup, error) {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true

	sarama.DebugLogger = log.New()

	consumerGroup, err := sarama.NewConsumerGroup(
		[]string{broker}, group, config)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize consumer group: %w", err)
	}

	fmt.Println("Successfully initialized consumer group")
	return consumerGroup, nil
}

func NewConsumer(ctx context.Context, topic string, otpStore repository.OTPStore) (*Consumer, error) {
	consumerGroup, err := initializeConsumerGroup()
	if err != nil {
		log.Printf("initialization error: %v", err)
		return nil, err
	}

	consumer := &Consumer{otpStore: otpStore}
	go func() {
		defer consumerGroup.Close()
		for {
			if err := consumerGroup.Consume(ctx, []string{topic}, consumer); err != nil {
				log.Printf("error from consumer: %v", err)
				if ctx.Err() != nil {
					return
				}
			}
			if ctx.Err() != nil {
				return
			}
			// Added a sleep to avoid tight loop
			time.Sleep(time.Second)
		}
	}()

	return consumer, nil
}
