package handlers

type Handler struct {
	// DB *db.DB
}

func NewHandler() *Handler {
	// TODO: figure out db situation later
	// db, err := db.Init("")
	// if err != nil {
	// 	panic(err)
	// }

	// return &Handler{
	// 	DB: db,
	// }
	return &Handler{}
}
