package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	pb "github.com/pranavnaikp/train_ticketing/proto/protos"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func main() {

	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewTicketBookingServiceClient(conn)

	var choice int
	for {
		fmt.Println("\nSelect an action:")
		fmt.Println("1. Book Ticket")
		fmt.Println("2. Get User Detail")
		fmt.Println("3. Get Users By Section")
		fmt.Println("4. Remove User")
		fmt.Println("5. Modify User Seat")
        fmt.Println("6. Get Receipt")
		fmt.Println("0. Exit")
		fmt.Print("Enter choice: ")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			bookTicket(client)
		case 2:
			getUserDetail(client)
		case 3:
			getUsersBySection(client)
		case 4:
			removeUser(client)
		case 5:
			modifyUserSeat(client)
		case 6:
			getReceiptDetail(client)
		case 0:
			fmt.Println("Exiting...")
			os.Exit(0)
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}

func bookTicket(c pb.TicketBookingServiceClient) {
    var firstName, lastName, email string
    var seatNumber int

    fmt.Print("Enter first name: ")
    fmt.Scanln(&firstName)
    fmt.Print("Enter last name: ")
    fmt.Scanln(&lastName)
    fmt.Print("Enter email: ")
    fmt.Scanln(&email)
    fmt.Print("Enter seat number: ")
    fmt.Scanln(&seatNumber)

    ctx, cancel := context.WithTimeout(context.Background(), time.Second)
    defer cancel()

    r, err := c.BookTicket(ctx, &pb.BookTicketRequest{
        FirstName:  firstName,
        LastName:   lastName,
        Email:      email,
        SeatNumber: int32(seatNumber),
    })
    if err != nil {
        log.Fatalf("could not book ticket: %v", err)
    }

    // Displaying the booking details
    fmt.Println("\nBooking Successful!")
    fmt.Printf("Full Name: %s\n", r.GetFullName())
    fmt.Printf("Email: %s\n", r.GetEmail())
    fmt.Printf("From: %s\n", r.GetFrom())
    fmt.Printf("To: %s\n", r.GetTo())
    fmt.Printf("Receipt ID: %s\n", r.GetReceiptId())
    fmt.Printf("Price: $%.2f\n", r.GetPrice())
}


func getUserDetail(c pb.TicketBookingServiceClient) {
	var email string
	fmt.Print("Enter email: ")
	fmt.Scanln(&email)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.GetUserDetail(ctx, &pb.UserRequest{Email: email})
	if err != nil {
		log.Fatalf("could not get user detail: %v", err)
	}
	user := r.GetUser()
	fmt.Printf("User Details - Name: %s %s, Email: %s, Seat Number: %d\n", user.GetFirstName(), user.GetLastName(), user.GetEmail(), user.GetSeatNumber())
}

func getUsersBySection(c pb.TicketBookingServiceClient) {
	var section string
	fmt.Print("Enter section (A/B): ")
	fmt.Scanln(&section)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.GetUsersBySection(ctx, &pb.SectionRequest{Section: section})
	if err != nil {
		log.Fatalf("could not get users by section: %v", err)
	}
	for _, user := range r.GetUsers() {
		fmt.Printf("Name: %s %s, Email: %s, Seat Number: %d\n", user.GetFirstName(), user.GetLastName(), user.GetEmail(), user.GetSeatNumber())
	}
}

func removeUser(c pb.TicketBookingServiceClient) {
	var email string
	fmt.Print("Enter email of user to remove: ")
	fmt.Scanln(&email)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.RemoveUser(ctx, &pb.RemoveUserRequest{Email: email})
	if err != nil {
		log.Fatalf("could not remove user: %v", err)
	}
	fmt.Println(r.GetMessage())
}

func modifyUserSeat(c pb.TicketBookingServiceClient) {
	var email string
	var newSeatNumber int

	fmt.Print("Enter email of user: ")
	fmt.Scanln(&email)
	fmt.Print("Enter new seat number: ")
	fmt.Scanln(&newSeatNumber)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.ModifyUserSeat(ctx, &pb.ModifySeatRequest{
		Email:         email,
		NewSeatNumber: int32(newSeatNumber),
	})
	if err != nil {
		log.Fatalf("could not modify user seat: %v", err)
	}
	fmt.Println(r.GetMessage())
}

func getReceiptDetail(c pb.TicketBookingServiceClient) {
	var receiptID string
	fmt.Print("Enter receipt ID: ")
	fmt.Scanln(&receiptID)

	// Create a context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Call the GetReceipt method
	r, err := c.GetReceipt(ctx, &pb.ReceiptRequest{ReceiptId: receiptID})
	if err != nil {
		log.Fatalf("could not retrieve receipt: %v", err)
	}

	// Display the receipt details
	receipt := r.GetReceipt()
	fmt.Printf("Receipt Details:\n")
	fmt.Printf("Receipt ID: %s\n", receipt.GetReceiptId())
	fmt.Printf("From: %s\n", receipt.GetFrom())
	fmt.Printf("To: %s\n", receipt.GetTo())
	fmt.Printf("Price: $%.2f\n", receipt.GetPrice())
	fmt.Printf("Seat Number: %d\n", receipt.GetSeatNumber())
	fmt.Printf("User Name: %s\n", receipt.GetUserName())
	fmt.Printf("User Email: %s\n", receipt.GetUserEmail())
}
