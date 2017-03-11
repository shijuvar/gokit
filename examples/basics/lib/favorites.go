package lib

// Stores favorites
var favorites []string

// Initialization logic for the package
func init() {
	favorites = make([]string, 3)
	favorites[0] = "github.com/gorilla/mux"
	favorites[1] = "github.com/codegangsta/negroni"
	favorites[2] = "gopkg.in/mgo.v2"
}

// Add a favorite into the in-memory collection
func Add(favorite string) {
	favorites = append(favorites, favorite)
}

// Returns all favorite
func GetAll() []string {
	return favorites
}
