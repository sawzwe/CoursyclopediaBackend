package route

import (
	"BackendCoursyclopedia/db"
	"BackendCoursyclopedia/handler/auditloghandler"
	"BackendCoursyclopedia/handler/facultyhandler"
	"BackendCoursyclopedia/handler/majorhandler"
	"BackendCoursyclopedia/handler/subjecthandler"
	"BackendCoursyclopedia/handler/userhandler"
	"BackendCoursyclopedia/repository/facultyrepository"
	"BackendCoursyclopedia/repository/majorrepository"
	"BackendCoursyclopedia/repository/subjectrepository"
	userrepo "BackendCoursyclopedia/repository/userrepository"
	"BackendCoursyclopedia/service/facultyservice"
	"BackendCoursyclopedia/service/majorservice"
	"BackendCoursyclopedia/service/subjectservice"

	auditlogrepo "BackendCoursyclopedia/repository/auditlogrepository"

	auditlogsvc "BackendCoursyclopedia/service/auditlogservice"
	usersvc "BackendCoursyclopedia/service/userservice"

	"github.com/gofiber/fiber/v2"
)

// func Setup(app *fiber.App) {
// 	db.ConnectDB()

// 	userRepository := userrepo.NewUserRepository(db.DB)
// 	majorRepository := majorrepository.NewMajorRepository(db.DB)
// 	facultyRepository := facultyrepository.NewFacultyRepository(db.DB)
// 	auditlogRepository := auditlogrepo.NewAuditLogRepository(db.DB)

// 	userService := usersvc.NewUserService(userRepository)
// 	facultyService := facultyservice.NewFacultyService(facultyRepository)
// 	majorService := majorservice.NewMajorService(majorRepository, facultyRepository)
// 	auditlogService := auditlogsvc.NewAuditLogService(auditlogRepository)

// 	userHandler := userhandler.NewUserHandler(userService)
// 	facultyHandler := facultyhandler.NewFacultyHandler(facultyService)
// 	majorHandler := majorhandler.NewMajorHandler(majorService)
// 	auditlogHandler := auditloghandler.NewAuditLogHandler(auditlogService)

// 	userGroup := app.Group("/api/users")
// 	userGroup.Get("/getallusers", userHandler.GetUsers)
// 	userGroup.Get("/getoneuser/:id", userHandler.GetOneUser)
// 	userGroup.Post("/createoneuser", userHandler.CreateOneUser)
// 	userGroup.Delete("/deleteoneuser/:id", userHandler.DeleteOneUser)
// 	userGroup.Put("/updateoneuser/:id", userHandler.UpdateOneUser)

// 	faculyGroup := app.Group("/api/faculties")
// 	faculyGroup.Get("getallfaculties", facultyHandler.GetFaculties)
// 	faculyGroup.Post("createfaculty", facultyHandler.CreateFaculty)
// 	faculyGroup.Put("updatefaculty/:id", facultyHandler.UpdateFaculty)
// 	faculyGroup.Delete("deletefaculty/:id", facultyHandler.DeleteFaculty)

// 	majorGroup := app.Group("api/majors")
// 	majorGroup.Post("createmajor", majorHandler.CreateMajor)
// 	majorGroup.Delete("deletemajor/:id", majorHandler.DeleteMajor)
// 	majorGroup.Put("updatemajor/:id", majorHandler.UpdateMajor)

// 	auditlogGroup := app.Group("/api/auditlogs")
// 	auditlogGroup.Get("/getallauditlogs", auditlogHandler.GetAuditLogs)

// }

func Setup(app *fiber.App) {
	// Connect to the database
	db.ConnectDB()

	// Initialize repositories
	userRepository := userrepo.NewUserRepository(db.DB)
	majorRepository := majorrepository.NewMajorRepository(db.DB)
	facultyRepository := facultyrepository.NewFacultyRepository(db.DB)
	auditlogRepository := auditlogrepo.NewAuditLogRepository(db.DB)
	subjectRepository := subjectrepository.NewSubjectRepository(db.DB)

	// Initialize services
	userService := usersvc.NewUserService(userRepository)
	facultyService := facultyservice.NewFacultyService(facultyRepository)
	majorService := majorservice.NewMajorService(majorRepository, facultyRepository)
	auditlogService := auditlogsvc.NewAuditLogService(auditlogRepository)
	subjectService := subjectservice.NewSubjectService(subjectRepository, majorRepository)

	// Initialize handlers
	userHandler := userhandler.NewUserHandler(userService)
	facultyHandler := facultyhandler.NewFacultyHandler(facultyService)
	majorHandler := majorhandler.NewMajorHandler(majorService)
	auditlogHandler := auditloghandler.NewAuditLogHandler(auditlogService)
	subjectHandler := subjecthandler.NewSubjectHandler(subjectService)

	// Define routes
	// Root route
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Welcome to the API")
	})

	// User routes
	userGroup := app.Group("/api/users")
	userGroup.Get("/getallusers", userHandler.GetUsers)
	userGroup.Get("/getoneuser/:id", userHandler.GetOneUser)
	userGroup.Post("/createoneuser", userHandler.CreateOneUser)
	userGroup.Delete("/deleteoneuser/:id", userHandler.DeleteOneUser)
	userGroup.Put("/updateoneuser/:id", userHandler.UpdateOneUser)


	faculyGroup := app.Group("/api/faculties")
	faculyGroup.Get("/getallfaculties", facultyHandler.GetFaculties)
	faculyGroup.Post("/createfaculty", facultyHandler.CreateFaculty)
	faculyGroup.Put("/updatefaculty/:id", facultyHandler.UpdateFaculty)
	faculyGroup.Delete("/deletefaculty/:id", facultyHandler.DeleteFaculty)

	majorGroup := app.Group("api/majors")
	majorGroup.Get("/getallmajors", majorHandler.GetMajors)

	majorGroup.Post("/createmajor", majorHandler.CreateMajor)
	majorGroup.Delete("/deletemajor/:id", majorHandler.DeleteMajor)
	majorGroup.Put("/updatemajor/:id", majorHandler.UpdateMajor)

	// Audit log routes
	auditlogGroup := app.Group("/api/auditlogs")
	auditlogGroup.Get("/getallauditlogs", auditlogHandler.GetAuditLogs)


	subjectGroup := app.Group("api/subjects")
	subjectGroup.Get("/getallsubjects", subjectHandler.GetSubjects)
	subjectGroup.Post("/createsubject", subjectHandler.CreateSubject)
	subjectGroup.Delete("/deletesubject/:id", subjectHandler.DeleteSubject)
	subjectGroup.Put("/updatesubject/:id", subjectHandler.UpdateSubject)

}
