package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"sync"

	pb "github.com/pranavnaikp/train_ticketing/proto/protos"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)


type server struct {
	pb.UnimplementedTicketBookingServiceServer
	mu       sync.Mutex
	users    map[string]*pb.UserDetail // Map email to UserDetail
	seats    map[int32]string          // Map seat number to email
	receipts map[string]*pb.TicketReceipt
}

type TicketReceipt struct {
	ReceiptID  string
	From       string
	To         string
	Price      float32
	SeatNumber int32
	UserName   string
	UserEmail  string
}

func (s *server) BookTicket(ctx context.Context, in *pb.BookTicketRequest) (*pb.BookTicketResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	// Check if seat is already booked
	if _, exists := s.seats[in.SeatNumber]; exists {
		return nil, fmt.Errorf("seat %d is already booked", in.SeatNumber)
	}

	// Create and store the user detail
	userDetail := &pb.UserDetail{
		FirstName:  in.FirstName,
		LastName:   in.LastName,
		Email:      in.Email,
		SeatNumber: in.SeatNumber,
	}

	s.users[in.Email] = userDetail
	s.seats[in.SeatNumber] = in.Email

	
	receiptID := fmt.Sprintf("REC-%s-%d", in.Email, in.SeatNumber)

	return &pb.BookTicketResponse{
		ReceiptId: receiptID,
		Message:   "Ticket booked successfully",
	}, nil
}

func (s *server) GetUserDetail(ctx context.Context, in *pb.UserRequest) (*pb.UserResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	userDetail, exists := s.users[in.Email]
	if !exists {
		return nil, fmt.Errorf("no user found with email %s", in.Email)
	}

	return &pb.UserResponse{
		User: userDetail,
	}, nil
}

func (s *server) GetUsersBySection(ctx context.Context, in *pb.SectionRequest) (*pb.SectionResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	usersInSection := []*pb.UserDetail{}
	for _, user := range s.users {
		if (in.Section == "A" && user.SeatNumber <= 50) || (in.Section == "B" && user.SeatNumber > 50) {
			usersInSection = append(usersInSection, user)
		}
	}

	return &pb.SectionResponse{
		Users: usersInSection,
	}, nil
}

func (s *server) RemoveUser(ctx context.Context, in *pb.RemoveUserRequest) (*pb.RemoveUserResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	userDetail, exists := s.users[in.Email]
	if !exists {
		return nil, fmt.Errorf("no user found with email %s", in.Email)
	}

	delete(s.seats, userDetail.SeatNumber)
	delete(s.users, in.Email)

	return &pb.RemoveUserResponse{
		Message: "User removed successfully",
	}, nil
}

func (s *server) ModifyUserSeat(ctx context.Context, in *pb.ModifySeatRequest) (*pb.ModifySeatResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	// Check if new seat is available
	if _, exists := s.seats[in.NewSeatNumber]; exists {
		return nil, fmt.Errorf("seat %d is already booked", in.NewSeatNumber)
	}

	userDetail, exists := s.users[in.Email]
	if !exists {
		return nil, fmt.Errorf("no user found with email %s", in.Email)
	}

	// Update seat number
	delete(s.seats, userDetail.SeatNumber)
	userDetail.SeatNumber = in.NewSeatNumber
	s.seats[in.NewSeatNumber] = in.Email

	return &pb.ModifySeatResponse{
		Message: "Seat modified successfully",
	}, nil
}

func (s *server) GetReceipt(ctx context.Context, req *pb.ReceiptRequest) (*pb.ReceiptResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	receipt, exists := s.receipts[req.ReceiptId]
	if !exists {
		return nil, fmt.Errorf("receipt with ID %s not found", req.ReceiptId)
	}

	// Convert your internal receipt structure to the protobuf structure if they are different
	pbReceipt := &pb.TicketReceipt{
		ReceiptId:  receipt.ReceiptId,
		From:       receipt.From,
		To:         receipt.To,
		Price:      receipt.Price,
		SeatNumber: receipt.SeatNumber,
		UserName:   receipt.UserName,
		UserEmail:  receipt.UserEmail,
	}

	return &pb.ReceiptResponse{Receipt: pbReceipt}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterTicketBookingServiceServer(s, &server{
		users: make(map[string]*pb.UserDetail),
		seats: make(map[int32]string),
	})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
