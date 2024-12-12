// edge service

package edge

import (
	"fmt"
	"strconv"
	"time"

	"github.com/maathias/capstone/db"

	"github.com/gin-gonic/gin"
)

func Run() {
	fmt.Println("Running edge service")

	r := gin.Default()

	r.POST("/location/update", func(c *gin.Context) {
		// * POST /location/update
		// * Content-Type: application/x-www-form-urlencoded
		// * x-uname: 123
		// *
		// * long=123.456&lat=123.456

		uname := c.GetHeader("x-uname")
		_long := c.PostForm("long")
		_lat := c.PostForm("lat")

		long, err := strconv.ParseFloat(_long, 64)
		if err != nil {
			c.JSON(400, gin.H{"error": "Invalid longitude"})
			return
		}

		lat, err := strconv.ParseFloat(_lat, 64)
		if err != nil {
			c.JSON(400, gin.H{"error": "Invalid latitude"})
			return
		}

		db.GeoAdd( "locations", long, lat, uname)

		fmt.Println(uname, long, lat)

		// TODO: send to api service
	})
	
	r.GET("/location/radius", func(c *gin.Context) {
		// * GET /location/radius?
		// * 	long=123.456
		// * 	&lat=123.456
		// * 	&radius=1000

		_long := c.Query("long")
		_lat := c.Query("lat")
		_radius := c.Query("radius")

		long, err := strconv.ParseFloat(_long, 64)
		if err != nil {
			c.JSON(400, gin.H{"error": "Invalid longitude"})
			return
		}

		lat, err := strconv.ParseFloat(_lat, 64)
		if err != nil {
			c.JSON(400, gin.H{"error": "Invalid latitude"})
			return
		}

		radius, err := strconv.ParseFloat(_radius, 64)
		if err != nil {
			c.JSON(400, gin.H{"error": "Invalid radius"})
			return
		}

		users := db.GeoRadius("locations", long, lat, radius)

		fmt.Println(users)

		// TODO: paginate
	})

	r.GET("/distance", func(c *gin.Context) {
		// * GET /distance?
		// * 	uname=123
		// * 	&start=2020-01-01T00:00:00Z
		// * 	&end=2020-01-02T00:00:00Z

		uname := c.Query("uname")
		_start := c.Query("start")
		_end := c.Query("end")

		start, err := time.Parse(time.RFC3339, _start)
		if err != nil {
			c.JSON(400, gin.H{"error": "Invalid start time"})
			return
		}

		end, err := time.Parse(time.RFC3339, _end)
		if err != nil {
			c.JSON(400, gin.H{"error": "Invalid end time"})
			return
		}

		fmt.Println(uname, start, end)

		// TODO: query api for user distance in given time range
	})

	r.Run(":8080")
}