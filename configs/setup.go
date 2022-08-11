package configs

import "github.com/gin-gonic/gin"

func ConnectDB() {

}

func RunServer(Router **gin.Engine) {
	(*Router).Run(EnvPORT())
}
