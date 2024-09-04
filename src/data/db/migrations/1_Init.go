package migrations

import (
	"CarSaleAd-Web-Api/config"
	"CarSaleAd-Web-Api/constants"
	"CarSaleAd-Web-Api/data/db"
	"CarSaleAd-Web-Api/data/models"
	"CarSaleAd-Web-Api/pkg/logging"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var logger = logging.NewLogger(config.GetConfig())

func Up_1() {
	database := db.GetDb()

	createTables(database)
	createDefaultUserInformation(database)
	createCountry(database)
	createPropertyCategory(database)

}

func createTables(database *gorm.DB) {

	tables := []interface{}{}

	// country := models.Country{}
	// city := models.City{}
	// user := models.User{}
	// role := models.Role{}
	// userRole := models.UserRole{}

	//Basic
	tables = addNewTable(database, models.Country{}, tables)
	tables = addNewTable(database, models.City{}, tables)
	tables = addNewTable(database, models.File{}, tables)
	tables = addNewTable(database, models.PersianYear{}, tables)
	//Property

	tables = addNewTable(database, models.PropertyCategory{}, tables)
	tables = addNewTable(database, models.Property{}, tables)

	//User
	tables = addNewTable(database, models.User{}, tables)
	tables = addNewTable(database, models.Role{}, tables)
	tables = addNewTable(database, models.UserRole{}, tables)

	//Car
	tables = addNewTable(database, models.Company{}, tables)
	tables = addNewTable(database, models.Gearbox{}, tables)
	tables = addNewTable(database, models.Color{}, tables)
	tables = addNewTable(database, models.CarType{}, tables)

	tables = addNewTable(database, models.CarModel{}, tables)
	tables = addNewTable(database, models.CarModelColor{}, tables)
	tables = addNewTable(database, models.CarModelYear{}, tables)
	tables = addNewTable(database, models.CarModelImage{}, tables)
	tables = addNewTable(database, models.CarModelPriceHistory{}, tables)
	tables = addNewTable(database, models.CarModelProperty{}, tables)
	tables = addNewTable(database, models.CarModelComment{}, tables)

	err := database.Migrator().CreateTable(tables...)
	if err != nil {
		logger.Error(logging.Postgres, logging.Migration, "tables Not Created", nil)
	}
	logger.Info(logging.Postgres, logging.Migration, "tables created", nil)
}

func addNewTable(database *gorm.DB, model interface{}, tables []interface{}) []interface{} {
	if !database.Migrator().HasTable(model) {
		tables = append(tables, model)
	}
	return tables
}

func createDefaultUserInformation(database *gorm.DB) {

	adminRole := models.Role{Name: constants.AdminRoleName}
	createRoleIfNotExists(database, &adminRole)

	defaultRole := models.Role{Name: constants.DefaultRoleName}
	createRoleIfNotExists(database, &defaultRole)

	u := models.User{Username: constants.DefaultUserName, FirstName: "Test", LastName: "Test",
		MobileNumber: "09156207562", Email: "admin@admin.com"}
	pass := "12345678"
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	u.Password = string(hashedPassword)

	createAdminUserIfNotExists(database, &u, adminRole.Id)

}

func createRoleIfNotExists(database *gorm.DB, r *models.Role) {
	exists := 0
	database.
		Model(&models.Role{}).
		Select("1").
		Where("name = ?", r.Name).
		First(&exists)

	if exists == 0 {

		database.Create(r)
	}
}

func createAdminUserIfNotExists(database *gorm.DB, u *models.User, roleId int) {
	exists := 0
	database.
		Model(&models.User{}).
		Select("1").
		Where("username = ?", u.Username).
		First(&exists)

	if exists == 0 {
		database.Create(u)
		ur := models.UserRole{UserId: u.Id, RoleId: roleId}
		database.Create(&ur)
	}
}

func createCountry(database *gorm.DB) {
	count := 0
	database.Model(&models.Country{}).
		Select("count(*)").
		Find(&count)
	if count == 0 {
		database.Create(&models.Country{Name: "Iran", Cities: []models.City{
			{Name: "Tehran"},
			{Name: "Isfahan"},
			{Name: "Shiraz"},
			{Name: "Chalus"},
			{Name: "Ahwaz"},
		}})
		database.Create(&models.Country{Name: "USA", Cities: []models.City{
			{Name: "New York"},
			{Name: "Washington"},
		}})
		database.Create(&models.Country{Name: "Germany", Cities: []models.City{
			{Name: "Berlin"},
			{Name: "Munich"},
		}})
		database.Create(&models.Country{Name: "China", Cities: []models.City{
			{Name: "Beijing"},
			{Name: "Shanghai"},
		}})
		database.Create(&models.Country{Name: "Italy", Cities: []models.City{
			{Name: "Roma"},
			{Name: "Turin"},
		}})
		database.Create(&models.Country{Name: "France", Cities: []models.City{
			{Name: "Paris"},
			{Name: "Lyon"},
		}})
		database.Create(&models.Country{Name: "Japan", Cities: []models.City{
			{Name: "Tokyo"},
			{Name: "Kyoto"},
		}})
		database.Create(&models.Country{Name: "South Korea", Cities: []models.City{
			{Name: "Seoul"},
			{Name: "Ulsan"},
		}})

	}
}

func createPropertyCategory(database *gorm.DB) {
	count := 0
	database.
		Model(&models.PropertyCategory{}).
		Select("count(*)").
		Find(&count)

	if count == 0 {
		database.Create(&models.PropertyCategory{Name: "Body"})
		database.Create(&models.PropertyCategory{Name: "Engine"})
		database.Create(&models.PropertyCategory{Name: "Drivetrain"})
		database.Create(&models.PropertyCategory{Name: "Suspension"})
		database.Create(&models.PropertyCategory{Name: "Equipment"})
		database.Create(&models.PropertyCategory{Name: "Driver support systems"})
		database.Create(&models.PropertyCategory{Name: "Lights"})
		database.Create(&models.PropertyCategory{Name: "Multimedia"})
		database.Create(&models.PropertyCategory{Name: "Safety equipment"})
		database.Create(&models.PropertyCategory{Name: "Seats and steering wheel"})
		database.Create(&models.PropertyCategory{Name: "Windows & mirrors"})

	}
	createProperty(database, "Body")
	createProperty(database, "Engine")
	createProperty(database, "Drivetrain")
	createProperty(database, "Suspension")
	createProperty(database, "Equipment")
	createProperty(database, "Driver support systems")
	createProperty(database, "Lights")
	createProperty(database, "Multimedia")
	createProperty(database, "SafetyEquipment")
	createProperty(database, "Seats and steering wheel")
	createProperty(database, "Windows and mirrors")

}

func createProperty(database *gorm.DB, cat string) {
	count := 0
	catModel := models.PropertyCategory{}

	database.Model(models.PropertyCategory{}).
		Where("name = ?", cat).
		Find(&catModel)

	database.Model(&models.Property{}).
		Select("count(*)").
		Where("category_id = ?", catModel.Id).
		Find(&count)

	if count > 0 || catModel.Id == 0 {
		return
	}

	var props *[]models.Property

	switch cat {
	case "Body":
		props = getBodyProperties(catModel.Id)

	case "Engine":
		props = getEngineProperties(catModel.Id)
	case "Drivetrain":
		props = getDrivetrainProperties(catModel.Id)
	case "Suspension":
		props = getSuspensionProperties(catModel.Id)
	case "Equipment":
		props = getEquipmentProperties(catModel.Id)
	case "Driver support systems":
		props = getDriverSupportSystemsProperties(catModel.Id)
	case "Lights":
		props = getLightsProperties(catModel.Id)
	case "Multimedia":
		props = getMultimediaProperties(catModel.Id)
	case "SafetyEquipment":
		props = getSafetyEquipmentProperties(catModel.Id)
	case "Seats and steering wheel":
		props = getSeatsAndSteeringWheelProperties(catModel.Id)
	case "Windows and mirrors":
		props = getWindowsAndMirrorsProperties(catModel.Id)

	default:
		props = &([]models.Property{})

	}
	for _, prop := range *props {
		database.Create(&prop)

	}

}

func Down_1() {

}
