package router

import (
	"github.com/gorilla/mux"

	util "github.com/shijuvar/gokit/examples/bookmark-api/apputil"
	"github.com/shijuvar/gokit/examples/bookmark-api/controller"
)

// SetBookmarkRoutes registers routes for bookmark entity.
func SetBookmarkRoutes(router *mux.Router) *mux.Router {
	bookmarkRouter := mux.NewRouter()
	bookmarkRouter.HandleFunc("/bookmarks", controller.CreateBookmark).Methods("POST")
	bookmarkRouter.HandleFunc("/bookmarks/{id}", controller.UpdateBookmark).Methods("PUT")
	bookmarkRouter.HandleFunc("/bookmarks", controller.GetBookmarks).Methods("GET")
	bookmarkRouter.HandleFunc("/bookmarks/{id}", controller.GetBookmarkByID).Methods("GET")
	bookmarkRouter.HandleFunc("/bookmarks/users/{id}", controller.GetBookmarksByUser).Methods("GET")
	bookmarkRouter.HandleFunc("/bookmarks/{id}", controller.DeleteBookmark).Methods("DELETE")
	router.PathPrefix("/bookmarks").Handler(util.AuthorizeRequest(bookmarkRouter))
	return router
}
