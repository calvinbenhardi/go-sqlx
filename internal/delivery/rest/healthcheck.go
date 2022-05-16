package rest

import (
	"net/http"

	"github.com/calvinbenhardi/go-sqlx/internal/tool"
)

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	tool.RenderJSON(w, http.StatusOK, tool.Map{"status": "Ok"})
}
