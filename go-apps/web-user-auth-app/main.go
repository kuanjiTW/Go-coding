package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/glebarez/sqlite"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)


// ========== Data Structures (Model) ==========

// User represents the user data structure in the database
type User struct {
    // Embed the struct gorm.Model.
    //
    // Reference: https://pkg.go.dev/gorm.io/gorm#Model
    // type Model struct {
    //   ID        uint `gorm:"primarykey"`
    //   CreatedAt time.Time
    //   UpdatedAt time.Time
    //   DeletedAt DeletedAt `gorm:"index"`
    // }
	gorm.Model

    // A field declaration may be followed by an optional
    // string literal tag, which becomes an attribute for
    // all the fields in the corresponding field declaration.
    //
    // Field tags of Gorm (https://gorm.io/docs/models.html#Fields-Tags):
    // - uniqueIndex: the index is unique
    // - not null: specifies column as NOT NULL (a value is required)

    // Field tags of JSON (https://pkg.go.dev/encoding/json):
    // - json: The encoding of each struct field can be customized
    //         by the format string stored under the "json" key in
    //         the struct field's tag. The format string gives the
    //         name of the field, possibly followed by a comma-separated
    //         list of options.
	Email    string `gorm:"uniqueIndex;not null" json:"email"`

    // json: As a special case, if the field tag is "-", the field
    //       is always omitted.
	Password string `gorm:"not null" json:"-"`

	FullName string `json:"full_name"`

	Age      int    `json:"age"`
}

// RegisterInput represents registration data
type RegisterInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
	FullName string `json:"full_name" binding:"required"`
	Age      int    `json:"age"`
}

// LoginInput represents login data
type LoginInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}


// ========== Database ==========

var DB *gorm.DB
var jwtSecret = []byte("lsiduy!TIB&qQLoidxmwzq@il3je[o;il2uweo1i2ueh")

func ConnectDatabase() {
	db, err := gorm.Open(sqlite.Open("users.db"), &gorm.Config{})
	if err != nil { panic("Failed to connect to database!") }
    // AutoMigrate: Automatically migrate your schema, to keep your
    //              schema up to date.
    //              (https://gorm.io/docs/migration.html)
	db.AutoMigrate(&User{})
	DB = db
}


// ========== API Handlers (Control) ==========

// Context is the most important part of gin. It allows us to pass
// variables between middleware, manage the flow, validate the JSON
// of a request and render a JSON response for example.
func Register(c *gin.Context) {
	var input RegisterInput

    // If there is a binding error, the error is returned and it
    // is the developer’s responsibility to handle the request
    // and error appropriately.
	if err := c.ShouldBindJSON(&input); err != nil {
        // H is a shortcut for map[string]any
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

    // Make the hash of the password for security reason
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)

    // Create the corresponding User struct
	user := User{Email: input.Email, Password: string(hashedPassword), FullName: input.FullName, Age: input.Age}

    // Store the user data to the database
	if result := DB.Create(&user); result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User already exists"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Registration successful"})
}

func Login(c *gin.Context) {
	var input LoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var user User

    // Find the user with the specified email
	if err := DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

    // Check if the hash of the input password matches
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

    // Create JWT (JSON Web Token)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})
    // Sign the token
	tokenString, _ := token.SignedString(jwtSecret)

	// Set Cookie
	c.SetCookie("Authorization", tokenString, 3600*24, "/", "", false, true)
	c.JSON(http.StatusOK, gin.H{"message": "Login successful"})
}

func GetProfileAPI(c *gin.Context) {
	user, _ := c.Get("currentUser")
	c.JSON(http.StatusOK, gin.H{"profile": user})
}

// A middleware for authentication
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
        // Find the Authorization cookie
		tokenString, err := c.Cookie("Authorization")
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization cookie missing"})
			return
		}

        // Check if the token is valid
        //
        // Reference: https://pkg.go.dev/github.com/golang-jwt/jwt/v5#Parse
        // - func Parse(tokenString string, keyFunc Keyfunc, options ...ParserOption) (*Token, error)
        //   + Parse parses, validates, verifies the signature and returns the
        //     parsed token. keyFunc will receive the parsed token and should
        //     return the cryptographic key for verifying the signature.
		token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		})

        // Raise an error if the token is not valid
		if token == nil || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			return
		}
		claims, _ := token.Claims.(jwt.MapClaims)
		var user User
		DB.First(&user, uint(claims["sub"].(float64)))
		c.Set("currentUser", user)
		c.Next()
	}
}


// ========== Webpages (View) ==========

func ShowRegisterPage(c *gin.Context) {
	c.HTML(http.StatusOK, "register.html", nil)
}

func ShowLoginPage(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", nil)
}

func ShowProfilePage(c *gin.Context) {
	c.HTML(http.StatusOK, "profile.html", nil)
}


// ========== Main ==========

func main() {
    // Connect to the database
	ConnectDatabase()

    // Create an Engine instance with the Logger and Recovery
    // middleware already attached.
	r := gin.Default()

	// Load Templates
	r.LoadHTMLGlob("templates/*")

	// Router
	r.GET("/register", ShowRegisterPage)
	r.GET("/login", ShowLoginPage)
	r.GET("/profile", ShowProfilePage)

	// API Routes (JSON)
	r.POST("/register", Register)
	r.POST("/login", Login)

	// Protected Data Route
	protected := r.Group("/")
	protected.Use(AuthMiddleware())
	{
		protected.GET("/me", GetProfileAPI)
	}

    // Run on port 8080
	fmt.Println("Server running on http://localhost:8080/login")
	r.Run(":8080")
}
