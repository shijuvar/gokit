package routers

import (
	"github.com/gorilla/mux"

	util "github.com/shijuvar/gokit/examples/bookmark-api/apputil"
	"github.com/shijuvar/gokit/examples/bookmark-api/controllers"
)

// SetBookmarkRoutes registers routes for bookmark entity.
func SetBookmarkRoutes(router *mux.Router) *mux.Router {
	bookmarkRouter := mux.NewRouter()
	bookmarkRouter.HandleFunc("/bookmarks", controllers.CreateBookmark).Methods("POST")
	bookmarkRouter.HandleFunc("/bookmarks/{id}", controllers.UpdateBookmark).Methods("PUT")
	bookmarkRouter.HandleFunc("/bookmarks", controllers.GetBookmarks).Methods("GET")
	bookmarkRouter.HandleFunc("/bookmarks/{id}", controllers.GetBookmarkByID).Methods("GET")
	bookmarkRouter.HandleFunc("/bookmarks/users/{id}", controllers.GetBookmarksByUser).Methods("GET")
	bookmarkRouter.HandleFunc("/bookmarks/{id}", controllers.DeleteBookmark).Methods("DELETE")
	router.PathPrefix("/bookmarks").Handler(util.AuthorizeRequest(bookmarkRouter))
	return router
}
