package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	pb "github.com/mohamedelbalshy/demo-grpc/invoicer"
	"google.golang.org/grpc"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial("localhost:8001", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewInvoicerClient(conn)

	// Set up a http setver.
	r := gin.Default()
	r.GET("/rest/n/:name", func(c *gin.Context) {
		name := c.Param("name")

		// Contact the server and print out its response.
		req := &pb.CreateRequest{Amount: &pb.Amount{Amount: 10, Currency: "USD"}, From: "Mohamed", To: "Maya"}
		res, err := client.Create(c, req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"result": string(res.Docx[:]),
			"name":   name,
		})
	})

	// Run http server
	if err := r.Run(":8052"); err != nil {
		log.Fatalf("could not run server: %v", err)
	}
}
