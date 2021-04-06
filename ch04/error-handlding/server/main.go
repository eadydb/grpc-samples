package main

import (
	"context"
	"fmt"
	pb "github.com/eadydb/grpc-samples/ch04/error-handlding/proto"
	"github.com/golang/protobuf/ptypes/wrappers"
	epb "google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io"
	"log"
	"net"
	"strings"
	"sync"
)

const (
	port           = ":50051"
	orderBatchSize = 3
)

type server struct {
	orderMap map[string]*pb.Order
	mu       sync.Mutex
	pb.UnimplementedOrderManagementServer
}

func (s *server) AddOrder(ctx context.Context, orderReq *pb.Order) (*wrappers.StringValue, error) {

	if orderReq.Id == "-1" {
		log.Printf("Order ID is invalid! -> Received Order ID %s", orderReq.Id)
		errorStatus := status.New(codes.InvalidArgument, "Invalid information received")
		ds, err := errorStatus.WithDetails(
			&epb.BadRequest_FieldViolation{
				Field:       "ID",
				Description: fmt.Sprintf("Order ID received is not valid %s : %s", orderReq.Id, orderReq.Description),
			})
		if err != nil {
			return nil, errorStatus.Err()
		}
		return nil, ds.Err()
	} else {
		s.orderMap[orderReq.Id] = orderReq
		log.Println("Order : ", orderReq.Id, " -> Added")
		return &wrappers.StringValue{Value: "Order Added: " + orderReq.Id}, nil
	}
}

func (s *server) GetOrder(_ context.Context, orderId *wrappers.StringValue) (*pb.Order, error) {
	ord, exists := s.orderMap[orderId.Value]
	if exists {
		return ord, status.New(codes.OK, "").Err()
	}
	return nil, status.Errorf(codes.NotFound, "order dos not exist. :", orderId)
}

// Server-side Streaming RPC
func (s *server) SearchOrders(searchQuery *wrappers.StringValue, stream pb.OrderManagement_SearchOrdersServer) error {
	for key, order := range s.orderMap {
		log.Print(key, order)
		for _, itemStr := range order.Items {
			log.Print(itemStr)
			if strings.Contains(itemStr, searchQuery.Value) {
				// send the matching order in a stream
				err := stream.Send(order)
				if err != nil {
					return fmt.Errorf("error sending message to stream : %v", err)
				}
				log.Printf("Matching Order Found: , %s", key)
				break
			}
		}
	}
	return nil
}

// Client-side streaming RPC
func (s *server) UpdateOrders(stream pb.OrderManagement_UpdateOrdersServer) error {
	orderStr := "Updated Order IDs: "
	for {
		order, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&wrappers.StringValue{Value: "Orders processed " + orderStr})
		}

		if err != nil {
			return err
		}
		s.orderMap[order.Id] = order
		log.Printf("Order ID : %s - %s", order.Id, "Updated")
		orderStr += order.Id + ", "
	}
}

// Bi-Directional Streaming RPC
func (s *server) ProcessOrders(stream pb.OrderManagement_ProcessOrdersServer) error {
	batchMarker := 1
	var combinedShipmentMap = make(map[string]*pb.CombinedShipment)
	for {
		orderId, err := stream.Recv()
		log.Printf("Reading Proc order: %s", orderId)
		if err == io.EOF {
			log.Printf("EOF: %s", orderId)
			for _, shipment := range combinedShipmentMap {
				if err := stream.Send(shipment); err != nil {
					return err
				}
			}
		}

		if err != nil {
			log.Println(err)
			return err
		}

		destination := s.orderMap[orderId.GetValue()].Destination
		shipment, found := combinedShipmentMap[destination]

		if found {
			ord := s.orderMap[orderId.GetValue()]
			shipment.OrderList = append(shipment.OrderList, ord)
		} else {
			shipment := &pb.CombinedShipment{}
			comShip := &pb.CombinedShipment{Id: "cmb-" + (s.orderMap[orderId.GetValue()].Destination), Status: "Processed!"}
			ord := s.orderMap[orderId.GetValue()]
			comShip.OrderList = append(shipment.OrderList, ord)
			combinedShipmentMap[destination] = comShip
			log.Print(len(comShip.OrderList), comShip.GetId())
		}

		if batchMarker == orderBatchSize {
			for _, comb := range combinedShipmentMap {
				log.Printf("Shipping : %v -> %v", comb.Id, len(comb.OrderList))
				if err := stream.Send(comb); err != nil {
					return err
				}
			}
			batchMarker = 0
			combinedShipmentMap = make(map[string]*pb.CombinedShipment)
		} else {
			batchMarker++
		}

	}
}

// 初始化数据
func (s *server) initSampleData() {
	order := make(map[string]*pb.Order)
	order["102"] = &pb.Order{Id: "102", Items: []string{"Google Pixel 3A", "Mac Book Pro"}, Destination: "Mountain View, CA",
		Price: 1800.00}
	order["103"] = &pb.Order{Id: "103", Items: []string{"Apple Watch S4"}, Destination: "San Jose, CA", Price: 400.00}
	order["104"] = &pb.Order{Id: "104", Items: []string{"Google Home Mini", "Google Nest Hub"}, Destination: "Mountain View, CA",
		Price: 400.00}
	order["105"] = &pb.Order{Id: "105", Items: []string{"Amazon Echo"}, Destination: "San Jose, CA", Price: 30.00}
	order["106"] = &pb.Order{Id: "106", Items: []string{"Amazon Echo", "Apple iPhone XS"}, Destination: "Mountain View, CA", Price: 300.00}

	s.orderMap = order
}

func main() {
	ser := &server{}
	ser.initSampleData()

	list, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	pb.RegisterOrderManagementServer(s, ser)

	log.Printf("Starting gRPC listener on port %s", port)

	if err := s.Serve(list); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
