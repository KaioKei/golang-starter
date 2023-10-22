package client

import (
	"context"
	pb "golang_starter/internal/api/grpc/go-grpc/route-guide"
	"golang_starter/internal/api/grpc/go-grpc/route-guide/backend"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"io"
	"log"
	"math/rand"
	"time"
)

func FatalIfError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func randomPoint(r *rand.Rand) *pb.Point {
	lat := (r.Int31n(180) - 90) * 1e7
	long := (r.Int31n(360) - 180) * 1e7
	return &pb.Point{Latitude: lat, Longitude: long}
}

func createRandomPoints() []*pb.Point {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	pointCount := int(r.Int31n(100)) + 2 // Traverse at least two points
	var points []*pb.Point
	for i := 0; i < pointCount; i++ {
		points = append(points, randomPoint(r))
	}
	return points
}

// getFeature simply calls the simple RPC method GetFeature
// pass a context.Context object which lets us change our RPC’s behavior if necessary, such as
// time-out/cancel an RPC in flight.
func sendGetFeature(client pb.RouteGuideClient, ctx context.Context, point *pb.Point) (*pb.Feature, error) {
	ft, err := client.GetFeature(ctx, point)
	return ft, err
}

// sendListFeatures sends a rectangle and receives a stream of points inside this rectangle
func sendListFeatures(client pb.RouteGuideClient, ctx context.Context, rectangle *pb.Rectangle) []*pb.Feature {
	// init the result slice of features
	var features []*pb.Feature
	// open the stream
	stream, err := client.ListFeatures(ctx, rectangle)
	FatalIfError(err)

	// compute the stream for each time clock
	for {
		streamFeature, err := stream.Recv()
		// if error returns EOF, simply close the stream
		if err == io.EOF {
			break
		}
		// any other error should fail the code
		FatalIfError(err)
		// extract the feature from the proto message
		features = append(features, streamFeature)
	}
	err = stream.CloseSend()
	FatalIfError(err)
	return features
}

// sendRecordRoute sends a stream of points and the server registers them
func sendRecordRoute(client pb.RouteGuideClient, ctx context.Context, points []*pb.Point) *pb.RouteSummary {
	stream, err := client.RecordRoute(ctx)
	FatalIfError(err)

	for _, point := range points {
		if err := stream.Send(point); err != nil {
			log.Fatalf("%v.Send(%v) = %v", stream, point, err)
		}
	}
	// This is a client stream, so we get the response by ending the stream from the client side
	// Since the server is waiting for the client to finish sending requests, it returns its response at
	//the end of the communication.
	reply, err := stream.CloseAndRecv()
	FatalIfError(err)

	return reply
}

func sendRouteChat(client pb.RouteGuideClient, ctx context.Context) {
	stream, err := client.RouteChat(ctx)
	FatalIfError(err)

	// channel to wait for incoming values
	waitChannel := make(chan struct{})

	// go routine to listen to incoming data
	go func() {
		for {
			in, err := stream.Recv()
			if err == io.EOF {
				close(waitChannel)
				return
			}
			FatalIfError(err)
			log.Printf("Point: (%d, %d) has a message: '%s'", in.Location.Latitude,
				in.Location.Latitude, in.Message)
		}
	}()

	for _, note := range backend.SavedNotes {
		err := stream.Send(note)
		FatalIfError(err)
	}
	stream.CloseSend()

	// blocking code execution until waitChannel receives a data or is closed
	<-waitChannel
}

func start(serverAddr string, method string) {
	// You can use DialOptions to set the auth credentials (for example, TLS, GCE credentials, or JWT
	// credentials) in grpc.Dial when a service requires them.
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	// Init the connection
	connection, err := grpc.Dial(serverAddr, opts...)
	FatalIfError(err)

	// Will close the connection after the 'main' function has ended
	defer connection.Close()

	// Init the client using our definition from the protobuf and the initialised connection
	client := pb.NewRouteGuideClient(connection)
	log.Println("Open client to", serverAddr)

	switch {
	case method == "getfeature":

		// Eiffel tower point
		point := &pb.Point{Latitude: 48858370, Longitude: 2294481}
		log.Println("Get feature for point:", point)
		feature, err := sendGetFeature(client, context.Background(), point)
		FatalIfError(err)
		log.Println(feature)
	case method == "listfeatures":
		log.Println("Send listFeatures message")
		// Rectangle points that should contain effeil tower point
		p1 := &pb.Point{Latitude: 48860806, Longitude: 2290437}
		p2 := &pb.Point{Latitude: 48855989, Longitude: 2297761}
		r := &pb.Rectangle{Lo: p1, Hi: p2}
		features := sendListFeatures(client, context.Background(), r)
		log.Printf("Features in Rectangle '%v' are: '%v'", r, features)
	case method == "sendrecordroute":
		// create a random number of random points of 10 points
		routePoints := createRandomPoints()
		log.Printf("Traversing %d points.", len(routePoints))
		route := sendRecordRoute(client, context.Background(), routePoints)
		log.Printf("Recorded route: {points: %d, distance: %dm, features: %d}",
			route.PointCount, route.Distance, route.FeatureCount)
	case method == "routechat":
		sendRouteChat(client, context.Background())
	case method == "debug":
		log.Println("It works")
	default:
		log.Fatal(&backend.UnknownGrpcMethodError{Name: "getfeature"})
	}

}

func Run(serverAddr string, method string) {
	start(serverAddr, method)
}
