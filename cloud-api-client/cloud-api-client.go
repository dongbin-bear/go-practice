package main

import (
	"context"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	_ "github.com/dgrijalva/jwt-go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"log"
	"time"

	cloud_api_service_pb "github.com/cloud/cloud_api_service/v0/github.com/bearrobotics/pennybot/api/cloud/v0/cloud_api_service_go_proto"
	//cloud_api_service_pb "github.com/bearrobotics/pennybot/api/cloud/v0/cloud_api_service_go_proto"
)

func main() {
	// Set the address
	serverAddress := "cloud-api-server-6svom534oq-uw.a.run.app"
	serverPort := 5123

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Create a gRPC connection with WithInsecure for development only
	conn, err := grpc.DialContext(ctx, fmt.Sprintf("%s:%d", serverAddress, serverPort), grpc.WithInsecure())
	if err != nil {
		// Check for timeout error specifically
		if status.Code(err) == codes.DeadlineExceeded {
			log.Fatalf("Connection timed out: %v", err)
		} else {
			log.Fatalf("Failed to connect to gRPC server: %v", err)
		}
	}
	defer conn.Close()

	log.Printf("Success")
	// Create a gRPC client
	client := cloud_api_service_pb.NewCloudAPIServiceClient(conn)

	// Create a ListRobotIDsRequest with any necessary parameters
	request := &cloud_api_service_pb.ListRobotIDsRequest{
		Filter: &cloud_api_service_pb.RobotFilter{LocationId: "cus1_location1"},
		// Add your request parameters here if needed
	}

	// Create a JWT token
	token, err := createJWT()
	if err != nil {
		log.Fatalf("Failed to create JWT token: %v", err)
	}

	ctx = metadata.AppendToOutgoingContext(ctx, "authorization", "Bearer "+token)

	log.Printf("Before request")

	// Make the gRPC call to ListRobotIDs with the JWT included
	response, err := client.ListRobotIDs(ctx, request)
	if err != nil {
		log.Fatalf("Failed to call ListRobotIDs RPC: %v", err)
	}

	// Process the response
	for _, robotID := range response.RobotIds {
		fmt.Printf("Robot ID: %s\n", robotID)
	}
}

func createJWT() (string, error) {
	// Set the expiration time to one hour from now
	expirationTime := time.Now().Add(time.Hour)

	// Create the token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": "user-id-here",
		"iss":    "Bear",
		"sub":    "auth/subject-here",
		"exp":    expirationTime.Unix(),
	})

	// Sign the token with your secret key
	secretKey := []byte("")
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
