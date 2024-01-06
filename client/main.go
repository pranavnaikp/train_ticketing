package main

import (
    "context"
    "fmt"
    "log"
    "os"
    "time"

    "google.golang.org/grpc"
    pb "github.com/pranavnaikp/train_ticketing/proto/protos"
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
    log.Printf("Booking Response: %s", r.GetMessage())
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



