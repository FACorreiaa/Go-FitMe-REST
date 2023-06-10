package books

//type Config struct {
//	postgresConfig postgres.Config
//}
//
//func NewConfig(postgresConfig postgres.Config) Config {
//	return Config{postgresConfig: postgresConfig}
//}
//
//type Book struct {
//	db.DefaultModel
//	Title  string `json:"title"`
//	Author string `json:"author"`
//	Rating int    `json:"rating"`
//}
//
//func GetBooks(c *fiber.Ctx) error {
//	database := db.NewPostgres(Config)
//	var books []Book
//	database
//	return c.JSON(books)
//}
//
//func GetBook(c *fiber.Ctx) error {
//	id := c.Params("id")
//	database := db.DB
//	var book Book
//	err := database.First(&book, id).Error
//
//	if errors.Is(err, gorm.ErrRecordNotFound) {
//		return handlers.EntityNotFound("No book found")
//	} else if err != nil {
//		return handlers.Unexpected(err.Error())
//	}
//
//	return c.JSON(book)
//}
//
//func NewBook(c *fiber.Ctx) error {
//	database := db.DB
//	book := new(Book)
//	if err := c.BodyParser(book); err != nil {
//		return handlers.BadRequest("Invalid params")
//	}
//	database.Create(&book)
//	return c.JSON(book)
//}
//
//func UpdateBook(c *fiber.Ctx) error {
//	id := c.Params("id")
//	database := db.DB
//	var book Book
//	err := database.First(&book, id).Error
//
//	if errors.Is(err, gorm.ErrRecordNotFound) {
//		return handlers.EntityNotFound("No book found")
//	} else if err != nil {
//		return handlers.Unexpected(err.Error())
//	}
//
//	updatedBook := new(Book)
//
//	if err := c.BodyParser(updatedBook); err != nil {
//		return handlers.BadRequest("Invalid params")
//	}
//
//	updatedBook = &Book{Title: updatedBook.Title, Author: updatedBook.Author, Rating: updatedBook.Rating}
//
//	if err = database.Model(&book).Updates(updatedBook).Error; err != nil {
//		return handlers.Unexpected(err.Error())
//	}
//
//	return c.SendStatus(204)
//}
//
//func DeleteBook(c *fiber.Ctx) error {
//	id := c.Params("id")
//	database := db.DB
//
//	var book Book
//	err := database.First(&book, id).Error
//
//	if errors.Is(err, gorm.ErrRecordNotFound) {
//		return handlers.EntityNotFound("No book found")
//	} else if err != nil {
//		return handlers.Unexpected(err.Error())
//	}
//
//	database.Delete(&book)
//	return c.SendStatus(204)
//}
