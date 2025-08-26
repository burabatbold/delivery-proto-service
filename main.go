package main

import (
	"log"

	adminProto "github.com/burabatbold/delivery-proto-service/generate/admin-service/grpc/protos"
	customerProto "github.com/burabatbold/delivery-proto-service/generate/customer-service/grpc/protos"
	deliveryProto "github.com/burabatbold/delivery-proto-service/generate/delivery-service/grpc/proto"
	notificationProto "github.com/burabatbold/delivery-proto-service/generate/notification-service/grpc/proto"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ProtoClient interface {
	NotificationClient() notificationProto.NotificationServiceClient
	DeliveryClient() deliveryProto.DeliveryServiceClient
	CustomerClient() customerProto.CustomerServiceClient
	AdminClient() adminProto.AdminServiceClient
}

type notificationClient struct {
	config *viper.Viper
}

func NewProtoClient(config *viper.Viper) ProtoClient {
	return &notificationClient{
		config: config,
	}
}

func (c *notificationClient) NotificationClient() notificationProto.NotificationServiceClient {
	conn, err := grpc.NewClient(c.config.GetString("notification.address"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to notification service: %v", err)
	}

	return notificationProto.NewNotificationServiceClient(conn)
}

func (c *notificationClient) DeliveryClient() deliveryProto.DeliveryServiceClient {
	conn, err := grpc.NewClient(c.config.GetString("delivery.address"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to delivery service: %v", err)
	}

	return deliveryProto.NewDeliveryServiceClient(conn)
}

func (c *notificationClient) CustomerClient() customerProto.CustomerServiceClient {
	conn, err := grpc.NewClient(c.config.GetString("customer.address"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to customer service: %v", err)
	}

	return customerProto.NewCustomerServiceClient(conn)
}

func (c *notificationClient) AdminClient() adminProto.AdminServiceClient {
	conn, err := grpc.NewClient(c.config.GetString("admin.address"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to admin service: %v", err)
	}

	return adminProto.NewAdminServiceClient(conn)
}
