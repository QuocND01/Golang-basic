package auth

// Xử lý Google OAuth
import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var GoogleOAuthConfig = &oauth2.Config{
	ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
	ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
	RedirectURL:  os.Getenv("GOOGLE_REDIRECT_URL"),
	Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
	Endpoint:     google.Endpoint,
}

func GoogleLogin(c *gin.Context) {
	url := GoogleOAuthConfig.AuthCodeURL("randomstate")
	c.Redirect(http.StatusTemporaryRedirect, url)
}

func GoogleCallback(c *gin.Context) {
	code := c.Query("code")

	token, err := GoogleOAuthConfig.Exchange(context.Background(), code)
	if err != nil {
		log.Println("Error exchange token:", err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	client := GoogleOAuthConfig.Client(context.Background(), token)
	resp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil || resp.StatusCode != http.StatusOK {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	email := "user@example.com"

	// Tạo JWT
	jwtToken, err := GenerateJWT(email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate JWT"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": jwtToken})
}
