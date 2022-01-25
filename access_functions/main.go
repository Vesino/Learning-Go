package accessfunctions

// Now this struct is available, that's why we have to write a description
// CarPublic contains only the brand and the year atributes of a car instance
type CarPublic struct {
	Brand string
	Year  string
}

type car struct {
	brand string
	year  string
}
