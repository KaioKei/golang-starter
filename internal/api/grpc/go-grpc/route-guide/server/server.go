package server

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/proto"
	pb "golang_starter/internal/api/grpc/go-grpc/route-guide"
	"golang_starter/internal/api/grpc/go-grpc/route-guide/backend"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
	"sync"
	"time"
)

// server implements all the methods in the protobuf file

// routeGuideServer server's structure
// Use this structure to store data to the server structure and manage it easily (eg. loaded data,
// etc ...)
type routeGuideServer struct {
	pb.UnimplementedRouteGuideServer
	savedFeatures []*pb.Feature // read-only after initialized

	// the Mutex will protect our routeNotes.
	// Mutex stands for 'mutual exclusion locks'.
	// Package 'sync' provides basic synchronization primitives such as mutual exclusion locks.
	// Other than the Once and WaitGroup types, most are intended for use by low-level library routines.
	// Higher-level synchronization is better done via channels and communication.
	//
	// A Mutex is a mutual exclusion lock. The zero value for a Mutex is an unlocked mutex.
	// A Mutex must not be copied after first use.
	// In the terminology of the Go memory model, the n'th call to Unlock “synchronizes before” the
	// m'th call to Lock for any n < m.
	// A successful call to TryLock is equivalent to a call to Lock.
	// A failed call to TryLock does not establish any “synchronizes before” relation at all.
	mu         sync.Mutex
	routeNotes map[string][]*pb.RouteNote
}

// GetFeature expects a Point and returns a unique feature from this Point
func (s *routeGuideServer) GetFeature(ctx context.Context, point *pb.Point) (*pb.Feature, error) {
	log.Println("Received GetFeature message for point:", point)
	for _, feature := range backend.Features {
		if feature.Location.Latitude == point.Latitude && feature.Location.Longitude == point.Longitude {
			// return the feature AND a nil error to tell gROC that we have finished dealing with the
			// client
			return feature, nil
		}
	}
	// no feature found, return unnamed feature
	return &pb.Feature{Name: "Unknown", Location: point}, nil
}

// ListFeatures expects a Rectangle and returns a stream of Features
// pb.RouteGuide_ListFeaturesServer is the compiled interface by protoc for grpc stream, from your
// route definition in the protobuf file. We need to use this interface to handle the stream from the
// server.
func (s *routeGuideServer) ListFeatures(rectangle *pb.Rectangle, stream pb.RouteGuide_ListFeaturesServer) error {
	log.Println("Received ListFeatures message for rectangle:", rectangle)
	for _, feature := range s.savedFeatures {
		// Use the backend function InRange to check if the feature's Point location is inside the
		// given rectangle
		if backend.InRange(feature.Location, rectangle) {
			// Sends back into the stream with the client, every feature inside the rectangle.
			// if stream fails, returns an error (if statement with short condition)
			if err := stream.Send(feature); err != nil {
				return err
			}
		}
	}
	// return a nil error to tell gRPC that we’ve finished writing responses
	return nil
}

// RecordRoute expects a stream of Point and returns a RouteSummary
// RouteGuide_RecordRouteServer is the compiled interface by protoc for grpc stream, from your
// route definition in the protobuf file. We need to use this interface to answer with stream from the
// client.
func (s *routeGuideServer) RecordRoute(stream pb.RouteGuide_RecordRouteServer) error {
	log.Println("Received RecordRoute stream connexion. Listen for route points ...")
	// init RouteSummary fields
	var pointCount, featureCount, distance int32
	// start computing route
	var lastPoint *pb.Point
	startTime := time.Now()
	// infinite loop
	// while the client does not close the stream
	// while the stream does not fail
	// while the last point has not been reached
	for {
		// The server receives client messages from the stream using its Recv()
		// The server needs to check the error returned from Recv() after each call.
		// If this is nil, the stream is still good and it can continue reading;
		// If it is io.EOF the message stream has ended and the server can return its RouteSummary.
		point, err := stream.Recv()
		if err == io.EOF {
			log.Println("The client has closed the stream.")
			endTime := time.Now()
			return stream.SendAndClose(&pb.RouteSummary{
				PointCount:   pointCount,
				FeatureCount: featureCount,
				Distance:     distance,
				ElapsedTime:  int32(endTime.Sub(startTime).Seconds()),
			})
		}
		// if the stream fails, return an error
		if err != nil {
			return err
		}
		// else, continue to count the points and the feature into the summary
		log.Println("Received new point:", point)
		pointCount++
		for _, feature := range s.savedFeatures {
			if proto.Equal(feature.Location, point) {
				featureCount++
			}
		}
		// the first time, the last point is nil because we need at least one point
		if lastPoint != nil {
			distance += backend.CalcDistance(lastPoint, point)
		}
		// update last point
		lastPoint = point
	}
}

// RouteChat expects a stream of RouteNote and returns a stream of RouteNote
// RouteGuide_RouteChatServer is the compiled interface by protoc for grpc stream, from your
// route definition in the protobuf file. We need to use this interface to answer with stream.
// We need only one stream to handle both the client and the server streams.
func (s *routeGuideServer) RouteChat(streamClient pb.RouteGuide_RouteChatServer) error {
	log.Println("Received RouteChat stream connexion. Listen for locations ...")
	// initialize the route notes
	s.routeNotes = make(map[string][]*pb.RouteNote)
	// infinite loop for the chat
	for {
		in, err := streamClient.Recv()
		// client has closed the communication
		if err == io.EOF {
			return nil
		}
		// stream failed
		if err != nil {
			return err
		}

		// serialize the location into a string text
		log.Println("Received new location:", in.Location)
		key := backend.Serialize(in.Location)

		// we will write the new location to a routeNote
		// protect the server's routeNotes with a mutex first
		s.mu.Lock()
		// add to the routeNotes the client input location, using the serialized location as a key
		s.routeNotes[key] = append(s.routeNotes[key], in)
		// Note: this copy prevents blocking other clients while serving this one.
		// We don't need to do a deep copy, because elements in the slice are
		// insert-only and never modified.
		rn := make([]*pb.RouteNote, len(s.routeNotes[key]))
		copy(rn, s.routeNotes[key])
		// unlocking the routeNotes for usage
		s.mu.Unlock()

		// look for notes to be sent to the client
		for _, note := range s.routeNotes[key] {
			// The server uses the stream’s Send() method rather than SendAndClose() because it’s
			// writing multiple responses.
			log.Println("Send note back to the client:", note)
			if err := streamClient.Send(note); err != nil {
				return err
			}
		}
	}

}

// newServer function is used to initialize the server and its data
func newServer() *routeGuideServer {
	// init empty server
	// maybe init the saved feature to an empty list in the server ?
	s := &routeGuideServer{}
	// load the backend's Features list into the server
	s.savedFeatures = backend.Features
	return s
}

func start(host string, port int) {
	listen, err := net.Listen("tcp", fmt.Sprintf("%s:%d", host, port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterRouteGuideServer(grpcServer, newServer())
	log.Printf("Listen on %s:%d", host, port)
	// Serve until the process is killed or Stop() is called
	grpcServer.Serve(listen)
}

func Run(host string, port int) {
	start(host, port)
}
